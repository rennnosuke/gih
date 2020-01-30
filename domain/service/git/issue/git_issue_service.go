package issue

import (
	"github.com/rennnosuke/gih/domain/model/entity"
	"github.com/rennnosuke/gih/domain/repository"
)

type GitIssueService struct {
	Repo *repository.GitRepository
}

func (s *GitIssueService) GetIssue(id int) (*entity.Issue, error) {
	return (*s.Repo).GetIssue(id)
}

func (s *GitIssueService) GetIssues() (*[]entity.Issue, error) {
	return (*s.Repo).GetIssues()
}

func (s *GitIssueService) CreateIssue(title, description string) (*entity.Issue, error) {
	r := repository.IssueCreateRequest{Title: title, Description: description}
	return (*s.Repo).CreateIssue(&r)
}

func (s *GitIssueService) UpdateIssue(id int, title, description string) (*entity.Issue, error) {
	r := repository.IssueUpdateRequest{
		IssueId: id,
		IssueCreateRequest: repository.IssueCreateRequest{
			Title:       title,
			Description: description,
		},
	}
	return (*s.Repo).UpdateIssue(&r)
}

func (s *GitIssueService) CloseIssue(id int) (*entity.Issue, error) {
	return (*s.Repo).CloseIssue(id)
}
