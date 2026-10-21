package directory_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestDirectory(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
