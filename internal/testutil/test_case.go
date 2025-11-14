package testutil

import (
	"io"
	"net/http"
	"net/http/httptest"
	"scaffold/internal/router"
	"testing"
)

type TestCase struct {
	T      *testing.T
	Router *http.ServeMux
	Server *httptest.Server
}

func NewTestCase(t *testing.T) *TestCase {
	r := router.New()
	s := httptest.NewServer(r)

	return &TestCase{
		T:      t,
		Router: r,
		Server: s,
	}
}

func (t *TestCase) Close() {
	t.Server.Close()
}

func (tc *TestCase) Get(path string) *AssertableResponse {
	return tc.do(http.MethodGet, path, nil)
}

func (tc *TestCase) Post(path string, body io.Reader) *AssertableResponse {
	return tc.do(http.MethodPost, path, body)
}

func (tc *TestCase) do(method, path string, body io.Reader) *AssertableResponse {
	req, err := http.NewRequest(method, tc.Server.URL+path, body)
	if err != nil {
		tc.T.Fatal(err)
	}

	res, err := tc.Server.Client().Do(req)
	if err != nil {
		tc.T.Fatal(err)
	}

	return NewAssertableResponse(res, tc.T)
}
