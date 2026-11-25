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

func TestReplay(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
