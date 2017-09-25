package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ef2k/dev"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	cfg := &dev.Config{
		Path: wd,
	}

	if err := dev.Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
}
