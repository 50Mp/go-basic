package server

import (
	playerhandler "github.com/50Mph/go-api/modules/player/playerHandler"
	playerrepository "github.com/50Mph/go-api/modules/player/playerRepository"
	playeruscase "github.com/50Mph/go-api/modules/player/playerUscase"
)

func (s *server) playerService() {
	playerRepository := playerrepository.NewPlayerRepository(s.db)
	playerUscase := playeruscase.NewPlayerUsecase(playerRepository)
	playerHandler := playerhandler.NewPlayerHandler(*s.config, playerUscase)
	playerGrpchandler := playerhandler.NewPlayerGrpcHandler(playerUscase)

	PlayerQuery := playerHandler.NewQueryHandler(*s.config, playerUscase)
	_ = playerHandler
	_ = playerGrpchandler
	_ = PlayerQuery

	player := s.app.Group("/api/v1/player")

	_ = player

}
