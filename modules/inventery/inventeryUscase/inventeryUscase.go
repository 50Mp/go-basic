package inventeryuscase

import inventeryrepository "github.com/50Mph/go-api/modules/inventery/inventeryRepository"

type (
	InventeryUscase interface{}
	inventerUscase  struct {
		inventerRepository inventeryrepository.InventeryRepositoryService
	}
)

func NewInebnteryUscase(inventerRepository inventeryrepository.InventeryRepositoryService) InventeryUscase {
	return inventerUscase{
		inventerRepository: inventerRepository,
	}
}
