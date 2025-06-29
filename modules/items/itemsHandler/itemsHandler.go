package itemshandler

import itemsuscase "github.com/50Mph/go-api/modules/items/itemsUscase"

type (
	ItemHttpHandler interface {}
	// Item represents an item in the inventory.
	itemHttphandler struct {
		itemsUcase itemsuscase.ItemsUcase
	}
)

func NewItemHttpHandler(itemsUcase itemsuscase.ItemsUcase) ItemGrpcHandler {
	return itemHttphandler{
		itemsUcase: itemsUcase,
	}
}
