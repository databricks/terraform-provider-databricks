package user

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/iam_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourceName = "users"

func DataSourceUsers() datasource.DataSource {
	return &UsersDataSource{}
}

var _ datasource.DataSourceWithConfigure = &UsersDataSource{}

type UsersDataSource struct {
	Client *common.DatabricksClient
}

type UsersList struct {
	Filter          types.String `tfsdk:"filter"`
	ExtraAttributes types.String `tfsdk:"extra_attributes"`
	Users           types.List   `tfsdk:"users"`
}

func (UsersList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["users"] = attrs["users"].SetComputed().SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["extra_attributes"] = attrs["extra_attributes"].SetOptional()

	return attrs
}

func (UsersList) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"users": reflect.TypeOf(iam_tf.User{}),
	}
}

func (d *UsersDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(dataSourceName)
}

func (d *UsersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, UsersList{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *UsersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *UsersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var usersInfo UsersList
	attributes := "id,userName,displayName,externalId"

	resp.Diagnostics.Append(req.Config.Get(ctx, &usersInfo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !(usersInfo.ExtraAttributes.IsNull()) {
		attributes += ","
		attributes += usersInfo.ExtraAttributes.String()
	}

	var users []iam.User
	var err error

	if d.Client.Config.IsAccountClient() {
		a, diags := d.Client.GetAccountClient()
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		users, err = a.Users.ListAll(ctx, iam.ListAccountUsersRequest{Filter: usersInfo.Filter.ValueString(), Attributes: attributes})
		if err != nil {
			resp.Diagnostics.AddError("Error listing account users", err.Error())
		}
	} else {
		w, diags := d.Client.GetWorkspaceClient()
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		users, err = w.Users.ListAll(ctx, iam.ListUsersRequest{Filter: usersInfo.Filter.ValueString()})
		if err != nil {
			resp.Diagnostics.AddError("Error listing workspace users", err.Error())
		}
	}
	tfUsers := []attr.Value{}
	for _, user := range users {
		var tfUser iam_tf.User
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, user, &tfUser)...)
		if resp.Diagnostics.HasError() {
			return
		}
		tfUsers = append(tfUsers, tfUser.ToObjectValue(ctx))
	}

	usersInfo.Users = types.ListValueMust(iam_tf.User{}.Type(ctx), tfUsers)

	resp.Diagnostics.Append(resp.State.Set(ctx, usersInfo)...)
}
