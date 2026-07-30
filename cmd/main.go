package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/pankajhamal/ECOMMERCE-API/internal/env"
)

func main(){
	ctx := context.Background()
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", 
			"host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}
	fmt.Println("Environment variable", cfg.db.dsn)

		//Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	//Database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err !=nil {
		panic(err)
	}
	defer conn.Close(ctx)

	logger.Info("Connected to database", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
		db: conn,
	}



	if err := api.run(api.mount()); err != nil{
		slog.Error("Server has failed to start", "error", err)
		// log.Printf("Server has failed to start, err: %s", err)
		os.Exit(1)
	}

}