// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package account_iam_group_v2

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/iamv2"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourcesName = "account_iam_groups_v2"

var _ datasource.DataSourceWithConfigure = &GroupsDataSource{}

func DataSourceGroups() datasource.DataSource {
	return &GroupsDataSource{}
}

// GroupsData extends the main model with additional fields.
type GroupsData struct {
	AccountIamV2 types.List `tfsdk:"groups"`
	// Optional. Allows filtering groups by group name or external id.
	Filter types.String `tfsdk:"filter"`
	// The maximum number of groups to return. The service may return fewer than
	// this value. If not provided, defaults to 1000, which is also the maximum
	// allowed. Requests for more than the maximum are clamped to 1000.
	PageSize types.Int64 `tfsdk:"page_size"`
}

func (GroupsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"groups": reflect.TypeOf(GroupData{}),
	}
}

func (m GroupsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()

	attrs["groups"] = attrs["groups"].SetComputed()
	return attrs
}

type GroupsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *GroupsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *GroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, GroupsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Group",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *GroupsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *GroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config GroupsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest iamv2.ListGroupsRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.IamV2.ListGroupsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list account_iam_groups_v2", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var group GroupData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &group)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, group.ToObjectValue(ctx))
	}

	config.AccountIamV2 = types.ListValueMust(GroupData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
