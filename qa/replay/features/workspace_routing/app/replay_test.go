package workspace_routing_app_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestWorkspaceRouting(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
