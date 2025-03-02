package itemsrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	ItemsRepository interface {
		// Define methods that the ItemsRepository should implement
	}

	itemRepository struct {
		db *mongo.Client
	}
)

// NewItemRepository creates a new instance of itemRepository
func NewItemRepository(db *mongo.Client) ItemsRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) itemConn(ctx context.Context) *mongo.Database {
	return r.db.Database("item_db")
}
