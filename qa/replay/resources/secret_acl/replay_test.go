package secretacl_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestSecretAcl(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
