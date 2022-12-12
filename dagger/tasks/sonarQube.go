package tasks

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"dagger.io/dagger"
)

func Scan(ctx context.Context) (err error) {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	CheckIfError(err)
	defer client.Close()

	sonarVersion := os.Getenv("SONAR_VERSION")
	sonarHost := os.Getenv("SONAR_WEB_HOST")
	sonarToken := os.Getenv("SONAR_TOKEN")
	sonarProject := os.Getenv("SONAR_PROJECT")

	image := fmt.Sprintf("sonarsource/sonar-scanner-cli:%s", sonarVersion)
	now := strconv.Itoa(int(time.Now().Unix()))
	src := client.Host().Directory(os.Getenv("APP_DIRECTORY"))

	output, err := client.
		Container(dagger.ContainerOpts{Platform: "linux/amd64"}).
		From(image).
		WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithEnvVariable("IGNORE_DAGGER_CACHE", now).
		WithEnvVariable("SONAR_HOST_URL", sonarHost).
		WithEnvVariable("SONAR_LOGIN", sonarToken).
		WithEnvVariable("SONAR_SCANNER_OPTS", "-Dsonar.projectKey="+sonarProject).
		WithEnvVariable("SONAR_ES_BOOTSTRAP_CHECKS_DISABLE", "true").
		WithExec([]string{
			"sonar-scanner",
			"-X",
		}).
		Stdout(ctx)
	CheckIfError(err)

	Info("\n%s", output)

	return
}
