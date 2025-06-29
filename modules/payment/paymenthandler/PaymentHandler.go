package paymenthandler

import (
	"github.com/50Mph/go-api/config"
	paymentuscase "github.com/50Mph/go-api/modules/payment/paymentUscase"
)

type (
	paymentHandler struct {
		cfg            config.Config
		paymentUsecase paymentuscase.PaymentUsecaseService
	}
)

func NewPaymentHandler(cfg config.Config, paymentUsecase paymentuscase.PaymentUsecaseService) *paymentHandler {
	return &paymentHandler{cfg: cfg,
		paymentUsecase: paymentUsecase}
}
