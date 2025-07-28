// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clean_room_assets_clean_room_asset

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/cleanrooms"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/cleanrooms_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourcesName = "clean_room_assets_clean_room_assets"

var _ datasource.DataSourceWithConfigure = &CleanRoomAssetsDataSource{}

func DataSourceCleanRoomAssets() datasource.DataSource {
	return &CleanRoomAssetsDataSource{}
}

type CleanRoomAssetsList struct {
	cleanrooms_tf.ListCleanRoomAssetsRequest
	CleanRoomAssets types.List `tfsdk:"assets"`
}

func (c CleanRoomAssetsList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["assets"] = attrs["assets"].SetComputed()
	return attrs
}

func (CleanRoomAssetsList) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"assets": reflect.TypeOf(cleanrooms_tf.CleanRoomAsset{}),
	}
}

type CleanRoomAssetsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *CleanRoomAssetsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *CleanRoomAssetsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, CleanRoomAssetsList{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks CleanRoomAsset",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *CleanRoomAssetsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *CleanRoomAssetsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config CleanRoomAssetsList
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest cleanrooms.ListCleanRoomAssetsRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.CleanRoomAssets.ListAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list clean_room_assets_clean_room_assets", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var clean_room_asset cleanrooms_tf.CleanRoomAsset
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &clean_room_asset)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, clean_room_asset.ToObjectValue(ctx))
	}

	var newState CleanRoomAssetsList
	newState.CleanRoomAssets = types.ListValueMust(cleanrooms_tf.CleanRoomAsset{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
