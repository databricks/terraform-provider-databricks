// Package mlflow_model_decoding_test runs the databricks_mlflow_model journeys that pin
// how success responses are decoded: missing envelopes, extra fields, timestamps, nulls,
// wrong types, undecodable bodies and unexpected answers to the PATCH and the DELETE.
// They are a separate package from ../mlflow_model so that each package stays within the
// go test timeout and CI balances them across replay shards.
package mlflow_model_decoding_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestReplay runs the journeys on the classic SDK path of databricks_mlflow_model. It
// clears the opt-in, so that the environment of the test process can't switch the path.
// The journeys in testdata_classic pin what only the classic path does.
func TestReplay(t *testing.T) {
	t.Setenv("USE_SDK_GO_RESOURCES", "")
	terraformtest.Run(t, "testdata")
	terraformtest.Run(t, "testdata_classic")
}

// TestReplaySdkGo runs the journeys of testdata on the modular SDK path of the resource.
// Each journey in testdata_sdkgo is the twin of the one with its name in
// testdata_classic, and its header says where the modular path differs and why.
func TestReplaySdkGo(t *testing.T) {
	t.Setenv("USE_SDK_GO_RESOURCES", "databricks_mlflow_model")
	terraformtest.Run(t, "testdata")
	terraformtest.Run(t, "testdata_sdkgo")
}
