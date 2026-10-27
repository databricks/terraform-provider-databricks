package workspaceconf_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestWorkspaceConf(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
