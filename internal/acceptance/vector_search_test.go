package acceptance

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

func TestAccVectorSearchEndpoint(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	switch cloudEnv {
	case "aws", "azure":
	default:
		t.Skipf("not available on %s", cloudEnv)
	}

	name := fmt.Sprintf("terraform-test-vector-search-%[1]s",
		acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	workspaceLevel(t, step{
		Template: fmt.Sprintf(`
			resource "databricks_vector_search_endpoint" "this" {
				name          = "%s"
				endpoint_type = "STANDARD"
			  }
			`, name),
	},
	)
}
