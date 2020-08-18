package acceptance

import (
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/databrickslabs/databricks-terraform/provider"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func AccTest(t *testing.T, tc resource.TestCase) {
	// each test - create new instance of provider.
	tc.Providers = map[string]terraform.ResourceProvider{
		"databricks": provider.DatabricksProvider(),
	}

	// this allows to debug from VSCode if it's launched with CLOUD_ENV var
	cloudEnv := os.Getenv("CLOUD_ENV")
	tc.IsUnitTest = cloudEnv != ""

	if cloudEnv != "" {
		// let's be more chatty in integration test logs
		for i, s := range tc.Steps {
			if s.Config != "" {
				t.Logf("Test %s (%s) step %d config is:\n%s", 
					t.Name(), cloudEnv, i,
					internal.TrimLeadingWhitespace(s.Config))
			}
		}
	}

	resource.Test(t, tc)
}

// ResourceCheck calls back a function with client and resource id
func ResourceCheck(name string,
	cb func(client *common.DatabricksClient, id string) error) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}
		client := common.CommonEnvironmentClient()
		return cb(client, rs.Primary.ID)
	}
}
