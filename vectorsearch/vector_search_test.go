package vectorsearch_test

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
)

func TestUcAccVectorSearchEndpoint(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	if acceptance.IsGcp(t) {
		acceptance.Skipf(t)("not available on GCP")
	}

	name := fmt.Sprintf("terraform-test-vector-search-%[1]s",
		acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: fmt.Sprintf(`
			resource "databricks_vector_search_endpoint" "this" {
				name          = "%s"
				endpoint_type = "STANDARD"
			}

			resource "databricks_permissions" "this" {
				vector_search_endpoint_id = databricks_vector_search_endpoint.this.endpoint_id

				access_control {
					group_name = "users"
					permission_level = "CAN_USE"
				}
			}
			`, name),
	},
	)
}
