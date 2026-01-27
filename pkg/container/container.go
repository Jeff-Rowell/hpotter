package container

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/netip"
	"strings"
	"sync"
	"time"

	"github.com/hpotter/pkg/dockerclient"
	"github.com/hpotter/pkg/oneway"
	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/network"
	"github.com/moby/moby/client"
)

// Start is a wrapper around the Docker API client. It ensures the given
// job can run successfully and does a series of checks. It will check if the
// container image associated with the job exists. Start creates the container
// along with two threads for writing requests and reading responses to and
// from the running container, respectively.
func (ct *ContainerThread) Start(job *ContainerJob) {
	container, err := ct.CreateContainer(job)
	if err != nil {
		job.ErrChan <- err
		return
	}

	err = ct.DockerClient.StartContainer(container)
	if err != nil {
		job.ErrChan <- err
		return
	}
	slog.Debug("started container", "job", job.Conn, "container", container)

	ct.DockerClient.Messages <- dockerclient.ActorMessage{
		Type:  dockerclient.Append,
		Value: container,
	}

	ip, err := ct.DockerClient.GetContainerIP(container)
	if err != nil {
		ct.handleError(job, container, err)
		return
	}

	containerConn, err := connectToContainer(job, ip)
	if err != nil {
		ct.handleError(job, container, err)
		return
	}

	var wg sync.WaitGroup
	errChan := make(chan error, 2)
	requestJob := oneway.OnewayJob{
		Source:      job.Conn,
		Destination: containerConn,
		Direction:   "request",
		Wg:          &wg,
		Ctx:         job.Ctx,
		ErrChan:     errChan,
	}
	responseJob := oneway.OnewayJob{
		Source:      containerConn,
		Destination: job.Conn,
		Direction:   "response",
		Wg:          &wg,
		Ctx:         job.Ctx,
		ErrChan:     errChan,
	}

	failed := false

	go func() {
		err := <-errChan
		if slog.Default().Enabled(context.Background(), slog.LevelDebug) {
			slog.Error("one-way thread error received", "error", err)
		}

		failed = true
		ct.writeData(containerConn, job, container)
		ct.DockerClient.Shutdown(container)
	}()

	wg.Add(1)
	go oneway.Thread(&requestJob)

	wg.Add(1)
	go oneway.Thread(&responseJob)

	wg.Wait()

	if !failed {
		ct.writeData(containerConn, job, container)
	}

	slog.Debug("one-way threads complete, container thread exiting")
}

// writeData attempts to orchestrate writing the recorded session data,
// connection information, and container/image information to the database.
func (ct *ContainerThread) writeData(cc net.Conn, j *ContainerJob, c string) {
	sd, err := j.Svc.Recorder.GetData(j.Ctx, ct.DockerClient.Client, c)
	if err != nil {
		j.ErrChan <- err
		slog.Error("session recorder received error", "error", err)
		return
	}

	j.Svc.Recorder.SaveSession(j.Conn, cc, sd, j.Db, j.Svc.Image, j.Started)
	slog.Debug("session data recorder", "sessionData", fmt.Sprintf("%+v", sd))
}

// handleError takes a job, container id, and error message and helps
// orchestrate a container cleanup when an error happens. The given error
// is written to the job's error channel to communicate upstream that it
// failed, then the Docker API client is called to gracefully shutdown the
// given container id.
func (d *ContainerThread) handleError(job *ContainerJob, c string, err error) {
	job.ErrChan <- err
	shutdownErr := d.DockerClient.Shutdown(c)
	if shutdownErr != nil {
		job.ErrChan <- shutdownErr
	}
}

// CreateContainer gets the host and container config and then calls
// the Docker API to create a container from j. CreateContainer returns
// an error if encountered, otherwise it returns the ID string from the
// container creation response.
func (ct *ContainerThread) CreateContainer(j *ContainerJob) (string, error) {
	hostConf, err := ct.getHostConf(j)
	if err != nil {
		return "", err
	}

	containerConf, err := ct.getContainerConf(j)
	if err != nil {
		return "", err
	}

	containerOpts := client.ContainerCreateOptions{
		Config:     containerConf,
		HostConfig: hostConf,
	}

	createResponse, err := ct.DockerClient.Client.ContainerCreate(ct.DockerClient.Ctx, containerOpts)
	if err != nil {
		return "", err
	}

	return createResponse.ID, nil
}

// connectToContainer attempts to create a connection to a container given
// the job specifying the container network details and the ip from calling
// [getContainerIP]. If the connection is successful it will be returned,
// otherwise nil will be retuned along with an error.
func connectToContainer(job *ContainerJob, ip string) (net.Conn, error) {
	var addr string
	if strings.Contains(ip, ":") {
		addr = fmt.Sprintf("[%s]:%s", ip, job.Svc.ContainerPort)
	} else {
		addr = fmt.Sprintf("%s:%s", ip, job.Svc.ContainerPort)
	}

	var err error
	var c net.Conn
	for range 10 {
		c, err = net.DialTimeout(job.Svc.ListenProto, addr, 5*time.Second)
		if err == nil {
			return c, nil
		}
		time.Sleep(2 * time.Second)
	}

	return nil, err
}

// getContainerConf uses the configuration settings in j and creates a
// container configuration that can be used to create a container using
// the Docker API. getContainerConf will return the container configuration
// if it was successfully created or an error if encountered.
func (ct *ContainerThread) getContainerConf(j *ContainerJob) (*container.Config, error) {
	p := fmt.Sprintf("%s/%s", j.Svc.ContainerPort, j.Svc.ListenProto)
	port, err := network.ParsePort(p)
	if err != nil {
		return nil, err
	}
	portSet := make(map[network.Port]struct{}, 1)
	portSet[port] = struct{}{}

	cConf := container.Config{
		Hostname:     "diagon-alley",
		Domainname:   "hogwarts",
		Image:        j.Svc.Image,
		ExposedPorts: portSet,
		Env:          j.Svc.Env,
	}
	return &cConf, nil
}

// getHostConf uses the configuration settings in j and creates a
// host configuration that can be used to create a container using
// the Docker API. getHostConf will return the host configuration
// if it was successfully created or an error if encountered.
func (ct *ContainerThread) getHostConf(j *ContainerJob) (*container.HostConfig, error) {
	var err error
	var listenOn netip.Addr

	if j.Svc.ListenAddress == "" {
		listenOn, err = netip.ParseAddr("127.0.0.1")
	} else {
		listenOn, err = netip.ParseAddr(j.Svc.ListenAddress)
	}

	if err != nil {
		return nil, err
	}

	p := fmt.Sprintf("%s/%s", j.Svc.ListenPort, j.Svc.ListenProto)
	port, err := network.ParsePort(p)
	if err != nil {
		return nil, err
	}

	portBinding := network.PortBinding{
		HostIP:   listenOn,
		HostPort: j.Svc.ListenPort,
	}

	portMap := make(map[network.Port][]network.PortBinding)
	portMap[port] = []network.PortBinding{portBinding}

	hostConfig := container.HostConfig{
		NetworkMode: "bridge",
		AutoRemove:  true,
		Privileged:  false,
	}
	return &hostConfig, nil
}
