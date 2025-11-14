package handler

import "net/http"

func Dashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Dashboard"))
	}
}
