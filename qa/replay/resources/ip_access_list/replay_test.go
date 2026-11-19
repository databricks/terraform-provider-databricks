package ipaccesslist_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestIPAccessList(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
