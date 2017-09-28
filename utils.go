package dev

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func timeNow() string {
	t := time.Now()
	return t.Format("15:04:05")
}

func printHeaderTime() {
	fmt.Printf("---\n%s \n", timeNow())
}

type tempfile struct {
	f *os.File
}

func (t *tempfile) remove() {
	_ = os.Remove(t.f.Name())
}

func newtempfile() *tempfile {
	f, _ := ioutil.TempFile("", "devtmp")
	return &tempfile{
		f,
	}
}

func clearScreen(f *os.File) {
	clear := &command{
		name:   "clear",
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
	_ = clear.run()
}
