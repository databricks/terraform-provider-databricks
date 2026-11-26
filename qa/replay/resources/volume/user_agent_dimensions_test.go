package volume_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/replay/resources/volume/headercheck"
)

// TestVolumeUserAgentDimensions pins the exact User-Agent of every volume call for
// each environment the classic SDK detects.
func TestVolumeUserAgentDimensions(t *testing.T) {
	headercheck.UserAgentDimensions(t, false)
}

// TestVolumeUserAgentExtra pins DATABRICKS_USER_AGENT_EXTRA on every volume call.
// The provider reads it once at start-up, so the test runs itself in a new process.
func TestVolumeUserAgentExtra(t *testing.T) {
	headercheck.UserAgentExtra(t, false)
}

// TestVolumeTraceparent pins the W3C traceparent header of every volume call,
// including both attempts of each retried call.
func TestVolumeTraceparent(t *testing.T) {
	headercheck.Traceparent(t, false)
}
