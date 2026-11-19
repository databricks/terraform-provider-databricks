// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package account_iam_workspace_assignment_v2

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
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "account_iam_workspace_assignment_v2"

var _ datasource.DataSourceWithConfigure = &WorkspaceAssignmentDataSource{}

func DataSourceWorkspaceAssignment() datasource.DataSource {
	return &WorkspaceAssignmentDataSource{}
}

type WorkspaceAssignmentDataSource struct {
	Client *autogen.DatabricksClient
}

// WorkspaceAssignmentData extends the main model with additional fields.
type WorkspaceAssignmentData struct {
	// The ID of the account that contains the workspace.
	AccountId types.String `tfsdk:"account_id"`
	// Every entitlement the principal has in the workspace, granted directly or
	// through a group. List responses leave it empty.
	EffectiveEntitlements types.Set `tfsdk:"effective_entitlements"`
	// The entitlements granted directly by this assignment. Create and update
	// set exactly this list, even if the principal also gets some of these
	// entitlements through a group. List responses leave it empty.
	Entitlements types.Set `tfsdk:"entitlements"`
	// The resource name of the assignment, in one of these formats: - Account
	// APIs:
	// `accounts/{account_id}/workspaces/{workspace_id}/workspace-assignments/{principal_id}`
	// - Workspace APIs: `workspace-assignments/{principal_id}`
	Name types.String `tfsdk:"name"`
	// The resource name of the workspace, in the format
	// `accounts/{account_id}/workspaces/{workspace_id}`. Workspace APIs return
	// the same format.
	Parent types.String `tfsdk:"parent"`
	// The Databricks ID of the assigned user, service principal, or group. It
	// can't be changed after the assignment is created.
	PrincipalId types.Int64 `tfsdk:"principal_id"`
	// Whether the assigned principal is a user, service principal, or group.
	PrincipalType types.String `tfsdk:"principal_type"`
	// The ID of the workspace.
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// WorkspaceAssignmentData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m WorkspaceAssignmentData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_entitlements": reflect.TypeOf(types.String{}),
		"entitlements":           reflect.TypeOf(types.String{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceAssignmentData
// only implements ToObjectValue() and Type().
func (m WorkspaceAssignmentData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":             m.AccountId,
			"effective_entitlements": m.EffectiveEntitlements,
			"entitlements":           m.Entitlements,
			"name":                   m.Name,
			"parent":                 m.Parent,
			"principal_id":           m.PrincipalId,
			"principal_type":         m.PrincipalType,
			"workspace_id":           m.WorkspaceId,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m WorkspaceAssignmentData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id": types.StringType,
			"effective_entitlements": basetypes.SetType{
				ElemType: types.StringType,
			},
			"entitlements": basetypes.SetType{
				ElemType: types.StringType,
			},
			"name":           types.StringType,
			"parent":         types.StringType,
			"principal_id":   types.Int64Type,
			"principal_type": types.StringType,
			"workspace_id":   types.Int64Type,
		},
	}
}

func (m WorkspaceAssignmentData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["effective_entitlements"] = attrs["effective_entitlements"].SetComputed()
	attrs["entitlements"] = attrs["entitlements"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["parent"] = attrs["parent"].SetComputed()
	attrs["principal_id"] = attrs["principal_id"].SetComputed()
	attrs["principal_type"] = attrs["principal_type"].SetComputed()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()

	return attrs
}

func (r *WorkspaceAssignmentDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *WorkspaceAssignmentDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, WorkspaceAssignmentData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks WorkspaceAssignment",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *WorkspaceAssignmentDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *WorkspaceAssignmentDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config WorkspaceAssignmentData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest iamv2.GetWorkspaceAssignmentRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.IamV2.GetWorkspaceAssignment(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get account_iam_workspace_assignment_v2", err.Error())
		return
	}

	var newState WorkspaceAssignmentData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
