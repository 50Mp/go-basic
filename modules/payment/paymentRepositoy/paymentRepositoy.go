package paymentrepositoy

import "go.mongodb.org/mongo-driver/mongo"

type (
	PaymentRepositoryService interface{}
	paymentRepository        struct {
		db *mongo.Client
	}
)

func NewPaymentRepository(db *mongo.Client) PaymentRepositoryService {
	return paymentRepository{db: db}
}
func (r *paymentRepository) authConn() *mongo.Database {
	return r.db.Database("payment_db")
}
