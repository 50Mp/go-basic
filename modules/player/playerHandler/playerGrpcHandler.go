package playerhandler

import playeruscase "github.com/50Mph/go-api/modules/player/playerUscase"

type (
	playerGrpchandler struct {
		playerUscase playeruscase.PlayerUsecase
	}
)

func NewPlayerGrpcHandler(playerUscase playeruscase.PlayerUsecase) playeruscase.PlayerUsecase {
	return playerGrpchandler{
		playerUscase: playerUscase,
	}
}
