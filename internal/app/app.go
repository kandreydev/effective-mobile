package app

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kandreydev/effective-mobile/database"
	"github.com/kandreydev/effective-mobile/internal/config"
	"github.com/kandreydev/effective-mobile/internal/handle"
	"github.com/kandreydev/effective-mobile/internal/repository"
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

	// REPOSITORY
	repo := repository.NewSubscriptionsRepo(pool)

	// HANDLER
	h := handle.New(repo, log)

	// ROUTER
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	api := router.Group("/api/v1/")
	{
		api.GET("subscriptions", h.ListSubscriptions)
		api.POST("/subscriptions", h.CreateSubscription)
		api.GET("/subscriptions/:id", h.GetSubscription)
		api.PUT("/subscriptions/:id", h.UpdateSubscription)
		api.DELETE("/subscriptions/:id", h.DeleteSubscription)

		api.GET("/billing/total", h.CalculateTotalCost)
	}
	// GET PORT via flag
	var port string
	flag.StringVar(&port, "port", "8080", "port for server")
	flag.Parse()
	log.Info("starting server", "port", port)

	// START SERVER 
	// TODO :stick to Graceful Shutdown pattern

	if err := router.Run(":" + port); err != nil {
		log.Error("failed to start server", slog.String("error", err.Error()))

		return errors.Wrap(err, "server failed")
	}

	return nil
}

func setupLogger(_ *config.Config) *slog.Logger {
	// TODO: use cfg to setup
	_, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}

	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	return log
}
