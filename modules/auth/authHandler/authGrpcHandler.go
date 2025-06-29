package authhandler

import authuscase "github.com/50Mph/go-api/modules/auth/authUscase"

type (
	AuthGrpcHandler interface {}
	authGrpcHandler struct {
		authUscase authuscase.AuthUscaseService
	}
)

func NewAuthGrpcHandler(authUscase authuscase.AuthUscaseService) AuthGrpcHandler {
	return authGrpcHandler{
		authUscase: authUscase,
	}
}
