// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres_branch

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/postgres"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "postgres_branch"

var _ datasource.DataSourceWithConfigure = &BranchDataSource{}

func DataSourceBranch() datasource.DataSource {
	return &BranchDataSource{}
}

type BranchDataSource struct {
	Client *autogen.DatabricksClient
}

// BranchData extends the main model with additional fields.
type BranchData struct {
	// A timestamp indicating when the branch was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The branch's state, indicating if it is initializing, ready for use, or
	// archived.
	CurrentState types.String `tfsdk:"current_state"`
	// Whether the branch is the project's default branch. This field is only
	// returned on create/update responses. See effective_default for the value
	// that is actually applied to the branch.
	Default types.Bool `tfsdk:"default"`
	// Whether the branch is the project's default branch.
	EffectiveDefault types.Bool `tfsdk:"effective_default"`
	// Whether the branch is protected.
	EffectiveIsProtected types.Bool `tfsdk:"effective_is_protected"`
	// The name of the source branch from which this branch was created. Format:
	// projects/{project_id}/branches/{branch_id}
	EffectiveSourceBranch types.String `tfsdk:"effective_source_branch"`
	// The Log Sequence Number (LSN) on the source branch from which this branch
	// was created.
	EffectiveSourceBranchLsn types.String `tfsdk:"effective_source_branch_lsn"`
	// The point in time on the source branch from which this branch was
	// created.
	EffectiveSourceBranchTime timetypes.RFC3339 `tfsdk:"effective_source_branch_time"`
	// Whether the branch is protected.
	IsProtected types.Bool `tfsdk:"is_protected"`
	// The logical size of the branch.
	LogicalSizeBytes types.Int64 `tfsdk:"logical_size_bytes"`
	// The resource name of the branch. Format:
	// projects/{project_id}/branches/{branch_id}
	Name types.String `tfsdk:"name"`
	// The project containing this branch. Format: projects/{project_id}
	Parent types.String `tfsdk:"parent"`
	// The pending state of the branch, if a state transition is in progress.
	PendingState types.String `tfsdk:"pending_state"`
	// The name of the source branch from which this branch was created. Format:
	// projects/{project_id}/branches/{branch_id}
	SourceBranch types.String `tfsdk:"source_branch"`
	// The Log Sequence Number (LSN) on the source branch from which this branch
	// was created.
	SourceBranchLsn types.String `tfsdk:"source_branch_lsn"`
	// The point in time on the source branch from which this branch was
	// created.
	SourceBranchTime timetypes.RFC3339 `tfsdk:"source_branch_time"`
	// A timestamp indicating when the `current_state` began.
	StateChangeTime timetypes.RFC3339 `tfsdk:"state_change_time"`
	// System generated unique ID for the branch.
	Uid types.String `tfsdk:"uid"`
	// A timestamp indicating when the branch was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// BranchData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m BranchData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BranchData
// only implements ToObjectValue() and Type().
func (m BranchData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":                  m.CreateTime,
			"current_state":                m.CurrentState,
			"default":                      m.Default,
			"effective_default":            m.EffectiveDefault,
			"effective_is_protected":       m.EffectiveIsProtected,
			"effective_source_branch":      m.EffectiveSourceBranch,
			"effective_source_branch_lsn":  m.EffectiveSourceBranchLsn,
			"effective_source_branch_time": m.EffectiveSourceBranchTime,
			"is_protected":                 m.IsProtected,
			"logical_size_bytes":           m.LogicalSizeBytes,
			"name":                         m.Name,
			"parent":                       m.Parent,
			"pending_state":                m.PendingState,
			"source_branch":                m.SourceBranch,
			"source_branch_lsn":            m.SourceBranchLsn,
			"source_branch_time":           m.SourceBranchTime,
			"state_change_time":            m.StateChangeTime,
			"uid":                          m.Uid,
			"update_time":                  m.UpdateTime,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m BranchData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":                  timetypes.RFC3339{}.Type(ctx),
			"current_state":                types.StringType,
			"default":                      types.BoolType,
			"effective_default":            types.BoolType,
			"effective_is_protected":       types.BoolType,
			"effective_source_branch":      types.StringType,
			"effective_source_branch_lsn":  types.StringType,
			"effective_source_branch_time": timetypes.RFC3339{}.Type(ctx),
			"is_protected":                 types.BoolType,
			"logical_size_bytes":           types.Int64Type,
			"name":                         types.StringType,
			"parent":                       types.StringType,
			"pending_state":                types.StringType,
			"source_branch":                types.StringType,
			"source_branch_lsn":            types.StringType,
			"source_branch_time":           timetypes.RFC3339{}.Type(ctx),
			"state_change_time":            timetypes.RFC3339{}.Type(ctx),
			"uid":                          types.StringType,
			"update_time":                  timetypes.RFC3339{}.Type(ctx),
		},
	}
}

func (m BranchData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["current_state"] = attrs["current_state"].SetComputed()
	attrs["default"] = attrs["default"].SetComputed()
	attrs["effective_default"] = attrs["effective_default"].SetComputed()
	attrs["effective_is_protected"] = attrs["effective_is_protected"].SetComputed()
	attrs["effective_source_branch"] = attrs["effective_source_branch"].SetComputed()
	attrs["effective_source_branch_lsn"] = attrs["effective_source_branch_lsn"].SetComputed()
	attrs["effective_source_branch_time"] = attrs["effective_source_branch_time"].SetComputed()
	attrs["is_protected"] = attrs["is_protected"].SetComputed()
	attrs["logical_size_bytes"] = attrs["logical_size_bytes"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["parent"] = attrs["parent"].SetComputed()
	attrs["pending_state"] = attrs["pending_state"].SetComputed()
	attrs["source_branch"] = attrs["source_branch"].SetComputed()
	attrs["source_branch_lsn"] = attrs["source_branch_lsn"].SetComputed()
	attrs["source_branch_time"] = attrs["source_branch_time"].SetComputed()
	attrs["state_change_time"] = attrs["state_change_time"].SetComputed()
	attrs["uid"] = attrs["uid"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

func (r *BranchDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *BranchDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, BranchData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Branch",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *BranchDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *BranchDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config BranchData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest postgres.GetBranchRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Postgres.GetBranch(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get postgres_branch", err.Error())
		return
	}

	var newState BranchData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
