package main

import (
	"job/internal/jobserver"
	"os"
)

func main() {
	command := jobserver.NewJobServerCommand()
	err := command.Execute()
	if err != nil {
		os.Exit(1)
	}
}
