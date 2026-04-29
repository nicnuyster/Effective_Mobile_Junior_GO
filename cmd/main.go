package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/nicnuyster/Effective_Mobile_Junior_GO/internal/env"
)

func main() {
	ctx := context.Background()

	// основной конфиг
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			dsn: env.GetString("DBSTRING", "host=localhost, user=postgres, password=postgres, dbname=efm, sslmode=disable"),
		},
	}

	// slog логгер
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// наша бд
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	logger.Info("БД соединина", "dsn", cfg.db.dsn)

	//
	api := application{
		config: cfg,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("Серверу плохо;(", "error", err)
		os.Exit(1)
	}
}
