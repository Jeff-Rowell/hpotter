package database

import (
	"time"

	"gorm.io/gorm"
)

// Connection represents the connections table in the database for storing
// connection information collected by the honeypot.
type Connection struct {
	gorm.Model
	ContainerImage  string        // The container image
	SourceIpAddress string        // The source IP adddress of the connection
	DestIpAddress   string        // The destination IP address of the connection
	SourcePort      int           // The source port of the connection
	DestPort        int           // The destination port of the connection
	Duration        time.Duration // The duration of the connection
	Latitude        float32       // The latitude of the IP address
	Longitude       float32       // The longitude of the IP address
	Country         string        // The country of the IP address
	Region          string        // The region of the IP address
	City            string        // The city of the IP address
	Zipcode         string        // The zipcode of the IP address
	Credentials     []Credential  // The credentials sent over the connection
	Payloads        []Payload     // The payloads/commands sent over the connection
}

// Credential represents the credentials table in the database for storing
// credential information collected by the honeypot.
type Credential struct {
	gorm.Model
	ConnectionID uint   // The associated [Connection] relationship
	Username     string // The username sent over the connection
	Password     string // The password sent over the connection
}

// Payload repesents the payloads table in the database for storing payload
// and command data collected by the honeypot.
type Payload struct {
	gorm.Model
	ConnectionID uint   // The associated [Connection] relationship
	Data         string // The payload/command sent over the connection
}
