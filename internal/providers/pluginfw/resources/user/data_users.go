package user

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func DataSourceUsers() datasource.DataSource {
	return &UsersDataSource{}
}

var _ datasource.DataSourceWithConfigure = &UsersDataSource{}

type UsersDataSource struct {
	Client *common.DatabricksClient
}

type UsersInfo struct {
	DisplayNameContains string     `json:"display_name_contains,omitempty" tf:"computed"`
	UserNameContains    string     `json:"user_name_contains,omitempty" tf:"computed"`
	Users               []iam.User `json:"users,omitempty" tf:"computed"`
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

// DiagsFromError helps us create an error diagnostic from an error.
func DiagsFromError(summary string, err error) diag.Diagnostics {
	return diag.Diagnostics{
		diag.NewErrorDiagnostic(summary, err.Error()),
	}
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

	var filter string

	if userInfo.DisplayNameContains != "" {
		filter = fmt.Sprintf("displayName co \"%s\"", userInfo.DisplayNameContains)
	} else if userInfo.UserNameContains != "" {
		filter = fmt.Sprintf("userName co \"%s\"", userInfo.UserNameContains)
	}

	var users []iam.User
	var err error

	if d.Client.Config.IsAccountClient() {
		a, diags := d.Client.GetAccountClient()
		if AppendDiagAndCheckErrors(resp, diags) {
			return
		}
		users, err = a.Users.ListAll(ctx, iam.ListAccountUsersRequest{Filter: filter})
		if err != nil && AppendDiagAndCheckErrors(resp, DiagsFromError("Error listing account users", err)) {
			return
		}
	} else {
		w, diags := d.Client.GetWorkspaceClient()
		if AppendDiagAndCheckErrors(resp, diags) {
			return
		}
		users, err = w.Users.ListAll(ctx, iam.ListUsersRequest{Filter: filter})
		if err != nil && AppendDiagAndCheckErrors(resp, DiagsFromError("Error listing workspace users", err)) {
			return
		}
	}

	userInfo.Users = users
	resp.Diagnostics.Append(resp.State.Set(ctx, userInfo)...)
}
