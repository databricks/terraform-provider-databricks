package unity_catalog_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestUnityCatalogWorkspaceRouting(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
