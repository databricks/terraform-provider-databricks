package grant_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestGrant(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
