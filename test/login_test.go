package test

import (
	"net/http"
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
}
