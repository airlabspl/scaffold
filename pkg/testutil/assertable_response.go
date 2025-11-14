package testutil

import (
	"net/http"
	"testing"
)

type AssertableResponse struct {
	Response *http.Response
	T        *testing.T
}

func NewAssertableResponse(response *http.Response, t *testing.T) *AssertableResponse {
	return &AssertableResponse{
		Response: response,
		T:        t,
	}
}

func (ar *AssertableResponse) AssertStatus(expected int) *AssertableResponse {
	if ar.Response.StatusCode != expected {
		ar.T.Errorf("expected status %d, got %d", expected, ar.Response.StatusCode)
	}

	return ar
}

func (ar *AssertableResponse) AssertUrl(expected string) *AssertableResponse {
	if ar.Response.Request.URL.Path != expected {
		ar.T.Errorf("expected url %s, got %s", expected, ar.Response.Request.URL.Path)
	}

	return ar
}
