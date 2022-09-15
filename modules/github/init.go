package github

import (
	"context"
	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
	"time"
)

var AccessToken = "unknown"

var Client *github.Client

func Init() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: AccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	tc.Timeout = time.Second * 30

	Client = github.NewClient(tc)
}
