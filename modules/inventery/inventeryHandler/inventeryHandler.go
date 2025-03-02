package inventeryhandler

import inventeryuscase "github.com/50Mph/go-api/modules/inventery/inventeryUscase"

type (
	InventeryHandler interface{}

	inventeryHandler struct {
		inventeryUscase inventeryuscase.InventeryUscase
	}
)

func NewInventeryHandler(inventeryUscase inventeryuscase.InventeryUscase) InventeryHandler {
	return inventeryHandler{
		inventeryUscase: inventeryUscase,
	}
}
