package notion

import (
	"bytes"
	"context"
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

	// Services for communicating with different parts of the api
	common service // Reuse a single struct instead of allocating one for each service on the heap.
	Page   *PageService
	User   *UserService
	Blocks *BlocksService
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

	// Setup services
	c.Page = (*PageService)(&c.common)
	c.User = (*UserService)(&c.common)
	c.Blocks = (*BlocksService)(&c.common)

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
		buf = &bytes.Buffer{}
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

	if c.AuthToken == "" {
		return nil, errors.New("no authorization token provided")
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AuthToken))

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer interface,
// the raw response body will be written to v, without attempting to first
// decode it. If v is nil, and no error happens, the response is returned as is.
//
// The provided ctx must be non-nil, if it is nil an error is returned. If it
// is canceled or times out, ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.BareDo(ctx, req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}
	return resp, err
}

// BareDo sends an API request and lets you handle the api response. If an error
// or API Error occurs, the error will contain more information. Otherwise you
// are supposed to read and close the response's Body. If rate limit is exceeded
// and reset time is in the future, BareDo returns *RateLimitError immediately
// without making a network API call.
//
// The provided ctx must be non-nil, if it is nil an error is returned. If it is
// canceled or times out, ctx.Err() will be returned.
func (c *Client) BareDo(ctx context.Context, req *http.Request) (*Response, error) {
	// TODO: add rate limit checks
	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}

	response := newResponse(resp)

	return response, err
}

type Response struct {
	*http.Response

	// TODO: add pagination stuff and rate limiting stuff here
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	// TODO: get pagination / rate limit stuff here (check github)
	return response
}
