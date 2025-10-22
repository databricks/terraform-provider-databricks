// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database_database_catalog

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

const dataSourcesName = "database_database_catalogs"

var _ datasource.DataSourceWithConfigure = &DatabaseCatalogsDataSource{}

func DataSourceDatabaseCatalogs() datasource.DataSource {
	return &DatabaseCatalogsDataSource{}
}

// DatabaseCatalogsData extends the main model with additional fields.
type DatabaseCatalogsData struct {
	Database types.List `tfsdk:"database_catalogs"`
	// Name of the instance to get database catalogs for.
	InstanceName types.String `tfsdk:"instance_name"`
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"page_size"`
}

func (DatabaseCatalogsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_catalogs": reflect.TypeOf(DatabaseCatalogData{}),
	}
}

func (m DatabaseCatalogsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["instance_name"] = attrs["instance_name"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["database_catalogs"] = attrs["database_catalogs"].SetComputed()
	return attrs
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *DatabaseCatalogsData) SyncFieldsDuringRead(ctx context.Context, from DatabaseCatalogsData) {
}

type DatabaseCatalogsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *DatabaseCatalogsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *DatabaseCatalogsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, DatabaseCatalogsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks DatabaseCatalog",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *DatabaseCatalogsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *DatabaseCatalogsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config DatabaseCatalogsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest database.ListDatabaseCatalogsRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Database.ListDatabaseCatalogsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list database_database_catalogs", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var database_catalog DatabaseCatalogData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &database_catalog)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, database_catalog.ToObjectValue(ctx))
	}

	var newState DatabaseCatalogsData
	newState.Database = types.ListValueMust(DatabaseCatalogData{}.Type(ctx), results)
	newState.SyncFieldsDuringRead(ctx, config)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
