package paymenthandler

import (
	"github.com/50Mph/go-api/config"
	paymentuscase "github.com/50Mph/go-api/modules/payment/paymentUscase"
)

type  (
	// PaymentGrpcHandler defines the interface for the payment gRPC handler.
	PaymentGrpcHandler interface {}
	// paymentGrpcHandler implements the PaymentGrpcHandler interface.
	paymentGrpcHandler struct {
		cfg config.Config
		//
		paymentUsecase paymentuscase.PaymentUsecaseService
	}
)
func NewPaymentGrpcHandler(cfg config.Config, paymentUsecase paymentuscase.PaymentUsecaseService) PaymentGrpcHandler {
	return paymentGrpcHandler{
		cfg:          cfg,
		paymentUsecase: paymentUsecase}
}
