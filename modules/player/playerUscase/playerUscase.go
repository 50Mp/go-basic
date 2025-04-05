package playeruscase

import playerrepository "github.com/50Mph/go-api/modules/player/playerRepository"

type (
	// PlayerUsecase represents the player usecase interface.
	PlayerUsecase interface {
		// Define the methods for the player usecase here
	}
	// playerUsecase is the concrete implementation of PlayerUsecase.
	playerUsecase struct {
		playerRepository playerrepository.PlayerRepository
	}
)

func NewPlayerUsecase(playerRepository playerrepository.PlayerRepository) PlayerUsecase {
	return playerUsecase{playerRepository: playerRepository}
}

func (u *playerUsecase) authConn() playerrepository.PlayerRepository {
	return u.playerRepository
}
