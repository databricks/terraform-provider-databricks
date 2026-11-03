package awscrossaccountpolicy_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestAwsCrossaccountPolicy(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
