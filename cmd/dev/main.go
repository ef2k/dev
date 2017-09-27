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

	devfile, err := dev.ReadConfig(wd)
	if err != nil {
		log.Fatal(err)
	}

	if err := dev.Run(devfile); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
}
