package paymenthandler

import paymentuscase "github.com/50Mph/go-api/modules/payment/paymentUscase"

type (
	paymentHandler struct {
		paymentUsecase paymentuscase.PaymentUsecaseService
	}
)

func NewPaymentHandler(paymentUsecase paymentuscase.PaymentUsecaseService) *paymentHandler {
	return &paymentHandler{paymentUsecase: paymentUsecase}
}
