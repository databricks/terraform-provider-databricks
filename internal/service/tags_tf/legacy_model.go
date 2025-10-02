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

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateTagPolicyRequest_SdkV2 struct {
	TagPolicy types.List `tfsdk:"tag_policy"`
}

func (to *CreateTagPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateTagPolicyRequest_SdkV2) {
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

func (to *CreateTagPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateTagPolicyRequest_SdkV2) {
	if !from.TagPolicy.IsNull() && !from.TagPolicy.IsUnknown() {
		if toTagPolicy, ok := to.GetTagPolicy(ctx); ok {
			if fromTagPolicy, ok := from.GetTagPolicy(ctx); ok {
				toTagPolicy.SyncFieldsDuringRead(ctx, fromTagPolicy)
				to.SetTagPolicy(ctx, toTagPolicy)
			}
		}
	}
}

func (m CreateTagPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["tag_policy"] = attrs["tag_policy"].SetRequired()
	attrs["tag_policy"] = attrs["tag_policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateTagPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateTagPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_policy": reflect.TypeOf(TagPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTagPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateTagPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_policy": m.TagPolicy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateTagPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateTagPolicyRequest_SdkV2) GetTagPolicy(ctx context.Context) (TagPolicy_SdkV2, bool) {
	var e TagPolicy_SdkV2
	if m.TagPolicy.IsNull() || m.TagPolicy.IsUnknown() {
		return e, false
	}
	var v []TagPolicy_SdkV2
	d := m.TagPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTagPolicy sets the value of the TagPolicy field in CreateTagPolicyRequest_SdkV2.
func (m *CreateTagPolicyRequest_SdkV2) SetTagPolicy(ctx context.Context, v TagPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tag_policy"]
	m.TagPolicy = types.ListValueMust(t, vs)
}

type DeleteTagPolicyRequest_SdkV2 struct {
	TagKey types.String `tfsdk:"-"`
}

func (to *DeleteTagPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteTagPolicyRequest_SdkV2) {
}

func (to *DeleteTagPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteTagPolicyRequest_SdkV2) {
}

func (m DeleteTagPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteTagPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTagPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteTagPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_key": m.TagKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteTagPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tag_key": types.StringType,
		},
	}
}

type GetTagPolicyRequest_SdkV2 struct {
	TagKey types.String `tfsdk:"-"`
}

func (to *GetTagPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetTagPolicyRequest_SdkV2) {
}

func (to *GetTagPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetTagPolicyRequest_SdkV2) {
}

func (m GetTagPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetTagPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTagPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetTagPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_key": m.TagKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetTagPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tag_key": types.StringType,
		},
	}
}

type ListTagPoliciesRequest_SdkV2 struct {
	// The maximum number of results to return in this request. Fewer results
	// may be returned than requested. If unspecified or set to 0, this defaults
	// to 1000. The maximum value is 1000; values above 1000 will be coerced
	// down to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// An optional page token received from a previous list tag policies call.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListTagPoliciesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTagPoliciesRequest_SdkV2) {
}

func (to *ListTagPoliciesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListTagPoliciesRequest_SdkV2) {
}

func (m ListTagPoliciesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListTagPoliciesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTagPoliciesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListTagPoliciesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListTagPoliciesRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ListTagPoliciesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTagPoliciesResponse_SdkV2) {
	if !from.TagPolicies.IsNull() && !from.TagPolicies.IsUnknown() && to.TagPolicies.IsNull() && len(from.TagPolicies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TagPolicies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TagPolicies = from.TagPolicies
	}
}

func (to *ListTagPoliciesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListTagPoliciesResponse_SdkV2) {
	if !from.TagPolicies.IsNull() && !from.TagPolicies.IsUnknown() && to.TagPolicies.IsNull() && len(from.TagPolicies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TagPolicies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TagPolicies = from.TagPolicies
	}
}

func (m ListTagPoliciesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListTagPoliciesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_policies": reflect.TypeOf(TagPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTagPoliciesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListTagPoliciesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"tag_policies":    m.TagPolicies,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListTagPoliciesResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ListTagPoliciesResponse_SdkV2) GetTagPolicies(ctx context.Context) ([]TagPolicy_SdkV2, bool) {
	if m.TagPolicies.IsNull() || m.TagPolicies.IsUnknown() {
		return nil, false
	}
	var v []TagPolicy_SdkV2
	d := m.TagPolicies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTagPolicies sets the value of the TagPolicies field in ListTagPoliciesResponse_SdkV2.
func (m *ListTagPoliciesResponse_SdkV2) SetTagPolicies(ctx context.Context, v []TagPolicy_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tag_policies"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TagPolicies = types.ListValueMust(t, vs)
}

type TagPolicy_SdkV2 struct {
	// Timestamp when the tag policy was created
	CreateTime types.String `tfsdk:"create_time"`

	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`

	TagKey types.String `tfsdk:"tag_key"`
	// Timestamp when the tag policy was last updated
	UpdateTime types.String `tfsdk:"update_time"`

	Values types.List `tfsdk:"values"`
}

func (to *TagPolicy_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TagPolicy_SdkV2) {
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (to *TagPolicy_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TagPolicy_SdkV2) {
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (m TagPolicy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TagPolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"values": reflect.TypeOf(Value_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TagPolicy_SdkV2
// only implements ToObjectValue() and Type().
func (m TagPolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m TagPolicy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": types.StringType,
			"description": types.StringType,
			"id":          types.StringType,
			"tag_key":     types.StringType,
			"update_time": types.StringType,
			"values": basetypes.ListType{
				ElemType: Value_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetValues returns the value of the Values field in TagPolicy_SdkV2 as
// a slice of Value_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *TagPolicy_SdkV2) GetValues(ctx context.Context) ([]Value_SdkV2, bool) {
	if m.Values.IsNull() || m.Values.IsUnknown() {
		return nil, false
	}
	var v []Value_SdkV2
	d := m.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in TagPolicy_SdkV2.
func (m *TagPolicy_SdkV2) SetValues(ctx context.Context, v []Value_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Values = types.ListValueMust(t, vs)
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

func (to *UpdateTagPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateTagPolicyRequest_SdkV2) {
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

func (to *UpdateTagPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateTagPolicyRequest_SdkV2) {
	if !from.TagPolicy.IsNull() && !from.TagPolicy.IsUnknown() {
		if toTagPolicy, ok := to.GetTagPolicy(ctx); ok {
			if fromTagPolicy, ok := from.GetTagPolicy(ctx); ok {
				toTagPolicy.SyncFieldsDuringRead(ctx, fromTagPolicy)
				to.SetTagPolicy(ctx, toTagPolicy)
			}
		}
	}
}

func (m UpdateTagPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["tag_policy"] = attrs["tag_policy"].SetRequired()
	attrs["tag_policy"] = attrs["tag_policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateTagPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_policy": reflect.TypeOf(TagPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateTagPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateTagPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tag_key":     m.TagKey,
			"tag_policy":  m.TagPolicy,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateTagPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *UpdateTagPolicyRequest_SdkV2) GetTagPolicy(ctx context.Context) (TagPolicy_SdkV2, bool) {
	var e TagPolicy_SdkV2
	if m.TagPolicy.IsNull() || m.TagPolicy.IsUnknown() {
		return e, false
	}
	var v []TagPolicy_SdkV2
	d := m.TagPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTagPolicy sets the value of the TagPolicy field in UpdateTagPolicyRequest_SdkV2.
func (m *UpdateTagPolicyRequest_SdkV2) SetTagPolicy(ctx context.Context, v TagPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tag_policy"]
	m.TagPolicy = types.ListValueMust(t, vs)
}

type Value_SdkV2 struct {
	Name types.String `tfsdk:"name"`
}

func (to *Value_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Value_SdkV2) {
}

func (to *Value_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Value_SdkV2) {
}

func (m Value_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Value_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Value_SdkV2
// only implements ToObjectValue() and Type().
func (m Value_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Value_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}
