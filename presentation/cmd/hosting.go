package cmd

import (
	"fmt"
	"regexp"
)

type GitHost int

const (
	GitHub GitHost = iota + 1
)

func (g *GitHost) String() string {
	switch *g {
	case GitHub:
		return "GitHub"
	default:
		return ""
	}
}

const (
	RegexGithubRepositoryUrl = `^https?://github.com/(.+)/(.+)$`
)

func parseRepositoryUrl(path string) (GitHost, string, string, error) {
	r := regexp.MustCompile(RegexGithubRepositoryUrl)
	if r.MatchString(path) {
		s := r.FindStringSubmatch(path)
		org, repo := s[1], s[2]
		return GitHub, org, repo, nil
	}
	return GitHub, "", "", fmt.Errorf("Failed to parse url: %s\n", path)
}
