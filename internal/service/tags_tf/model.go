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

type CreateTagAssignmentRequest struct {
	TagAssignment types.Object `tfsdk:"tag_assignment"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateTagAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateTagAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_assignment": reflect.TypeOf(TagAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTagAssignmentRequest
// only implements ToObjectValue() and Type().
func (o CreateTagAssignmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_assignment": o.TagAssignment,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateTagAssignmentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tag_assignment": TagAssignment{}.Type(ctx),
		},
	}
}

// GetTagAssignment returns the value of the TagAssignment field in CreateTagAssignmentRequest as
// a TagAssignment value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateTagAssignmentRequest) GetTagAssignment(ctx context.Context) (TagAssignment, bool) {
	var e TagAssignment
	if o.TagAssignment.IsNull() || o.TagAssignment.IsUnknown() {
		return e, false
	}
	var v []TagAssignment
	d := o.TagAssignment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTagAssignment sets the value of the TagAssignment field in CreateTagAssignmentRequest.
func (o *CreateTagAssignmentRequest) SetTagAssignment(ctx context.Context, v TagAssignment) {
	vs := v.ToObjectValue(ctx)
	o.TagAssignment = vs
}

type CreateTagPolicyRequest struct {
	TagPolicy types.Object `tfsdk:"tag_policy"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateTagPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateTagPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_policy": reflect.TypeOf(TagPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTagPolicyRequest
// only implements ToObjectValue() and Type().
func (o CreateTagPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_policy": o.TagPolicy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateTagPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tag_policy": TagPolicy{}.Type(ctx),
		},
	}
}

// GetTagPolicy returns the value of the TagPolicy field in CreateTagPolicyRequest as
// a TagPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateTagPolicyRequest) GetTagPolicy(ctx context.Context) (TagPolicy, bool) {
	var e TagPolicy
	if o.TagPolicy.IsNull() || o.TagPolicy.IsUnknown() {
		return e, false
	}
	var v []TagPolicy
	d := o.TagPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTagPolicy sets the value of the TagPolicy field in CreateTagPolicyRequest.
func (o *CreateTagPolicyRequest) SetTagPolicy(ctx context.Context, v TagPolicy) {
	vs := v.ToObjectValue(ctx)
	o.TagPolicy = vs
}

type DeleteTagAssignmentRequest struct {
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
func (a DeleteTagAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTagAssignmentRequest
// only implements ToObjectValue() and Type().
func (o DeleteTagAssignmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_id":   o.EntityId,
			"entity_type": o.EntityType,
			"tag_key":     o.TagKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteTagAssignmentRequest) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTagPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteTagPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTagPolicyRequest
// only implements ToObjectValue() and Type().
func (o DeleteTagPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_key": o.TagKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteTagPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tag_key": types.StringType,
		},
	}
}

type GetTagAssignmentRequest struct {
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
func (a GetTagAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTagAssignmentRequest
// only implements ToObjectValue() and Type().
func (o GetTagAssignmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_id":   o.EntityId,
			"entity_type": o.EntityType,
			"tag_key":     o.TagKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetTagAssignmentRequest) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetTagPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetTagPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTagPolicyRequest
// only implements ToObjectValue() and Type().
func (o GetTagPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_key": o.TagKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetTagPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tag_key": types.StringType,
		},
	}
}

type ListTagAssignmentsRequest struct {
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
func (a ListTagAssignmentsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTagAssignmentsRequest
// only implements ToObjectValue() and Type().
func (o ListTagAssignmentsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListTagAssignmentsRequest) Type(ctx context.Context) attr.Type {
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
	NextPageToken types.String `tfsdk:"next_page_token"`

	TagAssignments types.List `tfsdk:"tag_assignments"`
}

func (newState *ListTagAssignmentsResponse) SyncFieldsDuringCreateOrUpdate(plan ListTagAssignmentsResponse) {
}

func (newState *ListTagAssignmentsResponse) SyncFieldsDuringRead(existingState ListTagAssignmentsResponse) {
}

func (c ListTagAssignmentsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListTagAssignmentsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_assignments": reflect.TypeOf(TagAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTagAssignmentsResponse
// only implements ToObjectValue() and Type().
func (o ListTagAssignmentsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"tag_assignments": o.TagAssignments,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListTagAssignmentsResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListTagAssignmentsResponse) GetTagAssignments(ctx context.Context) ([]TagAssignment, bool) {
	if o.TagAssignments.IsNull() || o.TagAssignments.IsUnknown() {
		return nil, false
	}
	var v []TagAssignment
	d := o.TagAssignments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTagAssignments sets the value of the TagAssignments field in ListTagAssignmentsResponse.
func (o *ListTagAssignmentsResponse) SetTagAssignments(ctx context.Context, v []TagAssignment) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tag_assignments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TagAssignments = types.ListValueMust(t, vs)
}

type ListTagPoliciesRequest struct {
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
func (a ListTagPoliciesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTagPoliciesRequest
// only implements ToObjectValue() and Type().
func (o ListTagPoliciesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListTagPoliciesRequest) Type(ctx context.Context) attr.Type {
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

func (newState *ListTagPoliciesResponse) SyncFieldsDuringCreateOrUpdate(plan ListTagPoliciesResponse) {
}

func (newState *ListTagPoliciesResponse) SyncFieldsDuringRead(existingState ListTagPoliciesResponse) {
}

func (c ListTagPoliciesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListTagPoliciesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_policies": reflect.TypeOf(TagPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTagPoliciesResponse
// only implements ToObjectValue() and Type().
func (o ListTagPoliciesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"tag_policies":    o.TagPolicies,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListTagPoliciesResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListTagPoliciesResponse) GetTagPolicies(ctx context.Context) ([]TagPolicy, bool) {
	if o.TagPolicies.IsNull() || o.TagPolicies.IsUnknown() {
		return nil, false
	}
	var v []TagPolicy
	d := o.TagPolicies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTagPolicies sets the value of the TagPolicies field in ListTagPoliciesResponse.
func (o *ListTagPoliciesResponse) SetTagPolicies(ctx context.Context, v []TagPolicy) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tag_policies"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TagPolicies = types.ListValueMust(t, vs)
}

type TagAssignment struct {
	EntityId types.String `tfsdk:"entity_id"`

	EntityType types.String `tfsdk:"entity_type"`

	TagKey types.String `tfsdk:"tag_key"`

	TagValue types.String `tfsdk:"tag_value"`
}

func (newState *TagAssignment) SyncFieldsDuringCreateOrUpdate(plan TagAssignment) {
}

func (newState *TagAssignment) SyncFieldsDuringRead(existingState TagAssignment) {
}

func (c TagAssignment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TagAssignment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TagAssignment
// only implements ToObjectValue() and Type().
func (o TagAssignment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o TagAssignment) Type(ctx context.Context) attr.Type {
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
	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`

	TagKey types.String `tfsdk:"tag_key"`

	Values types.List `tfsdk:"values"`
}

func (newState *TagPolicy) SyncFieldsDuringCreateOrUpdate(plan TagPolicy) {
}

func (newState *TagPolicy) SyncFieldsDuringRead(existingState TagPolicy) {
}

func (c TagPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TagPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"values": reflect.TypeOf(Value{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TagPolicy
// only implements ToObjectValue() and Type().
func (o TagPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o TagPolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"id":          types.StringType,
			"tag_key":     types.StringType,
			"values": basetypes.ListType{
				ElemType: Value{}.Type(ctx),
			},
		},
	}
}

// GetValues returns the value of the Values field in TagPolicy as
// a slice of Value values.
// If the field is unknown or null, the boolean return value is false.
func (o *TagPolicy) GetValues(ctx context.Context) ([]Value, bool) {
	if o.Values.IsNull() || o.Values.IsUnknown() {
		return nil, false
	}
	var v []Value
	d := o.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in TagPolicy.
func (o *TagPolicy) SetValues(ctx context.Context, v []Value) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Values = types.ListValueMust(t, vs)
}

type UpdateTagAssignmentRequest struct {
	EntityId types.String `tfsdk:"-"`

	EntityType types.String `tfsdk:"-"`

	TagAssignment types.Object `tfsdk:"tag_assignment"`

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
func (a UpdateTagAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_assignment": reflect.TypeOf(TagAssignment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateTagAssignmentRequest
// only implements ToObjectValue() and Type().
func (o UpdateTagAssignmentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateTagAssignmentRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateTagAssignmentRequest) GetTagAssignment(ctx context.Context) (TagAssignment, bool) {
	var e TagAssignment
	if o.TagAssignment.IsNull() || o.TagAssignment.IsUnknown() {
		return e, false
	}
	var v []TagAssignment
	d := o.TagAssignment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTagAssignment sets the value of the TagAssignment field in UpdateTagAssignmentRequest.
func (o *UpdateTagAssignmentRequest) SetTagAssignment(ctx context.Context, v TagAssignment) {
	vs := v.ToObjectValue(ctx)
	o.TagAssignment = vs
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateTagPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateTagPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_policy": reflect.TypeOf(TagPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateTagPolicyRequest
// only implements ToObjectValue() and Type().
func (o UpdateTagPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_key":     o.TagKey,
			"tag_policy":  o.TagPolicy,
			"update_mask": o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateTagPolicyRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateTagPolicyRequest) GetTagPolicy(ctx context.Context) (TagPolicy, bool) {
	var e TagPolicy
	if o.TagPolicy.IsNull() || o.TagPolicy.IsUnknown() {
		return e, false
	}
	var v []TagPolicy
	d := o.TagPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTagPolicy sets the value of the TagPolicy field in UpdateTagPolicyRequest.
func (o *UpdateTagPolicyRequest) SetTagPolicy(ctx context.Context, v TagPolicy) {
	vs := v.ToObjectValue(ctx)
	o.TagPolicy = vs
}

type Value struct {
	Name types.String `tfsdk:"name"`
}

func (newState *Value) SyncFieldsDuringCreateOrUpdate(plan Value) {
}

func (newState *Value) SyncFieldsDuringRead(existingState Value) {
}

func (c Value) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Value) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Value
// only implements ToObjectValue() and Type().
func (o Value) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Value) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}
