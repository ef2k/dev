package dev

import (
	"errors"
	"io/ioutil"
	"path"

	yaml "gopkg.in/yaml.v2"
)

const devfile = "dev.yaml"

// Devfile holds the unmarshaled yaml content of a devfile
type Devfile struct {
	Tasks []Task
	Watch string
}

// Task wraps a system command along with its meta data.
type Task struct {
	Cmd     string
	Title   string
	Watch   string
	Default bool
}

// ReadConfig finds and reads a Devfile at the given directory path.
func ReadConfig(dirPath string) (*Devfile, error) {
	configPath := path.Join(dirPath, devfile)
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		msg := "No configuration found. Run: \n\n\t$ dev init\n"
		return nil, errors.New(msg)
	}

	var devfile Devfile
	if err := yaml.Unmarshal(b, &devfile); err != nil {
		return nil, err
	}

	return &devfile, nil
}
