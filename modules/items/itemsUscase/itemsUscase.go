package itemsuscase

import (
	itemsrepository "github.com/50Mph/go-api/modules/items/itemsRepository"
)

type (
	ItemsUcase interface{}
	itemsUcase struct {
		itemsRepository itemsrepository.ItemsRepository
	}
)

func NewItemsUscae(itemsRepository itemsrepository.ItemsRepository) ItemsUcase {
	return itemsUcase{itemsRepository: itemsRepository}
}
