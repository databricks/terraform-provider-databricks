package spark_version_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestSparkVersion(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
