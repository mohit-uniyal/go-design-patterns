package main

import "fmt"

type OrderPlaced struct {
	OrderId string
}

type Subscriber interface {
	Notify(orderPlaced OrderPlaced)
}

type EmailService struct{}

func (es *EmailService) Notify(orderPlaced OrderPlaced) {
	fmt.Println("Email sent successfully")
}

type AnalyticsService struct{}

func (es *AnalyticsService) Notify(orderPlaced OrderPlaced) {
	fmt.Println("Analytics saved successfully")
}

type EventBus struct {
	subscribers []Subscriber
}

func (e *EventBus) Subscribe(subscriber Subscriber) {
	e.subscribers = append(e.subscribers, subscriber)
}

func (e *EventBus) Publish(orderPlaced OrderPlaced) error {
	// validations

	for _, subscriber := range e.subscribers {
		subscriber.Notify(orderPlaced)
	}

	return nil
}

type OrderService struct {
	eventBus *EventBus
}

func (os *OrderService) PlaceOrder() {

	orderPlaced := OrderPlaced{OrderId: "ipad"}

	os.eventBus.Publish(orderPlaced)
}

func main() {

	eventBus := &EventBus{}

	emailService := &EmailService{}
	eventBus.Subscribe(emailService)

	analyticsService := &AnalyticsService{}
	eventBus.Subscribe(analyticsService)

	orderService := &OrderService{
		eventBus: eventBus,
	}

	orderService.PlaceOrder()

}
