package itemshandler

import itemsuscase "github.com/50Mph/go-api/modules/items/itemsUscase"

type (
	// Item represents an item in the inventory.
	Item struct {
		itemsUcase itemsuscase.ItemsUcase
	}
)

func NewItemHttpHandler(itemsUcase itemsuscase.ItemsUcase) *Item {
	return &Item{
		itemsUcase: itemsUcase,
	}
}
