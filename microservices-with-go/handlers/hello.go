package handlers

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type Hello struct {
	l *slog.Logger
}

func NewHello(s *slog.Logger) *Hello {
	return &Hello{s}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Info(r.RemoteAddr)
	data, err := io.ReadAll(r.Body)
	if err != nil {
		h.l.Error(err.Error())
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write([]byte("Sad"))
		http.Error(rw, "Failed", http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Hello %s", data)
}
