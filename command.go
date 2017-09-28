package dev

import (
	"io"
	"os/exec"
	"strings"
)

type command struct {
	name   string
	args   []string
	stderr io.Writer
	stdout io.Writer
}

func (c *command) run() error {
	cmd := exec.Command(c.name, c.args...)
	cmd.Stdout = c.stdout
	cmd.Stderr = c.stderr
	return cmd.Run()
}

func parseCmd(cmd string) (name string, args []string) {
	parts := strings.Split(cmd, " ")
	if len(parts) > 1 {
		name = parts[0]
		args = parts[1:]
	} else if len(parts) > 0 {
		name = parts[0]
	}
	return
}
