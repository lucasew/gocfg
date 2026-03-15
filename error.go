package gocfg

import (
	"log"
)

// reportError is a centralized error reporting function.
// All unexpected or unrecoverable errors must be funneled through this function instead of failing silently.
func reportError(err error, context map[string]interface{}) {
	log.Printf("ERROR: %v | Context: %v\n", err, context)
}
