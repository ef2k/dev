package dev

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

func runTask(t Task) error {
	name, args := parseCmd(t.Cmd)

	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	cmd := &command{
		name:   name,
		args:   args,
		stdout: stdout,
		stderr: stderr,
	}

	run := color.GreenString("[running]")
	fmt.Printf("%s %s\n", run, t.Title)

	t0 := time.Now()
	err := cmd.run()
	dt := time.Since(t0)

	fmt.Printf("          %s\n", t.Cmd)
	fmt.Printf("          %s\n", dt)

	stdoutb := stdout.Bytes()
	stderrb := stderr.Bytes()

	os.Stdout.Write(stdoutb)
	os.Stderr.Write(stderrb)

	return err
}

func runTasks(tasks []Task) error {
	var err error
	for _, t := range tasks {
		err = runTask(t)
		if err != nil {
			break
		}
	}
	return err
}
