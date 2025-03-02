package middlerwareuscase

import midlerwarerepository "github.com/50Mph/go-api/modules/middleware/midlerwareRepository"

type (
	MiddlerwareUscase interface{}

	middlerwareUscase struct {
		middlerwareRepository midlerwarerepository.MiddlerwareRepository
	}
)

func NewMiddlerwareUscase(middlerwareRepository midlerwarerepository.MiddlerwareRepository) MiddlerwareUscase {
	return middlerwareUscase{middlerwareRepository: middlerwareRepository}
}
