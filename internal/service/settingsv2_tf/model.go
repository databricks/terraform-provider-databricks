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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type BooleanMessage struct {
	Value types.Bool `tfsdk:"value"`
}

func (newState *BooleanMessage) SyncFieldsDuringCreateOrUpdate(plan BooleanMessage) {
}

func (newState *BooleanMessage) SyncFieldsDuringRead(existingState BooleanMessage) {
}

func (c BooleanMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a BooleanMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BooleanMessage
// only implements ToObjectValue() and Type().
func (o BooleanMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BooleanMessage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.BoolType,
		},
	}
}

type GetPublicAccountSettingRequest struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPublicAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPublicAccountSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublicAccountSettingRequest
// only implements ToObjectValue() and Type().
func (o GetPublicAccountSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPublicAccountSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetPublicWorkspaceSettingRequest struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPublicWorkspaceSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPublicWorkspaceSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublicWorkspaceSettingRequest
// only implements ToObjectValue() and Type().
func (o GetPublicWorkspaceSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPublicWorkspaceSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type IntegerMessage struct {
	Value types.Int64 `tfsdk:"value"`
}

func (newState *IntegerMessage) SyncFieldsDuringCreateOrUpdate(plan IntegerMessage) {
}

func (newState *IntegerMessage) SyncFieldsDuringRead(existingState IntegerMessage) {
}

func (c IntegerMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a IntegerMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IntegerMessage
// only implements ToObjectValue() and Type().
func (o IntegerMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o IntegerMessage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.Int64Type,
		},
	}
}

type ListAccountSettingsMetadataRequest struct {
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
func (a ListAccountSettingsMetadataRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountSettingsMetadataRequest
// only implements ToObjectValue() and Type().
func (o ListAccountSettingsMetadataRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountSettingsMetadataRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListAccountSettingsMetadataResponse struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of all settings available via public APIs and their metadata
	SettingsMetadata types.List `tfsdk:"settings_metadata"`
}

func (newState *ListAccountSettingsMetadataResponse) SyncFieldsDuringCreateOrUpdate(plan ListAccountSettingsMetadataResponse) {
}

func (newState *ListAccountSettingsMetadataResponse) SyncFieldsDuringRead(existingState ListAccountSettingsMetadataResponse) {
}

func (c ListAccountSettingsMetadataResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListAccountSettingsMetadataResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"settings_metadata": reflect.TypeOf(SettingsMetadata{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountSettingsMetadataResponse
// only implements ToObjectValue() and Type().
func (o ListAccountSettingsMetadataResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":   o.NextPageToken,
			"settings_metadata": o.SettingsMetadata,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountSettingsMetadataResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"settings_metadata": basetypes.ListType{
				ElemType: SettingsMetadata{}.Type(ctx),
			},
		},
	}
}

// GetSettingsMetadata returns the value of the SettingsMetadata field in ListAccountSettingsMetadataResponse as
// a slice of SettingsMetadata values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListAccountSettingsMetadataResponse) GetSettingsMetadata(ctx context.Context) ([]SettingsMetadata, bool) {
	if o.SettingsMetadata.IsNull() || o.SettingsMetadata.IsUnknown() {
		return nil, false
	}
	var v []SettingsMetadata
	d := o.SettingsMetadata.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSettingsMetadata sets the value of the SettingsMetadata field in ListAccountSettingsMetadataResponse.
func (o *ListAccountSettingsMetadataResponse) SetSettingsMetadata(ctx context.Context, v []SettingsMetadata) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["settings_metadata"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SettingsMetadata = types.ListValueMust(t, vs)
}

type ListWorkspaceSettingsMetadataRequest struct {
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
func (a ListWorkspaceSettingsMetadataRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceSettingsMetadataRequest
// only implements ToObjectValue() and Type().
func (o ListWorkspaceSettingsMetadataRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListWorkspaceSettingsMetadataRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListWorkspaceSettingsMetadataResponse struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of all settings available via public APIs and their metadata
	SettingsMetadata types.List `tfsdk:"settings_metadata"`
}

func (newState *ListWorkspaceSettingsMetadataResponse) SyncFieldsDuringCreateOrUpdate(plan ListWorkspaceSettingsMetadataResponse) {
}

func (newState *ListWorkspaceSettingsMetadataResponse) SyncFieldsDuringRead(existingState ListWorkspaceSettingsMetadataResponse) {
}

func (c ListWorkspaceSettingsMetadataResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListWorkspaceSettingsMetadataResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"settings_metadata": reflect.TypeOf(SettingsMetadata{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWorkspaceSettingsMetadataResponse
// only implements ToObjectValue() and Type().
func (o ListWorkspaceSettingsMetadataResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":   o.NextPageToken,
			"settings_metadata": o.SettingsMetadata,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListWorkspaceSettingsMetadataResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"settings_metadata": basetypes.ListType{
				ElemType: SettingsMetadata{}.Type(ctx),
			},
		},
	}
}

// GetSettingsMetadata returns the value of the SettingsMetadata field in ListWorkspaceSettingsMetadataResponse as
// a slice of SettingsMetadata values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListWorkspaceSettingsMetadataResponse) GetSettingsMetadata(ctx context.Context) ([]SettingsMetadata, bool) {
	if o.SettingsMetadata.IsNull() || o.SettingsMetadata.IsUnknown() {
		return nil, false
	}
	var v []SettingsMetadata
	d := o.SettingsMetadata.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSettingsMetadata sets the value of the SettingsMetadata field in ListWorkspaceSettingsMetadataResponse.
func (o *ListWorkspaceSettingsMetadataResponse) SetSettingsMetadata(ctx context.Context, v []SettingsMetadata) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["settings_metadata"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SettingsMetadata = types.ListValueMust(t, vs)
}

type PatchPublicAccountSettingRequest struct {
	Name types.String `tfsdk:"-"`

	Setting types.Object `tfsdk:"setting"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PatchPublicAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PatchPublicAccountSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(Setting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchPublicAccountSettingRequest
// only implements ToObjectValue() and Type().
func (o PatchPublicAccountSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    o.Name,
			"setting": o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PatchPublicAccountSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":    types.StringType,
			"setting": Setting{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in PatchPublicAccountSettingRequest as
// a Setting value.
// If the field is unknown or null, the boolean return value is false.
func (o *PatchPublicAccountSettingRequest) GetSetting(ctx context.Context) (Setting, bool) {
	var e Setting
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []Setting
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in PatchPublicAccountSettingRequest.
func (o *PatchPublicAccountSettingRequest) SetSetting(ctx context.Context, v Setting) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
}

type PatchPublicWorkspaceSettingRequest struct {
	Name types.String `tfsdk:"-"`

	Setting types.Object `tfsdk:"setting"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PatchPublicWorkspaceSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PatchPublicWorkspaceSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(Setting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchPublicWorkspaceSettingRequest
// only implements ToObjectValue() and Type().
func (o PatchPublicWorkspaceSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    o.Name,
			"setting": o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PatchPublicWorkspaceSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":    types.StringType,
			"setting": Setting{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in PatchPublicWorkspaceSettingRequest as
// a Setting value.
// If the field is unknown or null, the boolean return value is false.
func (o *PatchPublicWorkspaceSettingRequest) GetSetting(ctx context.Context) (Setting, bool) {
	var e Setting
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []Setting
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in PatchPublicWorkspaceSettingRequest.
func (o *PatchPublicWorkspaceSettingRequest) SetSetting(ctx context.Context, v Setting) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
}

type Setting struct {
	BooleanVal types.Object `tfsdk:"boolean_val"`

	EffectiveBooleanVal types.Object `tfsdk:"effective_boolean_val"`

	EffectiveIntegerVal types.Object `tfsdk:"effective_integer_val"`

	EffectiveStringVal types.Object `tfsdk:"effective_string_val"`

	IntegerVal types.Object `tfsdk:"integer_val"`
	// Name of the setting.
	Name types.String `tfsdk:"name"`

	StringVal types.Object `tfsdk:"string_val"`
}

func (newState *Setting) SyncFieldsDuringCreateOrUpdate(plan Setting) {
}

func (newState *Setting) SyncFieldsDuringRead(existingState Setting) {
}

func (c Setting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["boolean_val"] = attrs["boolean_val"].SetOptional()
	attrs["effective_boolean_val"] = attrs["effective_boolean_val"].SetComputed()
	attrs["effective_integer_val"] = attrs["effective_integer_val"].SetComputed()
	attrs["effective_string_val"] = attrs["effective_string_val"].SetComputed()
	attrs["integer_val"] = attrs["integer_val"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["string_val"] = attrs["string_val"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Setting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Setting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val":           reflect.TypeOf(BooleanMessage{}),
		"effective_boolean_val": reflect.TypeOf(BooleanMessage{}),
		"effective_integer_val": reflect.TypeOf(IntegerMessage{}),
		"effective_string_val":  reflect.TypeOf(StringMessage{}),
		"integer_val":           reflect.TypeOf(IntegerMessage{}),
		"string_val":            reflect.TypeOf(StringMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Setting
// only implements ToObjectValue() and Type().
func (o Setting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Setting) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"boolean_val":           BooleanMessage{}.Type(ctx),
			"effective_boolean_val": BooleanMessage{}.Type(ctx),
			"effective_integer_val": IntegerMessage{}.Type(ctx),
			"effective_string_val":  StringMessage{}.Type(ctx),
			"integer_val":           IntegerMessage{}.Type(ctx),
			"name":                  types.StringType,
			"string_val":            StringMessage{}.Type(ctx),
		},
	}
}

// GetBooleanVal returns the value of the BooleanVal field in Setting as
// a BooleanMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetBooleanVal(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if o.BooleanVal.IsNull() || o.BooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage
	d := o.BooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetBooleanVal sets the value of the BooleanVal field in Setting.
func (o *Setting) SetBooleanVal(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	o.BooleanVal = vs
}

// GetEffectiveBooleanVal returns the value of the EffectiveBooleanVal field in Setting as
// a BooleanMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetEffectiveBooleanVal(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if o.EffectiveBooleanVal.IsNull() || o.EffectiveBooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage
	d := o.EffectiveBooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetEffectiveBooleanVal sets the value of the EffectiveBooleanVal field in Setting.
func (o *Setting) SetEffectiveBooleanVal(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	o.EffectiveBooleanVal = vs
}

// GetEffectiveIntegerVal returns the value of the EffectiveIntegerVal field in Setting as
// a IntegerMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetEffectiveIntegerVal(ctx context.Context) (IntegerMessage, bool) {
	var e IntegerMessage
	if o.EffectiveIntegerVal.IsNull() || o.EffectiveIntegerVal.IsUnknown() {
		return e, false
	}
	var v []IntegerMessage
	d := o.EffectiveIntegerVal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetEffectiveIntegerVal sets the value of the EffectiveIntegerVal field in Setting.
func (o *Setting) SetEffectiveIntegerVal(ctx context.Context, v IntegerMessage) {
	vs := v.ToObjectValue(ctx)
	o.EffectiveIntegerVal = vs
}

// GetEffectiveStringVal returns the value of the EffectiveStringVal field in Setting as
// a StringMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetEffectiveStringVal(ctx context.Context) (StringMessage, bool) {
	var e StringMessage
	if o.EffectiveStringVal.IsNull() || o.EffectiveStringVal.IsUnknown() {
		return e, false
	}
	var v []StringMessage
	d := o.EffectiveStringVal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetEffectiveStringVal sets the value of the EffectiveStringVal field in Setting.
func (o *Setting) SetEffectiveStringVal(ctx context.Context, v StringMessage) {
	vs := v.ToObjectValue(ctx)
	o.EffectiveStringVal = vs
}

// GetIntegerVal returns the value of the IntegerVal field in Setting as
// a IntegerMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetIntegerVal(ctx context.Context) (IntegerMessage, bool) {
	var e IntegerMessage
	if o.IntegerVal.IsNull() || o.IntegerVal.IsUnknown() {
		return e, false
	}
	var v []IntegerMessage
	d := o.IntegerVal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetIntegerVal sets the value of the IntegerVal field in Setting.
func (o *Setting) SetIntegerVal(ctx context.Context, v IntegerMessage) {
	vs := v.ToObjectValue(ctx)
	o.IntegerVal = vs
}

// GetStringVal returns the value of the StringVal field in Setting as
// a StringMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *Setting) GetStringVal(ctx context.Context) (StringMessage, bool) {
	var e StringMessage
	if o.StringVal.IsNull() || o.StringVal.IsUnknown() {
		return e, false
	}
	var v []StringMessage
	d := o.StringVal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetStringVal sets the value of the StringVal field in Setting.
func (o *Setting) SetStringVal(ctx context.Context, v StringMessage) {
	vs := v.ToObjectValue(ctx)
	o.StringVal = vs
}

type SettingsMetadata struct {
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

func (newState *SettingsMetadata) SyncFieldsDuringCreateOrUpdate(plan SettingsMetadata) {
}

func (newState *SettingsMetadata) SyncFieldsDuringRead(existingState SettingsMetadata) {
}

func (c SettingsMetadata) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SettingsMetadata) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SettingsMetadata
// only implements ToObjectValue() and Type().
func (o SettingsMetadata) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o SettingsMetadata) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"docs_link":   types.StringType,
			"name":        types.StringType,
			"type":        types.StringType,
		},
	}
}

type StringMessage struct {
	// Represents a generic string value.
	Value types.String `tfsdk:"value"`
}

func (newState *StringMessage) SyncFieldsDuringCreateOrUpdate(plan StringMessage) {
}

func (newState *StringMessage) SyncFieldsDuringRead(existingState StringMessage) {
}

func (c StringMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a StringMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StringMessage
// only implements ToObjectValue() and Type().
func (o StringMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StringMessage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.StringType,
		},
	}
}
