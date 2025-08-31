package app

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/kandreydev/effective-mobile/internal/config"
)
func Run(ctx context.Context) error {
	// CONFIG
	cfg := config.MustLoad()
	// LOGS
	log := setupLogger(cfg)
	_ = log
	_ = cfg

	return nil
}
func setupLogger(cfg *config.Config) *slog.Logger {
	//TODO: use config to setup
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	log := slog.New(slog.NewJSONHandler(file, nil))
	return log
}