package group_role_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestGroupRole(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
