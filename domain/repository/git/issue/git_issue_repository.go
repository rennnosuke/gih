package issue

import "github.com/rennnosuke/gih/domain/model/entity"

type GitIssueRepository interface {
	// issue manipulation
	GetIssue(issueId int) (*entity.Issue, error)
	GetIssues() (*[]entity.Issue, error)
	CreateIssue(r *IssueCreateRequest) (*entity.Issue, error)
	UpdateIssue(r *IssueUpdateRequest) (*entity.Issue, error)
	CloseIssue(id int) (*entity.Issue, error)
}

type IssueCreateRequest struct {
	Title       string
	Description string
}

type IssueUpdateRequest struct {
	IssueId int
	IssueCreateRequest
}
