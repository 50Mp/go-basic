package itemshandler

import (
	itemsuscase "github.com/50Mph/go-api/modules/items/itemsUscase"
)

type (
	itemGrpcHandler struct {
		itemuscase itemsuscase.ItemsUcase
	}
)

func NewItemGrpcHandler(itemuscase itemsuscase.ItemsUcase) itemsuscase.ItemsUcase {
	return itemGrpcHandler{
		itemuscase: itemuscase,
	}

}
