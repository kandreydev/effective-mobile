package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

func GetConn(ctx context.Context, connStr string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "create pgxpool")
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, errors.Wrap(err, "Unable to ping database")
	}

	return pool, nil
}
