package paymenthandler

import paymentuscase "github.com/50Mph/go-api/modules/payment/paymentUscase"


type paymetGrpucHandler struct {
	paymentUsecase paymentuscase.PaymentUsecaseService
}

func NewPaymentGrpcHandler(paymentUsecase paymentuscase.PaymentUsecaseService) *paymetGrpucHandler {
	return &paymetGrpucHandler{paymentUsecase: paymentUsecase}
}