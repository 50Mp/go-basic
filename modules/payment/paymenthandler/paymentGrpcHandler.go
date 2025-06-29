package paymenthandler

import (
	"github.com/50Mph/go-api/config"
	paymentuscase "github.com/50Mph/go-api/modules/payment/paymentUscase"
)

type paymetGrpucHandler struct {
	cfg config.Config
	//
	paymentUsecase paymentuscase.PaymentUsecaseService
}

func NewPaymentGrpcHandler(cfg config.Config, paymentUsecase paymentuscase.PaymentUsecaseService) *paymetGrpucHandler {
	return &paymetGrpucHandler{
		cfg:            cfg,
		paymentUsecase: paymentUsecase}
}
