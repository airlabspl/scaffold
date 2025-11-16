package testutil

import (
	"net/http"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

type AssertableResponse struct {
	Response *http.Response
	Document *goquery.Document
	T        *testing.T
}

func NewAssertableResponse(res *http.Response, t *testing.T) *AssertableResponse {
	doc, _ := goquery.NewDocumentFromReader(res.Body)

	return &AssertableResponse{
		Response: res,
		Document: doc,
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

func (ar *AssertableResponse) AssertVisible(selector string) *AssertableResponse {
	if ar.Document.Find(selector).Length() == 0 {
		ar.T.Errorf("element %s does not exist", selector)
	}

	return ar
}

func (ar *AssertableResponse) AssertText(text string) *AssertableResponse {
	if !strings.Contains(ar.Document.Text(), text) {
		ar.T.Errorf("expected to see %s, got %s", text, ar.Document.Text())
	}

	return ar
}
