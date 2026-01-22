// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package postgres_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type Branch_SdkV2 struct {
	// A timestamp indicating when the branch was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The resource name of the branch. This field is output-only and
	// constructed by the system. Format:
	// `projects/{project_id}/branches/{branch_id}`
	Name types.String `tfsdk:"name"`
	// The project containing this branch (API resource hierarchy). Format:
	// projects/{project_id}
	//
	// Note: This field indicates where the branch exists in the resource
	// hierarchy. For point-in-time branching from another branch, see
	// `spec.source_branch`.
	Parent types.String `tfsdk:"parent"`
	// The spec contains the branch configuration.
	Spec types.List `tfsdk:"spec"`
	// The current status of a Branch.
	Status types.List `tfsdk:"status"`
	// System-generated unique ID for the branch.
	Uid types.String `tfsdk:"uid"`
	// A timestamp indicating when the branch was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *Branch_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Branch_SdkV2) {
	if !from.Spec.IsUnknown() && !from.Spec.IsNull() {
		// Spec is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Spec = from.Spec
	}
	if !from.Spec.IsNull() && !from.Spec.IsUnknown() {
		if toSpec, ok := to.GetSpec(ctx); ok {
			if fromSpec, ok := from.GetSpec(ctx); ok {
				// Recursively sync the fields of Spec
				toSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromSpec)
				to.SetSpec(ctx, toSpec)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				// Recursively sync the fields of Status
				toStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
}

func (to *Branch_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Branch_SdkV2) {
	if !from.Spec.IsUnknown() && !from.Spec.IsNull() {
		// Spec is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Spec = from.Spec
	}
	if !from.Spec.IsNull() && !from.Spec.IsUnknown() {
		if toSpec, ok := to.GetSpec(ctx); ok {
			if fromSpec, ok := from.GetSpec(ctx); ok {
				toSpec.SyncFieldsDuringRead(ctx, fromSpec)
				to.SetSpec(ctx, toSpec)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				toStatus.SyncFieldsDuringRead(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
}

func (m Branch_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["parent"] = attrs["parent"].SetComputed()
	attrs["spec"] = attrs["spec"].SetOptional()
	attrs["spec"] = attrs["spec"].SetComputed()
	attrs["spec"] = attrs["spec"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["spec"] = attrs["spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetComputed()
	attrs["status"] = attrs["status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["uid"] = attrs["uid"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Branch.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Branch_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spec":   reflect.TypeOf(BranchSpec_SdkV2{}),
		"status": reflect.TypeOf(BranchStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Branch_SdkV2
// only implements ToObjectValue() and Type().
func (m Branch_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time": m.CreateTime,
			"name":        m.Name,
			"parent":      m.Parent,
			"spec":        m.Spec,
			"status":      m.Status,
			"uid":         m.Uid,
			"update_time": m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Branch_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": timetypes.RFC3339{}.Type(ctx),
			"name":        types.StringType,
			"parent":      types.StringType,
			"spec": basetypes.ListType{
				ElemType: BranchSpec_SdkV2{}.Type(ctx),
			},
			"status": basetypes.ListType{
				ElemType: BranchStatus_SdkV2{}.Type(ctx),
			},
			"uid":         types.StringType,
			"update_time": timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetSpec returns the value of the Spec field in Branch_SdkV2 as
// a BranchSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Branch_SdkV2) GetSpec(ctx context.Context) (BranchSpec_SdkV2, bool) {
	var e BranchSpec_SdkV2
	if m.Spec.IsNull() || m.Spec.IsUnknown() {
		return e, false
	}
	var v []BranchSpec_SdkV2
	d := m.Spec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSpec sets the value of the Spec field in Branch_SdkV2.
func (m *Branch_SdkV2) SetSpec(ctx context.Context, v BranchSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spec"]
	m.Spec = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in Branch_SdkV2 as
// a BranchStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Branch_SdkV2) GetStatus(ctx context.Context) (BranchStatus_SdkV2, bool) {
	var e BranchStatus_SdkV2
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v []BranchStatus_SdkV2
	d := m.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in Branch_SdkV2.
func (m *Branch_SdkV2) SetStatus(ctx context.Context, v BranchStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	m.Status = types.ListValueMust(t, vs)
}

type BranchOperationMetadata_SdkV2 struct {
}

func (to *BranchOperationMetadata_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BranchOperationMetadata_SdkV2) {
}

func (to *BranchOperationMetadata_SdkV2) SyncFieldsDuringRead(ctx context.Context, from BranchOperationMetadata_SdkV2) {
}

func (m BranchOperationMetadata_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BranchOperationMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m BranchOperationMetadata_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BranchOperationMetadata_SdkV2
// only implements ToObjectValue() and Type().
func (m BranchOperationMetadata_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m BranchOperationMetadata_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type BranchSpec_SdkV2 struct {
	// Absolute expiration timestamp. When set, the branch will expire at this
	// time.
	ExpireTime timetypes.RFC3339 `tfsdk:"expire_time"`
	// When set to true, protects the branch from deletion and reset. Associated
	// compute endpoints and the project cannot be deleted while the branch is
	// protected.
	IsProtected types.Bool `tfsdk:"is_protected"`
	// Explicitly disable expiration. When set to true, the branch will not
	// expire. If set to false, the request is invalid; provide either ttl or
	// expire_time instead.
	NoExpiry types.Bool `tfsdk:"no_expiry"`
	// The name of the source branch from which this branch was created (data
	// lineage for point-in-time recovery). If not specified, defaults to the
	// project's default branch. Format:
	// projects/{project_id}/branches/{branch_id}
	SourceBranch types.String `tfsdk:"source_branch"`
	// The Log Sequence Number (LSN) on the source branch from which this branch
	// was created.
	SourceBranchLsn types.String `tfsdk:"source_branch_lsn"`
	// The point in time on the source branch from which this branch was
	// created.
	SourceBranchTime timetypes.RFC3339 `tfsdk:"source_branch_time"`
	// Relative time-to-live duration. When set, the branch will expire at
	// creation_time + ttl.
	Ttl timetypes.GoDuration `tfsdk:"ttl"`
}

func (to *BranchSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BranchSpec_SdkV2) {
}

func (to *BranchSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from BranchSpec_SdkV2) {
}

func (m BranchSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["expire_time"] = attrs["expire_time"].SetOptional()
	attrs["is_protected"] = attrs["is_protected"].SetOptional()
	attrs["no_expiry"] = attrs["no_expiry"].SetOptional()
	attrs["source_branch"] = attrs["source_branch"].SetOptional()
	attrs["source_branch"] = attrs["source_branch"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["source_branch_lsn"] = attrs["source_branch_lsn"].SetOptional()
	attrs["source_branch_lsn"] = attrs["source_branch_lsn"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["source_branch_time"] = attrs["source_branch_time"].SetOptional()
	attrs["source_branch_time"] = attrs["source_branch_time"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["ttl"] = attrs["ttl"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BranchSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m BranchSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BranchSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m BranchSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"expire_time":        m.ExpireTime,
			"is_protected":       m.IsProtected,
			"no_expiry":          m.NoExpiry,
			"source_branch":      m.SourceBranch,
			"source_branch_lsn":  m.SourceBranchLsn,
			"source_branch_time": m.SourceBranchTime,
			"ttl":                m.Ttl,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BranchSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"expire_time":        timetypes.RFC3339{}.Type(ctx),
			"is_protected":       types.BoolType,
			"no_expiry":          types.BoolType,
			"source_branch":      types.StringType,
			"source_branch_lsn":  types.StringType,
			"source_branch_time": timetypes.RFC3339{}.Type(ctx),
			"ttl":                timetypes.GoDuration{}.Type(ctx),
		},
	}
}

type BranchStatus_SdkV2 struct {
	// The branch's state, indicating if it is initializing, ready for use, or
	// archived.
	CurrentState types.String `tfsdk:"current_state"`
	// Whether the branch is the project's default branch.
	Default types.Bool `tfsdk:"default"`
	// Absolute expiration time for the branch. Empty if expiration is disabled.
	ExpireTime timetypes.RFC3339 `tfsdk:"expire_time"`
	// Whether the branch is protected.
	IsProtected types.Bool `tfsdk:"is_protected"`
	// The logical size of the branch.
	LogicalSizeBytes types.Int64 `tfsdk:"logical_size_bytes"`
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
}

func (to *BranchStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BranchStatus_SdkV2) {
}

func (to *BranchStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, from BranchStatus_SdkV2) {
}

func (m BranchStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["current_state"] = attrs["current_state"].SetComputed()
	attrs["default"] = attrs["default"].SetComputed()
	attrs["expire_time"] = attrs["expire_time"].SetComputed()
	attrs["is_protected"] = attrs["is_protected"].SetComputed()
	attrs["logical_size_bytes"] = attrs["logical_size_bytes"].SetComputed()
	attrs["pending_state"] = attrs["pending_state"].SetComputed()
	attrs["source_branch"] = attrs["source_branch"].SetComputed()
	attrs["source_branch_lsn"] = attrs["source_branch_lsn"].SetComputed()
	attrs["source_branch_time"] = attrs["source_branch_time"].SetComputed()
	attrs["state_change_time"] = attrs["state_change_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BranchStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m BranchStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BranchStatus_SdkV2
// only implements ToObjectValue() and Type().
func (m BranchStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"current_state":      m.CurrentState,
			"default":            m.Default,
			"expire_time":        m.ExpireTime,
			"is_protected":       m.IsProtected,
			"logical_size_bytes": m.LogicalSizeBytes,
			"pending_state":      m.PendingState,
			"source_branch":      m.SourceBranch,
			"source_branch_lsn":  m.SourceBranchLsn,
			"source_branch_time": m.SourceBranchTime,
			"state_change_time":  m.StateChangeTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BranchStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"current_state":      types.StringType,
			"default":            types.BoolType,
			"expire_time":        timetypes.RFC3339{}.Type(ctx),
			"is_protected":       types.BoolType,
			"logical_size_bytes": types.Int64Type,
			"pending_state":      types.StringType,
			"source_branch":      types.StringType,
			"source_branch_lsn":  types.StringType,
			"source_branch_time": timetypes.RFC3339{}.Type(ctx),
			"state_change_time":  timetypes.RFC3339{}.Type(ctx),
		},
	}
}

type CreateBranchRequest_SdkV2 struct {
	// The Branch to create.
	Branch types.List `tfsdk:"branch"`
	// The ID to use for the Branch. This becomes the final component of the
	// branch's resource name. The ID must be 1-63 characters long, start with a
	// lowercase letter, and contain only lowercase letters, numbers, and
	// hyphens (RFC 1123). Examples: - With custom ID: `staging` → name
	// becomes `projects/{project_id}/branches/staging` - Without custom ID:
	// system generates slug → name becomes
	// `projects/{project_id}/branches/br-example-name-x1y2z3a4`
	BranchId types.String `tfsdk:"-"`
	// The Project where this Branch will be created. Format:
	// projects/{project_id}
	Parent types.String `tfsdk:"-"`
}

func (to *CreateBranchRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateBranchRequest_SdkV2) {
	if !from.Branch.IsNull() && !from.Branch.IsUnknown() {
		if toBranch, ok := to.GetBranch(ctx); ok {
			if fromBranch, ok := from.GetBranch(ctx); ok {
				// Recursively sync the fields of Branch
				toBranch.SyncFieldsDuringCreateOrUpdate(ctx, fromBranch)
				to.SetBranch(ctx, toBranch)
			}
		}
	}
}

func (to *CreateBranchRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateBranchRequest_SdkV2) {
	if !from.Branch.IsNull() && !from.Branch.IsUnknown() {
		if toBranch, ok := to.GetBranch(ctx); ok {
			if fromBranch, ok := from.GetBranch(ctx); ok {
				toBranch.SyncFieldsDuringRead(ctx, fromBranch)
				to.SetBranch(ctx, toBranch)
			}
		}
	}
}

func (m CreateBranchRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch"] = attrs["branch"].SetRequired()
	attrs["branch"] = attrs["branch"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["branch_id"] = attrs["branch_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateBranchRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateBranchRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"branch": reflect.TypeOf(Branch_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBranchRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateBranchRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch":    m.Branch,
			"branch_id": m.BranchId,
			"parent":    m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateBranchRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch": basetypes.ListType{
				ElemType: Branch_SdkV2{}.Type(ctx),
			},
			"branch_id": types.StringType,
			"parent":    types.StringType,
		},
	}
}

// GetBranch returns the value of the Branch field in CreateBranchRequest_SdkV2 as
// a Branch_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateBranchRequest_SdkV2) GetBranch(ctx context.Context) (Branch_SdkV2, bool) {
	var e Branch_SdkV2
	if m.Branch.IsNull() || m.Branch.IsUnknown() {
		return e, false
	}
	var v []Branch_SdkV2
	d := m.Branch.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBranch sets the value of the Branch field in CreateBranchRequest_SdkV2.
func (m *CreateBranchRequest_SdkV2) SetBranch(ctx context.Context, v Branch_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["branch"]
	m.Branch = types.ListValueMust(t, vs)
}

type CreateEndpointRequest_SdkV2 struct {
	// The Endpoint to create.
	Endpoint types.List `tfsdk:"endpoint"`
	// The ID to use for the Endpoint. This becomes the final component of the
	// endpoint's resource name. The ID must be 1-63 characters long, start with
	// a lowercase letter, and contain only lowercase letters, numbers, and
	// hyphens (RFC 1123). Examples: - With custom ID: `primary` → name
	// becomes `projects/{project_id}/branches/{branch_id}/endpoints/primary` -
	// Without custom ID: system generates slug → name becomes
	// `projects/{project_id}/branches/{branch_id}/endpoints/ep-example-name-x1y2z3a4`
	EndpointId types.String `tfsdk:"-"`
	// The Branch where this Endpoint will be created. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent types.String `tfsdk:"-"`
}

func (to *CreateEndpointRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateEndpointRequest_SdkV2) {
	if !from.Endpoint.IsNull() && !from.Endpoint.IsUnknown() {
		if toEndpoint, ok := to.GetEndpoint(ctx); ok {
			if fromEndpoint, ok := from.GetEndpoint(ctx); ok {
				// Recursively sync the fields of Endpoint
				toEndpoint.SyncFieldsDuringCreateOrUpdate(ctx, fromEndpoint)
				to.SetEndpoint(ctx, toEndpoint)
			}
		}
	}
}

func (to *CreateEndpointRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateEndpointRequest_SdkV2) {
	if !from.Endpoint.IsNull() && !from.Endpoint.IsUnknown() {
		if toEndpoint, ok := to.GetEndpoint(ctx); ok {
			if fromEndpoint, ok := from.GetEndpoint(ctx); ok {
				toEndpoint.SyncFieldsDuringRead(ctx, fromEndpoint)
				to.SetEndpoint(ctx, toEndpoint)
			}
		}
	}
}

func (m CreateEndpointRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint"] = attrs["endpoint"].SetRequired()
	attrs["endpoint"] = attrs["endpoint"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["endpoint_id"] = attrs["endpoint_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoint": reflect.TypeOf(Endpoint_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint":    m.Endpoint,
			"endpoint_id": m.EndpointId,
			"parent":      m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint": basetypes.ListType{
				ElemType: Endpoint_SdkV2{}.Type(ctx),
			},
			"endpoint_id": types.StringType,
			"parent":      types.StringType,
		},
	}
}

// GetEndpoint returns the value of the Endpoint field in CreateEndpointRequest_SdkV2 as
// a Endpoint_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateEndpointRequest_SdkV2) GetEndpoint(ctx context.Context) (Endpoint_SdkV2, bool) {
	var e Endpoint_SdkV2
	if m.Endpoint.IsNull() || m.Endpoint.IsUnknown() {
		return e, false
	}
	var v []Endpoint_SdkV2
	d := m.Endpoint.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEndpoint sets the value of the Endpoint field in CreateEndpointRequest_SdkV2.
func (m *CreateEndpointRequest_SdkV2) SetEndpoint(ctx context.Context, v Endpoint_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["endpoint"]
	m.Endpoint = types.ListValueMust(t, vs)
}

type CreateProjectRequest_SdkV2 struct {
	// The Project to create.
	Project types.List `tfsdk:"project"`
	// The ID to use for the Project. This becomes the final component of the
	// project's resource name. The ID must be 1-63 characters long, start with
	// a lowercase letter, and contain only lowercase letters, numbers, and
	// hyphens (RFC 1123). Examples: - With custom ID: `production` → name
	// becomes `projects/production` - Without custom ID: system generates UUID
	// → name becomes `projects/a7f89b2c-3d4e-5f6g-7h8i-9j0k1l2m3n4o`
	ProjectId types.String `tfsdk:"-"`
}

func (to *CreateProjectRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateProjectRequest_SdkV2) {
	if !from.Project.IsNull() && !from.Project.IsUnknown() {
		if toProject, ok := to.GetProject(ctx); ok {
			if fromProject, ok := from.GetProject(ctx); ok {
				// Recursively sync the fields of Project
				toProject.SyncFieldsDuringCreateOrUpdate(ctx, fromProject)
				to.SetProject(ctx, toProject)
			}
		}
	}
}

func (to *CreateProjectRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateProjectRequest_SdkV2) {
	if !from.Project.IsNull() && !from.Project.IsUnknown() {
		if toProject, ok := to.GetProject(ctx); ok {
			if fromProject, ok := from.GetProject(ctx); ok {
				toProject.SyncFieldsDuringRead(ctx, fromProject)
				to.SetProject(ctx, toProject)
			}
		}
	}
}

func (m CreateProjectRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project"] = attrs["project"].SetRequired()
	attrs["project"] = attrs["project"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["project_id"] = attrs["project_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateProjectRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateProjectRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"project": reflect.TypeOf(Project_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateProjectRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateProjectRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"project":    m.Project,
			"project_id": m.ProjectId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateProjectRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"project": basetypes.ListType{
				ElemType: Project_SdkV2{}.Type(ctx),
			},
			"project_id": types.StringType,
		},
	}
}

// GetProject returns the value of the Project field in CreateProjectRequest_SdkV2 as
// a Project_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateProjectRequest_SdkV2) GetProject(ctx context.Context) (Project_SdkV2, bool) {
	var e Project_SdkV2
	if m.Project.IsNull() || m.Project.IsUnknown() {
		return e, false
	}
	var v []Project_SdkV2
	d := m.Project.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetProject sets the value of the Project field in CreateProjectRequest_SdkV2.
func (m *CreateProjectRequest_SdkV2) SetProject(ctx context.Context, v Project_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["project"]
	m.Project = types.ListValueMust(t, vs)
}

type CreateRoleRequest_SdkV2 struct {
	// The Branch where this Role is created. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent types.String `tfsdk:"-"`
	// The desired specification of a Role.
	Role types.List `tfsdk:"role"`
	// The ID to use for the Role, which will become the final component of the
	// role's resource name. This ID becomes the role in Postgres.
	//
	// This value should be 4-63 characters, and valid characters are lowercase
	// letters, numbers, and hyphens, as defined by RFC 1123.
	RoleId types.String `tfsdk:"-"`
}

func (to *CreateRoleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateRoleRequest_SdkV2) {
	if !from.Role.IsNull() && !from.Role.IsUnknown() {
		if toRole, ok := to.GetRole(ctx); ok {
			if fromRole, ok := from.GetRole(ctx); ok {
				// Recursively sync the fields of Role
				toRole.SyncFieldsDuringCreateOrUpdate(ctx, fromRole)
				to.SetRole(ctx, toRole)
			}
		}
	}
}

func (to *CreateRoleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateRoleRequest_SdkV2) {
	if !from.Role.IsNull() && !from.Role.IsUnknown() {
		if toRole, ok := to.GetRole(ctx); ok {
			if fromRole, ok := from.GetRole(ctx); ok {
				toRole.SyncFieldsDuringRead(ctx, fromRole)
				to.SetRole(ctx, toRole)
			}
		}
	}
}

func (m CreateRoleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["role"] = attrs["role"].SetRequired()
	attrs["role"] = attrs["role"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["role_id"] = attrs["role_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateRoleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateRoleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"role": reflect.TypeOf(Role_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRoleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateRoleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parent":  m.Parent,
			"role":    m.Role,
			"role_id": m.RoleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateRoleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent": types.StringType,
			"role": basetypes.ListType{
				ElemType: Role_SdkV2{}.Type(ctx),
			},
			"role_id": types.StringType,
		},
	}
}

// GetRole returns the value of the Role field in CreateRoleRequest_SdkV2 as
// a Role_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateRoleRequest_SdkV2) GetRole(ctx context.Context) (Role_SdkV2, bool) {
	var e Role_SdkV2
	if m.Role.IsNull() || m.Role.IsUnknown() {
		return e, false
	}
	var v []Role_SdkV2
	d := m.Role.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRole sets the value of the Role field in CreateRoleRequest_SdkV2.
func (m *CreateRoleRequest_SdkV2) SetRole(ctx context.Context, v Role_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["role"]
	m.Role = types.ListValueMust(t, vs)
}

type DatabaseCredential_SdkV2 struct {
	// Timestamp in UTC of when this credential expires.
	ExpireTime timetypes.RFC3339 `tfsdk:"expire_time"`
	// The OAuth token that can be used as a password when connecting to a
	// database.
	Token types.String `tfsdk:"token"`
}

func (to *DatabaseCredential_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseCredential_SdkV2) {
}

func (to *DatabaseCredential_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DatabaseCredential_SdkV2) {
}

func (m DatabaseCredential_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["expire_time"] = attrs["expire_time"].SetOptional()
	attrs["token"] = attrs["token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabaseCredential.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabaseCredential_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseCredential_SdkV2
// only implements ToObjectValue() and Type().
func (m DatabaseCredential_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"expire_time": m.ExpireTime,
			"token":       m.Token,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseCredential_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"expire_time": timetypes.RFC3339{}.Type(ctx),
			"token":       types.StringType,
		},
	}
}

// Databricks Error that is returned by all Databricks APIs.
type DatabricksServiceExceptionWithDetailsProto_SdkV2 struct {
	Details types.List `tfsdk:"details"`

	ErrorCode types.String `tfsdk:"error_code"`

	Message types.String `tfsdk:"message"`

	StackTrace types.String `tfsdk:"stack_trace"`
}

func (to *DatabricksServiceExceptionWithDetailsProto_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabricksServiceExceptionWithDetailsProto_SdkV2) {
	if !from.Details.IsNull() && !from.Details.IsUnknown() && to.Details.IsNull() && len(from.Details.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Details, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Details = from.Details
	}
}

func (to *DatabricksServiceExceptionWithDetailsProto_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DatabricksServiceExceptionWithDetailsProto_SdkV2) {
	if !from.Details.IsNull() && !from.Details.IsUnknown() && to.Details.IsNull() && len(from.Details.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Details, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Details = from.Details
	}
}

func (m DatabricksServiceExceptionWithDetailsProto_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["details"] = attrs["details"].SetOptional()
	attrs["error_code"] = attrs["error_code"].SetOptional()
	attrs["message"] = attrs["message"].SetOptional()
	attrs["stack_trace"] = attrs["stack_trace"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabricksServiceExceptionWithDetailsProto.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabricksServiceExceptionWithDetailsProto_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"details": reflect.TypeOf(types.Object{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabricksServiceExceptionWithDetailsProto_SdkV2
// only implements ToObjectValue() and Type().
func (m DatabricksServiceExceptionWithDetailsProto_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"details":     m.Details,
			"error_code":  m.ErrorCode,
			"message":     m.Message,
			"stack_trace": m.StackTrace,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabricksServiceExceptionWithDetailsProto_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"details": basetypes.ListType{
				ElemType: types.ObjectType{},
			},
			"error_code":  types.StringType,
			"message":     types.StringType,
			"stack_trace": types.StringType,
		},
	}
}

// GetDetails returns the value of the Details field in DatabricksServiceExceptionWithDetailsProto_SdkV2 as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabricksServiceExceptionWithDetailsProto_SdkV2) GetDetails(ctx context.Context) ([]types.Object, bool) {
	if m.Details.IsNull() || m.Details.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := m.Details.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDetails sets the value of the Details field in DatabricksServiceExceptionWithDetailsProto_SdkV2.
func (m *DatabricksServiceExceptionWithDetailsProto_SdkV2) SetDetails(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["details"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Details = types.ListValueMust(t, vs)
}

type DeleteBranchRequest_SdkV2 struct {
	// The name of the Branch to delete. Format:
	// projects/{project_id}/branches/{branch_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteBranchRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteBranchRequest_SdkV2) {
}

func (to *DeleteBranchRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteBranchRequest_SdkV2) {
}

func (m DeleteBranchRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteBranchRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteBranchRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteBranchRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteBranchRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteBranchRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteEndpointRequest_SdkV2 struct {
	// The name of the Endpoint to delete. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteEndpointRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteEndpointRequest_SdkV2) {
}

func (to *DeleteEndpointRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteEndpointRequest_SdkV2) {
}

func (m DeleteEndpointRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteProjectRequest_SdkV2 struct {
	// The name of the Project to delete. Format: projects/{project_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteProjectRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteProjectRequest_SdkV2) {
}

func (to *DeleteProjectRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteProjectRequest_SdkV2) {
}

func (m DeleteProjectRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteProjectRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteProjectRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteProjectRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteProjectRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteProjectRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteRoleRequest_SdkV2 struct {
	// The resource name of the postgres role. Format:
	// projects/{project_id}/branches/{branch_id}/roles/{role_id}
	Name types.String `tfsdk:"-"`
	// Reassign objects. If this is set, all objects owned by the role are
	// reassigned to the role specified in this parameter.
	//
	// NOTE: setting this requires spinning up a compute to succeed, since it
	// involves running SQL queries.
	//
	// TODO: #LKB-7187 implement reassign_owned_to on LBM side. This might
	// end-up being a synchronous query when this parameter is used.
	ReassignOwnedTo types.String `tfsdk:"-"`
}

func (to *DeleteRoleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRoleRequest_SdkV2) {
}

func (to *DeleteRoleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteRoleRequest_SdkV2) {
}

func (m DeleteRoleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["reassign_owned_to"] = attrs["reassign_owned_to"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRoleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteRoleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRoleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteRoleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":              m.Name,
			"reassign_owned_to": m.ReassignOwnedTo,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRoleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":              types.StringType,
			"reassign_owned_to": types.StringType,
		},
	}
}

type Endpoint_SdkV2 struct {
	// A timestamp indicating when the compute endpoint was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The resource name of the endpoint. This field is output-only and
	// constructed by the system. Format:
	// `projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}`
	Name types.String `tfsdk:"name"`
	// The branch containing this endpoint (API resource hierarchy). Format:
	// projects/{project_id}/branches/{branch_id}
	Parent types.String `tfsdk:"parent"`
	// The spec contains the compute endpoint configuration, including
	// autoscaling limits, suspend timeout, and disabled state.
	Spec types.List `tfsdk:"spec"`
	// Current operational status of the compute endpoint.
	Status types.List `tfsdk:"status"`
	// System-generated unique ID for the endpoint.
	Uid types.String `tfsdk:"uid"`
	// A timestamp indicating when the compute endpoint was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *Endpoint_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Endpoint_SdkV2) {
	if !from.Spec.IsUnknown() && !from.Spec.IsNull() {
		// Spec is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Spec = from.Spec
	}
	if !from.Spec.IsNull() && !from.Spec.IsUnknown() {
		if toSpec, ok := to.GetSpec(ctx); ok {
			if fromSpec, ok := from.GetSpec(ctx); ok {
				// Recursively sync the fields of Spec
				toSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromSpec)
				to.SetSpec(ctx, toSpec)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				// Recursively sync the fields of Status
				toStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
}

func (to *Endpoint_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Endpoint_SdkV2) {
	if !from.Spec.IsUnknown() && !from.Spec.IsNull() {
		// Spec is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Spec = from.Spec
	}
	if !from.Spec.IsNull() && !from.Spec.IsUnknown() {
		if toSpec, ok := to.GetSpec(ctx); ok {
			if fromSpec, ok := from.GetSpec(ctx); ok {
				toSpec.SyncFieldsDuringRead(ctx, fromSpec)
				to.SetSpec(ctx, toSpec)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				toStatus.SyncFieldsDuringRead(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
}

func (m Endpoint_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["parent"] = attrs["parent"].SetComputed()
	attrs["spec"] = attrs["spec"].SetOptional()
	attrs["spec"] = attrs["spec"].SetComputed()
	attrs["spec"] = attrs["spec"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["spec"] = attrs["spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetComputed()
	attrs["status"] = attrs["status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["uid"] = attrs["uid"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Endpoint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Endpoint_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spec":   reflect.TypeOf(EndpointSpec_SdkV2{}),
		"status": reflect.TypeOf(EndpointStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Endpoint_SdkV2
// only implements ToObjectValue() and Type().
func (m Endpoint_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time": m.CreateTime,
			"name":        m.Name,
			"parent":      m.Parent,
			"spec":        m.Spec,
			"status":      m.Status,
			"uid":         m.Uid,
			"update_time": m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Endpoint_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": timetypes.RFC3339{}.Type(ctx),
			"name":        types.StringType,
			"parent":      types.StringType,
			"spec": basetypes.ListType{
				ElemType: EndpointSpec_SdkV2{}.Type(ctx),
			},
			"status": basetypes.ListType{
				ElemType: EndpointStatus_SdkV2{}.Type(ctx),
			},
			"uid":         types.StringType,
			"update_time": timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetSpec returns the value of the Spec field in Endpoint_SdkV2 as
// a EndpointSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint_SdkV2) GetSpec(ctx context.Context) (EndpointSpec_SdkV2, bool) {
	var e EndpointSpec_SdkV2
	if m.Spec.IsNull() || m.Spec.IsUnknown() {
		return e, false
	}
	var v []EndpointSpec_SdkV2
	d := m.Spec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSpec sets the value of the Spec field in Endpoint_SdkV2.
func (m *Endpoint_SdkV2) SetSpec(ctx context.Context, v EndpointSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spec"]
	m.Spec = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in Endpoint_SdkV2 as
// a EndpointStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint_SdkV2) GetStatus(ctx context.Context) (EndpointStatus_SdkV2, bool) {
	var e EndpointStatus_SdkV2
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v []EndpointStatus_SdkV2
	d := m.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in Endpoint_SdkV2.
func (m *Endpoint_SdkV2) SetStatus(ctx context.Context, v EndpointStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	m.Status = types.ListValueMust(t, vs)
}

// Encapsulates various hostnames (r/w or r/o, pooled or not) for an endpoint.
type EndpointHosts_SdkV2 struct {
	// The hostname to connect to this endpoint. For read-write endpoints, this
	// is a read-write hostname which connects to the primary compute. For
	// read-only endpoints, this is a read-only hostname which allows read-only
	// operations.
	Host types.String `tfsdk:"host"`
}

func (to *EndpointHosts_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointHosts_SdkV2) {
}

func (to *EndpointHosts_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointHosts_SdkV2) {
}

func (m EndpointHosts_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["host"] = attrs["host"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointHosts.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointHosts_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointHosts_SdkV2
// only implements ToObjectValue() and Type().
func (m EndpointHosts_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"host": m.Host,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointHosts_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"host": types.StringType,
		},
	}
}

type EndpointOperationMetadata_SdkV2 struct {
}

func (to *EndpointOperationMetadata_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointOperationMetadata_SdkV2) {
}

func (to *EndpointOperationMetadata_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointOperationMetadata_SdkV2) {
}

func (m EndpointOperationMetadata_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointOperationMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointOperationMetadata_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointOperationMetadata_SdkV2
// only implements ToObjectValue() and Type().
func (m EndpointOperationMetadata_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointOperationMetadata_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// A collection of settings for a compute endpoint.
type EndpointSettings_SdkV2 struct {
	// A raw representation of Postgres settings.
	PgSettings types.Map `tfsdk:"pg_settings"`
}

func (to *EndpointSettings_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointSettings_SdkV2) {
}

func (to *EndpointSettings_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointSettings_SdkV2) {
}

func (m EndpointSettings_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["pg_settings"] = attrs["pg_settings"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointSettings_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"pg_settings": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointSettings_SdkV2
// only implements ToObjectValue() and Type().
func (m EndpointSettings_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pg_settings": m.PgSettings,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointSettings_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pg_settings": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetPgSettings returns the value of the PgSettings field in EndpointSettings_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointSettings_SdkV2) GetPgSettings(ctx context.Context) (map[string]types.String, bool) {
	if m.PgSettings.IsNull() || m.PgSettings.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.PgSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPgSettings sets the value of the PgSettings field in EndpointSettings_SdkV2.
func (m *EndpointSettings_SdkV2) SetPgSettings(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["pg_settings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PgSettings = types.MapValueMust(t, vs)
}

type EndpointSpec_SdkV2 struct {
	// The maximum number of Compute Units. Minimum value is 0.5.
	AutoscalingLimitMaxCu types.Float64 `tfsdk:"autoscaling_limit_max_cu"`
	// The minimum number of Compute Units. Minimum value is 0.5.
	AutoscalingLimitMinCu types.Float64 `tfsdk:"autoscaling_limit_min_cu"`
	// Whether to restrict connections to the compute endpoint. Enabling this
	// option schedules a suspend compute operation. A disabled compute endpoint
	// cannot be enabled by a connection or console action.
	Disabled types.Bool `tfsdk:"disabled"`
	// The endpoint type. A branch can only have one READ_WRITE endpoint.
	EndpointType types.String `tfsdk:"endpoint_type"`
	// When set to true, explicitly disables automatic suspension (never
	// suspend). Should be set to true when provided.
	NoSuspension types.Bool `tfsdk:"no_suspension"`

	Settings types.List `tfsdk:"settings"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended. If specified should be between 60s and 604800s (1 minute to 1
	// week).
	SuspendTimeoutDuration timetypes.GoDuration `tfsdk:"suspend_timeout_duration"`
}

func (to *EndpointSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointSpec_SdkV2) {
	if !from.Settings.IsNull() && !from.Settings.IsUnknown() {
		if toSettings, ok := to.GetSettings(ctx); ok {
			if fromSettings, ok := from.GetSettings(ctx); ok {
				// Recursively sync the fields of Settings
				toSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromSettings)
				to.SetSettings(ctx, toSettings)
			}
		}
	}
}

func (to *EndpointSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointSpec_SdkV2) {
	if !from.Settings.IsNull() && !from.Settings.IsUnknown() {
		if toSettings, ok := to.GetSettings(ctx); ok {
			if fromSettings, ok := from.GetSettings(ctx); ok {
				toSettings.SyncFieldsDuringRead(ctx, fromSettings)
				to.SetSettings(ctx, toSettings)
			}
		}
	}
}

func (m EndpointSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autoscaling_limit_max_cu"] = attrs["autoscaling_limit_max_cu"].SetOptional()
	attrs["autoscaling_limit_min_cu"] = attrs["autoscaling_limit_min_cu"].SetOptional()
	attrs["disabled"] = attrs["disabled"].SetOptional()
	attrs["endpoint_type"] = attrs["endpoint_type"].SetRequired()
	attrs["endpoint_type"] = attrs["endpoint_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["no_suspension"] = attrs["no_suspension"].SetOptional()
	attrs["settings"] = attrs["settings"].SetOptional()
	attrs["settings"] = attrs["settings"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["suspend_timeout_duration"] = attrs["suspend_timeout_duration"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"settings": reflect.TypeOf(EndpointSettings_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m EndpointSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autoscaling_limit_max_cu": m.AutoscalingLimitMaxCu,
			"autoscaling_limit_min_cu": m.AutoscalingLimitMinCu,
			"disabled":                 m.Disabled,
			"endpoint_type":            m.EndpointType,
			"no_suspension":            m.NoSuspension,
			"settings":                 m.Settings,
			"suspend_timeout_duration": m.SuspendTimeoutDuration,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscaling_limit_max_cu": types.Float64Type,
			"autoscaling_limit_min_cu": types.Float64Type,
			"disabled":                 types.BoolType,
			"endpoint_type":            types.StringType,
			"no_suspension":            types.BoolType,
			"settings": basetypes.ListType{
				ElemType: EndpointSettings_SdkV2{}.Type(ctx),
			},
			"suspend_timeout_duration": timetypes.GoDuration{}.Type(ctx),
		},
	}
}

// GetSettings returns the value of the Settings field in EndpointSpec_SdkV2 as
// a EndpointSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointSpec_SdkV2) GetSettings(ctx context.Context) (EndpointSettings_SdkV2, bool) {
	var e EndpointSettings_SdkV2
	if m.Settings.IsNull() || m.Settings.IsUnknown() {
		return e, false
	}
	var v []EndpointSettings_SdkV2
	d := m.Settings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSettings sets the value of the Settings field in EndpointSpec_SdkV2.
func (m *EndpointSpec_SdkV2) SetSettings(ctx context.Context, v EndpointSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["settings"]
	m.Settings = types.ListValueMust(t, vs)
}

type EndpointStatus_SdkV2 struct {
	// The maximum number of Compute Units.
	AutoscalingLimitMaxCu types.Float64 `tfsdk:"autoscaling_limit_max_cu"`
	// The minimum number of Compute Units.
	AutoscalingLimitMinCu types.Float64 `tfsdk:"autoscaling_limit_min_cu"`

	CurrentState types.String `tfsdk:"current_state"`
	// Whether to restrict connections to the compute endpoint. Enabling this
	// option schedules a suspend compute operation. A disabled compute endpoint
	// cannot be enabled by a connection or console action.
	Disabled types.Bool `tfsdk:"disabled"`
	// The endpoint type. A branch can only have one READ_WRITE endpoint.
	EndpointType types.String `tfsdk:"endpoint_type"`
	// Contains host information for connecting to the endpoint.
	Hosts types.List `tfsdk:"hosts"`

	PendingState types.String `tfsdk:"pending_state"`

	Settings types.List `tfsdk:"settings"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended.
	SuspendTimeoutDuration timetypes.GoDuration `tfsdk:"suspend_timeout_duration"`
}

func (to *EndpointStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointStatus_SdkV2) {
	if !from.Hosts.IsNull() && !from.Hosts.IsUnknown() {
		if toHosts, ok := to.GetHosts(ctx); ok {
			if fromHosts, ok := from.GetHosts(ctx); ok {
				// Recursively sync the fields of Hosts
				toHosts.SyncFieldsDuringCreateOrUpdate(ctx, fromHosts)
				to.SetHosts(ctx, toHosts)
			}
		}
	}
	if !from.Settings.IsNull() && !from.Settings.IsUnknown() {
		if toSettings, ok := to.GetSettings(ctx); ok {
			if fromSettings, ok := from.GetSettings(ctx); ok {
				// Recursively sync the fields of Settings
				toSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromSettings)
				to.SetSettings(ctx, toSettings)
			}
		}
	}
}

func (to *EndpointStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointStatus_SdkV2) {
	if !from.Hosts.IsNull() && !from.Hosts.IsUnknown() {
		if toHosts, ok := to.GetHosts(ctx); ok {
			if fromHosts, ok := from.GetHosts(ctx); ok {
				toHosts.SyncFieldsDuringRead(ctx, fromHosts)
				to.SetHosts(ctx, toHosts)
			}
		}
	}
	if !from.Settings.IsNull() && !from.Settings.IsUnknown() {
		if toSettings, ok := to.GetSettings(ctx); ok {
			if fromSettings, ok := from.GetSettings(ctx); ok {
				toSettings.SyncFieldsDuringRead(ctx, fromSettings)
				to.SetSettings(ctx, toSettings)
			}
		}
	}
}

func (m EndpointStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autoscaling_limit_max_cu"] = attrs["autoscaling_limit_max_cu"].SetComputed()
	attrs["autoscaling_limit_min_cu"] = attrs["autoscaling_limit_min_cu"].SetComputed()
	attrs["current_state"] = attrs["current_state"].SetComputed()
	attrs["disabled"] = attrs["disabled"].SetComputed()
	attrs["endpoint_type"] = attrs["endpoint_type"].SetComputed()
	attrs["hosts"] = attrs["hosts"].SetComputed()
	attrs["hosts"] = attrs["hosts"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["pending_state"] = attrs["pending_state"].SetComputed()
	attrs["settings"] = attrs["settings"].SetComputed()
	attrs["settings"] = attrs["settings"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["suspend_timeout_duration"] = attrs["suspend_timeout_duration"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"hosts":    reflect.TypeOf(EndpointHosts_SdkV2{}),
		"settings": reflect.TypeOf(EndpointSettings_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointStatus_SdkV2
// only implements ToObjectValue() and Type().
func (m EndpointStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autoscaling_limit_max_cu": m.AutoscalingLimitMaxCu,
			"autoscaling_limit_min_cu": m.AutoscalingLimitMinCu,
			"current_state":            m.CurrentState,
			"disabled":                 m.Disabled,
			"endpoint_type":            m.EndpointType,
			"hosts":                    m.Hosts,
			"pending_state":            m.PendingState,
			"settings":                 m.Settings,
			"suspend_timeout_duration": m.SuspendTimeoutDuration,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscaling_limit_max_cu": types.Float64Type,
			"autoscaling_limit_min_cu": types.Float64Type,
			"current_state":            types.StringType,
			"disabled":                 types.BoolType,
			"endpoint_type":            types.StringType,
			"hosts": basetypes.ListType{
				ElemType: EndpointHosts_SdkV2{}.Type(ctx),
			},
			"pending_state": types.StringType,
			"settings": basetypes.ListType{
				ElemType: EndpointSettings_SdkV2{}.Type(ctx),
			},
			"suspend_timeout_duration": timetypes.GoDuration{}.Type(ctx),
		},
	}
}

// GetHosts returns the value of the Hosts field in EndpointStatus_SdkV2 as
// a EndpointHosts_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointStatus_SdkV2) GetHosts(ctx context.Context) (EndpointHosts_SdkV2, bool) {
	var e EndpointHosts_SdkV2
	if m.Hosts.IsNull() || m.Hosts.IsUnknown() {
		return e, false
	}
	var v []EndpointHosts_SdkV2
	d := m.Hosts.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHosts sets the value of the Hosts field in EndpointStatus_SdkV2.
func (m *EndpointStatus_SdkV2) SetHosts(ctx context.Context, v EndpointHosts_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["hosts"]
	m.Hosts = types.ListValueMust(t, vs)
}

// GetSettings returns the value of the Settings field in EndpointStatus_SdkV2 as
// a EndpointSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointStatus_SdkV2) GetSettings(ctx context.Context) (EndpointSettings_SdkV2, bool) {
	var e EndpointSettings_SdkV2
	if m.Settings.IsNull() || m.Settings.IsUnknown() {
		return e, false
	}
	var v []EndpointSettings_SdkV2
	d := m.Settings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSettings sets the value of the Settings field in EndpointStatus_SdkV2.
func (m *EndpointStatus_SdkV2) SetSettings(ctx context.Context, v EndpointSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["settings"]
	m.Settings = types.ListValueMust(t, vs)
}

type GenerateDatabaseCredentialRequest_SdkV2 struct {
	// The returned token will be scoped to UC tables with the specified
	// permissions.
	Claims types.List `tfsdk:"claims"`
	// This field is not yet supported. The endpoint for which this credential
	// will be generated. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Endpoint types.String `tfsdk:"endpoint"`
}

func (to *GenerateDatabaseCredentialRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenerateDatabaseCredentialRequest_SdkV2) {
	if !from.Claims.IsNull() && !from.Claims.IsUnknown() && to.Claims.IsNull() && len(from.Claims.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Claims, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Claims = from.Claims
	}
}

func (to *GenerateDatabaseCredentialRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GenerateDatabaseCredentialRequest_SdkV2) {
	if !from.Claims.IsNull() && !from.Claims.IsUnknown() && to.Claims.IsNull() && len(from.Claims.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Claims, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Claims = from.Claims
	}
}

func (m GenerateDatabaseCredentialRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["claims"] = attrs["claims"].SetOptional()
	attrs["endpoint"] = attrs["endpoint"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenerateDatabaseCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenerateDatabaseCredentialRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"claims": reflect.TypeOf(RequestedClaims_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenerateDatabaseCredentialRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GenerateDatabaseCredentialRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"claims":   m.Claims,
			"endpoint": m.Endpoint,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenerateDatabaseCredentialRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"claims": basetypes.ListType{
				ElemType: RequestedClaims_SdkV2{}.Type(ctx),
			},
			"endpoint": types.StringType,
		},
	}
}

// GetClaims returns the value of the Claims field in GenerateDatabaseCredentialRequest_SdkV2 as
// a slice of RequestedClaims_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenerateDatabaseCredentialRequest_SdkV2) GetClaims(ctx context.Context) ([]RequestedClaims_SdkV2, bool) {
	if m.Claims.IsNull() || m.Claims.IsUnknown() {
		return nil, false
	}
	var v []RequestedClaims_SdkV2
	d := m.Claims.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClaims sets the value of the Claims field in GenerateDatabaseCredentialRequest_SdkV2.
func (m *GenerateDatabaseCredentialRequest_SdkV2) SetClaims(ctx context.Context, v []RequestedClaims_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["claims"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Claims = types.ListValueMust(t, vs)
}

type GetBranchRequest_SdkV2 struct {
	// The resource name of the branch to retrieve. Format:
	// `projects/{project_id}/branches/{branch_id}`
	Name types.String `tfsdk:"-"`
}

func (to *GetBranchRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetBranchRequest_SdkV2) {
}

func (to *GetBranchRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetBranchRequest_SdkV2) {
}

func (m GetBranchRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetBranchRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetBranchRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetBranchRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetBranchRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetBranchRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetEndpointRequest_SdkV2 struct {
	// The resource name of the endpoint to retrieve. Format:
	// `projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}`
	Name types.String `tfsdk:"-"`
}

func (to *GetEndpointRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEndpointRequest_SdkV2) {
}

func (to *GetEndpointRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetEndpointRequest_SdkV2) {
}

func (m GetEndpointRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetOperationRequest_SdkV2 struct {
	// The name of the operation resource.
	Name types.String `tfsdk:"-"`
}

func (to *GetOperationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetOperationRequest_SdkV2) {
}

func (to *GetOperationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetOperationRequest_SdkV2) {
}

func (m GetOperationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetOperationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetOperationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetOperationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetOperationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetOperationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetProjectRequest_SdkV2 struct {
	// The resource name of the project to retrieve. Format:
	// `projects/{project_id}`
	Name types.String `tfsdk:"-"`
}

func (to *GetProjectRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetProjectRequest_SdkV2) {
}

func (to *GetProjectRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetProjectRequest_SdkV2) {
}

func (m GetProjectRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetProjectRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetProjectRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetProjectRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetProjectRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetProjectRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetRoleRequest_SdkV2 struct {
	// The name of the Role to retrieve. Format:
	// projects/{project_id}/branches/{branch_id}/roles/{role_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetRoleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRoleRequest_SdkV2) {
}

func (to *GetRoleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetRoleRequest_SdkV2) {
}

func (m GetRoleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRoleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetRoleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRoleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetRoleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRoleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ListBranchesRequest_SdkV2 struct {
	// Upper bound for items returned. Cannot be negative.
	PageSize types.Int64 `tfsdk:"-"`
	// Page token from a previous response. If not provided, returns the first
	// page.
	PageToken types.String `tfsdk:"-"`
	// The Project that owns this collection of branches. Format:
	// projects/{project_id}
	Parent types.String `tfsdk:"-"`
}

func (to *ListBranchesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListBranchesRequest_SdkV2) {
}

func (to *ListBranchesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListBranchesRequest_SdkV2) {
}

func (m ListBranchesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListBranchesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListBranchesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListBranchesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListBranchesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListBranchesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListBranchesResponse_SdkV2 struct {
	// List of database branches in the project.
	Branches types.List `tfsdk:"branches"`
	// Token to request the next page of database branches.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListBranchesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListBranchesResponse_SdkV2) {
	if !from.Branches.IsNull() && !from.Branches.IsUnknown() && to.Branches.IsNull() && len(from.Branches.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Branches, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Branches = from.Branches
	}
}

func (to *ListBranchesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListBranchesResponse_SdkV2) {
	if !from.Branches.IsNull() && !from.Branches.IsUnknown() && to.Branches.IsNull() && len(from.Branches.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Branches, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Branches = from.Branches
	}
}

func (m ListBranchesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branches"] = attrs["branches"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListBranchesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListBranchesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"branches": reflect.TypeOf(Branch_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListBranchesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListBranchesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branches":        m.Branches,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListBranchesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branches": basetypes.ListType{
				ElemType: Branch_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetBranches returns the value of the Branches field in ListBranchesResponse_SdkV2 as
// a slice of Branch_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListBranchesResponse_SdkV2) GetBranches(ctx context.Context) ([]Branch_SdkV2, bool) {
	if m.Branches.IsNull() || m.Branches.IsUnknown() {
		return nil, false
	}
	var v []Branch_SdkV2
	d := m.Branches.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBranches sets the value of the Branches field in ListBranchesResponse_SdkV2.
func (m *ListBranchesResponse_SdkV2) SetBranches(ctx context.Context, v []Branch_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["branches"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Branches = types.ListValueMust(t, vs)
}

type ListEndpointsRequest_SdkV2 struct {
	// Upper bound for items returned. Cannot be negative.
	PageSize types.Int64 `tfsdk:"-"`
	// Page token from a previous response. If not provided, returns the first
	// page.
	PageToken types.String `tfsdk:"-"`
	// The Branch that owns this collection of endpoints. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent types.String `tfsdk:"-"`
}

func (to *ListEndpointsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEndpointsRequest_SdkV2) {
}

func (to *ListEndpointsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListEndpointsRequest_SdkV2) {
}

func (m ListEndpointsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListEndpointsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListEndpointsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListEndpointsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListEndpointsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListEndpointsResponse_SdkV2 struct {
	// List of compute endpoints in the branch.
	Endpoints types.List `tfsdk:"endpoints"`
	// Token to request the next page of compute endpoints.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListEndpointsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEndpointsResponse_SdkV2) {
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() && to.Endpoints.IsNull() && len(from.Endpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Endpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Endpoints = from.Endpoints
	}
}

func (to *ListEndpointsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListEndpointsResponse_SdkV2) {
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() && to.Endpoints.IsNull() && len(from.Endpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Endpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Endpoints = from.Endpoints
	}
}

func (m ListEndpointsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoints"] = attrs["endpoints"].SetComputed()
	attrs["next_page_token"] = attrs["next_page_token"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListEndpointsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListEndpointsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoints": reflect.TypeOf(Endpoint_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListEndpointsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoints":       m.Endpoints,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListEndpointsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoints": basetypes.ListType{
				ElemType: Endpoint_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetEndpoints returns the value of the Endpoints field in ListEndpointsResponse_SdkV2 as
// a slice of Endpoint_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListEndpointsResponse_SdkV2) GetEndpoints(ctx context.Context) ([]Endpoint_SdkV2, bool) {
	if m.Endpoints.IsNull() || m.Endpoints.IsUnknown() {
		return nil, false
	}
	var v []Endpoint_SdkV2
	d := m.Endpoints.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEndpoints sets the value of the Endpoints field in ListEndpointsResponse_SdkV2.
func (m *ListEndpointsResponse_SdkV2) SetEndpoints(ctx context.Context, v []Endpoint_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["endpoints"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Endpoints = types.ListValueMust(t, vs)
}

type ListProjectsRequest_SdkV2 struct {
	// Upper bound for items returned. Cannot be negative.
	PageSize types.Int64 `tfsdk:"-"`
	// Page token from a previous response. If not provided, returns the first
	// page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListProjectsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProjectsRequest_SdkV2) {
}

func (to *ListProjectsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListProjectsRequest_SdkV2) {
}

func (m ListProjectsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProjectsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListProjectsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProjectsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListProjectsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListProjectsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListProjectsResponse_SdkV2 struct {
	// Token to request the next page of database projects.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of all database projects in the workspace that the user has
	// permission to access.
	Projects types.List `tfsdk:"projects"`
}

func (to *ListProjectsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProjectsResponse_SdkV2) {
	if !from.Projects.IsNull() && !from.Projects.IsUnknown() && to.Projects.IsNull() && len(from.Projects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Projects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Projects = from.Projects
	}
}

func (to *ListProjectsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListProjectsResponse_SdkV2) {
	if !from.Projects.IsNull() && !from.Projects.IsUnknown() && to.Projects.IsNull() && len(from.Projects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Projects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Projects = from.Projects
	}
}

func (m ListProjectsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["projects"] = attrs["projects"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListProjectsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListProjectsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"projects": reflect.TypeOf(Project_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProjectsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListProjectsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"projects":        m.Projects,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListProjectsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"projects": basetypes.ListType{
				ElemType: Project_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetProjects returns the value of the Projects field in ListProjectsResponse_SdkV2 as
// a slice of Project_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListProjectsResponse_SdkV2) GetProjects(ctx context.Context) ([]Project_SdkV2, bool) {
	if m.Projects.IsNull() || m.Projects.IsUnknown() {
		return nil, false
	}
	var v []Project_SdkV2
	d := m.Projects.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProjects sets the value of the Projects field in ListProjectsResponse_SdkV2.
func (m *ListProjectsResponse_SdkV2) SetProjects(ctx context.Context, v []Project_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["projects"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Projects = types.ListValueMust(t, vs)
}

type ListRolesRequest_SdkV2 struct {
	// Upper bound for items returned. Cannot be negative.
	PageSize types.Int64 `tfsdk:"-"`
	// Page token from a previous response. If not provided, returns the first
	// page.
	PageToken types.String `tfsdk:"-"`
	// The Branch that owns this collection of roles. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent types.String `tfsdk:"-"`
}

func (to *ListRolesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListRolesRequest_SdkV2) {
}

func (to *ListRolesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListRolesRequest_SdkV2) {
}

func (m ListRolesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRolesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListRolesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRolesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListRolesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListRolesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListRolesResponse_SdkV2 struct {
	// Token to request the next page of Postgres roles.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of Postgres roles in the branch.
	Roles types.List `tfsdk:"roles"`
}

func (to *ListRolesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListRolesResponse_SdkV2) {
	if !from.Roles.IsNull() && !from.Roles.IsUnknown() && to.Roles.IsNull() && len(from.Roles.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Roles, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Roles = from.Roles
	}
}

func (to *ListRolesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListRolesResponse_SdkV2) {
	if !from.Roles.IsNull() && !from.Roles.IsUnknown() && to.Roles.IsNull() && len(from.Roles.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Roles, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Roles = from.Roles
	}
}

func (m ListRolesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["roles"] = attrs["roles"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRolesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListRolesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"roles": reflect.TypeOf(Role_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRolesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListRolesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"roles":           m.Roles,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListRolesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"roles": basetypes.ListType{
				ElemType: Role_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRoles returns the value of the Roles field in ListRolesResponse_SdkV2 as
// a slice of Role_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListRolesResponse_SdkV2) GetRoles(ctx context.Context) ([]Role_SdkV2, bool) {
	if m.Roles.IsNull() || m.Roles.IsUnknown() {
		return nil, false
	}
	var v []Role_SdkV2
	d := m.Roles.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRoles sets the value of the Roles field in ListRolesResponse_SdkV2.
func (m *ListRolesResponse_SdkV2) SetRoles(ctx context.Context, v []Role_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["roles"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Roles = types.ListValueMust(t, vs)
}

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation_SdkV2 struct {
	// If the value is `false`, it means the operation is still in progress. If
	// `true`, the operation is completed, and either `error` or `response` is
	// available.
	Done types.Bool `tfsdk:"done"`
	// The error result of the operation in case of failure or cancellation.
	Error types.List `tfsdk:"error"`
	// Service-specific metadata associated with the operation. It typically
	// contains progress information and common metadata such as create time.
	// Some services might not provide such metadata.
	Metadata types.Object `tfsdk:"metadata"`
	// The server-assigned name, which is only unique within the same service
	// that originally returns it. If you use the default HTTP mapping, the
	// `name` should be a resource name ending with `operations/{unique_id}`.
	Name types.String `tfsdk:"name"`
	// The normal, successful response of the operation.
	Response types.Object `tfsdk:"response"`
}

func (to *Operation_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Operation_SdkV2) {
	if !from.Error.IsNull() && !from.Error.IsUnknown() {
		if toError, ok := to.GetError(ctx); ok {
			if fromError, ok := from.GetError(ctx); ok {
				// Recursively sync the fields of Error
				toError.SyncFieldsDuringCreateOrUpdate(ctx, fromError)
				to.SetError(ctx, toError)
			}
		}
	}
}

func (to *Operation_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Operation_SdkV2) {
	if !from.Error.IsNull() && !from.Error.IsUnknown() {
		if toError, ok := to.GetError(ctx); ok {
			if fromError, ok := from.GetError(ctx); ok {
				toError.SyncFieldsDuringRead(ctx, fromError)
				to.SetError(ctx, toError)
			}
		}
	}
}

func (m Operation_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["done"] = attrs["done"].SetOptional()
	attrs["error"] = attrs["error"].SetOptional()
	attrs["error"] = attrs["error"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["metadata"] = attrs["metadata"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["response"] = attrs["response"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Operation.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Operation_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error": reflect.TypeOf(DatabricksServiceExceptionWithDetailsProto_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Operation_SdkV2
// only implements ToObjectValue() and Type().
func (m Operation_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"done":     m.Done,
			"error":    m.Error,
			"metadata": m.Metadata,
			"name":     m.Name,
			"response": m.Response,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Operation_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"done": types.BoolType,
			"error": basetypes.ListType{
				ElemType: DatabricksServiceExceptionWithDetailsProto_SdkV2{}.Type(ctx),
			},
			"metadata": types.ObjectType{},
			"name":     types.StringType,
			"response": types.ObjectType{},
		},
	}
}

// GetError returns the value of the Error field in Operation_SdkV2 as
// a DatabricksServiceExceptionWithDetailsProto_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Operation_SdkV2) GetError(ctx context.Context) (DatabricksServiceExceptionWithDetailsProto_SdkV2, bool) {
	var e DatabricksServiceExceptionWithDetailsProto_SdkV2
	if m.Error.IsNull() || m.Error.IsUnknown() {
		return e, false
	}
	var v []DatabricksServiceExceptionWithDetailsProto_SdkV2
	d := m.Error.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetError sets the value of the Error field in Operation_SdkV2.
func (m *Operation_SdkV2) SetError(ctx context.Context, v DatabricksServiceExceptionWithDetailsProto_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["error"]
	m.Error = types.ListValueMust(t, vs)
}

type Project_SdkV2 struct {
	// A timestamp indicating when the project was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The resource name of the project. This field is output-only and
	// constructed by the system. Format: `projects/{project_id}`
	Name types.String `tfsdk:"name"`
	// The spec contains the project configuration, including display_name,
	// pg_version (Postgres version), history_retention_duration, and
	// default_endpoint_settings.
	Spec types.List `tfsdk:"spec"`
	// The current status of a Project.
	Status types.List `tfsdk:"status"`
	// System-generated unique ID for the project.
	Uid types.String `tfsdk:"uid"`
	// A timestamp indicating when the project was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *Project_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Project_SdkV2) {
	if !from.Spec.IsUnknown() && !from.Spec.IsNull() {
		// Spec is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Spec = from.Spec
	}
	if !from.Spec.IsNull() && !from.Spec.IsUnknown() {
		if toSpec, ok := to.GetSpec(ctx); ok {
			if fromSpec, ok := from.GetSpec(ctx); ok {
				// Recursively sync the fields of Spec
				toSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromSpec)
				to.SetSpec(ctx, toSpec)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				// Recursively sync the fields of Status
				toStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
}

func (to *Project_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Project_SdkV2) {
	if !from.Spec.IsUnknown() && !from.Spec.IsNull() {
		// Spec is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Spec = from.Spec
	}
	if !from.Spec.IsNull() && !from.Spec.IsUnknown() {
		if toSpec, ok := to.GetSpec(ctx); ok {
			if fromSpec, ok := from.GetSpec(ctx); ok {
				toSpec.SyncFieldsDuringRead(ctx, fromSpec)
				to.SetSpec(ctx, toSpec)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				toStatus.SyncFieldsDuringRead(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
}

func (m Project_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["spec"] = attrs["spec"].SetOptional()
	attrs["spec"] = attrs["spec"].SetComputed()
	attrs["spec"] = attrs["spec"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["spec"] = attrs["spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetComputed()
	attrs["status"] = attrs["status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["uid"] = attrs["uid"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Project.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Project_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spec":   reflect.TypeOf(ProjectSpec_SdkV2{}),
		"status": reflect.TypeOf(ProjectStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Project_SdkV2
// only implements ToObjectValue() and Type().
func (m Project_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time": m.CreateTime,
			"name":        m.Name,
			"spec":        m.Spec,
			"status":      m.Status,
			"uid":         m.Uid,
			"update_time": m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Project_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": timetypes.RFC3339{}.Type(ctx),
			"name":        types.StringType,
			"spec": basetypes.ListType{
				ElemType: ProjectSpec_SdkV2{}.Type(ctx),
			},
			"status": basetypes.ListType{
				ElemType: ProjectStatus_SdkV2{}.Type(ctx),
			},
			"uid":         types.StringType,
			"update_time": timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetSpec returns the value of the Spec field in Project_SdkV2 as
// a ProjectSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Project_SdkV2) GetSpec(ctx context.Context) (ProjectSpec_SdkV2, bool) {
	var e ProjectSpec_SdkV2
	if m.Spec.IsNull() || m.Spec.IsUnknown() {
		return e, false
	}
	var v []ProjectSpec_SdkV2
	d := m.Spec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSpec sets the value of the Spec field in Project_SdkV2.
func (m *Project_SdkV2) SetSpec(ctx context.Context, v ProjectSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spec"]
	m.Spec = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in Project_SdkV2 as
// a ProjectStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Project_SdkV2) GetStatus(ctx context.Context) (ProjectStatus_SdkV2, bool) {
	var e ProjectStatus_SdkV2
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v []ProjectStatus_SdkV2
	d := m.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in Project_SdkV2.
func (m *Project_SdkV2) SetStatus(ctx context.Context, v ProjectStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	m.Status = types.ListValueMust(t, vs)
}

// A collection of settings for a compute endpoint.
type ProjectDefaultEndpointSettings_SdkV2 struct {
	// The maximum number of Compute Units. Minimum value is 0.5.
	AutoscalingLimitMaxCu types.Float64 `tfsdk:"autoscaling_limit_max_cu"`
	// The minimum number of Compute Units. Minimum value is 0.5.
	AutoscalingLimitMinCu types.Float64 `tfsdk:"autoscaling_limit_min_cu"`
	// When set to true, explicitly disables automatic suspension (never
	// suspend). Should be set to true when provided.
	NoSuspension types.Bool `tfsdk:"no_suspension"`
	// A raw representation of Postgres settings.
	PgSettings types.Map `tfsdk:"pg_settings"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended. If specified should be between 60s and 604800s (1 minute to 1
	// week).
	SuspendTimeoutDuration timetypes.GoDuration `tfsdk:"suspend_timeout_duration"`
}

func (to *ProjectDefaultEndpointSettings_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProjectDefaultEndpointSettings_SdkV2) {
}

func (to *ProjectDefaultEndpointSettings_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ProjectDefaultEndpointSettings_SdkV2) {
}

func (m ProjectDefaultEndpointSettings_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autoscaling_limit_max_cu"] = attrs["autoscaling_limit_max_cu"].SetOptional()
	attrs["autoscaling_limit_min_cu"] = attrs["autoscaling_limit_min_cu"].SetOptional()
	attrs["no_suspension"] = attrs["no_suspension"].SetOptional()
	attrs["pg_settings"] = attrs["pg_settings"].SetOptional()
	attrs["suspend_timeout_duration"] = attrs["suspend_timeout_duration"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProjectDefaultEndpointSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ProjectDefaultEndpointSettings_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"pg_settings": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProjectDefaultEndpointSettings_SdkV2
// only implements ToObjectValue() and Type().
func (m ProjectDefaultEndpointSettings_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autoscaling_limit_max_cu": m.AutoscalingLimitMaxCu,
			"autoscaling_limit_min_cu": m.AutoscalingLimitMinCu,
			"no_suspension":            m.NoSuspension,
			"pg_settings":              m.PgSettings,
			"suspend_timeout_duration": m.SuspendTimeoutDuration,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ProjectDefaultEndpointSettings_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscaling_limit_max_cu": types.Float64Type,
			"autoscaling_limit_min_cu": types.Float64Type,
			"no_suspension":            types.BoolType,
			"pg_settings": basetypes.MapType{
				ElemType: types.StringType,
			},
			"suspend_timeout_duration": timetypes.GoDuration{}.Type(ctx),
		},
	}
}

// GetPgSettings returns the value of the PgSettings field in ProjectDefaultEndpointSettings_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ProjectDefaultEndpointSettings_SdkV2) GetPgSettings(ctx context.Context) (map[string]types.String, bool) {
	if m.PgSettings.IsNull() || m.PgSettings.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.PgSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPgSettings sets the value of the PgSettings field in ProjectDefaultEndpointSettings_SdkV2.
func (m *ProjectDefaultEndpointSettings_SdkV2) SetPgSettings(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["pg_settings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PgSettings = types.MapValueMust(t, vs)
}

type ProjectOperationMetadata_SdkV2 struct {
}

func (to *ProjectOperationMetadata_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProjectOperationMetadata_SdkV2) {
}

func (to *ProjectOperationMetadata_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ProjectOperationMetadata_SdkV2) {
}

func (m ProjectOperationMetadata_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProjectOperationMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ProjectOperationMetadata_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProjectOperationMetadata_SdkV2
// only implements ToObjectValue() and Type().
func (m ProjectOperationMetadata_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ProjectOperationMetadata_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ProjectSpec_SdkV2 struct {
	DefaultEndpointSettings types.List `tfsdk:"default_endpoint_settings"`
	// Human-readable project name. Length should be between 1 and 256
	// characters.
	DisplayName types.String `tfsdk:"display_name"`
	// The number of seconds to retain the shared history for point in time
	// recovery for all branches in this project. Value should be between 0s and
	// 2592000s (up to 30 days).
	HistoryRetentionDuration timetypes.GoDuration `tfsdk:"history_retention_duration"`
	// The major Postgres version number. Supported versions are 16 and 17.
	PgVersion types.Int64 `tfsdk:"pg_version"`
}

func (to *ProjectSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProjectSpec_SdkV2) {
	if !from.DefaultEndpointSettings.IsNull() && !from.DefaultEndpointSettings.IsUnknown() {
		if toDefaultEndpointSettings, ok := to.GetDefaultEndpointSettings(ctx); ok {
			if fromDefaultEndpointSettings, ok := from.GetDefaultEndpointSettings(ctx); ok {
				// Recursively sync the fields of DefaultEndpointSettings
				toDefaultEndpointSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromDefaultEndpointSettings)
				to.SetDefaultEndpointSettings(ctx, toDefaultEndpointSettings)
			}
		}
	}
}

func (to *ProjectSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ProjectSpec_SdkV2) {
	if !from.DefaultEndpointSettings.IsNull() && !from.DefaultEndpointSettings.IsUnknown() {
		if toDefaultEndpointSettings, ok := to.GetDefaultEndpointSettings(ctx); ok {
			if fromDefaultEndpointSettings, ok := from.GetDefaultEndpointSettings(ctx); ok {
				toDefaultEndpointSettings.SyncFieldsDuringRead(ctx, fromDefaultEndpointSettings)
				to.SetDefaultEndpointSettings(ctx, toDefaultEndpointSettings)
			}
		}
	}
}

func (m ProjectSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_endpoint_settings"] = attrs["default_endpoint_settings"].SetOptional()
	attrs["default_endpoint_settings"] = attrs["default_endpoint_settings"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["history_retention_duration"] = attrs["history_retention_duration"].SetOptional()
	attrs["pg_version"] = attrs["pg_version"].SetOptional()
	attrs["pg_version"] = attrs["pg_version"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProjectSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ProjectSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_endpoint_settings": reflect.TypeOf(ProjectDefaultEndpointSettings_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProjectSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m ProjectSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_endpoint_settings":  m.DefaultEndpointSettings,
			"display_name":               m.DisplayName,
			"history_retention_duration": m.HistoryRetentionDuration,
			"pg_version":                 m.PgVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ProjectSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_endpoint_settings": basetypes.ListType{
				ElemType: ProjectDefaultEndpointSettings_SdkV2{}.Type(ctx),
			},
			"display_name":               types.StringType,
			"history_retention_duration": timetypes.GoDuration{}.Type(ctx),
			"pg_version":                 types.Int64Type,
		},
	}
}

// GetDefaultEndpointSettings returns the value of the DefaultEndpointSettings field in ProjectSpec_SdkV2 as
// a ProjectDefaultEndpointSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ProjectSpec_SdkV2) GetDefaultEndpointSettings(ctx context.Context) (ProjectDefaultEndpointSettings_SdkV2, bool) {
	var e ProjectDefaultEndpointSettings_SdkV2
	if m.DefaultEndpointSettings.IsNull() || m.DefaultEndpointSettings.IsUnknown() {
		return e, false
	}
	var v []ProjectDefaultEndpointSettings_SdkV2
	d := m.DefaultEndpointSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDefaultEndpointSettings sets the value of the DefaultEndpointSettings field in ProjectSpec_SdkV2.
func (m *ProjectSpec_SdkV2) SetDefaultEndpointSettings(ctx context.Context, v ProjectDefaultEndpointSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["default_endpoint_settings"]
	m.DefaultEndpointSettings = types.ListValueMust(t, vs)
}

type ProjectStatus_SdkV2 struct {
	// The logical size limit for a branch.
	BranchLogicalSizeLimitBytes types.Int64 `tfsdk:"branch_logical_size_limit_bytes"`
	// The effective default endpoint settings.
	DefaultEndpointSettings types.List `tfsdk:"default_endpoint_settings"`
	// The effective human-readable project name.
	DisplayName types.String `tfsdk:"display_name"`
	// The effective number of seconds to retain the shared history for point in
	// time recovery.
	HistoryRetentionDuration timetypes.GoDuration `tfsdk:"history_retention_duration"`
	// The email of the project owner.
	Owner types.String `tfsdk:"owner"`
	// The effective major Postgres version number.
	PgVersion types.Int64 `tfsdk:"pg_version"`
	// The current space occupied by the project in storage.
	SyntheticStorageSizeBytes types.Int64 `tfsdk:"synthetic_storage_size_bytes"`
}

func (to *ProjectStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProjectStatus_SdkV2) {
	if !from.DefaultEndpointSettings.IsNull() && !from.DefaultEndpointSettings.IsUnknown() {
		if toDefaultEndpointSettings, ok := to.GetDefaultEndpointSettings(ctx); ok {
			if fromDefaultEndpointSettings, ok := from.GetDefaultEndpointSettings(ctx); ok {
				// Recursively sync the fields of DefaultEndpointSettings
				toDefaultEndpointSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromDefaultEndpointSettings)
				to.SetDefaultEndpointSettings(ctx, toDefaultEndpointSettings)
			}
		}
	}
}

func (to *ProjectStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ProjectStatus_SdkV2) {
	if !from.DefaultEndpointSettings.IsNull() && !from.DefaultEndpointSettings.IsUnknown() {
		if toDefaultEndpointSettings, ok := to.GetDefaultEndpointSettings(ctx); ok {
			if fromDefaultEndpointSettings, ok := from.GetDefaultEndpointSettings(ctx); ok {
				toDefaultEndpointSettings.SyncFieldsDuringRead(ctx, fromDefaultEndpointSettings)
				to.SetDefaultEndpointSettings(ctx, toDefaultEndpointSettings)
			}
		}
	}
}

func (m ProjectStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch_logical_size_limit_bytes"] = attrs["branch_logical_size_limit_bytes"].SetComputed()
	attrs["default_endpoint_settings"] = attrs["default_endpoint_settings"].SetComputed()
	attrs["default_endpoint_settings"] = attrs["default_endpoint_settings"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["history_retention_duration"] = attrs["history_retention_duration"].SetComputed()
	attrs["owner"] = attrs["owner"].SetComputed()
	attrs["pg_version"] = attrs["pg_version"].SetComputed()
	attrs["synthetic_storage_size_bytes"] = attrs["synthetic_storage_size_bytes"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProjectStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ProjectStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_endpoint_settings": reflect.TypeOf(ProjectDefaultEndpointSettings_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProjectStatus_SdkV2
// only implements ToObjectValue() and Type().
func (m ProjectStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch_logical_size_limit_bytes": m.BranchLogicalSizeLimitBytes,
			"default_endpoint_settings":       m.DefaultEndpointSettings,
			"display_name":                    m.DisplayName,
			"history_retention_duration":      m.HistoryRetentionDuration,
			"owner":                           m.Owner,
			"pg_version":                      m.PgVersion,
			"synthetic_storage_size_bytes":    m.SyntheticStorageSizeBytes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ProjectStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch_logical_size_limit_bytes": types.Int64Type,
			"default_endpoint_settings": basetypes.ListType{
				ElemType: ProjectDefaultEndpointSettings_SdkV2{}.Type(ctx),
			},
			"display_name":                 types.StringType,
			"history_retention_duration":   timetypes.GoDuration{}.Type(ctx),
			"owner":                        types.StringType,
			"pg_version":                   types.Int64Type,
			"synthetic_storage_size_bytes": types.Int64Type,
		},
	}
}

// GetDefaultEndpointSettings returns the value of the DefaultEndpointSettings field in ProjectStatus_SdkV2 as
// a ProjectDefaultEndpointSettings_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ProjectStatus_SdkV2) GetDefaultEndpointSettings(ctx context.Context) (ProjectDefaultEndpointSettings_SdkV2, bool) {
	var e ProjectDefaultEndpointSettings_SdkV2
	if m.DefaultEndpointSettings.IsNull() || m.DefaultEndpointSettings.IsUnknown() {
		return e, false
	}
	var v []ProjectDefaultEndpointSettings_SdkV2
	d := m.DefaultEndpointSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDefaultEndpointSettings sets the value of the DefaultEndpointSettings field in ProjectStatus_SdkV2.
func (m *ProjectStatus_SdkV2) SetDefaultEndpointSettings(ctx context.Context, v ProjectDefaultEndpointSettings_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["default_endpoint_settings"]
	m.DefaultEndpointSettings = types.ListValueMust(t, vs)
}

type RequestedClaims_SdkV2 struct {
	PermissionSet types.String `tfsdk:"permission_set"`

	Resources types.List `tfsdk:"resources"`
}

func (to *RequestedClaims_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RequestedClaims_SdkV2) {
	if !from.Resources.IsNull() && !from.Resources.IsUnknown() && to.Resources.IsNull() && len(from.Resources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Resources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Resources = from.Resources
	}
}

func (to *RequestedClaims_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RequestedClaims_SdkV2) {
	if !from.Resources.IsNull() && !from.Resources.IsUnknown() && to.Resources.IsNull() && len(from.Resources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Resources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Resources = from.Resources
	}
}

func (m RequestedClaims_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_set"] = attrs["permission_set"].SetOptional()
	attrs["resources"] = attrs["resources"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RequestedClaims.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RequestedClaims_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"resources": reflect.TypeOf(RequestedResource_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RequestedClaims_SdkV2
// only implements ToObjectValue() and Type().
func (m RequestedClaims_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_set": m.PermissionSet,
			"resources":      m.Resources,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RequestedClaims_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_set": types.StringType,
			"resources": basetypes.ListType{
				ElemType: RequestedResource_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetResources returns the value of the Resources field in RequestedClaims_SdkV2 as
// a slice of RequestedResource_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *RequestedClaims_SdkV2) GetResources(ctx context.Context) ([]RequestedResource_SdkV2, bool) {
	if m.Resources.IsNull() || m.Resources.IsUnknown() {
		return nil, false
	}
	var v []RequestedResource_SdkV2
	d := m.Resources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResources sets the value of the Resources field in RequestedClaims_SdkV2.
func (m *RequestedClaims_SdkV2) SetResources(ctx context.Context, v []RequestedResource_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Resources = types.ListValueMust(t, vs)
}

type RequestedResource_SdkV2 struct {
	TableName types.String `tfsdk:"table_name"`

	UnspecifiedResourceName types.String `tfsdk:"unspecified_resource_name"`
}

func (to *RequestedResource_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RequestedResource_SdkV2) {
}

func (to *RequestedResource_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RequestedResource_SdkV2) {
}

func (m RequestedResource_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["table_name"] = attrs["table_name"].SetOptional()
	attrs["unspecified_resource_name"] = attrs["unspecified_resource_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RequestedResource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RequestedResource_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RequestedResource_SdkV2
// only implements ToObjectValue() and Type().
func (m RequestedResource_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table_name":                m.TableName,
			"unspecified_resource_name": m.UnspecifiedResourceName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RequestedResource_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_name":                types.StringType,
			"unspecified_resource_name": types.StringType,
		},
	}
}

// Role represents a Postgres role within a Branch.
type Role_SdkV2 struct {
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The resource name of the role. Format:
	// projects/{project_id}/branches/{branch_id}/roles/{role_id}
	Name types.String `tfsdk:"name"`
	// The Branch where this Role exists. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent types.String `tfsdk:"parent"`
	// The spec contains the role configuration, including identity type,
	// authentication method, and role attributes.
	Spec types.List `tfsdk:"spec"`
	// Current status of the role, including its identity type, authentication
	// method, and role attributes.
	Status types.List `tfsdk:"status"`

	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *Role_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Role_SdkV2) {
	if !from.Spec.IsUnknown() && !from.Spec.IsNull() {
		// Spec is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Spec = from.Spec
	}
	if !from.Spec.IsNull() && !from.Spec.IsUnknown() {
		if toSpec, ok := to.GetSpec(ctx); ok {
			if fromSpec, ok := from.GetSpec(ctx); ok {
				// Recursively sync the fields of Spec
				toSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromSpec)
				to.SetSpec(ctx, toSpec)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				// Recursively sync the fields of Status
				toStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
}

func (to *Role_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Role_SdkV2) {
	if !from.Spec.IsUnknown() && !from.Spec.IsNull() {
		// Spec is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Spec = from.Spec
	}
	if !from.Spec.IsNull() && !from.Spec.IsUnknown() {
		if toSpec, ok := to.GetSpec(ctx); ok {
			if fromSpec, ok := from.GetSpec(ctx); ok {
				toSpec.SyncFieldsDuringRead(ctx, fromSpec)
				to.SetSpec(ctx, toSpec)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				toStatus.SyncFieldsDuringRead(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
}

func (m Role_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["parent"] = attrs["parent"].SetComputed()
	attrs["spec"] = attrs["spec"].SetOptional()
	attrs["spec"] = attrs["spec"].SetComputed()
	attrs["spec"] = attrs["spec"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["spec"] = attrs["spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetComputed()
	attrs["status"] = attrs["status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Role.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Role_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spec":   reflect.TypeOf(RoleRoleSpec_SdkV2{}),
		"status": reflect.TypeOf(RoleRoleStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Role_SdkV2
// only implements ToObjectValue() and Type().
func (m Role_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time": m.CreateTime,
			"name":        m.Name,
			"parent":      m.Parent,
			"spec":        m.Spec,
			"status":      m.Status,
			"update_time": m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Role_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": timetypes.RFC3339{}.Type(ctx),
			"name":        types.StringType,
			"parent":      types.StringType,
			"spec": basetypes.ListType{
				ElemType: RoleRoleSpec_SdkV2{}.Type(ctx),
			},
			"status": basetypes.ListType{
				ElemType: RoleRoleStatus_SdkV2{}.Type(ctx),
			},
			"update_time": timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetSpec returns the value of the Spec field in Role_SdkV2 as
// a RoleRoleSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Role_SdkV2) GetSpec(ctx context.Context) (RoleRoleSpec_SdkV2, bool) {
	var e RoleRoleSpec_SdkV2
	if m.Spec.IsNull() || m.Spec.IsUnknown() {
		return e, false
	}
	var v []RoleRoleSpec_SdkV2
	d := m.Spec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSpec sets the value of the Spec field in Role_SdkV2.
func (m *Role_SdkV2) SetSpec(ctx context.Context, v RoleRoleSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spec"]
	m.Spec = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in Role_SdkV2 as
// a RoleRoleStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Role_SdkV2) GetStatus(ctx context.Context) (RoleRoleStatus_SdkV2, bool) {
	var e RoleRoleStatus_SdkV2
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v []RoleRoleStatus_SdkV2
	d := m.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in Role_SdkV2.
func (m *Role_SdkV2) SetStatus(ctx context.Context, v RoleRoleStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	m.Status = types.ListValueMust(t, vs)
}

type RoleOperationMetadata_SdkV2 struct {
}

func (to *RoleOperationMetadata_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RoleOperationMetadata_SdkV2) {
}

func (to *RoleOperationMetadata_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RoleOperationMetadata_SdkV2) {
}

func (m RoleOperationMetadata_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RoleOperationMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RoleOperationMetadata_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RoleOperationMetadata_SdkV2
// only implements ToObjectValue() and Type().
func (m RoleOperationMetadata_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m RoleOperationMetadata_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RoleRoleSpec_SdkV2 struct {
	// If auth_method is left unspecified, a meaningful authentication method is
	// derived from the identity_type: * For the managed identities, OAUTH is
	// used. * For the regular postgres roles, authentication based on postgres
	// passwords is used.
	//
	// NOTE: this is ignored for the Databricks identity type GROUP, and
	// NO_LOGIN is implicitly assumed instead for the GROUP identity type.
	AuthMethod types.String `tfsdk:"auth_method"`
	// The type of role. When specifying a managed-identity, the chosen role_id
	// must be a valid:
	//
	// * application ID for SERVICE_PRINCIPAL * user email for USER * group name
	// for GROUP
	IdentityType types.String `tfsdk:"identity_type"`
}

func (to *RoleRoleSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RoleRoleSpec_SdkV2) {
}

func (to *RoleRoleSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RoleRoleSpec_SdkV2) {
}

func (m RoleRoleSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auth_method"] = attrs["auth_method"].SetOptional()
	attrs["identity_type"] = attrs["identity_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RoleRoleSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RoleRoleSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RoleRoleSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m RoleRoleSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auth_method":   m.AuthMethod,
			"identity_type": m.IdentityType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RoleRoleSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auth_method":   types.StringType,
			"identity_type": types.StringType,
		},
	}
}

type RoleRoleStatus_SdkV2 struct {
	AuthMethod types.String `tfsdk:"auth_method"`
	// The type of the role.
	IdentityType types.String `tfsdk:"identity_type"`
}

func (to *RoleRoleStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RoleRoleStatus_SdkV2) {
}

func (to *RoleRoleStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RoleRoleStatus_SdkV2) {
}

func (m RoleRoleStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auth_method"] = attrs["auth_method"].SetOptional()
	attrs["identity_type"] = attrs["identity_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RoleRoleStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RoleRoleStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RoleRoleStatus_SdkV2
// only implements ToObjectValue() and Type().
func (m RoleRoleStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auth_method":   m.AuthMethod,
			"identity_type": m.IdentityType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RoleRoleStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auth_method":   types.StringType,
			"identity_type": types.StringType,
		},
	}
}

type UpdateBranchRequest_SdkV2 struct {
	// The Branch to update.
	//
	// The branch's `name` field is used to identify the branch to update.
	// Format: projects/{project_id}/branches/{branch_id}
	Branch types.List `tfsdk:"branch"`
	// The resource name of the branch. This field is output-only and
	// constructed by the system. Format:
	// `projects/{project_id}/branches/{branch_id}`
	Name types.String `tfsdk:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateBranchRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateBranchRequest_SdkV2) {
	if !from.Branch.IsNull() && !from.Branch.IsUnknown() {
		if toBranch, ok := to.GetBranch(ctx); ok {
			if fromBranch, ok := from.GetBranch(ctx); ok {
				// Recursively sync the fields of Branch
				toBranch.SyncFieldsDuringCreateOrUpdate(ctx, fromBranch)
				to.SetBranch(ctx, toBranch)
			}
		}
	}
}

func (to *UpdateBranchRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateBranchRequest_SdkV2) {
	if !from.Branch.IsNull() && !from.Branch.IsUnknown() {
		if toBranch, ok := to.GetBranch(ctx); ok {
			if fromBranch, ok := from.GetBranch(ctx); ok {
				toBranch.SyncFieldsDuringRead(ctx, fromBranch)
				to.SetBranch(ctx, toBranch)
			}
		}
	}
}

func (m UpdateBranchRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch"] = attrs["branch"].SetRequired()
	attrs["branch"] = attrs["branch"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateBranchRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateBranchRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"branch": reflect.TypeOf(Branch_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateBranchRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateBranchRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch":      m.Branch,
			"name":        m.Name,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateBranchRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch": basetypes.ListType{
				ElemType: Branch_SdkV2{}.Type(ctx),
			},
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetBranch returns the value of the Branch field in UpdateBranchRequest_SdkV2 as
// a Branch_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateBranchRequest_SdkV2) GetBranch(ctx context.Context) (Branch_SdkV2, bool) {
	var e Branch_SdkV2
	if m.Branch.IsNull() || m.Branch.IsUnknown() {
		return e, false
	}
	var v []Branch_SdkV2
	d := m.Branch.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBranch sets the value of the Branch field in UpdateBranchRequest_SdkV2.
func (m *UpdateBranchRequest_SdkV2) SetBranch(ctx context.Context, v Branch_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["branch"]
	m.Branch = types.ListValueMust(t, vs)
}

type UpdateEndpointRequest_SdkV2 struct {
	// The Endpoint to update.
	//
	// The endpoint's `name` field is used to identify the endpoint to update.
	// Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Endpoint types.List `tfsdk:"endpoint"`
	// The resource name of the endpoint. This field is output-only and
	// constructed by the system. Format:
	// `projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}`
	Name types.String `tfsdk:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateEndpointRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEndpointRequest_SdkV2) {
	if !from.Endpoint.IsNull() && !from.Endpoint.IsUnknown() {
		if toEndpoint, ok := to.GetEndpoint(ctx); ok {
			if fromEndpoint, ok := from.GetEndpoint(ctx); ok {
				// Recursively sync the fields of Endpoint
				toEndpoint.SyncFieldsDuringCreateOrUpdate(ctx, fromEndpoint)
				to.SetEndpoint(ctx, toEndpoint)
			}
		}
	}
}

func (to *UpdateEndpointRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateEndpointRequest_SdkV2) {
	if !from.Endpoint.IsNull() && !from.Endpoint.IsUnknown() {
		if toEndpoint, ok := to.GetEndpoint(ctx); ok {
			if fromEndpoint, ok := from.GetEndpoint(ctx); ok {
				toEndpoint.SyncFieldsDuringRead(ctx, fromEndpoint)
				to.SetEndpoint(ctx, toEndpoint)
			}
		}
	}
}

func (m UpdateEndpointRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint"] = attrs["endpoint"].SetRequired()
	attrs["endpoint"] = attrs["endpoint"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoint": reflect.TypeOf(Endpoint_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint":    m.Endpoint,
			"name":        m.Name,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint": basetypes.ListType{
				ElemType: Endpoint_SdkV2{}.Type(ctx),
			},
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetEndpoint returns the value of the Endpoint field in UpdateEndpointRequest_SdkV2 as
// a Endpoint_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateEndpointRequest_SdkV2) GetEndpoint(ctx context.Context) (Endpoint_SdkV2, bool) {
	var e Endpoint_SdkV2
	if m.Endpoint.IsNull() || m.Endpoint.IsUnknown() {
		return e, false
	}
	var v []Endpoint_SdkV2
	d := m.Endpoint.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEndpoint sets the value of the Endpoint field in UpdateEndpointRequest_SdkV2.
func (m *UpdateEndpointRequest_SdkV2) SetEndpoint(ctx context.Context, v Endpoint_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["endpoint"]
	m.Endpoint = types.ListValueMust(t, vs)
}

type UpdateProjectRequest_SdkV2 struct {
	// The resource name of the project. This field is output-only and
	// constructed by the system. Format: `projects/{project_id}`
	Name types.String `tfsdk:"-"`
	// The Project to update.
	//
	// The project's `name` field is used to identify the project to update.
	// Format: projects/{project_id}
	Project types.List `tfsdk:"project"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateProjectRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateProjectRequest_SdkV2) {
	if !from.Project.IsNull() && !from.Project.IsUnknown() {
		if toProject, ok := to.GetProject(ctx); ok {
			if fromProject, ok := from.GetProject(ctx); ok {
				// Recursively sync the fields of Project
				toProject.SyncFieldsDuringCreateOrUpdate(ctx, fromProject)
				to.SetProject(ctx, toProject)
			}
		}
	}
}

func (to *UpdateProjectRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateProjectRequest_SdkV2) {
	if !from.Project.IsNull() && !from.Project.IsUnknown() {
		if toProject, ok := to.GetProject(ctx); ok {
			if fromProject, ok := from.GetProject(ctx); ok {
				toProject.SyncFieldsDuringRead(ctx, fromProject)
				to.SetProject(ctx, toProject)
			}
		}
	}
}

func (m UpdateProjectRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project"] = attrs["project"].SetRequired()
	attrs["project"] = attrs["project"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateProjectRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateProjectRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"project": reflect.TypeOf(Project_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProjectRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateProjectRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":        m.Name,
			"project":     m.Project,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateProjectRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"project": basetypes.ListType{
				ElemType: Project_SdkV2{}.Type(ctx),
			},
			"update_mask": types.StringType,
		},
	}
}

// GetProject returns the value of the Project field in UpdateProjectRequest_SdkV2 as
// a Project_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateProjectRequest_SdkV2) GetProject(ctx context.Context) (Project_SdkV2, bool) {
	var e Project_SdkV2
	if m.Project.IsNull() || m.Project.IsUnknown() {
		return e, false
	}
	var v []Project_SdkV2
	d := m.Project.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetProject sets the value of the Project field in UpdateProjectRequest_SdkV2.
func (m *UpdateProjectRequest_SdkV2) SetProject(ctx context.Context, v Project_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["project"]
	m.Project = types.ListValueMust(t, vs)
}
