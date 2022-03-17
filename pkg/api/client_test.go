package api_test

import (
	"net/http"
	"testing"

	"cryptounifier-go-sdk/pkg/api"
)

func TestAPIClient_NewRequest(t *testing.T) {
	// Arrange
	client := api.NewDefaultClient()
	method := http.MethodGet
	path := "/merchant/invoice-hash"
	headers := map[string]string{"x-key": "$ecret-key"}
	query := map[string]string{"hash": "HASH"}

	// Act
	req, err := client.NewRequest(method, path, headers, query, nil)

	// Assert
	if err != nil {
		t.Errorf("invalid request: %v", err)
	}
	if req.URL.Path != path {
		t.Errorf("expected %s, got: %s", path, req.URL.Path)
	}
	if req.Header.Get("x-key") != "$ecret-key" {
		t.Errorf("expected %s, got: %s", "$ecret-key", req.Header.Get("x-key"))
	}
	if req.URL.Query().Get("hash") != "HASH" {
		t.Errorf("expected %s, got: %s", "HASH", req.URL.Query().Get("hash"))
	}
}
