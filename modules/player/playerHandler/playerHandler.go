package playerhandler

import (
	"github.com/50Mph/go-api/config"
	playeruscase "github.com/50Mph/go-api/modules/player/playerUscase"
)

type (
	// PlayerHandler defines the interface for player HTTP handlers.

	PlayerHttpHandler interface{}
	playerHttpHandler struct {
		// playerUsecase is the usecase for player operations.
		cfg           config.Config
		playerUsecase playeruscase.PlayerUsecase
	}
)
// NewPlayerHandler creates a new PlayerHandler
func NewPlayerHandler(cfg config.Config, playerUsecase playeruscase.PlayerUsecase) PlayerHttpHandler {
	return playerHttpHandler{
		cfg:           cfg,
		playerUsecase: playerUsecase,
	}
}

// NewQueryHandler creates a new PlayerQuery handler
func  NewQueryHttpHandler(cfg config.Config, playerUsecase playeruscase.PlayerUsecase) PlayerHttpHandler {
	return playerHttpHandler{
		cfg:           cfg,
		playerUsecase: playerUsecase,
	}
}
