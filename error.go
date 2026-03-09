package gocfg

import "log"

// reportError is the centralized error reporting function.
func reportError(err error, context map[string]interface{}) {
	if err == nil {
		return
	}
	log.Printf("[SECURITY] Error reported: %v | Context: %v", err, context)
}
