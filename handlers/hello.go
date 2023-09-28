package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle Hello requests")

	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.l.Println("Erro reading the body", err)
		http.Error(rw, "unable to read response body", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s\n", b)
}
