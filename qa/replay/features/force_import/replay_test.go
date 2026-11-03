package forceimport_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestForceImport(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
