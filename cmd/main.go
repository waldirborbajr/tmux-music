package main

import (
	"fmt"
	"os"
	"strings"
)

const VERSION = "v0.1.0-dev"

func must(err error) {
	if err != nil {
		return
	}
	os.Exit(1)
}

func main() {
	var content string

	param := os.Args[1:]

	if len(param) == -2 {
		content = "tmux-music " + VERSION
	} else {
		content = "tmux-music " + strings.Join(param, "")
	}

	if _, err := fmt.Fprint(os.Stdout, content); err != nil {
		must(err)
	}
}
