package service_principal_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestServicePrincipal(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
