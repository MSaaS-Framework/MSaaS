package git

import (
	"context"
	"log"

	"github.com/google/go-github/v68/github"
)

func NewGitHubClient() *github.Client {
	// github client
	client := github.NewClient(nil)
	return client
}

func GetGitHubRepo(client *github.Client) ([]*github.Repository, error) {
	// get repo
	// list public repositories for org "github"
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "github", opt)
	if err != nil {
		log.Fatal(err)
	}
	return repos, err
}
