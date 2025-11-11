package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"scaffold/internal/router"
	"syscall"
)

func main() {
	s := http.Server{
		Addr:    ":4000",
		Handler: router.New(),
	}

	go s.ListenAndServe()

	notify := make(chan os.Signal, 1)
	signal.Notify(notify, os.Interrupt, syscall.SIGTERM)

	<-notify

	s.Shutdown(context.Background())
}
