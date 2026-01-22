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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type Branch struct {
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
	Spec types.Object `tfsdk:"spec"`
	// The current status of a Branch.
	Status types.Object `tfsdk:"status"`
	// System-generated unique ID for the branch.
	Uid types.String `tfsdk:"uid"`
	// A timestamp indicating when the branch was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *Branch) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Branch) {
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

func (to *Branch) SyncFieldsDuringRead(ctx context.Context, from Branch) {
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

func (m Branch) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["parent"] = attrs["parent"].SetComputed()
	attrs["spec"] = attrs["spec"].SetOptional()
	attrs["spec"] = attrs["spec"].SetComputed()
	attrs["spec"] = attrs["spec"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetComputed()
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
func (m Branch) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spec":   reflect.TypeOf(BranchSpec{}),
		"status": reflect.TypeOf(BranchStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Branch
// only implements ToObjectValue() and Type().
func (m Branch) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Branch) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": timetypes.RFC3339{}.Type(ctx),
			"name":        types.StringType,
			"parent":      types.StringType,
			"spec":        BranchSpec{}.Type(ctx),
			"status":      BranchStatus{}.Type(ctx),
			"uid":         types.StringType,
			"update_time": timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetSpec returns the value of the Spec field in Branch as
// a BranchSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *Branch) GetSpec(ctx context.Context) (BranchSpec, bool) {
	var e BranchSpec
	if m.Spec.IsNull() || m.Spec.IsUnknown() {
		return e, false
	}
	var v BranchSpec
	d := m.Spec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSpec sets the value of the Spec field in Branch.
func (m *Branch) SetSpec(ctx context.Context, v BranchSpec) {
	vs := v.ToObjectValue(ctx)
	m.Spec = vs
}

// GetStatus returns the value of the Status field in Branch as
// a BranchStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *Branch) GetStatus(ctx context.Context) (BranchStatus, bool) {
	var e BranchStatus
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v BranchStatus
	d := m.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in Branch.
func (m *Branch) SetStatus(ctx context.Context, v BranchStatus) {
	vs := v.ToObjectValue(ctx)
	m.Status = vs
}

type BranchOperationMetadata struct {
}

func (to *BranchOperationMetadata) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BranchOperationMetadata) {
}

func (to *BranchOperationMetadata) SyncFieldsDuringRead(ctx context.Context, from BranchOperationMetadata) {
}

func (m BranchOperationMetadata) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BranchOperationMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m BranchOperationMetadata) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BranchOperationMetadata
// only implements ToObjectValue() and Type().
func (m BranchOperationMetadata) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m BranchOperationMetadata) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type BranchSpec struct {
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

func (to *BranchSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BranchSpec) {
}

func (to *BranchSpec) SyncFieldsDuringRead(ctx context.Context, from BranchSpec) {
}

func (m BranchSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m BranchSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BranchSpec
// only implements ToObjectValue() and Type().
func (m BranchSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m BranchSpec) Type(ctx context.Context) attr.Type {
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

type BranchStatus struct {
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

func (to *BranchStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BranchStatus) {
}

func (to *BranchStatus) SyncFieldsDuringRead(ctx context.Context, from BranchStatus) {
}

func (m BranchStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m BranchStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BranchStatus
// only implements ToObjectValue() and Type().
func (m BranchStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m BranchStatus) Type(ctx context.Context) attr.Type {
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

type CreateBranchRequest struct {
	// The Branch to create.
	Branch types.Object `tfsdk:"branch"`
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

func (to *CreateBranchRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateBranchRequest) {
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

func (to *CreateBranchRequest) SyncFieldsDuringRead(ctx context.Context, from CreateBranchRequest) {
	if !from.Branch.IsNull() && !from.Branch.IsUnknown() {
		if toBranch, ok := to.GetBranch(ctx); ok {
			if fromBranch, ok := from.GetBranch(ctx); ok {
				toBranch.SyncFieldsDuringRead(ctx, fromBranch)
				to.SetBranch(ctx, toBranch)
			}
		}
	}
}

func (m CreateBranchRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch"] = attrs["branch"].SetRequired()
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
func (m CreateBranchRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"branch": reflect.TypeOf(Branch{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateBranchRequest
// only implements ToObjectValue() and Type().
func (m CreateBranchRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch":    m.Branch,
			"branch_id": m.BranchId,
			"parent":    m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateBranchRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch":    Branch{}.Type(ctx),
			"branch_id": types.StringType,
			"parent":    types.StringType,
		},
	}
}

// GetBranch returns the value of the Branch field in CreateBranchRequest as
// a Branch value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateBranchRequest) GetBranch(ctx context.Context) (Branch, bool) {
	var e Branch
	if m.Branch.IsNull() || m.Branch.IsUnknown() {
		return e, false
	}
	var v Branch
	d := m.Branch.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBranch sets the value of the Branch field in CreateBranchRequest.
func (m *CreateBranchRequest) SetBranch(ctx context.Context, v Branch) {
	vs := v.ToObjectValue(ctx)
	m.Branch = vs
}

type CreateEndpointRequest struct {
	// The Endpoint to create.
	Endpoint types.Object `tfsdk:"endpoint"`
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

func (to *CreateEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateEndpointRequest) {
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

func (to *CreateEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from CreateEndpointRequest) {
	if !from.Endpoint.IsNull() && !from.Endpoint.IsUnknown() {
		if toEndpoint, ok := to.GetEndpoint(ctx); ok {
			if fromEndpoint, ok := from.GetEndpoint(ctx); ok {
				toEndpoint.SyncFieldsDuringRead(ctx, fromEndpoint)
				to.SetEndpoint(ctx, toEndpoint)
			}
		}
	}
}

func (m CreateEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint"] = attrs["endpoint"].SetRequired()
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
func (m CreateEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoint": reflect.TypeOf(Endpoint{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateEndpointRequest
// only implements ToObjectValue() and Type().
func (m CreateEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint":    m.Endpoint,
			"endpoint_id": m.EndpointId,
			"parent":      m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint":    Endpoint{}.Type(ctx),
			"endpoint_id": types.StringType,
			"parent":      types.StringType,
		},
	}
}

// GetEndpoint returns the value of the Endpoint field in CreateEndpointRequest as
// a Endpoint value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateEndpointRequest) GetEndpoint(ctx context.Context) (Endpoint, bool) {
	var e Endpoint
	if m.Endpoint.IsNull() || m.Endpoint.IsUnknown() {
		return e, false
	}
	var v Endpoint
	d := m.Endpoint.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEndpoint sets the value of the Endpoint field in CreateEndpointRequest.
func (m *CreateEndpointRequest) SetEndpoint(ctx context.Context, v Endpoint) {
	vs := v.ToObjectValue(ctx)
	m.Endpoint = vs
}

type CreateProjectRequest struct {
	// The Project to create.
	Project types.Object `tfsdk:"project"`
	// The ID to use for the Project. This becomes the final component of the
	// project's resource name. The ID must be 1-63 characters long, start with
	// a lowercase letter, and contain only lowercase letters, numbers, and
	// hyphens (RFC 1123). Examples: - With custom ID: `production` → name
	// becomes `projects/production` - Without custom ID: system generates UUID
	// → name becomes `projects/a7f89b2c-3d4e-5f6g-7h8i-9j0k1l2m3n4o`
	ProjectId types.String `tfsdk:"-"`
}

func (to *CreateProjectRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateProjectRequest) {
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

func (to *CreateProjectRequest) SyncFieldsDuringRead(ctx context.Context, from CreateProjectRequest) {
	if !from.Project.IsNull() && !from.Project.IsUnknown() {
		if toProject, ok := to.GetProject(ctx); ok {
			if fromProject, ok := from.GetProject(ctx); ok {
				toProject.SyncFieldsDuringRead(ctx, fromProject)
				to.SetProject(ctx, toProject)
			}
		}
	}
}

func (m CreateProjectRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project"] = attrs["project"].SetRequired()
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
func (m CreateProjectRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"project": reflect.TypeOf(Project{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateProjectRequest
// only implements ToObjectValue() and Type().
func (m CreateProjectRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"project":    m.Project,
			"project_id": m.ProjectId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateProjectRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"project":    Project{}.Type(ctx),
			"project_id": types.StringType,
		},
	}
}

// GetProject returns the value of the Project field in CreateProjectRequest as
// a Project value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateProjectRequest) GetProject(ctx context.Context) (Project, bool) {
	var e Project
	if m.Project.IsNull() || m.Project.IsUnknown() {
		return e, false
	}
	var v Project
	d := m.Project.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProject sets the value of the Project field in CreateProjectRequest.
func (m *CreateProjectRequest) SetProject(ctx context.Context, v Project) {
	vs := v.ToObjectValue(ctx)
	m.Project = vs
}

type CreateRoleRequest struct {
	// The Branch where this Role is created. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent types.String `tfsdk:"-"`
	// The desired specification of a Role.
	Role types.Object `tfsdk:"role"`
	// The ID to use for the Role, which will become the final component of the
	// role's resource name. This ID becomes the role in Postgres.
	//
	// This value should be 4-63 characters, and valid characters are lowercase
	// letters, numbers, and hyphens, as defined by RFC 1123.
	RoleId types.String `tfsdk:"-"`
}

func (to *CreateRoleRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateRoleRequest) {
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

func (to *CreateRoleRequest) SyncFieldsDuringRead(ctx context.Context, from CreateRoleRequest) {
	if !from.Role.IsNull() && !from.Role.IsUnknown() {
		if toRole, ok := to.GetRole(ctx); ok {
			if fromRole, ok := from.GetRole(ctx); ok {
				toRole.SyncFieldsDuringRead(ctx, fromRole)
				to.SetRole(ctx, toRole)
			}
		}
	}
}

func (m CreateRoleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["role"] = attrs["role"].SetRequired()
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
func (m CreateRoleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"role": reflect.TypeOf(Role{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRoleRequest
// only implements ToObjectValue() and Type().
func (m CreateRoleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parent":  m.Parent,
			"role":    m.Role,
			"role_id": m.RoleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateRoleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent":  types.StringType,
			"role":    Role{}.Type(ctx),
			"role_id": types.StringType,
		},
	}
}

// GetRole returns the value of the Role field in CreateRoleRequest as
// a Role value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateRoleRequest) GetRole(ctx context.Context) (Role, bool) {
	var e Role
	if m.Role.IsNull() || m.Role.IsUnknown() {
		return e, false
	}
	var v Role
	d := m.Role.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRole sets the value of the Role field in CreateRoleRequest.
func (m *CreateRoleRequest) SetRole(ctx context.Context, v Role) {
	vs := v.ToObjectValue(ctx)
	m.Role = vs
}

type DatabaseCredential struct {
	// Timestamp in UTC of when this credential expires.
	ExpireTime timetypes.RFC3339 `tfsdk:"expire_time"`
	// The OAuth token that can be used as a password when connecting to a
	// database.
	Token types.String `tfsdk:"token"`
}

func (to *DatabaseCredential) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseCredential) {
}

func (to *DatabaseCredential) SyncFieldsDuringRead(ctx context.Context, from DatabaseCredential) {
}

func (m DatabaseCredential) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DatabaseCredential) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseCredential
// only implements ToObjectValue() and Type().
func (m DatabaseCredential) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"expire_time": m.ExpireTime,
			"token":       m.Token,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseCredential) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"expire_time": timetypes.RFC3339{}.Type(ctx),
			"token":       types.StringType,
		},
	}
}

// Databricks Error that is returned by all Databricks APIs.
type DatabricksServiceExceptionWithDetailsProto struct {
	Details types.List `tfsdk:"details"`

	ErrorCode types.String `tfsdk:"error_code"`

	Message types.String `tfsdk:"message"`

	StackTrace types.String `tfsdk:"stack_trace"`
}

func (to *DatabricksServiceExceptionWithDetailsProto) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabricksServiceExceptionWithDetailsProto) {
	if !from.Details.IsNull() && !from.Details.IsUnknown() && to.Details.IsNull() && len(from.Details.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Details, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Details = from.Details
	}
}

func (to *DatabricksServiceExceptionWithDetailsProto) SyncFieldsDuringRead(ctx context.Context, from DatabricksServiceExceptionWithDetailsProto) {
	if !from.Details.IsNull() && !from.Details.IsUnknown() && to.Details.IsNull() && len(from.Details.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Details, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Details = from.Details
	}
}

func (m DatabricksServiceExceptionWithDetailsProto) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DatabricksServiceExceptionWithDetailsProto) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"details": reflect.TypeOf(types.Object{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabricksServiceExceptionWithDetailsProto
// only implements ToObjectValue() and Type().
func (m DatabricksServiceExceptionWithDetailsProto) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m DatabricksServiceExceptionWithDetailsProto) Type(ctx context.Context) attr.Type {
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

// GetDetails returns the value of the Details field in DatabricksServiceExceptionWithDetailsProto as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabricksServiceExceptionWithDetailsProto) GetDetails(ctx context.Context) ([]types.Object, bool) {
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

// SetDetails sets the value of the Details field in DatabricksServiceExceptionWithDetailsProto.
func (m *DatabricksServiceExceptionWithDetailsProto) SetDetails(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["details"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Details = types.ListValueMust(t, vs)
}

type DeleteBranchRequest struct {
	// The name of the Branch to delete. Format:
	// projects/{project_id}/branches/{branch_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteBranchRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteBranchRequest) {
}

func (to *DeleteBranchRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteBranchRequest) {
}

func (m DeleteBranchRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteBranchRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteBranchRequest
// only implements ToObjectValue() and Type().
func (m DeleteBranchRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteBranchRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteEndpointRequest struct {
	// The name of the Endpoint to delete. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteEndpointRequest) {
}

func (to *DeleteEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteEndpointRequest) {
}

func (m DeleteEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteEndpointRequest
// only implements ToObjectValue() and Type().
func (m DeleteEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteProjectRequest struct {
	// The name of the Project to delete. Format: projects/{project_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteProjectRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteProjectRequest) {
}

func (to *DeleteProjectRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteProjectRequest) {
}

func (m DeleteProjectRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteProjectRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteProjectRequest
// only implements ToObjectValue() and Type().
func (m DeleteProjectRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteProjectRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteRoleRequest struct {
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

func (to *DeleteRoleRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRoleRequest) {
}

func (to *DeleteRoleRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteRoleRequest) {
}

func (m DeleteRoleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteRoleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRoleRequest
// only implements ToObjectValue() and Type().
func (m DeleteRoleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":              m.Name,
			"reassign_owned_to": m.ReassignOwnedTo,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRoleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":              types.StringType,
			"reassign_owned_to": types.StringType,
		},
	}
}

type Endpoint struct {
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
	Spec types.Object `tfsdk:"spec"`
	// Current operational status of the compute endpoint.
	Status types.Object `tfsdk:"status"`
	// System-generated unique ID for the endpoint.
	Uid types.String `tfsdk:"uid"`
	// A timestamp indicating when the compute endpoint was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *Endpoint) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Endpoint) {
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

func (to *Endpoint) SyncFieldsDuringRead(ctx context.Context, from Endpoint) {
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

func (m Endpoint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["parent"] = attrs["parent"].SetComputed()
	attrs["spec"] = attrs["spec"].SetOptional()
	attrs["spec"] = attrs["spec"].SetComputed()
	attrs["spec"] = attrs["spec"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetComputed()
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
func (m Endpoint) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spec":   reflect.TypeOf(EndpointSpec{}),
		"status": reflect.TypeOf(EndpointStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Endpoint
// only implements ToObjectValue() and Type().
func (m Endpoint) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Endpoint) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": timetypes.RFC3339{}.Type(ctx),
			"name":        types.StringType,
			"parent":      types.StringType,
			"spec":        EndpointSpec{}.Type(ctx),
			"status":      EndpointStatus{}.Type(ctx),
			"uid":         types.StringType,
			"update_time": timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetSpec returns the value of the Spec field in Endpoint as
// a EndpointSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint) GetSpec(ctx context.Context) (EndpointSpec, bool) {
	var e EndpointSpec
	if m.Spec.IsNull() || m.Spec.IsUnknown() {
		return e, false
	}
	var v EndpointSpec
	d := m.Spec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSpec sets the value of the Spec field in Endpoint.
func (m *Endpoint) SetSpec(ctx context.Context, v EndpointSpec) {
	vs := v.ToObjectValue(ctx)
	m.Spec = vs
}

// GetStatus returns the value of the Status field in Endpoint as
// a EndpointStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint) GetStatus(ctx context.Context) (EndpointStatus, bool) {
	var e EndpointStatus
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v EndpointStatus
	d := m.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in Endpoint.
func (m *Endpoint) SetStatus(ctx context.Context, v EndpointStatus) {
	vs := v.ToObjectValue(ctx)
	m.Status = vs
}

// Encapsulates various hostnames (r/w or r/o, pooled or not) for an endpoint.
type EndpointHosts struct {
	// The hostname to connect to this endpoint. For read-write endpoints, this
	// is a read-write hostname which connects to the primary compute. For
	// read-only endpoints, this is a read-only hostname which allows read-only
	// operations.
	Host types.String `tfsdk:"host"`
}

func (to *EndpointHosts) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointHosts) {
}

func (to *EndpointHosts) SyncFieldsDuringRead(ctx context.Context, from EndpointHosts) {
}

func (m EndpointHosts) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EndpointHosts) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointHosts
// only implements ToObjectValue() and Type().
func (m EndpointHosts) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"host": m.Host,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointHosts) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"host": types.StringType,
		},
	}
}

type EndpointOperationMetadata struct {
}

func (to *EndpointOperationMetadata) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointOperationMetadata) {
}

func (to *EndpointOperationMetadata) SyncFieldsDuringRead(ctx context.Context, from EndpointOperationMetadata) {
}

func (m EndpointOperationMetadata) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointOperationMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointOperationMetadata) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointOperationMetadata
// only implements ToObjectValue() and Type().
func (m EndpointOperationMetadata) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointOperationMetadata) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// A collection of settings for a compute endpoint.
type EndpointSettings struct {
	// A raw representation of Postgres settings.
	PgSettings types.Map `tfsdk:"pg_settings"`
}

func (to *EndpointSettings) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointSettings) {
}

func (to *EndpointSettings) SyncFieldsDuringRead(ctx context.Context, from EndpointSettings) {
}

func (m EndpointSettings) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EndpointSettings) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"pg_settings": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointSettings
// only implements ToObjectValue() and Type().
func (m EndpointSettings) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pg_settings": m.PgSettings,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointSettings) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pg_settings": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetPgSettings returns the value of the PgSettings field in EndpointSettings as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointSettings) GetPgSettings(ctx context.Context) (map[string]types.String, bool) {
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

// SetPgSettings sets the value of the PgSettings field in EndpointSettings.
func (m *EndpointSettings) SetPgSettings(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["pg_settings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PgSettings = types.MapValueMust(t, vs)
}

type EndpointSpec struct {
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

	Settings types.Object `tfsdk:"settings"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended. If specified should be between 60s and 604800s (1 minute to 1
	// week).
	SuspendTimeoutDuration timetypes.GoDuration `tfsdk:"suspend_timeout_duration"`
}

func (to *EndpointSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointSpec) {
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

func (to *EndpointSpec) SyncFieldsDuringRead(ctx context.Context, from EndpointSpec) {
	if !from.Settings.IsNull() && !from.Settings.IsUnknown() {
		if toSettings, ok := to.GetSettings(ctx); ok {
			if fromSettings, ok := from.GetSettings(ctx); ok {
				toSettings.SyncFieldsDuringRead(ctx, fromSettings)
				to.SetSettings(ctx, toSettings)
			}
		}
	}
}

func (m EndpointSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autoscaling_limit_max_cu"] = attrs["autoscaling_limit_max_cu"].SetOptional()
	attrs["autoscaling_limit_min_cu"] = attrs["autoscaling_limit_min_cu"].SetOptional()
	attrs["disabled"] = attrs["disabled"].SetOptional()
	attrs["endpoint_type"] = attrs["endpoint_type"].SetRequired()
	attrs["endpoint_type"] = attrs["endpoint_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["no_suspension"] = attrs["no_suspension"].SetOptional()
	attrs["settings"] = attrs["settings"].SetOptional()
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
func (m EndpointSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"settings": reflect.TypeOf(EndpointSettings{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointSpec
// only implements ToObjectValue() and Type().
func (m EndpointSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m EndpointSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscaling_limit_max_cu": types.Float64Type,
			"autoscaling_limit_min_cu": types.Float64Type,
			"disabled":                 types.BoolType,
			"endpoint_type":            types.StringType,
			"no_suspension":            types.BoolType,
			"settings":                 EndpointSettings{}.Type(ctx),
			"suspend_timeout_duration": timetypes.GoDuration{}.Type(ctx),
		},
	}
}

// GetSettings returns the value of the Settings field in EndpointSpec as
// a EndpointSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointSpec) GetSettings(ctx context.Context) (EndpointSettings, bool) {
	var e EndpointSettings
	if m.Settings.IsNull() || m.Settings.IsUnknown() {
		return e, false
	}
	var v EndpointSettings
	d := m.Settings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSettings sets the value of the Settings field in EndpointSpec.
func (m *EndpointSpec) SetSettings(ctx context.Context, v EndpointSettings) {
	vs := v.ToObjectValue(ctx)
	m.Settings = vs
}

type EndpointStatus struct {
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
	Hosts types.Object `tfsdk:"hosts"`

	PendingState types.String `tfsdk:"pending_state"`

	Settings types.Object `tfsdk:"settings"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended.
	SuspendTimeoutDuration timetypes.GoDuration `tfsdk:"suspend_timeout_duration"`
}

func (to *EndpointStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointStatus) {
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

func (to *EndpointStatus) SyncFieldsDuringRead(ctx context.Context, from EndpointStatus) {
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

func (m EndpointStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autoscaling_limit_max_cu"] = attrs["autoscaling_limit_max_cu"].SetComputed()
	attrs["autoscaling_limit_min_cu"] = attrs["autoscaling_limit_min_cu"].SetComputed()
	attrs["current_state"] = attrs["current_state"].SetComputed()
	attrs["disabled"] = attrs["disabled"].SetComputed()
	attrs["endpoint_type"] = attrs["endpoint_type"].SetComputed()
	attrs["hosts"] = attrs["hosts"].SetComputed()
	attrs["pending_state"] = attrs["pending_state"].SetComputed()
	attrs["settings"] = attrs["settings"].SetComputed()
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
func (m EndpointStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"hosts":    reflect.TypeOf(EndpointHosts{}),
		"settings": reflect.TypeOf(EndpointSettings{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointStatus
// only implements ToObjectValue() and Type().
func (m EndpointStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m EndpointStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscaling_limit_max_cu": types.Float64Type,
			"autoscaling_limit_min_cu": types.Float64Type,
			"current_state":            types.StringType,
			"disabled":                 types.BoolType,
			"endpoint_type":            types.StringType,
			"hosts":                    EndpointHosts{}.Type(ctx),
			"pending_state":            types.StringType,
			"settings":                 EndpointSettings{}.Type(ctx),
			"suspend_timeout_duration": timetypes.GoDuration{}.Type(ctx),
		},
	}
}

// GetHosts returns the value of the Hosts field in EndpointStatus as
// a EndpointHosts value.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointStatus) GetHosts(ctx context.Context) (EndpointHosts, bool) {
	var e EndpointHosts
	if m.Hosts.IsNull() || m.Hosts.IsUnknown() {
		return e, false
	}
	var v EndpointHosts
	d := m.Hosts.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetHosts sets the value of the Hosts field in EndpointStatus.
func (m *EndpointStatus) SetHosts(ctx context.Context, v EndpointHosts) {
	vs := v.ToObjectValue(ctx)
	m.Hosts = vs
}

// GetSettings returns the value of the Settings field in EndpointStatus as
// a EndpointSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointStatus) GetSettings(ctx context.Context) (EndpointSettings, bool) {
	var e EndpointSettings
	if m.Settings.IsNull() || m.Settings.IsUnknown() {
		return e, false
	}
	var v EndpointSettings
	d := m.Settings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSettings sets the value of the Settings field in EndpointStatus.
func (m *EndpointStatus) SetSettings(ctx context.Context, v EndpointSettings) {
	vs := v.ToObjectValue(ctx)
	m.Settings = vs
}

type GenerateDatabaseCredentialRequest struct {
	// The returned token will be scoped to UC tables with the specified
	// permissions.
	Claims types.List `tfsdk:"claims"`
	// This field is not yet supported. The endpoint for which this credential
	// will be generated. Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Endpoint types.String `tfsdk:"endpoint"`
}

func (to *GenerateDatabaseCredentialRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenerateDatabaseCredentialRequest) {
	if !from.Claims.IsNull() && !from.Claims.IsUnknown() && to.Claims.IsNull() && len(from.Claims.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Claims, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Claims = from.Claims
	}
}

func (to *GenerateDatabaseCredentialRequest) SyncFieldsDuringRead(ctx context.Context, from GenerateDatabaseCredentialRequest) {
	if !from.Claims.IsNull() && !from.Claims.IsUnknown() && to.Claims.IsNull() && len(from.Claims.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Claims, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Claims = from.Claims
	}
}

func (m GenerateDatabaseCredentialRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GenerateDatabaseCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"claims": reflect.TypeOf(RequestedClaims{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenerateDatabaseCredentialRequest
// only implements ToObjectValue() and Type().
func (m GenerateDatabaseCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"claims":   m.Claims,
			"endpoint": m.Endpoint,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenerateDatabaseCredentialRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"claims": basetypes.ListType{
				ElemType: RequestedClaims{}.Type(ctx),
			},
			"endpoint": types.StringType,
		},
	}
}

// GetClaims returns the value of the Claims field in GenerateDatabaseCredentialRequest as
// a slice of RequestedClaims values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenerateDatabaseCredentialRequest) GetClaims(ctx context.Context) ([]RequestedClaims, bool) {
	if m.Claims.IsNull() || m.Claims.IsUnknown() {
		return nil, false
	}
	var v []RequestedClaims
	d := m.Claims.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClaims sets the value of the Claims field in GenerateDatabaseCredentialRequest.
func (m *GenerateDatabaseCredentialRequest) SetClaims(ctx context.Context, v []RequestedClaims) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["claims"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Claims = types.ListValueMust(t, vs)
}

type GetBranchRequest struct {
	// The resource name of the branch to retrieve. Format:
	// `projects/{project_id}/branches/{branch_id}`
	Name types.String `tfsdk:"-"`
}

func (to *GetBranchRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetBranchRequest) {
}

func (to *GetBranchRequest) SyncFieldsDuringRead(ctx context.Context, from GetBranchRequest) {
}

func (m GetBranchRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetBranchRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetBranchRequest
// only implements ToObjectValue() and Type().
func (m GetBranchRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetBranchRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetEndpointRequest struct {
	// The resource name of the endpoint to retrieve. Format:
	// `projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}`
	Name types.String `tfsdk:"-"`
}

func (to *GetEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEndpointRequest) {
}

func (to *GetEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from GetEndpointRequest) {
}

func (m GetEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEndpointRequest
// only implements ToObjectValue() and Type().
func (m GetEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetOperationRequest struct {
	// The name of the operation resource.
	Name types.String `tfsdk:"-"`
}

func (to *GetOperationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetOperationRequest) {
}

func (to *GetOperationRequest) SyncFieldsDuringRead(ctx context.Context, from GetOperationRequest) {
}

func (m GetOperationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetOperationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetOperationRequest
// only implements ToObjectValue() and Type().
func (m GetOperationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetOperationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetProjectRequest struct {
	// The resource name of the project to retrieve. Format:
	// `projects/{project_id}`
	Name types.String `tfsdk:"-"`
}

func (to *GetProjectRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetProjectRequest) {
}

func (to *GetProjectRequest) SyncFieldsDuringRead(ctx context.Context, from GetProjectRequest) {
}

func (m GetProjectRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetProjectRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetProjectRequest
// only implements ToObjectValue() and Type().
func (m GetProjectRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetProjectRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetRoleRequest struct {
	// The name of the Role to retrieve. Format:
	// projects/{project_id}/branches/{branch_id}/roles/{role_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetRoleRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRoleRequest) {
}

func (to *GetRoleRequest) SyncFieldsDuringRead(ctx context.Context, from GetRoleRequest) {
}

func (m GetRoleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetRoleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRoleRequest
// only implements ToObjectValue() and Type().
func (m GetRoleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRoleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ListBranchesRequest struct {
	// Upper bound for items returned. Cannot be negative.
	PageSize types.Int64 `tfsdk:"-"`
	// Page token from a previous response. If not provided, returns the first
	// page.
	PageToken types.String `tfsdk:"-"`
	// The Project that owns this collection of branches. Format:
	// projects/{project_id}
	Parent types.String `tfsdk:"-"`
}

func (to *ListBranchesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListBranchesRequest) {
}

func (to *ListBranchesRequest) SyncFieldsDuringRead(ctx context.Context, from ListBranchesRequest) {
}

func (m ListBranchesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListBranchesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListBranchesRequest
// only implements ToObjectValue() and Type().
func (m ListBranchesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListBranchesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListBranchesResponse struct {
	// List of database branches in the project.
	Branches types.List `tfsdk:"branches"`
	// Token to request the next page of database branches.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListBranchesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListBranchesResponse) {
	if !from.Branches.IsNull() && !from.Branches.IsUnknown() && to.Branches.IsNull() && len(from.Branches.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Branches, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Branches = from.Branches
	}
}

func (to *ListBranchesResponse) SyncFieldsDuringRead(ctx context.Context, from ListBranchesResponse) {
	if !from.Branches.IsNull() && !from.Branches.IsUnknown() && to.Branches.IsNull() && len(from.Branches.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Branches, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Branches = from.Branches
	}
}

func (m ListBranchesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListBranchesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"branches": reflect.TypeOf(Branch{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListBranchesResponse
// only implements ToObjectValue() and Type().
func (m ListBranchesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branches":        m.Branches,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListBranchesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branches": basetypes.ListType{
				ElemType: Branch{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetBranches returns the value of the Branches field in ListBranchesResponse as
// a slice of Branch values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListBranchesResponse) GetBranches(ctx context.Context) ([]Branch, bool) {
	if m.Branches.IsNull() || m.Branches.IsUnknown() {
		return nil, false
	}
	var v []Branch
	d := m.Branches.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBranches sets the value of the Branches field in ListBranchesResponse.
func (m *ListBranchesResponse) SetBranches(ctx context.Context, v []Branch) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["branches"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Branches = types.ListValueMust(t, vs)
}

type ListEndpointsRequest struct {
	// Upper bound for items returned. Cannot be negative.
	PageSize types.Int64 `tfsdk:"-"`
	// Page token from a previous response. If not provided, returns the first
	// page.
	PageToken types.String `tfsdk:"-"`
	// The Branch that owns this collection of endpoints. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent types.String `tfsdk:"-"`
}

func (to *ListEndpointsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEndpointsRequest) {
}

func (to *ListEndpointsRequest) SyncFieldsDuringRead(ctx context.Context, from ListEndpointsRequest) {
}

func (m ListEndpointsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListEndpointsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointsRequest
// only implements ToObjectValue() and Type().
func (m ListEndpointsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListEndpointsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListEndpointsResponse struct {
	// List of compute endpoints in the branch.
	Endpoints types.List `tfsdk:"endpoints"`
	// Token to request the next page of compute endpoints.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListEndpointsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEndpointsResponse) {
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() && to.Endpoints.IsNull() && len(from.Endpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Endpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Endpoints = from.Endpoints
	}
}

func (to *ListEndpointsResponse) SyncFieldsDuringRead(ctx context.Context, from ListEndpointsResponse) {
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() && to.Endpoints.IsNull() && len(from.Endpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Endpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Endpoints = from.Endpoints
	}
}

func (m ListEndpointsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListEndpointsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoints": reflect.TypeOf(Endpoint{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointsResponse
// only implements ToObjectValue() and Type().
func (m ListEndpointsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoints":       m.Endpoints,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListEndpointsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoints": basetypes.ListType{
				ElemType: Endpoint{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetEndpoints returns the value of the Endpoints field in ListEndpointsResponse as
// a slice of Endpoint values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListEndpointsResponse) GetEndpoints(ctx context.Context) ([]Endpoint, bool) {
	if m.Endpoints.IsNull() || m.Endpoints.IsUnknown() {
		return nil, false
	}
	var v []Endpoint
	d := m.Endpoints.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEndpoints sets the value of the Endpoints field in ListEndpointsResponse.
func (m *ListEndpointsResponse) SetEndpoints(ctx context.Context, v []Endpoint) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["endpoints"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Endpoints = types.ListValueMust(t, vs)
}

type ListProjectsRequest struct {
	// Upper bound for items returned. Cannot be negative.
	PageSize types.Int64 `tfsdk:"-"`
	// Page token from a previous response. If not provided, returns the first
	// page.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListProjectsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProjectsRequest) {
}

func (to *ListProjectsRequest) SyncFieldsDuringRead(ctx context.Context, from ListProjectsRequest) {
}

func (m ListProjectsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListProjectsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProjectsRequest
// only implements ToObjectValue() and Type().
func (m ListProjectsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListProjectsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListProjectsResponse struct {
	// Token to request the next page of database projects.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of all database projects in the workspace that the user has
	// permission to access.
	Projects types.List `tfsdk:"projects"`
}

func (to *ListProjectsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListProjectsResponse) {
	if !from.Projects.IsNull() && !from.Projects.IsUnknown() && to.Projects.IsNull() && len(from.Projects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Projects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Projects = from.Projects
	}
}

func (to *ListProjectsResponse) SyncFieldsDuringRead(ctx context.Context, from ListProjectsResponse) {
	if !from.Projects.IsNull() && !from.Projects.IsUnknown() && to.Projects.IsNull() && len(from.Projects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Projects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Projects = from.Projects
	}
}

func (m ListProjectsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListProjectsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"projects": reflect.TypeOf(Project{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListProjectsResponse
// only implements ToObjectValue() and Type().
func (m ListProjectsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"projects":        m.Projects,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListProjectsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"projects": basetypes.ListType{
				ElemType: Project{}.Type(ctx),
			},
		},
	}
}

// GetProjects returns the value of the Projects field in ListProjectsResponse as
// a slice of Project values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListProjectsResponse) GetProjects(ctx context.Context) ([]Project, bool) {
	if m.Projects.IsNull() || m.Projects.IsUnknown() {
		return nil, false
	}
	var v []Project
	d := m.Projects.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProjects sets the value of the Projects field in ListProjectsResponse.
func (m *ListProjectsResponse) SetProjects(ctx context.Context, v []Project) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["projects"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Projects = types.ListValueMust(t, vs)
}

type ListRolesRequest struct {
	// Upper bound for items returned. Cannot be negative.
	PageSize types.Int64 `tfsdk:"-"`
	// Page token from a previous response. If not provided, returns the first
	// page.
	PageToken types.String `tfsdk:"-"`
	// The Branch that owns this collection of roles. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent types.String `tfsdk:"-"`
}

func (to *ListRolesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListRolesRequest) {
}

func (to *ListRolesRequest) SyncFieldsDuringRead(ctx context.Context, from ListRolesRequest) {
}

func (m ListRolesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListRolesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRolesRequest
// only implements ToObjectValue() and Type().
func (m ListRolesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListRolesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListRolesResponse struct {
	// Token to request the next page of Postgres roles.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of Postgres roles in the branch.
	Roles types.List `tfsdk:"roles"`
}

func (to *ListRolesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListRolesResponse) {
	if !from.Roles.IsNull() && !from.Roles.IsUnknown() && to.Roles.IsNull() && len(from.Roles.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Roles, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Roles = from.Roles
	}
}

func (to *ListRolesResponse) SyncFieldsDuringRead(ctx context.Context, from ListRolesResponse) {
	if !from.Roles.IsNull() && !from.Roles.IsUnknown() && to.Roles.IsNull() && len(from.Roles.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Roles, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Roles = from.Roles
	}
}

func (m ListRolesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListRolesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"roles": reflect.TypeOf(Role{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRolesResponse
// only implements ToObjectValue() and Type().
func (m ListRolesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"roles":           m.Roles,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListRolesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"roles": basetypes.ListType{
				ElemType: Role{}.Type(ctx),
			},
		},
	}
}

// GetRoles returns the value of the Roles field in ListRolesResponse as
// a slice of Role values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListRolesResponse) GetRoles(ctx context.Context) ([]Role, bool) {
	if m.Roles.IsNull() || m.Roles.IsUnknown() {
		return nil, false
	}
	var v []Role
	d := m.Roles.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRoles sets the value of the Roles field in ListRolesResponse.
func (m *ListRolesResponse) SetRoles(ctx context.Context, v []Role) {
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
type Operation struct {
	// If the value is `false`, it means the operation is still in progress. If
	// `true`, the operation is completed, and either `error` or `response` is
	// available.
	Done types.Bool `tfsdk:"done"`
	// The error result of the operation in case of failure or cancellation.
	Error types.Object `tfsdk:"error"`
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

func (to *Operation) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Operation) {
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

func (to *Operation) SyncFieldsDuringRead(ctx context.Context, from Operation) {
	if !from.Error.IsNull() && !from.Error.IsUnknown() {
		if toError, ok := to.GetError(ctx); ok {
			if fromError, ok := from.GetError(ctx); ok {
				toError.SyncFieldsDuringRead(ctx, fromError)
				to.SetError(ctx, toError)
			}
		}
	}
}

func (m Operation) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["done"] = attrs["done"].SetOptional()
	attrs["error"] = attrs["error"].SetOptional()
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
func (m Operation) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error": reflect.TypeOf(DatabricksServiceExceptionWithDetailsProto{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Operation
// only implements ToObjectValue() and Type().
func (m Operation) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Operation) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"done":     types.BoolType,
			"error":    DatabricksServiceExceptionWithDetailsProto{}.Type(ctx),
			"metadata": types.ObjectType{},
			"name":     types.StringType,
			"response": types.ObjectType{},
		},
	}
}

// GetError returns the value of the Error field in Operation as
// a DatabricksServiceExceptionWithDetailsProto value.
// If the field is unknown or null, the boolean return value is false.
func (m *Operation) GetError(ctx context.Context) (DatabricksServiceExceptionWithDetailsProto, bool) {
	var e DatabricksServiceExceptionWithDetailsProto
	if m.Error.IsNull() || m.Error.IsUnknown() {
		return e, false
	}
	var v DatabricksServiceExceptionWithDetailsProto
	d := m.Error.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetError sets the value of the Error field in Operation.
func (m *Operation) SetError(ctx context.Context, v DatabricksServiceExceptionWithDetailsProto) {
	vs := v.ToObjectValue(ctx)
	m.Error = vs
}

type Project struct {
	// A timestamp indicating when the project was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The resource name of the project. This field is output-only and
	// constructed by the system. Format: `projects/{project_id}`
	Name types.String `tfsdk:"name"`
	// The spec contains the project configuration, including display_name,
	// pg_version (Postgres version), history_retention_duration, and
	// default_endpoint_settings.
	Spec types.Object `tfsdk:"spec"`
	// The current status of a Project.
	Status types.Object `tfsdk:"status"`
	// System-generated unique ID for the project.
	Uid types.String `tfsdk:"uid"`
	// A timestamp indicating when the project was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *Project) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Project) {
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

func (to *Project) SyncFieldsDuringRead(ctx context.Context, from Project) {
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

func (m Project) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["spec"] = attrs["spec"].SetOptional()
	attrs["spec"] = attrs["spec"].SetComputed()
	attrs["spec"] = attrs["spec"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetComputed()
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
func (m Project) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spec":   reflect.TypeOf(ProjectSpec{}),
		"status": reflect.TypeOf(ProjectStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Project
// only implements ToObjectValue() and Type().
func (m Project) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Project) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": timetypes.RFC3339{}.Type(ctx),
			"name":        types.StringType,
			"spec":        ProjectSpec{}.Type(ctx),
			"status":      ProjectStatus{}.Type(ctx),
			"uid":         types.StringType,
			"update_time": timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetSpec returns the value of the Spec field in Project as
// a ProjectSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *Project) GetSpec(ctx context.Context) (ProjectSpec, bool) {
	var e ProjectSpec
	if m.Spec.IsNull() || m.Spec.IsUnknown() {
		return e, false
	}
	var v ProjectSpec
	d := m.Spec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSpec sets the value of the Spec field in Project.
func (m *Project) SetSpec(ctx context.Context, v ProjectSpec) {
	vs := v.ToObjectValue(ctx)
	m.Spec = vs
}

// GetStatus returns the value of the Status field in Project as
// a ProjectStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *Project) GetStatus(ctx context.Context) (ProjectStatus, bool) {
	var e ProjectStatus
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v ProjectStatus
	d := m.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in Project.
func (m *Project) SetStatus(ctx context.Context, v ProjectStatus) {
	vs := v.ToObjectValue(ctx)
	m.Status = vs
}

// A collection of settings for a compute endpoint.
type ProjectDefaultEndpointSettings struct {
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

func (to *ProjectDefaultEndpointSettings) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProjectDefaultEndpointSettings) {
}

func (to *ProjectDefaultEndpointSettings) SyncFieldsDuringRead(ctx context.Context, from ProjectDefaultEndpointSettings) {
}

func (m ProjectDefaultEndpointSettings) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ProjectDefaultEndpointSettings) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"pg_settings": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProjectDefaultEndpointSettings
// only implements ToObjectValue() and Type().
func (m ProjectDefaultEndpointSettings) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ProjectDefaultEndpointSettings) Type(ctx context.Context) attr.Type {
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

// GetPgSettings returns the value of the PgSettings field in ProjectDefaultEndpointSettings as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ProjectDefaultEndpointSettings) GetPgSettings(ctx context.Context) (map[string]types.String, bool) {
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

// SetPgSettings sets the value of the PgSettings field in ProjectDefaultEndpointSettings.
func (m *ProjectDefaultEndpointSettings) SetPgSettings(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["pg_settings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PgSettings = types.MapValueMust(t, vs)
}

type ProjectOperationMetadata struct {
}

func (to *ProjectOperationMetadata) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProjectOperationMetadata) {
}

func (to *ProjectOperationMetadata) SyncFieldsDuringRead(ctx context.Context, from ProjectOperationMetadata) {
}

func (m ProjectOperationMetadata) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProjectOperationMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ProjectOperationMetadata) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProjectOperationMetadata
// only implements ToObjectValue() and Type().
func (m ProjectOperationMetadata) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ProjectOperationMetadata) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ProjectSpec struct {
	DefaultEndpointSettings types.Object `tfsdk:"default_endpoint_settings"`
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

func (to *ProjectSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProjectSpec) {
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

func (to *ProjectSpec) SyncFieldsDuringRead(ctx context.Context, from ProjectSpec) {
	if !from.DefaultEndpointSettings.IsNull() && !from.DefaultEndpointSettings.IsUnknown() {
		if toDefaultEndpointSettings, ok := to.GetDefaultEndpointSettings(ctx); ok {
			if fromDefaultEndpointSettings, ok := from.GetDefaultEndpointSettings(ctx); ok {
				toDefaultEndpointSettings.SyncFieldsDuringRead(ctx, fromDefaultEndpointSettings)
				to.SetDefaultEndpointSettings(ctx, toDefaultEndpointSettings)
			}
		}
	}
}

func (m ProjectSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_endpoint_settings"] = attrs["default_endpoint_settings"].SetOptional()
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
func (m ProjectSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_endpoint_settings": reflect.TypeOf(ProjectDefaultEndpointSettings{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProjectSpec
// only implements ToObjectValue() and Type().
func (m ProjectSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ProjectSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_endpoint_settings":  ProjectDefaultEndpointSettings{}.Type(ctx),
			"display_name":               types.StringType,
			"history_retention_duration": timetypes.GoDuration{}.Type(ctx),
			"pg_version":                 types.Int64Type,
		},
	}
}

// GetDefaultEndpointSettings returns the value of the DefaultEndpointSettings field in ProjectSpec as
// a ProjectDefaultEndpointSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *ProjectSpec) GetDefaultEndpointSettings(ctx context.Context) (ProjectDefaultEndpointSettings, bool) {
	var e ProjectDefaultEndpointSettings
	if m.DefaultEndpointSettings.IsNull() || m.DefaultEndpointSettings.IsUnknown() {
		return e, false
	}
	var v ProjectDefaultEndpointSettings
	d := m.DefaultEndpointSettings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDefaultEndpointSettings sets the value of the DefaultEndpointSettings field in ProjectSpec.
func (m *ProjectSpec) SetDefaultEndpointSettings(ctx context.Context, v ProjectDefaultEndpointSettings) {
	vs := v.ToObjectValue(ctx)
	m.DefaultEndpointSettings = vs
}

type ProjectStatus struct {
	// The logical size limit for a branch.
	BranchLogicalSizeLimitBytes types.Int64 `tfsdk:"branch_logical_size_limit_bytes"`
	// The effective default endpoint settings.
	DefaultEndpointSettings types.Object `tfsdk:"default_endpoint_settings"`
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

func (to *ProjectStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ProjectStatus) {
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

func (to *ProjectStatus) SyncFieldsDuringRead(ctx context.Context, from ProjectStatus) {
	if !from.DefaultEndpointSettings.IsNull() && !from.DefaultEndpointSettings.IsUnknown() {
		if toDefaultEndpointSettings, ok := to.GetDefaultEndpointSettings(ctx); ok {
			if fromDefaultEndpointSettings, ok := from.GetDefaultEndpointSettings(ctx); ok {
				toDefaultEndpointSettings.SyncFieldsDuringRead(ctx, fromDefaultEndpointSettings)
				to.SetDefaultEndpointSettings(ctx, toDefaultEndpointSettings)
			}
		}
	}
}

func (m ProjectStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch_logical_size_limit_bytes"] = attrs["branch_logical_size_limit_bytes"].SetComputed()
	attrs["default_endpoint_settings"] = attrs["default_endpoint_settings"].SetComputed()
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
func (m ProjectStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_endpoint_settings": reflect.TypeOf(ProjectDefaultEndpointSettings{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProjectStatus
// only implements ToObjectValue() and Type().
func (m ProjectStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ProjectStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch_logical_size_limit_bytes": types.Int64Type,
			"default_endpoint_settings":       ProjectDefaultEndpointSettings{}.Type(ctx),
			"display_name":                    types.StringType,
			"history_retention_duration":      timetypes.GoDuration{}.Type(ctx),
			"owner":                           types.StringType,
			"pg_version":                      types.Int64Type,
			"synthetic_storage_size_bytes":    types.Int64Type,
		},
	}
}

// GetDefaultEndpointSettings returns the value of the DefaultEndpointSettings field in ProjectStatus as
// a ProjectDefaultEndpointSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *ProjectStatus) GetDefaultEndpointSettings(ctx context.Context) (ProjectDefaultEndpointSettings, bool) {
	var e ProjectDefaultEndpointSettings
	if m.DefaultEndpointSettings.IsNull() || m.DefaultEndpointSettings.IsUnknown() {
		return e, false
	}
	var v ProjectDefaultEndpointSettings
	d := m.DefaultEndpointSettings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDefaultEndpointSettings sets the value of the DefaultEndpointSettings field in ProjectStatus.
func (m *ProjectStatus) SetDefaultEndpointSettings(ctx context.Context, v ProjectDefaultEndpointSettings) {
	vs := v.ToObjectValue(ctx)
	m.DefaultEndpointSettings = vs
}

type RequestedClaims struct {
	PermissionSet types.String `tfsdk:"permission_set"`

	Resources types.List `tfsdk:"resources"`
}

func (to *RequestedClaims) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RequestedClaims) {
	if !from.Resources.IsNull() && !from.Resources.IsUnknown() && to.Resources.IsNull() && len(from.Resources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Resources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Resources = from.Resources
	}
}

func (to *RequestedClaims) SyncFieldsDuringRead(ctx context.Context, from RequestedClaims) {
	if !from.Resources.IsNull() && !from.Resources.IsUnknown() && to.Resources.IsNull() && len(from.Resources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Resources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Resources = from.Resources
	}
}

func (m RequestedClaims) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RequestedClaims) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"resources": reflect.TypeOf(RequestedResource{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RequestedClaims
// only implements ToObjectValue() and Type().
func (m RequestedClaims) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_set": m.PermissionSet,
			"resources":      m.Resources,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RequestedClaims) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_set": types.StringType,
			"resources": basetypes.ListType{
				ElemType: RequestedResource{}.Type(ctx),
			},
		},
	}
}

// GetResources returns the value of the Resources field in RequestedClaims as
// a slice of RequestedResource values.
// If the field is unknown or null, the boolean return value is false.
func (m *RequestedClaims) GetResources(ctx context.Context) ([]RequestedResource, bool) {
	if m.Resources.IsNull() || m.Resources.IsUnknown() {
		return nil, false
	}
	var v []RequestedResource
	d := m.Resources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResources sets the value of the Resources field in RequestedClaims.
func (m *RequestedClaims) SetResources(ctx context.Context, v []RequestedResource) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Resources = types.ListValueMust(t, vs)
}

type RequestedResource struct {
	TableName types.String `tfsdk:"table_name"`

	UnspecifiedResourceName types.String `tfsdk:"unspecified_resource_name"`
}

func (to *RequestedResource) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RequestedResource) {
}

func (to *RequestedResource) SyncFieldsDuringRead(ctx context.Context, from RequestedResource) {
}

func (m RequestedResource) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RequestedResource) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RequestedResource
// only implements ToObjectValue() and Type().
func (m RequestedResource) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table_name":                m.TableName,
			"unspecified_resource_name": m.UnspecifiedResourceName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RequestedResource) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_name":                types.StringType,
			"unspecified_resource_name": types.StringType,
		},
	}
}

// Role represents a Postgres role within a Branch.
type Role struct {
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The resource name of the role. Format:
	// projects/{project_id}/branches/{branch_id}/roles/{role_id}
	Name types.String `tfsdk:"name"`
	// The Branch where this Role exists. Format:
	// projects/{project_id}/branches/{branch_id}
	Parent types.String `tfsdk:"parent"`
	// The spec contains the role configuration, including identity type,
	// authentication method, and role attributes.
	Spec types.Object `tfsdk:"spec"`
	// Current status of the role, including its identity type, authentication
	// method, and role attributes.
	Status types.Object `tfsdk:"status"`

	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *Role) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Role) {
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

func (to *Role) SyncFieldsDuringRead(ctx context.Context, from Role) {
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

func (m Role) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["parent"] = attrs["parent"].SetComputed()
	attrs["spec"] = attrs["spec"].SetOptional()
	attrs["spec"] = attrs["spec"].SetComputed()
	attrs["spec"] = attrs["spec"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetComputed()
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
func (m Role) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spec":   reflect.TypeOf(RoleRoleSpec{}),
		"status": reflect.TypeOf(RoleRoleStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Role
// only implements ToObjectValue() and Type().
func (m Role) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Role) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": timetypes.RFC3339{}.Type(ctx),
			"name":        types.StringType,
			"parent":      types.StringType,
			"spec":        RoleRoleSpec{}.Type(ctx),
			"status":      RoleRoleStatus{}.Type(ctx),
			"update_time": timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetSpec returns the value of the Spec field in Role as
// a RoleRoleSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *Role) GetSpec(ctx context.Context) (RoleRoleSpec, bool) {
	var e RoleRoleSpec
	if m.Spec.IsNull() || m.Spec.IsUnknown() {
		return e, false
	}
	var v RoleRoleSpec
	d := m.Spec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSpec sets the value of the Spec field in Role.
func (m *Role) SetSpec(ctx context.Context, v RoleRoleSpec) {
	vs := v.ToObjectValue(ctx)
	m.Spec = vs
}

// GetStatus returns the value of the Status field in Role as
// a RoleRoleStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *Role) GetStatus(ctx context.Context) (RoleRoleStatus, bool) {
	var e RoleRoleStatus
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v RoleRoleStatus
	d := m.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in Role.
func (m *Role) SetStatus(ctx context.Context, v RoleRoleStatus) {
	vs := v.ToObjectValue(ctx)
	m.Status = vs
}

type RoleOperationMetadata struct {
}

func (to *RoleOperationMetadata) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RoleOperationMetadata) {
}

func (to *RoleOperationMetadata) SyncFieldsDuringRead(ctx context.Context, from RoleOperationMetadata) {
}

func (m RoleOperationMetadata) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RoleOperationMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RoleOperationMetadata) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RoleOperationMetadata
// only implements ToObjectValue() and Type().
func (m RoleOperationMetadata) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m RoleOperationMetadata) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RoleRoleSpec struct {
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

func (to *RoleRoleSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RoleRoleSpec) {
}

func (to *RoleRoleSpec) SyncFieldsDuringRead(ctx context.Context, from RoleRoleSpec) {
}

func (m RoleRoleSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RoleRoleSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RoleRoleSpec
// only implements ToObjectValue() and Type().
func (m RoleRoleSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auth_method":   m.AuthMethod,
			"identity_type": m.IdentityType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RoleRoleSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auth_method":   types.StringType,
			"identity_type": types.StringType,
		},
	}
}

type RoleRoleStatus struct {
	AuthMethod types.String `tfsdk:"auth_method"`
	// The type of the role.
	IdentityType types.String `tfsdk:"identity_type"`
}

func (to *RoleRoleStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RoleRoleStatus) {
}

func (to *RoleRoleStatus) SyncFieldsDuringRead(ctx context.Context, from RoleRoleStatus) {
}

func (m RoleRoleStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RoleRoleStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RoleRoleStatus
// only implements ToObjectValue() and Type().
func (m RoleRoleStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auth_method":   m.AuthMethod,
			"identity_type": m.IdentityType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RoleRoleStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auth_method":   types.StringType,
			"identity_type": types.StringType,
		},
	}
}

type UpdateBranchRequest struct {
	// The Branch to update.
	//
	// The branch's `name` field is used to identify the branch to update.
	// Format: projects/{project_id}/branches/{branch_id}
	Branch types.Object `tfsdk:"branch"`
	// The resource name of the branch. This field is output-only and
	// constructed by the system. Format:
	// `projects/{project_id}/branches/{branch_id}`
	Name types.String `tfsdk:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateBranchRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateBranchRequest) {
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

func (to *UpdateBranchRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateBranchRequest) {
	if !from.Branch.IsNull() && !from.Branch.IsUnknown() {
		if toBranch, ok := to.GetBranch(ctx); ok {
			if fromBranch, ok := from.GetBranch(ctx); ok {
				toBranch.SyncFieldsDuringRead(ctx, fromBranch)
				to.SetBranch(ctx, toBranch)
			}
		}
	}
}

func (m UpdateBranchRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch"] = attrs["branch"].SetRequired()
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
func (m UpdateBranchRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"branch": reflect.TypeOf(Branch{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateBranchRequest
// only implements ToObjectValue() and Type().
func (m UpdateBranchRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch":      m.Branch,
			"name":        m.Name,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateBranchRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch":      Branch{}.Type(ctx),
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetBranch returns the value of the Branch field in UpdateBranchRequest as
// a Branch value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateBranchRequest) GetBranch(ctx context.Context) (Branch, bool) {
	var e Branch
	if m.Branch.IsNull() || m.Branch.IsUnknown() {
		return e, false
	}
	var v Branch
	d := m.Branch.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBranch sets the value of the Branch field in UpdateBranchRequest.
func (m *UpdateBranchRequest) SetBranch(ctx context.Context, v Branch) {
	vs := v.ToObjectValue(ctx)
	m.Branch = vs
}

type UpdateEndpointRequest struct {
	// The Endpoint to update.
	//
	// The endpoint's `name` field is used to identify the endpoint to update.
	// Format:
	// projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}
	Endpoint types.Object `tfsdk:"endpoint"`
	// The resource name of the endpoint. This field is output-only and
	// constructed by the system. Format:
	// `projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}`
	Name types.String `tfsdk:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEndpointRequest) {
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

func (to *UpdateEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateEndpointRequest) {
	if !from.Endpoint.IsNull() && !from.Endpoint.IsUnknown() {
		if toEndpoint, ok := to.GetEndpoint(ctx); ok {
			if fromEndpoint, ok := from.GetEndpoint(ctx); ok {
				toEndpoint.SyncFieldsDuringRead(ctx, fromEndpoint)
				to.SetEndpoint(ctx, toEndpoint)
			}
		}
	}
}

func (m UpdateEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint"] = attrs["endpoint"].SetRequired()
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
func (m UpdateEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoint": reflect.TypeOf(Endpoint{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEndpointRequest
// only implements ToObjectValue() and Type().
func (m UpdateEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint":    m.Endpoint,
			"name":        m.Name,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint":    Endpoint{}.Type(ctx),
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetEndpoint returns the value of the Endpoint field in UpdateEndpointRequest as
// a Endpoint value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateEndpointRequest) GetEndpoint(ctx context.Context) (Endpoint, bool) {
	var e Endpoint
	if m.Endpoint.IsNull() || m.Endpoint.IsUnknown() {
		return e, false
	}
	var v Endpoint
	d := m.Endpoint.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEndpoint sets the value of the Endpoint field in UpdateEndpointRequest.
func (m *UpdateEndpointRequest) SetEndpoint(ctx context.Context, v Endpoint) {
	vs := v.ToObjectValue(ctx)
	m.Endpoint = vs
}

type UpdateProjectRequest struct {
	// The resource name of the project. This field is output-only and
	// constructed by the system. Format: `projects/{project_id}`
	Name types.String `tfsdk:"-"`
	// The Project to update.
	//
	// The project's `name` field is used to identify the project to update.
	// Format: projects/{project_id}
	Project types.Object `tfsdk:"project"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateProjectRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateProjectRequest) {
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

func (to *UpdateProjectRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateProjectRequest) {
	if !from.Project.IsNull() && !from.Project.IsUnknown() {
		if toProject, ok := to.GetProject(ctx); ok {
			if fromProject, ok := from.GetProject(ctx); ok {
				toProject.SyncFieldsDuringRead(ctx, fromProject)
				to.SetProject(ctx, toProject)
			}
		}
	}
}

func (m UpdateProjectRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project"] = attrs["project"].SetRequired()
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
func (m UpdateProjectRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"project": reflect.TypeOf(Project{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProjectRequest
// only implements ToObjectValue() and Type().
func (m UpdateProjectRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":        m.Name,
			"project":     m.Project,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateProjectRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":        types.StringType,
			"project":     Project{}.Type(ctx),
			"update_mask": types.StringType,
		},
	}
}

// GetProject returns the value of the Project field in UpdateProjectRequest as
// a Project value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateProjectRequest) GetProject(ctx context.Context) (Project, bool) {
	var e Project
	if m.Project.IsNull() || m.Project.IsUnknown() {
		return e, false
	}
	var v Project
	d := m.Project.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProject sets the value of the Project field in UpdateProjectRequest.
func (m *UpdateProjectRequest) SetProject(ctx context.Context, v Project) {
	vs := v.ToObjectValue(ctx)
	m.Project = vs
}
