package user

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/iam_tf"
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

type UsersInfo struct {
	Filter types.String  `json:"filter,omitempty"`
	Users  []iam_tf.User `json:"users,omitempty" tf:"computed"`
}

func (d *UsersDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(dataSourceName)
}

func (d *UsersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(UsersInfo{}, nil)
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
	var usersInfo UsersInfo

	resp.Diagnostics.Append(req.Config.Get(ctx, &usersInfo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var users []iam.User
	var err error

	if d.Client.Config.IsAccountClient() {
		a, diags := d.Client.GetAccountClient()
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		users, err = a.Users.ListAll(ctx, iam.ListAccountUsersRequest{Filter: usersInfo.Filter.ValueString()})
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

	for _, user := range users {
		var tfUser iam_tf.User
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, user, &tfUser)...)
		if resp.Diagnostics.HasError() {
			return
		}
		usersInfo.Users = append(usersInfo.Users, tfUser)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, usersInfo)...)
}
