// Package main ...
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/praveenmahasena/gocacheserver/internal"
)

func main() {
	go func() {
		time.Sleep(5 * time.Second)
		os.Exit(-1)
	}()
	if err := internal.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(0)
	}
}
