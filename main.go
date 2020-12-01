package main

import (
	"os"
	"github.com/elastic/beats/filebeat/cmd"
)
func main() {
	if err := cmd.RoodCmd.Execute();err != nil {
		os.Exit(1)
	}
}
