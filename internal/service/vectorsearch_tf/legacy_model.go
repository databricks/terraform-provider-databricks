// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package vectorsearch_tf

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

type ColumnInfo_SdkV2 struct {
	// Name of the column.
	Name types.String `tfsdk:"name"`
}

func (to *ColumnInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ColumnInfo_SdkV2) {
}

func (to *ColumnInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ColumnInfo_SdkV2) {
}

func (m ColumnInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ColumnInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ColumnInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ColumnInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m ColumnInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ColumnInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type CreateEndpoint_SdkV2 struct {
	// The budget policy id to be applied
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// Type of endpoint
	EndpointType types.String `tfsdk:"endpoint_type"`
	// Name of the vector search endpoint
	Name types.String `tfsdk:"name"`
	// The usage policy id to be applied once we've migrated to usage policies
	UsagePolicyId types.String `tfsdk:"usage_policy_id"`
}

func (to *CreateEndpoint_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateEndpoint_SdkV2) {
}

func (to *CreateEndpoint_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateEndpoint_SdkV2) {
}

func (m CreateEndpoint_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["endpoint_type"] = attrs["endpoint_type"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["usage_policy_id"] = attrs["usage_policy_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateEndpoint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateEndpoint_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateEndpoint_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateEndpoint_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget_policy_id": m.BudgetPolicyId,
			"endpoint_type":    m.EndpointType,
			"name":             m.Name,
			"usage_policy_id":  m.UsagePolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateEndpoint_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_policy_id": types.StringType,
			"endpoint_type":    types.StringType,
			"name":             types.StringType,
			"usage_policy_id":  types.StringType,
		},
	}
}

type CreateVectorIndexRequest_SdkV2 struct {
	// Specification for Delta Sync Index. Required if `index_type` is
	// `DELTA_SYNC`.
	DeltaSyncIndexSpec types.List `tfsdk:"delta_sync_index_spec"`
	// Specification for Direct Vector Access Index. Required if `index_type` is
	// `DIRECT_ACCESS`.
	DirectAccessIndexSpec types.List `tfsdk:"direct_access_index_spec"`
	// Name of the endpoint to be used for serving the index
	EndpointName types.String `tfsdk:"endpoint_name"`

	IndexType types.String `tfsdk:"index_type"`
	// Name of the index
	Name types.String `tfsdk:"name"`
	// Primary key of the index
	PrimaryKey types.String `tfsdk:"primary_key"`
}

func (to *CreateVectorIndexRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateVectorIndexRequest_SdkV2) {
	if !from.DeltaSyncIndexSpec.IsNull() && !from.DeltaSyncIndexSpec.IsUnknown() {
		if toDeltaSyncIndexSpec, ok := to.GetDeltaSyncIndexSpec(ctx); ok {
			if fromDeltaSyncIndexSpec, ok := from.GetDeltaSyncIndexSpec(ctx); ok {
				// Recursively sync the fields of DeltaSyncIndexSpec
				toDeltaSyncIndexSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromDeltaSyncIndexSpec)
				to.SetDeltaSyncIndexSpec(ctx, toDeltaSyncIndexSpec)
			}
		}
	}
	if !from.DirectAccessIndexSpec.IsNull() && !from.DirectAccessIndexSpec.IsUnknown() {
		if toDirectAccessIndexSpec, ok := to.GetDirectAccessIndexSpec(ctx); ok {
			if fromDirectAccessIndexSpec, ok := from.GetDirectAccessIndexSpec(ctx); ok {
				// Recursively sync the fields of DirectAccessIndexSpec
				toDirectAccessIndexSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromDirectAccessIndexSpec)
				to.SetDirectAccessIndexSpec(ctx, toDirectAccessIndexSpec)
			}
		}
	}
}

func (to *CreateVectorIndexRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateVectorIndexRequest_SdkV2) {
	if !from.DeltaSyncIndexSpec.IsNull() && !from.DeltaSyncIndexSpec.IsUnknown() {
		if toDeltaSyncIndexSpec, ok := to.GetDeltaSyncIndexSpec(ctx); ok {
			if fromDeltaSyncIndexSpec, ok := from.GetDeltaSyncIndexSpec(ctx); ok {
				toDeltaSyncIndexSpec.SyncFieldsDuringRead(ctx, fromDeltaSyncIndexSpec)
				to.SetDeltaSyncIndexSpec(ctx, toDeltaSyncIndexSpec)
			}
		}
	}
	if !from.DirectAccessIndexSpec.IsNull() && !from.DirectAccessIndexSpec.IsUnknown() {
		if toDirectAccessIndexSpec, ok := to.GetDirectAccessIndexSpec(ctx); ok {
			if fromDirectAccessIndexSpec, ok := from.GetDirectAccessIndexSpec(ctx); ok {
				toDirectAccessIndexSpec.SyncFieldsDuringRead(ctx, fromDirectAccessIndexSpec)
				to.SetDirectAccessIndexSpec(ctx, toDirectAccessIndexSpec)
			}
		}
	}
}

func (m CreateVectorIndexRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["delta_sync_index_spec"] = attrs["delta_sync_index_spec"].SetOptional()
	attrs["delta_sync_index_spec"] = attrs["delta_sync_index_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["direct_access_index_spec"] = attrs["direct_access_index_spec"].SetOptional()
	attrs["direct_access_index_spec"] = attrs["direct_access_index_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["endpoint_name"] = attrs["endpoint_name"].SetRequired()
	attrs["index_type"] = attrs["index_type"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["primary_key"] = attrs["primary_key"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateVectorIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateVectorIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"delta_sync_index_spec":    reflect.TypeOf(DeltaSyncVectorIndexSpecRequest_SdkV2{}),
		"direct_access_index_spec": reflect.TypeOf(DirectAccessVectorIndexSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateVectorIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateVectorIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"delta_sync_index_spec":    m.DeltaSyncIndexSpec,
			"direct_access_index_spec": m.DirectAccessIndexSpec,
			"endpoint_name":            m.EndpointName,
			"index_type":               m.IndexType,
			"name":                     m.Name,
			"primary_key":              m.PrimaryKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateVectorIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"delta_sync_index_spec": basetypes.ListType{
				ElemType: DeltaSyncVectorIndexSpecRequest_SdkV2{}.Type(ctx),
			},
			"direct_access_index_spec": basetypes.ListType{
				ElemType: DirectAccessVectorIndexSpec_SdkV2{}.Type(ctx),
			},
			"endpoint_name": types.StringType,
			"index_type":    types.StringType,
			"name":          types.StringType,
			"primary_key":   types.StringType,
		},
	}
}

// GetDeltaSyncIndexSpec returns the value of the DeltaSyncIndexSpec field in CreateVectorIndexRequest_SdkV2 as
// a DeltaSyncVectorIndexSpecRequest_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateVectorIndexRequest_SdkV2) GetDeltaSyncIndexSpec(ctx context.Context) (DeltaSyncVectorIndexSpecRequest_SdkV2, bool) {
	var e DeltaSyncVectorIndexSpecRequest_SdkV2
	if m.DeltaSyncIndexSpec.IsNull() || m.DeltaSyncIndexSpec.IsUnknown() {
		return e, false
	}
	var v []DeltaSyncVectorIndexSpecRequest_SdkV2
	d := m.DeltaSyncIndexSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDeltaSyncIndexSpec sets the value of the DeltaSyncIndexSpec field in CreateVectorIndexRequest_SdkV2.
func (m *CreateVectorIndexRequest_SdkV2) SetDeltaSyncIndexSpec(ctx context.Context, v DeltaSyncVectorIndexSpecRequest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["delta_sync_index_spec"]
	m.DeltaSyncIndexSpec = types.ListValueMust(t, vs)
}

// GetDirectAccessIndexSpec returns the value of the DirectAccessIndexSpec field in CreateVectorIndexRequest_SdkV2 as
// a DirectAccessVectorIndexSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateVectorIndexRequest_SdkV2) GetDirectAccessIndexSpec(ctx context.Context) (DirectAccessVectorIndexSpec_SdkV2, bool) {
	var e DirectAccessVectorIndexSpec_SdkV2
	if m.DirectAccessIndexSpec.IsNull() || m.DirectAccessIndexSpec.IsUnknown() {
		return e, false
	}
	var v []DirectAccessVectorIndexSpec_SdkV2
	d := m.DirectAccessIndexSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDirectAccessIndexSpec sets the value of the DirectAccessIndexSpec field in CreateVectorIndexRequest_SdkV2.
func (m *CreateVectorIndexRequest_SdkV2) SetDirectAccessIndexSpec(ctx context.Context, v DirectAccessVectorIndexSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["direct_access_index_spec"]
	m.DirectAccessIndexSpec = types.ListValueMust(t, vs)
}

type CustomTag_SdkV2 struct {
	// Key field for a vector search endpoint tag.
	Key types.String `tfsdk:"key"`
	// [Optional] Value field for a vector search endpoint tag.
	Value types.String `tfsdk:"value"`
}

func (to *CustomTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CustomTag_SdkV2) {
}

func (to *CustomTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CustomTag_SdkV2) {
}

func (m CustomTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetRequired()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CustomTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CustomTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomTag_SdkV2
// only implements ToObjectValue() and Type().
func (m CustomTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CustomTag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type DeleteDataResult_SdkV2 struct {
	// List of primary keys for rows that failed to process.
	FailedPrimaryKeys types.List `tfsdk:"failed_primary_keys"`
	// Count of successfully processed rows.
	SuccessRowCount types.Int64 `tfsdk:"success_row_count"`
}

func (to *DeleteDataResult_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDataResult_SdkV2) {
	if !from.FailedPrimaryKeys.IsNull() && !from.FailedPrimaryKeys.IsUnknown() && to.FailedPrimaryKeys.IsNull() && len(from.FailedPrimaryKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FailedPrimaryKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FailedPrimaryKeys = from.FailedPrimaryKeys
	}
}

func (to *DeleteDataResult_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDataResult_SdkV2) {
	if !from.FailedPrimaryKeys.IsNull() && !from.FailedPrimaryKeys.IsUnknown() && to.FailedPrimaryKeys.IsNull() && len(from.FailedPrimaryKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FailedPrimaryKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FailedPrimaryKeys = from.FailedPrimaryKeys
	}
}

func (m DeleteDataResult_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["failed_primary_keys"] = attrs["failed_primary_keys"].SetOptional()
	attrs["success_row_count"] = attrs["success_row_count"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDataResult.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDataResult_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"failed_primary_keys": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDataResult_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDataResult_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"failed_primary_keys": m.FailedPrimaryKeys,
			"success_row_count":   m.SuccessRowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDataResult_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"failed_primary_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
			"success_row_count": types.Int64Type,
		},
	}
}

// GetFailedPrimaryKeys returns the value of the FailedPrimaryKeys field in DeleteDataResult_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeleteDataResult_SdkV2) GetFailedPrimaryKeys(ctx context.Context) ([]types.String, bool) {
	if m.FailedPrimaryKeys.IsNull() || m.FailedPrimaryKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.FailedPrimaryKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFailedPrimaryKeys sets the value of the FailedPrimaryKeys field in DeleteDataResult_SdkV2.
func (m *DeleteDataResult_SdkV2) SetFailedPrimaryKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["failed_primary_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FailedPrimaryKeys = types.ListValueMust(t, vs)
}

type DeleteDataVectorIndexRequest_SdkV2 struct {
	// Name of the vector index where data is to be deleted. Must be a Direct
	// Vector Access Index.
	IndexName types.String `tfsdk:"-"`
	// List of primary keys for the data to be deleted.
	PrimaryKeys types.List `tfsdk:"-"`
}

func (to *DeleteDataVectorIndexRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDataVectorIndexRequest_SdkV2) {
}

func (to *DeleteDataVectorIndexRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDataVectorIndexRequest_SdkV2) {
}

func (m DeleteDataVectorIndexRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["index_name"] = attrs["index_name"].SetRequired()
	attrs["primary_keys"] = attrs["primary_keys"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDataVectorIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDataVectorIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"primary_keys": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDataVectorIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDataVectorIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name":   m.IndexName,
			"primary_keys": m.PrimaryKeys,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDataVectorIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name": types.StringType,
			"primary_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetPrimaryKeys returns the value of the PrimaryKeys field in DeleteDataVectorIndexRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeleteDataVectorIndexRequest_SdkV2) GetPrimaryKeys(ctx context.Context) ([]types.String, bool) {
	if m.PrimaryKeys.IsNull() || m.PrimaryKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.PrimaryKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrimaryKeys sets the value of the PrimaryKeys field in DeleteDataVectorIndexRequest_SdkV2.
func (m *DeleteDataVectorIndexRequest_SdkV2) SetPrimaryKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["primary_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PrimaryKeys = types.ListValueMust(t, vs)
}

type DeleteDataVectorIndexResponse_SdkV2 struct {
	// Result of the upsert or delete operation.
	Result types.List `tfsdk:"result"`
	// Status of the delete operation.
	Status types.String `tfsdk:"status"`
}

func (to *DeleteDataVectorIndexResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDataVectorIndexResponse_SdkV2) {
	if !from.Result.IsNull() && !from.Result.IsUnknown() {
		if toResult, ok := to.GetResult(ctx); ok {
			if fromResult, ok := from.GetResult(ctx); ok {
				// Recursively sync the fields of Result
				toResult.SyncFieldsDuringCreateOrUpdate(ctx, fromResult)
				to.SetResult(ctx, toResult)
			}
		}
	}
}

func (to *DeleteDataVectorIndexResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDataVectorIndexResponse_SdkV2) {
	if !from.Result.IsNull() && !from.Result.IsUnknown() {
		if toResult, ok := to.GetResult(ctx); ok {
			if fromResult, ok := from.GetResult(ctx); ok {
				toResult.SyncFieldsDuringRead(ctx, fromResult)
				to.SetResult(ctx, toResult)
			}
		}
	}
}

func (m DeleteDataVectorIndexResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["result"] = attrs["result"].SetOptional()
	attrs["result"] = attrs["result"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDataVectorIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDataVectorIndexResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"result": reflect.TypeOf(DeleteDataResult_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDataVectorIndexResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDataVectorIndexResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"result": m.Result,
			"status": m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDataVectorIndexResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"result": basetypes.ListType{
				ElemType: DeleteDataResult_SdkV2{}.Type(ctx),
			},
			"status": types.StringType,
		},
	}
}

// GetResult returns the value of the Result field in DeleteDataVectorIndexResponse_SdkV2 as
// a DeleteDataResult_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DeleteDataVectorIndexResponse_SdkV2) GetResult(ctx context.Context) (DeleteDataResult_SdkV2, bool) {
	var e DeleteDataResult_SdkV2
	if m.Result.IsNull() || m.Result.IsUnknown() {
		return e, false
	}
	var v []DeleteDataResult_SdkV2
	d := m.Result.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetResult sets the value of the Result field in DeleteDataVectorIndexResponse_SdkV2.
func (m *DeleteDataVectorIndexResponse_SdkV2) SetResult(ctx context.Context, v DeleteDataResult_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["result"]
	m.Result = types.ListValueMust(t, vs)
}

type DeleteEndpointRequest_SdkV2 struct {
	// Name of the vector search endpoint
	EndpointName types.String `tfsdk:"-"`
}

func (to *DeleteEndpointRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteEndpointRequest_SdkV2) {
}

func (to *DeleteEndpointRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteEndpointRequest_SdkV2) {
}

func (m DeleteEndpointRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint_name"] = attrs["endpoint_name"].SetRequired()

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
			"endpoint_name": m.EndpointName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_name": types.StringType,
		},
	}
}

type DeleteEndpointResponse_SdkV2 struct {
}

func (to *DeleteEndpointResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteEndpointResponse_SdkV2) {
}

func (to *DeleteEndpointResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteEndpointResponse_SdkV2) {
}

func (m DeleteEndpointResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteEndpointResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteEndpointResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteEndpointResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteEndpointResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteEndpointResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteIndexRequest_SdkV2 struct {
	// Name of the index
	IndexName types.String `tfsdk:"-"`
}

func (to *DeleteIndexRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteIndexRequest_SdkV2) {
}

func (to *DeleteIndexRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteIndexRequest_SdkV2) {
}

func (m DeleteIndexRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["index_name"] = attrs["index_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name": m.IndexName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name": types.StringType,
		},
	}
}

type DeleteIndexResponse_SdkV2 struct {
}

func (to *DeleteIndexResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteIndexResponse_SdkV2) {
}

func (to *DeleteIndexResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteIndexResponse_SdkV2) {
}

func (m DeleteIndexResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteIndexResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteIndexResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteIndexResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteIndexResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeltaSyncVectorIndexSpecRequest_SdkV2 struct {
	// [Optional] Select the columns to sync with the vector index. If you leave
	// this field blank, all columns from the source table are synced with the
	// index. The primary key column and embedding source column or embedding
	// vector column are always synced.
	ColumnsToSync types.List `tfsdk:"columns_to_sync"`
	// The budget policy id applied to the vector search index
	EffectiveBudgetPolicyId types.String `tfsdk:"effective_budget_policy_id"`

	EffectiveUsagePolicyId types.String `tfsdk:"effective_usage_policy_id"`
	// The columns that contain the embedding source.
	EmbeddingSourceColumns types.List `tfsdk:"embedding_source_columns"`
	// The columns that contain the embedding vectors.
	EmbeddingVectorColumns types.List `tfsdk:"embedding_vector_columns"`
	// [Optional] Name of the Delta table to sync the vector index contents and
	// computed embeddings to.
	EmbeddingWritebackTable types.String `tfsdk:"embedding_writeback_table"`
	// Pipeline execution mode. - `TRIGGERED`: If the pipeline uses the
	// triggered execution mode, the system stops processing after successfully
	// refreshing the source table in the pipeline once, ensuring the table is
	// updated based on the data available when the update started. -
	// `CONTINUOUS`: If the pipeline uses continuous execution, the pipeline
	// processes new data as it arrives in the source table to keep vector index
	// fresh.
	PipelineType types.String `tfsdk:"pipeline_type"`
	// The name of the source table.
	SourceTable types.String `tfsdk:"source_table"`
}

func (to *DeltaSyncVectorIndexSpecRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeltaSyncVectorIndexSpecRequest_SdkV2) {
	if !from.ColumnsToSync.IsNull() && !from.ColumnsToSync.IsUnknown() && to.ColumnsToSync.IsNull() && len(from.ColumnsToSync.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnsToSync, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnsToSync = from.ColumnsToSync
	}
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() && to.EmbeddingSourceColumns.IsNull() && len(from.EmbeddingSourceColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingSourceColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingSourceColumns = from.EmbeddingSourceColumns
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() && to.EmbeddingVectorColumns.IsNull() && len(from.EmbeddingVectorColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingVectorColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingVectorColumns = from.EmbeddingVectorColumns
	}
}

func (to *DeltaSyncVectorIndexSpecRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeltaSyncVectorIndexSpecRequest_SdkV2) {
	if !from.ColumnsToSync.IsNull() && !from.ColumnsToSync.IsUnknown() && to.ColumnsToSync.IsNull() && len(from.ColumnsToSync.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnsToSync, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnsToSync = from.ColumnsToSync
	}
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() && to.EmbeddingSourceColumns.IsNull() && len(from.EmbeddingSourceColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingSourceColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingSourceColumns = from.EmbeddingSourceColumns
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() && to.EmbeddingVectorColumns.IsNull() && len(from.EmbeddingVectorColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingVectorColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingVectorColumns = from.EmbeddingVectorColumns
	}
}

func (m DeltaSyncVectorIndexSpecRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["columns_to_sync"] = attrs["columns_to_sync"].SetOptional()
	attrs["effective_budget_policy_id"] = attrs["effective_budget_policy_id"].SetOptional()
	attrs["effective_usage_policy_id"] = attrs["effective_usage_policy_id"].SetOptional()
	attrs["embedding_source_columns"] = attrs["embedding_source_columns"].SetOptional()
	attrs["embedding_vector_columns"] = attrs["embedding_vector_columns"].SetOptional()
	attrs["embedding_writeback_table"] = attrs["embedding_writeback_table"].SetOptional()
	attrs["pipeline_type"] = attrs["pipeline_type"].SetOptional()
	attrs["source_table"] = attrs["source_table"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeltaSyncVectorIndexSpecRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeltaSyncVectorIndexSpecRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns_to_sync":          reflect.TypeOf(types.String{}),
		"embedding_source_columns": reflect.TypeOf(EmbeddingSourceColumn_SdkV2{}),
		"embedding_vector_columns": reflect.TypeOf(EmbeddingVectorColumn_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaSyncVectorIndexSpecRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeltaSyncVectorIndexSpecRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns_to_sync":            m.ColumnsToSync,
			"effective_budget_policy_id": m.EffectiveBudgetPolicyId,
			"effective_usage_policy_id":  m.EffectiveUsagePolicyId,
			"embedding_source_columns":   m.EmbeddingSourceColumns,
			"embedding_vector_columns":   m.EmbeddingVectorColumns,
			"embedding_writeback_table":  m.EmbeddingWritebackTable,
			"pipeline_type":              m.PipelineType,
			"source_table":               m.SourceTable,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeltaSyncVectorIndexSpecRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns_to_sync": basetypes.ListType{
				ElemType: types.StringType,
			},
			"effective_budget_policy_id": types.StringType,
			"effective_usage_policy_id":  types.StringType,
			"embedding_source_columns": basetypes.ListType{
				ElemType: EmbeddingSourceColumn_SdkV2{}.Type(ctx),
			},
			"embedding_vector_columns": basetypes.ListType{
				ElemType: EmbeddingVectorColumn_SdkV2{}.Type(ctx),
			},
			"embedding_writeback_table": types.StringType,
			"pipeline_type":             types.StringType,
			"source_table":              types.StringType,
		},
	}
}

// GetColumnsToSync returns the value of the ColumnsToSync field in DeltaSyncVectorIndexSpecRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSyncVectorIndexSpecRequest_SdkV2) GetColumnsToSync(ctx context.Context) ([]types.String, bool) {
	if m.ColumnsToSync.IsNull() || m.ColumnsToSync.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ColumnsToSync.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumnsToSync sets the value of the ColumnsToSync field in DeltaSyncVectorIndexSpecRequest_SdkV2.
func (m *DeltaSyncVectorIndexSpecRequest_SdkV2) SetColumnsToSync(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns_to_sync"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ColumnsToSync = types.ListValueMust(t, vs)
}

// GetEmbeddingSourceColumns returns the value of the EmbeddingSourceColumns field in DeltaSyncVectorIndexSpecRequest_SdkV2 as
// a slice of EmbeddingSourceColumn_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSyncVectorIndexSpecRequest_SdkV2) GetEmbeddingSourceColumns(ctx context.Context) ([]EmbeddingSourceColumn_SdkV2, bool) {
	if m.EmbeddingSourceColumns.IsNull() || m.EmbeddingSourceColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingSourceColumn_SdkV2
	d := m.EmbeddingSourceColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingSourceColumns sets the value of the EmbeddingSourceColumns field in DeltaSyncVectorIndexSpecRequest_SdkV2.
func (m *DeltaSyncVectorIndexSpecRequest_SdkV2) SetEmbeddingSourceColumns(ctx context.Context, v []EmbeddingSourceColumn_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_source_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmbeddingSourceColumns = types.ListValueMust(t, vs)
}

// GetEmbeddingVectorColumns returns the value of the EmbeddingVectorColumns field in DeltaSyncVectorIndexSpecRequest_SdkV2 as
// a slice of EmbeddingVectorColumn_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSyncVectorIndexSpecRequest_SdkV2) GetEmbeddingVectorColumns(ctx context.Context) ([]EmbeddingVectorColumn_SdkV2, bool) {
	if m.EmbeddingVectorColumns.IsNull() || m.EmbeddingVectorColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingVectorColumn_SdkV2
	d := m.EmbeddingVectorColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingVectorColumns sets the value of the EmbeddingVectorColumns field in DeltaSyncVectorIndexSpecRequest_SdkV2.
func (m *DeltaSyncVectorIndexSpecRequest_SdkV2) SetEmbeddingVectorColumns(ctx context.Context, v []EmbeddingVectorColumn_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_vector_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmbeddingVectorColumns = types.ListValueMust(t, vs)
}

type DeltaSyncVectorIndexSpecResponse_SdkV2 struct {
	// The budget policy id applied to the vector search index
	EffectiveBudgetPolicyId types.String `tfsdk:"effective_budget_policy_id"`

	EffectiveUsagePolicyId types.String `tfsdk:"effective_usage_policy_id"`
	// The columns that contain the embedding source.
	EmbeddingSourceColumns types.List `tfsdk:"embedding_source_columns"`
	// The columns that contain the embedding vectors.
	EmbeddingVectorColumns types.List `tfsdk:"embedding_vector_columns"`
	// [Optional] Name of the Delta table to sync the vector index contents and
	// computed embeddings to.
	EmbeddingWritebackTable types.String `tfsdk:"embedding_writeback_table"`
	// The ID of the pipeline that is used to sync the index.
	PipelineId types.String `tfsdk:"pipeline_id"`
	// Pipeline execution mode. - `TRIGGERED`: If the pipeline uses the
	// triggered execution mode, the system stops processing after successfully
	// refreshing the source table in the pipeline once, ensuring the table is
	// updated based on the data available when the update started. -
	// `CONTINUOUS`: If the pipeline uses continuous execution, the pipeline
	// processes new data as it arrives in the source table to keep vector index
	// fresh.
	PipelineType types.String `tfsdk:"pipeline_type"`
	// The name of the source table.
	SourceTable types.String `tfsdk:"source_table"`
}

func (to *DeltaSyncVectorIndexSpecResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeltaSyncVectorIndexSpecResponse_SdkV2) {
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() && to.EmbeddingSourceColumns.IsNull() && len(from.EmbeddingSourceColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingSourceColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingSourceColumns = from.EmbeddingSourceColumns
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() && to.EmbeddingVectorColumns.IsNull() && len(from.EmbeddingVectorColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingVectorColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingVectorColumns = from.EmbeddingVectorColumns
	}
}

func (to *DeltaSyncVectorIndexSpecResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeltaSyncVectorIndexSpecResponse_SdkV2) {
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() && to.EmbeddingSourceColumns.IsNull() && len(from.EmbeddingSourceColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingSourceColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingSourceColumns = from.EmbeddingSourceColumns
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() && to.EmbeddingVectorColumns.IsNull() && len(from.EmbeddingVectorColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingVectorColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingVectorColumns = from.EmbeddingVectorColumns
	}
}

func (m DeltaSyncVectorIndexSpecResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["effective_budget_policy_id"] = attrs["effective_budget_policy_id"].SetOptional()
	attrs["effective_usage_policy_id"] = attrs["effective_usage_policy_id"].SetOptional()
	attrs["embedding_source_columns"] = attrs["embedding_source_columns"].SetOptional()
	attrs["embedding_vector_columns"] = attrs["embedding_vector_columns"].SetOptional()
	attrs["embedding_writeback_table"] = attrs["embedding_writeback_table"].SetOptional()
	attrs["pipeline_id"] = attrs["pipeline_id"].SetOptional()
	attrs["pipeline_type"] = attrs["pipeline_type"].SetOptional()
	attrs["source_table"] = attrs["source_table"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeltaSyncVectorIndexSpecResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeltaSyncVectorIndexSpecResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"embedding_source_columns": reflect.TypeOf(EmbeddingSourceColumn_SdkV2{}),
		"embedding_vector_columns": reflect.TypeOf(EmbeddingVectorColumn_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaSyncVectorIndexSpecResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeltaSyncVectorIndexSpecResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"effective_budget_policy_id": m.EffectiveBudgetPolicyId,
			"effective_usage_policy_id":  m.EffectiveUsagePolicyId,
			"embedding_source_columns":   m.EmbeddingSourceColumns,
			"embedding_vector_columns":   m.EmbeddingVectorColumns,
			"embedding_writeback_table":  m.EmbeddingWritebackTable,
			"pipeline_id":                m.PipelineId,
			"pipeline_type":              m.PipelineType,
			"source_table":               m.SourceTable,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeltaSyncVectorIndexSpecResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"effective_budget_policy_id": types.StringType,
			"effective_usage_policy_id":  types.StringType,
			"embedding_source_columns": basetypes.ListType{
				ElemType: EmbeddingSourceColumn_SdkV2{}.Type(ctx),
			},
			"embedding_vector_columns": basetypes.ListType{
				ElemType: EmbeddingVectorColumn_SdkV2{}.Type(ctx),
			},
			"embedding_writeback_table": types.StringType,
			"pipeline_id":               types.StringType,
			"pipeline_type":             types.StringType,
			"source_table":              types.StringType,
		},
	}
}

// GetEmbeddingSourceColumns returns the value of the EmbeddingSourceColumns field in DeltaSyncVectorIndexSpecResponse_SdkV2 as
// a slice of EmbeddingSourceColumn_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSyncVectorIndexSpecResponse_SdkV2) GetEmbeddingSourceColumns(ctx context.Context) ([]EmbeddingSourceColumn_SdkV2, bool) {
	if m.EmbeddingSourceColumns.IsNull() || m.EmbeddingSourceColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingSourceColumn_SdkV2
	d := m.EmbeddingSourceColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingSourceColumns sets the value of the EmbeddingSourceColumns field in DeltaSyncVectorIndexSpecResponse_SdkV2.
func (m *DeltaSyncVectorIndexSpecResponse_SdkV2) SetEmbeddingSourceColumns(ctx context.Context, v []EmbeddingSourceColumn_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_source_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmbeddingSourceColumns = types.ListValueMust(t, vs)
}

// GetEmbeddingVectorColumns returns the value of the EmbeddingVectorColumns field in DeltaSyncVectorIndexSpecResponse_SdkV2 as
// a slice of EmbeddingVectorColumn_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSyncVectorIndexSpecResponse_SdkV2) GetEmbeddingVectorColumns(ctx context.Context) ([]EmbeddingVectorColumn_SdkV2, bool) {
	if m.EmbeddingVectorColumns.IsNull() || m.EmbeddingVectorColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingVectorColumn_SdkV2
	d := m.EmbeddingVectorColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingVectorColumns sets the value of the EmbeddingVectorColumns field in DeltaSyncVectorIndexSpecResponse_SdkV2.
func (m *DeltaSyncVectorIndexSpecResponse_SdkV2) SetEmbeddingVectorColumns(ctx context.Context, v []EmbeddingVectorColumn_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_vector_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmbeddingVectorColumns = types.ListValueMust(t, vs)
}

type DirectAccessVectorIndexSpec_SdkV2 struct {
	// The columns that contain the embedding source. The format should be
	// array[double].
	EmbeddingSourceColumns types.List `tfsdk:"embedding_source_columns"`
	// The columns that contain the embedding vectors. The format should be
	// array[double].
	EmbeddingVectorColumns types.List `tfsdk:"embedding_vector_columns"`
	// The schema of the index in JSON format. Supported types are `integer`,
	// `long`, `float`, `double`, `boolean`, `string`, `date`, `timestamp`.
	// Supported types for vector column: `array<float>`, `array<double>`,`.
	SchemaJson types.String `tfsdk:"schema_json"`
}

func (to *DirectAccessVectorIndexSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DirectAccessVectorIndexSpec_SdkV2) {
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() && to.EmbeddingSourceColumns.IsNull() && len(from.EmbeddingSourceColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingSourceColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingSourceColumns = from.EmbeddingSourceColumns
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() && to.EmbeddingVectorColumns.IsNull() && len(from.EmbeddingVectorColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingVectorColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingVectorColumns = from.EmbeddingVectorColumns
	}
}

func (to *DirectAccessVectorIndexSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DirectAccessVectorIndexSpec_SdkV2) {
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() && to.EmbeddingSourceColumns.IsNull() && len(from.EmbeddingSourceColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingSourceColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingSourceColumns = from.EmbeddingSourceColumns
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() && to.EmbeddingVectorColumns.IsNull() && len(from.EmbeddingVectorColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingVectorColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingVectorColumns = from.EmbeddingVectorColumns
	}
}

func (m DirectAccessVectorIndexSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["embedding_source_columns"] = attrs["embedding_source_columns"].SetOptional()
	attrs["embedding_vector_columns"] = attrs["embedding_vector_columns"].SetOptional()
	attrs["schema_json"] = attrs["schema_json"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DirectAccessVectorIndexSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DirectAccessVectorIndexSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"embedding_source_columns": reflect.TypeOf(EmbeddingSourceColumn_SdkV2{}),
		"embedding_vector_columns": reflect.TypeOf(EmbeddingVectorColumn_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DirectAccessVectorIndexSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m DirectAccessVectorIndexSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding_source_columns": m.EmbeddingSourceColumns,
			"embedding_vector_columns": m.EmbeddingVectorColumns,
			"schema_json":              m.SchemaJson,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DirectAccessVectorIndexSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"embedding_source_columns": basetypes.ListType{
				ElemType: EmbeddingSourceColumn_SdkV2{}.Type(ctx),
			},
			"embedding_vector_columns": basetypes.ListType{
				ElemType: EmbeddingVectorColumn_SdkV2{}.Type(ctx),
			},
			"schema_json": types.StringType,
		},
	}
}

// GetEmbeddingSourceColumns returns the value of the EmbeddingSourceColumns field in DirectAccessVectorIndexSpec_SdkV2 as
// a slice of EmbeddingSourceColumn_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *DirectAccessVectorIndexSpec_SdkV2) GetEmbeddingSourceColumns(ctx context.Context) ([]EmbeddingSourceColumn_SdkV2, bool) {
	if m.EmbeddingSourceColumns.IsNull() || m.EmbeddingSourceColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingSourceColumn_SdkV2
	d := m.EmbeddingSourceColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingSourceColumns sets the value of the EmbeddingSourceColumns field in DirectAccessVectorIndexSpec_SdkV2.
func (m *DirectAccessVectorIndexSpec_SdkV2) SetEmbeddingSourceColumns(ctx context.Context, v []EmbeddingSourceColumn_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_source_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmbeddingSourceColumns = types.ListValueMust(t, vs)
}

// GetEmbeddingVectorColumns returns the value of the EmbeddingVectorColumns field in DirectAccessVectorIndexSpec_SdkV2 as
// a slice of EmbeddingVectorColumn_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *DirectAccessVectorIndexSpec_SdkV2) GetEmbeddingVectorColumns(ctx context.Context) ([]EmbeddingVectorColumn_SdkV2, bool) {
	if m.EmbeddingVectorColumns.IsNull() || m.EmbeddingVectorColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingVectorColumn_SdkV2
	d := m.EmbeddingVectorColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingVectorColumns sets the value of the EmbeddingVectorColumns field in DirectAccessVectorIndexSpec_SdkV2.
func (m *DirectAccessVectorIndexSpec_SdkV2) SetEmbeddingVectorColumns(ctx context.Context, v []EmbeddingVectorColumn_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_vector_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmbeddingVectorColumns = types.ListValueMust(t, vs)
}

type EmbeddingSourceColumn_SdkV2 struct {
	// Name of the embedding model endpoint, used by default for both ingestion
	// and querying.
	EmbeddingModelEndpointName types.String `tfsdk:"embedding_model_endpoint_name"`
	// Name of the embedding model endpoint which, if specified, is used for
	// querying (not ingestion).
	ModelEndpointNameForQuery types.String `tfsdk:"model_endpoint_name_for_query"`
	// Name of the column
	Name types.String `tfsdk:"name"`
}

func (to *EmbeddingSourceColumn_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EmbeddingSourceColumn_SdkV2) {
}

func (to *EmbeddingSourceColumn_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EmbeddingSourceColumn_SdkV2) {
}

func (m EmbeddingSourceColumn_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["embedding_model_endpoint_name"] = attrs["embedding_model_endpoint_name"].SetOptional()
	attrs["model_endpoint_name_for_query"] = attrs["model_endpoint_name_for_query"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EmbeddingSourceColumn.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EmbeddingSourceColumn_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EmbeddingSourceColumn_SdkV2
// only implements ToObjectValue() and Type().
func (m EmbeddingSourceColumn_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding_model_endpoint_name": m.EmbeddingModelEndpointName,
			"model_endpoint_name_for_query": m.ModelEndpointNameForQuery,
			"name":                          m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EmbeddingSourceColumn_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"embedding_model_endpoint_name": types.StringType,
			"model_endpoint_name_for_query": types.StringType,
			"name":                          types.StringType,
		},
	}
}

type EmbeddingVectorColumn_SdkV2 struct {
	// Dimension of the embedding vector
	EmbeddingDimension types.Int64 `tfsdk:"embedding_dimension"`
	// Name of the column
	Name types.String `tfsdk:"name"`
}

func (to *EmbeddingVectorColumn_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EmbeddingVectorColumn_SdkV2) {
}

func (to *EmbeddingVectorColumn_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EmbeddingVectorColumn_SdkV2) {
}

func (m EmbeddingVectorColumn_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["embedding_dimension"] = attrs["embedding_dimension"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EmbeddingVectorColumn.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EmbeddingVectorColumn_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EmbeddingVectorColumn_SdkV2
// only implements ToObjectValue() and Type().
func (m EmbeddingVectorColumn_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding_dimension": m.EmbeddingDimension,
			"name":                m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EmbeddingVectorColumn_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"embedding_dimension": types.Int64Type,
			"name":                types.StringType,
		},
	}
}

type EndpointInfo_SdkV2 struct {
	// Timestamp of endpoint creation
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`
	// Creator of the endpoint
	Creator types.String `tfsdk:"creator"`
	// The custom tags assigned to the endpoint
	CustomTags types.List `tfsdk:"custom_tags"`
	// The budget policy id applied to the endpoint
	EffectiveBudgetPolicyId types.String `tfsdk:"effective_budget_policy_id"`
	// Current status of the endpoint
	EndpointStatus types.List `tfsdk:"endpoint_status"`
	// Type of endpoint
	EndpointType types.String `tfsdk:"endpoint_type"`
	// Unique identifier of the endpoint
	Id types.String `tfsdk:"id"`
	// Timestamp of last update to the endpoint
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// User who last updated the endpoint
	LastUpdatedUser types.String `tfsdk:"last_updated_user"`
	// Name of the vector search endpoint
	Name types.String `tfsdk:"name"`
	// Number of indexes on the endpoint
	NumIndexes types.Int64 `tfsdk:"num_indexes"`
}

func (to *EndpointInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointInfo_SdkV2) {
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
	if !from.EndpointStatus.IsNull() && !from.EndpointStatus.IsUnknown() {
		if toEndpointStatus, ok := to.GetEndpointStatus(ctx); ok {
			if fromEndpointStatus, ok := from.GetEndpointStatus(ctx); ok {
				// Recursively sync the fields of EndpointStatus
				toEndpointStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromEndpointStatus)
				to.SetEndpointStatus(ctx, toEndpointStatus)
			}
		}
	}
}

func (to *EndpointInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointInfo_SdkV2) {
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
	if !from.EndpointStatus.IsNull() && !from.EndpointStatus.IsUnknown() {
		if toEndpointStatus, ok := to.GetEndpointStatus(ctx); ok {
			if fromEndpointStatus, ok := from.GetEndpointStatus(ctx); ok {
				toEndpointStatus.SyncFieldsDuringRead(ctx, fromEndpointStatus)
				to.SetEndpointStatus(ctx, toEndpointStatus)
			}
		}
	}
}

func (m EndpointInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creation_timestamp"] = attrs["creation_timestamp"].SetOptional()
	attrs["creator"] = attrs["creator"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["effective_budget_policy_id"] = attrs["effective_budget_policy_id"].SetOptional()
	attrs["endpoint_status"] = attrs["endpoint_status"].SetOptional()
	attrs["endpoint_status"] = attrs["endpoint_status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["endpoint_type"] = attrs["endpoint_type"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetOptional()
	attrs["last_updated_user"] = attrs["last_updated_user"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["num_indexes"] = attrs["num_indexes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_tags":     reflect.TypeOf(CustomTag_SdkV2{}),
		"endpoint_status": reflect.TypeOf(EndpointStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m EndpointInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_timestamp":         m.CreationTimestamp,
			"creator":                    m.Creator,
			"custom_tags":                m.CustomTags,
			"effective_budget_policy_id": m.EffectiveBudgetPolicyId,
			"endpoint_status":            m.EndpointStatus,
			"endpoint_type":              m.EndpointType,
			"id":                         m.Id,
			"last_updated_timestamp":     m.LastUpdatedTimestamp,
			"last_updated_user":          m.LastUpdatedUser,
			"name":                       m.Name,
			"num_indexes":                m.NumIndexes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_timestamp": types.Int64Type,
			"creator":            types.StringType,
			"custom_tags": basetypes.ListType{
				ElemType: CustomTag_SdkV2{}.Type(ctx),
			},
			"effective_budget_policy_id": types.StringType,
			"endpoint_status": basetypes.ListType{
				ElemType: EndpointStatus_SdkV2{}.Type(ctx),
			},
			"endpoint_type":          types.StringType,
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"last_updated_user":      types.StringType,
			"name":                   types.StringType,
			"num_indexes":            types.Int64Type,
		},
	}
}

// GetCustomTags returns the value of the CustomTags field in EndpointInfo_SdkV2 as
// a slice of CustomTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointInfo_SdkV2) GetCustomTags(ctx context.Context) ([]CustomTag_SdkV2, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v []CustomTag_SdkV2
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in EndpointInfo_SdkV2.
func (m *EndpointInfo_SdkV2) SetCustomTags(ctx context.Context, v []CustomTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.ListValueMust(t, vs)
}

// GetEndpointStatus returns the value of the EndpointStatus field in EndpointInfo_SdkV2 as
// a EndpointStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointInfo_SdkV2) GetEndpointStatus(ctx context.Context) (EndpointStatus_SdkV2, bool) {
	var e EndpointStatus_SdkV2
	if m.EndpointStatus.IsNull() || m.EndpointStatus.IsUnknown() {
		return e, false
	}
	var v []EndpointStatus_SdkV2
	d := m.EndpointStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEndpointStatus sets the value of the EndpointStatus field in EndpointInfo_SdkV2.
func (m *EndpointInfo_SdkV2) SetEndpointStatus(ctx context.Context, v EndpointStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["endpoint_status"]
	m.EndpointStatus = types.ListValueMust(t, vs)
}

// Status information of an endpoint
type EndpointStatus_SdkV2 struct {
	// Additional status message
	Message types.String `tfsdk:"message"`
	// Current state of the endpoint
	State types.String `tfsdk:"state"`
}

func (to *EndpointStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointStatus_SdkV2) {
}

func (to *EndpointStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointStatus_SdkV2) {
}

func (m EndpointStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["message"] = attrs["message"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()

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
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointStatus_SdkV2
// only implements ToObjectValue() and Type().
func (m EndpointStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": m.Message,
			"state":   m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"state":   types.StringType,
		},
	}
}

type GetEndpointRequest_SdkV2 struct {
	// Name of the endpoint
	EndpointName types.String `tfsdk:"-"`
}

func (to *GetEndpointRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEndpointRequest_SdkV2) {
}

func (to *GetEndpointRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetEndpointRequest_SdkV2) {
}

func (m GetEndpointRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint_name"] = attrs["endpoint_name"].SetRequired()

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
			"endpoint_name": m.EndpointName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_name": types.StringType,
		},
	}
}

type GetIndexRequest_SdkV2 struct {
	// If true, the URL returned for the index is guaranteed to be compatible
	// with the reranker. Currently this means we return the CP URL regardless
	// of how the index is being accessed. If not set or set to false, the URL
	// may still be compatible with the reranker depending on what URL we
	// return.
	EnsureRerankerCompatible types.Bool `tfsdk:"-"`
	// Name of the index
	IndexName types.String `tfsdk:"-"`
}

func (to *GetIndexRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetIndexRequest_SdkV2) {
}

func (to *GetIndexRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetIndexRequest_SdkV2) {
}

func (m GetIndexRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["index_name"] = attrs["index_name"].SetRequired()
	attrs["ensure_reranker_compatible"] = attrs["ensure_reranker_compatible"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ensure_reranker_compatible": m.EnsureRerankerCompatible,
			"index_name":                 m.IndexName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ensure_reranker_compatible": types.BoolType,
			"index_name":                 types.StringType,
		},
	}
}

type ListEndpointResponse_SdkV2 struct {
	// An array of Endpoint objects
	Endpoints types.List `tfsdk:"endpoints"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListEndpointResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEndpointResponse_SdkV2) {
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() && to.Endpoints.IsNull() && len(from.Endpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Endpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Endpoints = from.Endpoints
	}
}

func (to *ListEndpointResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListEndpointResponse_SdkV2) {
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() && to.Endpoints.IsNull() && len(from.Endpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Endpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Endpoints = from.Endpoints
	}
}

func (m ListEndpointResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoints"] = attrs["endpoints"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListEndpointResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListEndpointResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoints": reflect.TypeOf(EndpointInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListEndpointResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoints":       m.Endpoints,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListEndpointResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoints": basetypes.ListType{
				ElemType: EndpointInfo_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetEndpoints returns the value of the Endpoints field in ListEndpointResponse_SdkV2 as
// a slice of EndpointInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListEndpointResponse_SdkV2) GetEndpoints(ctx context.Context) ([]EndpointInfo_SdkV2, bool) {
	if m.Endpoints.IsNull() || m.Endpoints.IsUnknown() {
		return nil, false
	}
	var v []EndpointInfo_SdkV2
	d := m.Endpoints.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEndpoints sets the value of the Endpoints field in ListEndpointResponse_SdkV2.
func (m *ListEndpointResponse_SdkV2) SetEndpoints(ctx context.Context, v []EndpointInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["endpoints"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Endpoints = types.ListValueMust(t, vs)
}

type ListEndpointsRequest_SdkV2 struct {
	// Token for pagination
	PageToken types.String `tfsdk:"-"`
}

func (to *ListEndpointsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEndpointsRequest_SdkV2) {
}

func (to *ListEndpointsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListEndpointsRequest_SdkV2) {
}

func (m ListEndpointsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()

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
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListEndpointsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_token": types.StringType,
		},
	}
}

type ListIndexesRequest_SdkV2 struct {
	// Name of the endpoint
	EndpointName types.String `tfsdk:"-"`
	// Token for pagination
	PageToken types.String `tfsdk:"-"`
}

func (to *ListIndexesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListIndexesRequest_SdkV2) {
}

func (to *ListIndexesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListIndexesRequest_SdkV2) {
}

func (m ListIndexesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint_name"] = attrs["endpoint_name"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListIndexesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListIndexesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListIndexesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListIndexesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_name": m.EndpointName,
			"page_token":    m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListIndexesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_name": types.StringType,
			"page_token":    types.StringType,
		},
	}
}

type ListValue_SdkV2 struct {
	// Repeated field of dynamically typed values.
	Values types.List `tfsdk:"values"`
}

func (to *ListValue_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListValue_SdkV2) {
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (to *ListValue_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListValue_SdkV2) {
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (m ListValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["values"] = attrs["values"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"values": reflect.TypeOf(Value_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListValue_SdkV2
// only implements ToObjectValue() and Type().
func (m ListValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"values": m.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListValue_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"values": basetypes.ListType{
				ElemType: Value_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetValues returns the value of the Values field in ListValue_SdkV2 as
// a slice of Value_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListValue_SdkV2) GetValues(ctx context.Context) ([]Value_SdkV2, bool) {
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

// SetValues sets the value of the Values field in ListValue_SdkV2.
func (m *ListValue_SdkV2) SetValues(ctx context.Context, v []Value_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Values = types.ListValueMust(t, vs)
}

type ListVectorIndexesResponse_SdkV2 struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`

	VectorIndexes types.List `tfsdk:"vector_indexes"`
}

func (to *ListVectorIndexesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListVectorIndexesResponse_SdkV2) {
	if !from.VectorIndexes.IsNull() && !from.VectorIndexes.IsUnknown() && to.VectorIndexes.IsNull() && len(from.VectorIndexes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for VectorIndexes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.VectorIndexes = from.VectorIndexes
	}
}

func (to *ListVectorIndexesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListVectorIndexesResponse_SdkV2) {
	if !from.VectorIndexes.IsNull() && !from.VectorIndexes.IsUnknown() && to.VectorIndexes.IsNull() && len(from.VectorIndexes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for VectorIndexes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.VectorIndexes = from.VectorIndexes
	}
}

func (m ListVectorIndexesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["vector_indexes"] = attrs["vector_indexes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListVectorIndexesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListVectorIndexesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"vector_indexes": reflect.TypeOf(MiniVectorIndex_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListVectorIndexesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListVectorIndexesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"vector_indexes":  m.VectorIndexes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListVectorIndexesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"vector_indexes": basetypes.ListType{
				ElemType: MiniVectorIndex_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetVectorIndexes returns the value of the VectorIndexes field in ListVectorIndexesResponse_SdkV2 as
// a slice of MiniVectorIndex_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListVectorIndexesResponse_SdkV2) GetVectorIndexes(ctx context.Context) ([]MiniVectorIndex_SdkV2, bool) {
	if m.VectorIndexes.IsNull() || m.VectorIndexes.IsUnknown() {
		return nil, false
	}
	var v []MiniVectorIndex_SdkV2
	d := m.VectorIndexes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVectorIndexes sets the value of the VectorIndexes field in ListVectorIndexesResponse_SdkV2.
func (m *ListVectorIndexesResponse_SdkV2) SetVectorIndexes(ctx context.Context, v []MiniVectorIndex_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["vector_indexes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.VectorIndexes = types.ListValueMust(t, vs)
}

// Key-value pair.
type MapStringValueEntry_SdkV2 struct {
	// Column name.
	Key types.String `tfsdk:"key"`
	// Column value, nullable.
	Value types.List `tfsdk:"value"`
}

func (to *MapStringValueEntry_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MapStringValueEntry_SdkV2) {
	if !from.Value.IsNull() && !from.Value.IsUnknown() {
		if toValue, ok := to.GetValue(ctx); ok {
			if fromValue, ok := from.GetValue(ctx); ok {
				// Recursively sync the fields of Value
				toValue.SyncFieldsDuringCreateOrUpdate(ctx, fromValue)
				to.SetValue(ctx, toValue)
			}
		}
	}
}

func (to *MapStringValueEntry_SdkV2) SyncFieldsDuringRead(ctx context.Context, from MapStringValueEntry_SdkV2) {
	if !from.Value.IsNull() && !from.Value.IsUnknown() {
		if toValue, ok := to.GetValue(ctx); ok {
			if fromValue, ok := from.GetValue(ctx); ok {
				toValue.SyncFieldsDuringRead(ctx, fromValue)
				to.SetValue(ctx, toValue)
			}
		}
	}
}

func (m MapStringValueEntry_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()
	attrs["value"] = attrs["value"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MapStringValueEntry.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m MapStringValueEntry_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"value": reflect.TypeOf(Value_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MapStringValueEntry_SdkV2
// only implements ToObjectValue() and Type().
func (m MapStringValueEntry_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m MapStringValueEntry_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key": types.StringType,
			"value": basetypes.ListType{
				ElemType: Value_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetValue returns the value of the Value field in MapStringValueEntry_SdkV2 as
// a Value_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *MapStringValueEntry_SdkV2) GetValue(ctx context.Context) (Value_SdkV2, bool) {
	var e Value_SdkV2
	if m.Value.IsNull() || m.Value.IsUnknown() {
		return e, false
	}
	var v []Value_SdkV2
	d := m.Value.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetValue sets the value of the Value field in MapStringValueEntry_SdkV2.
func (m *MapStringValueEntry_SdkV2) SetValue(ctx context.Context, v Value_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["value"]
	m.Value = types.ListValueMust(t, vs)
}

type MiniVectorIndex_SdkV2 struct {
	// The user who created the index.
	Creator types.String `tfsdk:"creator"`
	// Name of the endpoint associated with the index
	EndpointName types.String `tfsdk:"endpoint_name"`

	IndexType types.String `tfsdk:"index_type"`
	// Name of the index
	Name types.String `tfsdk:"name"`
	// Primary key of the index
	PrimaryKey types.String `tfsdk:"primary_key"`
}

func (to *MiniVectorIndex_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MiniVectorIndex_SdkV2) {
}

func (to *MiniVectorIndex_SdkV2) SyncFieldsDuringRead(ctx context.Context, from MiniVectorIndex_SdkV2) {
}

func (m MiniVectorIndex_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creator"] = attrs["creator"].SetOptional()
	attrs["endpoint_name"] = attrs["endpoint_name"].SetOptional()
	attrs["index_type"] = attrs["index_type"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["primary_key"] = attrs["primary_key"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MiniVectorIndex.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m MiniVectorIndex_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MiniVectorIndex_SdkV2
// only implements ToObjectValue() and Type().
func (m MiniVectorIndex_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creator":       m.Creator,
			"endpoint_name": m.EndpointName,
			"index_type":    m.IndexType,
			"name":          m.Name,
			"primary_key":   m.PrimaryKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (m MiniVectorIndex_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creator":       types.StringType,
			"endpoint_name": types.StringType,
			"index_type":    types.StringType,
			"name":          types.StringType,
			"primary_key":   types.StringType,
		},
	}
}

type PatchEndpointBudgetPolicyRequest_SdkV2 struct {
	// The budget policy id to be applied (hima-sheth) TODO: remove this once
	// we've migrated to usage policies
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// Name of the vector search endpoint
	EndpointName types.String `tfsdk:"-"`
}

func (to *PatchEndpointBudgetPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PatchEndpointBudgetPolicyRequest_SdkV2) {
}

func (to *PatchEndpointBudgetPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PatchEndpointBudgetPolicyRequest_SdkV2) {
}

func (m PatchEndpointBudgetPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetRequired()
	attrs["endpoint_name"] = attrs["endpoint_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PatchEndpointBudgetPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PatchEndpointBudgetPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchEndpointBudgetPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m PatchEndpointBudgetPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget_policy_id": m.BudgetPolicyId,
			"endpoint_name":    m.EndpointName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PatchEndpointBudgetPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_policy_id": types.StringType,
			"endpoint_name":    types.StringType,
		},
	}
}

type PatchEndpointBudgetPolicyResponse_SdkV2 struct {
	// The budget policy applied to the vector search endpoint.
	EffectiveBudgetPolicyId types.String `tfsdk:"effective_budget_policy_id"`
}

func (to *PatchEndpointBudgetPolicyResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PatchEndpointBudgetPolicyResponse_SdkV2) {
}

func (to *PatchEndpointBudgetPolicyResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PatchEndpointBudgetPolicyResponse_SdkV2) {
}

func (m PatchEndpointBudgetPolicyResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["effective_budget_policy_id"] = attrs["effective_budget_policy_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PatchEndpointBudgetPolicyResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PatchEndpointBudgetPolicyResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchEndpointBudgetPolicyResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m PatchEndpointBudgetPolicyResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"effective_budget_policy_id": m.EffectiveBudgetPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PatchEndpointBudgetPolicyResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"effective_budget_policy_id": types.StringType,
		},
	}
}

// Request payload for getting next page of results.
type QueryVectorIndexNextPageRequest_SdkV2 struct {
	// Name of the endpoint.
	EndpointName types.String `tfsdk:"endpoint_name"`
	// Name of the vector index to query.
	IndexName types.String `tfsdk:"-"`
	// Page token returned from previous `QueryVectorIndex` or
	// `QueryVectorIndexNextPage` API.
	PageToken types.String `tfsdk:"page_token"`
}

func (to *QueryVectorIndexNextPageRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryVectorIndexNextPageRequest_SdkV2) {
}

func (to *QueryVectorIndexNextPageRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from QueryVectorIndexNextPageRequest_SdkV2) {
}

func (m QueryVectorIndexNextPageRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint_name"] = attrs["endpoint_name"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["index_name"] = attrs["index_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryVectorIndexNextPageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m QueryVectorIndexNextPageRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryVectorIndexNextPageRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m QueryVectorIndexNextPageRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_name": m.EndpointName,
			"index_name":    m.IndexName,
			"page_token":    m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryVectorIndexNextPageRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_name": types.StringType,
			"index_name":    types.StringType,
			"page_token":    types.StringType,
		},
	}
}

type QueryVectorIndexRequest_SdkV2 struct {
	// List of column names to include in the response.
	Columns types.List `tfsdk:"columns"`
	// Column names used to retrieve data to send to the reranker.
	ColumnsToRerank types.List `tfsdk:"columns_to_rerank"`
	// JSON string representing query filters.
	//
	// Example filters:
	//
	// - `{"id <": 5}`: Filter for id less than 5. - `{"id >": 5}`: Filter for
	// id greater than 5. - `{"id <=": 5}`: Filter for id less than equal to 5.
	// - `{"id >=": 5}`: Filter for id greater than equal to 5. - `{"id": 5}`:
	// Filter for id equal to 5.
	FiltersJson types.String `tfsdk:"filters_json"`
	// Name of the vector index to query.
	IndexName types.String `tfsdk:"-"`
	// Number of results to return. Defaults to 10.
	NumResults types.Int64 `tfsdk:"num_results"`
	// Query text. Required for Delta Sync Index using model endpoint.
	QueryText types.String `tfsdk:"query_text"`
	// The query type to use. Choices are `ANN` and `HYBRID`. Defaults to `ANN`.
	QueryType types.String `tfsdk:"query_type"`
	// Query vector. Required for Direct Vector Access Index and Delta Sync
	// Index using self-managed vectors.
	QueryVector types.List `tfsdk:"query_vector"`

	Reranker types.List `tfsdk:"reranker"`
	// Threshold for the approximate nearest neighbor search. Defaults to 0.0.
	ScoreThreshold types.Float64 `tfsdk:"score_threshold"`
}

func (to *QueryVectorIndexRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryVectorIndexRequest_SdkV2) {
	if !from.ColumnsToRerank.IsNull() && !from.ColumnsToRerank.IsUnknown() && to.ColumnsToRerank.IsNull() && len(from.ColumnsToRerank.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnsToRerank, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnsToRerank = from.ColumnsToRerank
	}
	if !from.QueryVector.IsNull() && !from.QueryVector.IsUnknown() && to.QueryVector.IsNull() && len(from.QueryVector.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for QueryVector, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.QueryVector = from.QueryVector
	}
	if !from.Reranker.IsNull() && !from.Reranker.IsUnknown() {
		if toReranker, ok := to.GetReranker(ctx); ok {
			if fromReranker, ok := from.GetReranker(ctx); ok {
				// Recursively sync the fields of Reranker
				toReranker.SyncFieldsDuringCreateOrUpdate(ctx, fromReranker)
				to.SetReranker(ctx, toReranker)
			}
		}
	}
}

func (to *QueryVectorIndexRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from QueryVectorIndexRequest_SdkV2) {
	if !from.ColumnsToRerank.IsNull() && !from.ColumnsToRerank.IsUnknown() && to.ColumnsToRerank.IsNull() && len(from.ColumnsToRerank.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnsToRerank, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnsToRerank = from.ColumnsToRerank
	}
	if !from.QueryVector.IsNull() && !from.QueryVector.IsUnknown() && to.QueryVector.IsNull() && len(from.QueryVector.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for QueryVector, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.QueryVector = from.QueryVector
	}
	if !from.Reranker.IsNull() && !from.Reranker.IsUnknown() {
		if toReranker, ok := to.GetReranker(ctx); ok {
			if fromReranker, ok := from.GetReranker(ctx); ok {
				toReranker.SyncFieldsDuringRead(ctx, fromReranker)
				to.SetReranker(ctx, toReranker)
			}
		}
	}
}

func (m QueryVectorIndexRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["columns"] = attrs["columns"].SetRequired()
	attrs["columns_to_rerank"] = attrs["columns_to_rerank"].SetOptional()
	attrs["filters_json"] = attrs["filters_json"].SetOptional()
	attrs["num_results"] = attrs["num_results"].SetOptional()
	attrs["query_text"] = attrs["query_text"].SetOptional()
	attrs["query_type"] = attrs["query_type"].SetOptional()
	attrs["query_vector"] = attrs["query_vector"].SetOptional()
	attrs["reranker"] = attrs["reranker"].SetOptional()
	attrs["reranker"] = attrs["reranker"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["score_threshold"] = attrs["score_threshold"].SetOptional()
	attrs["index_name"] = attrs["index_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryVectorIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m QueryVectorIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns":           reflect.TypeOf(types.String{}),
		"columns_to_rerank": reflect.TypeOf(types.String{}),
		"query_vector":      reflect.TypeOf(types.Float64{}),
		"reranker":          reflect.TypeOf(RerankerConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryVectorIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m QueryVectorIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns":           m.Columns,
			"columns_to_rerank": m.ColumnsToRerank,
			"filters_json":      m.FiltersJson,
			"index_name":        m.IndexName,
			"num_results":       m.NumResults,
			"query_text":        m.QueryText,
			"query_type":        m.QueryType,
			"query_vector":      m.QueryVector,
			"reranker":          m.Reranker,
			"score_threshold":   m.ScoreThreshold,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryVectorIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"columns_to_rerank": basetypes.ListType{
				ElemType: types.StringType,
			},
			"filters_json": types.StringType,
			"index_name":   types.StringType,
			"num_results":  types.Int64Type,
			"query_text":   types.StringType,
			"query_type":   types.StringType,
			"query_vector": basetypes.ListType{
				ElemType: types.Float64Type,
			},
			"reranker": basetypes.ListType{
				ElemType: RerankerConfig_SdkV2{}.Type(ctx),
			},
			"score_threshold": types.Float64Type,
		},
	}
}

// GetColumns returns the value of the Columns field in QueryVectorIndexRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryVectorIndexRequest_SdkV2) GetColumns(ctx context.Context) ([]types.String, bool) {
	if m.Columns.IsNull() || m.Columns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in QueryVectorIndexRequest_SdkV2.
func (m *QueryVectorIndexRequest_SdkV2) SetColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Columns = types.ListValueMust(t, vs)
}

// GetColumnsToRerank returns the value of the ColumnsToRerank field in QueryVectorIndexRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryVectorIndexRequest_SdkV2) GetColumnsToRerank(ctx context.Context) ([]types.String, bool) {
	if m.ColumnsToRerank.IsNull() || m.ColumnsToRerank.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ColumnsToRerank.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumnsToRerank sets the value of the ColumnsToRerank field in QueryVectorIndexRequest_SdkV2.
func (m *QueryVectorIndexRequest_SdkV2) SetColumnsToRerank(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns_to_rerank"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ColumnsToRerank = types.ListValueMust(t, vs)
}

// GetQueryVector returns the value of the QueryVector field in QueryVectorIndexRequest_SdkV2 as
// a slice of types.Float64 values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryVectorIndexRequest_SdkV2) GetQueryVector(ctx context.Context) ([]types.Float64, bool) {
	if m.QueryVector.IsNull() || m.QueryVector.IsUnknown() {
		return nil, false
	}
	var v []types.Float64
	d := m.QueryVector.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQueryVector sets the value of the QueryVector field in QueryVectorIndexRequest_SdkV2.
func (m *QueryVectorIndexRequest_SdkV2) SetQueryVector(ctx context.Context, v []types.Float64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["query_vector"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.QueryVector = types.ListValueMust(t, vs)
}

// GetReranker returns the value of the Reranker field in QueryVectorIndexRequest_SdkV2 as
// a RerankerConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryVectorIndexRequest_SdkV2) GetReranker(ctx context.Context) (RerankerConfig_SdkV2, bool) {
	var e RerankerConfig_SdkV2
	if m.Reranker.IsNull() || m.Reranker.IsUnknown() {
		return e, false
	}
	var v []RerankerConfig_SdkV2
	d := m.Reranker.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetReranker sets the value of the Reranker field in QueryVectorIndexRequest_SdkV2.
func (m *QueryVectorIndexRequest_SdkV2) SetReranker(ctx context.Context, v RerankerConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["reranker"]
	m.Reranker = types.ListValueMust(t, vs)
}

type QueryVectorIndexResponse_SdkV2 struct {
	// Metadata about the result set.
	Manifest types.List `tfsdk:"manifest"`
	// [Optional] Token that can be used in `QueryVectorIndexNextPage` API to
	// get next page of results. If more than 1000 results satisfy the query,
	// they are returned in groups of 1000. Empty value means no more results.
	// The maximum number of results that can be returned is 10,000.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Data returned in the query result.
	Result types.List `tfsdk:"result"`
}

func (to *QueryVectorIndexResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryVectorIndexResponse_SdkV2) {
	if !from.Manifest.IsNull() && !from.Manifest.IsUnknown() {
		if toManifest, ok := to.GetManifest(ctx); ok {
			if fromManifest, ok := from.GetManifest(ctx); ok {
				// Recursively sync the fields of Manifest
				toManifest.SyncFieldsDuringCreateOrUpdate(ctx, fromManifest)
				to.SetManifest(ctx, toManifest)
			}
		}
	}
	if !from.Result.IsNull() && !from.Result.IsUnknown() {
		if toResult, ok := to.GetResult(ctx); ok {
			if fromResult, ok := from.GetResult(ctx); ok {
				// Recursively sync the fields of Result
				toResult.SyncFieldsDuringCreateOrUpdate(ctx, fromResult)
				to.SetResult(ctx, toResult)
			}
		}
	}
}

func (to *QueryVectorIndexResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from QueryVectorIndexResponse_SdkV2) {
	if !from.Manifest.IsNull() && !from.Manifest.IsUnknown() {
		if toManifest, ok := to.GetManifest(ctx); ok {
			if fromManifest, ok := from.GetManifest(ctx); ok {
				toManifest.SyncFieldsDuringRead(ctx, fromManifest)
				to.SetManifest(ctx, toManifest)
			}
		}
	}
	if !from.Result.IsNull() && !from.Result.IsUnknown() {
		if toResult, ok := to.GetResult(ctx); ok {
			if fromResult, ok := from.GetResult(ctx); ok {
				toResult.SyncFieldsDuringRead(ctx, fromResult)
				to.SetResult(ctx, toResult)
			}
		}
	}
}

func (m QueryVectorIndexResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["manifest"] = attrs["manifest"].SetOptional()
	attrs["manifest"] = attrs["manifest"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["result"] = attrs["result"].SetOptional()
	attrs["result"] = attrs["result"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryVectorIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m QueryVectorIndexResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"manifest": reflect.TypeOf(ResultManifest_SdkV2{}),
		"result":   reflect.TypeOf(ResultData_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryVectorIndexResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m QueryVectorIndexResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"manifest":        m.Manifest,
			"next_page_token": m.NextPageToken,
			"result":          m.Result,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryVectorIndexResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"manifest": basetypes.ListType{
				ElemType: ResultManifest_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
			"result": basetypes.ListType{
				ElemType: ResultData_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetManifest returns the value of the Manifest field in QueryVectorIndexResponse_SdkV2 as
// a ResultManifest_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryVectorIndexResponse_SdkV2) GetManifest(ctx context.Context) (ResultManifest_SdkV2, bool) {
	var e ResultManifest_SdkV2
	if m.Manifest.IsNull() || m.Manifest.IsUnknown() {
		return e, false
	}
	var v []ResultManifest_SdkV2
	d := m.Manifest.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetManifest sets the value of the Manifest field in QueryVectorIndexResponse_SdkV2.
func (m *QueryVectorIndexResponse_SdkV2) SetManifest(ctx context.Context, v ResultManifest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["manifest"]
	m.Manifest = types.ListValueMust(t, vs)
}

// GetResult returns the value of the Result field in QueryVectorIndexResponse_SdkV2 as
// a ResultData_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryVectorIndexResponse_SdkV2) GetResult(ctx context.Context) (ResultData_SdkV2, bool) {
	var e ResultData_SdkV2
	if m.Result.IsNull() || m.Result.IsUnknown() {
		return e, false
	}
	var v []ResultData_SdkV2
	d := m.Result.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetResult sets the value of the Result field in QueryVectorIndexResponse_SdkV2.
func (m *QueryVectorIndexResponse_SdkV2) SetResult(ctx context.Context, v ResultData_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["result"]
	m.Result = types.ListValueMust(t, vs)
}

type RerankerConfig_SdkV2 struct {
	Model types.String `tfsdk:"model"`

	Parameters types.List `tfsdk:"parameters"`
}

func (to *RerankerConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RerankerConfig_SdkV2) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() {
		if toParameters, ok := to.GetParameters(ctx); ok {
			if fromParameters, ok := from.GetParameters(ctx); ok {
				// Recursively sync the fields of Parameters
				toParameters.SyncFieldsDuringCreateOrUpdate(ctx, fromParameters)
				to.SetParameters(ctx, toParameters)
			}
		}
	}
}

func (to *RerankerConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RerankerConfig_SdkV2) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() {
		if toParameters, ok := to.GetParameters(ctx); ok {
			if fromParameters, ok := from.GetParameters(ctx); ok {
				toParameters.SyncFieldsDuringRead(ctx, fromParameters)
				to.SetParameters(ctx, toParameters)
			}
		}
	}
}

func (m RerankerConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model"] = attrs["model"].SetOptional()
	attrs["parameters"] = attrs["parameters"].SetOptional()
	attrs["parameters"] = attrs["parameters"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RerankerConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RerankerConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(RerankerConfigRerankerParameters_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RerankerConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m RerankerConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model":      m.Model,
			"parameters": m.Parameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RerankerConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model": types.StringType,
			"parameters": basetypes.ListType{
				ElemType: RerankerConfigRerankerParameters_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetParameters returns the value of the Parameters field in RerankerConfig_SdkV2 as
// a RerankerConfigRerankerParameters_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *RerankerConfig_SdkV2) GetParameters(ctx context.Context) (RerankerConfigRerankerParameters_SdkV2, bool) {
	var e RerankerConfigRerankerParameters_SdkV2
	if m.Parameters.IsNull() || m.Parameters.IsUnknown() {
		return e, false
	}
	var v []RerankerConfigRerankerParameters_SdkV2
	d := m.Parameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetParameters sets the value of the Parameters field in RerankerConfig_SdkV2.
func (m *RerankerConfig_SdkV2) SetParameters(ctx context.Context, v RerankerConfigRerankerParameters_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["parameters"]
	m.Parameters = types.ListValueMust(t, vs)
}

type RerankerConfigRerankerParameters_SdkV2 struct {
	ColumnsToRerank types.List `tfsdk:"columns_to_rerank"`
}

func (to *RerankerConfigRerankerParameters_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RerankerConfigRerankerParameters_SdkV2) {
	if !from.ColumnsToRerank.IsNull() && !from.ColumnsToRerank.IsUnknown() && to.ColumnsToRerank.IsNull() && len(from.ColumnsToRerank.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnsToRerank, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnsToRerank = from.ColumnsToRerank
	}
}

func (to *RerankerConfigRerankerParameters_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RerankerConfigRerankerParameters_SdkV2) {
	if !from.ColumnsToRerank.IsNull() && !from.ColumnsToRerank.IsUnknown() && to.ColumnsToRerank.IsNull() && len(from.ColumnsToRerank.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnsToRerank, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnsToRerank = from.ColumnsToRerank
	}
}

func (m RerankerConfigRerankerParameters_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["columns_to_rerank"] = attrs["columns_to_rerank"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RerankerConfigRerankerParameters.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RerankerConfigRerankerParameters_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns_to_rerank": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RerankerConfigRerankerParameters_SdkV2
// only implements ToObjectValue() and Type().
func (m RerankerConfigRerankerParameters_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns_to_rerank": m.ColumnsToRerank,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RerankerConfigRerankerParameters_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns_to_rerank": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetColumnsToRerank returns the value of the ColumnsToRerank field in RerankerConfigRerankerParameters_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *RerankerConfigRerankerParameters_SdkV2) GetColumnsToRerank(ctx context.Context) ([]types.String, bool) {
	if m.ColumnsToRerank.IsNull() || m.ColumnsToRerank.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ColumnsToRerank.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumnsToRerank sets the value of the ColumnsToRerank field in RerankerConfigRerankerParameters_SdkV2.
func (m *RerankerConfigRerankerParameters_SdkV2) SetColumnsToRerank(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns_to_rerank"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ColumnsToRerank = types.ListValueMust(t, vs)
}

// Data returned in the query result.
type ResultData_SdkV2 struct {
	// Data rows returned in the query.
	DataArray types.List `tfsdk:"data_array"`
	// Number of rows in the result set.
	RowCount types.Int64 `tfsdk:"row_count"`
}

func (to *ResultData_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResultData_SdkV2) {
	if !from.DataArray.IsNull() && !from.DataArray.IsUnknown() && to.DataArray.IsNull() && len(from.DataArray.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DataArray, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DataArray = from.DataArray
	}
}

func (to *ResultData_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResultData_SdkV2) {
	if !from.DataArray.IsNull() && !from.DataArray.IsUnknown() && to.DataArray.IsNull() && len(from.DataArray.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DataArray, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DataArray = from.DataArray
	}
}

func (m ResultData_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["data_array"] = attrs["data_array"].SetOptional()
	attrs["row_count"] = attrs["row_count"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResultData.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ResultData_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_array": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultData_SdkV2
// only implements ToObjectValue() and Type().
func (m ResultData_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_array": m.DataArray,
			"row_count":  m.RowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResultData_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_array": basetypes.ListType{
				ElemType: basetypes.ListType{
					ElemType: types.StringType,
				},
			},
			"row_count": types.Int64Type,
		},
	}
}

// GetDataArray returns the value of the DataArray field in ResultData_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ResultData_SdkV2) GetDataArray(ctx context.Context) ([]types.String, bool) {
	if m.DataArray.IsNull() || m.DataArray.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.DataArray.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataArray sets the value of the DataArray field in ResultData_SdkV2.
func (m *ResultData_SdkV2) SetDataArray(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["data_array"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DataArray = types.ListValueMust(t, vs)
}

// Metadata about the result set.
type ResultManifest_SdkV2 struct {
	// Number of columns in the result set.
	ColumnCount types.Int64 `tfsdk:"column_count"`
	// Information about each column in the result set.
	Columns types.List `tfsdk:"columns"`
}

func (to *ResultManifest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResultManifest_SdkV2) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
}

func (to *ResultManifest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ResultManifest_SdkV2) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
}

func (m ResultManifest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["column_count"] = attrs["column_count"].SetOptional()
	attrs["columns"] = attrs["columns"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResultManifest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ResultManifest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(ColumnInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultManifest_SdkV2
// only implements ToObjectValue() and Type().
func (m ResultManifest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column_count": m.ColumnCount,
			"columns":      m.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResultManifest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"column_count": types.Int64Type,
			"columns": basetypes.ListType{
				ElemType: ColumnInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetColumns returns the value of the Columns field in ResultManifest_SdkV2 as
// a slice of ColumnInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ResultManifest_SdkV2) GetColumns(ctx context.Context) ([]ColumnInfo_SdkV2, bool) {
	if m.Columns.IsNull() || m.Columns.IsUnknown() {
		return nil, false
	}
	var v []ColumnInfo_SdkV2
	d := m.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in ResultManifest_SdkV2.
func (m *ResultManifest_SdkV2) SetColumns(ctx context.Context, v []ColumnInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Columns = types.ListValueMust(t, vs)
}

type ScanVectorIndexRequest_SdkV2 struct {
	// Name of the vector index to scan.
	IndexName types.String `tfsdk:"-"`
	// Primary key of the last entry returned in the previous scan.
	LastPrimaryKey types.String `tfsdk:"last_primary_key"`
	// Number of results to return. Defaults to 10.
	NumResults types.Int64 `tfsdk:"num_results"`
}

func (to *ScanVectorIndexRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ScanVectorIndexRequest_SdkV2) {
}

func (to *ScanVectorIndexRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ScanVectorIndexRequest_SdkV2) {
}

func (m ScanVectorIndexRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["last_primary_key"] = attrs["last_primary_key"].SetOptional()
	attrs["num_results"] = attrs["num_results"].SetOptional()
	attrs["index_name"] = attrs["index_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ScanVectorIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ScanVectorIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ScanVectorIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ScanVectorIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name":       m.IndexName,
			"last_primary_key": m.LastPrimaryKey,
			"num_results":      m.NumResults,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ScanVectorIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name":       types.StringType,
			"last_primary_key": types.StringType,
			"num_results":      types.Int64Type,
		},
	}
}

// Response to a scan vector index request.
type ScanVectorIndexResponse_SdkV2 struct {
	// List of data entries
	Data types.List `tfsdk:"data"`
	// Primary key of the last entry.
	LastPrimaryKey types.String `tfsdk:"last_primary_key"`
}

func (to *ScanVectorIndexResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ScanVectorIndexResponse_SdkV2) {
	if !from.Data.IsNull() && !from.Data.IsUnknown() && to.Data.IsNull() && len(from.Data.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Data, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Data = from.Data
	}
}

func (to *ScanVectorIndexResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ScanVectorIndexResponse_SdkV2) {
	if !from.Data.IsNull() && !from.Data.IsUnknown() && to.Data.IsNull() && len(from.Data.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Data, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Data = from.Data
	}
}

func (m ScanVectorIndexResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["data"] = attrs["data"].SetOptional()
	attrs["last_primary_key"] = attrs["last_primary_key"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ScanVectorIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ScanVectorIndexResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data": reflect.TypeOf(Struct_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ScanVectorIndexResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ScanVectorIndexResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data":             m.Data,
			"last_primary_key": m.LastPrimaryKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ScanVectorIndexResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data": basetypes.ListType{
				ElemType: Struct_SdkV2{}.Type(ctx),
			},
			"last_primary_key": types.StringType,
		},
	}
}

// GetData returns the value of the Data field in ScanVectorIndexResponse_SdkV2 as
// a slice of Struct_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ScanVectorIndexResponse_SdkV2) GetData(ctx context.Context) ([]Struct_SdkV2, bool) {
	if m.Data.IsNull() || m.Data.IsUnknown() {
		return nil, false
	}
	var v []Struct_SdkV2
	d := m.Data.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetData sets the value of the Data field in ScanVectorIndexResponse_SdkV2.
func (m *ScanVectorIndexResponse_SdkV2) SetData(ctx context.Context, v []Struct_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["data"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Data = types.ListValueMust(t, vs)
}

type Struct_SdkV2 struct {
	// Data entry, corresponding to a row in a vector index.
	Fields types.List `tfsdk:"fields"`
}

func (to *Struct_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Struct_SdkV2) {
	if !from.Fields.IsNull() && !from.Fields.IsUnknown() && to.Fields.IsNull() && len(from.Fields.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Fields, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Fields = from.Fields
	}
}

func (to *Struct_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Struct_SdkV2) {
	if !from.Fields.IsNull() && !from.Fields.IsUnknown() && to.Fields.IsNull() && len(from.Fields.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Fields, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Fields = from.Fields
	}
}

func (m Struct_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["fields"] = attrs["fields"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Struct.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Struct_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"fields": reflect.TypeOf(MapStringValueEntry_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Struct_SdkV2
// only implements ToObjectValue() and Type().
func (m Struct_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"fields": m.Fields,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Struct_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"fields": basetypes.ListType{
				ElemType: MapStringValueEntry_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetFields returns the value of the Fields field in Struct_SdkV2 as
// a slice of MapStringValueEntry_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *Struct_SdkV2) GetFields(ctx context.Context) ([]MapStringValueEntry_SdkV2, bool) {
	if m.Fields.IsNull() || m.Fields.IsUnknown() {
		return nil, false
	}
	var v []MapStringValueEntry_SdkV2
	d := m.Fields.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFields sets the value of the Fields field in Struct_SdkV2.
func (m *Struct_SdkV2) SetFields(ctx context.Context, v []MapStringValueEntry_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["fields"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Fields = types.ListValueMust(t, vs)
}

type SyncIndexRequest_SdkV2 struct {
	// Name of the vector index to synchronize. Must be a Delta Sync Index.
	IndexName types.String `tfsdk:"-"`
}

func (to *SyncIndexRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncIndexRequest_SdkV2) {
}

func (to *SyncIndexRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SyncIndexRequest_SdkV2) {
}

func (m SyncIndexRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["index_name"] = attrs["index_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SyncIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m SyncIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name": m.IndexName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SyncIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name": types.StringType,
		},
	}
}

type SyncIndexResponse_SdkV2 struct {
}

func (to *SyncIndexResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncIndexResponse_SdkV2) {
}

func (to *SyncIndexResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SyncIndexResponse_SdkV2) {
}

func (m SyncIndexResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SyncIndexResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncIndexResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m SyncIndexResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m SyncIndexResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateEndpointCustomTagsRequest_SdkV2 struct {
	// The new custom tags for the vector search endpoint
	CustomTags types.List `tfsdk:"custom_tags"`
	// Name of the vector search endpoint
	EndpointName types.String `tfsdk:"-"`
}

func (to *UpdateEndpointCustomTagsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEndpointCustomTagsRequest_SdkV2) {
}

func (to *UpdateEndpointCustomTagsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateEndpointCustomTagsRequest_SdkV2) {
}

func (m UpdateEndpointCustomTagsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["custom_tags"] = attrs["custom_tags"].SetRequired()
	attrs["endpoint_name"] = attrs["endpoint_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateEndpointCustomTagsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateEndpointCustomTagsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_tags": reflect.TypeOf(CustomTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEndpointCustomTagsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateEndpointCustomTagsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"custom_tags":   m.CustomTags,
			"endpoint_name": m.EndpointName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateEndpointCustomTagsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"custom_tags": basetypes.ListType{
				ElemType: CustomTag_SdkV2{}.Type(ctx),
			},
			"endpoint_name": types.StringType,
		},
	}
}

// GetCustomTags returns the value of the CustomTags field in UpdateEndpointCustomTagsRequest_SdkV2 as
// a slice of CustomTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateEndpointCustomTagsRequest_SdkV2) GetCustomTags(ctx context.Context) ([]CustomTag_SdkV2, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v []CustomTag_SdkV2
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in UpdateEndpointCustomTagsRequest_SdkV2.
func (m *UpdateEndpointCustomTagsRequest_SdkV2) SetCustomTags(ctx context.Context, v []CustomTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.ListValueMust(t, vs)
}

type UpdateEndpointCustomTagsResponse_SdkV2 struct {
	// All the custom tags that are applied to the vector search endpoint.
	CustomTags types.List `tfsdk:"custom_tags"`
	// The name of the vector search endpoint whose custom tags were updated.
	Name types.String `tfsdk:"name"`
}

func (to *UpdateEndpointCustomTagsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEndpointCustomTagsResponse_SdkV2) {
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
}

func (to *UpdateEndpointCustomTagsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateEndpointCustomTagsResponse_SdkV2) {
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
}

func (m UpdateEndpointCustomTagsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateEndpointCustomTagsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateEndpointCustomTagsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_tags": reflect.TypeOf(CustomTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEndpointCustomTagsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateEndpointCustomTagsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"custom_tags": m.CustomTags,
			"name":        m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateEndpointCustomTagsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"custom_tags": basetypes.ListType{
				ElemType: CustomTag_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
		},
	}
}

// GetCustomTags returns the value of the CustomTags field in UpdateEndpointCustomTagsResponse_SdkV2 as
// a slice of CustomTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateEndpointCustomTagsResponse_SdkV2) GetCustomTags(ctx context.Context) ([]CustomTag_SdkV2, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v []CustomTag_SdkV2
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in UpdateEndpointCustomTagsResponse_SdkV2.
func (m *UpdateEndpointCustomTagsResponse_SdkV2) SetCustomTags(ctx context.Context, v []CustomTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.ListValueMust(t, vs)
}

type UpdateVectorIndexUsagePolicyRequest_SdkV2 struct {
	// Name of the vector search index
	IndexName types.String `tfsdk:"-"`
}

func (to *UpdateVectorIndexUsagePolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateVectorIndexUsagePolicyRequest_SdkV2) {
}

func (to *UpdateVectorIndexUsagePolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateVectorIndexUsagePolicyRequest_SdkV2) {
}

func (m UpdateVectorIndexUsagePolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["index_name"] = attrs["index_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateVectorIndexUsagePolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateVectorIndexUsagePolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateVectorIndexUsagePolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateVectorIndexUsagePolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name": m.IndexName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateVectorIndexUsagePolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name": types.StringType,
		},
	}
}

type UpdateVectorIndexUsagePolicyResponse_SdkV2 struct {
}

func (to *UpdateVectorIndexUsagePolicyResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateVectorIndexUsagePolicyResponse_SdkV2) {
}

func (to *UpdateVectorIndexUsagePolicyResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateVectorIndexUsagePolicyResponse_SdkV2) {
}

func (m UpdateVectorIndexUsagePolicyResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateVectorIndexUsagePolicyResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateVectorIndexUsagePolicyResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateVectorIndexUsagePolicyResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateVectorIndexUsagePolicyResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateVectorIndexUsagePolicyResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpsertDataResult_SdkV2 struct {
	// List of primary keys for rows that failed to process.
	FailedPrimaryKeys types.List `tfsdk:"failed_primary_keys"`
	// Count of successfully processed rows.
	SuccessRowCount types.Int64 `tfsdk:"success_row_count"`
}

func (to *UpsertDataResult_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpsertDataResult_SdkV2) {
	if !from.FailedPrimaryKeys.IsNull() && !from.FailedPrimaryKeys.IsUnknown() && to.FailedPrimaryKeys.IsNull() && len(from.FailedPrimaryKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FailedPrimaryKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FailedPrimaryKeys = from.FailedPrimaryKeys
	}
}

func (to *UpsertDataResult_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpsertDataResult_SdkV2) {
	if !from.FailedPrimaryKeys.IsNull() && !from.FailedPrimaryKeys.IsUnknown() && to.FailedPrimaryKeys.IsNull() && len(from.FailedPrimaryKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FailedPrimaryKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FailedPrimaryKeys = from.FailedPrimaryKeys
	}
}

func (m UpsertDataResult_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["failed_primary_keys"] = attrs["failed_primary_keys"].SetOptional()
	attrs["success_row_count"] = attrs["success_row_count"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpsertDataResult.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpsertDataResult_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"failed_primary_keys": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpsertDataResult_SdkV2
// only implements ToObjectValue() and Type().
func (m UpsertDataResult_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"failed_primary_keys": m.FailedPrimaryKeys,
			"success_row_count":   m.SuccessRowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpsertDataResult_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"failed_primary_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
			"success_row_count": types.Int64Type,
		},
	}
}

// GetFailedPrimaryKeys returns the value of the FailedPrimaryKeys field in UpsertDataResult_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpsertDataResult_SdkV2) GetFailedPrimaryKeys(ctx context.Context) ([]types.String, bool) {
	if m.FailedPrimaryKeys.IsNull() || m.FailedPrimaryKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.FailedPrimaryKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFailedPrimaryKeys sets the value of the FailedPrimaryKeys field in UpsertDataResult_SdkV2.
func (m *UpsertDataResult_SdkV2) SetFailedPrimaryKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["failed_primary_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FailedPrimaryKeys = types.ListValueMust(t, vs)
}

type UpsertDataVectorIndexRequest_SdkV2 struct {
	// Name of the vector index where data is to be upserted. Must be a Direct
	// Vector Access Index.
	IndexName types.String `tfsdk:"-"`
	// JSON string representing the data to be upserted.
	InputsJson types.String `tfsdk:"inputs_json"`
}

func (to *UpsertDataVectorIndexRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpsertDataVectorIndexRequest_SdkV2) {
}

func (to *UpsertDataVectorIndexRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpsertDataVectorIndexRequest_SdkV2) {
}

func (m UpsertDataVectorIndexRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inputs_json"] = attrs["inputs_json"].SetRequired()
	attrs["index_name"] = attrs["index_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpsertDataVectorIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpsertDataVectorIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpsertDataVectorIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpsertDataVectorIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name":  m.IndexName,
			"inputs_json": m.InputsJson,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpsertDataVectorIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name":  types.StringType,
			"inputs_json": types.StringType,
		},
	}
}

type UpsertDataVectorIndexResponse_SdkV2 struct {
	// Result of the upsert or delete operation.
	Result types.List `tfsdk:"result"`
	// Status of the upsert operation.
	Status types.String `tfsdk:"status"`
}

func (to *UpsertDataVectorIndexResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpsertDataVectorIndexResponse_SdkV2) {
	if !from.Result.IsNull() && !from.Result.IsUnknown() {
		if toResult, ok := to.GetResult(ctx); ok {
			if fromResult, ok := from.GetResult(ctx); ok {
				// Recursively sync the fields of Result
				toResult.SyncFieldsDuringCreateOrUpdate(ctx, fromResult)
				to.SetResult(ctx, toResult)
			}
		}
	}
}

func (to *UpsertDataVectorIndexResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpsertDataVectorIndexResponse_SdkV2) {
	if !from.Result.IsNull() && !from.Result.IsUnknown() {
		if toResult, ok := to.GetResult(ctx); ok {
			if fromResult, ok := from.GetResult(ctx); ok {
				toResult.SyncFieldsDuringRead(ctx, fromResult)
				to.SetResult(ctx, toResult)
			}
		}
	}
}

func (m UpsertDataVectorIndexResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["result"] = attrs["result"].SetOptional()
	attrs["result"] = attrs["result"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpsertDataVectorIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpsertDataVectorIndexResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"result": reflect.TypeOf(UpsertDataResult_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpsertDataVectorIndexResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m UpsertDataVectorIndexResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"result": m.Result,
			"status": m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpsertDataVectorIndexResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"result": basetypes.ListType{
				ElemType: UpsertDataResult_SdkV2{}.Type(ctx),
			},
			"status": types.StringType,
		},
	}
}

// GetResult returns the value of the Result field in UpsertDataVectorIndexResponse_SdkV2 as
// a UpsertDataResult_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpsertDataVectorIndexResponse_SdkV2) GetResult(ctx context.Context) (UpsertDataResult_SdkV2, bool) {
	var e UpsertDataResult_SdkV2
	if m.Result.IsNull() || m.Result.IsUnknown() {
		return e, false
	}
	var v []UpsertDataResult_SdkV2
	d := m.Result.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetResult sets the value of the Result field in UpsertDataVectorIndexResponse_SdkV2.
func (m *UpsertDataVectorIndexResponse_SdkV2) SetResult(ctx context.Context, v UpsertDataResult_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["result"]
	m.Result = types.ListValueMust(t, vs)
}

type Value_SdkV2 struct {
	BoolValue types.Bool `tfsdk:"bool_value"`

	ListValue types.List `tfsdk:"list_value"`

	NumberValue types.Float64 `tfsdk:"number_value"`

	StringValue types.String `tfsdk:"string_value"`

	StructValue types.List `tfsdk:"struct_value"`
}

func (to *Value_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Value_SdkV2) {
	if !from.ListValue.IsNull() && !from.ListValue.IsUnknown() {
		if toListValue, ok := to.GetListValue(ctx); ok {
			if fromListValue, ok := from.GetListValue(ctx); ok {
				// Recursively sync the fields of ListValue
				toListValue.SyncFieldsDuringCreateOrUpdate(ctx, fromListValue)
				to.SetListValue(ctx, toListValue)
			}
		}
	}
	if !from.StructValue.IsNull() && !from.StructValue.IsUnknown() {
		if toStructValue, ok := to.GetStructValue(ctx); ok {
			if fromStructValue, ok := from.GetStructValue(ctx); ok {
				// Recursively sync the fields of StructValue
				toStructValue.SyncFieldsDuringCreateOrUpdate(ctx, fromStructValue)
				to.SetStructValue(ctx, toStructValue)
			}
		}
	}
}

func (to *Value_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Value_SdkV2) {
	if !from.ListValue.IsNull() && !from.ListValue.IsUnknown() {
		if toListValue, ok := to.GetListValue(ctx); ok {
			if fromListValue, ok := from.GetListValue(ctx); ok {
				toListValue.SyncFieldsDuringRead(ctx, fromListValue)
				to.SetListValue(ctx, toListValue)
			}
		}
	}
	if !from.StructValue.IsNull() && !from.StructValue.IsUnknown() {
		if toStructValue, ok := to.GetStructValue(ctx); ok {
			if fromStructValue, ok := from.GetStructValue(ctx); ok {
				toStructValue.SyncFieldsDuringRead(ctx, fromStructValue)
				to.SetStructValue(ctx, toStructValue)
			}
		}
	}
}

func (m Value_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["bool_value"] = attrs["bool_value"].SetOptional()
	attrs["list_value"] = attrs["list_value"].SetOptional()
	attrs["list_value"] = attrs["list_value"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["number_value"] = attrs["number_value"].SetOptional()
	attrs["string_value"] = attrs["string_value"].SetOptional()
	attrs["struct_value"] = attrs["struct_value"].SetOptional()
	attrs["struct_value"] = attrs["struct_value"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

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
	return map[string]reflect.Type{
		"list_value":   reflect.TypeOf(ListValue_SdkV2{}),
		"struct_value": reflect.TypeOf(Struct_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Value_SdkV2
// only implements ToObjectValue() and Type().
func (m Value_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bool_value":   m.BoolValue,
			"list_value":   m.ListValue,
			"number_value": m.NumberValue,
			"string_value": m.StringValue,
			"struct_value": m.StructValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Value_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bool_value": types.BoolType,
			"list_value": basetypes.ListType{
				ElemType: ListValue_SdkV2{}.Type(ctx),
			},
			"number_value": types.Float64Type,
			"string_value": types.StringType,
			"struct_value": basetypes.ListType{
				ElemType: Struct_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetListValue returns the value of the ListValue field in Value_SdkV2 as
// a ListValue_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Value_SdkV2) GetListValue(ctx context.Context) (ListValue_SdkV2, bool) {
	var e ListValue_SdkV2
	if m.ListValue.IsNull() || m.ListValue.IsUnknown() {
		return e, false
	}
	var v []ListValue_SdkV2
	d := m.ListValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetListValue sets the value of the ListValue field in Value_SdkV2.
func (m *Value_SdkV2) SetListValue(ctx context.Context, v ListValue_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["list_value"]
	m.ListValue = types.ListValueMust(t, vs)
}

// GetStructValue returns the value of the StructValue field in Value_SdkV2 as
// a Struct_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Value_SdkV2) GetStructValue(ctx context.Context) (Struct_SdkV2, bool) {
	var e Struct_SdkV2
	if m.StructValue.IsNull() || m.StructValue.IsUnknown() {
		return e, false
	}
	var v []Struct_SdkV2
	d := m.StructValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStructValue sets the value of the StructValue field in Value_SdkV2.
func (m *Value_SdkV2) SetStructValue(ctx context.Context, v Struct_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["struct_value"]
	m.StructValue = types.ListValueMust(t, vs)
}

type VectorIndex_SdkV2 struct {
	// The user who created the index.
	Creator types.String `tfsdk:"creator"`

	DeltaSyncIndexSpec types.List `tfsdk:"delta_sync_index_spec"`

	DirectAccessIndexSpec types.List `tfsdk:"direct_access_index_spec"`
	// Name of the endpoint associated with the index
	EndpointName types.String `tfsdk:"endpoint_name"`

	IndexType types.String `tfsdk:"index_type"`
	// Name of the index
	Name types.String `tfsdk:"name"`
	// Primary key of the index
	PrimaryKey types.String `tfsdk:"primary_key"`

	Status types.List `tfsdk:"status"`
}

func (to *VectorIndex_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from VectorIndex_SdkV2) {
	if !from.DeltaSyncIndexSpec.IsNull() && !from.DeltaSyncIndexSpec.IsUnknown() {
		if toDeltaSyncIndexSpec, ok := to.GetDeltaSyncIndexSpec(ctx); ok {
			if fromDeltaSyncIndexSpec, ok := from.GetDeltaSyncIndexSpec(ctx); ok {
				// Recursively sync the fields of DeltaSyncIndexSpec
				toDeltaSyncIndexSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromDeltaSyncIndexSpec)
				to.SetDeltaSyncIndexSpec(ctx, toDeltaSyncIndexSpec)
			}
		}
	}
	if !from.DirectAccessIndexSpec.IsNull() && !from.DirectAccessIndexSpec.IsUnknown() {
		if toDirectAccessIndexSpec, ok := to.GetDirectAccessIndexSpec(ctx); ok {
			if fromDirectAccessIndexSpec, ok := from.GetDirectAccessIndexSpec(ctx); ok {
				// Recursively sync the fields of DirectAccessIndexSpec
				toDirectAccessIndexSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromDirectAccessIndexSpec)
				to.SetDirectAccessIndexSpec(ctx, toDirectAccessIndexSpec)
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

func (to *VectorIndex_SdkV2) SyncFieldsDuringRead(ctx context.Context, from VectorIndex_SdkV2) {
	if !from.DeltaSyncIndexSpec.IsNull() && !from.DeltaSyncIndexSpec.IsUnknown() {
		if toDeltaSyncIndexSpec, ok := to.GetDeltaSyncIndexSpec(ctx); ok {
			if fromDeltaSyncIndexSpec, ok := from.GetDeltaSyncIndexSpec(ctx); ok {
				toDeltaSyncIndexSpec.SyncFieldsDuringRead(ctx, fromDeltaSyncIndexSpec)
				to.SetDeltaSyncIndexSpec(ctx, toDeltaSyncIndexSpec)
			}
		}
	}
	if !from.DirectAccessIndexSpec.IsNull() && !from.DirectAccessIndexSpec.IsUnknown() {
		if toDirectAccessIndexSpec, ok := to.GetDirectAccessIndexSpec(ctx); ok {
			if fromDirectAccessIndexSpec, ok := from.GetDirectAccessIndexSpec(ctx); ok {
				toDirectAccessIndexSpec.SyncFieldsDuringRead(ctx, fromDirectAccessIndexSpec)
				to.SetDirectAccessIndexSpec(ctx, toDirectAccessIndexSpec)
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

func (m VectorIndex_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creator"] = attrs["creator"].SetOptional()
	attrs["delta_sync_index_spec"] = attrs["delta_sync_index_spec"].SetOptional()
	attrs["delta_sync_index_spec"] = attrs["delta_sync_index_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["direct_access_index_spec"] = attrs["direct_access_index_spec"].SetOptional()
	attrs["direct_access_index_spec"] = attrs["direct_access_index_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["endpoint_name"] = attrs["endpoint_name"].SetOptional()
	attrs["index_type"] = attrs["index_type"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["primary_key"] = attrs["primary_key"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["status"] = attrs["status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in VectorIndex.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m VectorIndex_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"delta_sync_index_spec":    reflect.TypeOf(DeltaSyncVectorIndexSpecResponse_SdkV2{}),
		"direct_access_index_spec": reflect.TypeOf(DirectAccessVectorIndexSpec_SdkV2{}),
		"status":                   reflect.TypeOf(VectorIndexStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VectorIndex_SdkV2
// only implements ToObjectValue() and Type().
func (m VectorIndex_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creator":                  m.Creator,
			"delta_sync_index_spec":    m.DeltaSyncIndexSpec,
			"direct_access_index_spec": m.DirectAccessIndexSpec,
			"endpoint_name":            m.EndpointName,
			"index_type":               m.IndexType,
			"name":                     m.Name,
			"primary_key":              m.PrimaryKey,
			"status":                   m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m VectorIndex_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creator": types.StringType,
			"delta_sync_index_spec": basetypes.ListType{
				ElemType: DeltaSyncVectorIndexSpecResponse_SdkV2{}.Type(ctx),
			},
			"direct_access_index_spec": basetypes.ListType{
				ElemType: DirectAccessVectorIndexSpec_SdkV2{}.Type(ctx),
			},
			"endpoint_name": types.StringType,
			"index_type":    types.StringType,
			"name":          types.StringType,
			"primary_key":   types.StringType,
			"status": basetypes.ListType{
				ElemType: VectorIndexStatus_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetDeltaSyncIndexSpec returns the value of the DeltaSyncIndexSpec field in VectorIndex_SdkV2 as
// a DeltaSyncVectorIndexSpecResponse_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *VectorIndex_SdkV2) GetDeltaSyncIndexSpec(ctx context.Context) (DeltaSyncVectorIndexSpecResponse_SdkV2, bool) {
	var e DeltaSyncVectorIndexSpecResponse_SdkV2
	if m.DeltaSyncIndexSpec.IsNull() || m.DeltaSyncIndexSpec.IsUnknown() {
		return e, false
	}
	var v []DeltaSyncVectorIndexSpecResponse_SdkV2
	d := m.DeltaSyncIndexSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDeltaSyncIndexSpec sets the value of the DeltaSyncIndexSpec field in VectorIndex_SdkV2.
func (m *VectorIndex_SdkV2) SetDeltaSyncIndexSpec(ctx context.Context, v DeltaSyncVectorIndexSpecResponse_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["delta_sync_index_spec"]
	m.DeltaSyncIndexSpec = types.ListValueMust(t, vs)
}

// GetDirectAccessIndexSpec returns the value of the DirectAccessIndexSpec field in VectorIndex_SdkV2 as
// a DirectAccessVectorIndexSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *VectorIndex_SdkV2) GetDirectAccessIndexSpec(ctx context.Context) (DirectAccessVectorIndexSpec_SdkV2, bool) {
	var e DirectAccessVectorIndexSpec_SdkV2
	if m.DirectAccessIndexSpec.IsNull() || m.DirectAccessIndexSpec.IsUnknown() {
		return e, false
	}
	var v []DirectAccessVectorIndexSpec_SdkV2
	d := m.DirectAccessIndexSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDirectAccessIndexSpec sets the value of the DirectAccessIndexSpec field in VectorIndex_SdkV2.
func (m *VectorIndex_SdkV2) SetDirectAccessIndexSpec(ctx context.Context, v DirectAccessVectorIndexSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["direct_access_index_spec"]
	m.DirectAccessIndexSpec = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in VectorIndex_SdkV2 as
// a VectorIndexStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *VectorIndex_SdkV2) GetStatus(ctx context.Context) (VectorIndexStatus_SdkV2, bool) {
	var e VectorIndexStatus_SdkV2
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v []VectorIndexStatus_SdkV2
	d := m.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in VectorIndex_SdkV2.
func (m *VectorIndex_SdkV2) SetStatus(ctx context.Context, v VectorIndexStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	m.Status = types.ListValueMust(t, vs)
}

type VectorIndexStatus_SdkV2 struct {
	// Index API Url to be used to perform operations on the index
	IndexUrl types.String `tfsdk:"index_url"`
	// Number of rows indexed
	IndexedRowCount types.Int64 `tfsdk:"indexed_row_count"`
	// Message associated with the index status
	Message types.String `tfsdk:"message"`
	// Whether the index is ready for search
	Ready types.Bool `tfsdk:"ready"`
}

func (to *VectorIndexStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from VectorIndexStatus_SdkV2) {
}

func (to *VectorIndexStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, from VectorIndexStatus_SdkV2) {
}

func (m VectorIndexStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["index_url"] = attrs["index_url"].SetOptional()
	attrs["indexed_row_count"] = attrs["indexed_row_count"].SetOptional()
	attrs["message"] = attrs["message"].SetOptional()
	attrs["ready"] = attrs["ready"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in VectorIndexStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m VectorIndexStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VectorIndexStatus_SdkV2
// only implements ToObjectValue() and Type().
func (m VectorIndexStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_url":         m.IndexUrl,
			"indexed_row_count": m.IndexedRowCount,
			"message":           m.Message,
			"ready":             m.Ready,
		})
}

// Type implements basetypes.ObjectValuable.
func (m VectorIndexStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_url":         types.StringType,
			"indexed_row_count": types.Int64Type,
			"message":           types.StringType,
			"ready":             types.BoolType,
		},
	}
}
