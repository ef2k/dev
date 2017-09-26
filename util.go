package dev

import (
	"fmt"
	"time"
)

func timeNow() string {
	t := time.Now()
	return t.Format("15:04:05")
}

func printHeaderTime() {
	fmt.Printf("---\n%s \n", timeNow())
}
