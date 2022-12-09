package tasks

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"dagger.io/dagger"
)

func Test(ctx context.Context) error {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	CheckIfError(err)
	defer client.Close()

	src := client.Host().Directory(os.Getenv("APP_DIRECTORY"), dagger.HostDirectoryOpts{
		Exclude: []string{
			".circleci/*",
			".github/*",
			"dagger/*",
			"terraform/*",
			"output/*",
		},
	})

	testoutput := client.Directory()
	cacheKey := "gomods"
	cache := client.CacheVolume(cacheKey)

	now := strconv.Itoa(int(time.Now().Unix()))

	// multiplatform tests
	goversions := []string{"1.18", "1.19"}
	platforms := []dagger.Platform{"linux/amd64", "linux/arm64"}

	for _, plat := range platforms {
		for _, goversion := range goversions {
			image := fmt.Sprintf("golang:%s", goversion)
			builder := client.Container(dagger.ContainerOpts{Platform: plat}).
				From(image).
				WithMountedDirectory("/src", src).
				WithWorkdir("/src").
				WithMountedCache("/cache", cache).
				WithEnvVariable("GOMODCACHE", "/cache").
				WithEnvVariable("IGNORE_DAGGER_CACHE", now).
				WithExec([]string{"sh", "-c", "go test > /src/test.out"})

			output, err := builder.
				WithExec([]string{"sh", "-c", "go test"}).
				WithEnvVariable("IGNORE_DAGGER_CACHE", now).Stdout(ctx)

			CheckIfError(err)

			Info("Platform: %s\nGO Version: %s\n\n%s", plat, goversion, output)

			local, err := strconv.ParseBool(os.Getenv("LOCAL"))
			CheckIfError(err)

			if local {
				outputfile := fmt.Sprintf("output/%s/%s.out", string(plat), goversion)
				testoutput = testoutput.WithFile(
					outputfile,
					builder.File("/src/test.out"),
				)
			}
		}
	}

	_, err = testoutput.Export(ctx, ".")
	CheckIfError(err)

	return nil
}
