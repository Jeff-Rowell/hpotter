package parser

import "github.com/hpotter/pkg/session"

// Parser is the interface that wraps the basic Parse method.
type Parser interface {
	// Parse parses a JSON or YAML file defined by an absolute path p.
	// It returns a pointer a Config for the honey pot and an error, if any.
	// Parse returns a non-nil error if the config file cannot be
	// found or if the contents of the config file are invalid.
	Parse(p string) (c *Config, err error)
}

// YAMLParser is the implementation of the Parser interface for YAML files.
type YAMLParser struct{}

// YAMLConfig contains the configuration for the honey pot. The configuration
// is populated from calling configparser.Parse() with a path to a YAML
// or YAML configuration file.
type Config struct {
	// Services is the slice of Service objects parse from the YAML
	// configuration file.
	Services []Service

	// Database is the database configuration fields parsed from the YAML
	// config file.
	Database Database
}

// Service is the definition of the ficticious honey pot service to run.
// The Service specifies details on how to run the honey pot service
// including what host-based address and port to listen on as well as
// what container port to forward the traffic to.
type Service struct {
	// Name is the name of the service and is used as a human readable
	// identifier of the service.
	Name string `yaml:"name"`

	// ListenAddress is the address that will be used to host the
	// ficticious honey pot service and will listen for new connections.
	// If nothing is provided, this will be set to localhost (127.0.0.1).
	ListenAddress string `yaml:"listen_address"`

	// ListenPort is the port that will used ot host the ficticious honey
	// pot service and will listen for new connections. This is required
	// and will not run unless provided.
	ListenPort string `yaml:"listen_port"`

	// ListenProto is the protocol for ListenPort
	ListenProto string `yaml:"listen_proto"`

	// Image is the container image to run for the ficticious honey pot service.
	// This is required and will not run unless provided.
	Image string `yaml:"image"`

	// ContainerPort is the container NAT port that the ListenPort will forward
	// traffic to. This is not required and defaults to the default port of
	// the provided image in Image.
	ContainerPort string `yaml:"container_port"`

	// GenerateCerts is used to determine whether or not a certificate and key
	// pair needs to be generated. This can only be used when UseTLS is true.
	GenerateCerts bool `yaml:"generate_certs"`

	// Env is a list of environment variables to be passed into the container
	// image. The format is "ENV_KEY=ENV_VALUE"
	Env []string `yaml:"env"`

	// UsernamePattern defines the log pattern matching usernames
	UsernamePattern string `yaml:"username_pattern"`

	// PasswordPattern defines the log pattern matching passwords
	PasswordPattern string `yaml:"password_pattern"`

	// PayloadPattern defines the log pattern matching the payload/commands
	PayloadPattern string `yaml:"payload_pattern"`

	// Recorder is the session recorder that keeps track of the usernames,
	// passwords, and commands/payloads run inside the container
	Recorder *session.Recorder
}

// Database is the database configuration elements to be parsed from the
// YAML config file.
type Database struct {
	// Image is the container image of the database
	Image string `yaml:"image"`

	// Port is the port to connect to the database on
	Port int `yaml:"port"`

	// Username is the username to connect to the database with
	Username string `yaml:"username"`

	// Pasword is the password to connect to the database with
	Password string `yaml:"password"`
}
