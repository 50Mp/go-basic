package itemshandler

import (
	itemsuscase "github.com/50Mph/go-api/modules/items/itemsUscase"
)

type (
	ItemGrpcHandler interface {}
	itemGrpcHandler struct {
		itemuscase itemsuscase.ItemsUcase
	}
)

func NewItemGrpcHandler(itemuscase itemsuscase.ItemsUcase) ItemGrpcHandler {
	return itemGrpcHandler{
		itemuscase: itemuscase,
	}

}
