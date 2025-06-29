package server

import (
	authhandler "github.com/50Mph/go-api/modules/auth/authHandler"
	authrepository "github.com/50Mph/go-api/modules/auth/authRepository"
	authuscase "github.com/50Mph/go-api/modules/auth/authUscase"
)

func (s *server) authService() {
	authRepository := authrepository.NewAuthRepository(s.db)
	authUsecase := authuscase.NewAuthUscase(authRepository)
	// HTTP handler
	authHandler := authhandler.NewAuthHandler(*s.config, authUsecase)
	// // gRPC handler
	authGrpc := authhandler.NewAuthGrpcHandler(authUsecase)

	_ = authGrpc
	_ = authHandler

	auth := s.app.Group("/api/v1/auth")

	_ = auth

}
