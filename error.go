package gocfg

import (
	"log"
)

// reportError is the centralized error reporting function for unhandled or unexpected errors.
func reportError(err error, context map[string]interface{}) {
	if err != nil {
		log.Printf("ERROR: %v | Context: %v\n", err, context)
		// Here we would integrate with Sentry or another backend if available.
	}
}
