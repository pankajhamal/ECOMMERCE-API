package main

import (
	"log/slog"
	"os"
)

func main(){
	cfg := config{
		addr: ":8080",
		
	}

	api := application{
		config: cfg,
	}

	//Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil{
		slog.Error("Server has failed to start", "error", err)
		// log.Printf("Server has failed to start, err: %s", err)
		os.Exit(1)
	}

}