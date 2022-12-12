package tasks

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
)

func GitHash() (shortHash string, err error) {

	repository := os.Getenv("GIT_REPOSITORY")
	user := os.Getenv("GIT_USER")
	token := os.Getenv("GIT_TOKEN")

	Info("git clone " + repository)
	// Clones the given repository, creating the remote, the local branches
	// and fetching the objects, everything in memory:
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: repository,
		Auth: &http.BasicAuth{
			Username: user,
			Password: token,
		},
	})
	CheckIfError(err)

	Info("git rev-parse --short HEAD")
	// ... retrieves the branch pointed by HEAD
	ref, err := r.Head()
	CheckIfError(err)
	hash := ref.Hash().String()

	shortHash = hash[0:7]

	return

}
