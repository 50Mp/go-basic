package server

import (
	itemshandler "github.com/50Mph/go-api/modules/items/itemsHandler"
	itemsrepository "github.com/50Mph/go-api/modules/items/itemsRepository"
	itemsuscase "github.com/50Mph/go-api/modules/items/itemsUscase"
)

func (s *server) intemServer() {

	itemsrepository := itemsrepository.NewItemRepository(s.db)
	itemsuscase := itemsuscase.NewItemsUscae(itemsrepository)

	itemshandlers := itemshandler.NewItem(itemsuscase)

	itemsGrpc := itemshandler.NewItemGrpcHandler(itemsuscase)

	_ = itemshandlers
	_ = itemsGrpc

}
