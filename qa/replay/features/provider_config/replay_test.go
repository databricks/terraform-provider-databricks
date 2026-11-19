package providerconfig_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestProviderConfig(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
