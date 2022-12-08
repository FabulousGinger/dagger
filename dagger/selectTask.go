package main

import (
	"context"

	"github.com/fabulousginger/dagger-demo/dagger/tasks"
)

func selectTask(args []string, task string) (err error) {
	ctx := context.Background()

	switch task {
	case "git":
		if len(args) < 2 {
			tasks.Info(gitUsage)
			return
		}
		shortHash, err := tasks.GitHash()
		tasks.CheckIfError(err)
		tasks.Info(shortHash)
	case "ecrlogin":
		err = tasks.ECRLogin()
	case "test":
		err = tasks.Test(ctx)
	case "push":
		_, err = tasks.Push(ctx)
	case "task":
		if len(args) < 2 {
			tasks.Info(ECSTaskUsage)
			return
		}
		_, err = tasks.ECSFargateTask(args[1])
	case "service":
		if len(args) < 2 {
			tasks.Info(ECSServiceUsage)
			return
		}
		err = tasks.ECSFargateService(args[1])
		tasks.CheckIfError(err)
	case "deploy":
		err = tasks.ECSDeploy(ctx)
	case "tf":
		if len(args) < 2 {
			tasks.Info(tfUsage)
			return
		}
		subtask := args[1]
		err = tasks.Tf(ctx, subtask)
	default:
		tasks.Info("Unknown task: %s", task)
		return
	}
	return
}
