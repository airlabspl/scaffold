package router

import (
	"net/http"
	"scaffold/internal/handler"
)

func New() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/", handler.Home())

	return r
}
