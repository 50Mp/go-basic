package playerhandler

import playeruscase "github.com/50Mph/go-api/modules/player/playerUscase"

type (
	// PlayerGrpcHandler defines the interface for player gRPC handlers.
	PlayerGrpcHandler interface {}

	// playerGrpchandler implements the PlayerGrpcHandler interface.
	playerGrpchandler struct {
		playerUscase playeruscase.PlayerUsecase
	}
)

func NewPlayerGrpcHandler(playerUscase playeruscase.PlayerUsecase) PlayerGrpcHandler {
	return playerGrpchandler{
		playerUscase: playerUscase,
	}
}
