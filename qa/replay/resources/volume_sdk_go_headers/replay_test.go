// Package volume_sdk_go_headers_test runs the databricks_volume User-Agent and
// traceparent checks with the resource on sdk-go uc/volumes/v1. They are a separate
// package from ../volume, which runs the same checks on the existing client, so that
// each package stays within the go test timeout and CI balances them across replay
// shards.
package volume_sdk_go_headers_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/replay/resources/volume/headercheck"
)

// TestVolumeUserAgentDimensionsWithSdkGo pins the User-Agents of
// TestVolumeUserAgentDimensions with the volume calls through sdk-go uc/volumes/v1,
// which append the sdk-go module segments.
func TestVolumeUserAgentDimensionsWithSdkGo(t *testing.T) {
	headercheck.UserAgentDimensions(t, true)
}

// TestVolumeUserAgentExtraWithSdkGo pins the User-Agent of TestVolumeUserAgentExtra with
// the volume calls through sdk-go uc/volumes/v1.
func TestVolumeUserAgentExtraWithSdkGo(t *testing.T) {
	headercheck.UserAgentExtra(t, true)
}

// TestVolumeTraceparentWithSdkGo pins the headers of TestVolumeTraceparent with the
// volume calls through sdk-go uc/volumes/v1, whose retries follow the classic client's
// rules.
func TestVolumeTraceparentWithSdkGo(t *testing.T) {
	headercheck.Traceparent(t, true)
}
