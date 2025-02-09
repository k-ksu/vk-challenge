package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client struct {
	*pgxpool.Pool
}

func NewClient(ctx context.Context, dsn string) (*Client, error) {
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.Connect: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pgxpool.Ping: %w", err)
	}

	return &Client{pool}, nil
}
