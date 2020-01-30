package registry

import (
	"github.com/rennnosuke/gih/domain/service/git/issue"
)

func NewGitService(accessToken, repositoryName, organization string) *issue.GitIssueService {
	repo := NewGitRepository(accessToken, repositoryName, organization)
	return &issue.GitIssueService{&repo}
}
