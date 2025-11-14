package handler

import (
	"net/http"
	"scaffold/html"
)

type loginData struct {
	Title string
}

func Login() http.HandlerFunc {
	t := html.New("base_layout.html", "login.html")

	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, loginData{
			Title: "Login",
		})
	}
}
