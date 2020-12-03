package common

import (
	"fmt"
)

var (
	version = "0.2.9"
)

// Version returns version of provider
func Version() string {
	return version
}

// UserAgent returns provider's user agent with the correct version
func UserAgent() string {
	return fmt.Sprintf("databricks-tf-provider/%s", Version())
}
