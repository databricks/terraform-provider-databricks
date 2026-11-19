package gitcredential_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestGitCredential(t *testing.T) {
	terraformtest.Run(t, "testdata")
}
