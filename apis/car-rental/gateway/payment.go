package gateway

import "fmt"

// Process ...
func Process(uuid string, amount float64, creditCard string) error {
	// No code required.
	// Assume this function will charge or refund to the given credit card
	// A negative amount means refund.
	fmt.Printf("Payment Processed: user: %v amount: %v", uuid, amount)
	return nil
}
