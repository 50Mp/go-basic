package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/50Mph/go-api/config"
	middlewarehandler "github.com/50Mph/go-api/modules/middleware/middlewareHandler"
	midlerwarerepository "github.com/50Mph/go-api/modules/middleware/midlerwareRepository"
	middlerwareuscase "github.com/50Mph/go-api/modules/middleware/midlerwareUscase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	server struct {
		// Echo instance
		app *echo.Echo
		// MongoDB client
		db *mongo.Client
		// config file
		config *config.Config
		// midlware ...
		middleware middlewarehandler.MiddlerwareHandler
	}
)

// NewServer initializes a new server instance
func newMiddleware(cfg *config.Config) middlewarehandler.MiddlerwareHandler {
	repo := midlerwarerepository.NewMiddlerwareRepository()
	uscase := middlerwareuscase.NewMiddlerwareUscase(repo)
	return middlewarehandler.NewMiddlerwareHandler(*cfg, uscase)

}

// http listener
func (s *server) httpListener() error {

	if err := s.app.Start(s.config.App.Url); err != nil && err != http.ErrServerClosed {
		// Log the error and shut down the server
		s.app.Logger.Errorf("Error starting server: %v", err)

		// Close the database connection
		if err := s.db.Disconnect(context.Background()); err != nil {
			s.app.Logger.Errorf("Error disconnecting from database: %v", err)
		}
		return err
	}
	return nil
}

// Graceful shutdown
func (s *server) gracefulShutdown(pcx context.Context, quit <-chan os.Signal) {
	log.Printf("start Service: %s ", s.config.App.Name)
	<-quit

	log.Println("Shutting down service...")

	ctx, cancel := context.WithTimeout(pcx, 5*time.Second)
	defer cancel()

	// Close the database connection
	if err := s.app.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down service: %v", err)
	} else {
		log.Println("Service shut down gracefully")
	}

}

//StartServer starts the server and listens for incoming requests

func Start(pcx context.Context, cfg *config.Config, db *mongo.Client) (*server, error) {
	s := &server{
		app:        echo.New(),
		db:         db,
		config:     cfg,
		middleware: newMiddleware(cfg),
	}

	// Middleware Timeout
	s.app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Error: Request timeout",
		Timeout:      30 * time.Second,
	}))

	//body limit
	s.app.Use(middleware.BodyLimit("2M"))
	// Middleware CORS
	switch s.config.App.Name {
	case "auth":
		s.authService()
	case "player":
		s.playerService()
	case "item":
		s.itemService()
	case "inventory":
		s.inventoryService()
	case "payment":
		s.paymentService()
	}

	//graceful shutdown
	// Create a channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	s.app.Use(middleware.Logger())

	go s.gracefulShutdown(pcx, quit)

	//http listener
	s.httpListener()

	return s, nil

}
