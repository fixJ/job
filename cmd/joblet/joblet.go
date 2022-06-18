package main

import (
	"job/internal/joblet"
	"os"
)

func main() {
	cmd := joblet.NewJobletCommand()
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
