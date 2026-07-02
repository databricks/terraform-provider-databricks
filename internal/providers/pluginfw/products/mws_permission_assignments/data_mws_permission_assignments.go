package mws_permission_assignments

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/iam_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourceName = "mws_permission_assignments"

func DataSourceMwsPermissionAssignments() datasource.DataSource {
	return &MwsPermissionAssignmentsDataSource{}
}

var _ datasource.DataSourceWithConfigure = &MwsPermissionAssignmentsDataSource{}

type MwsPermissionAssignmentsDataSource struct {
	Client *common.DatabricksClient
}

// MwsPermissionAssignmentsInfo is the Terraform schema for the
// databricks_mws_permission_assignments data source. WorkspaceId is the only
// input; PermissionAssignments holds the list read back from the account API.
type MwsPermissionAssignmentsInfo struct {
	WorkspaceId           types.Int64 `tfsdk:"workspace_id"`
	PermissionAssignments types.List  `tfsdk:"permission_assignments"`
}

func (MwsPermissionAssignmentsInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	attrs["permission_assignments"] = attrs["permission_assignments"].SetComputed()
	return attrs
}

func (MwsPermissionAssignmentsInfo) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_assignments": reflect.TypeOf(iam_tf.PermissionAssignment{}),
	}
}

func (d *MwsPermissionAssignmentsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(dataSourceName)
}

func (d *MwsPermissionAssignmentsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, MwsPermissionAssignmentsInfo{}, nil)
	resp.Schema = schema.Schema{
		Description: "Lists all workspace permission assignments for a workspace via the Databricks account API. This data source can only be used with an account-level provider.",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (d *MwsPermissionAssignmentsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

// flattenPermissionAssignments converts the SDK response into the list of
// Terraform SDK objects used in the data source state. It is kept separate from
// Read so the flatten/convert logic can be unit tested without a live client.
func flattenPermissionAssignments(ctx context.Context, assignments *iam.PermissionAssignments) ([]attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	result := []attr.Value{}
	if assignments == nil {
		return result, diags
	}
	for _, assignment := range assignments.PermissionAssignments {
		var tfAssignment iam_tf.PermissionAssignment
		diags.Append(converters.GoSdkToTfSdkStruct(ctx, assignment, &tfAssignment)...)
		if diags.HasError() {
			return result, diags
		}
		result = append(result, tfAssignment.ToObjectValue(ctx))
	}
	return result, diags
}

func (d *MwsPermissionAssignmentsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var info MwsPermissionAssignmentsInfo
	resp.Diagnostics.Append(req.Config.Get(ctx, &info)...)
	if resp.Diagnostics.HasError() {
		return
	}

	a, diags := d.Client.GetAccountClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	assignments, err := a.WorkspaceAssignment.ListByWorkspaceId(ctx, info.WorkspaceId.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError("failed to list workspace permission assignments", err.Error())
		return
	}

	tfAssignments, flattenDiags := flattenPermissionAssignments(ctx, assignments)
	resp.Diagnostics.Append(flattenDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	info.PermissionAssignments = types.ListValueMust(iam_tf.PermissionAssignment{}.Type(ctx), tfAssignments)
	resp.Diagnostics.Append(resp.State.Set(ctx, info)...)
}
