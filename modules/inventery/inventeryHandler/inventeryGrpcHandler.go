package inventeryhandler

import inventeryuscase "github.com/50Mph/go-api/modules/inventery/inventeryUscase"

type (
	InventeryGrpcHandler struct {
		inventeryUscase inventeryuscase.InventeryUscase
	}
)

func NewInventerGrpchandler(inventeryUscase inventeryuscase.InventeryUscase) *InventeryGrpcHandler {
	return &InventeryGrpcHandler{
		inventeryUscase: inventeryUscase,
	}
}
