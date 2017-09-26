package dev

import (
	"log"
)

func Run(wd string) error {
	// Read the devfile
	devfile, err := ReadConfig(wd)
	if err != nil {
		return err
	}

	// Validate the given watcher path
	devfile.Watch, err = validatePath(devfile.Watch)
	if err != nil {
		return err
	}

	// Run the very first build
	printHeaderTime()
	if err := build(devfile.Tasks); err != nil {
		return err
	}

	// If no watch path is given, then we're done.
	if devfile.Watch == "" {
		return nil
	}

	// Otherwise, watch  for changes.
	changed := make(chan struct{}, 1)
	errored := make(chan error, 1)
	closed := make(chan struct{}, 1)

	go watch(devfile.Watch, changed, errored, closed)
	for {
		select {
		case <-changed:
			printHeaderTime()
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
