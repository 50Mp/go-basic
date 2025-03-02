package authhandler

import (
	"github.com/50Mph/go-api/config"
	authuscase "github.com/50Mph/go-api/modules/auth/authUscase"
)

type (
	AuthHandlerService interface{}
	authHandler        struct {
		authUscase authuscase.AuthUscaseService
	}
)

func NewAuthHandler(cfg config.Config, authUscase authuscase.AuthUscaseService) AuthHandlerService {
	return authHandler{
		authUscase: authUscase,
	}
}
