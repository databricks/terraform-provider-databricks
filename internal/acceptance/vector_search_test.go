package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
)

func TestUcAccVectorSearchEndpoint(t *testing.T) {
	LoadUcwsEnv(t)
	if IsGcp(t) {
		skipf(t)("not available on GCP")
	}

	name := fmt.Sprintf("terraform-test-vector-search-%[1]s",
		acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	UnityWorkspaceLevel(t, Step{
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
