// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package account_iam_user_v2

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

const dataSourcesName = "account_iam_users_v2"

var _ datasource.DataSourceWithConfigure = &UsersDataSource{}

func DataSourceUsers() datasource.DataSource {
	return &UsersDataSource{}
}

// UsersData extends the main model with additional fields.
type UsersData struct {
	AccountIamV2 types.List `tfsdk:"users"`
	// Optional. Allows filtering users by username or external id.
	Filter types.String `tfsdk:"filter"`
	// The maximum number of users to return. The service may return fewer than
	// this value.
	PageSize types.Int64 `tfsdk:"page_size"`
}

func (UsersData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"users": reflect.TypeOf(UserData{}),
	}
}

func (m UsersData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()

	attrs["users"] = attrs["users"].SetComputed()
	return attrs
}

type UsersDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *UsersDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *UsersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, UsersData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks User",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *UsersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *UsersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config UsersData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest iamv2.ListUsersRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.IamV2.ListUsersAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list account_iam_users_v2", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var user UserData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &user)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, user.ToObjectValue(ctx))
	}

	config.AccountIamV2 = types.ListValueMust(UserData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
