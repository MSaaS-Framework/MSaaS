package git

// 참고 자료
// https://github.com/go-git/go-git/tree/master/_examples
// https://github.com/google/go-github

import (
	"os"
	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Repo struct {
	URL       string
	LocalPath string
}

func SetGitOptions() {
	// git options

}

func CloneRepo(repo Repo) {
	// git clone
	_, err := git.PlainClone(repo.LocalPath, false, &git.CloneOptions{
		URL:      repo.URL,
		Progress: os.Stdout,
	})
	if err != nil {
		panic(err)
	}
}

// Clone method 1. username and password in URL
// Clone method 2. personal access token
// Clone method 3. ssh private key
// Clone method 4. ssh agent

func PullRepo(repo Repo) {
	// git pull
	r, err := git.PlainOpen(repo.LocalPath)
	if err != nil {
		panic(err)
	}

	w, err := r.Worktree()
	if err != nil {
		panic(err)
	}

	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		panic(err)
	}
}

type Commit struct {
	Author  string
	Email   string
	Message string
}

func CommitAndPushRepo(repo Repo, commit Commit) {
	// git push
	r, err := git.PlainOpen(repo.LocalPath)
	if err != nil {
		panic(err)
	}

	w, err := r.Worktree()
	if err != nil {
		panic(err)
	}

	err = w.AddGlob(".")
	if err != nil {
		panic(err)
	}

	_, err = w.Commit(commit.Message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  commit.Author,
			Email: commit.Email,
			When:  time.Now(),
		},
	})
	if err != nil {
		panic(err)
	}

	err = r.Push(&git.PushOptions{
		RemoteName: "origin",
	})
	if err != nil {
		panic(err)
	}
}
