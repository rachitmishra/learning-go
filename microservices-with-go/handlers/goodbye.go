package handlers

import (
	"log/slog"
	"net/http"
)

type Goodbye struct {
}

func NewGoodbye() *Goodbye {
	return &Goodbye{}
}

func (gb *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(300)
	slog.Info("Goodbye!")
}
