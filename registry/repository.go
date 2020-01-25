package registry

import (
	"github.com/rennnosuke/gih/domain/repository"
	"github.com/rennnosuke/gih/infrastructure/github"
)

func NewGitRepository(accessToken, repositoryName, organization string) repository.GitRepository {
	return &github.RepositoryImpl{
		AccessToken:    accessToken,
		RepositoryName: repositoryName,
		Organization:   organization,
	}
}
