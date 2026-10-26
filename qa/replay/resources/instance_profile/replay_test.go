package instance_profile_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestInstanceProfile(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
