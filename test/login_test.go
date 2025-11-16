package test

import (
	"net/http"
	"net/url"
	"scaffold/pkg/testutil"
	"testing"
)

func TestLogin(t *testing.T) {
	t.Run("page works", func(t *testing.T) {
		tc := testutil.NewTestCase(t)
		defer tc.Close()

		tc.Get("/login").
			AssertStatus(http.StatusOK)
	})

	t.Run("form is visible", func(t *testing.T) {
		tc := testutil.NewTestCase(t)
		defer tc.Close()

		tc.Get("/login").
			AssertVisible(`form[hx-post="/login"]`).
			AssertVisible(`input[name="email"]`)
	})

	t.Run("email is required", func(t *testing.T) {
		tc := testutil.NewTestCase(t)
		defer tc.Close()

		tc.Post("/login", nil).
			AssertStatus(http.StatusBadRequest).
			AssertText("Email is required")
	})

	t.Run("email has to be valid", func(t *testing.T) {
		tc := testutil.NewTestCase(t)
		defer tc.Close()

		tc.Post("/login", url.Values{
			"email": {"invalid"},
		}).
			AssertStatus(http.StatusBadRequest).
			AssertText("Email must be a valid email address")
	})
}
