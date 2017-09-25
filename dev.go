package dev

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"

	"github.com/radovskyb/watcher"
)

type Config struct {
	Path string
}

type Task struct {
	Cmd     string
	Title   string
	Watch   string
	Default bool
}

type Devfile struct {
	Tasks []Task
	Watch string
}

func Run(c *Config) error {
	// Read the config file
	configPath := path.Join(c.Path, "devfile")
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		return errors.New("No configuration found. Create one:\n\n\t$ dev init\n")
	}

	// YAML to struct
	var devfile Devfile
	yaml.Unmarshal(b, &devfile)
	if err != nil {
		return err
	}

	// Validate/confirm the watcher path
	if devfile.Watch == "." {
		devfile.Watch = c.Path
	} else {
		p := devfile.Watch
		if _, err := os.Stat(p); os.IsNotExist(err) {
			return errors.New(fmt.Sprintf("The given watch path does not exist: %s", p))
		}
	}

	build := func() error {
		for _, t := range devfile.Tasks {
			cargs := strings.Split(t.Cmd, " ")
			cmd := exec.Command(cargs[0], cargs[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				return err
			} else {
				fmt.Printf("[ran] %s \n", t.Cmd)
			}
		}
		return nil
	}

	// Run the very first build
	if err := build(); err != nil {
		return err
	}

	// Start a watcher if defined in the config.
	if devfile.Watch != "" {
		w := watcher.New()
		w.SetMaxEvents(2)
		w.IgnoreHiddenFiles(true)

		if err := w.AddRecursive(c.Path); err != nil {
			log.Fatalln(err)
		}

		go func() {
			for {
				select {
				case event := <-w.Event:
					if !event.IsDir() {
						build()
					}
				case err := <-w.Error:
					log.Printf("[err] %v", err)
				case <-w.Closed:
					log.Print("Watcher chan closed")
				}
			}
		}()

		if err := w.Start(time.Millisecond * 200); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
