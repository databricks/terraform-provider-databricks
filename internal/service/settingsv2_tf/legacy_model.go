// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package settingsv2_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type BooleanMessage_SdkV2 struct {
	Value types.Bool `tfsdk:"value"`
}

func (newState *BooleanMessage_SdkV2) SyncFieldsDuringCreateOrUpdate(plan BooleanMessage_SdkV2) {
}

func (newState *BooleanMessage_SdkV2) SyncFieldsDuringRead(existingState BooleanMessage_SdkV2) {
}

func (c BooleanMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BooleanMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BooleanMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BooleanMessage_SdkV2
// only implements ToObjectValue() and Type().
func (o BooleanMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BooleanMessage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.BoolType,
		},
	}
}

type GetPublicAccountSettingRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPublicAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPublicAccountSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublicAccountSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPublicAccountSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPublicAccountSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetPublicWorkspaceSettingRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPublicWorkspaceSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPublicWorkspaceSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublicWorkspaceSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPublicWorkspaceSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPublicWorkspaceSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type IntegerMessage_SdkV2 struct {
	Value types.Int64 `tfsdk:"value"`
}

func (newState *IntegerMessage_SdkV2) SyncFieldsDuringCreateOrUpdate(plan IntegerMessage_SdkV2) {
}

func (newState *IntegerMessage_SdkV2) SyncFieldsDuringRead(existingState IntegerMessage_SdkV2) {
}

func (c IntegerMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in IntegerMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a IntegerMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IntegerMessage_SdkV2
// only implements ToObjectValue() and Type().
func (o IntegerMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o IntegerMessage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.Int64Type,
		},
	}
}

type ListAccountSettingsMetadataRequest_SdkV2 struct {
	// The maximum number of settings to return. The service may return fewer
	// than this value. If unspecified, at most 200 settings will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous
	// `ListAccountSettingsMetadataRequest` call. Provide this to retrieve the
	// subsequent page.
	//
	// When paginating, all other parameters provided to
	// `ListAccountSettingsMetadataRequest` must match the call that provided
	// the page token.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountSettingsMetadataRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountSettingsMetadataRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountSettingsMetadataRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAccountSettingsMetadataRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountSettingsMetadataRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListAccountSettingsMetadataResponse_SdkV2 struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of all settings available via public APIs and their metadata
	SettingsMetadata types.List `tfsdk:"settings_metadata"`
}

func (newState *ListAccountSettingsMetadataResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ListAccountSettingsMetadataResponse_SdkV2) {
}

func (newState *ListAccountSettingsMetadataResponse_SdkV2) SyncFieldsDuringRead(existingState ListAccountSettingsMetadataResponse_SdkV2) {
}

func (c ListAccountSettingsMetadataResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["settings_metadata"] = attrs["settings_metadata"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountSettingsMetadataResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountSettingsMetadataResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"settings_metadata": reflect.TypeOf(SettingsMetadata_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountSettingsMetadataResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListAccountSettingsMetadataResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":   o.NextPageToken,
			"settings_metadata": o.SettingsMetadata,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountSettingsMetadataResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"settings_metadata": basetypes.ListType{
				ElemType: SettingsMetadata_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSettingsMetadata returns the value of the SettingsMetadata field in ListAccountSettingsMetadataResponse_SdkV2 as
// a slice of SettingsMetadata_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAccountSettingsMetadataResponse_SdkV2) GetSettingsMetadata(ctx context.Context) ([]SettingsMetadata_SdkV2, bool) {
	if o.SettingsMetadata.IsNull() || o.SettingsMetadata.IsUnknown() {
		return nil, false
	}
	var v []SettingsMetadata_SdkV2
	d := o.SettingsMetadata.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSettingsMetadata sets the value of the SettingsMetadata field in ListAccountSettingsMetadataResponse_SdkV2.
func (o *ListAccountSettingsMetadataResponse_SdkV2) SetSettingsMetadata(ctx context.Context, v []SettingsMetadata_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["settings_metadata"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SettingsMetadata = types.ListValueMust(t, vs)
}

type ListWorkspaceSettingsMetadataRequest_SdkV2 struct {
	// The maximum number of settings to return. The service may return fewer
	// than this value. If unspecified, at most 200 settings will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous
	// `ListWorkspaceSettingsMetadataRequest` call. Provide this to retrieve the
	// subsequent page.
	//
	// When paginating, all other parameters provided to
	// `ListWorkspaceSettingsMetadataRequest` must match the call that provided
	// the page token.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceSettingsMetadataRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListWorkspaceSettingsMetadataRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceSettingsMetadataRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListWorkspaceSettingsMetadataRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListWorkspaceSettingsMetadataRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListWorkspaceSettingsMetadataResponse_SdkV2 struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of all settings available via public APIs and their metadata
	SettingsMetadata types.List `tfsdk:"settings_metadata"`
}

func (newState *ListWorkspaceSettingsMetadataResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(plan ListWorkspaceSettingsMetadataResponse_SdkV2) {
}

func (newState *ListWorkspaceSettingsMetadataResponse_SdkV2) SyncFieldsDuringRead(existingState ListWorkspaceSettingsMetadataResponse_SdkV2) {
}

func (c ListWorkspaceSettingsMetadataResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["settings_metadata"] = attrs["settings_metadata"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceSettingsMetadataResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListWorkspaceSettingsMetadataResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"settings_metadata": reflect.TypeOf(SettingsMetadata_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceSettingsMetadataResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListWorkspaceSettingsMetadataResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":   o.NextPageToken,
			"settings_metadata": o.SettingsMetadata,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListWorkspaceSettingsMetadataResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"settings_metadata": basetypes.ListType{
				ElemType: SettingsMetadata_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSettingsMetadata returns the value of the SettingsMetadata field in ListWorkspaceSettingsMetadataResponse_SdkV2 as
// a slice of SettingsMetadata_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListWorkspaceSettingsMetadataResponse_SdkV2) GetSettingsMetadata(ctx context.Context) ([]SettingsMetadata_SdkV2, bool) {
	if o.SettingsMetadata.IsNull() || o.SettingsMetadata.IsUnknown() {
		return nil, false
	}
	var v []SettingsMetadata_SdkV2
	d := o.SettingsMetadata.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSettingsMetadata sets the value of the SettingsMetadata field in ListWorkspaceSettingsMetadataResponse_SdkV2.
func (o *ListWorkspaceSettingsMetadataResponse_SdkV2) SetSettingsMetadata(ctx context.Context, v []SettingsMetadata_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["settings_metadata"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SettingsMetadata = types.ListValueMust(t, vs)
}

type PatchPublicAccountSettingRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`

	Setting types.List `tfsdk:"setting"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PatchPublicAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PatchPublicAccountSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(Setting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchPublicAccountSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o PatchPublicAccountSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    o.Name,
			"setting": o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PatchPublicAccountSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"setting": basetypes.ListType{
				ElemType: Setting_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in PatchPublicAccountSettingRequest_SdkV2 as
// a Setting_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PatchPublicAccountSettingRequest_SdkV2) GetSetting(ctx context.Context) (Setting_SdkV2, bool) {
	var e Setting_SdkV2
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []Setting_SdkV2
	d := o.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in PatchPublicAccountSettingRequest_SdkV2.
func (o *PatchPublicAccountSettingRequest_SdkV2) SetSetting(ctx context.Context, v Setting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	o.Setting = types.ListValueMust(t, vs)
}

type PatchPublicWorkspaceSettingRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`

	Setting types.List `tfsdk:"setting"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PatchPublicWorkspaceSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PatchPublicWorkspaceSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(Setting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchPublicWorkspaceSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o PatchPublicWorkspaceSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    o.Name,
			"setting": o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PatchPublicWorkspaceSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"setting": basetypes.ListType{
				ElemType: Setting_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in PatchPublicWorkspaceSettingRequest_SdkV2 as
// a Setting_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PatchPublicWorkspaceSettingRequest_SdkV2) GetSetting(ctx context.Context) (Setting_SdkV2, bool) {
	var e Setting_SdkV2
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []Setting_SdkV2
	d := o.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in PatchPublicWorkspaceSettingRequest_SdkV2.
func (o *PatchPublicWorkspaceSettingRequest_SdkV2) SetSetting(ctx context.Context, v Setting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	o.Setting = types.ListValueMust(t, vs)
}

type Setting_SdkV2 struct {
	BooleanVal types.List `tfsdk:"boolean_val"`

	EffectiveBooleanVal types.List `tfsdk:"effective_boolean_val"`

	EffectiveIntegerVal types.List `tfsdk:"effective_integer_val"`

	EffectiveStringVal types.List `tfsdk:"effective_string_val"`

	IntegerVal types.List `tfsdk:"integer_val"`
	// Name of the setting.
	Name types.String `tfsdk:"name"`

	StringVal types.List `tfsdk:"string_val"`
}

func (newState *Setting_SdkV2) SyncFieldsDuringCreateOrUpdate(plan Setting_SdkV2) {
}

func (newState *Setting_SdkV2) SyncFieldsDuringRead(existingState Setting_SdkV2) {
}

func (c Setting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["boolean_val"] = attrs["boolean_val"].SetOptional()
	attrs["boolean_val"] = attrs["boolean_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["effective_boolean_val"] = attrs["effective_boolean_val"].SetComputed()
	attrs["effective_boolean_val"] = attrs["effective_boolean_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["effective_integer_val"] = attrs["effective_integer_val"].SetComputed()
	attrs["effective_integer_val"] = attrs["effective_integer_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["effective_string_val"] = attrs["effective_string_val"].SetComputed()
	attrs["effective_string_val"] = attrs["effective_string_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["integer_val"] = attrs["integer_val"].SetOptional()
	attrs["integer_val"] = attrs["integer_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetOptional()
	attrs["string_val"] = attrs["string_val"].SetOptional()
	attrs["string_val"] = attrs["string_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Setting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Setting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val":           reflect.TypeOf(BooleanMessage_SdkV2{}),
		"effective_boolean_val": reflect.TypeOf(BooleanMessage_SdkV2{}),
		"effective_integer_val": reflect.TypeOf(IntegerMessage_SdkV2{}),
		"effective_string_val":  reflect.TypeOf(StringMessage_SdkV2{}),
		"integer_val":           reflect.TypeOf(IntegerMessage_SdkV2{}),
		"string_val":            reflect.TypeOf(StringMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Setting_SdkV2
// only implements ToObjectValue() and Type().
func (o Setting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":           o.BooleanVal,
			"effective_boolean_val": o.EffectiveBooleanVal,
			"effective_integer_val": o.EffectiveIntegerVal,
			"effective_string_val":  o.EffectiveStringVal,
			"integer_val":           o.IntegerVal,
			"name":                  o.Name,
			"string_val":            o.StringVal,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Setting_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"boolean_val": basetypes.ListType{
				ElemType: BooleanMessage_SdkV2{}.Type(ctx),
			},
			"effective_boolean_val": basetypes.ListType{
				ElemType: BooleanMessage_SdkV2{}.Type(ctx),
			},
			"effective_integer_val": basetypes.ListType{
				ElemType: IntegerMessage_SdkV2{}.Type(ctx),
			},
			"effective_string_val": basetypes.ListType{
				ElemType: StringMessage_SdkV2{}.Type(ctx),
			},
			"integer_val": basetypes.ListType{
				ElemType: IntegerMessage_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
			"string_val": basetypes.ListType{
				ElemType: StringMessage_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetBooleanVal returns the value of the BooleanVal field in Setting_SdkV2 as
// a BooleanMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetBooleanVal(ctx context.Context) (BooleanMessage_SdkV2, bool) {
	var e BooleanMessage_SdkV2
	if o.BooleanVal.IsNull() || o.BooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage_SdkV2
	d := o.BooleanVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBooleanVal sets the value of the BooleanVal field in Setting_SdkV2.
func (o *Setting_SdkV2) SetBooleanVal(ctx context.Context, v BooleanMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["boolean_val"]
	o.BooleanVal = types.ListValueMust(t, vs)
}

// GetEffectiveBooleanVal returns the value of the EffectiveBooleanVal field in Setting_SdkV2 as
// a BooleanMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetEffectiveBooleanVal(ctx context.Context) (BooleanMessage_SdkV2, bool) {
	var e BooleanMessage_SdkV2
	if o.EffectiveBooleanVal.IsNull() || o.EffectiveBooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage_SdkV2
	d := o.EffectiveBooleanVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectiveBooleanVal sets the value of the EffectiveBooleanVal field in Setting_SdkV2.
func (o *Setting_SdkV2) SetEffectiveBooleanVal(ctx context.Context, v BooleanMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_boolean_val"]
	o.EffectiveBooleanVal = types.ListValueMust(t, vs)
}

// GetEffectiveIntegerVal returns the value of the EffectiveIntegerVal field in Setting_SdkV2 as
// a IntegerMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetEffectiveIntegerVal(ctx context.Context) (IntegerMessage_SdkV2, bool) {
	var e IntegerMessage_SdkV2
	if o.EffectiveIntegerVal.IsNull() || o.EffectiveIntegerVal.IsUnknown() {
		return e, false
	}
	var v []IntegerMessage_SdkV2
	d := o.EffectiveIntegerVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectiveIntegerVal sets the value of the EffectiveIntegerVal field in Setting_SdkV2.
func (o *Setting_SdkV2) SetEffectiveIntegerVal(ctx context.Context, v IntegerMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_integer_val"]
	o.EffectiveIntegerVal = types.ListValueMust(t, vs)
}

// GetEffectiveStringVal returns the value of the EffectiveStringVal field in Setting_SdkV2 as
// a StringMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetEffectiveStringVal(ctx context.Context) (StringMessage_SdkV2, bool) {
	var e StringMessage_SdkV2
	if o.EffectiveStringVal.IsNull() || o.EffectiveStringVal.IsUnknown() {
		return e, false
	}
	var v []StringMessage_SdkV2
	d := o.EffectiveStringVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectiveStringVal sets the value of the EffectiveStringVal field in Setting_SdkV2.
func (o *Setting_SdkV2) SetEffectiveStringVal(ctx context.Context, v StringMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_string_val"]
	o.EffectiveStringVal = types.ListValueMust(t, vs)
}

// GetIntegerVal returns the value of the IntegerVal field in Setting_SdkV2 as
// a IntegerMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetIntegerVal(ctx context.Context) (IntegerMessage_SdkV2, bool) {
	var e IntegerMessage_SdkV2
	if o.IntegerVal.IsNull() || o.IntegerVal.IsUnknown() {
		return e, false
	}
	var v []IntegerMessage_SdkV2
	d := o.IntegerVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIntegerVal sets the value of the IntegerVal field in Setting_SdkV2.
func (o *Setting_SdkV2) SetIntegerVal(ctx context.Context, v IntegerMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["integer_val"]
	o.IntegerVal = types.ListValueMust(t, vs)
}

// GetStringVal returns the value of the StringVal field in Setting_SdkV2 as
// a StringMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting_SdkV2) GetStringVal(ctx context.Context) (StringMessage_SdkV2, bool) {
	var e StringMessage_SdkV2
	if o.StringVal.IsNull() || o.StringVal.IsUnknown() {
		return e, false
	}
	var v []StringMessage_SdkV2
	d := o.StringVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStringVal sets the value of the StringVal field in Setting_SdkV2.
func (o *Setting_SdkV2) SetStringVal(ctx context.Context, v StringMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["string_val"]
	o.StringVal = types.ListValueMust(t, vs)
}

type SettingsMetadata_SdkV2 struct {
	// Setting description for what this setting controls
	Description types.String `tfsdk:"description"`
	// Link to databricks documentation for the setting
	DocsLink types.String `tfsdk:"docs_link"`
	// Name of the setting.
	Name types.String `tfsdk:"name"`
	// Type of the setting. To set this setting, the value sent must match this
	// type.
	Type_ types.String `tfsdk:"type"`
}

func (newState *SettingsMetadata_SdkV2) SyncFieldsDuringCreateOrUpdate(plan SettingsMetadata_SdkV2) {
}

func (newState *SettingsMetadata_SdkV2) SyncFieldsDuringRead(existingState SettingsMetadata_SdkV2) {
}

func (c SettingsMetadata_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["docs_link"] = attrs["docs_link"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["type"] = attrs["type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SettingsMetadata.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SettingsMetadata_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SettingsMetadata_SdkV2
// only implements ToObjectValue() and Type().
func (o SettingsMetadata_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": o.Description,
			"docs_link":   o.DocsLink,
			"name":        o.Name,
			"type":        o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SettingsMetadata_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"docs_link":   types.StringType,
			"name":        types.StringType,
			"type":        types.StringType,
		},
	}
}

type StringMessage_SdkV2 struct {
	// Represents a generic string value.
	Value types.String `tfsdk:"value"`
}

func (newState *StringMessage_SdkV2) SyncFieldsDuringCreateOrUpdate(plan StringMessage_SdkV2) {
}

func (newState *StringMessage_SdkV2) SyncFieldsDuringRead(existingState StringMessage_SdkV2) {
}

func (c StringMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StringMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StringMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StringMessage_SdkV2
// only implements ToObjectValue() and Type().
func (o StringMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StringMessage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.StringType,
		},
	}
}
