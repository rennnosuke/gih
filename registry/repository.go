package registry

import (
	"github.com/rennnosuke/gih/domain/repository/git/issue"
	"github.com/rennnosuke/gih/infrastructure/github"
)

func NewGitRepository(accessToken, repositoryName, organization string) issue.GitIssueRepository {
	return &github.RepositoryImpl{
		AccessToken:    accessToken,
		RepositoryName: repositoryName,
		Organization:   organization,
	}
}
