package application

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"net/http"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func NewApp() *App {
	app := &App{
		router: loadRoutes(),
		rdb:    redis.NewClient(&redis.Options{}),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
