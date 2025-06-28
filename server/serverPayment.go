package server

import (
	paymentrepositoy "github.com/50Mph/go-api/modules/payment/paymentRepositoy"
	paymentuscase "github.com/50Mph/go-api/modules/payment/paymentUscase"
	"github.com/50Mph/go-api/modules/payment/paymenthandler"
)

func (s *server) paymentServer() {
	paymentRepository := paymentrepositoy.NewPaymentRepository(s.db)
	paymentUsecase := paymentuscase.NewPaymentUsecase(paymentRepository)
	paymentHandler := paymenthandler.NewPaymentHandler(paymentUsecase)
	paymentGrpchandler := paymenthandler.NewPaymentGrpcHandler(paymentUsecase)
	_ = paymentHandler
	_ = paymentGrpchandler

	
	payment := s.app.Group("/api/v1/payment")
	_ = payment
}
