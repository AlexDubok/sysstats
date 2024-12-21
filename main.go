package main

import (
	"sysstats-cli/cmd"
	"sysstats-cli/src/config"
)

func main() {
	config.LoadConfig()
	cmd.Execute()
}
