package dev

import (
	"fmt"
	"log"
	"time"
)

type Config struct {
	Path string
}

func timeNow() string {
	t := time.Now()
	return fmt.Sprintf("%d:%d:%d", t.Hour(), t.Minute(), t.Second())
}

func Run(c *Config) error {
	// Read the devfile
	devfile, err := ReadConfig(c.Path)
	if err != nil {
		return err
	}

	// Validate the given watcher path
	devfile.Watch, err = validatePath(devfile.Watch)
	if err != nil {
		return err
	}

	// Run the very first build
	if err := build(devfile.Tasks); err != nil {
		return err
	}

	// If no watch path is given, then we're done.
	if devfile.Watch == "" {
		return nil
	}

	// Watch  for changes.
	changed := make(chan struct{}, 1)
	errored := make(chan error, 1)
	closed := make(chan struct{}, 1)

	go watch(devfile.Watch, changed, errored, closed)
	for {
		select {
		case <-changed:
			fmt.Printf("---\n%s\n", timeNow())
			build(devfile.Tasks)
		case err := <-errored:
			log.Print(err)
		case <-closed:
			close(changed)
			close(errored)
			close(closed)
			break
		}
	}
	return nil
}
