// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database_synced_database_table

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
)

const dataSourcesName = "database_synced_database_tables"

var _ datasource.DataSourceWithConfigure = &SyncedDatabaseTablesDataSource{}

func DataSourceSyncedDatabaseTables() datasource.DataSource {
	return &SyncedDatabaseTablesDataSource{}
}

// SyncedDatabaseTablesData extends the main model with additional fields.
type SyncedDatabaseTablesData struct {
	Database types.List `tfsdk:"synced_tables"`
	// Name of the instance to get synced tables for.
	InstanceName types.String `tfsdk:"instance_name"`
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"page_size"`
}

func (SyncedDatabaseTablesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"synced_tables": reflect.TypeOf(SyncedDatabaseTableData{}),
	}
}

func (m SyncedDatabaseTablesData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["instance_name"] = attrs["instance_name"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["synced_tables"] = attrs["synced_tables"].SetComputed()
	return attrs
}

type SyncedDatabaseTablesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *SyncedDatabaseTablesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *SyncedDatabaseTablesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, SyncedDatabaseTablesData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks SyncedDatabaseTable",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *SyncedDatabaseTablesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *SyncedDatabaseTablesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config SyncedDatabaseTablesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest database.ListSyncedDatabaseTablesRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Database.ListSyncedDatabaseTablesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list database_synced_database_tables", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var synced_database_table SyncedDatabaseTableData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &synced_database_table)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, synced_database_table.ToObjectValue(ctx))
	}

	config.Database = types.ListValueMust(SyncedDatabaseTableData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
