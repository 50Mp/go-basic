package server

import (
	authhandler "github.com/50Mph/go-api/modules/auth/authHandler"
	authrepository "github.com/50Mph/go-api/modules/auth/authRepository"
	authuscase "github.com/50Mph/go-api/modules/auth/authUscase"
)

func (s *server) authServer() {
	authRepository := authrepository.NewAuthRepository(s.db)
	authUsecase := authuscase.NewAuthUscase(authRepository)
	//
	authHandler := authhandler.NewAuthHandler(*s.config, authUsecase)
	//
	authGrpc := authhandler.NewAuthGrpcHandler(*s.config)

	_ = authGrpc
	_ = authHandler

	auth := s.app.Group("/api/v1/auth")

	_ = auth

}
