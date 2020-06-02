package main

import (
	"github.com/elastic/beats/filebeat/cmd"
	_ "github.com/helplove/ccbbeats/fileout"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}