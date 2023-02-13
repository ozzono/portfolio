package gateway

import "fmt"

// Info ...
func Info(content map[string]interface{}) {
	// No code required.
	// Assume this function will log at info level.
	fmt.Printf("Logger::Info %v", content)
}

// Error ...
func Error(content map[string]interface{}) {
	// No code required.
	// Assume this function will log at error level.
	fmt.Printf("Logger::Error %v", content)
}
