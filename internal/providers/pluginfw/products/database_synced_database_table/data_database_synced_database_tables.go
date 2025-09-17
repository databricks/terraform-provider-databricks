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
	"github.com/databricks/terraform-provider-databricks/internal/service/database_tf"
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
}

func (SyncedDatabaseTablesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"synced_tables": reflect.TypeOf(database_tf.SyncedDatabaseTable{}),
	}
}

type SyncedDatabaseTablesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *SyncedDatabaseTablesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *SyncedDatabaseTablesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, SyncedDatabaseTablesData{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.SetComputed("synced_tables")
		return c
	})
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

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

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

	response, err := client.Database.ListSyncedDatabaseTablesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list database_synced_database_tables", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var synced_database_table database_tf.SyncedDatabaseTable
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &synced_database_table)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, synced_database_table.ToObjectValue(ctx))
	}

	var newState SyncedDatabaseTablesData
	newState.Database = types.ListValueMust(database_tf.SyncedDatabaseTable{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
