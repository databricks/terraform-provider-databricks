package group_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestGroup(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
