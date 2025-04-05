package inventeryrepository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	InventeryRepositoryService interface{}
	inventeryRepository        struct {
		db *mongo.Client
	}
)

func NewInventeryRepository(db *mongo.Client) InventeryRepositoryService {
	return inventeryRepository{db: db}
}
func (r *inventeryRepository) authConn() *mongo.Database {
	return r.db.Database("inventery_db")
}
