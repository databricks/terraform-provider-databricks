// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clean_room_auto_approval_rule

import (
	"context"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/cleanrooms"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/cleanrooms_tf"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

const dataSourceName = "clean_room_auto_approval_rule"

var _ datasource.DataSourceWithConfigure = &CleanRoomAutoApprovalRuleDataSource{}

func DataSourceCleanRoomAutoApprovalRule() datasource.DataSource {
	return &CleanRoomAutoApprovalRuleDataSource{}
}

type CleanRoomAutoApprovalRuleDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *CleanRoomAutoApprovalRuleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *CleanRoomAutoApprovalRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, cleanrooms_tf.CleanRoomAutoApprovalRule{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks CleanRoomAutoApprovalRule",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *CleanRoomAutoApprovalRuleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *CleanRoomAutoApprovalRuleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config cleanrooms_tf.CleanRoomAutoApprovalRule
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest cleanrooms.GetCleanRoomAutoApprovalRuleRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.CleanRoomAutoApprovalRules.Get(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get clean_room_auto_approval_rule", err.Error())
		return
	}

	var newState cleanrooms_tf.CleanRoomAutoApprovalRule
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(config)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
