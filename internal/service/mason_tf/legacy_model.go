// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package mason_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Request to append items to a session.
type AppendSessionItemsRequest_SdkV2 struct {
	// Items to append atomically in request order. Concurrent append requests
	// are serialized into one committed order without exposing a numeric
	// sequence in the public contract.
	Items types.List `tfsdk:"items"`
	// Resource name of the containing session, in the form
	// `session-stores/{session_store_id}/sessions/{session_id}`.
	Parent types.String `tfsdk:"-"`
}

func (to *AppendSessionItemsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppendSessionItemsRequest_SdkV2) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() {
		if toItems, ok := to.GetItems(ctx); ok {
			if fromItems, ok := from.GetItems(ctx); ok {
				// Recursively sync the fields of each Items element by position.
				for i := range toItems {
					if i < len(fromItems) {
						toItems[i].SyncFieldsDuringCreateOrUpdate(ctx, fromItems[i])
					}
				}
				to.SetItems(ctx, toItems)
			}
		}
	}
}

func (to *AppendSessionItemsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AppendSessionItemsRequest_SdkV2) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() {
		if toItems, ok := to.GetItems(ctx); ok {
			if fromItems, ok := from.GetItems(ctx); ok {
				for i := range toItems {
					if i < len(fromItems) {
						toItems[i].SyncFieldsDuringRead(ctx, fromItems[i])
					}
				}
				to.SetItems(ctx, toItems)
			}
		}
	}
}

func (m AppendSessionItemsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["items"] = attrs["items"].SetRequired()
	attrs["parent"] = attrs["parent"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppendSessionItemsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppendSessionItemsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(SessionItem_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppendSessionItemsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m AppendSessionItemsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"items":  m.Items,
			"parent": m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppendSessionItemsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"items": basetypes.ListType{
				ElemType: SessionItem_SdkV2{}.Type(ctx),
			},
			"parent": types.StringType,
		},
	}
}

// GetItems returns the value of the Items field in AppendSessionItemsRequest_SdkV2 as
// a slice of SessionItem_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *AppendSessionItemsRequest_SdkV2) GetItems(ctx context.Context) ([]SessionItem_SdkV2, bool) {
	if m.Items.IsNull() || m.Items.IsUnknown() {
		return nil, false
	}
	var v []SessionItem_SdkV2
	d := m.Items.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetItems sets the value of the Items field in AppendSessionItemsRequest_SdkV2.
func (m *AppendSessionItemsRequest_SdkV2) SetItems(ctx context.Context, v []SessionItem_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["items"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Items = types.ListValueMust(t, vs)
}

// Response containing appended items.
type AppendSessionItemsResponse_SdkV2 struct {
	// Persisted session items with service-assigned fields.
	SessionItems types.List `tfsdk:"session_items"`
}

func (to *AppendSessionItemsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppendSessionItemsResponse_SdkV2) {
	if !from.SessionItems.IsNull() && !from.SessionItems.IsUnknown() && to.SessionItems.IsNull() && len(from.SessionItems.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SessionItems, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SessionItems = from.SessionItems
	}
	if !from.SessionItems.IsNull() && !from.SessionItems.IsUnknown() {
		if toSessionItems, ok := to.GetSessionItems(ctx); ok {
			if fromSessionItems, ok := from.GetSessionItems(ctx); ok {
				// Recursively sync the fields of each SessionItems element by position.
				for i := range toSessionItems {
					if i < len(fromSessionItems) {
						toSessionItems[i].SyncFieldsDuringCreateOrUpdate(ctx, fromSessionItems[i])
					}
				}
				to.SetSessionItems(ctx, toSessionItems)
			}
		}
	}
}

func (to *AppendSessionItemsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AppendSessionItemsResponse_SdkV2) {
	if !from.SessionItems.IsNull() && !from.SessionItems.IsUnknown() && to.SessionItems.IsNull() && len(from.SessionItems.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SessionItems, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SessionItems = from.SessionItems
	}
	if !from.SessionItems.IsNull() && !from.SessionItems.IsUnknown() {
		if toSessionItems, ok := to.GetSessionItems(ctx); ok {
			if fromSessionItems, ok := from.GetSessionItems(ctx); ok {
				for i := range toSessionItems {
					if i < len(fromSessionItems) {
						toSessionItems[i].SyncFieldsDuringRead(ctx, fromSessionItems[i])
					}
				}
				to.SetSessionItems(ctx, toSessionItems)
			}
		}
	}
}

func (m AppendSessionItemsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["session_items"] = attrs["session_items"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppendSessionItemsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppendSessionItemsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"session_items": reflect.TypeOf(SessionItem_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppendSessionItemsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m AppendSessionItemsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"session_items": m.SessionItems,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppendSessionItemsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"session_items": basetypes.ListType{
				ElemType: SessionItem_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSessionItems returns the value of the SessionItems field in AppendSessionItemsResponse_SdkV2 as
// a slice of SessionItem_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *AppendSessionItemsResponse_SdkV2) GetSessionItems(ctx context.Context) ([]SessionItem_SdkV2, bool) {
	if m.SessionItems.IsNull() || m.SessionItems.IsUnknown() {
		return nil, false
	}
	var v []SessionItem_SdkV2
	d := m.SessionItems.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSessionItems sets the value of the SessionItems field in AppendSessionItemsResponse_SdkV2.
func (m *AppendSessionItemsResponse_SdkV2) SetSessionItems(ctx context.Context, v []SessionItem_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["session_items"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SessionItems = types.ListValueMust(t, vs)
}

// Request to clear all items from a session.
type ClearSessionItemsRequest_SdkV2 struct {
	// Resource name of the containing session, in the form
	// `session-stores/{session_store_id}/sessions/{session_id}`.
	Parent types.String `tfsdk:"-"`
}

func (to *ClearSessionItemsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClearSessionItemsRequest_SdkV2) {
}

func (to *ClearSessionItemsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ClearSessionItemsRequest_SdkV2) {
}

func (m ClearSessionItemsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClearSessionItemsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClearSessionItemsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClearSessionItemsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ClearSessionItemsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parent": m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClearSessionItemsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent": types.StringType,
		},
	}
}

// Response from clearing items from a session.
type ClearSessionItemsResponse_SdkV2 struct {
}

func (to *ClearSessionItemsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClearSessionItemsResponse_SdkV2) {
}

func (to *ClearSessionItemsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ClearSessionItemsResponse_SdkV2) {
}

func (m ClearSessionItemsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClearSessionItemsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClearSessionItemsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClearSessionItemsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ClearSessionItemsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ClearSessionItemsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type CreateManagedMemoryEntryRequest_SdkV2 struct {
	// The managed memory entry to create.
	ManagedMemoryEntry types.List `tfsdk:"managed_memory_entry"`
	// Optional caller-selected managed memory entry ID. The service generates
	// an ID when omitted.
	ManagedMemoryEntryId types.String `tfsdk:"-"`
	// Managed memory store that will contain the entry, in the form
	// `memory-stores/{managed_memory_store_id}`.
	Parent types.String `tfsdk:"-"`
}

func (to *CreateManagedMemoryEntryRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateManagedMemoryEntryRequest_SdkV2) {
	if !from.ManagedMemoryEntry.IsNull() && !from.ManagedMemoryEntry.IsUnknown() {
		if toManagedMemoryEntry, ok := to.GetManagedMemoryEntry(ctx); ok {
			if fromManagedMemoryEntry, ok := from.GetManagedMemoryEntry(ctx); ok {
				// Recursively sync the fields of ManagedMemoryEntry
				toManagedMemoryEntry.SyncFieldsDuringCreateOrUpdate(ctx, fromManagedMemoryEntry)
				to.SetManagedMemoryEntry(ctx, toManagedMemoryEntry)
			}
		}
	}
}

func (to *CreateManagedMemoryEntryRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateManagedMemoryEntryRequest_SdkV2) {
	if !from.ManagedMemoryEntry.IsNull() && !from.ManagedMemoryEntry.IsUnknown() {
		if toManagedMemoryEntry, ok := to.GetManagedMemoryEntry(ctx); ok {
			if fromManagedMemoryEntry, ok := from.GetManagedMemoryEntry(ctx); ok {
				toManagedMemoryEntry.SyncFieldsDuringRead(ctx, fromManagedMemoryEntry)
				to.SetManagedMemoryEntry(ctx, toManagedMemoryEntry)
			}
		}
	}
}

func (m CreateManagedMemoryEntryRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["managed_memory_entry"] = attrs["managed_memory_entry"].SetRequired()
	attrs["managed_memory_entry"] = attrs["managed_memory_entry"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["managed_memory_entry_id"] = attrs["managed_memory_entry_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateManagedMemoryEntryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateManagedMemoryEntryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"managed_memory_entry": reflect.TypeOf(ManagedMemoryEntry_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateManagedMemoryEntryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateManagedMemoryEntryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"managed_memory_entry":    m.ManagedMemoryEntry,
			"managed_memory_entry_id": m.ManagedMemoryEntryId,
			"parent":                  m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateManagedMemoryEntryRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"managed_memory_entry": basetypes.ListType{
				ElemType: ManagedMemoryEntry_SdkV2{}.Type(ctx),
			},
			"managed_memory_entry_id": types.StringType,
			"parent":                  types.StringType,
		},
	}
}

// GetManagedMemoryEntry returns the value of the ManagedMemoryEntry field in CreateManagedMemoryEntryRequest_SdkV2 as
// a ManagedMemoryEntry_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateManagedMemoryEntryRequest_SdkV2) GetManagedMemoryEntry(ctx context.Context) (ManagedMemoryEntry_SdkV2, bool) {
	var e ManagedMemoryEntry_SdkV2
	if m.ManagedMemoryEntry.IsNull() || m.ManagedMemoryEntry.IsUnknown() {
		return e, false
	}
	var v []ManagedMemoryEntry_SdkV2
	d := m.ManagedMemoryEntry.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetManagedMemoryEntry sets the value of the ManagedMemoryEntry field in CreateManagedMemoryEntryRequest_SdkV2.
func (m *CreateManagedMemoryEntryRequest_SdkV2) SetManagedMemoryEntry(ctx context.Context, v ManagedMemoryEntry_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["managed_memory_entry"]
	m.ManagedMemoryEntry = types.ListValueMust(t, vs)
}

type CreateManagedMemoryStoreRequest_SdkV2 struct {
	// The managed memory store to create.
	ManagedMemoryStore types.List `tfsdk:"managed_memory_store"`
	// Caller-provided, workspace-unique managed memory store ID. It must be
	// 3-56 characters, begin with a lowercase letter, contain only lowercase
	// letters, digits, and hyphens, and end with a letter or digit.
	ManagedMemoryStoreId types.String `tfsdk:"-"`
}

func (to *CreateManagedMemoryStoreRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateManagedMemoryStoreRequest_SdkV2) {
	if !from.ManagedMemoryStore.IsNull() && !from.ManagedMemoryStore.IsUnknown() {
		if toManagedMemoryStore, ok := to.GetManagedMemoryStore(ctx); ok {
			if fromManagedMemoryStore, ok := from.GetManagedMemoryStore(ctx); ok {
				// Recursively sync the fields of ManagedMemoryStore
				toManagedMemoryStore.SyncFieldsDuringCreateOrUpdate(ctx, fromManagedMemoryStore)
				to.SetManagedMemoryStore(ctx, toManagedMemoryStore)
			}
		}
	}
}

func (to *CreateManagedMemoryStoreRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateManagedMemoryStoreRequest_SdkV2) {
	if !from.ManagedMemoryStore.IsNull() && !from.ManagedMemoryStore.IsUnknown() {
		if toManagedMemoryStore, ok := to.GetManagedMemoryStore(ctx); ok {
			if fromManagedMemoryStore, ok := from.GetManagedMemoryStore(ctx); ok {
				toManagedMemoryStore.SyncFieldsDuringRead(ctx, fromManagedMemoryStore)
				to.SetManagedMemoryStore(ctx, toManagedMemoryStore)
			}
		}
	}
}

func (m CreateManagedMemoryStoreRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["managed_memory_store"] = attrs["managed_memory_store"].SetRequired()
	attrs["managed_memory_store"] = attrs["managed_memory_store"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["managed_memory_store_id"] = attrs["managed_memory_store_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateManagedMemoryStoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateManagedMemoryStoreRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"managed_memory_store": reflect.TypeOf(ManagedMemoryStore_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateManagedMemoryStoreRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateManagedMemoryStoreRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"managed_memory_store":    m.ManagedMemoryStore,
			"managed_memory_store_id": m.ManagedMemoryStoreId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateManagedMemoryStoreRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"managed_memory_store": basetypes.ListType{
				ElemType: ManagedMemoryStore_SdkV2{}.Type(ctx),
			},
			"managed_memory_store_id": types.StringType,
		},
	}
}

// GetManagedMemoryStore returns the value of the ManagedMemoryStore field in CreateManagedMemoryStoreRequest_SdkV2 as
// a ManagedMemoryStore_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateManagedMemoryStoreRequest_SdkV2) GetManagedMemoryStore(ctx context.Context) (ManagedMemoryStore_SdkV2, bool) {
	var e ManagedMemoryStore_SdkV2
	if m.ManagedMemoryStore.IsNull() || m.ManagedMemoryStore.IsUnknown() {
		return e, false
	}
	var v []ManagedMemoryStore_SdkV2
	d := m.ManagedMemoryStore.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetManagedMemoryStore sets the value of the ManagedMemoryStore field in CreateManagedMemoryStoreRequest_SdkV2.
func (m *CreateManagedMemoryStoreRequest_SdkV2) SetManagedMemoryStore(ctx context.Context, v ManagedMemoryStore_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["managed_memory_store"]
	m.ManagedMemoryStore = types.ListValueMust(t, vs)
}

type CreateSessionRequest_SdkV2 struct {
	// Resource name of the containing session store, in the form
	// `session-stores/{session_store_id}`.
	Parent types.String `tfsdk:"-"`
	// The session to create. `actor_id` is required. A session with
	// `parent_session_id` is a child and must use its parent's `actor_id`.
	// Independent forks are created only through `ForkSession`.
	Session types.List `tfsdk:"session"`
	// Optional caller-selected session ID. The service generates a UUID when
	// this field is omitted. The ID must be unique; a collision returns
	// `ALREADY_EXISTS`.
	SessionId types.String `tfsdk:"-"`
}

func (to *CreateSessionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateSessionRequest_SdkV2) {
	if !from.Session.IsNull() && !from.Session.IsUnknown() {
		if toSession, ok := to.GetSession(ctx); ok {
			if fromSession, ok := from.GetSession(ctx); ok {
				// Recursively sync the fields of Session
				toSession.SyncFieldsDuringCreateOrUpdate(ctx, fromSession)
				to.SetSession(ctx, toSession)
			}
		}
	}
}

func (to *CreateSessionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateSessionRequest_SdkV2) {
	if !from.Session.IsNull() && !from.Session.IsUnknown() {
		if toSession, ok := to.GetSession(ctx); ok {
			if fromSession, ok := from.GetSession(ctx); ok {
				toSession.SyncFieldsDuringRead(ctx, fromSession)
				to.SetSession(ctx, toSession)
			}
		}
	}
}

func (m CreateSessionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["session"] = attrs["session"].SetRequired()
	attrs["session"] = attrs["session"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["session_id"] = attrs["session_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateSessionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateSessionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"session": reflect.TypeOf(Session_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateSessionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateSessionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parent":     m.Parent,
			"session":    m.Session,
			"session_id": m.SessionId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateSessionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent": types.StringType,
			"session": basetypes.ListType{
				ElemType: Session_SdkV2{}.Type(ctx),
			},
			"session_id": types.StringType,
		},
	}
}

// GetSession returns the value of the Session field in CreateSessionRequest_SdkV2 as
// a Session_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateSessionRequest_SdkV2) GetSession(ctx context.Context) (Session_SdkV2, bool) {
	var e Session_SdkV2
	if m.Session.IsNull() || m.Session.IsUnknown() {
		return e, false
	}
	var v []Session_SdkV2
	d := m.Session.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSession sets the value of the Session field in CreateSessionRequest_SdkV2.
func (m *CreateSessionRequest_SdkV2) SetSession(ctx context.Context, v Session_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["session"]
	m.Session = types.ListValueMust(t, vs)
}

type CreateSessionStoreRequest_SdkV2 struct {
	// The session store to create.
	SessionStore types.List `tfsdk:"session_store"`
	// Caller-provided, workspace-unique session store ID. It must be 3-55
	// characters, begin with a lowercase letter, and contain only lowercase
	// letters, digits, and hyphens.
	SessionStoreId types.String `tfsdk:"-"`
}

func (to *CreateSessionStoreRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateSessionStoreRequest_SdkV2) {
	if !from.SessionStore.IsNull() && !from.SessionStore.IsUnknown() {
		if toSessionStore, ok := to.GetSessionStore(ctx); ok {
			if fromSessionStore, ok := from.GetSessionStore(ctx); ok {
				// Recursively sync the fields of SessionStore
				toSessionStore.SyncFieldsDuringCreateOrUpdate(ctx, fromSessionStore)
				to.SetSessionStore(ctx, toSessionStore)
			}
		}
	}
}

func (to *CreateSessionStoreRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateSessionStoreRequest_SdkV2) {
	if !from.SessionStore.IsNull() && !from.SessionStore.IsUnknown() {
		if toSessionStore, ok := to.GetSessionStore(ctx); ok {
			if fromSessionStore, ok := from.GetSessionStore(ctx); ok {
				toSessionStore.SyncFieldsDuringRead(ctx, fromSessionStore)
				to.SetSessionStore(ctx, toSessionStore)
			}
		}
	}
}

func (m CreateSessionStoreRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["session_store"] = attrs["session_store"].SetRequired()
	attrs["session_store"] = attrs["session_store"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["session_store_id"] = attrs["session_store_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateSessionStoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateSessionStoreRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"session_store": reflect.TypeOf(SessionStore_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateSessionStoreRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateSessionStoreRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"session_store":    m.SessionStore,
			"session_store_id": m.SessionStoreId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateSessionStoreRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"session_store": basetypes.ListType{
				ElemType: SessionStore_SdkV2{}.Type(ctx),
			},
			"session_store_id": types.StringType,
		},
	}
}

// GetSessionStore returns the value of the SessionStore field in CreateSessionStoreRequest_SdkV2 as
// a SessionStore_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateSessionStoreRequest_SdkV2) GetSessionStore(ctx context.Context) (SessionStore_SdkV2, bool) {
	var e SessionStore_SdkV2
	if m.SessionStore.IsNull() || m.SessionStore.IsUnknown() {
		return e, false
	}
	var v []SessionStore_SdkV2
	d := m.SessionStore.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSessionStore sets the value of the SessionStore field in CreateSessionStoreRequest_SdkV2.
func (m *CreateSessionStoreRequest_SdkV2) SetSessionStore(ctx context.Context, v SessionStore_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["session_store"]
	m.SessionStore = types.ListValueMust(t, vs)
}

type DeleteManagedMemoryEntryRequest_SdkV2 struct {
	// Resource name in the form
	// `memory-stores/{managed_memory_store_id}/entries/{managed_memory_entry_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteManagedMemoryEntryRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteManagedMemoryEntryRequest_SdkV2) {
}

func (to *DeleteManagedMemoryEntryRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteManagedMemoryEntryRequest_SdkV2) {
}

func (m DeleteManagedMemoryEntryRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteManagedMemoryEntryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteManagedMemoryEntryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteManagedMemoryEntryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteManagedMemoryEntryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteManagedMemoryEntryRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteManagedMemoryStoreRequest_SdkV2 struct {
	// Resource name in the form `memory-stores/{managed_memory_store_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteManagedMemoryStoreRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteManagedMemoryStoreRequest_SdkV2) {
}

func (to *DeleteManagedMemoryStoreRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteManagedMemoryStoreRequest_SdkV2) {
}

func (m DeleteManagedMemoryStoreRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteManagedMemoryStoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteManagedMemoryStoreRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteManagedMemoryStoreRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteManagedMemoryStoreRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteManagedMemoryStoreRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteSessionRequest_SdkV2 struct {
	// Resource name in the form
	// `session-stores/{session_store_id}/sessions/{session_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteSessionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteSessionRequest_SdkV2) {
}

func (to *DeleteSessionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteSessionRequest_SdkV2) {
}

func (m DeleteSessionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSessionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteSessionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSessionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteSessionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteSessionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteSessionStoreRequest_SdkV2 struct {
	// Resource name in the form `session-stores/{session_store_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteSessionStoreRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteSessionStoreRequest_SdkV2) {
}

func (to *DeleteSessionStoreRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteSessionStoreRequest_SdkV2) {
}

func (m DeleteSessionStoreRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSessionStoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteSessionStoreRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSessionStoreRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteSessionStoreRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteSessionStoreRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Request to synchronously extract memories from a single session.
type ExtractMemoriesRequest_SdkV2 struct {
	// When true, extract and return the entries without writing them to the
	// memory store. Defaults to false, which persists the extracted entries and
	// returns them.
	DryRun types.Bool `tfsdk:"dry_run"`
	// Instructions steering what is extracted from the session.
	Instructions types.String `tfsdk:"instructions"`
	// Managed memory store the extracted entries are written to, in the form
	// `memory-stores/{managed_memory_store_id}`.
	MemoryStore types.String `tfsdk:"memory_store"`
	// Identifier of the session whose transcript is distilled into memories.
	SessionId types.String `tfsdk:"-"`
	// Session store containing the session, in the form
	// `session-stores/{session_store_id}`.
	SessionStore types.String `tfsdk:"-"`
}

func (to *ExtractMemoriesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExtractMemoriesRequest_SdkV2) {
}

func (to *ExtractMemoriesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExtractMemoriesRequest_SdkV2) {
}

func (m ExtractMemoriesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dry_run"] = attrs["dry_run"].SetOptional()
	attrs["instructions"] = attrs["instructions"].SetOptional()
	attrs["memory_store"] = attrs["memory_store"].SetRequired()
	attrs["session_store"] = attrs["session_store"].SetRequired()
	attrs["session_id"] = attrs["session_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExtractMemoriesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExtractMemoriesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExtractMemoriesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ExtractMemoriesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dry_run":       m.DryRun,
			"instructions":  m.Instructions,
			"memory_store":  m.MemoryStore,
			"session_id":    m.SessionId,
			"session_store": m.SessionStore,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExtractMemoriesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dry_run":       types.BoolType,
			"instructions":  types.StringType,
			"memory_store":  types.StringType,
			"session_id":    types.StringType,
			"session_store": types.StringType,
		},
	}
}

// Result of a single-session memory extraction.
type ExtractMemoriesResponse_SdkV2 struct {
	// The memory entries written by this extraction.
	Entries types.List `tfsdk:"entries"`
	// Correlation identifier for this extraction, for logging and tracing. Not
	// a fetchable resource.
	Name types.String `tfsdk:"name"`
}

func (to *ExtractMemoriesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExtractMemoriesResponse_SdkV2) {
	if !from.Entries.IsNull() && !from.Entries.IsUnknown() && to.Entries.IsNull() && len(from.Entries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Entries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Entries = from.Entries
	}
	if !from.Entries.IsNull() && !from.Entries.IsUnknown() {
		if toEntries, ok := to.GetEntries(ctx); ok {
			if fromEntries, ok := from.GetEntries(ctx); ok {
				// Recursively sync the fields of each Entries element by position.
				for i := range toEntries {
					if i < len(fromEntries) {
						toEntries[i].SyncFieldsDuringCreateOrUpdate(ctx, fromEntries[i])
					}
				}
				to.SetEntries(ctx, toEntries)
			}
		}
	}
}

func (to *ExtractMemoriesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExtractMemoriesResponse_SdkV2) {
	if !from.Entries.IsNull() && !from.Entries.IsUnknown() && to.Entries.IsNull() && len(from.Entries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Entries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Entries = from.Entries
	}
	if !from.Entries.IsNull() && !from.Entries.IsUnknown() {
		if toEntries, ok := to.GetEntries(ctx); ok {
			if fromEntries, ok := from.GetEntries(ctx); ok {
				for i := range toEntries {
					if i < len(fromEntries) {
						toEntries[i].SyncFieldsDuringRead(ctx, fromEntries[i])
					}
				}
				to.SetEntries(ctx, toEntries)
			}
		}
	}
}

func (m ExtractMemoriesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["entries"] = attrs["entries"].SetComputed()
	attrs["name"] = attrs["name"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExtractMemoriesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExtractMemoriesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"entries": reflect.TypeOf(ManagedMemoryEntry_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExtractMemoriesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ExtractMemoriesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entries": m.Entries,
			"name":    m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExtractMemoriesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entries": basetypes.ListType{
				ElemType: ManagedMemoryEntry_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
		},
	}
}

// GetEntries returns the value of the Entries field in ExtractMemoriesResponse_SdkV2 as
// a slice of ManagedMemoryEntry_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExtractMemoriesResponse_SdkV2) GetEntries(ctx context.Context) ([]ManagedMemoryEntry_SdkV2, bool) {
	if m.Entries.IsNull() || m.Entries.IsUnknown() {
		return nil, false
	}
	var v []ManagedMemoryEntry_SdkV2
	d := m.Entries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEntries sets the value of the Entries field in ExtractMemoriesResponse_SdkV2.
func (m *ExtractMemoriesResponse_SdkV2) SetEntries(ctx context.Context, v []ManagedMemoryEntry_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["entries"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Entries = types.ListValueMust(t, vs)
}

// Request to fork a session.
type ForkSessionRequest_SdkV2 struct {
	// Opaque caller-provided identifier for the application actor associated
	// with the forked session.
	ActorId types.String `tfsdk:"actor_id"`
	// Optional metadata for the fork.
	Metadata types.Map `tfsdk:"metadata"`
	// Resource name of the containing session store, in the form
	// `session-stores/{session_store_id}`.
	Parent types.String `tfsdk:"-"`
	// Optional unique ID for the forked session. A collision returns
	// `ALREADY_EXISTS`.
	SessionId types.String `tfsdk:"session_id"`
	// ID of the session to copy.
	SourceSessionId types.String `tfsdk:"source_session_id"`
	// Optional last item ID to copy through, inclusively. When omitted, the
	// fork atomically copies all items committed before the fork operation
	// begins.
	UpToItemId types.String `tfsdk:"up_to_item_id"`
}

func (to *ForkSessionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ForkSessionRequest_SdkV2) {
}

func (to *ForkSessionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ForkSessionRequest_SdkV2) {
}

func (m ForkSessionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["actor_id"] = attrs["actor_id"].SetRequired()
	attrs["metadata"] = attrs["metadata"].SetOptional()
	attrs["session_id"] = attrs["session_id"].SetOptional()
	attrs["source_session_id"] = attrs["source_session_id"].SetRequired()
	attrs["up_to_item_id"] = attrs["up_to_item_id"].SetOptional()
	attrs["parent"] = attrs["parent"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ForkSessionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ForkSessionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metadata": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ForkSessionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ForkSessionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"actor_id":          m.ActorId,
			"metadata":          m.Metadata,
			"parent":            m.Parent,
			"session_id":        m.SessionId,
			"source_session_id": m.SourceSessionId,
			"up_to_item_id":     m.UpToItemId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ForkSessionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"actor_id": types.StringType,
			"metadata": basetypes.MapType{
				ElemType: types.StringType,
			},
			"parent":            types.StringType,
			"session_id":        types.StringType,
			"source_session_id": types.StringType,
			"up_to_item_id":     types.StringType,
		},
	}
}

// GetMetadata returns the value of the Metadata field in ForkSessionRequest_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ForkSessionRequest_SdkV2) GetMetadata(ctx context.Context) (map[string]types.String, bool) {
	if m.Metadata.IsNull() || m.Metadata.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Metadata.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetadata sets the value of the Metadata field in ForkSessionRequest_SdkV2.
func (m *ForkSessionRequest_SdkV2) SetMetadata(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["metadata"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Metadata = types.MapValueMust(t, vs)
}

// Response from forking a session.
type ForkSessionResponse_SdkV2 struct {
	// The newly-created independent top-level session.
	Session types.List `tfsdk:"session"`
}

func (to *ForkSessionResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ForkSessionResponse_SdkV2) {
	if !from.Session.IsNull() && !from.Session.IsUnknown() {
		if toSession, ok := to.GetSession(ctx); ok {
			if fromSession, ok := from.GetSession(ctx); ok {
				// Recursively sync the fields of Session
				toSession.SyncFieldsDuringCreateOrUpdate(ctx, fromSession)
				to.SetSession(ctx, toSession)
			}
		}
	}
}

func (to *ForkSessionResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ForkSessionResponse_SdkV2) {
	if !from.Session.IsNull() && !from.Session.IsUnknown() {
		if toSession, ok := to.GetSession(ctx); ok {
			if fromSession, ok := from.GetSession(ctx); ok {
				toSession.SyncFieldsDuringRead(ctx, fromSession)
				to.SetSession(ctx, toSession)
			}
		}
	}
}

func (m ForkSessionResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["session"] = attrs["session"].SetOptional()
	attrs["session"] = attrs["session"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ForkSessionResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ForkSessionResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"session": reflect.TypeOf(Session_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ForkSessionResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ForkSessionResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"session": m.Session,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ForkSessionResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"session": basetypes.ListType{
				ElemType: Session_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSession returns the value of the Session field in ForkSessionResponse_SdkV2 as
// a Session_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ForkSessionResponse_SdkV2) GetSession(ctx context.Context) (Session_SdkV2, bool) {
	var e Session_SdkV2
	if m.Session.IsNull() || m.Session.IsUnknown() {
		return e, false
	}
	var v []Session_SdkV2
	d := m.Session.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSession sets the value of the Session field in ForkSessionResponse_SdkV2.
func (m *ForkSessionResponse_SdkV2) SetSession(ctx context.Context, v Session_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["session"]
	m.Session = types.ListValueMust(t, vs)
}

type GetManagedMemoryEntryRequest_SdkV2 struct {
	// Resource name in the form
	// `memory-stores/{managed_memory_store_id}/entries/{managed_memory_entry_id}`.
	Name types.String `tfsdk:"-"`
	// Fields to return, using proto field names such as `content` (not
	// `contents`). An omitted or empty mask returns the full entry, including
	// `content`; a non-empty mask returns only the requested fields.
	ReadMask types.String `tfsdk:"-"`
}

func (to *GetManagedMemoryEntryRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetManagedMemoryEntryRequest_SdkV2) {
}

func (to *GetManagedMemoryEntryRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetManagedMemoryEntryRequest_SdkV2) {
}

func (m GetManagedMemoryEntryRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["read_mask"] = attrs["read_mask"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetManagedMemoryEntryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetManagedMemoryEntryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetManagedMemoryEntryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetManagedMemoryEntryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":      m.Name,
			"read_mask": m.ReadMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetManagedMemoryEntryRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":      types.StringType,
			"read_mask": types.StringType,
		},
	}
}

type GetManagedMemoryStoreRequest_SdkV2 struct {
	// Resource name in the form `memory-stores/{managed_memory_store_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetManagedMemoryStoreRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetManagedMemoryStoreRequest_SdkV2) {
}

func (to *GetManagedMemoryStoreRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetManagedMemoryStoreRequest_SdkV2) {
}

func (m GetManagedMemoryStoreRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetManagedMemoryStoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetManagedMemoryStoreRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetManagedMemoryStoreRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetManagedMemoryStoreRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetManagedMemoryStoreRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetSessionRequest_SdkV2 struct {
	// Resource name in the form
	// `session-stores/{session_store_id}/sessions/{session_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetSessionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSessionRequest_SdkV2) {
}

func (to *GetSessionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetSessionRequest_SdkV2) {
}

func (m GetSessionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSessionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetSessionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSessionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetSessionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSessionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetSessionStoreRequest_SdkV2 struct {
	// Resource name in the form `session-stores/{session_store_id}`.
	Name types.String `tfsdk:"-"`
}

func (to *GetSessionStoreRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSessionStoreRequest_SdkV2) {
}

func (to *GetSessionStoreRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetSessionStoreRequest_SdkV2) {
}

func (m GetSessionStoreRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSessionStoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetSessionStoreRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSessionStoreRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetSessionStoreRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSessionStoreRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ListManagedMemoryEntriesRequest_SdkV2 struct {
	// Customer-provided identifier for the actor whose entries are listed.
	ActorId types.String `tfsdk:"-"`
	// Maximum number of entries to return. The service may return fewer entries
	// than requested. Defaults to 10; must be between 1 and 100.
	PageSize types.Int64 `tfsdk:"-"`
	// Opaque pagination token from a previous ListManagedMemoryEntries
	// response.
	PageToken types.String `tfsdk:"-"`
	// Managed memory store whose entries are listed, in the form
	// `memory-stores/{managed_memory_store_id}`.
	Parent types.String `tfsdk:"-"`
	// Optional path prefix used to restrict entries within the actor partition.
	PathPrefix types.String `tfsdk:"-"`
	// Fields to return in each entry, using proto field names such as `content`
	// (not `contents`). An omitted or empty mask returns each full entry,
	// including `content`; a non-empty mask returns only the requested fields.
	ReadMask types.String `tfsdk:"-"`
	// Optional session identifier. When set, only entries with this exact
	// `session_id` are returned. Omitted-session (cross-session) entries are
	// not included. Ignored when path is set.
	SessionId types.String `tfsdk:"-"`
}

func (to *ListManagedMemoryEntriesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListManagedMemoryEntriesRequest_SdkV2) {
}

func (to *ListManagedMemoryEntriesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListManagedMemoryEntriesRequest_SdkV2) {
}

func (m ListManagedMemoryEntriesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["actor_id"] = attrs["actor_id"].SetRequired()
	attrs["path_prefix"] = attrs["path_prefix"].SetOptional()
	attrs["session_id"] = attrs["session_id"].SetOptional()
	attrs["read_mask"] = attrs["read_mask"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListManagedMemoryEntriesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListManagedMemoryEntriesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListManagedMemoryEntriesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListManagedMemoryEntriesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"actor_id":    m.ActorId,
			"page_size":   m.PageSize,
			"page_token":  m.PageToken,
			"parent":      m.Parent,
			"path_prefix": m.PathPrefix,
			"read_mask":   m.ReadMask,
			"session_id":  m.SessionId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListManagedMemoryEntriesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"actor_id":    types.StringType,
			"page_size":   types.Int64Type,
			"page_token":  types.StringType,
			"parent":      types.StringType,
			"path_prefix": types.StringType,
			"read_mask":   types.StringType,
			"session_id":  types.StringType,
		},
	}
}

// Response containing managed memory entries.
type ListManagedMemoryEntriesResponse_SdkV2 struct {
	// Managed memory entries matching the request and its read mask.
	ManagedMemoryEntries types.List `tfsdk:"managed_memory_entries"`
	// Opaque pagination token. This field is omitted when there are no more
	// results.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListManagedMemoryEntriesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListManagedMemoryEntriesResponse_SdkV2) {
	if !from.ManagedMemoryEntries.IsNull() && !from.ManagedMemoryEntries.IsUnknown() && to.ManagedMemoryEntries.IsNull() && len(from.ManagedMemoryEntries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ManagedMemoryEntries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ManagedMemoryEntries = from.ManagedMemoryEntries
	}
	if !from.ManagedMemoryEntries.IsNull() && !from.ManagedMemoryEntries.IsUnknown() {
		if toManagedMemoryEntries, ok := to.GetManagedMemoryEntries(ctx); ok {
			if fromManagedMemoryEntries, ok := from.GetManagedMemoryEntries(ctx); ok {
				// Recursively sync the fields of each ManagedMemoryEntries element by position.
				for i := range toManagedMemoryEntries {
					if i < len(fromManagedMemoryEntries) {
						toManagedMemoryEntries[i].SyncFieldsDuringCreateOrUpdate(ctx, fromManagedMemoryEntries[i])
					}
				}
				to.SetManagedMemoryEntries(ctx, toManagedMemoryEntries)
			}
		}
	}
}

func (to *ListManagedMemoryEntriesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListManagedMemoryEntriesResponse_SdkV2) {
	if !from.ManagedMemoryEntries.IsNull() && !from.ManagedMemoryEntries.IsUnknown() && to.ManagedMemoryEntries.IsNull() && len(from.ManagedMemoryEntries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ManagedMemoryEntries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ManagedMemoryEntries = from.ManagedMemoryEntries
	}
	if !from.ManagedMemoryEntries.IsNull() && !from.ManagedMemoryEntries.IsUnknown() {
		if toManagedMemoryEntries, ok := to.GetManagedMemoryEntries(ctx); ok {
			if fromManagedMemoryEntries, ok := from.GetManagedMemoryEntries(ctx); ok {
				for i := range toManagedMemoryEntries {
					if i < len(fromManagedMemoryEntries) {
						toManagedMemoryEntries[i].SyncFieldsDuringRead(ctx, fromManagedMemoryEntries[i])
					}
				}
				to.SetManagedMemoryEntries(ctx, toManagedMemoryEntries)
			}
		}
	}
}

func (m ListManagedMemoryEntriesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["managed_memory_entries"] = attrs["managed_memory_entries"].SetComputed()
	attrs["next_page_token"] = attrs["next_page_token"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListManagedMemoryEntriesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListManagedMemoryEntriesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"managed_memory_entries": reflect.TypeOf(ManagedMemoryEntry_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListManagedMemoryEntriesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListManagedMemoryEntriesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"managed_memory_entries": m.ManagedMemoryEntries,
			"next_page_token":        m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListManagedMemoryEntriesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"managed_memory_entries": basetypes.ListType{
				ElemType: ManagedMemoryEntry_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetManagedMemoryEntries returns the value of the ManagedMemoryEntries field in ListManagedMemoryEntriesResponse_SdkV2 as
// a slice of ManagedMemoryEntry_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListManagedMemoryEntriesResponse_SdkV2) GetManagedMemoryEntries(ctx context.Context) ([]ManagedMemoryEntry_SdkV2, bool) {
	if m.ManagedMemoryEntries.IsNull() || m.ManagedMemoryEntries.IsUnknown() {
		return nil, false
	}
	var v []ManagedMemoryEntry_SdkV2
	d := m.ManagedMemoryEntries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetManagedMemoryEntries sets the value of the ManagedMemoryEntries field in ListManagedMemoryEntriesResponse_SdkV2.
func (m *ListManagedMemoryEntriesResponse_SdkV2) SetManagedMemoryEntries(ctx context.Context, v []ManagedMemoryEntry_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["managed_memory_entries"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ManagedMemoryEntries = types.ListValueMust(t, vs)
}

type ListManagedMemoryStoresRequest_SdkV2 struct {
	// Maximum number of stores to return. The service may return fewer stores
	// than requested. Defaults to 10; must be between 1 and 100.
	PageSize types.Int64 `tfsdk:"-"`
	// Opaque pagination token from a previous ListManagedMemoryStores response.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListManagedMemoryStoresRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListManagedMemoryStoresRequest_SdkV2) {
}

func (to *ListManagedMemoryStoresRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListManagedMemoryStoresRequest_SdkV2) {
}

func (m ListManagedMemoryStoresRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListManagedMemoryStoresRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListManagedMemoryStoresRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListManagedMemoryStoresRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListManagedMemoryStoresRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListManagedMemoryStoresRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// Response containing managed memory stores in the caller's workspace.
type ListManagedMemoryStoresResponse_SdkV2 struct {
	// Managed memory stores in the caller's workspace.
	ManagedMemoryStores types.List `tfsdk:"managed_memory_stores"`
	// Opaque pagination token. This field is omitted when there are no more
	// results.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListManagedMemoryStoresResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListManagedMemoryStoresResponse_SdkV2) {
	if !from.ManagedMemoryStores.IsNull() && !from.ManagedMemoryStores.IsUnknown() && to.ManagedMemoryStores.IsNull() && len(from.ManagedMemoryStores.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ManagedMemoryStores, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ManagedMemoryStores = from.ManagedMemoryStores
	}
	if !from.ManagedMemoryStores.IsNull() && !from.ManagedMemoryStores.IsUnknown() {
		if toManagedMemoryStores, ok := to.GetManagedMemoryStores(ctx); ok {
			if fromManagedMemoryStores, ok := from.GetManagedMemoryStores(ctx); ok {
				// Recursively sync the fields of each ManagedMemoryStores element by position.
				for i := range toManagedMemoryStores {
					if i < len(fromManagedMemoryStores) {
						toManagedMemoryStores[i].SyncFieldsDuringCreateOrUpdate(ctx, fromManagedMemoryStores[i])
					}
				}
				to.SetManagedMemoryStores(ctx, toManagedMemoryStores)
			}
		}
	}
}

func (to *ListManagedMemoryStoresResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListManagedMemoryStoresResponse_SdkV2) {
	if !from.ManagedMemoryStores.IsNull() && !from.ManagedMemoryStores.IsUnknown() && to.ManagedMemoryStores.IsNull() && len(from.ManagedMemoryStores.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ManagedMemoryStores, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ManagedMemoryStores = from.ManagedMemoryStores
	}
	if !from.ManagedMemoryStores.IsNull() && !from.ManagedMemoryStores.IsUnknown() {
		if toManagedMemoryStores, ok := to.GetManagedMemoryStores(ctx); ok {
			if fromManagedMemoryStores, ok := from.GetManagedMemoryStores(ctx); ok {
				for i := range toManagedMemoryStores {
					if i < len(fromManagedMemoryStores) {
						toManagedMemoryStores[i].SyncFieldsDuringRead(ctx, fromManagedMemoryStores[i])
					}
				}
				to.SetManagedMemoryStores(ctx, toManagedMemoryStores)
			}
		}
	}
}

func (m ListManagedMemoryStoresResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["managed_memory_stores"] = attrs["managed_memory_stores"].SetComputed()
	attrs["next_page_token"] = attrs["next_page_token"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListManagedMemoryStoresResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListManagedMemoryStoresResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"managed_memory_stores": reflect.TypeOf(ManagedMemoryStore_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListManagedMemoryStoresResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListManagedMemoryStoresResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"managed_memory_stores": m.ManagedMemoryStores,
			"next_page_token":       m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListManagedMemoryStoresResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"managed_memory_stores": basetypes.ListType{
				ElemType: ManagedMemoryStore_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetManagedMemoryStores returns the value of the ManagedMemoryStores field in ListManagedMemoryStoresResponse_SdkV2 as
// a slice of ManagedMemoryStore_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListManagedMemoryStoresResponse_SdkV2) GetManagedMemoryStores(ctx context.Context) ([]ManagedMemoryStore_SdkV2, bool) {
	if m.ManagedMemoryStores.IsNull() || m.ManagedMemoryStores.IsUnknown() {
		return nil, false
	}
	var v []ManagedMemoryStore_SdkV2
	d := m.ManagedMemoryStores.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetManagedMemoryStores sets the value of the ManagedMemoryStores field in ListManagedMemoryStoresResponse_SdkV2.
func (m *ListManagedMemoryStoresResponse_SdkV2) SetManagedMemoryStores(ctx context.Context, v []ManagedMemoryStore_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["managed_memory_stores"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ManagedMemoryStores = types.ListValueMust(t, vs)
}

type ListSessionItemsRequest_SdkV2 struct {
	// Sort order. Supported values are `create_time asc` and `create_time
	// desc`. The default is `create_time desc`, which returns the most recently
	// appended items first. Equal timestamps are resolved by committed append
	// order in the requested direction.
	OrderBy types.String `tfsdk:"-"`
	// Maximum number of items to return. Defaults to 10; must be between 1 and
	// 100.
	PageSize types.Int64 `tfsdk:"-"`
	// Token returned by a previous list request.
	PageToken types.String `tfsdk:"-"`
	// Resource name of the containing session, in the form
	// `session-stores/{session_store_id}/sessions/{session_id}`.
	Parent types.String `tfsdk:"-"`
}

func (to *ListSessionItemsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSessionItemsRequest_SdkV2) {
}

func (to *ListSessionItemsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSessionItemsRequest_SdkV2) {
}

func (m ListSessionItemsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["order_by"] = attrs["order_by"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSessionItemsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSessionItemsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSessionItemsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListSessionItemsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"order_by":   m.OrderBy,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSessionItemsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"order_by":   types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// Response containing a page of session items.
type ListSessionItemsResponse_SdkV2 struct {
	// Token to retrieve the next page.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Session items in the requested page.
	SessionItems types.List `tfsdk:"session_items"`
}

func (to *ListSessionItemsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSessionItemsResponse_SdkV2) {
	if !from.SessionItems.IsNull() && !from.SessionItems.IsUnknown() && to.SessionItems.IsNull() && len(from.SessionItems.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SessionItems, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SessionItems = from.SessionItems
	}
	if !from.SessionItems.IsNull() && !from.SessionItems.IsUnknown() {
		if toSessionItems, ok := to.GetSessionItems(ctx); ok {
			if fromSessionItems, ok := from.GetSessionItems(ctx); ok {
				// Recursively sync the fields of each SessionItems element by position.
				for i := range toSessionItems {
					if i < len(fromSessionItems) {
						toSessionItems[i].SyncFieldsDuringCreateOrUpdate(ctx, fromSessionItems[i])
					}
				}
				to.SetSessionItems(ctx, toSessionItems)
			}
		}
	}
}

func (to *ListSessionItemsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSessionItemsResponse_SdkV2) {
	if !from.SessionItems.IsNull() && !from.SessionItems.IsUnknown() && to.SessionItems.IsNull() && len(from.SessionItems.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SessionItems, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SessionItems = from.SessionItems
	}
	if !from.SessionItems.IsNull() && !from.SessionItems.IsUnknown() {
		if toSessionItems, ok := to.GetSessionItems(ctx); ok {
			if fromSessionItems, ok := from.GetSessionItems(ctx); ok {
				for i := range toSessionItems {
					if i < len(fromSessionItems) {
						toSessionItems[i].SyncFieldsDuringRead(ctx, fromSessionItems[i])
					}
				}
				to.SetSessionItems(ctx, toSessionItems)
			}
		}
	}
}

func (m ListSessionItemsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["session_items"] = attrs["session_items"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSessionItemsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSessionItemsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"session_items": reflect.TypeOf(SessionItem_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSessionItemsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListSessionItemsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"session_items":   m.SessionItems,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSessionItemsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"session_items": basetypes.ListType{
				ElemType: SessionItem_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSessionItems returns the value of the SessionItems field in ListSessionItemsResponse_SdkV2 as
// a slice of SessionItem_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListSessionItemsResponse_SdkV2) GetSessionItems(ctx context.Context) ([]SessionItem_SdkV2, bool) {
	if m.SessionItems.IsNull() || m.SessionItems.IsUnknown() {
		return nil, false
	}
	var v []SessionItem_SdkV2
	d := m.SessionItems.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSessionItems sets the value of the SessionItems field in ListSessionItemsResponse_SdkV2.
func (m *ListSessionItemsResponse_SdkV2) SetSessionItems(ctx context.Context, v []SessionItem_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["session_items"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SessionItems = types.ListValueMust(t, vs)
}

type ListSessionStoresRequest_SdkV2 struct {
	// Maximum number of session stores to return. Defaults to 10; must be
	// between 1 and 100.
	PageSize types.Int64 `tfsdk:"-"`
	// Token returned by a previous list request.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListSessionStoresRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSessionStoresRequest_SdkV2) {
}

func (to *ListSessionStoresRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSessionStoresRequest_SdkV2) {
}

func (m ListSessionStoresRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSessionStoresRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSessionStoresRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSessionStoresRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListSessionStoresRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSessionStoresRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// Response containing a page of session stores.
type ListSessionStoresResponse_SdkV2 struct {
	// Token to retrieve the next page.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Session stores in the requested page.
	SessionStores types.List `tfsdk:"session_stores"`
}

func (to *ListSessionStoresResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSessionStoresResponse_SdkV2) {
	if !from.SessionStores.IsNull() && !from.SessionStores.IsUnknown() && to.SessionStores.IsNull() && len(from.SessionStores.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SessionStores, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SessionStores = from.SessionStores
	}
	if !from.SessionStores.IsNull() && !from.SessionStores.IsUnknown() {
		if toSessionStores, ok := to.GetSessionStores(ctx); ok {
			if fromSessionStores, ok := from.GetSessionStores(ctx); ok {
				// Recursively sync the fields of each SessionStores element by position.
				for i := range toSessionStores {
					if i < len(fromSessionStores) {
						toSessionStores[i].SyncFieldsDuringCreateOrUpdate(ctx, fromSessionStores[i])
					}
				}
				to.SetSessionStores(ctx, toSessionStores)
			}
		}
	}
}

func (to *ListSessionStoresResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSessionStoresResponse_SdkV2) {
	if !from.SessionStores.IsNull() && !from.SessionStores.IsUnknown() && to.SessionStores.IsNull() && len(from.SessionStores.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SessionStores, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SessionStores = from.SessionStores
	}
	if !from.SessionStores.IsNull() && !from.SessionStores.IsUnknown() {
		if toSessionStores, ok := to.GetSessionStores(ctx); ok {
			if fromSessionStores, ok := from.GetSessionStores(ctx); ok {
				for i := range toSessionStores {
					if i < len(fromSessionStores) {
						toSessionStores[i].SyncFieldsDuringRead(ctx, fromSessionStores[i])
					}
				}
				to.SetSessionStores(ctx, toSessionStores)
			}
		}
	}
}

func (m ListSessionStoresResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["session_stores"] = attrs["session_stores"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSessionStoresResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSessionStoresResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"session_stores": reflect.TypeOf(SessionStore_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSessionStoresResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListSessionStoresResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"session_stores":  m.SessionStores,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSessionStoresResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"session_stores": basetypes.ListType{
				ElemType: SessionStore_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSessionStores returns the value of the SessionStores field in ListSessionStoresResponse_SdkV2 as
// a slice of SessionStore_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListSessionStoresResponse_SdkV2) GetSessionStores(ctx context.Context) ([]SessionStore_SdkV2, bool) {
	if m.SessionStores.IsNull() || m.SessionStores.IsUnknown() {
		return nil, false
	}
	var v []SessionStore_SdkV2
	d := m.SessionStores.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSessionStores sets the value of the SessionStores field in ListSessionStoresResponse_SdkV2.
func (m *ListSessionStoresResponse_SdkV2) SetSessionStores(ctx context.Context, v []SessionStore_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["session_stores"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SessionStores = types.ListValueMust(t, vs)
}

type ListSessionsRequest_SdkV2 struct {
	// Filter expression. Supported fields include `actor_id` and `metadata`;
	// for example, `actor_id = "support-customer-123"`.
	Filter types.String `tfsdk:"-"`
	// Sort order. Defaults to `last_activity_time desc`. Page-token
	// continuation is exactly-once when ordering by `create_time` (immutable);
	// ordering by `last_activity_time` is best-effort, because that value
	// changes as a session gains activity, so a session updated between page
	// requests may be repeated or skipped. To enumerate every session exactly
	// once, order by `create_time`.
	OrderBy types.String `tfsdk:"-"`
	// Maximum number of sessions to return. Defaults to 10; must be between 1
	// and 100.
	PageSize types.Int64 `tfsdk:"-"`
	// Token returned by a previous list request.
	PageToken types.String `tfsdk:"-"`
	// Resource name of the containing session store, in the form
	// `session-stores/{session_store_id}`.
	Parent types.String `tfsdk:"-"`
}

func (to *ListSessionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSessionsRequest_SdkV2) {
}

func (to *ListSessionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSessionsRequest_SdkV2) {
}

func (m ListSessionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["order_by"] = attrs["order_by"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSessionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSessionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSessionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListSessionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter":     m.Filter,
			"order_by":   m.OrderBy,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSessionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter":     types.StringType,
			"order_by":   types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// Response containing a page of sessions.
type ListSessionsResponse_SdkV2 struct {
	// Token to retrieve the next page.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Sessions in the requested page.
	Sessions types.List `tfsdk:"sessions"`
}

func (to *ListSessionsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSessionsResponse_SdkV2) {
	if !from.Sessions.IsNull() && !from.Sessions.IsUnknown() && to.Sessions.IsNull() && len(from.Sessions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Sessions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Sessions = from.Sessions
	}
	if !from.Sessions.IsNull() && !from.Sessions.IsUnknown() {
		if toSessions, ok := to.GetSessions(ctx); ok {
			if fromSessions, ok := from.GetSessions(ctx); ok {
				// Recursively sync the fields of each Sessions element by position.
				for i := range toSessions {
					if i < len(fromSessions) {
						toSessions[i].SyncFieldsDuringCreateOrUpdate(ctx, fromSessions[i])
					}
				}
				to.SetSessions(ctx, toSessions)
			}
		}
	}
}

func (to *ListSessionsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSessionsResponse_SdkV2) {
	if !from.Sessions.IsNull() && !from.Sessions.IsUnknown() && to.Sessions.IsNull() && len(from.Sessions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Sessions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Sessions = from.Sessions
	}
	if !from.Sessions.IsNull() && !from.Sessions.IsUnknown() {
		if toSessions, ok := to.GetSessions(ctx); ok {
			if fromSessions, ok := from.GetSessions(ctx); ok {
				for i := range toSessions {
					if i < len(fromSessions) {
						toSessions[i].SyncFieldsDuringRead(ctx, fromSessions[i])
					}
				}
				to.SetSessions(ctx, toSessions)
			}
		}
	}
}

func (m ListSessionsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["sessions"] = attrs["sessions"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSessionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSessionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sessions": reflect.TypeOf(Session_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSessionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListSessionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"sessions":        m.Sessions,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSessionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"sessions": basetypes.ListType{
				ElemType: Session_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSessions returns the value of the Sessions field in ListSessionsResponse_SdkV2 as
// a slice of Session_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListSessionsResponse_SdkV2) GetSessions(ctx context.Context) ([]Session_SdkV2, bool) {
	if m.Sessions.IsNull() || m.Sessions.IsUnknown() {
		return nil, false
	}
	var v []Session_SdkV2
	d := m.Sessions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSessions sets the value of the Sessions field in ListSessionsResponse_SdkV2.
func (m *ListSessionsResponse_SdkV2) SetSessions(ctx context.Context, v []Session_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sessions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Sessions = types.ListValueMust(t, vs)
}

// A workspace-scoped entry in a managed memory store.
type ManagedMemoryEntry_SdkV2 struct {
	// Customer-provided identifier for the actor whose memory this entry
	// represents.
	ActorId types.String `tfsdk:"actor_id"`
	// Optional free-form memory content.
	Content types.String `tfsdk:"content"`
	// Time when the entry was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Human-readable description of the memory entry.
	Description types.String `tfsdk:"description"`
	// Resource name in the form
	// `memory-stores/{managed_memory_store_id}/entries/{managed_memory_entry_id}`.
	Name types.String `tfsdk:"name"`
	// Absolute, case-sensitive path identifying the entry within its actor and
	// optional session. Paths must begin with `/` and must not contain empty,
	// `.` or `..` segments.
	Path types.String `tfsdk:"path"`
	// Optional identifier for the session associated with this memory entry.
	// When omitted, the entry applies across the actor's sessions.
	SessionId types.String `tfsdk:"session_id"`
	// Which writer created this entry. Caller sets this on Create; immutable
	// after creation.
	SourceType types.String `tfsdk:"source_type"`
	// Time when the entry was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *ManagedMemoryEntry_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ManagedMemoryEntry_SdkV2) {
}

func (to *ManagedMemoryEntry_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ManagedMemoryEntry_SdkV2) {
}

func (m ManagedMemoryEntry_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["actor_id"] = attrs["actor_id"].SetRequired()
	attrs["actor_id"] = attrs["actor_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["content"] = attrs["content"].SetOptional()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["path"] = attrs["path"].SetRequired()
	attrs["path"] = attrs["path"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["session_id"] = attrs["session_id"].SetOptional()
	attrs["session_id"] = attrs["session_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["source_type"] = attrs["source_type"].SetOptional()
	attrs["source_type"] = attrs["source_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ManagedMemoryEntry.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ManagedMemoryEntry_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ManagedMemoryEntry_SdkV2
// only implements ToObjectValue() and Type().
func (m ManagedMemoryEntry_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"actor_id":    m.ActorId,
			"content":     m.Content,
			"create_time": m.CreateTime,
			"description": m.Description,
			"name":        m.Name,
			"path":        m.Path,
			"session_id":  m.SessionId,
			"source_type": m.SourceType,
			"update_time": m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ManagedMemoryEntry_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"actor_id":    types.StringType,
			"content":     types.StringType,
			"create_time": timetypes.RFC3339{}.Type(ctx),
			"description": types.StringType,
			"name":        types.StringType,
			"path":        types.StringType,
			"session_id":  types.StringType,
			"source_type": types.StringType,
			"update_time": timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// One relevance-ranked managed memory search result.
type ManagedMemoryEntrySearchResult_SdkV2 struct {
	// Managed memory entry matching the query.
	ManagedMemoryEntry types.List `tfsdk:"managed_memory_entry"`
	// Relevance score for the result. Higher scores are more relevant.
	Score types.Float64 `tfsdk:"score"`
}

func (to *ManagedMemoryEntrySearchResult_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ManagedMemoryEntrySearchResult_SdkV2) {
	if !from.ManagedMemoryEntry.IsNull() && !from.ManagedMemoryEntry.IsUnknown() {
		if toManagedMemoryEntry, ok := to.GetManagedMemoryEntry(ctx); ok {
			if fromManagedMemoryEntry, ok := from.GetManagedMemoryEntry(ctx); ok {
				// Recursively sync the fields of ManagedMemoryEntry
				toManagedMemoryEntry.SyncFieldsDuringCreateOrUpdate(ctx, fromManagedMemoryEntry)
				to.SetManagedMemoryEntry(ctx, toManagedMemoryEntry)
			}
		}
	}
}

func (to *ManagedMemoryEntrySearchResult_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ManagedMemoryEntrySearchResult_SdkV2) {
	if !from.ManagedMemoryEntry.IsNull() && !from.ManagedMemoryEntry.IsUnknown() {
		if toManagedMemoryEntry, ok := to.GetManagedMemoryEntry(ctx); ok {
			if fromManagedMemoryEntry, ok := from.GetManagedMemoryEntry(ctx); ok {
				toManagedMemoryEntry.SyncFieldsDuringRead(ctx, fromManagedMemoryEntry)
				to.SetManagedMemoryEntry(ctx, toManagedMemoryEntry)
			}
		}
	}
}

func (m ManagedMemoryEntrySearchResult_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["managed_memory_entry"] = attrs["managed_memory_entry"].SetComputed()
	attrs["managed_memory_entry"] = attrs["managed_memory_entry"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["score"] = attrs["score"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ManagedMemoryEntrySearchResult.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ManagedMemoryEntrySearchResult_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"managed_memory_entry": reflect.TypeOf(ManagedMemoryEntry_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ManagedMemoryEntrySearchResult_SdkV2
// only implements ToObjectValue() and Type().
func (m ManagedMemoryEntrySearchResult_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"managed_memory_entry": m.ManagedMemoryEntry,
			"score":                m.Score,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ManagedMemoryEntrySearchResult_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"managed_memory_entry": basetypes.ListType{
				ElemType: ManagedMemoryEntry_SdkV2{}.Type(ctx),
			},
			"score": types.Float64Type,
		},
	}
}

// GetManagedMemoryEntry returns the value of the ManagedMemoryEntry field in ManagedMemoryEntrySearchResult_SdkV2 as
// a ManagedMemoryEntry_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ManagedMemoryEntrySearchResult_SdkV2) GetManagedMemoryEntry(ctx context.Context) (ManagedMemoryEntry_SdkV2, bool) {
	var e ManagedMemoryEntry_SdkV2
	if m.ManagedMemoryEntry.IsNull() || m.ManagedMemoryEntry.IsUnknown() {
		return e, false
	}
	var v []ManagedMemoryEntry_SdkV2
	d := m.ManagedMemoryEntry.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetManagedMemoryEntry sets the value of the ManagedMemoryEntry field in ManagedMemoryEntrySearchResult_SdkV2.
func (m *ManagedMemoryEntrySearchResult_SdkV2) SetManagedMemoryEntry(ctx context.Context, v ManagedMemoryEntry_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["managed_memory_entry"]
	m.ManagedMemoryEntry = types.ListValueMust(t, vs)
}

// A workspace-scoped managed memory store backed by service-managed storage.
type ManagedMemoryStore_SdkV2 struct {
	// Time when the store was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Workspace-local user ID of the authenticated principal that created the
	// store. This is immutable server-set attribution and does not grant
	// access; authorization is evaluated from the authenticated request
	// context.
	CreatorUserId types.String `tfsdk:"creator_user_id"`
	// Human-readable description of the memory store.
	Description types.String `tfsdk:"description"`
	// Deprecated compatibility alias for the caller-provided managed memory
	// store ID. Canonical clients provide the ID through
	// `CreateMemoryStoreRequest.managed_memory_store_id` and use `name` as the
	// resource identifier.
	DisplayName types.String `tfsdk:"display_name"`
	// Resource name in the form `memory-stores/{managed_memory_store_id}`.
	Name types.String `tfsdk:"name"`
	// Deprecated alias for `creator_user_id`. This identifies the original
	// creator, not a transferable owner. Use `creator_user_id` instead.
	OwnerUserId types.String `tfsdk:"owner_user_id"`
	// Service-managed storage backing this memory store.
	StorageBackend types.List `tfsdk:"storage_backend"`
	// Time when the store was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// Workspace that owns the memory store.
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

func (to *ManagedMemoryStore_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ManagedMemoryStore_SdkV2) {
	if !from.StorageBackend.IsNull() && !from.StorageBackend.IsUnknown() {
		if toStorageBackend, ok := to.GetStorageBackend(ctx); ok {
			if fromStorageBackend, ok := from.GetStorageBackend(ctx); ok {
				// Recursively sync the fields of StorageBackend
				toStorageBackend.SyncFieldsDuringCreateOrUpdate(ctx, fromStorageBackend)
				to.SetStorageBackend(ctx, toStorageBackend)
			}
		}
	}
}

func (to *ManagedMemoryStore_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ManagedMemoryStore_SdkV2) {
	if !from.StorageBackend.IsNull() && !from.StorageBackend.IsUnknown() {
		if toStorageBackend, ok := to.GetStorageBackend(ctx); ok {
			if fromStorageBackend, ok := from.GetStorageBackend(ctx); ok {
				toStorageBackend.SyncFieldsDuringRead(ctx, fromStorageBackend)
				to.SetStorageBackend(ctx, toStorageBackend)
			}
		}
	}
}

func (m ManagedMemoryStore_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator_user_id"] = attrs["creator_user_id"].SetComputed()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["display_name"] = attrs["display_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetComputed()
	attrs["owner_user_id"] = attrs["owner_user_id"].SetComputed()
	attrs["storage_backend"] = attrs["storage_backend"].SetComputed()
	attrs["storage_backend"] = attrs["storage_backend"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ManagedMemoryStore.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ManagedMemoryStore_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"storage_backend": reflect.TypeOf(StorageBackend_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ManagedMemoryStore_SdkV2
// only implements ToObjectValue() and Type().
func (m ManagedMemoryStore_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":     m.CreateTime,
			"creator_user_id": m.CreatorUserId,
			"description":     m.Description,
			"display_name":    m.DisplayName,
			"name":            m.Name,
			"owner_user_id":   m.OwnerUserId,
			"storage_backend": m.StorageBackend,
			"update_time":     m.UpdateTime,
			"workspace_id":    m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ManagedMemoryStore_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":     timetypes.RFC3339{}.Type(ctx),
			"creator_user_id": types.StringType,
			"description":     types.StringType,
			"display_name":    types.StringType,
			"name":            types.StringType,
			"owner_user_id":   types.StringType,
			"storage_backend": basetypes.ListType{
				ElemType: StorageBackend_SdkV2{}.Type(ctx),
			},
			"update_time":  timetypes.RFC3339{}.Type(ctx),
			"workspace_id": types.Int64Type,
		},
	}
}

// GetStorageBackend returns the value of the StorageBackend field in ManagedMemoryStore_SdkV2 as
// a StorageBackend_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ManagedMemoryStore_SdkV2) GetStorageBackend(ctx context.Context) (StorageBackend_SdkV2, bool) {
	var e StorageBackend_SdkV2
	if m.StorageBackend.IsNull() || m.StorageBackend.IsUnknown() {
		return e, false
	}
	var v []StorageBackend_SdkV2
	d := m.StorageBackend.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStorageBackend sets the value of the StorageBackend field in ManagedMemoryStore_SdkV2.
func (m *ManagedMemoryStore_SdkV2) SetStorageBackend(ctx context.Context, v StorageBackend_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["storage_backend"]
	m.StorageBackend = types.ListValueMust(t, vs)
}

// Request to pop an item from a session.
type PopSessionItemRequest_SdkV2 struct {
	// Resource name of the containing session, in the form
	// `session-stores/{session_store_id}/sessions/{session_id}`.
	Parent types.String `tfsdk:"-"`
}

func (to *PopSessionItemRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PopSessionItemRequest_SdkV2) {
}

func (to *PopSessionItemRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PopSessionItemRequest_SdkV2) {
}

func (m PopSessionItemRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PopSessionItemRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PopSessionItemRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PopSessionItemRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m PopSessionItemRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parent": m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PopSessionItemRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent": types.StringType,
		},
	}
}

// Response containing the popped item.
type PopSessionItemResponse_SdkV2 struct {
	// Removed item, if any.
	Item types.List `tfsdk:"item"`
}

func (to *PopSessionItemResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PopSessionItemResponse_SdkV2) {
	if !from.Item.IsNull() && !from.Item.IsUnknown() {
		if toItem, ok := to.GetItem(ctx); ok {
			if fromItem, ok := from.GetItem(ctx); ok {
				// Recursively sync the fields of Item
				toItem.SyncFieldsDuringCreateOrUpdate(ctx, fromItem)
				to.SetItem(ctx, toItem)
			}
		}
	}
}

func (to *PopSessionItemResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PopSessionItemResponse_SdkV2) {
	if !from.Item.IsNull() && !from.Item.IsUnknown() {
		if toItem, ok := to.GetItem(ctx); ok {
			if fromItem, ok := from.GetItem(ctx); ok {
				toItem.SyncFieldsDuringRead(ctx, fromItem)
				to.SetItem(ctx, toItem)
			}
		}
	}
}

func (m PopSessionItemResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["item"] = attrs["item"].SetOptional()
	attrs["item"] = attrs["item"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PopSessionItemResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PopSessionItemResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"item": reflect.TypeOf(SessionItem_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PopSessionItemResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m PopSessionItemResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"item": m.Item,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PopSessionItemResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"item": basetypes.ListType{
				ElemType: SessionItem_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetItem returns the value of the Item field in PopSessionItemResponse_SdkV2 as
// a SessionItem_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PopSessionItemResponse_SdkV2) GetItem(ctx context.Context) (SessionItem_SdkV2, bool) {
	var e SessionItem_SdkV2
	if m.Item.IsNull() || m.Item.IsUnknown() {
		return e, false
	}
	var v []SessionItem_SdkV2
	d := m.Item.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetItem sets the value of the Item field in PopSessionItemResponse_SdkV2.
func (m *PopSessionItemResponse_SdkV2) SetItem(ctx context.Context, v SessionItem_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["item"]
	m.Item = types.ListValueMust(t, vs)
}

// Request to search managed memory entries by text query for one actor. Search
// returns a relevance-ranked top-N result set and does not currently paginate.
type SearchManagedMemoryEntriesRequest_SdkV2 struct {
	// Customer-provided identifier for the actor whose entries are searched.
	ActorId types.String `tfsdk:"actor_id"`
	// Deprecated alias for `page_size`. When both fields are set, their values
	// must match.
	Limit types.Int64 `tfsdk:"limit"`
	// Maximum number of relevance-ranked entries to return. Defaults to 10 and
	// must be between 1 and 100.
	PageSize types.Int64 `tfsdk:"page_size"`
	// Reserved for pagination compatibility. The server currently ignores this
	// field because Search returns a ranked top-N result set.
	PageToken types.String `tfsdk:"page_token"`
	// Managed memory store whose entries are searched, in the form
	// `memory-stores/{managed_memory_store_id}`.
	Parent types.String `tfsdk:"-"`
	// Optional absolute, case-sensitive path prefix used to restrict searched
	// entries within the actor partition. The prefix must begin with `/` and
	// must not contain empty, `.` or `..` segments.
	PathPrefix types.String `tfsdk:"path_prefix"`
	// Free-form search query.
	Query types.String `tfsdk:"query"`
	// Fields to return in each matching entry, using proto field names such as
	// `content` (not `contents`). An omitted or empty mask returns each full
	// entry, including `content`; a non-empty mask returns only the requested
	// fields. Search scores are always returned.
	//
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	ReadMask types.String `tfsdk:"read_mask"`
	// Optional session identifier. When set, only entries with this exact
	// `session_id` are searched. Omitted-session (cross-session) entries are
	// not included.
	SessionId types.String `tfsdk:"session_id"`
}

func (to *SearchManagedMemoryEntriesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchManagedMemoryEntriesRequest_SdkV2) {
}

func (to *SearchManagedMemoryEntriesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SearchManagedMemoryEntriesRequest_SdkV2) {
}

func (m SearchManagedMemoryEntriesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["actor_id"] = attrs["actor_id"].SetRequired()
	attrs["limit"] = attrs["limit"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["path_prefix"] = attrs["path_prefix"].SetOptional()
	attrs["query"] = attrs["query"].SetRequired()
	attrs["read_mask"] = attrs["read_mask"].SetOptional()
	attrs["session_id"] = attrs["session_id"].SetOptional()
	attrs["parent"] = attrs["parent"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchManagedMemoryEntriesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SearchManagedMemoryEntriesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchManagedMemoryEntriesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m SearchManagedMemoryEntriesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"actor_id":    m.ActorId,
			"limit":       m.Limit,
			"page_size":   m.PageSize,
			"page_token":  m.PageToken,
			"parent":      m.Parent,
			"path_prefix": m.PathPrefix,
			"query":       m.Query,
			"read_mask":   m.ReadMask,
			"session_id":  m.SessionId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchManagedMemoryEntriesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"actor_id":    types.StringType,
			"limit":       types.Int64Type,
			"page_size":   types.Int64Type,
			"page_token":  types.StringType,
			"parent":      types.StringType,
			"path_prefix": types.StringType,
			"query":       types.StringType,
			"read_mask":   types.StringType,
			"session_id":  types.StringType,
		},
	}
}

// Response containing managed memory entries ranked by relevance.
type SearchManagedMemoryEntriesResponse_SdkV2 struct {
	// Deprecated compatibility alias for clients migrating to `results`. This
	// contains the same entries in the same order, but omits their relevance
	// scores.
	ManagedMemoryEntries types.List `tfsdk:"managed_memory_entries"`
	// Opaque pagination token. Search currently returns an unpaginated ranked
	// top-N result set, so the server does not populate this field.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Canonical matching entries and relevance scores, ordered most relevant
	// first.
	Results types.List `tfsdk:"results"`
}

func (to *SearchManagedMemoryEntriesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchManagedMemoryEntriesResponse_SdkV2) {
	if !from.ManagedMemoryEntries.IsNull() && !from.ManagedMemoryEntries.IsUnknown() && to.ManagedMemoryEntries.IsNull() && len(from.ManagedMemoryEntries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ManagedMemoryEntries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ManagedMemoryEntries = from.ManagedMemoryEntries
	}
	if !from.ManagedMemoryEntries.IsNull() && !from.ManagedMemoryEntries.IsUnknown() {
		if toManagedMemoryEntries, ok := to.GetManagedMemoryEntries(ctx); ok {
			if fromManagedMemoryEntries, ok := from.GetManagedMemoryEntries(ctx); ok {
				// Recursively sync the fields of each ManagedMemoryEntries element by position.
				for i := range toManagedMemoryEntries {
					if i < len(fromManagedMemoryEntries) {
						toManagedMemoryEntries[i].SyncFieldsDuringCreateOrUpdate(ctx, fromManagedMemoryEntries[i])
					}
				}
				to.SetManagedMemoryEntries(ctx, toManagedMemoryEntries)
			}
		}
	}
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
	if !from.Results.IsNull() && !from.Results.IsUnknown() {
		if toResults, ok := to.GetResults(ctx); ok {
			if fromResults, ok := from.GetResults(ctx); ok {
				// Recursively sync the fields of each Results element by position.
				for i := range toResults {
					if i < len(fromResults) {
						toResults[i].SyncFieldsDuringCreateOrUpdate(ctx, fromResults[i])
					}
				}
				to.SetResults(ctx, toResults)
			}
		}
	}
}

func (to *SearchManagedMemoryEntriesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SearchManagedMemoryEntriesResponse_SdkV2) {
	if !from.ManagedMemoryEntries.IsNull() && !from.ManagedMemoryEntries.IsUnknown() && to.ManagedMemoryEntries.IsNull() && len(from.ManagedMemoryEntries.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ManagedMemoryEntries, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ManagedMemoryEntries = from.ManagedMemoryEntries
	}
	if !from.ManagedMemoryEntries.IsNull() && !from.ManagedMemoryEntries.IsUnknown() {
		if toManagedMemoryEntries, ok := to.GetManagedMemoryEntries(ctx); ok {
			if fromManagedMemoryEntries, ok := from.GetManagedMemoryEntries(ctx); ok {
				for i := range toManagedMemoryEntries {
					if i < len(fromManagedMemoryEntries) {
						toManagedMemoryEntries[i].SyncFieldsDuringRead(ctx, fromManagedMemoryEntries[i])
					}
				}
				to.SetManagedMemoryEntries(ctx, toManagedMemoryEntries)
			}
		}
	}
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
	if !from.Results.IsNull() && !from.Results.IsUnknown() {
		if toResults, ok := to.GetResults(ctx); ok {
			if fromResults, ok := from.GetResults(ctx); ok {
				for i := range toResults {
					if i < len(fromResults) {
						toResults[i].SyncFieldsDuringRead(ctx, fromResults[i])
					}
				}
				to.SetResults(ctx, toResults)
			}
		}
	}
}

func (m SearchManagedMemoryEntriesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["managed_memory_entries"] = attrs["managed_memory_entries"].SetComputed()
	attrs["next_page_token"] = attrs["next_page_token"].SetComputed()
	attrs["results"] = attrs["results"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchManagedMemoryEntriesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SearchManagedMemoryEntriesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"managed_memory_entries": reflect.TypeOf(ManagedMemoryEntry_SdkV2{}),
		"results":                reflect.TypeOf(ManagedMemoryEntrySearchResult_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchManagedMemoryEntriesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m SearchManagedMemoryEntriesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"managed_memory_entries": m.ManagedMemoryEntries,
			"next_page_token":        m.NextPageToken,
			"results":                m.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchManagedMemoryEntriesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"managed_memory_entries": basetypes.ListType{
				ElemType: ManagedMemoryEntry_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
			"results": basetypes.ListType{
				ElemType: ManagedMemoryEntrySearchResult_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetManagedMemoryEntries returns the value of the ManagedMemoryEntries field in SearchManagedMemoryEntriesResponse_SdkV2 as
// a slice of ManagedMemoryEntry_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchManagedMemoryEntriesResponse_SdkV2) GetManagedMemoryEntries(ctx context.Context) ([]ManagedMemoryEntry_SdkV2, bool) {
	if m.ManagedMemoryEntries.IsNull() || m.ManagedMemoryEntries.IsUnknown() {
		return nil, false
	}
	var v []ManagedMemoryEntry_SdkV2
	d := m.ManagedMemoryEntries.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetManagedMemoryEntries sets the value of the ManagedMemoryEntries field in SearchManagedMemoryEntriesResponse_SdkV2.
func (m *SearchManagedMemoryEntriesResponse_SdkV2) SetManagedMemoryEntries(ctx context.Context, v []ManagedMemoryEntry_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["managed_memory_entries"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ManagedMemoryEntries = types.ListValueMust(t, vs)
}

// GetResults returns the value of the Results field in SearchManagedMemoryEntriesResponse_SdkV2 as
// a slice of ManagedMemoryEntrySearchResult_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchManagedMemoryEntriesResponse_SdkV2) GetResults(ctx context.Context) ([]ManagedMemoryEntrySearchResult_SdkV2, bool) {
	if m.Results.IsNull() || m.Results.IsUnknown() {
		return nil, false
	}
	var v []ManagedMemoryEntrySearchResult_SdkV2
	d := m.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in SearchManagedMemoryEntriesResponse_SdkV2.
func (m *SearchManagedMemoryEntriesResponse_SdkV2) SetResults(ctx context.Context, v []ManagedMemoryEntrySearchResult_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Results = types.ListValueMust(t, vs)
}

// A durable logical interaction stored within a Session Store.
type Session_SdkV2 struct {
	// Opaque caller-provided identifier for the application actor associated
	// with the session.
	//
	// This is application data and has no Databricks authentication or
	// authorization semantics. Use the same value as the Managed Memory Entry
	// `actor_id` when storing memories associated with this actor. Every
	// session must set it. A child session must use the same value as its
	// parent.
	ActorId types.String `tfsdk:"actor_id"`
	// Time when the session was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Time when the session's item history was last mutated.
	LastActivityTime timetypes.RFC3339 `tfsdk:"last_activity_time"`
	// Mutable caller-defined string labels.
	Metadata types.Map `tfsdk:"metadata"`
	// Resource name in the form
	// `session-stores/{session_store_id}/sessions/{session_id}`.
	Name types.String `tfsdk:"name"`
	// Immediate parent session ID. Set only at creation for child sessions,
	// immutable thereafter, and restricted to the same store.
	ParentSessionId types.String `tfsdk:"parent_session_id"`
	// Top-level session ID in the spawn tree. This equals `session_id` for a
	// root or fork and is inherited transitively by child sessions.
	RootSessionId types.String `tfsdk:"root_session_id"`
	// Unique session ID. The service generates a UUID unless the caller
	// supplies `CreateSessionRequest.session_id`.
	SessionId types.String `tfsdk:"session_id"`
	// Time when session resource fields last changed.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *Session_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Session_SdkV2) {
}

func (to *Session_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Session_SdkV2) {
}

func (m Session_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["actor_id"] = attrs["actor_id"].SetRequired()
	attrs["actor_id"] = attrs["actor_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["last_activity_time"] = attrs["last_activity_time"].SetComputed()
	attrs["metadata"] = attrs["metadata"].SetOptional()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["parent_session_id"] = attrs["parent_session_id"].SetOptional()
	attrs["parent_session_id"] = attrs["parent_session_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["root_session_id"] = attrs["root_session_id"].SetComputed()
	attrs["session_id"] = attrs["session_id"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Session.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Session_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metadata": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Session_SdkV2
// only implements ToObjectValue() and Type().
func (m Session_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"actor_id":           m.ActorId,
			"create_time":        m.CreateTime,
			"last_activity_time": m.LastActivityTime,
			"metadata":           m.Metadata,
			"name":               m.Name,
			"parent_session_id":  m.ParentSessionId,
			"root_session_id":    m.RootSessionId,
			"session_id":         m.SessionId,
			"update_time":        m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Session_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"actor_id":           types.StringType,
			"create_time":        timetypes.RFC3339{}.Type(ctx),
			"last_activity_time": timetypes.RFC3339{}.Type(ctx),
			"metadata": basetypes.MapType{
				ElemType: types.StringType,
			},
			"name":              types.StringType,
			"parent_session_id": types.StringType,
			"root_session_id":   types.StringType,
			"session_id":        types.StringType,
			"update_time":       timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetMetadata returns the value of the Metadata field in Session_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Session_SdkV2) GetMetadata(ctx context.Context) (map[string]types.String, bool) {
	if m.Metadata.IsNull() || m.Metadata.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Metadata.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetadata sets the value of the Metadata field in Session_SdkV2.
func (m *Session_SdkV2) SetMetadata(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["metadata"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Metadata = types.MapValueMust(t, vs)
}

// A transcript entry in a session's history.
type SessionItem_SdkV2 struct {
	// Server-assigned time when the append commits. Values are nondecreasing
	// within a session. Item listing orders by this timestamp; equal timestamps
	// are resolved by committed append order.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Complete SDK-native, JSON-compatible item. The service stores and returns
	// this value without interpreting provider-specific fields such as `type`,
	// `role`, or `content`.
	Data jsontypes.Normalized `tfsdk:"data"`
	// Stable service-generated item ID.
	ItemId types.String `tfsdk:"item_id"`
}

func (to *SessionItem_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SessionItem_SdkV2) {
}

func (to *SessionItem_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SessionItem_SdkV2) {
}

func (m SessionItem_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["data"] = attrs["data"].SetRequired()
	attrs["item_id"] = attrs["item_id"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SessionItem.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SessionItem_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SessionItem_SdkV2
// only implements ToObjectValue() and Type().
func (m SessionItem_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time": m.CreateTime,
			"data":        m.Data,
			"item_id":     m.ItemId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SessionItem_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": timetypes.RFC3339{}.Type(ctx),
			"data":        jsontypes.NormalizedType{},
			"item_id":     types.StringType,
		},
	}
}

// A workspace-scoped session store.
type SessionStore_SdkV2 struct {
	// Time when the store was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Workspace-local user ID of the authenticated principal that created the
	// store. This is immutable server-set attribution and does not grant
	// access; authorization is evaluated from the authenticated request
	// context.
	CreatorUserId types.String `tfsdk:"creator_user_id"`
	// Human-readable description of the session store.
	Description types.String `tfsdk:"description"`
	// Mutable caller-defined string labels.
	Metadata types.Map `tfsdk:"metadata"`
	// Resource name in the form `session-stores/{session_store_id}`.
	Name types.String `tfsdk:"name"`
	// Time when the store was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *SessionStore_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SessionStore_SdkV2) {
}

func (to *SessionStore_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SessionStore_SdkV2) {
}

func (m SessionStore_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator_user_id"] = attrs["creator_user_id"].SetComputed()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["metadata"] = attrs["metadata"].SetOptional()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SessionStore.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SessionStore_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metadata": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SessionStore_SdkV2
// only implements ToObjectValue() and Type().
func (m SessionStore_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":     m.CreateTime,
			"creator_user_id": m.CreatorUserId,
			"description":     m.Description,
			"metadata":        m.Metadata,
			"name":            m.Name,
			"update_time":     m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SessionStore_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":     timetypes.RFC3339{}.Type(ctx),
			"creator_user_id": types.StringType,
			"description":     types.StringType,
			"metadata": basetypes.MapType{
				ElemType: types.StringType,
			},
			"name":        types.StringType,
			"update_time": timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetMetadata returns the value of the Metadata field in SessionStore_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SessionStore_SdkV2) GetMetadata(ctx context.Context) (map[string]types.String, bool) {
	if m.Metadata.IsNull() || m.Metadata.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Metadata.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetadata sets the value of the Metadata field in SessionStore_SdkV2.
func (m *SessionStore_SdkV2) SetMetadata(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["metadata"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Metadata = types.MapValueMust(t, vs)
}

// Service-managed storage backing a managed memory store.
type StorageBackend_SdkV2 struct {
	// Backend-specific identifier. For Lakebase, this is the project ID.
	BackendId types.String `tfsdk:"backend_id"`
	// Type of the storage backend.
	BackendType types.String `tfsdk:"backend_type"`
}

func (to *StorageBackend_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StorageBackend_SdkV2) {
}

func (to *StorageBackend_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StorageBackend_SdkV2) {
}

func (m StorageBackend_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["backend_id"] = attrs["backend_id"].SetComputed()
	attrs["backend_type"] = attrs["backend_type"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StorageBackend.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m StorageBackend_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StorageBackend_SdkV2
// only implements ToObjectValue() and Type().
func (m StorageBackend_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"backend_id":   m.BackendId,
			"backend_type": m.BackendType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StorageBackend_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"backend_id":   types.StringType,
			"backend_type": types.StringType,
		},
	}
}

type UpdateManagedMemoryEntryRequest_SdkV2 struct {
	// The managed memory entry to update.
	ManagedMemoryEntry types.List `tfsdk:"managed_memory_entry"`
	// Resource name in the form
	// `memory-stores/{managed_memory_store_id}/entries/{managed_memory_entry_id}`.
	Name types.String `tfsdk:"-"`
	// Fields to update. Only `content` and `description` may be updated.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateManagedMemoryEntryRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateManagedMemoryEntryRequest_SdkV2) {
	if !from.ManagedMemoryEntry.IsNull() && !from.ManagedMemoryEntry.IsUnknown() {
		if toManagedMemoryEntry, ok := to.GetManagedMemoryEntry(ctx); ok {
			if fromManagedMemoryEntry, ok := from.GetManagedMemoryEntry(ctx); ok {
				// Recursively sync the fields of ManagedMemoryEntry
				toManagedMemoryEntry.SyncFieldsDuringCreateOrUpdate(ctx, fromManagedMemoryEntry)
				to.SetManagedMemoryEntry(ctx, toManagedMemoryEntry)
			}
		}
	}
}

func (to *UpdateManagedMemoryEntryRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateManagedMemoryEntryRequest_SdkV2) {
	if !from.ManagedMemoryEntry.IsNull() && !from.ManagedMemoryEntry.IsUnknown() {
		if toManagedMemoryEntry, ok := to.GetManagedMemoryEntry(ctx); ok {
			if fromManagedMemoryEntry, ok := from.GetManagedMemoryEntry(ctx); ok {
				toManagedMemoryEntry.SyncFieldsDuringRead(ctx, fromManagedMemoryEntry)
				to.SetManagedMemoryEntry(ctx, toManagedMemoryEntry)
			}
		}
	}
}

func (m UpdateManagedMemoryEntryRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["managed_memory_entry"] = attrs["managed_memory_entry"].SetRequired()
	attrs["managed_memory_entry"] = attrs["managed_memory_entry"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetComputed()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateManagedMemoryEntryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateManagedMemoryEntryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"managed_memory_entry": reflect.TypeOf(ManagedMemoryEntry_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateManagedMemoryEntryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateManagedMemoryEntryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"managed_memory_entry": m.ManagedMemoryEntry,
			"name":                 m.Name,
			"update_mask":          m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateManagedMemoryEntryRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"managed_memory_entry": basetypes.ListType{
				ElemType: ManagedMemoryEntry_SdkV2{}.Type(ctx),
			},
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetManagedMemoryEntry returns the value of the ManagedMemoryEntry field in UpdateManagedMemoryEntryRequest_SdkV2 as
// a ManagedMemoryEntry_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateManagedMemoryEntryRequest_SdkV2) GetManagedMemoryEntry(ctx context.Context) (ManagedMemoryEntry_SdkV2, bool) {
	var e ManagedMemoryEntry_SdkV2
	if m.ManagedMemoryEntry.IsNull() || m.ManagedMemoryEntry.IsUnknown() {
		return e, false
	}
	var v []ManagedMemoryEntry_SdkV2
	d := m.ManagedMemoryEntry.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetManagedMemoryEntry sets the value of the ManagedMemoryEntry field in UpdateManagedMemoryEntryRequest_SdkV2.
func (m *UpdateManagedMemoryEntryRequest_SdkV2) SetManagedMemoryEntry(ctx context.Context, v ManagedMemoryEntry_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["managed_memory_entry"]
	m.ManagedMemoryEntry = types.ListValueMust(t, vs)
}

type UpdateManagedMemoryStoreRequest_SdkV2 struct {
	// The managed memory store to update. `name` is taken from the URL.
	ManagedMemoryStore types.List `tfsdk:"managed_memory_store"`
	// Resource name in the form `memory-stores/{managed_memory_store_id}`.
	Name types.String `tfsdk:"-"`
	// Only `description` may be updated.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateManagedMemoryStoreRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateManagedMemoryStoreRequest_SdkV2) {
	if !from.ManagedMemoryStore.IsNull() && !from.ManagedMemoryStore.IsUnknown() {
		if toManagedMemoryStore, ok := to.GetManagedMemoryStore(ctx); ok {
			if fromManagedMemoryStore, ok := from.GetManagedMemoryStore(ctx); ok {
				// Recursively sync the fields of ManagedMemoryStore
				toManagedMemoryStore.SyncFieldsDuringCreateOrUpdate(ctx, fromManagedMemoryStore)
				to.SetManagedMemoryStore(ctx, toManagedMemoryStore)
			}
		}
	}
}

func (to *UpdateManagedMemoryStoreRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateManagedMemoryStoreRequest_SdkV2) {
	if !from.ManagedMemoryStore.IsNull() && !from.ManagedMemoryStore.IsUnknown() {
		if toManagedMemoryStore, ok := to.GetManagedMemoryStore(ctx); ok {
			if fromManagedMemoryStore, ok := from.GetManagedMemoryStore(ctx); ok {
				toManagedMemoryStore.SyncFieldsDuringRead(ctx, fromManagedMemoryStore)
				to.SetManagedMemoryStore(ctx, toManagedMemoryStore)
			}
		}
	}
}

func (m UpdateManagedMemoryStoreRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["managed_memory_store"] = attrs["managed_memory_store"].SetRequired()
	attrs["managed_memory_store"] = attrs["managed_memory_store"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetComputed()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateManagedMemoryStoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateManagedMemoryStoreRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"managed_memory_store": reflect.TypeOf(ManagedMemoryStore_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateManagedMemoryStoreRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateManagedMemoryStoreRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"managed_memory_store": m.ManagedMemoryStore,
			"name":                 m.Name,
			"update_mask":          m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateManagedMemoryStoreRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"managed_memory_store": basetypes.ListType{
				ElemType: ManagedMemoryStore_SdkV2{}.Type(ctx),
			},
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetManagedMemoryStore returns the value of the ManagedMemoryStore field in UpdateManagedMemoryStoreRequest_SdkV2 as
// a ManagedMemoryStore_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateManagedMemoryStoreRequest_SdkV2) GetManagedMemoryStore(ctx context.Context) (ManagedMemoryStore_SdkV2, bool) {
	var e ManagedMemoryStore_SdkV2
	if m.ManagedMemoryStore.IsNull() || m.ManagedMemoryStore.IsUnknown() {
		return e, false
	}
	var v []ManagedMemoryStore_SdkV2
	d := m.ManagedMemoryStore.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetManagedMemoryStore sets the value of the ManagedMemoryStore field in UpdateManagedMemoryStoreRequest_SdkV2.
func (m *UpdateManagedMemoryStoreRequest_SdkV2) SetManagedMemoryStore(ctx context.Context, v ManagedMemoryStore_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["managed_memory_store"]
	m.ManagedMemoryStore = types.ListValueMust(t, vs)
}

type UpdateSessionRequest_SdkV2 struct {
	// Resource name in the form
	// `session-stores/{session_store_id}/sessions/{session_id}`.
	Name types.String `tfsdk:"-"`
	// Session to update.
	Session types.List `tfsdk:"session"`
	// Fields to update. Only `metadata` is mutable; any other path returns
	// `INVALID_PARAMETER_VALUE`.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateSessionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateSessionRequest_SdkV2) {
	if !from.Session.IsNull() && !from.Session.IsUnknown() {
		if toSession, ok := to.GetSession(ctx); ok {
			if fromSession, ok := from.GetSession(ctx); ok {
				// Recursively sync the fields of Session
				toSession.SyncFieldsDuringCreateOrUpdate(ctx, fromSession)
				to.SetSession(ctx, toSession)
			}
		}
	}
}

func (to *UpdateSessionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateSessionRequest_SdkV2) {
	if !from.Session.IsNull() && !from.Session.IsUnknown() {
		if toSession, ok := to.GetSession(ctx); ok {
			if fromSession, ok := from.GetSession(ctx); ok {
				toSession.SyncFieldsDuringRead(ctx, fromSession)
				to.SetSession(ctx, toSession)
			}
		}
	}
}

func (m UpdateSessionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["session"] = attrs["session"].SetRequired()
	attrs["session"] = attrs["session"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetComputed()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateSessionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateSessionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"session": reflect.TypeOf(Session_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSessionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateSessionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":        m.Name,
			"session":     m.Session,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateSessionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"session": basetypes.ListType{
				ElemType: Session_SdkV2{}.Type(ctx),
			},
			"update_mask": types.StringType,
		},
	}
}

// GetSession returns the value of the Session field in UpdateSessionRequest_SdkV2 as
// a Session_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateSessionRequest_SdkV2) GetSession(ctx context.Context) (Session_SdkV2, bool) {
	var e Session_SdkV2
	if m.Session.IsNull() || m.Session.IsUnknown() {
		return e, false
	}
	var v []Session_SdkV2
	d := m.Session.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSession sets the value of the Session field in UpdateSessionRequest_SdkV2.
func (m *UpdateSessionRequest_SdkV2) SetSession(ctx context.Context, v Session_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["session"]
	m.Session = types.ListValueMust(t, vs)
}

type UpdateSessionStoreRequest_SdkV2 struct {
	// Resource name in the form `session-stores/{session_store_id}`.
	Name types.String `tfsdk:"-"`
	// Session store to update.
	SessionStore types.List `tfsdk:"session_store"`
	// Fields to update. Only `description` and `metadata` are mutable; any
	// other path returns `INVALID_PARAMETER_VALUE`.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateSessionStoreRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateSessionStoreRequest_SdkV2) {
	if !from.SessionStore.IsNull() && !from.SessionStore.IsUnknown() {
		if toSessionStore, ok := to.GetSessionStore(ctx); ok {
			if fromSessionStore, ok := from.GetSessionStore(ctx); ok {
				// Recursively sync the fields of SessionStore
				toSessionStore.SyncFieldsDuringCreateOrUpdate(ctx, fromSessionStore)
				to.SetSessionStore(ctx, toSessionStore)
			}
		}
	}
}

func (to *UpdateSessionStoreRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateSessionStoreRequest_SdkV2) {
	if !from.SessionStore.IsNull() && !from.SessionStore.IsUnknown() {
		if toSessionStore, ok := to.GetSessionStore(ctx); ok {
			if fromSessionStore, ok := from.GetSessionStore(ctx); ok {
				toSessionStore.SyncFieldsDuringRead(ctx, fromSessionStore)
				to.SetSessionStore(ctx, toSessionStore)
			}
		}
	}
}

func (m UpdateSessionStoreRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["session_store"] = attrs["session_store"].SetRequired()
	attrs["session_store"] = attrs["session_store"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetComputed()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateSessionStoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateSessionStoreRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"session_store": reflect.TypeOf(SessionStore_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSessionStoreRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateSessionStoreRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":          m.Name,
			"session_store": m.SessionStore,
			"update_mask":   m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateSessionStoreRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"session_store": basetypes.ListType{
				ElemType: SessionStore_SdkV2{}.Type(ctx),
			},
			"update_mask": types.StringType,
		},
	}
}

// GetSessionStore returns the value of the SessionStore field in UpdateSessionStoreRequest_SdkV2 as
// a SessionStore_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateSessionStoreRequest_SdkV2) GetSessionStore(ctx context.Context) (SessionStore_SdkV2, bool) {
	var e SessionStore_SdkV2
	if m.SessionStore.IsNull() || m.SessionStore.IsUnknown() {
		return e, false
	}
	var v []SessionStore_SdkV2
	d := m.SessionStore.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSessionStore sets the value of the SessionStore field in UpdateSessionStoreRequest_SdkV2.
func (m *UpdateSessionStoreRequest_SdkV2) SetSessionStore(ctx context.Context, v SessionStore_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["session_store"]
	m.SessionStore = types.ListValueMust(t, vs)
}
