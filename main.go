package main

import (
	"os"
	"github.com/elastic/beats/filebeat/cmd"
	_ "github.com/qhhan/kafka_filebeat/kafkawithoffset"
)
func main() {
	if err := cmd.RoodCmd.Execute();err != nil {
		os.Exit(1)
	}
}
