package github

import (
	"context"
	"github.com/google/go-github/v47/github"
)

func GetLatestRelease(owner, repo string) (*github.RepositoryRelease, error) {
	t, _, e := Client.Repositories.GetLatestRelease(context.Background(), owner, repo)
	return t, e
}
