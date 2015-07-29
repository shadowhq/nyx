package nyx

import (
	git "github.com/libgit2/git2go"
)

func GetRepo(repo_url string) (*git.Repository, error) {
	repo, err := git.Clone(repo_url, "sample_repo", &git.CloneOptions{})
	if err != nil {
		// Pass error up to be dealt with by handler
		return nil, err
	}
	if repo == nil {
		// TODO: Figure out if this would happen without err returned
		panic(err)
	}

	return repo, nil
}
