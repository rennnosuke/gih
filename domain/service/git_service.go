package service

import (
	"gih/domain/model/entity"
	"gih/domain/repository"
)

type GitService struct {
	Repo *repository.GitRepository
}

func (s *GitService) GetIssue(id int) *entity.Issue {
	return (*s.Repo).GetIssue(id)
}

func (s *GitService) GetIssues() *[]entity.Issue {
	return (*s.Repo).GetIssues()
}

func (s *GitService) CreateIssue(title, description string) *entity.Issue {
	r := repository.IssueCreateRequest{Title: title, Description: description}
	return (*s.Repo).CreateIssue(&r)
}

func (s *GitService) UpdateIssue(id int, title, description string) *entity.Issue {
	r := repository.IssueUpdateRequest{
		IssueId: id,
		IssueCreateRequest: repository.IssueCreateRequest{
			Title:       title,
			Description: description,
		},
	}
	return (*s.Repo).UpdateIssue(&r)
}

func (s *GitService) CloseIssue(id int) *entity.Issue {
	return (*s.Repo).CloseIssue(id)
}
