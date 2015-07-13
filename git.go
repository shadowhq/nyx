package nyx

import (
	git "github.com/libgit2/git2go"
)

func GetRepo(clone_dir string) *git.Repository {
	repo, err := git.Clone(clone_dir, "sample_repo", &git.CloneOptions{})

	if err != nil {
		panic(err)
	}

	if repo == nil {
		panic(err)
	}

	return repo
}
