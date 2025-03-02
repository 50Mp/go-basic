package authrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	AuthrepositoryService interface{}
	authRepository        struct {
		db *mongo.Client
	}
)

func NewAuthRepository(db *mongo.Client) AuthrepositoryService {
	return authRepository{db: db}
}

func (r *authRepository) authConn(ctx context.Context) *mongo.Database {
	return r.db.Database("auth_db")
}
