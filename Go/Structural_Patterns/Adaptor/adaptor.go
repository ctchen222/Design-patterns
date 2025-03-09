package adaptor

import "fmt"

type PaymentGateway interface {
	Pay(amount float64)
}

type PayPal struct{}

func (p *PayPal) Pay(amount float64) {
	fmt.Printf("Paying %.2f using PayPal\n", amount)
}

type Stripe struct{}

func (s *Stripe) MakePayment(amount float64) {
	fmt.Printf("Paying %.2f using Stripe\n", amount)
}

type StripeAdaptor struct {
	StripeService *Stripe
}

func (s *StripeAdaptor) Pay(amount float64) {
	s.StripeService.MakePayment(amount)
}
