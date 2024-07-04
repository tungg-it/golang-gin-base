package recover

import (
	"fmt"
)

// GoroutineRecover is a function to recover from panic in concurrent condition
// It help to prevent the panic from crashing the application
func Goroutine() {
	if err := recover(); err != nil {
		fmt.Printf("[GoroutineRecover] recovered from panic: %v", err)
	}
}
