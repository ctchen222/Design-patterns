package strategy

import "fmt"

type PaymentStrategy interface {
	Pay(amount int)
}

type CreditCartPayment struct {
	cardNumber string
}

func (c *CreditCartPayment) Pay(amount int) {
	fmt.Printf("Pay %d with card %s\n", amount, c.cardNumber)
}

func NewCreditCardPayment(cardNumber string) *CreditCartPayment {
	return &CreditCartPayment{
		cardNumber: cardNumber,
	}
}

type PaypalPayment struct {
	email string
}

func (p *PaypalPayment) Pay(amount int) {
	fmt.Printf("Pay %d with paypal %s\n", amount, p.email)
}

func NewPaypalPayment(email string) *PaypalPayment {
	return &PaypalPayment{
		email: email,
	}
}

type LinePayment struct {
	mobile string
}

func (p *LinePayment) Pay(amount int) {
	fmt.Printf("Pay %d with linepay %s\n", amount, p.mobile)
}

func NewLinePayment(mobile string) *LinePayment {
	return &LinePayment{
		mobile: mobile,
	}
}

type PaymentContext struct {
	Strategy PaymentStrategy
}

func (p *PaymentContext) SetPayment(payment PaymentStrategy) {
	p.Strategy = payment
}

func NewPaymentContext(payment PaymentStrategy) *PaymentContext {
	return &PaymentContext{
		Strategy: payment,
	}
}
