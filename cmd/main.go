package main

import (
	"context"
	"log"

	"github.com/kandreydev/effective-mobile/internal/app"
)

func main() {
	if err := app.Run(context.Background()); err != nil {
		log.Fatalf("failed to run application: %v", err)
	}
}
