package user

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DataSourceUsers() datasource.DataSource {
	return &UsersDataSource{}
}

var _ datasource.DataSourceWithConfigure = &UsersDataSource{}

type UsersDataSource struct {
	Client *common.DatabricksClient
}

type UserData struct {
	Id          types.String `tfsdk:"id,omitempty" tf:"computed"`
	UserName    types.String `tfsdk:"user_name,omitempty" tf:"computed"`
	DisplayName types.String `tfsdk:"display_name,omitempty" tf:"computed"`
}

type UsersInfo struct {
	DisplayNameContains string     `json:"display_name_contains,omitempty" tf:"computed"`
	UserNameContains    string     `json:"user_name_contains,omitempty" tf:"computed"`
	Users               []UserData `json:"users,omitempty" tf:"computed"` // TODO: use  UserData[] or []iam_tf.ListAccountUserResponse / []iam_tf.ListUserResponse?
}

func (d *UsersDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "databricks_users"
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

// AppendDiagAndCheckErrors is a helper function that simplifies error handling by combining the appending of diagnostics and the checking of errors in a single step.
// It is particularly useful in data source and resource read operations where you want to append diagnostics and immediately determine if an error has occurred.
func AppendDiagAndCheckErrors(resp *datasource.ReadResponse, diags diag.Diagnostics) bool {
	resp.Diagnostics.Append(diags...)
	return resp.Diagnostics.HasError()
}

func validateFilters(data *UsersInfo) diag.Diagnostics {
	if data.DisplayNameContains != "" && data.UserNameContains != "" {
		return diag.Diagnostics{diag.NewErrorDiagnostic("Invalid configuration", "Exactly one of display_name_contains or user_name_contains should be specified, not both.")}
	}
	return nil
}

func (d *UsersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var userInfo UsersInfo

	diags := req.Config.Get(ctx, &userInfo)
	diags = append(diags, validateFilters(&userInfo)...)

	if AppendDiagAndCheckErrors(resp, diags) {
		return
	}

	if d.Client.Config.IsAccountClient() {
		a, diags := d.Client.GetAccountClient()
		if AppendDiagAndCheckErrors(resp, diags) {
			return
		}
		// TODO: Add retrieval of iterator at the account level.
	} else {
		w, diags := d.Client.GetWorkspaceClient()
		if AppendDiagAndCheckErrors(resp, diags) {
			return
		}
		// TODO: Add retreival of iterator at the workspace level.
	}
	// TODO: Continue setting the datasource.
}
