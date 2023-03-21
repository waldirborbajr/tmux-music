package main

import (
	"fmt"
	"os"
	"strings"
)

const VERSION = "v0.1.0"

func must(err error) {
	if err != nil {
		return
	}
	os.Exit(1)
}

func main() {
	var content string

	param := os.Args[1:]

	// fmt.Printf("Size %d \n", len(param))
	// fmt.Printf("Param %s \n\n", strings.Join(param, ""))

	if len(param) == 0 {
		content = "tmux-music " + VERSION
	} else {
		content = "tmux-music " + strings.Join(param, "")
	}

	if _, err := fmt.Fprintf(os.Stdout, content); err != nil {
		must(err)
	}
}
