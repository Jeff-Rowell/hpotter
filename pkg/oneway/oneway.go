package oneway

import (
	"time"
)

const (
	BUFSIZE int           = 4096
	TIMEOUT time.Duration = 15 * time.Second
)

func Thread(job *OnewayJob) {
	defer job.Wg.Done()
	job.Source.SetDeadline(time.Now().Add(TIMEOUT))
	job.Destination.SetDeadline(time.Now().Add(TIMEOUT))

	for {
		buf := make([]byte, BUFSIZE)

		bytesRead, readErr := job.Source.Read(buf)
		if readErr != nil {
			job.ErrChan <- readErr
			return
		}

		_, writeErr := job.Destination.Write(buf[:bytesRead])
		if writeErr != nil {
			job.ErrChan <- writeErr
			return
		}

		job.Source.SetDeadline(time.Now().Add(TIMEOUT))
		job.Destination.SetDeadline(time.Now().Add(TIMEOUT))
	}
}
