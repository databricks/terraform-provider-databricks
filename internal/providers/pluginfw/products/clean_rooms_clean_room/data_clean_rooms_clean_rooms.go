// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clean_rooms_clean_room

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

const dataSourcesName = "clean_rooms_clean_rooms"

var _ datasource.DataSourceWithConfigure = &CleanRoomsDataSource{}

func DataSourceCleanRooms() datasource.DataSource {
	return &CleanRoomsDataSource{}
}

// CleanRoomsData extends the main model with additional fields.
type CleanRoomsData struct {
	CleanRooms  types.List   `tfsdk:"clean_rooms"`
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

func (CleanRoomsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clean_rooms": reflect.TypeOf(cleanrooms_tf.CleanRoom{}),
	}
}

type CleanRoomsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *CleanRoomsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *CleanRoomsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, CleanRoomsData{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.SetComputed("clean_rooms")
		c.SetOptional("workspace_id")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks CleanRoom",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *CleanRoomsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *CleanRoomsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config CleanRoomsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest cleanrooms.ListCleanRoomsRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.CleanRooms.ListAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list clean_rooms_clean_rooms", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var clean_room cleanrooms_tf.CleanRoom
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &clean_room)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, clean_room.ToObjectValue(ctx))
	}

	var newState CleanRoomsData
	newState.CleanRooms = types.ListValueMust(cleanrooms_tf.CleanRoom{}.Type(ctx), results)
	newState.WorkspaceID = config.WorkspaceID
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
