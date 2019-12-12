package github

import (
	"encoding/json"
	"fmt"
	"os"
)

type IssueResponse struct {
	Id    int    `json:"id"`
	State string `json:"state"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func convertToIssueResponse(b *[]byte) *IssueResponse {
	res := IssueResponse{}
	err := json.Unmarshal(*b, &res)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Falied to parse JSON response body. %#v\n", string(*b))
		os.Exit(1)
	}
	return &res
}

func convertToIssueResponses(b *[]byte) *[]IssueResponse {
	var res []IssueResponse
	err := json.Unmarshal(*b, &res)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Falied to parse JSON response body. %#v\n", string(*b))
		os.Exit(1)
	}
	return &res
}
