package paymenthandler

import (
	"github.com/50Mph/go-api/config"
	paymentuscase "github.com/50Mph/go-api/modules/payment/paymentUscase"
)

type (
	// PaymentHttpsHandler defines the interface for the payment HTTP handler.
	PaymentHttpsHandler interface{}
	paymentHandler struct {
		cfg            config.Config
		paymentUsecase paymentuscase.PaymentUsecaseService
	}
)

func NewPaymentHandler(cfg config.Config, paymentUsecase paymentuscase.PaymentUsecaseService) PaymentHttpsHandler {
	return &paymentHandler{
		cfg: cfg,
		paymentUsecase: paymentUsecase,
	}
}
