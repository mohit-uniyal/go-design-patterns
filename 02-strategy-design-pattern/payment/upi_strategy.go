package payment

import "fmt"

type UPI struct{}

func (u UPI) Pay(amount float64) error {
	fmt.Printf("Paid ₹%.2f using UPI\n", amount)
	return nil
}
