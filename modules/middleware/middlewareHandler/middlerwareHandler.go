package middlewarehandler

import (
	"github.com/50Mph/go-api/config"
	middlerwareuscase "github.com/50Mph/go-api/modules/middleware/midlerwareUscase"
)

type (
	MiddlerwareHandler       interface{}
	middlewareUsecaseHandler struct {
		cfg               config.Config
		middlerwareUscase middlerwareuscase.MiddlerwareUscase
	}
)

func NewMiddlerwareHandler(cfg config.Config, middlerwareUscase middlerwareuscase.MiddlerwareUscase) MiddlerwareHandler {
	return middlewareUsecaseHandler{
		cfg:               cfg,
		middlerwareUscase: middlerwareUscase}
}
