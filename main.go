package main

import (
	"context"
	"log"
	"os"

	"github.com/50Mph/go-api/config"
	"github.com/50Mph/go-api/pkg/database"
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
	db := database.DbConn(ctx, &cfg)

	log.Print(db)
}
