package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"subscriptions/Backend/core/config"
	"subscriptions/Backend/core/db"
	httpx "subscriptions/Backend/core/http"
)

type App struct {
	cfg *config.Config
	srv *httpx.Server
}

func main() {
	app := NewApp()
	app.Run()
}

func NewApp() *App {

	cfg := mustLoadConfig()
	initLogger(cfg)

	initCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := db.InitPool(initCtx, cfg.DB.URL, cfg.DB.MaxConns); err != nil {
		slog.Error("db connect failed", "error", err)
		os.Exit(1)
	}

	router := httpx.ConstructRouter()
	srv := httpx.ConstructServer(cfg.HTTP.Addr, router)

	return &App{cfg: cfg, srv: srv}
}

func (a *App) Run() {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := a.srv.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("http server error", "error", err)
			stop()
		}
	}()

	<-ctx.Done()
	a.shutdown()
}

func (a *App) shutdown() {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("shutdown error", "error", err)
	}

	db.ClosePool()
}

func mustLoadConfig() *config.Config {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("config load failed", "error", err)
		os.Exit(1)
	}
	return cfg
}

func initLogger(cfg *config.Config) {
	logger := config.NewLogger(cfg.Log.Level)
	slog.SetDefault(logger)
}
