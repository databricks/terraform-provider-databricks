package gcp

import (
	"context"
	"reflect"

	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const unityCatalogDataSourceName = "gcp_unity_catalog_policy"

func DataSourceGcpUnityCatalogPolicy() datasource.DataSource {
	return &GcpUnityCatalogPolicyDataSource{}
}

var _ datasource.DataSourceWithConfigure = &GcpUnityCatalogPolicyDataSource{}

type GcpUnityCatalogPolicyDataSource struct {
	Client *common.DatabricksClient
}

type GcpUnityCatalogPolicyData struct {
	DatabricksGoogleServiceAccount types.String `tfsdk:"databricks_google_service_account"`
	Permissions                    types.List   `tfsdk:"permissions"`
	tfschema.Namespace
}

func (GcpUnityCatalogPolicyData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["databricks_google_service_account"] = attrs["databricks_google_service_account"].SetRequired()
	attrs["permissions"] = attrs["permissions"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	return attrs
}

func (GcpUnityCatalogPolicyData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions":     reflect.TypeOf(types.String{}),
		"provider_config": reflect.TypeOf(tfschema.ProviderConfigData{}),
	}
}

func (d *GcpUnityCatalogPolicyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(unityCatalogDataSourceName)
}

func (d *GcpUnityCatalogPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, GcpUnityCatalogPolicyData{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *GcpUnityCatalogPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *GcpUnityCatalogPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, unityCatalogDataSourceName)

	var data GcpUnityCatalogPolicyData
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	permValues := make([]attr.Value, len(gcpUnityCatalogPermissions))
	for i, p := range gcpUnityCatalogPermissions {
		permValues[i] = types.StringValue(p)
	}
	data.Permissions = types.ListValueMust(types.StringType, permValues)
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

// gcpUnityCatalogPermissions are the permissions required for the Unity Catalog
// file events custom role on GCP.
var gcpUnityCatalogPermissions = []string{
	"pubsub.subscriptions.consume",
	"pubsub.subscriptions.create",
	"pubsub.subscriptions.delete",
	"pubsub.subscriptions.get",
	"pubsub.subscriptions.list",
	"pubsub.subscriptions.update",
	"pubsub.topics.attachSubscription",
	"pubsub.topics.create",
	"pubsub.topics.delete",
	"pubsub.topics.detachSubscription",
	"pubsub.topics.get",
	"pubsub.topics.list",
	"pubsub.topics.update",
	"storage.buckets.update",
}
