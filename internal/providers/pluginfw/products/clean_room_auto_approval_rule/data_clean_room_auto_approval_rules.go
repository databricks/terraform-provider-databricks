// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clean_room_auto_approval_rule

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

const dataSourcesName = "clean_room_auto_approval_rules"

var _ datasource.DataSourceWithConfigure = &CleanRoomAutoApprovalRulesDataSource{}

func DataSourceCleanRoomAutoApprovalRules() datasource.DataSource {
	return &CleanRoomAutoApprovalRulesDataSource{}
}

// CleanRoomAutoApprovalRulesData extends the main model with additional fields.
type CleanRoomAutoApprovalRulesData struct {
	CleanRoomAutoApprovalRules types.List   `tfsdk:"rules"`
	WorkspaceID                types.String `tfsdk:"workspace_id"`
}

func (CleanRoomAutoApprovalRulesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rules": reflect.TypeOf(cleanrooms_tf.CleanRoomAutoApprovalRule{}),
	}
}

type CleanRoomAutoApprovalRulesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *CleanRoomAutoApprovalRulesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *CleanRoomAutoApprovalRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, CleanRoomAutoApprovalRulesData{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.SetComputed("rules")
		c.SetOptional("workspace_id")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks CleanRoomAutoApprovalRule",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *CleanRoomAutoApprovalRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *CleanRoomAutoApprovalRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config CleanRoomAutoApprovalRulesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest cleanrooms.ListCleanRoomAutoApprovalRulesRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.CleanRoomAutoApprovalRules.ListAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list clean_room_auto_approval_rules", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var clean_room_auto_approval_rule cleanrooms_tf.CleanRoomAutoApprovalRule
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &clean_room_auto_approval_rule)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, clean_room_auto_approval_rule.ToObjectValue(ctx))
	}

	var newState CleanRoomAutoApprovalRulesData
	newState.CleanRoomAutoApprovalRules = types.ListValueMust(cleanrooms_tf.CleanRoomAutoApprovalRule{}.Type(ctx), results)
	newState.WorkspaceID = config.WorkspaceID
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
