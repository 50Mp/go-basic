package server

import (
	inventeryhandler "github.com/50Mph/go-api/modules/inventery/inventeryHandler"
	inventeryrepository "github.com/50Mph/go-api/modules/inventery/inventeryRepository"
	inventeryuscase "github.com/50Mph/go-api/modules/inventery/inventeryUscase"
)

func (s *server) inventoryServer() {

	inventoryRepository := inventeryrepository.NewInventeryRepository(s.db)
	inventoryUsecase := inventeryuscase.NewInebnteryUscase(inventoryRepository)

	inventoryHandler := inventeryhandler.NewInventeryHandler(*s.config, inventoryUsecase)
	inventoryGrpc := inventeryhandler.NewInventerGrpchandler(inventoryUsecase)

	_ = inventoryHandler
	_ = inventoryGrpc

	inventory := s.app.Group("/api/v1/inventory")

	_ = inventory
}
