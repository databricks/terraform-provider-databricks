package secretscope_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestSecretScope(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
