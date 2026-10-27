package grants_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestGrants(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
