package payment

import "fmt"

type Payment struct {
	paymentStrategy PaymentStrategy
}

func NewPayment(paymentStrategy PaymentStrategy) *Payment {
	return &Payment{
		paymentStrategy: paymentStrategy,
	}
}

func (p *Payment) Checkout(amount float64) error {
	fmt.Println("check inventory")
	_ = p.paymentStrategy.Pay(amount)
	fmt.Println("place order")
	fmt.Println("send notification")

	return nil
}
