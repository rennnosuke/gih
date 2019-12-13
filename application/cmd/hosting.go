package cmd

import (
	"fmt"
	"regexp"
)

type GitHost int

const (
	Github GitHost = iota
)

const (
	RegexGithubRepositoryUrl = `^https?://github.com/(.+)/(.+)$`
)

func ParseRepositoryUrl(path string) (GitHost, string, string, error) {
	r := regexp.MustCompile(RegexGithubRepositoryUrl)
	if r.MatchString(path) {
		s := r.FindStringSubmatch(path)
		org, repo := s[1], s[2]
		return Github, org, repo, nil
	}
	return Github, "", "", fmt.Errorf("Failed to parse url: %s\n", path)
}
