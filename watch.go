package dev

import (
	"os"
	"time"

	"github.com/radovskyb/watcher"
)

// validatePath checks if the given path exists. Returns an expanded path
// string from the given path or an error if the path doesn't exist.
func validatePath(path string) (wd string, err error) {
	if path == "." {
		return os.Getwd()
	}
	if _, err = os.Stat(path); os.IsNotExist(err) {
		wd = path
	}
	return
}

// watch watches the given workingDir for changes and notifies of changed,
// errored, and closed states via the given channels.
func watch(workingDir string, changed chan struct{}, errored chan error, closed chan struct{}) {
	w := watcher.New()
	w.SetMaxEvents(2)
	w.IgnoreHiddenFiles(true)

	if err := w.AddRecursive(workingDir); err != nil {
		errored <- err
	}

	go func() {
		for {
			select {
			case event := <-w.Event:
				if !event.IsDir() {
					changed <- struct{}{}
				}
			case err := <-w.Error:
				errored <- err
			case c := <-w.Closed:
				closed <- c
			}
		}
	}()

	if err := w.Start(time.Millisecond * 200); err != nil {
		errored <- err
	}
}
