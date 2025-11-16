package testutil

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
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

func (tc *TestCase) Post(path string, values url.Values) *AssertableResponse {
	return tc.do(http.MethodPost, path, values)
}

func (tc *TestCase) do(method, path string, values url.Values) *AssertableResponse {
	body := bytes.NewBufferString(values.Encode())

	req, err := http.NewRequest(method, tc.Server.URL+path, body)
	if err != nil {
		tc.T.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := tc.Server.Client().Do(req)
	if err != nil {
		tc.T.Fatal(err)
	}

	return NewAssertableResponse(res, tc.T)
}
