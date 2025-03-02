package authhandler

import authuscase "github.com/50Mph/go-api/modules/auth/authUscase"

type (
	authGrpcHandler struct {
		authUscase authuscase.AuthUscaseService
	}
)

func NewAuthGrpcHandler(authUscase authuscase.AuthUscaseService) authuscase.AuthUscaseService {
	return authGrpcHandler{
		authUscase: authUscase,
	}
}
