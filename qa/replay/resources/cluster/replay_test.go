package cluster_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestCluster(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
