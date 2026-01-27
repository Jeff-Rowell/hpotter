package session

import (
	"regexp"
)

// Recorder defines the username, password, and payload patterns to be used
// to parse logs by a [SessionRecorder].
type Recorder struct {
	// UsernamePattern defines the log pattern matching usernames
	UsernamePattern *regexp.Regexp

	// PasswordPattern defines the log pattern matching passwords
	PasswordPattern *regexp.Regexp

	// PayloadPattern defines the log pattern matching the payload/commands
	PayloadPattern *regexp.Regexp
}

// Credential represents an obtained username and password combination.
type Credential struct {
	// Username is the username found in the logs
	Username string

	// Password is the password found in the logs
	Password string
}

// SessionData is the data obtained from a [SessionRecorder] and contains a
// slice of [Credential] and slice of data payloads or commands stored as a
// slice of strings.
type SessionData struct {
	// Credentials is a slice containing all of the [Credential] usernames and
	// passwords found in the given container's logs
	Credentials []Credential

	// Payloads is a slice containing all of the [Payload] request data or
	// commands found in the given container's logs
	Payloads []string
}

// SessionRecorder is the interface that groups the RecordSession and
// SaveSession methods.
type SessionRecorder interface {
	// RecordSession records a session given container id c and uses the
	// patterns defined in r to parse the container logs of c. The session
	// will be recorded until the io.Reader for c is closed or an error import
	// encountered. The recorded SessionData is returned over a channel if
	// successful, otherwise an error will be delivered over an error channel.
	GetSessionData(c string) (*SessionData, error)

	// SaveSession saves the SessionData to the given database connection.
	SaveSession(d *SessionData)
}

// IpInfo contains the geo-location lookup information for an IP address
type IpInfo struct {
	// The latitude of the IP address
	Latitude float32 `json:"lat"`

	// The longitude of the IP address
	Longitude float32 `json:"lon"`

	// The country of the IP address
	Country string `json:"country"`

	// The region of the IP address
	Region string `json:"region"`

	// The city of the IP address
	City string `json:"city"`

	// The zipcode of the IP address
	Zipcode string `json:"zipcode"`
}
