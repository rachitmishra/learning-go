package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"rachitmishra.com/learning-go/microservices-with-go/handlers"
)

func server() {
	l := slog.New(slog.Default().Handler())
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye()
	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	sm.Handle("/products", ph)
	sm.Handle("/product/:id", ph)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		l.Info(":9090")
		if err != nil {
			l.Error(err.Error())
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	<-sigChan

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

func main() {
	server()
}
