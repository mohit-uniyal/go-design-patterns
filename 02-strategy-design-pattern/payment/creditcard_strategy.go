package payment

import "fmt"

type CreditCard struct{}

func (c CreditCard) Pay(amount float64) error {
	fmt.Printf("Paid ₹%.2f using Credit Card\n", amount)
	return nil
}
