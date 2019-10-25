package main

import (
	"grok/cmd"
)

// Version of grok. Overwritten during build
var Version = "development"

func main() {
	cmd.Execute(Version)
}
