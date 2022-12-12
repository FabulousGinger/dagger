package tasks

import (
	"context"
	"os"
	"strconv"
	"time"

	"dagger.io/dagger"
)

func Terraform(ctx context.Context, subtask string) (err error) {
	var terraformCommand []string

	switch subtask {
	case "plan":
		terraformCommand = []string{"plan"}
	case "apply":
		terraformCommand = []string{"apply", "-auto-approve"}
	case "destroy":
		terraformCommand = []string{"apply", "-destroy", "-auto-approve"}
	}

	// Create client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	CheckIfError(err)
	defer client.Close()

	now := strconv.Itoa(int(time.Now().Unix()))

	// Load terraform directory
	terraformDirectory := client.Host().Directory(os.Getenv("TERRAFORM_DIRECTORY"))

	// Load terraform image, init, and run
	terraform, err := client.Container().
		From("hashicorp/terraform:"+os.Getenv("TERRAFORM_VERSION")).
		WithMountedDirectory("/terraform", terraformDirectory).
		WithWorkdir(os.Getenv("TERRAFORM_WORK_DIR")).
		WithEnvVariable("AWS_ACCESS_KEY_ID", os.Getenv("AWS_ACCESS_KEY_ID")).
		WithEnvVariable("AWS_SECRET_ACCESS_KEY", os.Getenv("AWS_SECRET_ACCESS_KEY")).
		WithEnvVariable("AWS_DEFAULT_REGION", os.Getenv("AWS_DEFAULT_REGION")).
		WithEnvVariable("IGNORE_DAGGER_CACHE", now).
		WithExec([]string{"init"}).
		WithExec(terraformCommand).Stdout(ctx)

	CheckIfError(err)

	Info(terraform)

	return

}
