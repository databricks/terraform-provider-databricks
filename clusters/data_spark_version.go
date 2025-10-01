package clusters

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceSparkVersion returns DBR version matching to the specification
func DataSourceSparkVersion() common.Resource {
	type SparkVersionDataStruct struct {
		compute.SparkVersionRequest
		common.ProviderConfig
	}
	return common.WorkspaceDataWithCustomizeFunc(func(ctx context.Context, data *SparkVersionDataStruct, w *databricks.WorkspaceClient) error {
		data.Id = ""
		version, err := w.Clusters.SelectSparkVersion(ctx, data.SparkVersionRequest)
		if err != nil {
			return err
		}
		data.Id = version
		return nil
	}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(s, "photon").SetDeprecated("Specify runtime_engine=\"PHOTON\" in the cluster configuration")
		common.CustomizeSchemaPath(s, "graviton").SetDeprecated("Not required anymore - it's automatically enabled on the Graviton-based node types")

		// Add provider_config customizations
		common.CustomizeSchemaPath(s, "provider_config").SetOptional()
		common.CustomizeSchemaPath(s, "provider_config", "workspace_id").SetRequired()

		return s
	})
}
