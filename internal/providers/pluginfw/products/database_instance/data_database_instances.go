// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database_instance

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourcesName = "database_instances"

var _ datasource.DataSourceWithConfigure = &DatabaseInstancesDataSource{}

func DataSourceDatabaseInstances() datasource.DataSource {
	return &DatabaseInstancesDataSource{}
}

type DatabaseInstancesList struct {
	catalog_tf.ListDatabaseInstancesRequest
	DatabaseInstances types.List `tfsdk:"database_instances"`
}

func (c DatabaseInstancesList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_instances"] = attrs["database_instances"].SetComputed()
	return attrs
}

func (DatabaseInstancesList) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instances": reflect.TypeOf(catalog_tf.DatabaseInstance{}),
	}
}

type DatabaseInstancesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *DatabaseInstancesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *DatabaseInstancesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, DatabaseInstancesList{}, nil)
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

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config DatabaseInstancesList
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest catalog.ListDatabaseInstancesRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.DatabaseInstances.ListDatabaseInstancesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list database_instances", err.Error())
		return
	}

	var database_instances = []attr.Value{}
	for _, item := range response {
		var database_instance catalog_tf.DatabaseInstance
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &database_instance)...)
		if resp.Diagnostics.HasError() {
			return
		}
		database_instances = append(database_instances, database_instance.ToObjectValue(ctx))
	}

	var newState DatabaseInstancesList
	newState.DatabaseInstances = types.ListValueMust(catalog_tf.DatabaseInstance{}.Type(ctx), database_instances)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
