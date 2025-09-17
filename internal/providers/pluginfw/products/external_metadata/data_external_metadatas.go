// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package external_metadata

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

const dataSourcesName = "external_metadatas"

var _ datasource.DataSourceWithConfigure = &ExternalMetadatasDataSource{}

func DataSourceExternalMetadatas() datasource.DataSource {
	return &ExternalMetadatasDataSource{}
}

// ExternalMetadatasData extends the main model with additional fields.
type ExternalMetadatasData struct {
	ExternalMetadata types.List `tfsdk:"external_metadata"`
}

func (ExternalMetadatasData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"external_metadata": reflect.TypeOf(catalog_tf.ExternalMetadata{}),
	}
}

type ExternalMetadatasDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *ExternalMetadatasDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *ExternalMetadatasDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ExternalMetadatasData{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.SetComputed("external_metadata")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks ExternalMetadata",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ExternalMetadatasDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *ExternalMetadatasDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config ExternalMetadatasData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest catalog.ListExternalMetadataRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.ExternalMetadata.ListExternalMetadataAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list external_metadatas", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var external_metadata catalog_tf.ExternalMetadata
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &external_metadata)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, external_metadata.ToObjectValue(ctx))
	}

	var newState ExternalMetadatasData
	newState.ExternalMetadata = types.ListValueMust(catalog_tf.ExternalMetadata{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
