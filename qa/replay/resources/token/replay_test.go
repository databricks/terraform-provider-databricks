package token_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestToken(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
