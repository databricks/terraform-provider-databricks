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

func (c CreateTagPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	var v TagPolicy
	d := o.TagPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTagPolicy sets the value of the TagPolicy field in CreateTagPolicyRequest.
func (o *CreateTagPolicyRequest) SetTagPolicy(ctx context.Context, v TagPolicy) {
	vs := v.ToObjectValue(ctx)
	o.TagPolicy = vs
}

type DeleteTagPolicyRequest struct {
	TagKey types.String `tfsdk:"-"`
}

func (to *DeleteTagPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteTagPolicyRequest) {
}

func (to *DeleteTagPolicyRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteTagPolicyRequest) {
}

func (c DeleteTagPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type GetTagPolicyRequest struct {
	TagKey types.String `tfsdk:"-"`
}

func (to *GetTagPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetTagPolicyRequest) {
}

func (to *GetTagPolicyRequest) SyncFieldsDuringRead(ctx context.Context, from GetTagPolicyRequest) {
}

func (c GetTagPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

func (c ListTagPoliciesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

func (c TagPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"create_time": o.CreateTime,
			"description": o.Description,
			"id":          o.Id,
			"tag_key":     o.TagKey,
			"update_time": o.UpdateTime,
			"values":      o.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TagPolicy) Type(ctx context.Context) attr.Type {
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

func (c UpdateTagPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	var v TagPolicy
	d := o.TagPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTagPolicy sets the value of the TagPolicy field in UpdateTagPolicyRequest.
func (o *UpdateTagPolicyRequest) SetTagPolicy(ctx context.Context, v TagPolicy) {
	vs := v.ToObjectValue(ctx)
	o.TagPolicy = vs
}

type Value struct {
	Name types.String `tfsdk:"name"`
}

func (to *Value) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Value) {
}

func (to *Value) SyncFieldsDuringRead(ctx context.Context, from Value) {
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
