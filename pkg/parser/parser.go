package parser

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/hpotter/pkg/session"
	"gopkg.in/yaml.v3"
)

const emptyPathError = "config file must be an absolute path."

// NewYAMLParser Creates a new YAML parser.
func NewYAMLParser() *YAMLParser {
	return &YAMLParser{}
}

// Parse parses a YAML file defined by an absolute path p.
// It returns a pointer a Config for the honey pot and an error, if any.
// Parse returns a non-nil error if the config file cannot be
// found or if the contents of the config file are invalid.
func (yp *YAMLParser) Parse(p string) (*Config, error) {
	if p == "" {
		return nil, fmt.Errorf(emptyPathError)
	}

	handle, err := os.Open(p)
	if err != nil {
		return nil, fmt.Errorf("failed to open file '%s'", p)
	}

	data, err := io.ReadAll(handle)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	err = yp.buildServiceRecorders(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (yp *YAMLParser) buildServiceRecorders(c *Config) error {
	for idx := range c.Services {
		if c.Services[idx].UsernamePattern == "" ||
			c.Services[idx].PasswordPattern == "" ||
			c.Services[idx].PayloadPattern == "" {
			return fmt.Errorf("invalid recorder pattern in service '%s'", c.Services[idx].Name)
		}

		username, err := regexp.Compile(c.Services[idx].UsernamePattern)
		if err != nil {
			return err
		}

		password, err := regexp.Compile(c.Services[idx].PasswordPattern)
		if err != nil {
			return err
		}

		payload, err := regexp.Compile(c.Services[idx].PayloadPattern)
		if err != nil {
			return err
		}

		recorder := &session.Recorder{
			UsernamePattern: username,
			PasswordPattern: password,
			PayloadPattern:  payload,
		}

		c.Services[idx].Recorder = recorder
	}

	return nil
}
