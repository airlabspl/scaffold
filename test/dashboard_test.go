package test

import (
	"scaffold/pkg/testutil"
	"testing"
)

func TestDashboard(t *testing.T) {
	t.Run("redirects guests to login page", func(t *testing.T) {
		tc := testutil.NewTestCase(t)
		defer tc.Close()

		tc.Get("/app").
			AssertUrl("/login")
	})
}
