// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package tags_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateTagAssignmentRequest struct {
	TagAssignment types.Object `tfsdk:"tag_assignment"`
}

func (to *CreateTagAssignmentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateTagAssignmentRequest) {
	if !from.TagAssignment.IsNull() && !from.TagAssignment.IsUnknown() {
		if toTagAssignment, ok := to.GetTagAssignment(ctx); ok {
			if fromTagAssignment, ok := from.GetTagAssignment(ctx); ok {
				// Recursively sync the fields of TagAssignment
				toTagAssignment.SyncFieldsDuringCreateOrUpdate(ctx, fromTagAssignment)
				to.SetTagAssignment(ctx, toTagAssignment)
			}
		}
	}
}

func (to *CreateTagAssignmentRequest) SyncFieldsDuringRead(ctx context.Context, from CreateTagAssignmentRequest) {
	if !from.TagAssignment.IsNull() && !from.TagAssignment.IsUnknown() {
		if toTagAssignment, ok := to.GetTagAssignment(ctx); ok {
			if fromTagAssignment, ok := from.GetTagAssignment(ctx); ok {
				toTagAssignment.SyncFieldsDuringRead(ctx, fromTagAssignment)
				to.SetTagAssignment(ctx, toTagAssignment)
			}
		}
	}
}

func (m CreateTagAssignmentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["tag_assignment"] = attrs["tag_assignment"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateTagAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateTagAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_assignment": reflect.TypeOf(TagAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTagAssignmentRequest
// only implements ToObjectValue() and Type().
func (m CreateTagAssignmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_assignment": m.TagAssignment,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateTagAssignmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tag_assignment": TagAssignment{}.Type(ctx),
		},
	}
}

// GetTagAssignment returns the value of the TagAssignment field in CreateTagAssignmentRequest as
// a TagAssignment value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateTagAssignmentRequest) GetTagAssignment(ctx context.Context) (TagAssignment, bool) {
	var e TagAssignment
	if m.TagAssignment.IsNull() || m.TagAssignment.IsUnknown() {
		return e, false
	}
	var v TagAssignment
	d := m.TagAssignment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTagAssignment sets the value of the TagAssignment field in CreateTagAssignmentRequest.
func (m *CreateTagAssignmentRequest) SetTagAssignment(ctx context.Context, v TagAssignment) {
	vs := v.ToObjectValue(ctx)
	m.TagAssignment = vs
}

type CreateTagPolicyRequest struct {
	TagPolicy types.Object `tfsdk:"tag_policy"`
}

func (to *CreateTagPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateTagPolicyRequest) {
	if !from.TagPolicy.IsNull() && !from.TagPolicy.IsUnknown() {
		if toTagPolicy, ok := to.GetTagPolicy(ctx); ok {
			if fromTagPolicy, ok := from.GetTagPolicy(ctx); ok {
				// Recursively sync the fields of TagPolicy
				toTagPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromTagPolicy)
				to.SetTagPolicy(ctx, toTagPolicy)
			}
		}
	}
}

func (to *CreateTagPolicyRequest) SyncFieldsDuringRead(ctx context.Context, from CreateTagPolicyRequest) {
	if !from.TagPolicy.IsNull() && !from.TagPolicy.IsUnknown() {
		if toTagPolicy, ok := to.GetTagPolicy(ctx); ok {
			if fromTagPolicy, ok := from.GetTagPolicy(ctx); ok {
				toTagPolicy.SyncFieldsDuringRead(ctx, fromTagPolicy)
				to.SetTagPolicy(ctx, toTagPolicy)
			}
		}
	}
}

func (m CreateTagPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["tag_policy"] = attrs["tag_policy"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateTagPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateTagPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_policy": reflect.TypeOf(TagPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTagPolicyRequest
// only implements ToObjectValue() and Type().
func (m CreateTagPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_policy": m.TagPolicy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateTagPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tag_policy": TagPolicy{}.Type(ctx),
		},
	}
}

// GetTagPolicy returns the value of the TagPolicy field in CreateTagPolicyRequest as
// a TagPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateTagPolicyRequest) GetTagPolicy(ctx context.Context) (TagPolicy, bool) {
	var e TagPolicy
	if m.TagPolicy.IsNull() || m.TagPolicy.IsUnknown() {
		return e, false
	}
	var v TagPolicy
	d := m.TagPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTagPolicy sets the value of the TagPolicy field in CreateTagPolicyRequest.
func (m *CreateTagPolicyRequest) SetTagPolicy(ctx context.Context, v TagPolicy) {
	vs := v.ToObjectValue(ctx)
	m.TagPolicy = vs
}

type DeleteTagAssignmentRequest struct {
	// The identifier of the entity to which the tag is assigned
	EntityId types.String `tfsdk:"-"`
	// The type of entity to which the tag is assigned. Allowed value is
	// dashboards
	EntityType types.String `tfsdk:"-"`
	// The key of the tag. The characters , . : / - = and leading/trailing
	// spaces are not allowed
	TagKey types.String `tfsdk:"-"`
}

func (to *DeleteTagAssignmentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteTagAssignmentRequest) {
}

func (to *DeleteTagAssignmentRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteTagAssignmentRequest) {
}

func (m DeleteTagAssignmentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["entity_type"] = attrs["entity_type"].SetRequired()
	attrs["entity_id"] = attrs["entity_id"].SetRequired()
	attrs["tag_key"] = attrs["tag_key"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTagAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteTagAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTagAssignmentRequest
// only implements ToObjectValue() and Type().
func (m DeleteTagAssignmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_id":   m.EntityId,
			"entity_type": m.EntityType,
			"tag_key":     m.TagKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteTagAssignmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_id":   types.StringType,
			"entity_type": types.StringType,
			"tag_key":     types.StringType,
		},
	}
}

type DeleteTagPolicyRequest struct {
	TagKey types.String `tfsdk:"-"`
}

func (to *DeleteTagPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteTagPolicyRequest) {
}

func (to *DeleteTagPolicyRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteTagPolicyRequest) {
}

func (m DeleteTagPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["tag_key"] = attrs["tag_key"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTagPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteTagPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTagPolicyRequest
// only implements ToObjectValue() and Type().
func (m DeleteTagPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_key": m.TagKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteTagPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tag_key": types.StringType,
		},
	}
}

type GetTagAssignmentRequest struct {
	// The identifier of the entity to which the tag is assigned
	EntityId types.String `tfsdk:"-"`
	// The type of entity to which the tag is assigned. Allowed value is
	// dashboards
	EntityType types.String `tfsdk:"-"`
	// The key of the tag. The characters , . : / - = and leading/trailing
	// spaces are not allowed
	TagKey types.String `tfsdk:"-"`
}

func (to *GetTagAssignmentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetTagAssignmentRequest) {
}

func (to *GetTagAssignmentRequest) SyncFieldsDuringRead(ctx context.Context, from GetTagAssignmentRequest) {
}

func (m GetTagAssignmentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["entity_type"] = attrs["entity_type"].SetRequired()
	attrs["entity_id"] = attrs["entity_id"].SetRequired()
	attrs["tag_key"] = attrs["tag_key"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetTagAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetTagAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTagAssignmentRequest
// only implements ToObjectValue() and Type().
func (m GetTagAssignmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_id":   m.EntityId,
			"entity_type": m.EntityType,
			"tag_key":     m.TagKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetTagAssignmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_id":   types.StringType,
			"entity_type": types.StringType,
			"tag_key":     types.StringType,
		},
	}
}

type GetTagPolicyRequest struct {
	TagKey types.String `tfsdk:"-"`
}

func (to *GetTagPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetTagPolicyRequest) {
}

func (to *GetTagPolicyRequest) SyncFieldsDuringRead(ctx context.Context, from GetTagPolicyRequest) {
}

func (m GetTagPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["tag_key"] = attrs["tag_key"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetTagPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetTagPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTagPolicyRequest
// only implements ToObjectValue() and Type().
func (m GetTagPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_key": m.TagKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetTagPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tag_key": types.StringType,
		},
	}
}

type ListTagAssignmentsRequest struct {
	// The identifier of the entity to which the tag is assigned
	EntityId types.String `tfsdk:"-"`
	// The type of entity to which the tag is assigned. Allowed value is
	// dashboards
	EntityType types.String `tfsdk:"-"`
	// Optional. Maximum number of tag assignments to return in a single page
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of tag assignments. Requests
	// first page if absent.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListTagAssignmentsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTagAssignmentsRequest) {
}

func (to *ListTagAssignmentsRequest) SyncFieldsDuringRead(ctx context.Context, from ListTagAssignmentsRequest) {
}

func (m ListTagAssignmentsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["entity_type"] = attrs["entity_type"].SetRequired()
	attrs["entity_id"] = attrs["entity_id"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTagAssignmentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListTagAssignmentsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTagAssignmentsRequest
// only implements ToObjectValue() and Type().
func (m ListTagAssignmentsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_id":   m.EntityId,
			"entity_type": m.EntityType,
			"page_size":   m.PageSize,
			"page_token":  m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListTagAssignmentsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_id":   types.StringType,
			"entity_type": types.StringType,
			"page_size":   types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListTagAssignmentsResponse struct {
	// Pagination token to request the next page of tag assignments
	NextPageToken types.String `tfsdk:"next_page_token"`

	TagAssignments types.List `tfsdk:"tag_assignments"`
}

func (to *ListTagAssignmentsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTagAssignmentsResponse) {
	if !from.TagAssignments.IsNull() && !from.TagAssignments.IsUnknown() && to.TagAssignments.IsNull() && len(from.TagAssignments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TagAssignments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TagAssignments = from.TagAssignments
	}
}

func (to *ListTagAssignmentsResponse) SyncFieldsDuringRead(ctx context.Context, from ListTagAssignmentsResponse) {
	if !from.TagAssignments.IsNull() && !from.TagAssignments.IsUnknown() && to.TagAssignments.IsNull() && len(from.TagAssignments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TagAssignments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TagAssignments = from.TagAssignments
	}
}

func (m ListTagAssignmentsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["tag_assignments"] = attrs["tag_assignments"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTagAssignmentsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListTagAssignmentsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_assignments": reflect.TypeOf(TagAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTagAssignmentsResponse
// only implements ToObjectValue() and Type().
func (m ListTagAssignmentsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"tag_assignments": m.TagAssignments,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListTagAssignmentsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"tag_assignments": basetypes.ListType{
				ElemType: TagAssignment{}.Type(ctx),
			},
		},
	}
}

// GetTagAssignments returns the value of the TagAssignments field in ListTagAssignmentsResponse as
// a slice of TagAssignment values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListTagAssignmentsResponse) GetTagAssignments(ctx context.Context) ([]TagAssignment, bool) {
	if m.TagAssignments.IsNull() || m.TagAssignments.IsUnknown() {
		return nil, false
	}
	var v []TagAssignment
	d := m.TagAssignments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTagAssignments sets the value of the TagAssignments field in ListTagAssignmentsResponse.
func (m *ListTagAssignmentsResponse) SetTagAssignments(ctx context.Context, v []TagAssignment) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tag_assignments"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TagAssignments = types.ListValueMust(t, vs)
}

type ListTagPoliciesRequest struct {
	// The maximum number of results to return in this request. Fewer results
	// may be returned than requested. If unspecified or set to 0, this defaults
	// to 1000. The maximum value is 1000; values above 1000 will be coerced
	// down to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// An optional page token received from a previous list tag policies call.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListTagPoliciesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTagPoliciesRequest) {
}

func (to *ListTagPoliciesRequest) SyncFieldsDuringRead(ctx context.Context, from ListTagPoliciesRequest) {
}

func (m ListTagPoliciesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTagPoliciesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListTagPoliciesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTagPoliciesRequest
// only implements ToObjectValue() and Type().
func (m ListTagPoliciesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListTagPoliciesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListTagPoliciesResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	TagPolicies types.List `tfsdk:"tag_policies"`
}

func (to *ListTagPoliciesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTagPoliciesResponse) {
	if !from.TagPolicies.IsNull() && !from.TagPolicies.IsUnknown() && to.TagPolicies.IsNull() && len(from.TagPolicies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TagPolicies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TagPolicies = from.TagPolicies
	}
}

func (to *ListTagPoliciesResponse) SyncFieldsDuringRead(ctx context.Context, from ListTagPoliciesResponse) {
	if !from.TagPolicies.IsNull() && !from.TagPolicies.IsUnknown() && to.TagPolicies.IsNull() && len(from.TagPolicies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TagPolicies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TagPolicies = from.TagPolicies
	}
}

func (m ListTagPoliciesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["tag_policies"] = attrs["tag_policies"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTagPoliciesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListTagPoliciesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_policies": reflect.TypeOf(TagPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTagPoliciesResponse
// only implements ToObjectValue() and Type().
func (m ListTagPoliciesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"tag_policies":    m.TagPolicies,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListTagPoliciesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"tag_policies": basetypes.ListType{
				ElemType: TagPolicy{}.Type(ctx),
			},
		},
	}
}

// GetTagPolicies returns the value of the TagPolicies field in ListTagPoliciesResponse as
// a slice of TagPolicy values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListTagPoliciesResponse) GetTagPolicies(ctx context.Context) ([]TagPolicy, bool) {
	if m.TagPolicies.IsNull() || m.TagPolicies.IsUnknown() {
		return nil, false
	}
	var v []TagPolicy
	d := m.TagPolicies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTagPolicies sets the value of the TagPolicies field in ListTagPoliciesResponse.
func (m *ListTagPoliciesResponse) SetTagPolicies(ctx context.Context, v []TagPolicy) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tag_policies"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TagPolicies = types.ListValueMust(t, vs)
}

type TagAssignment struct {
	// The identifier of the entity to which the tag is assigned
	EntityId types.String `tfsdk:"entity_id"`
	// The type of entity to which the tag is assigned. Allowed value is
	// dashboards
	EntityType types.String `tfsdk:"entity_type"`
	// The key of the tag. The characters , . : / - = and leading/trailing
	// spaces are not allowed
	TagKey types.String `tfsdk:"tag_key"`
	// The value of the tag
	TagValue types.String `tfsdk:"tag_value"`
}

func (to *TagAssignment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TagAssignment) {
}

func (to *TagAssignment) SyncFieldsDuringRead(ctx context.Context, from TagAssignment) {
}

func (m TagAssignment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["entity_id"] = attrs["entity_id"].SetRequired()
	attrs["entity_id"] = attrs["entity_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["entity_type"] = attrs["entity_type"].SetRequired()
	attrs["entity_type"] = attrs["entity_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["tag_key"] = attrs["tag_key"].SetRequired()
	attrs["tag_key"] = attrs["tag_key"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["tag_value"] = attrs["tag_value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TagAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TagAssignment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TagAssignment
// only implements ToObjectValue() and Type().
func (m TagAssignment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_id":   m.EntityId,
			"entity_type": m.EntityType,
			"tag_key":     m.TagKey,
			"tag_value":   m.TagValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TagAssignment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_id":   types.StringType,
			"entity_type": types.StringType,
			"tag_key":     types.StringType,
			"tag_value":   types.StringType,
		},
	}
}

type TagPolicy struct {
	// Timestamp when the tag policy was created
	CreateTime types.String `tfsdk:"create_time"`

	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`

	TagKey types.String `tfsdk:"tag_key"`
	// Timestamp when the tag policy was last updated
	UpdateTime types.String `tfsdk:"update_time"`

	Values types.List `tfsdk:"values"`
}

func (to *TagPolicy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TagPolicy) {
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (to *TagPolicy) SyncFieldsDuringRead(ctx context.Context, from TagPolicy) {
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (m TagPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["tag_key"] = attrs["tag_key"].SetRequired()
	attrs["tag_key"] = attrs["tag_key"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["values"] = attrs["values"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TagPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TagPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"values": reflect.TypeOf(Value{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TagPolicy
// only implements ToObjectValue() and Type().
func (m TagPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time": m.CreateTime,
			"description": m.Description,
			"id":          m.Id,
			"tag_key":     m.TagKey,
			"update_time": m.UpdateTime,
			"values":      m.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TagPolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": types.StringType,
			"description": types.StringType,
			"id":          types.StringType,
			"tag_key":     types.StringType,
			"update_time": types.StringType,
			"values": basetypes.ListType{
				ElemType: Value{}.Type(ctx),
			},
		},
	}
}

// GetValues returns the value of the Values field in TagPolicy as
// a slice of Value values.
// If the field is unknown or null, the boolean return value is false.
func (m *TagPolicy) GetValues(ctx context.Context) ([]Value, bool) {
	if m.Values.IsNull() || m.Values.IsUnknown() {
		return nil, false
	}
	var v []Value
	d := m.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in TagPolicy.
func (m *TagPolicy) SetValues(ctx context.Context, v []Value) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Values = types.ListValueMust(t, vs)
}

type UpdateTagAssignmentRequest struct {
	// The identifier of the entity to which the tag is assigned
	EntityId types.String `tfsdk:"-"`
	// The type of entity to which the tag is assigned. Allowed value is
	// dashboards
	EntityType types.String `tfsdk:"-"`

	TagAssignment types.Object `tfsdk:"tag_assignment"`
	// The key of the tag. The characters , . : / - = and leading/trailing
	// spaces are not allowed
	TagKey types.String `tfsdk:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateTagAssignmentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateTagAssignmentRequest) {
	if !from.TagAssignment.IsNull() && !from.TagAssignment.IsUnknown() {
		if toTagAssignment, ok := to.GetTagAssignment(ctx); ok {
			if fromTagAssignment, ok := from.GetTagAssignment(ctx); ok {
				// Recursively sync the fields of TagAssignment
				toTagAssignment.SyncFieldsDuringCreateOrUpdate(ctx, fromTagAssignment)
				to.SetTagAssignment(ctx, toTagAssignment)
			}
		}
	}
}

func (to *UpdateTagAssignmentRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateTagAssignmentRequest) {
	if !from.TagAssignment.IsNull() && !from.TagAssignment.IsUnknown() {
		if toTagAssignment, ok := to.GetTagAssignment(ctx); ok {
			if fromTagAssignment, ok := from.GetTagAssignment(ctx); ok {
				toTagAssignment.SyncFieldsDuringRead(ctx, fromTagAssignment)
				to.SetTagAssignment(ctx, toTagAssignment)
			}
		}
	}
}

func (m UpdateTagAssignmentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["tag_assignment"] = attrs["tag_assignment"].SetRequired()
	attrs["entity_type"] = attrs["entity_type"].SetRequired()
	attrs["entity_type"] = attrs["entity_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["entity_id"] = attrs["entity_id"].SetRequired()
	attrs["entity_id"] = attrs["entity_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["tag_key"] = attrs["tag_key"].SetRequired()
	attrs["tag_key"] = attrs["tag_key"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateTagAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateTagAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_assignment": reflect.TypeOf(TagAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateTagAssignmentRequest
// only implements ToObjectValue() and Type().
func (m UpdateTagAssignmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_id":      m.EntityId,
			"entity_type":    m.EntityType,
			"tag_assignment": m.TagAssignment,
			"tag_key":        m.TagKey,
			"update_mask":    m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateTagAssignmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_id":      types.StringType,
			"entity_type":    types.StringType,
			"tag_assignment": TagAssignment{}.Type(ctx),
			"tag_key":        types.StringType,
			"update_mask":    types.StringType,
		},
	}
}

// GetTagAssignment returns the value of the TagAssignment field in UpdateTagAssignmentRequest as
// a TagAssignment value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateTagAssignmentRequest) GetTagAssignment(ctx context.Context) (TagAssignment, bool) {
	var e TagAssignment
	if m.TagAssignment.IsNull() || m.TagAssignment.IsUnknown() {
		return e, false
	}
	var v TagAssignment
	d := m.TagAssignment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTagAssignment sets the value of the TagAssignment field in UpdateTagAssignmentRequest.
func (m *UpdateTagAssignmentRequest) SetTagAssignment(ctx context.Context, v TagAssignment) {
	vs := v.ToObjectValue(ctx)
	m.TagAssignment = vs
}

type UpdateTagPolicyRequest struct {
	TagKey types.String `tfsdk:"-"`

	TagPolicy types.Object `tfsdk:"tag_policy"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateTagPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateTagPolicyRequest) {
	if !from.TagPolicy.IsNull() && !from.TagPolicy.IsUnknown() {
		if toTagPolicy, ok := to.GetTagPolicy(ctx); ok {
			if fromTagPolicy, ok := from.GetTagPolicy(ctx); ok {
				// Recursively sync the fields of TagPolicy
				toTagPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromTagPolicy)
				to.SetTagPolicy(ctx, toTagPolicy)
			}
		}
	}
}

func (to *UpdateTagPolicyRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateTagPolicyRequest) {
	if !from.TagPolicy.IsNull() && !from.TagPolicy.IsUnknown() {
		if toTagPolicy, ok := to.GetTagPolicy(ctx); ok {
			if fromTagPolicy, ok := from.GetTagPolicy(ctx); ok {
				toTagPolicy.SyncFieldsDuringRead(ctx, fromTagPolicy)
				to.SetTagPolicy(ctx, toTagPolicy)
			}
		}
	}
}

func (m UpdateTagPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["tag_policy"] = attrs["tag_policy"].SetRequired()
	attrs["tag_key"] = attrs["tag_key"].SetRequired()
	attrs["tag_key"] = attrs["tag_key"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateTagPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateTagPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_policy": reflect.TypeOf(TagPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateTagPolicyRequest
// only implements ToObjectValue() and Type().
func (m UpdateTagPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_key":     m.TagKey,
			"tag_policy":  m.TagPolicy,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateTagPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tag_key":     types.StringType,
			"tag_policy":  TagPolicy{}.Type(ctx),
			"update_mask": types.StringType,
		},
	}
}

// GetTagPolicy returns the value of the TagPolicy field in UpdateTagPolicyRequest as
// a TagPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateTagPolicyRequest) GetTagPolicy(ctx context.Context) (TagPolicy, bool) {
	var e TagPolicy
	if m.TagPolicy.IsNull() || m.TagPolicy.IsUnknown() {
		return e, false
	}
	var v TagPolicy
	d := m.TagPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTagPolicy sets the value of the TagPolicy field in UpdateTagPolicyRequest.
func (m *UpdateTagPolicyRequest) SetTagPolicy(ctx context.Context, v TagPolicy) {
	vs := v.ToObjectValue(ctx)
	m.TagPolicy = vs
}

type Value struct {
	Name types.String `tfsdk:"name"`
}

func (to *Value) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Value) {
}

func (to *Value) SyncFieldsDuringRead(ctx context.Context, from Value) {
}

func (m Value) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Value.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Value) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Value
// only implements ToObjectValue() and Type().
func (m Value) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Value) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}
