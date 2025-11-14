package router

import (
	"net/http"
	"scaffold/internal/handler"
	"scaffold/static"
)

func New() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/", handler.Home())

	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(static.FS))))

	return r
}
