package session

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net"
	"strconv"
	"strings"
	"time"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"github.com/hpotter/pkg/database"
	"github.com/moby/moby/client"
)

const BUFFER_SIZE = 4096

// GetData parses the container logs from c and returns a [session.SessionData]
// containing all of the username, password, and payload/command data found.
// If an error is encountered nil is returned along with a non-nil error,
// otherwise the populated SessionData is returned with a nil error.
func (r *Recorder) GetData(ctx context.Context, cl *client.Client, c string) (*SessionData, error) {
	logs, err := r.readLogs(ctx, cl, c)
	if err != nil {
		return nil, err
	}

	creds := []Credential{}
	payloads := []string{}
	for _, l := range logs {
		u := r.UsernamePattern.FindStringSubmatch(l)
		p := r.PasswordPattern.FindStringSubmatch(l)

		if len(u) > 1 && len(p) > 1 {
			c := Credential{
				Username: u[1],
				Password: p[1],
			}
			creds = append(creds, c)
		}

		payload := r.PayloadPattern.FindStringSubmatch(l)

		if len(payload) > 1 {
			payloads = append(payloads, payload[1])
		}
	}

	return &SessionData{Credentials: creds, Payloads: payloads}, nil
}

// readLogs reads logs from the given container c using the Docker API.
// readLogs returns a slice of strings containing the parsed logs and a nil
// error if successful, otherwise a nil slice is returned with a non-nil
// error in the case that an error is encountered.
func (r *Recorder) readLogs(ctx context.Context, cl *client.Client, c string) ([]string, error) {
	logOpts := client.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	}
	reader, err := cl.ContainerLogs(ctx, c, logOpts)
	if err != nil {
		return nil, err
	}

	cLogs := []string{}

	for {
		bytes := make([]byte, BUFFER_SIZE)
		bytesRead, err := reader.Read(bytes)
		lines := strings.Split(string(bytes[:bytesRead]), "\n")
		cLogs = append(cLogs, lines...)

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}
	}

	return cLogs, nil

}

func (r *Recorder) SaveSession(s, d net.Conn, sd *SessionData, db *database.Database, i string, t time.Time) {
	sourceIp, sourcePort := r.parseIp(s)
	destIp, destPort := r.parseIp(d)
	ipInfo, err := r.getGeoLocationData(sourceIp)
	if err != nil {
		slog.Error("failed to get ip info", "error", err)
		return
	}

	connection := database.Connection{
		ContainerImage:  i,
		SourceIpAddress: sourceIp,
		DestIpAddress:   destIp,
		SourcePort:      sourcePort,
		DestPort:        destPort,
		Duration:        time.Since(t),
		Latitude:        ipInfo.Latitude,
		Longitude:       ipInfo.Longitude,
		Country:         ipInfo.Country,
		Region:          ipInfo.Region,
		City:            ipInfo.City,
		Zipcode:         ipInfo.Zipcode,
	}

	connId, err := db.CreateConnection(&connection)
	if err != nil {
		slog.Error("failed to add connection to db", "error", err)
		return
	}

	for _, credential := range sd.Credentials {
		cred := database.Credential{
			ConnectionID: connId,
			Username:     credential.Username,
			Password:     credential.Password,
		}

		err := db.CreateCredential(&cred)
		if err != nil {
			slog.Error("failed to create credential", "error", err)
		}
	}

	for _, payload := range sd.Payloads {
		pl := database.Payload{ConnectionID: connId, Data: payload}

		err := db.CreatePayload(&pl)
		if err != nil {
			slog.Error("failed to create payload", "error", err)
		}
	}
}

func (r *Recorder) parseIp(c net.Conn) (string, int) {
	addr := c.RemoteAddr().String()
	slog.Debug("parsing address", "addr", addr)

	var err error
	var addrIp string
	var addrPort int

	if strings.Contains(addr, "[") {
		addrParts := strings.Split(addr, "]:")
		addrIp = strings.TrimPrefix(addrParts[0], "[")
		addrPort, err = strconv.Atoi(addrParts[1])
		if err != nil {
			slog.Error("failed to parse port", "error", err)
		}
	} else {
		addrParts := strings.Split(addr, ":")
		addrIp = addrParts[0]
		addrPort, err = strconv.Atoi(addrParts[1])
		if err != nil {
			slog.Error("failed to parse port", "error", err)
		}
	}

	return addrIp, addrPort
}

func (r *Recorder) getGeoLocationData(ip string) (*IpInfo, error) {
	client := retryablehttp.NewClient()
	client.RetryMax = 5
	client.RetryWaitMin = 1 * time.Second
	client.RetryWaitMax = 30 * time.Second
	client.Logger = nil

	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)
	req, err := retryablehttp.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ipInfo := IpInfo{}
	if err = json.Unmarshal(respData, &ipInfo); err != nil {
		return nil, err
	}

	return &ipInfo, nil
}
