package repository

import (
	"gih/domain/model/entity"
)

type GitRepository interface {
	// issue manipulation
	GetIssue(issueId int) *entity.Issue
	GetIssues() *[]entity.Issue
	CreateIssue(r *IssueCreateRequest) *entity.Issue
	UpdateIssue(r *IssueUpdateRequest) *entity.Issue
	CloseIssue(id int) *entity.Issue
}

type IssueCreateRequest struct {
	Title       string
	Description string
}

type IssueUpdateRequest struct {
	IssueId int
	IssueCreateRequest
}
