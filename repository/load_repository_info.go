// Package repository provides functions to load information about a Git repository
package repository

import (
	"fmt"
	"github.com/go-git/go-git/v5"
)

// LoadRepositoryInfo loads Git repository information given a provided `path`.
// It returns an error if the path doesn't exist or the path doesn't contain a valid Git repository.
func LoadRepositoryInfo(path string) (*RepositoryInfo, error) {
	fmt.Println("Search for RepositoryInfo into: " + path)
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}
	remotes, err := repo.Remotes()
	if err != nil {
		return nil, err
	}
	remoteUrl := "unknown"
	for _, r := range remotes {
		if r.Config().Name == "origin" {
			remoteUrl = r.Config().URLs[0]
		}
	}
	head, err := repo.Head()
	if err != nil {
		return nil, err
	}

	wt, err := repo.Worktree()
	if err != nil {
		return nil, err
	}
	status, err := wt.Status()
	if err != nil {
		return nil, err
	}

	return &RepositoryInfo{
		Path:            path,
		Url:             remoteUrl,
		CurrentBranch:   head.Name().Short(),
		CurrentRevision: head.Hash().String(),
		IsUpToDate:      status.IsClean(),
	}, nil
}
