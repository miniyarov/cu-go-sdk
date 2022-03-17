package api_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"cryptounifier-go-sdk/pkg/api"
)

type mockClient struct {
	api.Client
	resp string
}

func newMockClient(response string) api.Client {
	return &mockClient{resp: response}
}

func (m *mockClient) DoRequest(r *http.Request) (interface{}, error) {
	var data interface{}
	err := json.Unmarshal([]byte(m.resp), &data)
	return data, err
}

func (m *mockClient) NewRequest(method string, path string, headers map[string]string, query map[string]string, body io.Reader) (*http.Request, error) {
	u, _ := url.Parse(path)
	return &http.Request{
		Method: method,
		URL:    u,
	}, nil
}
