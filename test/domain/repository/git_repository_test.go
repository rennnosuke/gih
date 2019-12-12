package repository

import (
	"fmt"
	"gih/domain/repository"
	"gih/registry"
	"os"
	"testing"
)

const (
	envKeyAccessToken    = "GIH_ACCESS_TOKEN"
	envKeyRepositoryName = "GIH_REPO_NAME"
	envKeyOrganization   = "GIH_ORG_NAME"
)

func TestGetIssue(t *testing.T) {

	accessToken := os.Getenv(envKeyAccessToken)
	repositoryName := os.Getenv(envKeyRepositoryName)
	organization := os.Getenv(envKeyOrganization)

	c := registry.NewGitRepository(accessToken, repositoryName, organization)

	result := c.GetIssue(10)
	fmt.Printf("%#v\n", result)

}

func TestGetIssues(t *testing.T) {

	accessToken := os.Getenv(envKeyAccessToken)
	repositoryName := os.Getenv(envKeyRepositoryName)
	organization := os.Getenv(envKeyOrganization)

	c := registry.NewGitRepository(accessToken, repositoryName, organization)

	result := c.GetIssues()
	fmt.Printf("%#v\n", result)

}

func TestCreateIssue(t *testing.T) {

	accessToken := os.Getenv(envKeyAccessToken)
	repositoryName := os.Getenv(envKeyRepositoryName)
	organization := os.Getenv(envKeyOrganization)

	c := registry.NewGitRepository(accessToken, repositoryName, organization)

	title := "Create Issue Test Title"
	description := "Create Issue Test Description."
	req := repository.IssueCreateRequest{Title: title, Description: description}
	result := c.CreateIssue(&req)
	fmt.Println(result)

}

func TestUpdateIssue(t *testing.T) {

	accessToken := os.Getenv(envKeyAccessToken)
	repositoryName := os.Getenv(envKeyRepositoryName)
	organization := os.Getenv(envKeyOrganization)

	c := registry.NewGitRepository(accessToken, repositoryName, organization)

	title := "Update Issue Test Title"
	description := "Update Issue Test Description."

	req := repository.IssueUpdateRequest{
		IssueId: 10,
		IssueCreateRequest: repository.IssueCreateRequest{
			Title:       title,
			Description: description,
		},
	}
	result := c.UpdateIssue(&req)
	fmt.Println(result)

}
