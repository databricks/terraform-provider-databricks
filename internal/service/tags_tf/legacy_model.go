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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateTagAssignmentRequest_SdkV2 struct {
	TagAssignment types.List `tfsdk:"tag_assignment"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateTagAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateTagAssignmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_assignment": reflect.TypeOf(TagAssignment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTagAssignmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateTagAssignmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_assignment": o.TagAssignment,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateTagAssignmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tag_assignment": basetypes.ListType{
				ElemType: TagAssignment_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTagAssignment returns the value of the TagAssignment field in CreateTagAssignmentRequest_SdkV2 as
// a TagAssignment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateTagAssignmentRequest_SdkV2) GetTagAssignment(ctx context.Context) (TagAssignment_SdkV2, bool) {
	var e TagAssignment_SdkV2
	if o.TagAssignment.IsNull() || o.TagAssignment.IsUnknown() {
		return e, false
	}
	var v []TagAssignment_SdkV2
	d := o.TagAssignment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTagAssignment sets the value of the TagAssignment field in CreateTagAssignmentRequest_SdkV2.
func (o *CreateTagAssignmentRequest_SdkV2) SetTagAssignment(ctx context.Context, v TagAssignment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tag_assignment"]
	o.TagAssignment = types.ListValueMust(t, vs)
}

type CreateTagPolicyRequest_SdkV2 struct {
	TagPolicy types.List `tfsdk:"tag_policy"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateTagPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateTagPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_policy": reflect.TypeOf(TagPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTagPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateTagPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_policy": o.TagPolicy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateTagPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tag_policy": basetypes.ListType{
				ElemType: TagPolicy_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTagPolicy returns the value of the TagPolicy field in CreateTagPolicyRequest_SdkV2 as
// a TagPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateTagPolicyRequest_SdkV2) GetTagPolicy(ctx context.Context) (TagPolicy_SdkV2, bool) {
	var e TagPolicy_SdkV2
	if o.TagPolicy.IsNull() || o.TagPolicy.IsUnknown() {
		return e, false
	}
	var v []TagPolicy_SdkV2
	d := o.TagPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTagPolicy sets the value of the TagPolicy field in CreateTagPolicyRequest_SdkV2.
func (o *CreateTagPolicyRequest_SdkV2) SetTagPolicy(ctx context.Context, v TagPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tag_policy"]
	o.TagPolicy = types.ListValueMust(t, vs)
}

type DeleteTagAssignmentRequest_SdkV2 struct {
	EntityId types.String `tfsdk:"-"`

	EntityType types.String `tfsdk:"-"`

	TagKey types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTagAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteTagAssignmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTagAssignmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteTagAssignmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_id":   o.EntityId,
			"entity_type": o.EntityType,
			"tag_key":     o.TagKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteTagAssignmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_id":   types.StringType,
			"entity_type": types.StringType,
			"tag_key":     types.StringType,
		},
	}
}

type DeleteTagPolicyRequest_SdkV2 struct {
	TagKey types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTagPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteTagPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTagPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteTagPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_key": o.TagKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteTagPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tag_key": types.StringType,
		},
	}
}

type GetTagAssignmentRequest_SdkV2 struct {
	EntityId types.String `tfsdk:"-"`

	EntityType types.String `tfsdk:"-"`

	TagKey types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetTagAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetTagAssignmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTagAssignmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetTagAssignmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_id":   o.EntityId,
			"entity_type": o.EntityType,
			"tag_key":     o.TagKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetTagAssignmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_id":   types.StringType,
			"entity_type": types.StringType,
			"tag_key":     types.StringType,
		},
	}
}

type GetTagPolicyRequest_SdkV2 struct {
	TagKey types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetTagPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetTagPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTagPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetTagPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_key": o.TagKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetTagPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tag_key": types.StringType,
		},
	}
}

type ListTagAssignmentsRequest_SdkV2 struct {
	EntityId types.String `tfsdk:"-"`

	EntityType types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTagAssignmentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListTagAssignmentsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTagAssignmentsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListTagAssignmentsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_id":   o.EntityId,
			"entity_type": o.EntityType,
			"page_size":   o.PageSize,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListTagAssignmentsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_id":   types.StringType,
			"entity_type": types.StringType,
			"page_size":   types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListTagAssignmentsResponse_SdkV2 struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	TagAssignments types.List `tfsdk:"tag_assignments"`
}

func (newState *ListTagAssignmentsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ListTagAssignmentsResponse_SdkV2) {
}

func (newState *ListTagAssignmentsResponse_SdkV2) SyncFieldsDuringRead(existingState ListTagAssignmentsResponse_SdkV2) {
}

func (c ListTagAssignmentsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListTagAssignmentsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_assignments": reflect.TypeOf(TagAssignment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTagAssignmentsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListTagAssignmentsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"tag_assignments": o.TagAssignments,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListTagAssignmentsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"tag_assignments": basetypes.ListType{
				ElemType: TagAssignment_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTagAssignments returns the value of the TagAssignments field in ListTagAssignmentsResponse_SdkV2 as
// a slice of TagAssignment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListTagAssignmentsResponse_SdkV2) GetTagAssignments(ctx context.Context) ([]TagAssignment_SdkV2, bool) {
	if o.TagAssignments.IsNull() || o.TagAssignments.IsUnknown() {
		return nil, false
	}
	var v []TagAssignment_SdkV2
	d := o.TagAssignments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTagAssignments sets the value of the TagAssignments field in ListTagAssignmentsResponse_SdkV2.
func (o *ListTagAssignmentsResponse_SdkV2) SetTagAssignments(ctx context.Context, v []TagAssignment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tag_assignments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TagAssignments = types.ListValueMust(t, vs)
}

type ListTagPoliciesRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTagPoliciesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListTagPoliciesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTagPoliciesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListTagPoliciesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListTagPoliciesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListTagPoliciesResponse_SdkV2 struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	TagPolicies types.List `tfsdk:"tag_policies"`
}

func (newState *ListTagPoliciesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ListTagPoliciesResponse_SdkV2) {
}

func (newState *ListTagPoliciesResponse_SdkV2) SyncFieldsDuringRead(existingState ListTagPoliciesResponse_SdkV2) {
}

func (c ListTagPoliciesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListTagPoliciesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_policies": reflect.TypeOf(TagPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTagPoliciesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListTagPoliciesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"tag_policies":    o.TagPolicies,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListTagPoliciesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"tag_policies": basetypes.ListType{
				ElemType: TagPolicy_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTagPolicies returns the value of the TagPolicies field in ListTagPoliciesResponse_SdkV2 as
// a slice of TagPolicy_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListTagPoliciesResponse_SdkV2) GetTagPolicies(ctx context.Context) ([]TagPolicy_SdkV2, bool) {
	if o.TagPolicies.IsNull() || o.TagPolicies.IsUnknown() {
		return nil, false
	}
	var v []TagPolicy_SdkV2
	d := o.TagPolicies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTagPolicies sets the value of the TagPolicies field in ListTagPoliciesResponse_SdkV2.
func (o *ListTagPoliciesResponse_SdkV2) SetTagPolicies(ctx context.Context, v []TagPolicy_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tag_policies"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TagPolicies = types.ListValueMust(t, vs)
}

type TagAssignment_SdkV2 struct {
	EntityId types.String `tfsdk:"entity_id"`

	EntityType types.String `tfsdk:"entity_type"`

	TagKey types.String `tfsdk:"tag_key"`

	TagValue types.String `tfsdk:"tag_value"`
}

func (newState *TagAssignment_SdkV2) SyncFieldsDuringCreateOrUpdate(plan TagAssignment_SdkV2) {
}

func (newState *TagAssignment_SdkV2) SyncFieldsDuringRead(existingState TagAssignment_SdkV2) {
}

func (c TagAssignment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["entity_id"] = attrs["entity_id"].SetRequired()
	attrs["entity_type"] = attrs["entity_type"].SetRequired()
	attrs["tag_key"] = attrs["tag_key"].SetRequired()
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
func (a TagAssignment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TagAssignment_SdkV2
// only implements ToObjectValue() and Type().
func (o TagAssignment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_id":   o.EntityId,
			"entity_type": o.EntityType,
			"tag_key":     o.TagKey,
			"tag_value":   o.TagValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TagAssignment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_id":   types.StringType,
			"entity_type": types.StringType,
			"tag_key":     types.StringType,
			"tag_value":   types.StringType,
		},
	}
}

type TagPolicy_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`

	TagKey types.String `tfsdk:"tag_key"`

	Values types.List `tfsdk:"values"`
}

func (newState *TagPolicy_SdkV2) SyncFieldsDuringCreateOrUpdate(plan TagPolicy_SdkV2) {
}

func (newState *TagPolicy_SdkV2) SyncFieldsDuringRead(existingState TagPolicy_SdkV2) {
}

func (c TagPolicy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["tag_key"] = attrs["tag_key"].SetRequired()
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
func (a TagPolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"values": reflect.TypeOf(Value_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TagPolicy_SdkV2
// only implements ToObjectValue() and Type().
func (o TagPolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": o.Description,
			"id":          o.Id,
			"tag_key":     o.TagKey,
			"values":      o.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TagPolicy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"id":          types.StringType,
			"tag_key":     types.StringType,
			"values": basetypes.ListType{
				ElemType: Value_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetValues returns the value of the Values field in TagPolicy_SdkV2 as
// a slice of Value_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *TagPolicy_SdkV2) GetValues(ctx context.Context) ([]Value_SdkV2, bool) {
	if o.Values.IsNull() || o.Values.IsUnknown() {
		return nil, false
	}
	var v []Value_SdkV2
	d := o.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in TagPolicy_SdkV2.
func (o *TagPolicy_SdkV2) SetValues(ctx context.Context, v []Value_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Values = types.ListValueMust(t, vs)
}

type UpdateTagAssignmentRequest_SdkV2 struct {
	EntityId types.String `tfsdk:"-"`

	EntityType types.String `tfsdk:"-"`

	TagAssignment types.List `tfsdk:"tag_assignment"`

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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateTagAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateTagAssignmentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_assignment": reflect.TypeOf(TagAssignment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateTagAssignmentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateTagAssignmentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_id":      o.EntityId,
			"entity_type":    o.EntityType,
			"tag_assignment": o.TagAssignment,
			"tag_key":        o.TagKey,
			"update_mask":    o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateTagAssignmentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_id":   types.StringType,
			"entity_type": types.StringType,
			"tag_assignment": basetypes.ListType{
				ElemType: TagAssignment_SdkV2{}.Type(ctx),
			},
			"tag_key":     types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetTagAssignment returns the value of the TagAssignment field in UpdateTagAssignmentRequest_SdkV2 as
// a TagAssignment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateTagAssignmentRequest_SdkV2) GetTagAssignment(ctx context.Context) (TagAssignment_SdkV2, bool) {
	var e TagAssignment_SdkV2
	if o.TagAssignment.IsNull() || o.TagAssignment.IsUnknown() {
		return e, false
	}
	var v []TagAssignment_SdkV2
	d := o.TagAssignment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTagAssignment sets the value of the TagAssignment field in UpdateTagAssignmentRequest_SdkV2.
func (o *UpdateTagAssignmentRequest_SdkV2) SetTagAssignment(ctx context.Context, v TagAssignment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tag_assignment"]
	o.TagAssignment = types.ListValueMust(t, vs)
}

type UpdateTagPolicyRequest_SdkV2 struct {
	TagKey types.String `tfsdk:"-"`

	TagPolicy types.List `tfsdk:"tag_policy"`
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateTagPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateTagPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_policy": reflect.TypeOf(TagPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateTagPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateTagPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_key":     o.TagKey,
			"tag_policy":  o.TagPolicy,
			"update_mask": o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateTagPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tag_key": types.StringType,
			"tag_policy": basetypes.ListType{
				ElemType: TagPolicy_SdkV2{}.Type(ctx),
			},
			"update_mask": types.StringType,
		},
	}
}

// GetTagPolicy returns the value of the TagPolicy field in UpdateTagPolicyRequest_SdkV2 as
// a TagPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateTagPolicyRequest_SdkV2) GetTagPolicy(ctx context.Context) (TagPolicy_SdkV2, bool) {
	var e TagPolicy_SdkV2
	if o.TagPolicy.IsNull() || o.TagPolicy.IsUnknown() {
		return e, false
	}
	var v []TagPolicy_SdkV2
	d := o.TagPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTagPolicy sets the value of the TagPolicy field in UpdateTagPolicyRequest_SdkV2.
func (o *UpdateTagPolicyRequest_SdkV2) SetTagPolicy(ctx context.Context, v TagPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tag_policy"]
	o.TagPolicy = types.ListValueMust(t, vs)
}

type Value_SdkV2 struct {
	Name types.String `tfsdk:"name"`
}

func (newState *Value_SdkV2) SyncFieldsDuringCreateOrUpdate(plan Value_SdkV2) {
}

func (newState *Value_SdkV2) SyncFieldsDuringRead(existingState Value_SdkV2) {
}

func (c Value_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Value_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Value_SdkV2
// only implements ToObjectValue() and Type().
func (o Value_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Value_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}
