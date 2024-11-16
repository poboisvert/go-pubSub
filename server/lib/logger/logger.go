package logger

import (
	"log"
)

// Info logs informational messages.
func Info(message string) {
	log.Printf("[INFO]: %s", message)
}

// Error logs error messages.
func Error(err error) {
	log.Printf("[ERROR]: %s", err)
}
