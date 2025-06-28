package server

import (
	playerhandler "github.com/50Mph/go-api/modules/player/playerHandler"
	playerrepository "github.com/50Mph/go-api/modules/player/playerRepository"
	playeruscase "github.com/50Mph/go-api/modules/player/playerUscase"
)

func (s *server) playerServer() {
	playerRepository := playerrepository.NewPlayerRepository(s.db)
	playerUscase := playeruscase.NewPlayerUsecase(playerRepository)
	playerHandler := playerhandler.NewPlayerHandler(playerUscase)
	playerGrpchandler := playerhandler.NewPlayerGrpcHandler(playerUscase)

	_ = playerHandler
	_ = playerGrpchandler

	player := s.app.Group("/api/v1/player")

	_ = player

}
 