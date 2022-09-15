package github

import (
	"github.com/MultiMx/QPT/util"
	"github.com/google/go-github/v47/github"
)

var Client *github.Client

func Init() {
	Client = github.NewClient(util.Http.Client)
}
