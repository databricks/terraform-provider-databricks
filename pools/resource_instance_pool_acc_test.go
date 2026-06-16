package pools_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

// TestAccInstancePool_EnableElasticDiskFalse is the live-API counterpart of
// TestResourceInstancePoolCreate_ElasticDiskDisabled. Before the omitempty fix,
// the Create request dropped enable_elastic_disk, the server applied its
// default of true, and the post-apply refresh-plan saw a true -> false diff
// that wanted to recreate the pool — an infinite loop. The implicit
// refresh-plan after Step.Apply fails the test on any non-empty diff, so
// reaching the end of this step verifies the fix end-to-end.
func TestAccInstancePool_EnableElasticDiskFalse(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
			data "databricks_node_type" "smallest" {
				local_disk = true
			}

			resource "databricks_instance_pool" "this" {
				instance_pool_name                    = "test-elastic-disk-false-{var.RANDOM}"
				min_idle_instances                    = 0
				max_capacity                          = 1
				node_type_id                          = data.databricks_node_type.smallest.id
				idle_instance_autotermination_minutes = 10
				enable_elastic_disk                   = false
			}
		`,
	})
}
