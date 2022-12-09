package tasks

import (
	"context"
	"os"
	"strconv"
	"time"

	"dagger.io/dagger"
)

func Push(ctx context.Context) (address string, err error) {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	CheckIfError(err)
	defer client.Close()

	gitHash, err := GitHash()
	CheckIfError(err)

	repository := os.Getenv("AWS_REPOSITORY")
	src := client.Host().Directory(os.Getenv("APP_DIRECTORY"))

	now := strconv.Itoa(int(time.Now().Unix()))

	Info("Building Docker image")
	daggerImg := client.Container().Build(src).
		WithEnvVariable("CGO_ENABLED", "0").
		WithEnvVariable("GOOS", "linux").
		WithEnvVariable("GOARCH", "amd64").
		WithEnvVariable("IGNORE_DAGGER_CACHE", now).
		WithEnvVariable("DOCKER_DEFAULT_PLATFORM", "linux/amd64")

	Info("Pushing Docker image")
	address, err = daggerImg.Publish(
		ctx,
		repository+":"+gitHash,
		dagger.ContainerPublishOpts{})
	CheckIfError(err)

	Info("Docker Image pushed successfully: %s", address)

	return

}
