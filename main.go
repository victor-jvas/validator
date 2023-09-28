package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"validator/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	hg := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", hg)

	s := http.Server{
		Addr:        ":9090",
		Handler:     sm,
		IdleTimeout: 10 * time.Second,
	}

	err := s.ListenAndServe()
	if err != nil {
		l.Println("Error initiating server %s\n", err)
		os.Exit(1)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, gracefully shutdown commence ", sig)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
	s.Shutdown(ctx)
}
