package payment

import "fmt"

type PayPal struct{}

func (p PayPal) Pay(amount float64) error {
	fmt.Printf("Paid ₹%.2f using PayPal\n", amount)
	return nil
}
