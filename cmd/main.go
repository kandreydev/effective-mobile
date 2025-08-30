package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/kandreydev/effective-mobile/internal/app"
)

func main() {
	if err := app.Run(context.Background()); err != nil {
		slog.Error("failed to run application", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
