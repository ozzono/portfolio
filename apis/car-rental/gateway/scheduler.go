package gateway

import "fmt"

// Schedule ...
func Schedule(duration int32, callback interface{}, args ...interface{}) {
	// No code required.
	// Assume this function schedule a callback for a duration of time in seconds
	fmt.Printf("Scheduler::Schedule: callback scheduled for %v seconds later with %d args %v\n", duration, len(args), args)
}
