package user_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestUser(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
