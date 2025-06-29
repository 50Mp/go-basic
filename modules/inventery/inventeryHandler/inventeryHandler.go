package inventeryhandler

import (
	"github.com/50Mph/go-api/config"
	inventeryuscase "github.com/50Mph/go-api/modules/inventery/inventeryUscase"
)

type (
	InventeryHttpsHandler interface {}
	inventeryHttphandler struct {
		cfg             config.Config
		inventeryUscase inventeryuscase.InventeryUscase
	}
)

func NewInventeryHandler(cfg config.Config, inventeryUscase inventeryuscase.InventeryUscase) InventeryHttpsHandler {
	return inventeryHttphandler{
		cfg:             cfg,
		inventeryUscase: inventeryUscase,
	}
}

func (h inventeryHttphandler) NewQueryHandler(cfg config.Config, inventeryUscase inventeryuscase.InventeryUscase) InventeryHttpsHandler {
	return inventeryHttphandler{
		cfg:            cfg ,
		inventeryUscase: inventeryUscase,
	}
}
