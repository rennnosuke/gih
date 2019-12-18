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

func (c *RepositoryImpl) GetIssue(issueId int) (*entity.Issue, error) {
	b, err := c.request("GET", c.endpoint(strconv.Itoa(issueId)), nil, nil)
	if err != nil {
		return nil, fmt.Errorf("GetIssue : %s", err)
	}
	res := convertToIssueResponse(b)
	return &entity.Issue{Number: res.Number, Title: res.Title, Description: res.Body}, nil
}

func (c *RepositoryImpl) GetIssues() (*[]entity.Issue, error) {
	b, err := c.request("GET", c.endpoint(), nil, nil)
	if err != nil {
		return nil, fmt.Errorf("GetIssues : %s", err)
	}
	res := convertToIssueResponses(b)
	var issues []entity.Issue
	for _, i := range *res {
		issues = append(issues, entity.Issue{Number: i.Number, Title: i.Title, Description: i.Body})
	}
	return &issues, nil
}

func (c *RepositoryImpl) CreateIssue(r *repository.IssueCreateRequest) (*entity.Issue, error) {
	req := IssueCreateRequest{r.Title, r.Description}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("CreateIssue : %s", err)
	}
	b, err := c.request("POST", c.endpoint(), nil, body)
	if err != nil {
		return nil, fmt.Errorf("CreateIssue : %s", err)
	}
	res := convertToIssueResponse(b)
	return &entity.Issue{Number: res.Number, Title: res.Title, Description: res.Body}, nil
}

func (c *RepositoryImpl) UpdateIssue(r *repository.IssueUpdateRequest) (*entity.Issue, error) {
	req := IssueUpdateRequest{r.Title, r.Description}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("UpdateIssue : %s", err)
	}
	b, err := c.request("PATCH", c.endpoint(strconv.Itoa(r.IssueId)), nil, body)
	if err != nil {
		return nil, fmt.Errorf("UpdateIssue : %s", err)
	}
	res := convertToIssueResponse(b)
	return &entity.Issue{Number: res.Number, Title: res.Title, Description: res.Body}, nil
}

func (c *RepositoryImpl) CloseIssue(id int) (*entity.Issue, error) {
	req := IssueCloseRequest{State: "closed"}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("CloseIssue : %s", err)
	}
	b, err := c.request("PATCH", c.endpoint(strconv.Itoa(id)), nil, body)
	if err != nil {
		return nil, fmt.Errorf("CloseIssue : %s", err)
	}
	res := convertToIssueResponse(b)
	return &entity.Issue{Number: res.Number, Title: res.Title, Description: res.Body}, nil
}

func (c *RepositoryImpl) request(method string, endpoint string, params *map[string]string, body []byte) (*[]byte, error) {

	url := endpoint
	queryParams := joinParams(params)
	if queryParams != "" {
		url += "?" + queryParams
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("request ( endpoint : %s ) - failed to create request : %s", endpoint, err)
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", accept)
	req.Header.Set("Authorization", fmt.Sprintf("token %s", c.AccessToken))

	httpClient := http.Client{Timeout: time.Duration(timeout) * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request ( endpoint : %s ) - request fail : %s", endpoint, err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("request ( endpoint : %s ) - failed to read response body : %s", endpoint, err)
	}

	if !isHttpStatusOK(resp) {
		return nil, fmt.Errorf("request ( endpoint : %s ) - failed to request [%d %s] : %s", endpoint, resp.StatusCode, resp.Status, err)
	}

	return &respBody, nil
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
