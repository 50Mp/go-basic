package playerhandler

import (
	"github.com/50Mph/go-api/config"
	playeruscase "github.com/50Mph/go-api/modules/player/playerUscase"
)

type (
	PlayerHandler struct {
		// playerUsecase is the usecase for player operations.
		cfg           config.Config
		playerUsecase playeruscase.PlayerUsecase
	}
)

// NewPlayerHandler creates a new PlayerHandler
func NewPlayerHandler(cfg config.Config, playerUsecase playeruscase.PlayerUsecase) PlayerHandler {
	return PlayerHandler{
		cfg:           cfg,
		playerUsecase: playerUsecase,
	}
}

// NewQueryHandler creates a new PlayerQuery handler
func (h *PlayerHandler) NewQueryHandler(cfg config.Config, NewQueryHandler playeruscase.PlayerUsecase) PlayerHandler {
	return PlayerHandler{
		cfg:           cfg,
		playerUsecase: h.playerUsecase,
	}
}
