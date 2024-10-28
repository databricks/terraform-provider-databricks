package cluster

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

const resourceName = "cluster"

var _ resource.ResourceWithConfigure = &ClusterResource{}

func ResourceCluster() resource.Resource {
	return &ClusterResource{}
}

type ClusterResource struct {
	Client *common.DatabricksClient
}

func (r *ClusterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksStagingName(resourceName)
}

func (r *ClusterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ClusterResource{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.SetRequired("assets_dir")
		c.SetRequired("output_schema_name")
		c.SetReadOnly("monitor_version")
		c.SetReadOnly("drift_metrics_table_name")
		c.SetReadOnly("profile_metrics_table_name")
		c.SetReadOnly("status")
		c.SetReadOnly("dashboard_id")
		c.SetReadOnly("schedule", "pause_status")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Cluster",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}
