package registry

import (
	"github.com/rennnosuke/gih/domain/service"
)

func NewGitService(accessToken, repositoryName, organization string) *service.GitService {
	repo := NewGitRepository(accessToken, repositoryName, organization)
	return &service.GitService{&repo}
}
