package service

import (
	"fmt"
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

	s := registry.NewGitService(accessToken, repositoryName, organization)

	result := s.GetIssue(10)
	fmt.Printf("%#v\n", result)

}

func TestGetIssues(t *testing.T) {

	accessToken := os.Getenv(envKeyAccessToken)
	repositoryName := os.Getenv(envKeyRepositoryName)
	organization := os.Getenv(envKeyOrganization)

	s := registry.NewGitService(accessToken, repositoryName, organization)

	result := s.GetIssues()
	fmt.Printf("%#v\n", result)

}

func TestCreateIssue(t *testing.T) {

	accessToken := os.Getenv(envKeyAccessToken)
	repositoryName := os.Getenv(envKeyRepositoryName)
	organization := os.Getenv(envKeyOrganization)

	title := "Create Issue Test Title"
	description := "Create Issue Test Description."

	s := registry.NewGitService(accessToken, repositoryName, organization)
	result := s.CreateIssue(title, description)
	fmt.Println(result)

}

func TestUpdateIssue(t *testing.T) {

	accessToken := os.Getenv(envKeyAccessToken)
	repositoryName := os.Getenv(envKeyRepositoryName)
	organization := os.Getenv(envKeyOrganization)

	title := "Update Issue Test Title"
	description := "Update Issue Test Description."

	s := registry.NewGitService(accessToken, repositoryName, organization)
	result := s.UpdateIssue(10, title, description)

	fmt.Println(result)

}
