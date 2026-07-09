package main

import "strategy-design-pattern/payment"

func main() {
	userPaymentMethod := "upi"

	// This is the factory design pattern
	// The wrapping of strategy choosing logic in a place(factory) is factory design pattern.
	paymentStrategy, _ := payment.GetPaymentStrategy(userPaymentMethod) //This is our factory method

	//1. kind of like dependency injection
	//2. but the dependency changes(behaviour) at runtime
	//3. choose when you have to choose between multiple things/strategies
	paymentClient := payment.NewPayment(paymentStrategy)

	paymentClient.Checkout(100)
}
