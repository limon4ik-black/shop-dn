package db

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func InitConnDb(log *slog.Logger) (*pgx.Conn, context.Context) {
	// postgres://authuser:authpass@postgres:5432/authdb?sslmode=disable
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://authuser:authpass@postgres:5432/authdb?sslmode=disable")
	if err != nil {
		log.Error("failed to connection: db", "error", err)
		os.Exit(1)
	}

	return conn, ctx
}
