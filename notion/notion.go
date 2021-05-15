package notion

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL   = "https://api.notion.com/v1/"
	userAgent        = "notiongo"
	defaultMediaType = "application/json"
	VERSION          = "0.1.0"
	notionVersion    = "2021-05-13"
)

// Client ...
type Client struct {
	client *http.Client // HTTP client used to communicate with the API.

	// BaseURL for API requests. Defaults to the public notion api.
	BaseURL *url.URL
	// UserAgent used for communicating with the Notion API
	UserAgent string

	AuthToken     string // AuthToken is the notion authorization token
	NotionVersion string // NotionVersion is the current version of notion

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// TODO: services for communicating with different parts of the api
}

type service struct {
	client *Client
}

// New will create a new instance of the library
func NewClient(token string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:        httpClient,
		BaseURL:       baseURL,
		UserAgent:     userAgent,
		AuthToken:     token,
		NotionVersion: notionVersion,
	}

	c.common.client = c
	return c
}

// NewRequest will make a new authenticated request to the notion api
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	if c.AuthToken != "" {
		return nil, errors.New("no authorization token provided")
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AuthToken))

	return req, nil
}
