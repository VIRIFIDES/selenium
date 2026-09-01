package selenium

import (
	"net/http"
	"testing"
)

func TestNewRequestHeaders(t *testing.T) {
	t.Parallel()

	request, err := newRequest(http.MethodPost, "http://webdriver.test/session", []byte(`{"browserName":"chrome"}`))
	if err != nil {
		t.Fatalf("newRequest() error = %v", err)
	}

	if got := request.Header.Get("Accept"); got != jsonContentType {
		t.Errorf("Accept header = %q, want %q", got, jsonContentType)
	}
	if got := request.Header.Get("Content-Type"); got != jsonContentType {
		t.Errorf("Content-Type header = %q, want %q", got, jsonContentType)
	}
}

func TestNewRequestWithoutBodyOmitsContentType(t *testing.T) {
	t.Parallel()

	request, err := newRequest(http.MethodGet, "http://webdriver.test/status", nil)
	if err != nil {
		t.Fatalf("newRequest() error = %v", err)
	}

	if got := request.Header.Get("Content-Type"); got != "" {
		t.Errorf("Content-Type header = %q, want empty", got)
	}
}
