package gocfg

import (
	"log"
)

// reportError is the centralized error reporting function.
// All unexpected or unrecoverable errors must be funneled through this function
// instead of failing silently.
func reportError(err error, context map[string]interface{}) {
	if err == nil {
		return
	}

	// If Sentry were available, it would be called here:
	// sentry.CaptureException(err)
	// sentry.WithScope(func(scope *sentry.Scope) {
	//     scope.SetContext("metadata", context)
	//     sentry.CaptureException(err)
	// })

	log.Printf("ERROR: %v | Context: %v\n", err, context)
}
