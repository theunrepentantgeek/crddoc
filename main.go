package main

import (
	"os"

	"github.com/theunrepentantgeek/crddoc/cmd"
)

func main() {
	log := cmd.CreateLogger()

	err := cmd.Execute(log)
	if err != nil {
		log.Error(err, "command failed")
		os.Exit(1)
	}
}
