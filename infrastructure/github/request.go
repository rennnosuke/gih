package github

type IssueCreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"body"`
}

type IssueUpdateRequest IssueCreateRequest

type IssueCloseRequest struct {
	State string `json:"state"`
}
