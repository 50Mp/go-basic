package main

import (
	"context"
	"log"
	"os"

	"github.com/50Mph/go-api/config"
	"github.com/50Mph/go-api/pkg/database"
	"github.com/50Mph/go-api/server"
)

func main() {

	ctx := context.Background()

	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		return os.Args[1]
	}())

	log.Println(cfg)

	//Database connections
	db, _ := database.DbConn(ctx, &cfg)

	defer db.Disconnect(ctx)

	// Initialize the server
	// server := server.NewServer(db, cfg)
	// server.Start()
	// Graceful shutdown

	server.Start(ctx, &cfg, db)
}
