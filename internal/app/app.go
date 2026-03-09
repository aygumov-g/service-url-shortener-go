package app

import (
	"context"

	"github.com/aygumov-g/service-url-shortener-go/internal/config"
	"github.com/aygumov-g/service-url-shortener-go/internal/infrastructure/db"
	"github.com/aygumov-g/service-url-shortener-go/internal/transport/http/server"
	"github.com/aygumov-g/service-url-shortener-go/pkg/logger"
)

type App struct {
	logger logger.Logger
	server *server.Server
	db     *db.Storage
}

func NewApp(ctx context.Context) (*App, error) {
	cfg := config.Load()
	log := logger.New()

	db, err := db.New(cfg.DB)
	if err != nil {
		return nil, err
	}

	server, err := buildHTTP(cfg, db)
	if err != nil {
		return nil, err
	}

	app := &App{
		logger: log,
		server: server,
		db:     db,
	}

	return app, nil
}

func (a *App) Run() {
	a.logger.Info("http server started", "addr", a.server.Addr())

	if err := a.server.Start(); err != nil {
		a.logger.Error("http server failed started", "error", err)
	}
}

func (a *App) Shutdown(ctx context.Context) {
	a.logger.Info("shutdown started")

	err := a.server.Shutdown(ctx)
	if err != nil {
		a.logger.Error("http server failed shutdown", "error", err)
	}

	err = a.db.Close()
	if err != nil {
		a.logger.Error("db failed shutdown", "error", err)
	}

	a.logger.Info("shutdown completed")
}
