package cmd

import (
	"fmt"
	"gih/domain/model/entity"
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

func listIssues(issues *[]entity.Issue) {
	fmt.Printf(formatPrintIssue, "ISSUEID", "TITLE", "DESCRIPTION", "STATE", "CREATED_AT")
	for _, i := range *issues {
		fmt.Printf(formatPrintIssue, i.ID, trimTitle(i.Title), trimDescription(i.Description), "", "")
	}
}

func printIssue(issue *entity.Issue, prefix string) {
	fmt.Printf("%s issue : %d\n[TITLE]\n%s\n\n[DESCRIPTION]\n%s\n", prefix, issue.ID,issue.Title,issue.Description)
}

func trimTitle(title string) string {
	return trim(title, titleMaxLength, "...")
}

func trimDescription(description string) string {
	return trim(description, descriptionMaxLength, "...")
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
