package main

import (
	"os"

	"github.com/fabulousginger/dagger/dagger/tasks"
	"github.com/joho/godotenv"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		tasks.Info(usage)
		return
	}

	task := args[0]

	err := godotenv.Load()
	tasks.CheckIfError(err)

	err = selectTask(args, task)
	tasks.CheckIfError(err)
}
