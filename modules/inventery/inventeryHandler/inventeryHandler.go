package inventeryhandler

import (
	"github.com/50Mph/go-api/config"
	inventeryuscase "github.com/50Mph/go-api/modules/inventery/inventeryUscase"
)

type inventeryHandler struct {
	cfg             config.Config
	inventeryUscase inventeryuscase.InventeryUscase
}

func NewInventeryHandler(cfg config.Config, inventeryUscase inventeryuscase.InventeryUscase) *inventeryHandler {
	return &inventeryHandler{
		cfg:             cfg,
		inventeryUscase: inventeryUscase,
	}
}

func (h inventeryHandler) NewQueryHandler(cfg config.Config, inventeryUscase inventeryuscase.InventeryUscase) *inventeryHandler {
	return &inventeryHandler{
		cfg:             cfg,
		inventeryUscase: inventeryUscase,
	}
}
