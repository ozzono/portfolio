package gateway

import "fmt"

// SendText ...
func SendText(phone string, message string) error {
	// No code required.
	// Assume this function will send a text message.
	fmt.Printf("Texter::SMS sent to %v. Msg: %v\n", phone, message)
	return nil
}
