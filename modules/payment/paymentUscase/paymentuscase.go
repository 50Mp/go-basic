package paymentuscase

import paymentrepositoy "github.com/50Mph/go-api/modules/payment/paymentRepositoy"

type (
	PaymentUsecaseService interface{}
	paymentUsecase        struct {
		paymentRepository paymentrepositoy.PaymentRepositoryService
	}
)

func NewPaymentUsecase(paymentRepository paymentrepositoy.PaymentRepositoryService) PaymentUsecaseService {
	return paymentUsecase{paymentRepository: paymentRepository}
}
func (p *paymentUsecase) authConn() paymentrepositoy.PaymentRepositoryService {
	return p.paymentRepository
}
