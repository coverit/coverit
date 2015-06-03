package main

import (
	"gopkg.in/libgit2/git2go.v22"
)

type GitRepo struct {
	Path string
	Repo *git.Repository
}

type Signature struct {
	Name  string
	Email string
}

func OpenRepo(path string) (*GitRepo, error) {

	repo, err := git.OpenRepository(path)

	if err != nil {
		return nil, err
	}

	return &GitRepo{
		Path: path,
		Repo: repo,
	}, nil
}

func (repo *GitRepo) Branch() (string, error) {

	ref, err := repo.Repo.Head()

	if err != nil {
		return "", err
	}

	return ref.Branch().Name()
}

func (repo *GitRepo) CommitSha() (string, error) {

	ref, err := repo.Repo.Head()

	if err != nil {
		return "", err
	}

	return ref.Target().String(), nil
}

func (repo *GitRepo) Author() (*Signature, error) {

	ref, err := repo.Repo.Head()

	if err != nil {
		return nil, err
	}

	commit, err := repo.Repo.LookupCommit(ref.Target())

	if err != nil {
		return nil, err
	}

	return &Signature{
		Name:  commit.Author().Name,
		Email: commit.Author().Email,
	}, nil
}

func (repo *GitRepo) Committer() (*Signature, error) {

	ref, err := repo.Repo.Head()

	if err != nil {
		return nil, err
	}

	commit, err := repo.Repo.LookupCommit(ref.Target())

	if err != nil {
		return nil, err
	}

	return &Signature{
		Name:  commit.Committer().Name,
		Email: commit.Committer().Email,
	}, nil
}

func (repo *GitRepo) CommitMessage() (string, error) {

	ref, err := repo.Repo.Head()

	if err != nil {
		return "", err
	}

	commit, err := repo.Repo.LookupCommit(ref.Target())

	if err != nil {
		return "", err
	}

	return commit.Message(), nil
}
