package dev

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"runtime"
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

	os.Stdout.Write(highlight(stdoutb))
	os.Stderr.Write(highlight(stderrb))

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

func highlight(b []byte) []byte {
	r := bufio.NewReader(bytes.NewReader(b))
	var buf bytes.Buffer
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		if bytes.Contains(line, []byte("FAIL")) ||
			bytes.Contains(line, []byte("error")) ||
			bytes.Contains(line, []byte(".go:")) {
			red := color.RedString(string(line))
			buf.Write(appendNewLine([]byte(red)))
		} else {
			buf.Write(appendNewLine(line))
		}
	}
	return buf.Bytes()
}

func appendNewLine(b []byte) (newLined []byte) {
	cr := []byte("\n")
	if runtime.GOOS == "windows" {
		cr = []byte("\r\n")
	}
	newLined = append(b, cr...)
	return
}
