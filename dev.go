package dev

import (
	"log"
	"os"
)

// Run invokes the commands found in the given Devfile. If a watcher is defined,
// the path is validated and a filesystem watcher is started. If an error occurs
// at any point of execution, it's immediately returned.
func Run(devfile *Devfile) error {
	// Validate watch path
	validPath, err := validatePath(devfile.Watch)
	if err != nil {
		return err
	}
	devfile.Watch = validPath

	// If no tasks were given, then we're done.
	// Otherwise, run them atleast once.
	if len(devfile.Tasks) <= 0 {
		return nil
	} else {
		if err := runTasks(devfile.Tasks); err != nil {
			return err
		}
	}

	// If no watcher is needed, then we're done.
	if devfile.Watch == "" {
		return nil
	}

	// Otherwise, watch for changes.
	changed := make(chan struct{}, 1)
	errored := make(chan error, 1)
	closed := make(chan struct{}, 1)
	go watch(devfile.Watch, changed, errored, closed)
	for {
		select {
		case <-changed:
			clearScreen(os.Stdout)
			runTasks(devfile.Tasks)
		case err := <-errored:
			log.Print(err)
		case <-closed:
			break
		}
	}

	return nil
}
