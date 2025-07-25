package db

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitConnDb(log *slog.Logger) (*pgxpool.Pool, context.Context) {
	// postgres://authuser:authpass@postgres:5432/authdb?sslmode=disable
	// conn, err := pgx.Connect(ctx, "postgres://psql_limon4ik_user:psql_limon4ik_password@localhost:5434/db?sslmode=disable")

	ctx := context.Background()
	connStr := "postgres://psql_limon4ik_user:psql_limon4ik_password@localhost:5434/db?sslmode=disable"

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Error("failed to parse pool config", "error", err)
		os.Exit(1)
	}

	config.MaxConns = 20
	config.MaxConnLifetime = 5 * time.Minute

	conn, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Error("failed to create pool connection", "error", err)
		os.Exit(1)
	}

	return conn, ctx
}
