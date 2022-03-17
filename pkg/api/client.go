package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const defaultUrl = "https://sdk-endpoint.io/api/v1/"

// Client is a public sdk client interface
type Client interface {
	Do(*http.Request) (*http.Response, error)
	DoRequest(*http.Request) (interface{}, error)
	NewRequest(string, string, map[string]string, map[string]string, io.Reader) (*http.Request, error)
}

// APIClient is a client instance
type APIClient struct {
	baseUrl *url.URL
	client  *http.Client
}

// NewDefaultClient constructs a new default client
func NewDefaultClient() Client {
	baseUrl, _ := url.Parse(defaultUrl)
	return &APIClient{
		baseUrl: baseUrl,
		client:  http.DefaultClient,
	}
}

// NewClient constructs a new client with given configuration
func NewClient(c *http.Client, u string) Client {
	baseUrl, _ := url.Parse(u)
	return &APIClient{
		baseUrl: baseUrl,
		client:  c,
	}
}

// Do performs a http request
func (c *APIClient) Do(r *http.Request) (*http.Response, error) {
	return c.client.Do(r)
}

// DoRequest performs a http request and unmarshals response
func (c *APIClient) DoRequest(r *http.Request) (interface{}, error) {
	resp, err := c.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewRequest constructs a new request
func (c *APIClient) NewRequest(method string, path string, header map[string]string, query map[string]string, body io.Reader) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	u := c.baseUrl.ResolveReference(rel)
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if query != nil {
		for key, value := range query {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	if header != nil {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	return req, nil
}
