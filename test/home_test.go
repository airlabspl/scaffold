package test

import (
	"net/http"
	"scaffold/internal/testutil"
	"testing"
)

func TestHome(t *testing.T) {
	t.Run("works", func(t *testing.T) {
		tc := testutil.NewTestCase(t)
		defer tc.Close()

		tc.Get("/").
			AssertStatus(http.StatusOK)
	})
}
