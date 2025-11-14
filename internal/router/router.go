package router

import (
	"net/http"
	"scaffold/internal/handler"
	"scaffold/internal/middleware"
	"scaffold/static"
)

func New() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/", handler.Home())

	a := http.NewServeMux()
	{
		a.Handle("/", handler.Dashboard())
	}

	r.Handle("/app", middleware.RequireAuthenticatedUser(a))
	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(static.FS))))

	return r
}
