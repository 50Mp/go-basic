package inventeryhandler

import inventeryuscase "github.com/50Mph/go-api/modules/inventery/inventeryUscase"

type (
	InventeryGrpcHandler interface {}
	inventeryGrpcHandler struct {
		inventeryUscase inventeryuscase.InventeryUscase
	}
)

func NewInventerGrpchandler(inventeryUscase inventeryuscase.InventeryUscase) InventeryGrpcHandler {
	return &inventeryGrpcHandler{
		inventeryUscase: inventeryUscase,
	}
}
