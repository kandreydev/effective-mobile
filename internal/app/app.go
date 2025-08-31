package app

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/kandreydev/effective-mobile/database"
	"github.com/kandreydev/effective-mobile/internal/config"
	"github.com/kandreydev/effective-mobile/internal/storage"
	"github.com/pkg/errors"
)

func Run(ctx context.Context) error {
	// CONFIG
	cfg := config.MustLoad()
	// LOG
	log := setupLogger(cfg)
	// MIGRATIONS
	if err := database.Migrate(cfg.DSN()); err != nil {
		log.Error("failed to run migrations", slog.String("error", err.Error()))

		return errors.Wrap(err, "migrations failed")
	}
	// STORAGE CONNECTION
	pool, err := storage.GetConn(ctx, cfg.DSN())
	if err != nil {
		log.Error("failed to connect to storage", slog.String("error", err.Error()))

		return errors.Wrap(err, "storage connection failed")
	}

	defer pool.Close()

	return nil
}

func setupLogger(_ *config.Config) *slog.Logger {
	// TODO: use cfg to setup
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}

	log := slog.New(slog.NewJSONHandler(file, nil))

	return log
}
