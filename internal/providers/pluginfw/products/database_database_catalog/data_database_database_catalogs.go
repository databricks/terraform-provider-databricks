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
	"github.com/databricks/terraform-provider-databricks/internal/service/database_tf"
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

// DatabaseCatalogsDataExtended extends the main model with additional fields.
type DatabaseCatalogsDataExtended struct {
	database_tf.ListDatabaseCatalogsRequest
	Database types.List `tfsdk:"database_catalogs"`
}

func (c DatabaseCatalogsDataExtended) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_catalogs"] = attrs["database_catalogs"].SetComputed()
	return attrs
}

func (DatabaseCatalogsDataExtended) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_catalogs": reflect.TypeOf(database_tf.DatabaseCatalog{}),
	}
}

type DatabaseCatalogsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *DatabaseCatalogsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *DatabaseCatalogsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, DatabaseCatalogsDataExtended{}, nil)
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

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config DatabaseCatalogsDataExtended
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest database.ListDatabaseCatalogsRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
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
		var database_catalog database_tf.DatabaseCatalog
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &database_catalog)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, database_catalog.ToObjectValue(ctx))
	}

	var newState DatabaseCatalogsDataExtended
	newState.Database = types.ListValueMust(database_tf.DatabaseCatalog{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
