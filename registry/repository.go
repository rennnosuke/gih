package registry

import (
	"gih/domain/repository"
	"gih/infrastructure/github"
)

func NewGitRepository(accessToken, repositoryName, organization string) repository.GitRepository {
	return &github.RepositoryImpl{
		AccessToken:    accessToken,
		RepositoryName: repositoryName,
		Organization:   organization,
	}
}
