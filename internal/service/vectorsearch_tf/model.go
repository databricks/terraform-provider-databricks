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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type ColumnInfo struct {
	// Name of the column.
	Name types.String `tfsdk:"name"`
}

func (to *ColumnInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ColumnInfo) {
}

func (to *ColumnInfo) SyncFieldsDuringRead(ctx context.Context, from ColumnInfo) {
}

func (m ColumnInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ColumnInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ColumnInfo
// only implements ToObjectValue() and Type().
func (m ColumnInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ColumnInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type CreateEndpoint struct {
	// The budget policy id to be applied
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// Type of endpoint
	EndpointType types.String `tfsdk:"endpoint_type"`
	// Name of the vector search endpoint
	Name types.String `tfsdk:"name"`
}

func (to *CreateEndpoint) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateEndpoint) {
}

func (to *CreateEndpoint) SyncFieldsDuringRead(ctx context.Context, from CreateEndpoint) {
}

func (m CreateEndpoint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["endpoint_type"] = attrs["endpoint_type"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateEndpoint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateEndpoint) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateEndpoint
// only implements ToObjectValue() and Type().
func (m CreateEndpoint) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget_policy_id": m.BudgetPolicyId,
			"endpoint_type":    m.EndpointType,
			"name":             m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateEndpoint) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_policy_id": types.StringType,
			"endpoint_type":    types.StringType,
			"name":             types.StringType,
		},
	}
}

type CreateVectorIndexRequest struct {
	// Specification for Delta Sync Index. Required if `index_type` is
	// `DELTA_SYNC`.
	DeltaSyncIndexSpec types.Object `tfsdk:"delta_sync_index_spec"`
	// Specification for Direct Vector Access Index. Required if `index_type` is
	// `DIRECT_ACCESS`.
	DirectAccessIndexSpec types.Object `tfsdk:"direct_access_index_spec"`
	// Name of the endpoint to be used for serving the index
	EndpointName types.String `tfsdk:"endpoint_name"`

	IndexType types.String `tfsdk:"index_type"`
	// Name of the index
	Name types.String `tfsdk:"name"`
	// Primary key of the index
	PrimaryKey types.String `tfsdk:"primary_key"`
}

func (to *CreateVectorIndexRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateVectorIndexRequest) {
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

func (to *CreateVectorIndexRequest) SyncFieldsDuringRead(ctx context.Context, from CreateVectorIndexRequest) {
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

func (m CreateVectorIndexRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["delta_sync_index_spec"] = attrs["delta_sync_index_spec"].SetOptional()
	attrs["direct_access_index_spec"] = attrs["direct_access_index_spec"].SetOptional()
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
func (m CreateVectorIndexRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"delta_sync_index_spec":    reflect.TypeOf(DeltaSyncVectorIndexSpecRequest{}),
		"direct_access_index_spec": reflect.TypeOf(DirectAccessVectorIndexSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateVectorIndexRequest
// only implements ToObjectValue() and Type().
func (m CreateVectorIndexRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CreateVectorIndexRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"delta_sync_index_spec":    DeltaSyncVectorIndexSpecRequest{}.Type(ctx),
			"direct_access_index_spec": DirectAccessVectorIndexSpec{}.Type(ctx),
			"endpoint_name":            types.StringType,
			"index_type":               types.StringType,
			"name":                     types.StringType,
			"primary_key":              types.StringType,
		},
	}
}

// GetDeltaSyncIndexSpec returns the value of the DeltaSyncIndexSpec field in CreateVectorIndexRequest as
// a DeltaSyncVectorIndexSpecRequest value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateVectorIndexRequest) GetDeltaSyncIndexSpec(ctx context.Context) (DeltaSyncVectorIndexSpecRequest, bool) {
	var e DeltaSyncVectorIndexSpecRequest
	if m.DeltaSyncIndexSpec.IsNull() || m.DeltaSyncIndexSpec.IsUnknown() {
		return e, false
	}
	var v DeltaSyncVectorIndexSpecRequest
	d := m.DeltaSyncIndexSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDeltaSyncIndexSpec sets the value of the DeltaSyncIndexSpec field in CreateVectorIndexRequest.
func (m *CreateVectorIndexRequest) SetDeltaSyncIndexSpec(ctx context.Context, v DeltaSyncVectorIndexSpecRequest) {
	vs := v.ToObjectValue(ctx)
	m.DeltaSyncIndexSpec = vs
}

// GetDirectAccessIndexSpec returns the value of the DirectAccessIndexSpec field in CreateVectorIndexRequest as
// a DirectAccessVectorIndexSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateVectorIndexRequest) GetDirectAccessIndexSpec(ctx context.Context) (DirectAccessVectorIndexSpec, bool) {
	var e DirectAccessVectorIndexSpec
	if m.DirectAccessIndexSpec.IsNull() || m.DirectAccessIndexSpec.IsUnknown() {
		return e, false
	}
	var v DirectAccessVectorIndexSpec
	d := m.DirectAccessIndexSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDirectAccessIndexSpec sets the value of the DirectAccessIndexSpec field in CreateVectorIndexRequest.
func (m *CreateVectorIndexRequest) SetDirectAccessIndexSpec(ctx context.Context, v DirectAccessVectorIndexSpec) {
	vs := v.ToObjectValue(ctx)
	m.DirectAccessIndexSpec = vs
}

type CustomTag struct {
	// Key field for a vector search endpoint tag.
	Key types.String `tfsdk:"key"`
	// [Optional] Value field for a vector search endpoint tag.
	Value types.String `tfsdk:"value"`
}

func (to *CustomTag) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CustomTag) {
}

func (to *CustomTag) SyncFieldsDuringRead(ctx context.Context, from CustomTag) {
}

func (m CustomTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CustomTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomTag
// only implements ToObjectValue() and Type().
func (m CustomTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CustomTag) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type DeleteDataResult struct {
	// List of primary keys for rows that failed to process.
	FailedPrimaryKeys types.List `tfsdk:"failed_primary_keys"`
	// Count of successfully processed rows.
	SuccessRowCount types.Int64 `tfsdk:"success_row_count"`
}

func (to *DeleteDataResult) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDataResult) {
	if !from.FailedPrimaryKeys.IsNull() && !from.FailedPrimaryKeys.IsUnknown() && to.FailedPrimaryKeys.IsNull() && len(from.FailedPrimaryKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FailedPrimaryKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FailedPrimaryKeys = from.FailedPrimaryKeys
	}
}

func (to *DeleteDataResult) SyncFieldsDuringRead(ctx context.Context, from DeleteDataResult) {
	if !from.FailedPrimaryKeys.IsNull() && !from.FailedPrimaryKeys.IsUnknown() && to.FailedPrimaryKeys.IsNull() && len(from.FailedPrimaryKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FailedPrimaryKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FailedPrimaryKeys = from.FailedPrimaryKeys
	}
}

func (m DeleteDataResult) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteDataResult) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"failed_primary_keys": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDataResult
// only implements ToObjectValue() and Type().
func (m DeleteDataResult) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"failed_primary_keys": m.FailedPrimaryKeys,
			"success_row_count":   m.SuccessRowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDataResult) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"failed_primary_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
			"success_row_count": types.Int64Type,
		},
	}
}

// GetFailedPrimaryKeys returns the value of the FailedPrimaryKeys field in DeleteDataResult as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeleteDataResult) GetFailedPrimaryKeys(ctx context.Context) ([]types.String, bool) {
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

// SetFailedPrimaryKeys sets the value of the FailedPrimaryKeys field in DeleteDataResult.
func (m *DeleteDataResult) SetFailedPrimaryKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["failed_primary_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FailedPrimaryKeys = types.ListValueMust(t, vs)
}

type DeleteDataVectorIndexRequest struct {
	// Name of the vector index where data is to be deleted. Must be a Direct
	// Vector Access Index.
	IndexName types.String `tfsdk:"-"`
	// List of primary keys for the data to be deleted.
	PrimaryKeys types.List `tfsdk:"-"`
}

func (to *DeleteDataVectorIndexRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDataVectorIndexRequest) {
}

func (to *DeleteDataVectorIndexRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteDataVectorIndexRequest) {
}

func (m DeleteDataVectorIndexRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteDataVectorIndexRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"primary_keys": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDataVectorIndexRequest
// only implements ToObjectValue() and Type().
func (m DeleteDataVectorIndexRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name":   m.IndexName,
			"primary_keys": m.PrimaryKeys,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDataVectorIndexRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name": types.StringType,
			"primary_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetPrimaryKeys returns the value of the PrimaryKeys field in DeleteDataVectorIndexRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeleteDataVectorIndexRequest) GetPrimaryKeys(ctx context.Context) ([]types.String, bool) {
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

// SetPrimaryKeys sets the value of the PrimaryKeys field in DeleteDataVectorIndexRequest.
func (m *DeleteDataVectorIndexRequest) SetPrimaryKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["primary_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PrimaryKeys = types.ListValueMust(t, vs)
}

type DeleteDataVectorIndexResponse struct {
	// Result of the upsert or delete operation.
	Result types.Object `tfsdk:"result"`
	// Status of the delete operation.
	Status types.String `tfsdk:"status"`
}

func (to *DeleteDataVectorIndexResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDataVectorIndexResponse) {
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

func (to *DeleteDataVectorIndexResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteDataVectorIndexResponse) {
	if !from.Result.IsNull() && !from.Result.IsUnknown() {
		if toResult, ok := to.GetResult(ctx); ok {
			if fromResult, ok := from.GetResult(ctx); ok {
				toResult.SyncFieldsDuringRead(ctx, fromResult)
				to.SetResult(ctx, toResult)
			}
		}
	}
}

func (m DeleteDataVectorIndexResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["result"] = attrs["result"].SetOptional()
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
func (m DeleteDataVectorIndexResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"result": reflect.TypeOf(DeleteDataResult{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDataVectorIndexResponse
// only implements ToObjectValue() and Type().
func (m DeleteDataVectorIndexResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"result": m.Result,
			"status": m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDataVectorIndexResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"result": DeleteDataResult{}.Type(ctx),
			"status": types.StringType,
		},
	}
}

// GetResult returns the value of the Result field in DeleteDataVectorIndexResponse as
// a DeleteDataResult value.
// If the field is unknown or null, the boolean return value is false.
func (m *DeleteDataVectorIndexResponse) GetResult(ctx context.Context) (DeleteDataResult, bool) {
	var e DeleteDataResult
	if m.Result.IsNull() || m.Result.IsUnknown() {
		return e, false
	}
	var v DeleteDataResult
	d := m.Result.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResult sets the value of the Result field in DeleteDataVectorIndexResponse.
func (m *DeleteDataVectorIndexResponse) SetResult(ctx context.Context, v DeleteDataResult) {
	vs := v.ToObjectValue(ctx)
	m.Result = vs
}

type DeleteEndpointRequest struct {
	// Name of the vector search endpoint
	EndpointName types.String `tfsdk:"-"`
}

func (to *DeleteEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteEndpointRequest) {
}

func (to *DeleteEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteEndpointRequest) {
}

func (m DeleteEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"endpoint_name": m.EndpointName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_name": types.StringType,
		},
	}
}

type DeleteEndpointResponse struct {
}

func (to *DeleteEndpointResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteEndpointResponse) {
}

func (to *DeleteEndpointResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteEndpointResponse) {
}

func (m DeleteEndpointResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteEndpointResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteEndpointResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteEndpointResponse
// only implements ToObjectValue() and Type().
func (m DeleteEndpointResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteEndpointResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteIndexRequest struct {
	// Name of the index
	IndexName types.String `tfsdk:"-"`
}

func (to *DeleteIndexRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteIndexRequest) {
}

func (to *DeleteIndexRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteIndexRequest) {
}

func (m DeleteIndexRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteIndexRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteIndexRequest
// only implements ToObjectValue() and Type().
func (m DeleteIndexRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name": m.IndexName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteIndexRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name": types.StringType,
		},
	}
}

type DeleteIndexResponse struct {
}

func (to *DeleteIndexResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteIndexResponse) {
}

func (to *DeleteIndexResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteIndexResponse) {
}

func (m DeleteIndexResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteIndexResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteIndexResponse
// only implements ToObjectValue() and Type().
func (m DeleteIndexResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteIndexResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeltaSyncVectorIndexSpecRequest struct {
	// [Optional] Select the columns to sync with the vector index. If you leave
	// this field blank, all columns from the source table are synced with the
	// index. The primary key column and embedding source column or embedding
	// vector column are always synced.
	ColumnsToSync types.List `tfsdk:"columns_to_sync"`
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

func (to *DeltaSyncVectorIndexSpecRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeltaSyncVectorIndexSpecRequest) {
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

func (to *DeltaSyncVectorIndexSpecRequest) SyncFieldsDuringRead(ctx context.Context, from DeltaSyncVectorIndexSpecRequest) {
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

func (m DeltaSyncVectorIndexSpecRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["columns_to_sync"] = attrs["columns_to_sync"].SetOptional()
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
func (m DeltaSyncVectorIndexSpecRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns_to_sync":          reflect.TypeOf(types.String{}),
		"embedding_source_columns": reflect.TypeOf(EmbeddingSourceColumn{}),
		"embedding_vector_columns": reflect.TypeOf(EmbeddingVectorColumn{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaSyncVectorIndexSpecRequest
// only implements ToObjectValue() and Type().
func (m DeltaSyncVectorIndexSpecRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns_to_sync":           m.ColumnsToSync,
			"embedding_source_columns":  m.EmbeddingSourceColumns,
			"embedding_vector_columns":  m.EmbeddingVectorColumns,
			"embedding_writeback_table": m.EmbeddingWritebackTable,
			"pipeline_type":             m.PipelineType,
			"source_table":              m.SourceTable,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeltaSyncVectorIndexSpecRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns_to_sync": basetypes.ListType{
				ElemType: types.StringType,
			},
			"embedding_source_columns": basetypes.ListType{
				ElemType: EmbeddingSourceColumn{}.Type(ctx),
			},
			"embedding_vector_columns": basetypes.ListType{
				ElemType: EmbeddingVectorColumn{}.Type(ctx),
			},
			"embedding_writeback_table": types.StringType,
			"pipeline_type":             types.StringType,
			"source_table":              types.StringType,
		},
	}
}

// GetColumnsToSync returns the value of the ColumnsToSync field in DeltaSyncVectorIndexSpecRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSyncVectorIndexSpecRequest) GetColumnsToSync(ctx context.Context) ([]types.String, bool) {
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

// SetColumnsToSync sets the value of the ColumnsToSync field in DeltaSyncVectorIndexSpecRequest.
func (m *DeltaSyncVectorIndexSpecRequest) SetColumnsToSync(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns_to_sync"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ColumnsToSync = types.ListValueMust(t, vs)
}

// GetEmbeddingSourceColumns returns the value of the EmbeddingSourceColumns field in DeltaSyncVectorIndexSpecRequest as
// a slice of EmbeddingSourceColumn values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSyncVectorIndexSpecRequest) GetEmbeddingSourceColumns(ctx context.Context) ([]EmbeddingSourceColumn, bool) {
	if m.EmbeddingSourceColumns.IsNull() || m.EmbeddingSourceColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingSourceColumn
	d := m.EmbeddingSourceColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingSourceColumns sets the value of the EmbeddingSourceColumns field in DeltaSyncVectorIndexSpecRequest.
func (m *DeltaSyncVectorIndexSpecRequest) SetEmbeddingSourceColumns(ctx context.Context, v []EmbeddingSourceColumn) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_source_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmbeddingSourceColumns = types.ListValueMust(t, vs)
}

// GetEmbeddingVectorColumns returns the value of the EmbeddingVectorColumns field in DeltaSyncVectorIndexSpecRequest as
// a slice of EmbeddingVectorColumn values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSyncVectorIndexSpecRequest) GetEmbeddingVectorColumns(ctx context.Context) ([]EmbeddingVectorColumn, bool) {
	if m.EmbeddingVectorColumns.IsNull() || m.EmbeddingVectorColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingVectorColumn
	d := m.EmbeddingVectorColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingVectorColumns sets the value of the EmbeddingVectorColumns field in DeltaSyncVectorIndexSpecRequest.
func (m *DeltaSyncVectorIndexSpecRequest) SetEmbeddingVectorColumns(ctx context.Context, v []EmbeddingVectorColumn) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_vector_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmbeddingVectorColumns = types.ListValueMust(t, vs)
}

type DeltaSyncVectorIndexSpecResponse struct {
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

func (to *DeltaSyncVectorIndexSpecResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeltaSyncVectorIndexSpecResponse) {
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

func (to *DeltaSyncVectorIndexSpecResponse) SyncFieldsDuringRead(ctx context.Context, from DeltaSyncVectorIndexSpecResponse) {
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

func (m DeltaSyncVectorIndexSpecResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeltaSyncVectorIndexSpecResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"embedding_source_columns": reflect.TypeOf(EmbeddingSourceColumn{}),
		"embedding_vector_columns": reflect.TypeOf(EmbeddingVectorColumn{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaSyncVectorIndexSpecResponse
// only implements ToObjectValue() and Type().
func (m DeltaSyncVectorIndexSpecResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding_source_columns":  m.EmbeddingSourceColumns,
			"embedding_vector_columns":  m.EmbeddingVectorColumns,
			"embedding_writeback_table": m.EmbeddingWritebackTable,
			"pipeline_id":               m.PipelineId,
			"pipeline_type":             m.PipelineType,
			"source_table":              m.SourceTable,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeltaSyncVectorIndexSpecResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"embedding_source_columns": basetypes.ListType{
				ElemType: EmbeddingSourceColumn{}.Type(ctx),
			},
			"embedding_vector_columns": basetypes.ListType{
				ElemType: EmbeddingVectorColumn{}.Type(ctx),
			},
			"embedding_writeback_table": types.StringType,
			"pipeline_id":               types.StringType,
			"pipeline_type":             types.StringType,
			"source_table":              types.StringType,
		},
	}
}

// GetEmbeddingSourceColumns returns the value of the EmbeddingSourceColumns field in DeltaSyncVectorIndexSpecResponse as
// a slice of EmbeddingSourceColumn values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSyncVectorIndexSpecResponse) GetEmbeddingSourceColumns(ctx context.Context) ([]EmbeddingSourceColumn, bool) {
	if m.EmbeddingSourceColumns.IsNull() || m.EmbeddingSourceColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingSourceColumn
	d := m.EmbeddingSourceColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingSourceColumns sets the value of the EmbeddingSourceColumns field in DeltaSyncVectorIndexSpecResponse.
func (m *DeltaSyncVectorIndexSpecResponse) SetEmbeddingSourceColumns(ctx context.Context, v []EmbeddingSourceColumn) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_source_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmbeddingSourceColumns = types.ListValueMust(t, vs)
}

// GetEmbeddingVectorColumns returns the value of the EmbeddingVectorColumns field in DeltaSyncVectorIndexSpecResponse as
// a slice of EmbeddingVectorColumn values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSyncVectorIndexSpecResponse) GetEmbeddingVectorColumns(ctx context.Context) ([]EmbeddingVectorColumn, bool) {
	if m.EmbeddingVectorColumns.IsNull() || m.EmbeddingVectorColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingVectorColumn
	d := m.EmbeddingVectorColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingVectorColumns sets the value of the EmbeddingVectorColumns field in DeltaSyncVectorIndexSpecResponse.
func (m *DeltaSyncVectorIndexSpecResponse) SetEmbeddingVectorColumns(ctx context.Context, v []EmbeddingVectorColumn) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_vector_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmbeddingVectorColumns = types.ListValueMust(t, vs)
}

type DirectAccessVectorIndexSpec struct {
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

func (to *DirectAccessVectorIndexSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DirectAccessVectorIndexSpec) {
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

func (to *DirectAccessVectorIndexSpec) SyncFieldsDuringRead(ctx context.Context, from DirectAccessVectorIndexSpec) {
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

func (m DirectAccessVectorIndexSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DirectAccessVectorIndexSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"embedding_source_columns": reflect.TypeOf(EmbeddingSourceColumn{}),
		"embedding_vector_columns": reflect.TypeOf(EmbeddingVectorColumn{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DirectAccessVectorIndexSpec
// only implements ToObjectValue() and Type().
func (m DirectAccessVectorIndexSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding_source_columns": m.EmbeddingSourceColumns,
			"embedding_vector_columns": m.EmbeddingVectorColumns,
			"schema_json":              m.SchemaJson,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DirectAccessVectorIndexSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"embedding_source_columns": basetypes.ListType{
				ElemType: EmbeddingSourceColumn{}.Type(ctx),
			},
			"embedding_vector_columns": basetypes.ListType{
				ElemType: EmbeddingVectorColumn{}.Type(ctx),
			},
			"schema_json": types.StringType,
		},
	}
}

// GetEmbeddingSourceColumns returns the value of the EmbeddingSourceColumns field in DirectAccessVectorIndexSpec as
// a slice of EmbeddingSourceColumn values.
// If the field is unknown or null, the boolean return value is false.
func (m *DirectAccessVectorIndexSpec) GetEmbeddingSourceColumns(ctx context.Context) ([]EmbeddingSourceColumn, bool) {
	if m.EmbeddingSourceColumns.IsNull() || m.EmbeddingSourceColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingSourceColumn
	d := m.EmbeddingSourceColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingSourceColumns sets the value of the EmbeddingSourceColumns field in DirectAccessVectorIndexSpec.
func (m *DirectAccessVectorIndexSpec) SetEmbeddingSourceColumns(ctx context.Context, v []EmbeddingSourceColumn) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_source_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmbeddingSourceColumns = types.ListValueMust(t, vs)
}

// GetEmbeddingVectorColumns returns the value of the EmbeddingVectorColumns field in DirectAccessVectorIndexSpec as
// a slice of EmbeddingVectorColumn values.
// If the field is unknown or null, the boolean return value is false.
func (m *DirectAccessVectorIndexSpec) GetEmbeddingVectorColumns(ctx context.Context) ([]EmbeddingVectorColumn, bool) {
	if m.EmbeddingVectorColumns.IsNull() || m.EmbeddingVectorColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingVectorColumn
	d := m.EmbeddingVectorColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingVectorColumns sets the value of the EmbeddingVectorColumns field in DirectAccessVectorIndexSpec.
func (m *DirectAccessVectorIndexSpec) SetEmbeddingVectorColumns(ctx context.Context, v []EmbeddingVectorColumn) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_vector_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EmbeddingVectorColumns = types.ListValueMust(t, vs)
}

type EmbeddingSourceColumn struct {
	// Name of the embedding model endpoint, used by default for both ingestion
	// and querying.
	EmbeddingModelEndpointName types.String `tfsdk:"embedding_model_endpoint_name"`
	// Name of the embedding model endpoint which, if specified, is used for
	// querying (not ingestion).
	ModelEndpointNameForQuery types.String `tfsdk:"model_endpoint_name_for_query"`
	// Name of the column
	Name types.String `tfsdk:"name"`
}

func (to *EmbeddingSourceColumn) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EmbeddingSourceColumn) {
}

func (to *EmbeddingSourceColumn) SyncFieldsDuringRead(ctx context.Context, from EmbeddingSourceColumn) {
}

func (m EmbeddingSourceColumn) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EmbeddingSourceColumn) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EmbeddingSourceColumn
// only implements ToObjectValue() and Type().
func (m EmbeddingSourceColumn) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding_model_endpoint_name": m.EmbeddingModelEndpointName,
			"model_endpoint_name_for_query": m.ModelEndpointNameForQuery,
			"name":                          m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EmbeddingSourceColumn) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"embedding_model_endpoint_name": types.StringType,
			"model_endpoint_name_for_query": types.StringType,
			"name":                          types.StringType,
		},
	}
}

type EmbeddingVectorColumn struct {
	// Dimension of the embedding vector
	EmbeddingDimension types.Int64 `tfsdk:"embedding_dimension"`
	// Name of the column
	Name types.String `tfsdk:"name"`
}

func (to *EmbeddingVectorColumn) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EmbeddingVectorColumn) {
}

func (to *EmbeddingVectorColumn) SyncFieldsDuringRead(ctx context.Context, from EmbeddingVectorColumn) {
}

func (m EmbeddingVectorColumn) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EmbeddingVectorColumn) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EmbeddingVectorColumn
// only implements ToObjectValue() and Type().
func (m EmbeddingVectorColumn) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding_dimension": m.EmbeddingDimension,
			"name":                m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EmbeddingVectorColumn) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"embedding_dimension": types.Int64Type,
			"name":                types.StringType,
		},
	}
}

type EndpointInfo struct {
	// Timestamp of endpoint creation
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`
	// Creator of the endpoint
	Creator types.String `tfsdk:"creator"`
	// The custom tags assigned to the endpoint
	CustomTags types.List `tfsdk:"custom_tags"`
	// The budget policy id applied to the endpoint
	EffectiveBudgetPolicyId types.String `tfsdk:"effective_budget_policy_id"`
	// Current status of the endpoint
	EndpointStatus types.Object `tfsdk:"endpoint_status"`
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

func (to *EndpointInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointInfo) {
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

func (to *EndpointInfo) SyncFieldsDuringRead(ctx context.Context, from EndpointInfo) {
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

func (m EndpointInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creation_timestamp"] = attrs["creation_timestamp"].SetOptional()
	attrs["creator"] = attrs["creator"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["effective_budget_policy_id"] = attrs["effective_budget_policy_id"].SetOptional()
	attrs["endpoint_status"] = attrs["endpoint_status"].SetOptional()
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
func (m EndpointInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_tags":     reflect.TypeOf(CustomTag{}),
		"endpoint_status": reflect.TypeOf(EndpointStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointInfo
// only implements ToObjectValue() and Type().
func (m EndpointInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m EndpointInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_timestamp": types.Int64Type,
			"creator":            types.StringType,
			"custom_tags": basetypes.ListType{
				ElemType: CustomTag{}.Type(ctx),
			},
			"effective_budget_policy_id": types.StringType,
			"endpoint_status":            EndpointStatus{}.Type(ctx),
			"endpoint_type":              types.StringType,
			"id":                         types.StringType,
			"last_updated_timestamp":     types.Int64Type,
			"last_updated_user":          types.StringType,
			"name":                       types.StringType,
			"num_indexes":                types.Int64Type,
		},
	}
}

// GetCustomTags returns the value of the CustomTags field in EndpointInfo as
// a slice of CustomTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointInfo) GetCustomTags(ctx context.Context) ([]CustomTag, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v []CustomTag
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in EndpointInfo.
func (m *EndpointInfo) SetCustomTags(ctx context.Context, v []CustomTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.ListValueMust(t, vs)
}

// GetEndpointStatus returns the value of the EndpointStatus field in EndpointInfo as
// a EndpointStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointInfo) GetEndpointStatus(ctx context.Context) (EndpointStatus, bool) {
	var e EndpointStatus
	if m.EndpointStatus.IsNull() || m.EndpointStatus.IsUnknown() {
		return e, false
	}
	var v EndpointStatus
	d := m.EndpointStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEndpointStatus sets the value of the EndpointStatus field in EndpointInfo.
func (m *EndpointInfo) SetEndpointStatus(ctx context.Context, v EndpointStatus) {
	vs := v.ToObjectValue(ctx)
	m.EndpointStatus = vs
}

// Status information of an endpoint
type EndpointStatus struct {
	// Additional status message
	Message types.String `tfsdk:"message"`
	// Current state of the endpoint
	State types.String `tfsdk:"state"`
}

func (to *EndpointStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointStatus) {
}

func (to *EndpointStatus) SyncFieldsDuringRead(ctx context.Context, from EndpointStatus) {
}

func (m EndpointStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EndpointStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointStatus
// only implements ToObjectValue() and Type().
func (m EndpointStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": m.Message,
			"state":   m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"state":   types.StringType,
		},
	}
}

type GetEndpointRequest struct {
	// Name of the endpoint
	EndpointName types.String `tfsdk:"-"`
}

func (to *GetEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEndpointRequest) {
}

func (to *GetEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from GetEndpointRequest) {
}

func (m GetEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"endpoint_name": m.EndpointName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_name": types.StringType,
		},
	}
}

type GetIndexRequest struct {
	// If true, the URL returned for the index is guaranteed to be compatible
	// with the reranker. Currently this means we return the CP URL regardless
	// of how the index is being accessed. If not set or set to false, the URL
	// may still be compatible with the reranker depending on what URL we
	// return.
	EnsureRerankerCompatible types.Bool `tfsdk:"-"`
	// Name of the index
	IndexName types.String `tfsdk:"-"`
}

func (to *GetIndexRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetIndexRequest) {
}

func (to *GetIndexRequest) SyncFieldsDuringRead(ctx context.Context, from GetIndexRequest) {
}

func (m GetIndexRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetIndexRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetIndexRequest
// only implements ToObjectValue() and Type().
func (m GetIndexRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ensure_reranker_compatible": m.EnsureRerankerCompatible,
			"index_name":                 m.IndexName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetIndexRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ensure_reranker_compatible": types.BoolType,
			"index_name":                 types.StringType,
		},
	}
}

type ListEndpointResponse struct {
	// An array of Endpoint objects
	Endpoints types.List `tfsdk:"endpoints"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListEndpointResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEndpointResponse) {
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() && to.Endpoints.IsNull() && len(from.Endpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Endpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Endpoints = from.Endpoints
	}
}

func (to *ListEndpointResponse) SyncFieldsDuringRead(ctx context.Context, from ListEndpointResponse) {
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() && to.Endpoints.IsNull() && len(from.Endpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Endpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Endpoints = from.Endpoints
	}
}

func (m ListEndpointResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListEndpointResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoints": reflect.TypeOf(EndpointInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointResponse
// only implements ToObjectValue() and Type().
func (m ListEndpointResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoints":       m.Endpoints,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListEndpointResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoints": basetypes.ListType{
				ElemType: EndpointInfo{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetEndpoints returns the value of the Endpoints field in ListEndpointResponse as
// a slice of EndpointInfo values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListEndpointResponse) GetEndpoints(ctx context.Context) ([]EndpointInfo, bool) {
	if m.Endpoints.IsNull() || m.Endpoints.IsUnknown() {
		return nil, false
	}
	var v []EndpointInfo
	d := m.Endpoints.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEndpoints sets the value of the Endpoints field in ListEndpointResponse.
func (m *ListEndpointResponse) SetEndpoints(ctx context.Context, v []EndpointInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["endpoints"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Endpoints = types.ListValueMust(t, vs)
}

type ListEndpointsRequest struct {
	// Token for pagination
	PageToken types.String `tfsdk:"-"`
}

func (to *ListEndpointsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEndpointsRequest) {
}

func (to *ListEndpointsRequest) SyncFieldsDuringRead(ctx context.Context, from ListEndpointsRequest) {
}

func (m ListEndpointsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListEndpointsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_token": types.StringType,
		},
	}
}

type ListIndexesRequest struct {
	// Name of the endpoint
	EndpointName types.String `tfsdk:"-"`
	// Token for pagination
	PageToken types.String `tfsdk:"-"`
}

func (to *ListIndexesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListIndexesRequest) {
}

func (to *ListIndexesRequest) SyncFieldsDuringRead(ctx context.Context, from ListIndexesRequest) {
}

func (m ListIndexesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListIndexesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListIndexesRequest
// only implements ToObjectValue() and Type().
func (m ListIndexesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_name": m.EndpointName,
			"page_token":    m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListIndexesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_name": types.StringType,
			"page_token":    types.StringType,
		},
	}
}

type ListValue struct {
	// Repeated field of dynamically typed values.
	Values types.List `tfsdk:"values"`
}

func (to *ListValue) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListValue) {
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (to *ListValue) SyncFieldsDuringRead(ctx context.Context, from ListValue) {
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (m ListValue) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListValue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"values": reflect.TypeOf(Value{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListValue
// only implements ToObjectValue() and Type().
func (m ListValue) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"values": m.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListValue) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"values": basetypes.ListType{
				ElemType: Value{}.Type(ctx),
			},
		},
	}
}

// GetValues returns the value of the Values field in ListValue as
// a slice of Value values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListValue) GetValues(ctx context.Context) ([]Value, bool) {
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

// SetValues sets the value of the Values field in ListValue.
func (m *ListValue) SetValues(ctx context.Context, v []Value) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Values = types.ListValueMust(t, vs)
}

type ListVectorIndexesResponse struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`

	VectorIndexes types.List `tfsdk:"vector_indexes"`
}

func (to *ListVectorIndexesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListVectorIndexesResponse) {
	if !from.VectorIndexes.IsNull() && !from.VectorIndexes.IsUnknown() && to.VectorIndexes.IsNull() && len(from.VectorIndexes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for VectorIndexes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.VectorIndexes = from.VectorIndexes
	}
}

func (to *ListVectorIndexesResponse) SyncFieldsDuringRead(ctx context.Context, from ListVectorIndexesResponse) {
	if !from.VectorIndexes.IsNull() && !from.VectorIndexes.IsUnknown() && to.VectorIndexes.IsNull() && len(from.VectorIndexes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for VectorIndexes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.VectorIndexes = from.VectorIndexes
	}
}

func (m ListVectorIndexesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListVectorIndexesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"vector_indexes": reflect.TypeOf(MiniVectorIndex{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListVectorIndexesResponse
// only implements ToObjectValue() and Type().
func (m ListVectorIndexesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"vector_indexes":  m.VectorIndexes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListVectorIndexesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"vector_indexes": basetypes.ListType{
				ElemType: MiniVectorIndex{}.Type(ctx),
			},
		},
	}
}

// GetVectorIndexes returns the value of the VectorIndexes field in ListVectorIndexesResponse as
// a slice of MiniVectorIndex values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListVectorIndexesResponse) GetVectorIndexes(ctx context.Context) ([]MiniVectorIndex, bool) {
	if m.VectorIndexes.IsNull() || m.VectorIndexes.IsUnknown() {
		return nil, false
	}
	var v []MiniVectorIndex
	d := m.VectorIndexes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVectorIndexes sets the value of the VectorIndexes field in ListVectorIndexesResponse.
func (m *ListVectorIndexesResponse) SetVectorIndexes(ctx context.Context, v []MiniVectorIndex) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["vector_indexes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.VectorIndexes = types.ListValueMust(t, vs)
}

// Key-value pair.
type MapStringValueEntry struct {
	// Column name.
	Key types.String `tfsdk:"key"`
	// Column value, nullable.
	Value types.Object `tfsdk:"value"`
}

func (to *MapStringValueEntry) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MapStringValueEntry) {
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

func (to *MapStringValueEntry) SyncFieldsDuringRead(ctx context.Context, from MapStringValueEntry) {
	if !from.Value.IsNull() && !from.Value.IsUnknown() {
		if toValue, ok := to.GetValue(ctx); ok {
			if fromValue, ok := from.GetValue(ctx); ok {
				toValue.SyncFieldsDuringRead(ctx, fromValue)
				to.SetValue(ctx, toValue)
			}
		}
	}
}

func (m MapStringValueEntry) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MapStringValueEntry.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m MapStringValueEntry) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"value": reflect.TypeOf(Value{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MapStringValueEntry
// only implements ToObjectValue() and Type().
func (m MapStringValueEntry) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m MapStringValueEntry) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": Value{}.Type(ctx),
		},
	}
}

// GetValue returns the value of the Value field in MapStringValueEntry as
// a Value value.
// If the field is unknown or null, the boolean return value is false.
func (m *MapStringValueEntry) GetValue(ctx context.Context) (Value, bool) {
	var e Value
	if m.Value.IsNull() || m.Value.IsUnknown() {
		return e, false
	}
	var v Value
	d := m.Value.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValue sets the value of the Value field in MapStringValueEntry.
func (m *MapStringValueEntry) SetValue(ctx context.Context, v Value) {
	vs := v.ToObjectValue(ctx)
	m.Value = vs
}

type MiniVectorIndex struct {
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

func (to *MiniVectorIndex) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MiniVectorIndex) {
}

func (to *MiniVectorIndex) SyncFieldsDuringRead(ctx context.Context, from MiniVectorIndex) {
}

func (m MiniVectorIndex) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m MiniVectorIndex) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MiniVectorIndex
// only implements ToObjectValue() and Type().
func (m MiniVectorIndex) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m MiniVectorIndex) Type(ctx context.Context) attr.Type {
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

type PatchEndpointBudgetPolicyRequest struct {
	// The budget policy id to be applied (hima-sheth) TODO: remove this once
	// we've migrated to usage policies
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// Name of the vector search endpoint
	EndpointName types.String `tfsdk:"-"`
}

func (to *PatchEndpointBudgetPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PatchEndpointBudgetPolicyRequest) {
}

func (to *PatchEndpointBudgetPolicyRequest) SyncFieldsDuringRead(ctx context.Context, from PatchEndpointBudgetPolicyRequest) {
}

func (m PatchEndpointBudgetPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PatchEndpointBudgetPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchEndpointBudgetPolicyRequest
// only implements ToObjectValue() and Type().
func (m PatchEndpointBudgetPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget_policy_id": m.BudgetPolicyId,
			"endpoint_name":    m.EndpointName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PatchEndpointBudgetPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_policy_id": types.StringType,
			"endpoint_name":    types.StringType,
		},
	}
}

type PatchEndpointBudgetPolicyResponse struct {
	// The budget policy applied to the vector search endpoint.
	EffectiveBudgetPolicyId types.String `tfsdk:"effective_budget_policy_id"`
}

func (to *PatchEndpointBudgetPolicyResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PatchEndpointBudgetPolicyResponse) {
}

func (to *PatchEndpointBudgetPolicyResponse) SyncFieldsDuringRead(ctx context.Context, from PatchEndpointBudgetPolicyResponse) {
}

func (m PatchEndpointBudgetPolicyResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PatchEndpointBudgetPolicyResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchEndpointBudgetPolicyResponse
// only implements ToObjectValue() and Type().
func (m PatchEndpointBudgetPolicyResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"effective_budget_policy_id": m.EffectiveBudgetPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PatchEndpointBudgetPolicyResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"effective_budget_policy_id": types.StringType,
		},
	}
}

// Request payload for getting next page of results.
type QueryVectorIndexNextPageRequest struct {
	// Name of the endpoint.
	EndpointName types.String `tfsdk:"endpoint_name"`
	// Name of the vector index to query.
	IndexName types.String `tfsdk:"-"`
	// Page token returned from previous `QueryVectorIndex` or
	// `QueryVectorIndexNextPage` API.
	PageToken types.String `tfsdk:"page_token"`
}

func (to *QueryVectorIndexNextPageRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryVectorIndexNextPageRequest) {
}

func (to *QueryVectorIndexNextPageRequest) SyncFieldsDuringRead(ctx context.Context, from QueryVectorIndexNextPageRequest) {
}

func (m QueryVectorIndexNextPageRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m QueryVectorIndexNextPageRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryVectorIndexNextPageRequest
// only implements ToObjectValue() and Type().
func (m QueryVectorIndexNextPageRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_name": m.EndpointName,
			"index_name":    m.IndexName,
			"page_token":    m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryVectorIndexNextPageRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_name": types.StringType,
			"index_name":    types.StringType,
			"page_token":    types.StringType,
		},
	}
}

type QueryVectorIndexRequest struct {
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

	Reranker types.Object `tfsdk:"reranker"`
	// Threshold for the approximate nearest neighbor search. Defaults to 0.0.
	ScoreThreshold types.Float64 `tfsdk:"score_threshold"`
}

func (to *QueryVectorIndexRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryVectorIndexRequest) {
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

func (to *QueryVectorIndexRequest) SyncFieldsDuringRead(ctx context.Context, from QueryVectorIndexRequest) {
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

func (m QueryVectorIndexRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["columns"] = attrs["columns"].SetRequired()
	attrs["columns_to_rerank"] = attrs["columns_to_rerank"].SetOptional()
	attrs["filters_json"] = attrs["filters_json"].SetOptional()
	attrs["num_results"] = attrs["num_results"].SetOptional()
	attrs["query_text"] = attrs["query_text"].SetOptional()
	attrs["query_type"] = attrs["query_type"].SetOptional()
	attrs["query_vector"] = attrs["query_vector"].SetOptional()
	attrs["reranker"] = attrs["reranker"].SetOptional()
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
func (m QueryVectorIndexRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns":           reflect.TypeOf(types.String{}),
		"columns_to_rerank": reflect.TypeOf(types.String{}),
		"query_vector":      reflect.TypeOf(types.Float64{}),
		"reranker":          reflect.TypeOf(RerankerConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryVectorIndexRequest
// only implements ToObjectValue() and Type().
func (m QueryVectorIndexRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m QueryVectorIndexRequest) Type(ctx context.Context) attr.Type {
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
			"reranker":        RerankerConfig{}.Type(ctx),
			"score_threshold": types.Float64Type,
		},
	}
}

// GetColumns returns the value of the Columns field in QueryVectorIndexRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryVectorIndexRequest) GetColumns(ctx context.Context) ([]types.String, bool) {
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

// SetColumns sets the value of the Columns field in QueryVectorIndexRequest.
func (m *QueryVectorIndexRequest) SetColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Columns = types.ListValueMust(t, vs)
}

// GetColumnsToRerank returns the value of the ColumnsToRerank field in QueryVectorIndexRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryVectorIndexRequest) GetColumnsToRerank(ctx context.Context) ([]types.String, bool) {
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

// SetColumnsToRerank sets the value of the ColumnsToRerank field in QueryVectorIndexRequest.
func (m *QueryVectorIndexRequest) SetColumnsToRerank(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns_to_rerank"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ColumnsToRerank = types.ListValueMust(t, vs)
}

// GetQueryVector returns the value of the QueryVector field in QueryVectorIndexRequest as
// a slice of types.Float64 values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryVectorIndexRequest) GetQueryVector(ctx context.Context) ([]types.Float64, bool) {
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

// SetQueryVector sets the value of the QueryVector field in QueryVectorIndexRequest.
func (m *QueryVectorIndexRequest) SetQueryVector(ctx context.Context, v []types.Float64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["query_vector"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.QueryVector = types.ListValueMust(t, vs)
}

// GetReranker returns the value of the Reranker field in QueryVectorIndexRequest as
// a RerankerConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryVectorIndexRequest) GetReranker(ctx context.Context) (RerankerConfig, bool) {
	var e RerankerConfig
	if m.Reranker.IsNull() || m.Reranker.IsUnknown() {
		return e, false
	}
	var v RerankerConfig
	d := m.Reranker.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetReranker sets the value of the Reranker field in QueryVectorIndexRequest.
func (m *QueryVectorIndexRequest) SetReranker(ctx context.Context, v RerankerConfig) {
	vs := v.ToObjectValue(ctx)
	m.Reranker = vs
}

type QueryVectorIndexResponse struct {
	// Metadata about the result set.
	Manifest types.Object `tfsdk:"manifest"`
	// [Optional] Token that can be used in `QueryVectorIndexNextPage` API to
	// get next page of results. If more than 1000 results satisfy the query,
	// they are returned in groups of 1000. Empty value means no more results.
	// The maximum number of results that can be returned is 10,000.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Data returned in the query result.
	Result types.Object `tfsdk:"result"`
}

func (to *QueryVectorIndexResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryVectorIndexResponse) {
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

func (to *QueryVectorIndexResponse) SyncFieldsDuringRead(ctx context.Context, from QueryVectorIndexResponse) {
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

func (m QueryVectorIndexResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["manifest"] = attrs["manifest"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["result"] = attrs["result"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryVectorIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m QueryVectorIndexResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"manifest": reflect.TypeOf(ResultManifest{}),
		"result":   reflect.TypeOf(ResultData{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryVectorIndexResponse
// only implements ToObjectValue() and Type().
func (m QueryVectorIndexResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"manifest":        m.Manifest,
			"next_page_token": m.NextPageToken,
			"result":          m.Result,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryVectorIndexResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"manifest":        ResultManifest{}.Type(ctx),
			"next_page_token": types.StringType,
			"result":          ResultData{}.Type(ctx),
		},
	}
}

// GetManifest returns the value of the Manifest field in QueryVectorIndexResponse as
// a ResultManifest value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryVectorIndexResponse) GetManifest(ctx context.Context) (ResultManifest, bool) {
	var e ResultManifest
	if m.Manifest.IsNull() || m.Manifest.IsUnknown() {
		return e, false
	}
	var v ResultManifest
	d := m.Manifest.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetManifest sets the value of the Manifest field in QueryVectorIndexResponse.
func (m *QueryVectorIndexResponse) SetManifest(ctx context.Context, v ResultManifest) {
	vs := v.ToObjectValue(ctx)
	m.Manifest = vs
}

// GetResult returns the value of the Result field in QueryVectorIndexResponse as
// a ResultData value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryVectorIndexResponse) GetResult(ctx context.Context) (ResultData, bool) {
	var e ResultData
	if m.Result.IsNull() || m.Result.IsUnknown() {
		return e, false
	}
	var v ResultData
	d := m.Result.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResult sets the value of the Result field in QueryVectorIndexResponse.
func (m *QueryVectorIndexResponse) SetResult(ctx context.Context, v ResultData) {
	vs := v.ToObjectValue(ctx)
	m.Result = vs
}

type RerankerConfig struct {
	Model types.String `tfsdk:"model"`

	Parameters types.Object `tfsdk:"parameters"`
}

func (to *RerankerConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RerankerConfig) {
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

func (to *RerankerConfig) SyncFieldsDuringRead(ctx context.Context, from RerankerConfig) {
	if !from.Parameters.IsNull() && !from.Parameters.IsUnknown() {
		if toParameters, ok := to.GetParameters(ctx); ok {
			if fromParameters, ok := from.GetParameters(ctx); ok {
				toParameters.SyncFieldsDuringRead(ctx, fromParameters)
				to.SetParameters(ctx, toParameters)
			}
		}
	}
}

func (m RerankerConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model"] = attrs["model"].SetOptional()
	attrs["parameters"] = attrs["parameters"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RerankerConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RerankerConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"parameters": reflect.TypeOf(RerankerConfigRerankerParameters{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RerankerConfig
// only implements ToObjectValue() and Type().
func (m RerankerConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model":      m.Model,
			"parameters": m.Parameters,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RerankerConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model":      types.StringType,
			"parameters": RerankerConfigRerankerParameters{}.Type(ctx),
		},
	}
}

// GetParameters returns the value of the Parameters field in RerankerConfig as
// a RerankerConfigRerankerParameters value.
// If the field is unknown or null, the boolean return value is false.
func (m *RerankerConfig) GetParameters(ctx context.Context) (RerankerConfigRerankerParameters, bool) {
	var e RerankerConfigRerankerParameters
	if m.Parameters.IsNull() || m.Parameters.IsUnknown() {
		return e, false
	}
	var v RerankerConfigRerankerParameters
	d := m.Parameters.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParameters sets the value of the Parameters field in RerankerConfig.
func (m *RerankerConfig) SetParameters(ctx context.Context, v RerankerConfigRerankerParameters) {
	vs := v.ToObjectValue(ctx)
	m.Parameters = vs
}

type RerankerConfigRerankerParameters struct {
	ColumnsToRerank types.List `tfsdk:"columns_to_rerank"`
}

func (to *RerankerConfigRerankerParameters) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RerankerConfigRerankerParameters) {
	if !from.ColumnsToRerank.IsNull() && !from.ColumnsToRerank.IsUnknown() && to.ColumnsToRerank.IsNull() && len(from.ColumnsToRerank.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnsToRerank, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnsToRerank = from.ColumnsToRerank
	}
}

func (to *RerankerConfigRerankerParameters) SyncFieldsDuringRead(ctx context.Context, from RerankerConfigRerankerParameters) {
	if !from.ColumnsToRerank.IsNull() && !from.ColumnsToRerank.IsUnknown() && to.ColumnsToRerank.IsNull() && len(from.ColumnsToRerank.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnsToRerank, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnsToRerank = from.ColumnsToRerank
	}
}

func (m RerankerConfigRerankerParameters) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RerankerConfigRerankerParameters) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns_to_rerank": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RerankerConfigRerankerParameters
// only implements ToObjectValue() and Type().
func (m RerankerConfigRerankerParameters) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns_to_rerank": m.ColumnsToRerank,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RerankerConfigRerankerParameters) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns_to_rerank": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetColumnsToRerank returns the value of the ColumnsToRerank field in RerankerConfigRerankerParameters as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *RerankerConfigRerankerParameters) GetColumnsToRerank(ctx context.Context) ([]types.String, bool) {
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

// SetColumnsToRerank sets the value of the ColumnsToRerank field in RerankerConfigRerankerParameters.
func (m *RerankerConfigRerankerParameters) SetColumnsToRerank(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns_to_rerank"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ColumnsToRerank = types.ListValueMust(t, vs)
}

// Data returned in the query result.
type ResultData struct {
	// Data rows returned in the query.
	DataArray types.List `tfsdk:"data_array"`
	// Number of rows in the result set.
	RowCount types.Int64 `tfsdk:"row_count"`
}

func (to *ResultData) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResultData) {
	if !from.DataArray.IsNull() && !from.DataArray.IsUnknown() && to.DataArray.IsNull() && len(from.DataArray.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DataArray, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DataArray = from.DataArray
	}
}

func (to *ResultData) SyncFieldsDuringRead(ctx context.Context, from ResultData) {
	if !from.DataArray.IsNull() && !from.DataArray.IsUnknown() && to.DataArray.IsNull() && len(from.DataArray.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DataArray, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DataArray = from.DataArray
	}
}

func (m ResultData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ResultData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_array": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultData
// only implements ToObjectValue() and Type().
func (m ResultData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_array": m.DataArray,
			"row_count":  m.RowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResultData) Type(ctx context.Context) attr.Type {
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

// GetDataArray returns the value of the DataArray field in ResultData as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ResultData) GetDataArray(ctx context.Context) ([]types.String, bool) {
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

// SetDataArray sets the value of the DataArray field in ResultData.
func (m *ResultData) SetDataArray(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["data_array"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DataArray = types.ListValueMust(t, vs)
}

// Metadata about the result set.
type ResultManifest struct {
	// Number of columns in the result set.
	ColumnCount types.Int64 `tfsdk:"column_count"`
	// Information about each column in the result set.
	Columns types.List `tfsdk:"columns"`
}

func (to *ResultManifest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResultManifest) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
}

func (to *ResultManifest) SyncFieldsDuringRead(ctx context.Context, from ResultManifest) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
}

func (m ResultManifest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ResultManifest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(ColumnInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultManifest
// only implements ToObjectValue() and Type().
func (m ResultManifest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column_count": m.ColumnCount,
			"columns":      m.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ResultManifest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"column_count": types.Int64Type,
			"columns": basetypes.ListType{
				ElemType: ColumnInfo{}.Type(ctx),
			},
		},
	}
}

// GetColumns returns the value of the Columns field in ResultManifest as
// a slice of ColumnInfo values.
// If the field is unknown or null, the boolean return value is false.
func (m *ResultManifest) GetColumns(ctx context.Context) ([]ColumnInfo, bool) {
	if m.Columns.IsNull() || m.Columns.IsUnknown() {
		return nil, false
	}
	var v []ColumnInfo
	d := m.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in ResultManifest.
func (m *ResultManifest) SetColumns(ctx context.Context, v []ColumnInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Columns = types.ListValueMust(t, vs)
}

type ScanVectorIndexRequest struct {
	// Name of the vector index to scan.
	IndexName types.String `tfsdk:"-"`
	// Primary key of the last entry returned in the previous scan.
	LastPrimaryKey types.String `tfsdk:"last_primary_key"`
	// Number of results to return. Defaults to 10.
	NumResults types.Int64 `tfsdk:"num_results"`
}

func (to *ScanVectorIndexRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ScanVectorIndexRequest) {
}

func (to *ScanVectorIndexRequest) SyncFieldsDuringRead(ctx context.Context, from ScanVectorIndexRequest) {
}

func (m ScanVectorIndexRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ScanVectorIndexRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ScanVectorIndexRequest
// only implements ToObjectValue() and Type().
func (m ScanVectorIndexRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name":       m.IndexName,
			"last_primary_key": m.LastPrimaryKey,
			"num_results":      m.NumResults,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ScanVectorIndexRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name":       types.StringType,
			"last_primary_key": types.StringType,
			"num_results":      types.Int64Type,
		},
	}
}

// Response to a scan vector index request.
type ScanVectorIndexResponse struct {
	// List of data entries
	Data types.List `tfsdk:"data"`
	// Primary key of the last entry.
	LastPrimaryKey types.String `tfsdk:"last_primary_key"`
}

func (to *ScanVectorIndexResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ScanVectorIndexResponse) {
	if !from.Data.IsNull() && !from.Data.IsUnknown() && to.Data.IsNull() && len(from.Data.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Data, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Data = from.Data
	}
}

func (to *ScanVectorIndexResponse) SyncFieldsDuringRead(ctx context.Context, from ScanVectorIndexResponse) {
	if !from.Data.IsNull() && !from.Data.IsUnknown() && to.Data.IsNull() && len(from.Data.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Data, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Data = from.Data
	}
}

func (m ScanVectorIndexResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ScanVectorIndexResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data": reflect.TypeOf(Struct{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ScanVectorIndexResponse
// only implements ToObjectValue() and Type().
func (m ScanVectorIndexResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data":             m.Data,
			"last_primary_key": m.LastPrimaryKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ScanVectorIndexResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data": basetypes.ListType{
				ElemType: Struct{}.Type(ctx),
			},
			"last_primary_key": types.StringType,
		},
	}
}

// GetData returns the value of the Data field in ScanVectorIndexResponse as
// a slice of Struct values.
// If the field is unknown or null, the boolean return value is false.
func (m *ScanVectorIndexResponse) GetData(ctx context.Context) ([]Struct, bool) {
	if m.Data.IsNull() || m.Data.IsUnknown() {
		return nil, false
	}
	var v []Struct
	d := m.Data.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetData sets the value of the Data field in ScanVectorIndexResponse.
func (m *ScanVectorIndexResponse) SetData(ctx context.Context, v []Struct) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["data"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Data = types.ListValueMust(t, vs)
}

type Struct struct {
	// Data entry, corresponding to a row in a vector index.
	Fields types.List `tfsdk:"fields"`
}

func (to *Struct) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Struct) {
	if !from.Fields.IsNull() && !from.Fields.IsUnknown() && to.Fields.IsNull() && len(from.Fields.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Fields, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Fields = from.Fields
	}
}

func (to *Struct) SyncFieldsDuringRead(ctx context.Context, from Struct) {
	if !from.Fields.IsNull() && !from.Fields.IsUnknown() && to.Fields.IsNull() && len(from.Fields.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Fields, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Fields = from.Fields
	}
}

func (m Struct) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Struct) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"fields": reflect.TypeOf(MapStringValueEntry{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Struct
// only implements ToObjectValue() and Type().
func (m Struct) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"fields": m.Fields,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Struct) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"fields": basetypes.ListType{
				ElemType: MapStringValueEntry{}.Type(ctx),
			},
		},
	}
}

// GetFields returns the value of the Fields field in Struct as
// a slice of MapStringValueEntry values.
// If the field is unknown or null, the boolean return value is false.
func (m *Struct) GetFields(ctx context.Context) ([]MapStringValueEntry, bool) {
	if m.Fields.IsNull() || m.Fields.IsUnknown() {
		return nil, false
	}
	var v []MapStringValueEntry
	d := m.Fields.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFields sets the value of the Fields field in Struct.
func (m *Struct) SetFields(ctx context.Context, v []MapStringValueEntry) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["fields"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Fields = types.ListValueMust(t, vs)
}

type SyncIndexRequest struct {
	// Name of the vector index to synchronize. Must be a Delta Sync Index.
	IndexName types.String `tfsdk:"-"`
}

func (to *SyncIndexRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncIndexRequest) {
}

func (to *SyncIndexRequest) SyncFieldsDuringRead(ctx context.Context, from SyncIndexRequest) {
}

func (m SyncIndexRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SyncIndexRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncIndexRequest
// only implements ToObjectValue() and Type().
func (m SyncIndexRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name": m.IndexName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SyncIndexRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name": types.StringType,
		},
	}
}

type SyncIndexResponse struct {
}

func (to *SyncIndexResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncIndexResponse) {
}

func (to *SyncIndexResponse) SyncFieldsDuringRead(ctx context.Context, from SyncIndexResponse) {
}

func (m SyncIndexResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SyncIndexResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncIndexResponse
// only implements ToObjectValue() and Type().
func (m SyncIndexResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m SyncIndexResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateEndpointCustomTagsRequest struct {
	// The new custom tags for the vector search endpoint
	CustomTags types.List `tfsdk:"custom_tags"`
	// Name of the vector search endpoint
	EndpointName types.String `tfsdk:"-"`
}

func (to *UpdateEndpointCustomTagsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEndpointCustomTagsRequest) {
}

func (to *UpdateEndpointCustomTagsRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateEndpointCustomTagsRequest) {
}

func (m UpdateEndpointCustomTagsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateEndpointCustomTagsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_tags": reflect.TypeOf(CustomTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEndpointCustomTagsRequest
// only implements ToObjectValue() and Type().
func (m UpdateEndpointCustomTagsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"custom_tags":   m.CustomTags,
			"endpoint_name": m.EndpointName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateEndpointCustomTagsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"custom_tags": basetypes.ListType{
				ElemType: CustomTag{}.Type(ctx),
			},
			"endpoint_name": types.StringType,
		},
	}
}

// GetCustomTags returns the value of the CustomTags field in UpdateEndpointCustomTagsRequest as
// a slice of CustomTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateEndpointCustomTagsRequest) GetCustomTags(ctx context.Context) ([]CustomTag, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v []CustomTag
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in UpdateEndpointCustomTagsRequest.
func (m *UpdateEndpointCustomTagsRequest) SetCustomTags(ctx context.Context, v []CustomTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.ListValueMust(t, vs)
}

type UpdateEndpointCustomTagsResponse struct {
	// All the custom tags that are applied to the vector search endpoint.
	CustomTags types.List `tfsdk:"custom_tags"`
	// The name of the vector search endpoint whose custom tags were updated.
	Name types.String `tfsdk:"name"`
}

func (to *UpdateEndpointCustomTagsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEndpointCustomTagsResponse) {
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
}

func (to *UpdateEndpointCustomTagsResponse) SyncFieldsDuringRead(ctx context.Context, from UpdateEndpointCustomTagsResponse) {
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
}

func (m UpdateEndpointCustomTagsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateEndpointCustomTagsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_tags": reflect.TypeOf(CustomTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEndpointCustomTagsResponse
// only implements ToObjectValue() and Type().
func (m UpdateEndpointCustomTagsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"custom_tags": m.CustomTags,
			"name":        m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateEndpointCustomTagsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"custom_tags": basetypes.ListType{
				ElemType: CustomTag{}.Type(ctx),
			},
			"name": types.StringType,
		},
	}
}

// GetCustomTags returns the value of the CustomTags field in UpdateEndpointCustomTagsResponse as
// a slice of CustomTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateEndpointCustomTagsResponse) GetCustomTags(ctx context.Context) ([]CustomTag, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v []CustomTag
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in UpdateEndpointCustomTagsResponse.
func (m *UpdateEndpointCustomTagsResponse) SetCustomTags(ctx context.Context, v []CustomTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.ListValueMust(t, vs)
}

type UpsertDataResult struct {
	// List of primary keys for rows that failed to process.
	FailedPrimaryKeys types.List `tfsdk:"failed_primary_keys"`
	// Count of successfully processed rows.
	SuccessRowCount types.Int64 `tfsdk:"success_row_count"`
}

func (to *UpsertDataResult) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpsertDataResult) {
	if !from.FailedPrimaryKeys.IsNull() && !from.FailedPrimaryKeys.IsUnknown() && to.FailedPrimaryKeys.IsNull() && len(from.FailedPrimaryKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FailedPrimaryKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FailedPrimaryKeys = from.FailedPrimaryKeys
	}
}

func (to *UpsertDataResult) SyncFieldsDuringRead(ctx context.Context, from UpsertDataResult) {
	if !from.FailedPrimaryKeys.IsNull() && !from.FailedPrimaryKeys.IsUnknown() && to.FailedPrimaryKeys.IsNull() && len(from.FailedPrimaryKeys.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FailedPrimaryKeys, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FailedPrimaryKeys = from.FailedPrimaryKeys
	}
}

func (m UpsertDataResult) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpsertDataResult) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"failed_primary_keys": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpsertDataResult
// only implements ToObjectValue() and Type().
func (m UpsertDataResult) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"failed_primary_keys": m.FailedPrimaryKeys,
			"success_row_count":   m.SuccessRowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpsertDataResult) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"failed_primary_keys": basetypes.ListType{
				ElemType: types.StringType,
			},
			"success_row_count": types.Int64Type,
		},
	}
}

// GetFailedPrimaryKeys returns the value of the FailedPrimaryKeys field in UpsertDataResult as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpsertDataResult) GetFailedPrimaryKeys(ctx context.Context) ([]types.String, bool) {
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

// SetFailedPrimaryKeys sets the value of the FailedPrimaryKeys field in UpsertDataResult.
func (m *UpsertDataResult) SetFailedPrimaryKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["failed_primary_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FailedPrimaryKeys = types.ListValueMust(t, vs)
}

type UpsertDataVectorIndexRequest struct {
	// Name of the vector index where data is to be upserted. Must be a Direct
	// Vector Access Index.
	IndexName types.String `tfsdk:"-"`
	// JSON string representing the data to be upserted.
	InputsJson types.String `tfsdk:"inputs_json"`
}

func (to *UpsertDataVectorIndexRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpsertDataVectorIndexRequest) {
}

func (to *UpsertDataVectorIndexRequest) SyncFieldsDuringRead(ctx context.Context, from UpsertDataVectorIndexRequest) {
}

func (m UpsertDataVectorIndexRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpsertDataVectorIndexRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpsertDataVectorIndexRequest
// only implements ToObjectValue() and Type().
func (m UpsertDataVectorIndexRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name":  m.IndexName,
			"inputs_json": m.InputsJson,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpsertDataVectorIndexRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name":  types.StringType,
			"inputs_json": types.StringType,
		},
	}
}

type UpsertDataVectorIndexResponse struct {
	// Result of the upsert or delete operation.
	Result types.Object `tfsdk:"result"`
	// Status of the upsert operation.
	Status types.String `tfsdk:"status"`
}

func (to *UpsertDataVectorIndexResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpsertDataVectorIndexResponse) {
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

func (to *UpsertDataVectorIndexResponse) SyncFieldsDuringRead(ctx context.Context, from UpsertDataVectorIndexResponse) {
	if !from.Result.IsNull() && !from.Result.IsUnknown() {
		if toResult, ok := to.GetResult(ctx); ok {
			if fromResult, ok := from.GetResult(ctx); ok {
				toResult.SyncFieldsDuringRead(ctx, fromResult)
				to.SetResult(ctx, toResult)
			}
		}
	}
}

func (m UpsertDataVectorIndexResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["result"] = attrs["result"].SetOptional()
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
func (m UpsertDataVectorIndexResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"result": reflect.TypeOf(UpsertDataResult{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpsertDataVectorIndexResponse
// only implements ToObjectValue() and Type().
func (m UpsertDataVectorIndexResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"result": m.Result,
			"status": m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpsertDataVectorIndexResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"result": UpsertDataResult{}.Type(ctx),
			"status": types.StringType,
		},
	}
}

// GetResult returns the value of the Result field in UpsertDataVectorIndexResponse as
// a UpsertDataResult value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpsertDataVectorIndexResponse) GetResult(ctx context.Context) (UpsertDataResult, bool) {
	var e UpsertDataResult
	if m.Result.IsNull() || m.Result.IsUnknown() {
		return e, false
	}
	var v UpsertDataResult
	d := m.Result.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResult sets the value of the Result field in UpsertDataVectorIndexResponse.
func (m *UpsertDataVectorIndexResponse) SetResult(ctx context.Context, v UpsertDataResult) {
	vs := v.ToObjectValue(ctx)
	m.Result = vs
}

type Value struct {
	BoolValue types.Bool `tfsdk:"bool_value"`

	ListValue types.Object `tfsdk:"list_value"`

	NumberValue types.Float64 `tfsdk:"number_value"`

	StringValue types.String `tfsdk:"string_value"`

	StructValue types.Object `tfsdk:"struct_value"`
}

func (to *Value) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Value) {
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

func (to *Value) SyncFieldsDuringRead(ctx context.Context, from Value) {
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

func (m Value) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["bool_value"] = attrs["bool_value"].SetOptional()
	attrs["list_value"] = attrs["list_value"].SetOptional()
	attrs["number_value"] = attrs["number_value"].SetOptional()
	attrs["string_value"] = attrs["string_value"].SetOptional()
	attrs["struct_value"] = attrs["struct_value"].SetOptional()

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
	return map[string]reflect.Type{
		"list_value":   reflect.TypeOf(ListValue{}),
		"struct_value": reflect.TypeOf(Struct{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Value
// only implements ToObjectValue() and Type().
func (m Value) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Value) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bool_value":   types.BoolType,
			"list_value":   ListValue{}.Type(ctx),
			"number_value": types.Float64Type,
			"string_value": types.StringType,
			"struct_value": Struct{}.Type(ctx),
		},
	}
}

// GetListValue returns the value of the ListValue field in Value as
// a ListValue value.
// If the field is unknown or null, the boolean return value is false.
func (m *Value) GetListValue(ctx context.Context) (ListValue, bool) {
	var e ListValue
	if m.ListValue.IsNull() || m.ListValue.IsUnknown() {
		return e, false
	}
	var v ListValue
	d := m.ListValue.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetListValue sets the value of the ListValue field in Value.
func (m *Value) SetListValue(ctx context.Context, v ListValue) {
	vs := v.ToObjectValue(ctx)
	m.ListValue = vs
}

// GetStructValue returns the value of the StructValue field in Value as
// a Struct value.
// If the field is unknown or null, the boolean return value is false.
func (m *Value) GetStructValue(ctx context.Context) (Struct, bool) {
	var e Struct
	if m.StructValue.IsNull() || m.StructValue.IsUnknown() {
		return e, false
	}
	var v Struct
	d := m.StructValue.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStructValue sets the value of the StructValue field in Value.
func (m *Value) SetStructValue(ctx context.Context, v Struct) {
	vs := v.ToObjectValue(ctx)
	m.StructValue = vs
}

type VectorIndex struct {
	// The user who created the index.
	Creator types.String `tfsdk:"creator"`

	DeltaSyncIndexSpec types.Object `tfsdk:"delta_sync_index_spec"`

	DirectAccessIndexSpec types.Object `tfsdk:"direct_access_index_spec"`
	// Name of the endpoint associated with the index
	EndpointName types.String `tfsdk:"endpoint_name"`

	IndexType types.String `tfsdk:"index_type"`
	// Name of the index
	Name types.String `tfsdk:"name"`
	// Primary key of the index
	PrimaryKey types.String `tfsdk:"primary_key"`

	Status types.Object `tfsdk:"status"`
}

func (to *VectorIndex) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from VectorIndex) {
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

func (to *VectorIndex) SyncFieldsDuringRead(ctx context.Context, from VectorIndex) {
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

func (m VectorIndex) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creator"] = attrs["creator"].SetOptional()
	attrs["delta_sync_index_spec"] = attrs["delta_sync_index_spec"].SetOptional()
	attrs["direct_access_index_spec"] = attrs["direct_access_index_spec"].SetOptional()
	attrs["endpoint_name"] = attrs["endpoint_name"].SetOptional()
	attrs["index_type"] = attrs["index_type"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["primary_key"] = attrs["primary_key"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in VectorIndex.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m VectorIndex) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"delta_sync_index_spec":    reflect.TypeOf(DeltaSyncVectorIndexSpecResponse{}),
		"direct_access_index_spec": reflect.TypeOf(DirectAccessVectorIndexSpec{}),
		"status":                   reflect.TypeOf(VectorIndexStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VectorIndex
// only implements ToObjectValue() and Type().
func (m VectorIndex) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m VectorIndex) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creator":                  types.StringType,
			"delta_sync_index_spec":    DeltaSyncVectorIndexSpecResponse{}.Type(ctx),
			"direct_access_index_spec": DirectAccessVectorIndexSpec{}.Type(ctx),
			"endpoint_name":            types.StringType,
			"index_type":               types.StringType,
			"name":                     types.StringType,
			"primary_key":              types.StringType,
			"status":                   VectorIndexStatus{}.Type(ctx),
		},
	}
}

// GetDeltaSyncIndexSpec returns the value of the DeltaSyncIndexSpec field in VectorIndex as
// a DeltaSyncVectorIndexSpecResponse value.
// If the field is unknown or null, the boolean return value is false.
func (m *VectorIndex) GetDeltaSyncIndexSpec(ctx context.Context) (DeltaSyncVectorIndexSpecResponse, bool) {
	var e DeltaSyncVectorIndexSpecResponse
	if m.DeltaSyncIndexSpec.IsNull() || m.DeltaSyncIndexSpec.IsUnknown() {
		return e, false
	}
	var v DeltaSyncVectorIndexSpecResponse
	d := m.DeltaSyncIndexSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDeltaSyncIndexSpec sets the value of the DeltaSyncIndexSpec field in VectorIndex.
func (m *VectorIndex) SetDeltaSyncIndexSpec(ctx context.Context, v DeltaSyncVectorIndexSpecResponse) {
	vs := v.ToObjectValue(ctx)
	m.DeltaSyncIndexSpec = vs
}

// GetDirectAccessIndexSpec returns the value of the DirectAccessIndexSpec field in VectorIndex as
// a DirectAccessVectorIndexSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *VectorIndex) GetDirectAccessIndexSpec(ctx context.Context) (DirectAccessVectorIndexSpec, bool) {
	var e DirectAccessVectorIndexSpec
	if m.DirectAccessIndexSpec.IsNull() || m.DirectAccessIndexSpec.IsUnknown() {
		return e, false
	}
	var v DirectAccessVectorIndexSpec
	d := m.DirectAccessIndexSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDirectAccessIndexSpec sets the value of the DirectAccessIndexSpec field in VectorIndex.
func (m *VectorIndex) SetDirectAccessIndexSpec(ctx context.Context, v DirectAccessVectorIndexSpec) {
	vs := v.ToObjectValue(ctx)
	m.DirectAccessIndexSpec = vs
}

// GetStatus returns the value of the Status field in VectorIndex as
// a VectorIndexStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *VectorIndex) GetStatus(ctx context.Context) (VectorIndexStatus, bool) {
	var e VectorIndexStatus
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v VectorIndexStatus
	d := m.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in VectorIndex.
func (m *VectorIndex) SetStatus(ctx context.Context, v VectorIndexStatus) {
	vs := v.ToObjectValue(ctx)
	m.Status = vs
}

type VectorIndexStatus struct {
	// Index API Url to be used to perform operations on the index
	IndexUrl types.String `tfsdk:"index_url"`
	// Number of rows indexed
	IndexedRowCount types.Int64 `tfsdk:"indexed_row_count"`
	// Message associated with the index status
	Message types.String `tfsdk:"message"`
	// Whether the index is ready for search
	Ready types.Bool `tfsdk:"ready"`
}

func (to *VectorIndexStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from VectorIndexStatus) {
}

func (to *VectorIndexStatus) SyncFieldsDuringRead(ctx context.Context, from VectorIndexStatus) {
}

func (m VectorIndexStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m VectorIndexStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VectorIndexStatus
// only implements ToObjectValue() and Type().
func (m VectorIndexStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m VectorIndexStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_url":         types.StringType,
			"indexed_row_count": types.Int64Type,
			"message":           types.StringType,
			"ready":             types.BoolType,
		},
	}
}
