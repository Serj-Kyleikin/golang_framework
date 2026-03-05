package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	once sync.Once
	pool *pgxpool.Pool
	err  error
)

func InitPool(ctx context.Context, url string, maxConns int32) (*pgxpool.Pool, error) {
	once.Do(func() {
		cfg, e := pgxpool.ParseConfig(url)
		if e != nil {
			err = fmt.Errorf("parse pgxpool config: %w", e)
			return
		}

		if maxConns > 0 {
			cfg.MaxConns = maxConns
		}

		pool, err = pgxpool.NewWithConfig(ctx, cfg)
		if err != nil {
			err = fmt.Errorf("create pgxpool: %w", err)
			return
		}

		if e := pool.Ping(ctx); e != nil {
			pool.Close()
			pool = nil
			err = fmt.Errorf("db ping failed: %w", e)
			return
		}
	})

	return pool, err
}

func MustPool() *pgxpool.Pool {
	if pool == nil {
		panic("db pool is not initialized: call db.InitPool() on app startup")
	}
	return pool
}

func ClosePool() {
	if pool != nil {
		pool.Close()
	}
}
