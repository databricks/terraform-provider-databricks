package cluster_policy_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestClusterPolicy(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
