package handler

import (
	"net/http"
	"scaffold/html"
)

type homeData struct {
	Title string
}

func Home() http.HandlerFunc {
	t := html.New("base_layout.html", "home.html")

	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, homeData{
			Title: "Home",
		})
	}
}
