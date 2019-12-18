package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gih/domain/model/entity"
	"gih/domain/repository"
	"github.com/PuerkitoBio/urlesc"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	baseUrl     = "https://api.github.com"
	accept      = "application/vnd.github.v3+json"
	contentType = "application/json"
	timeout     = 30 // second
)

type RepositoryImpl struct {
	AccessToken    string
	RepositoryName string
	Organization   string
}

func (c *RepositoryImpl) GetIssue(issueId int) *entity.Issue {
	b := c.request("GET", c.endpoint(strconv.Itoa(issueId)), nil, nil)
	res := convertToIssueResponse(b)
	return &entity.Issue{Number: res.Number, Title: res.Title, Description: res.Body}
}

func (c *RepositoryImpl) GetIssues() *[]entity.Issue {
	b := c.request("GET", c.endpoint(), nil, nil)
	res := convertToIssueResponses(b)

	var issues []entity.Issue
	for _, i := range *res {
		issues = append(issues, entity.Issue{Number: i.Number, Title: i.Title, Description: i.Body})
	}
	return &issues
}

func (c *RepositoryImpl) CreateIssue(r *repository.IssueCreateRequest) *entity.Issue {
	req := IssueCreateRequest{r.Title, r.Description}
	body, err := json.Marshal(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Falied to create JSON request body. %#v\n", req)
		os.Exit(1)
	}
	b := c.request("POST", c.endpoint(), nil, body)
	res := convertToIssueResponse(b)
	return &entity.Issue{Number: res.Number, Title: res.Title, Description: res.Body}
}

func (c *RepositoryImpl) UpdateIssue(r *repository.IssueUpdateRequest) *entity.Issue {
	req := IssueUpdateRequest{r.Title, r.Description}
	body, err := json.Marshal(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Falied to update JSON request body. %#v", req)
		os.Exit(1)
	}
	b := c.request("PATCH", c.endpoint(strconv.Itoa(r.IssueId)), nil, body)
	res := convertToIssueResponse(b)
	return &entity.Issue{Number: res.Number, Title: res.Title, Description: res.Body}
}

func (c *RepositoryImpl) CloseIssue(id int) *entity.Issue {
	req := IssueCloseRequest{State: "closed"}
	body, err := json.Marshal(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Falied to update JSON request body. %#v", req)
		os.Exit(1)
	}
	b := c.request("PATCH", c.endpoint(strconv.Itoa(id)), nil, body)
	res := convertToIssueResponse(b)
	return &entity.Issue{Number: res.Number, Title: res.Title, Description: res.Body}
}

func (c *RepositoryImpl) request(method string, endpoint string, params *map[string]string, body []byte) *[]byte {

	url := endpoint
	queryParams := joinParams(params)
	if queryParams != "" {
		url += "?" + queryParams
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create new request : method=%s, url=%s, params=%v, body=%s\n", method, endpoint, params, string(body))
		os.Exit(1)
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", accept)
	req.Header.Set("Authorization", fmt.Sprintf("token %s", c.AccessToken))

	httpClient := http.Client{Timeout: time.Duration(timeout) * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute http request : method=%s, url=%s, params=%v, body=%v\n", method, endpoint, params, body)
		os.Exit(1)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read respose body : %v\n", resp.Body)
		os.Exit(1)
	}

	if !isHttpStatusOK(resp) {
		var formatted bytes.Buffer
		err = json.Indent(&formatted, respBody, "", "\t")
		fmt.Printf("Failed to request : %d\nEndpoint :%s\nResponse body : \n%s \n", resp.StatusCode, url, string(formatted.Bytes()))
		os.Exit(1)
	}

	return &respBody
}

func joinParams(params *map[string]string) string {
	if params == nil {
		return ""
	}
	var query string
	for k, v := range *params {
		query += fmt.Sprintf("%s=%s&", k, v)
	}
	return urlesc.QueryEscape(query[:len(query)-1])
}

func (c *RepositoryImpl) endpoint(s ...string) string {
	return baseUrl + path.Join("/repos", c.Organization, c.RepositoryName, "issues", path.Join(s...))
}

func isHttpStatusOK(resp *http.Response) bool {
	return resp.StatusCode == http.StatusOK ||
		resp.StatusCode == http.StatusCreated ||
		resp.StatusCode == http.StatusAccepted
}
