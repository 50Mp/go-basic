package midlerwarerepository

type (
	MiddlerwareRepository interface{}
	middlerwareRepository struct{}
)

func NewMiddlerwareRepository() MiddlerwareRepository {
	return &middlerwareRepository{}
}