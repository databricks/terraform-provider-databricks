// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database_instance

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/database"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourcesName = "database_instances"

var _ datasource.DataSourceWithConfigure = &DatabaseInstancesDataSource{}

func DataSourceDatabaseInstances() datasource.DataSource {
	return &DatabaseInstancesDataSource{}
}

// DatabaseInstancesData extends the main model with additional fields.
type DatabaseInstancesData struct {
	Database           types.List   `tfsdk:"database_instances"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (DatabaseInstancesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instances": reflect.TypeOf(DatabaseInstanceData{}),
		"provider_config":    reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m DatabaseInstancesData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_instances"] = attrs["database_instances"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type DatabaseInstancesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *DatabaseInstancesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *DatabaseInstancesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, DatabaseInstancesData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks DatabaseInstance",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *DatabaseInstancesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *DatabaseInstancesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config DatabaseInstancesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest database.ListDatabaseInstancesRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namespace ProviderConfigData
	resp.Diagnostics.Append(config.ProviderConfigData.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProvider(ctx, namespace.WorkspaceID.ValueString())

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Database.ListDatabaseInstancesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list database_instances", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var database_instance DatabaseInstanceData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &database_instance)...)
		if resp.Diagnostics.HasError() {
			return
		}
		database_instance.ProviderConfigData = config.ProviderConfigData

		results = append(results, database_instance.ToObjectValue(ctx))
	}

	var newState DatabaseInstancesData
	newState.Database = types.ListValueMust(DatabaseInstanceData{}.Type(ctx), results)
	newState.ProviderConfigData = config.ProviderConfigData

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
