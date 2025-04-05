package playerrepository

import "go.mongodb.org/mongo-driver/mongo"

type (
	// PlayerRepository represents the player repository interface.
	PlayerRepository interface{}
	// playerRepository is the concrete implementation of PlayerRepository.
	playerRepository struct {
		// db is the database connection.
		db *mongo.Client
	}
)

func NewPlayerRepository(db *mongo.Client) PlayerRepository {
	return playerRepository{db: db}
}
func (r *playerRepository) authConn() *mongo.Database {
	return r.db.Database("player_db")
}
