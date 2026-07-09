package payment

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64) error
}

func GetPaymentStrategy(paymentStrategy string) (PaymentStrategy, error) {
	switch paymentStrategy {
	case "creditcard":
		return CreditCard{}, nil
	case "upi":
		return UPI{}, nil
	case "paypal":
		return PayPal{}, nil
	default:
		return nil, fmt.Errorf("not a valid payment strategy")
	}
}
