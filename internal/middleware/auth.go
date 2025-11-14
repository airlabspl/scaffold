package middleware

import (
	"net/http"
	"scaffold/internal/model"
)

type userKeyType string

const userKey userKeyType = "user"

func RequireAuthenticatedUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(userKey).(*model.User)
		if !ok || user == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
