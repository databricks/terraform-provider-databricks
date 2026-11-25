package lro

import "time"

// LroOptions is the options for the Long Running Operations.
// DO NOT USE THIS OPTION. This option is still under development
// and can be updated in the future without notice.
type LroOptions struct {
	// Timeout is the timeout for the Long Running Operations.
	// If not set, the default timeout is 20 minutes.
	Timeout time.Duration
}
