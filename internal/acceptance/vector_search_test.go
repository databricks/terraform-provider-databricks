package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
)

func TestUcAccVectorSearchEndpoint(t *testing.T) {
	loadUcwsEnv(t)
	if isGcp(t) {
		skipf(t)("not available on GCP")
	}

	name := fmt.Sprintf("terraform-test-vector-search-%[1]s",
		acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	unityWorkspaceLevel(t, LegacyStep{
		Template: fmt.Sprintf(`
			resource "databricks_vector_search_endpoint" "this" {
				name          = "%s"
				endpoint_type = "STANDARD"
			  }
			`, name),
	},
	)
}
