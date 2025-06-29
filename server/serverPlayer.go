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

	playerqueryHandler := playerhandler.NewQueryHttpHandler(*s.config, playerUscase)
	_ = playerqueryHandler
	_ = playerHandler
	_ = playerGrpchandler
	player := s.app.Group("/api/v1/player")

	_ = player

}
