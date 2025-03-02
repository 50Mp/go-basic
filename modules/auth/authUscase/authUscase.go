package authuscase

import authrepository "github.com/50Mph/go-api/modules/auth/authRepository"

type (
	AuthUscaseService interface{}
	authUscase        struct {
		authRepository authrepository.AuthrepositoryService
	}
)

func NewAuthUscase(authRepository authrepository.AuthrepositoryService) AuthUscaseService {
	return authUscase{authRepository}
}
