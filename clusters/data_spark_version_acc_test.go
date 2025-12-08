package clusters_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccSparkVersion(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `data "databricks_spark_version" "latest" {}`,
		Check: func(state *terraform.State) error {
			id, ok := state.RootModule().Resources["data.databricks_spark_version.latest"].Primary.Attributes["id"]
			if !ok {
				return fmt.Errorf("data.databricks_spark_version.latest not found in state")
			}
			if !strings.HasSuffix(id, "-scala2.13") {
				return fmt.Errorf("data.databricks_spark_version.latest id is not a scala2.13 version")
			}
			return nil
		},
	})
}
