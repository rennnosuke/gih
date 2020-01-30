package cmd

import (
	"fmt"
	"github.com/rennnosuke/gih/domain/model/entity"
	"github.com/rennnosuke/gih/domain/service/git/issue"
	"github.com/urfave/cli/v2"
	"regexp"
	"strconv"
	"unicode/utf8"
)

const (
	titleMaxLength       = 35
	descriptionMaxLength = 50
)

var formatPrintIssue = "%-10v%-" + strconv.Itoa(titleMaxLength) + "s%-" + strconv.Itoa(descriptionMaxLength) + "s%-10s%-10s\n"

var newLineRegex = regexp.MustCompile(`\r?\n`)

func listIssues(s *issue.GitIssueService) error {
	issues, err := s.GetIssues()
	if err != nil {
		return err
	}
	fmt.Printf(formatPrintIssue, "ISSUEID", "TITLE", "DESCRIPTION", "STATE", "CREATED_AT")
	for _, i := range *issues {
		title := trim(i.Title, titleMaxLength, "...")
		description := trim(i.Description, descriptionMaxLength, "...")
		fmt.Printf(formatPrintIssue, i.Number, title, description, "", "")
	}
	return nil
}

func createIssue(service *issue.GitIssueService, context *cli.Context) error {
	title := context.Args().Get(0)
	description := context.Args().Get(1)
	issue, err := service.CreateIssue(title, description)
	if err != nil {
		return err
	}
	printIssue(issue, "create")
	return nil
}

func updateIssue(service *issue.GitIssueService, context *cli.Context) error {
	issueNumberStr := context.Args().Get(0)
	issueNumber, err := strconv.Atoi(issueNumberStr)
	if err != nil {
		return err
	}
	title := context.Args().Get(1)
	description := context.Args().Get(2)
	issue, err := service.UpdateIssue(issueNumber, title, description)
	if err != nil {
		return err
	}
	printIssue(issue, "update")
	return nil
}

func closeIssue(service *issue.GitIssueService, context *cli.Context) error {
	sid := context.Args().Get(0)
	id, err := strconv.Atoi(sid)
	if err != nil {
		return err
	}
	issue, err := service.CloseIssue(id)
	if err != nil {
		return err
	}
	printIssue(issue, "delete")
	return nil
}

func printIssue(issue *entity.Issue, prefix string) {
	fmt.Printf("%s issue : %d\n[TITLE]\n%s\n\n[DESCRIPTION]\n%s\n", prefix, issue.Number, issue.Title, issue.Description)
}

func trim(s string, maxLength int, omitPostFix string) string {
	s = newLineRegex.ReplaceAllString(s, " ")
	if len(s) > maxLength-len(omitPostFix) {
		return byteCountRoundedByRune(s, maxLength-len(omitPostFix)) + omitPostFix
	}
	return s
}

func byteCountRoundedByRune(s string, nByte int) string {

	if nByte > len(s) {
		nByte = len(s)
	}

	lastRuneCount := utf8.RuneCountInString(s[:nByte])

	for i := nByte - 1; i >= 0; i-- {
		if c := utf8.RuneCountInString(s[:i]); c > lastRuneCount {
			return s[:i+1]
		} else {
			lastRuneCount = c
		}
	}

	return ""
}
