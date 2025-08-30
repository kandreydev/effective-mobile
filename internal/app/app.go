package app

import (
	"context"

	"github.com/kandreydev/effective-mobile/internal/config"
	"github.com/pkg/errors"
)

func Run(ctx context.Context) error {
	cfg, err := config.Load()
	if err != nil {
		return errors.Wrap(err, "Init config")
	}

	_ = cfg

	return nil
}
