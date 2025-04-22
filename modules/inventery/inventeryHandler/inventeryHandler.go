package inventeryhandler

import (
	"github.com/50Mph/go-api/config"
	inventeryuscase "github.com/50Mph/go-api/modules/inventery/inventeryUscase"
)

type (
	InventeryHandler interface{}

	inventeryHandler struct {
		cfg             config.Config
		inventeryUscase inventeryuscase.InventeryUscase
	}
)

func NewInventeryHandler(cfg config.Config, inventeryUscase inventeryuscase.InventeryUscase) InventeryHandler {
	return inventeryHandler{
		cfg: cfg,

		inventeryUscase: inventeryUscase,
	}
}
