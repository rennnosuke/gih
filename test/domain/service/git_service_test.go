package service

import (
	"fmt"
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

	s := registry.NewGitService(accessToken, repositoryName, organization)

	result, err := s.GetIssue(10)
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

	s := registry.NewGitService(accessToken, repositoryName, organization)

	result, err := s.GetIssues()
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

	title := "Create Issue Test Title"
	description := "Create Issue Test Description."

	s := registry.NewGitService(accessToken, repositoryName, organization)

	result, err := s.CreateIssue(title, description)
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

	title := "Update Issue Test Title"
	description := "Update Issue Test Description."

	s := registry.NewGitService(accessToken, repositoryName, organization)

	result, err := s.UpdateIssue(10, title, description)
	if err != nil {
		fmt.Printf("%s\n", err)
		t.Fail()
	}

	fmt.Println(result)

}
