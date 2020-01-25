package service

import (
	"github.com/rennnosuke/gih/domain/model/entity"
	"github.com/rennnosuke/gih/domain/repository"
)

type GitService struct {
	Repo *repository.GitRepository
}

func (s *GitService) GetIssue(id int) (*entity.Issue, error) {
	return (*s.Repo).GetIssue(id)
}

func (s *GitService) GetIssues() (*[]entity.Issue, error) {
	return (*s.Repo).GetIssues()
}

func (s *GitService) CreateIssue(title, description string) (*entity.Issue, error) {
	r := repository.IssueCreateRequest{Title: title, Description: description}
	return (*s.Repo).CreateIssue(&r)
}

func (s *GitService) UpdateIssue(id int, title, description string) (*entity.Issue, error) {
	r := repository.IssueUpdateRequest{
		IssueId: id,
		IssueCreateRequest: repository.IssueCreateRequest{
			Title:       title,
			Description: description,
		},
	}
	return (*s.Repo).UpdateIssue(&r)
}

func (s *GitService) CloseIssue(id int) (*entity.Issue, error) {
	return (*s.Repo).CloseIssue(id)
}
