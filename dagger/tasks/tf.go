package tasks

import (
	"context"
	"os"
	"strconv"
	"time"

	"dagger.io/dagger"
)

func Tf(ctx context.Context, subtask string) error {
	var tfcommand []string

	switch subtask {
	case "plan":
		tfcommand = []string{"plan"}
	case "apply":
		tfcommand = []string{"apply", "-auto-approve"}
	case "destroy":
		tfcommand = []string{"apply", "-destroy", "-auto-approve"}
	}

	// Create client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}
	defer client.Close()

	now := strconv.Itoa(int(time.Now().Unix()))

	// Load terraform directory
	tfdirectory := client.Host().Directory(os.Getenv("TERRAFORM_DIRECTORY"))

	// Load terraform image, init, and run
	tf := client.Container().
		From("hashicorp/terraform:"+os.Getenv("TERRAFORM_VERSION")).
		WithMountedDirectory("/terraform", tfdirectory).
		WithWorkdir(os.Getenv("TERRAFORM_WORK_DIR")).
		WithEnvVariable("AWS_ACCESS_KEY_ID", os.Getenv("AWS_ACCESS_KEY_ID")).
		WithEnvVariable("AWS_SECRET_ACCESS_KEY", os.Getenv("AWS_SECRET_ACCESS_KEY")).
		WithEnvVariable("AWS_DEFAULT_REGION", os.Getenv("AWS_DEFAULT_REGION")).
		WithEnvVariable("CACHEBUSTER", now).
		WithExec([]string{"init"}).
		WithExec(tfcommand)

	// Execute against dagger engine
	_, err = tf.ExitCode(ctx)

	return err

}
