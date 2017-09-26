package dev

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func command(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
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

func runTask(t Task) error {
	name, args := parseCmd(t.Cmd)
	return command(name, args...)
}

func build(tasks []Task) error {
	var err error
	for _, t := range tasks {
		fmt.Printf("%s %s\n      %s \n\n", "[run]", t.Title, t.Cmd)
		err = runTask(t)
		if err != nil {
			break
		}
	}
	return err
}
