package repository

import (
	"fmt"
	"github.com/rennnosuke/gih/domain/repository"
	"github.com/rennnosuke/gih/registry"
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

	result, err := c.GetIssue(10)
	if err != nil {
		fmt.Printf("%s\n", err)
		t.Fail()
	}

	fmt.Printf("%#v\n", result)

}

func TestGetIssues(t *testing.T) {

	accessToken := os.Getenv(envKeyAccessToken)
	repositoryName := os.Getenv(envKeyRepositoryName)
	organization := os.Getenv(envKeyOrganization)

	c := registry.NewGitRepository(accessToken, repositoryName, organization)

	result, err := c.GetIssues()
	if err != nil {
		fmt.Printf("%s\n", err)
		t.Fail()
	}

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

	result, err := c.CreateIssue(&req)
	if err != nil {
		fmt.Printf("%s\n", err)
		t.Fail()
	}

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

	result, err := c.UpdateIssue(&req)

	if err != nil {
		fmt.Printf("%s\n", err)
		t.Fail()
	}

	fmt.Println(result)

}
