package main

import (
	"fmt"
	"net/http"
	"time"
)

func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("==> Authentication: Checking Token")

		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			fmt.Println("Authentication Failed")
			return
		}

		fmt.Println("Authentication Successful")

		// Call the next middleware/handler
		next(w, r)
	}
}

func Logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request received for ", r.Method, r.Pattern)
		next(w, r)
	}
}

func Timing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Println("Timing Start")
		next(w, r)
		fmt.Println("request completed in ", time.Since(start))
	}
}

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Order placed")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Order placed successfully"))
}

func main() {

	handler := Authentication(
		Logging(
			Timing(PlaceOrder),
		),
	)

	http.HandleFunc("/order", handler)

	http.ListenAndServe(":8080", nil)
}
