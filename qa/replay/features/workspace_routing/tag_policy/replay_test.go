package workspace_routing_tag_policy_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestWorkspaceRouting(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
