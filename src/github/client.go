package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/urlesc"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	baseUrl     = "https://api.github.com"
	accept      = "application/vnd.github.v3+json"
	contentType = "application/json"
	timeout     = 30 // second
)

type Client struct {
	AccessToken    string
	RepositoryName string
	Organization   string
}

type IssueCreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (c *Client) SearchIssues(searchWords []string) *[]byte {
	params := make(map[string]string)
	params["q"] = strings.Join(searchWords, " ")
	return c.request("GET", c.endpoint(), &params, nil)
}

func (c *Client) CreateIssue(title, description string) *[]byte {
	req := IssueCreateRequest{title, description}
	body, err := json.Marshal(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Falied to create JSON request body. %#v", req)
		os.Exit(1)
	}
	return c.request("POST", c.endpoint(), nil, body)
}

func (c *Client) request(method string, endpoint string, params *map[string]string, body []byte) *[]byte {

	url := endpoint
	queryParams := joinParams(params)
	if queryParams != "" {
		url += "?" + queryParams
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create new request : method=%s, url=%s, params=%v, body=%v\n", method, endpoint, params, body)
		os.Exit(1)
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", accept)
	req.Header.Set("Authorization", fmt.Sprintf("token %s", c.AccessToken))

	httpClient := http.Client{Timeout: time.Duration(timeout) * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to http request : method=%s, url=%s, params=%v, body=%v\n", method, endpoint, params, body)
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
		fmt.Printf("Failed to request : %d\nResponse body : \n%s \n", resp.StatusCode, string(formatted.Bytes()))
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

func (c *Client) endpoint() string {
	return fmt.Sprintf("%s/repos/%s/%s/issues", baseUrl, c.Organization, c.RepositoryName)
}

func isHttpStatusOK(resp *http.Response) bool {
	return resp.StatusCode == http.StatusOK ||
		resp.StatusCode == http.StatusCreated ||
		resp.StatusCode == http.StatusAccepted
}
