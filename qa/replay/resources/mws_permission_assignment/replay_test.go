package mws_permission_assignment_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestMwsPermissionAssignment(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
