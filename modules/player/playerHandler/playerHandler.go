package playerhandler

import playeruscase "github.com/50Mph/go-api/modules/player/playerUscase"

type (
	PlayerHandler struct {
		// playerUsecase is the usecase for player operations.
		playerUsecase playeruscase.PlayerUsecase
	}
)

func NewPlayerHandler(playerUsecase playeruscase.PlayerUsecase) PlayerHandler {
	return PlayerHandler{
		playerUsecase: playerUsecase,
	}
}
func (h *PlayerHandler) authConn() playeruscase.PlayerUsecase {
	return h.playerUsecase
}
