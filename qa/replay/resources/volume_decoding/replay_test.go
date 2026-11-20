// Package volume_decoding_test runs the databricks_volume journeys that pin how 200
// response bodies are decoded: unknown, missing and wrongly typed fields, JSON shape, HTML
// and the request log in decode errors. They are a separate package from ../volume so that
// each package stays within the go test timeout and CI balances them across replay shards.
package volume_decoding_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
