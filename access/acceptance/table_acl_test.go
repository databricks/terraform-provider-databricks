package acceptance

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
)

func TestAccTableACL(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_table_acl" "this" {
				table = "table_acl_{var.RANDOM}"
			
				grant {
					principal = "tf-{var.RANDOM}@example.com"
					privileges = ["ALL PRIVILEGES"]
				}
			
				grant {
					principal = "tfg-{var.RANDOM}"
					privileges = ["SELECT", "READ", "MODIFY"]
				}
			
				deny {
					principal = "users"
					privileges = ["SELECT", "READ"]
				}
			}`,
		},
	})
}
