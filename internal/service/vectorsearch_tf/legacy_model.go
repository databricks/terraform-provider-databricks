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

func (newState *ColumnInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ColumnInfo_SdkV2) {
}

func (newState *ColumnInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState ColumnInfo_SdkV2) {
}

func (c ColumnInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ColumnInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ColumnInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o ColumnInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ColumnInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type CreateEndpoint_SdkV2 struct {
	// Type of endpoint.
	EndpointType types.String `tfsdk:"endpoint_type"`
	// Name of endpoint
	Name types.String `tfsdk:"name"`
}

func (newState *CreateEndpoint_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateEndpoint_SdkV2) {
}

func (newState *CreateEndpoint_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateEndpoint_SdkV2) {
}

func (c CreateEndpoint_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateEndpoint_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateEndpoint_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateEndpoint_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_type": o.EndpointType,
			"name":          o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateEndpoint_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_type": types.StringType,
			"name":          types.StringType,
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
	// There are 2 types of Vector Search indexes:
	//
	// - `DELTA_SYNC`: An index that automatically syncs with a source Delta
	// Table, automatically and incrementally updating the index as the
	// underlying data in the Delta Table changes. - `DIRECT_ACCESS`: An index
	// that supports direct read and write of vectors and metadata through our
	// REST and SDK APIs. With this model, the user manages index updates.
	IndexType types.String `tfsdk:"index_type"`
	// Name of the index
	Name types.String `tfsdk:"name"`
	// Primary key of the index
	PrimaryKey types.String `tfsdk:"primary_key"`
}

func (newState *CreateVectorIndexRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateVectorIndexRequest_SdkV2) {
}

func (newState *CreateVectorIndexRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateVectorIndexRequest_SdkV2) {
}

func (c CreateVectorIndexRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateVectorIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"delta_sync_index_spec":    reflect.TypeOf(DeltaSyncVectorIndexSpecRequest_SdkV2{}),
		"direct_access_index_spec": reflect.TypeOf(DirectAccessVectorIndexSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateVectorIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateVectorIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"delta_sync_index_spec":    o.DeltaSyncIndexSpec,
			"direct_access_index_spec": o.DirectAccessIndexSpec,
			"endpoint_name":            o.EndpointName,
			"index_type":               o.IndexType,
			"name":                     o.Name,
			"primary_key":              o.PrimaryKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateVectorIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateVectorIndexRequest_SdkV2) GetDeltaSyncIndexSpec(ctx context.Context) (DeltaSyncVectorIndexSpecRequest_SdkV2, bool) {
	var e DeltaSyncVectorIndexSpecRequest_SdkV2
	if o.DeltaSyncIndexSpec.IsNull() || o.DeltaSyncIndexSpec.IsUnknown() {
		return e, false
	}
	var v []DeltaSyncVectorIndexSpecRequest_SdkV2
	d := o.DeltaSyncIndexSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDeltaSyncIndexSpec sets the value of the DeltaSyncIndexSpec field in CreateVectorIndexRequest_SdkV2.
func (o *CreateVectorIndexRequest_SdkV2) SetDeltaSyncIndexSpec(ctx context.Context, v DeltaSyncVectorIndexSpecRequest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["delta_sync_index_spec"]
	o.DeltaSyncIndexSpec = types.ListValueMust(t, vs)
}

// GetDirectAccessIndexSpec returns the value of the DirectAccessIndexSpec field in CreateVectorIndexRequest_SdkV2 as
// a DirectAccessVectorIndexSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateVectorIndexRequest_SdkV2) GetDirectAccessIndexSpec(ctx context.Context) (DirectAccessVectorIndexSpec_SdkV2, bool) {
	var e DirectAccessVectorIndexSpec_SdkV2
	if o.DirectAccessIndexSpec.IsNull() || o.DirectAccessIndexSpec.IsUnknown() {
		return e, false
	}
	var v []DirectAccessVectorIndexSpec_SdkV2
	d := o.DirectAccessIndexSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDirectAccessIndexSpec sets the value of the DirectAccessIndexSpec field in CreateVectorIndexRequest_SdkV2.
func (o *CreateVectorIndexRequest_SdkV2) SetDirectAccessIndexSpec(ctx context.Context, v DirectAccessVectorIndexSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["direct_access_index_spec"]
	o.DirectAccessIndexSpec = types.ListValueMust(t, vs)
}

type CreateVectorIndexResponse_SdkV2 struct {
	VectorIndex types.List `tfsdk:"vector_index"`
}

func (newState *CreateVectorIndexResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateVectorIndexResponse_SdkV2) {
}

func (newState *CreateVectorIndexResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateVectorIndexResponse_SdkV2) {
}

func (c CreateVectorIndexResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["vector_index"] = attrs["vector_index"].SetOptional()
	attrs["vector_index"] = attrs["vector_index"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateVectorIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateVectorIndexResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"vector_index": reflect.TypeOf(VectorIndex_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateVectorIndexResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateVectorIndexResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"vector_index": o.VectorIndex,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateVectorIndexResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"vector_index": basetypes.ListType{
				ElemType: VectorIndex_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetVectorIndex returns the value of the VectorIndex field in CreateVectorIndexResponse_SdkV2 as
// a VectorIndex_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateVectorIndexResponse_SdkV2) GetVectorIndex(ctx context.Context) (VectorIndex_SdkV2, bool) {
	var e VectorIndex_SdkV2
	if o.VectorIndex.IsNull() || o.VectorIndex.IsUnknown() {
		return e, false
	}
	var v []VectorIndex_SdkV2
	d := o.VectorIndex.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVectorIndex sets the value of the VectorIndex field in CreateVectorIndexResponse_SdkV2.
func (o *CreateVectorIndexResponse_SdkV2) SetVectorIndex(ctx context.Context, v VectorIndex_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["vector_index"]
	o.VectorIndex = types.ListValueMust(t, vs)
}

// Result of the upsert or delete operation.
type DeleteDataResult_SdkV2 struct {
	// List of primary keys for rows that failed to process.
	FailedPrimaryKeys types.List `tfsdk:"failed_primary_keys"`
	// Count of successfully processed rows.
	SuccessRowCount types.Int64 `tfsdk:"success_row_count"`
}

func (newState *DeleteDataResult_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDataResult_SdkV2) {
}

func (newState *DeleteDataResult_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteDataResult_SdkV2) {
}

func (c DeleteDataResult_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteDataResult_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"failed_primary_keys": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDataResult_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDataResult_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"failed_primary_keys": o.FailedPrimaryKeys,
			"success_row_count":   o.SuccessRowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDataResult_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *DeleteDataResult_SdkV2) GetFailedPrimaryKeys(ctx context.Context) ([]types.String, bool) {
	if o.FailedPrimaryKeys.IsNull() || o.FailedPrimaryKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.FailedPrimaryKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFailedPrimaryKeys sets the value of the FailedPrimaryKeys field in DeleteDataResult_SdkV2.
func (o *DeleteDataResult_SdkV2) SetFailedPrimaryKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["failed_primary_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.FailedPrimaryKeys = types.ListValueMust(t, vs)
}

// Request payload for deleting data from a vector index.
type DeleteDataVectorIndexRequest_SdkV2 struct {
	// Name of the vector index where data is to be deleted. Must be a Direct
	// Vector Access Index.
	IndexName types.String `tfsdk:"-"`
	// List of primary keys for the data to be deleted.
	PrimaryKeys types.List `tfsdk:"primary_keys"`
}

func (newState *DeleteDataVectorIndexRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDataVectorIndexRequest_SdkV2) {
}

func (newState *DeleteDataVectorIndexRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteDataVectorIndexRequest_SdkV2) {
}

func (c DeleteDataVectorIndexRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteDataVectorIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"primary_keys": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDataVectorIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDataVectorIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name":   o.IndexName,
			"primary_keys": o.PrimaryKeys,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDataVectorIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *DeleteDataVectorIndexRequest_SdkV2) GetPrimaryKeys(ctx context.Context) ([]types.String, bool) {
	if o.PrimaryKeys.IsNull() || o.PrimaryKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.PrimaryKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrimaryKeys sets the value of the PrimaryKeys field in DeleteDataVectorIndexRequest_SdkV2.
func (o *DeleteDataVectorIndexRequest_SdkV2) SetPrimaryKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["primary_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PrimaryKeys = types.ListValueMust(t, vs)
}

// Response to a delete data vector index request.
type DeleteDataVectorIndexResponse_SdkV2 struct {
	// Result of the upsert or delete operation.
	Result types.List `tfsdk:"result"`
	// Status of the delete operation.
	Status types.String `tfsdk:"status"`
}

func (newState *DeleteDataVectorIndexResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDataVectorIndexResponse_SdkV2) {
}

func (newState *DeleteDataVectorIndexResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteDataVectorIndexResponse_SdkV2) {
}

func (c DeleteDataVectorIndexResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteDataVectorIndexResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"result": reflect.TypeOf(DeleteDataResult_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDataVectorIndexResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDataVectorIndexResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"result": o.Result,
			"status": o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDataVectorIndexResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *DeleteDataVectorIndexResponse_SdkV2) GetResult(ctx context.Context) (DeleteDataResult_SdkV2, bool) {
	var e DeleteDataResult_SdkV2
	if o.Result.IsNull() || o.Result.IsUnknown() {
		return e, false
	}
	var v []DeleteDataResult_SdkV2
	d := o.Result.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetResult sets the value of the Result field in DeleteDataVectorIndexResponse_SdkV2.
func (o *DeleteDataVectorIndexResponse_SdkV2) SetResult(ctx context.Context, v DeleteDataResult_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["result"]
	o.Result = types.ListValueMust(t, vs)
}

// Delete an endpoint
type DeleteEndpointRequest_SdkV2 struct {
	// Name of the endpoint
	EndpointName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_name": o.EndpointName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_name": types.StringType,
		},
	}
}

type DeleteEndpointResponse_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteEndpointResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteEndpointResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteEndpointResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteEndpointResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteEndpointResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete an index
type DeleteIndexRequest_SdkV2 struct {
	// Name of the index
	IndexName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name": o.IndexName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name": types.StringType,
		},
	}
}

type DeleteIndexResponse_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteIndexResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteIndexResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteIndexResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteIndexResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
	// The columns that contain the embedding source.
	EmbeddingSourceColumns types.List `tfsdk:"embedding_source_columns"`
	// The columns that contain the embedding vectors. The format should be
	// array[double].
	EmbeddingVectorColumns types.List `tfsdk:"embedding_vector_columns"`
	// [Optional] Automatically sync the vector index contents and computed
	// embeddings to the specified Delta table. The only supported table name is
	// the index name with the suffix `_writeback_table`.
	EmbeddingWritebackTable types.String `tfsdk:"embedding_writeback_table"`
	// Pipeline execution mode.
	//
	// - `TRIGGERED`: If the pipeline uses the triggered execution mode, the
	// system stops processing after successfully refreshing the source table in
	// the pipeline once, ensuring the table is updated based on the data
	// available when the update started. - `CONTINUOUS`: If the pipeline uses
	// continuous execution, the pipeline processes new data as it arrives in
	// the source table to keep vector index fresh.
	PipelineType types.String `tfsdk:"pipeline_type"`
	// The name of the source table.
	SourceTable types.String `tfsdk:"source_table"`
}

func (newState *DeltaSyncVectorIndexSpecRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeltaSyncVectorIndexSpecRequest_SdkV2) {
}

func (newState *DeltaSyncVectorIndexSpecRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeltaSyncVectorIndexSpecRequest_SdkV2) {
}

func (c DeltaSyncVectorIndexSpecRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeltaSyncVectorIndexSpecRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns_to_sync":          reflect.TypeOf(types.String{}),
		"embedding_source_columns": reflect.TypeOf(EmbeddingSourceColumn_SdkV2{}),
		"embedding_vector_columns": reflect.TypeOf(EmbeddingVectorColumn_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaSyncVectorIndexSpecRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeltaSyncVectorIndexSpecRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns_to_sync":           o.ColumnsToSync,
			"embedding_source_columns":  o.EmbeddingSourceColumns,
			"embedding_vector_columns":  o.EmbeddingVectorColumns,
			"embedding_writeback_table": o.EmbeddingWritebackTable,
			"pipeline_type":             o.PipelineType,
			"source_table":              o.SourceTable,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeltaSyncVectorIndexSpecRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns_to_sync": basetypes.ListType{
				ElemType: types.StringType,
			},
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
func (o *DeltaSyncVectorIndexSpecRequest_SdkV2) GetColumnsToSync(ctx context.Context) ([]types.String, bool) {
	if o.ColumnsToSync.IsNull() || o.ColumnsToSync.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ColumnsToSync.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumnsToSync sets the value of the ColumnsToSync field in DeltaSyncVectorIndexSpecRequest_SdkV2.
func (o *DeltaSyncVectorIndexSpecRequest_SdkV2) SetColumnsToSync(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["columns_to_sync"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ColumnsToSync = types.ListValueMust(t, vs)
}

// GetEmbeddingSourceColumns returns the value of the EmbeddingSourceColumns field in DeltaSyncVectorIndexSpecRequest_SdkV2 as
// a slice of EmbeddingSourceColumn_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *DeltaSyncVectorIndexSpecRequest_SdkV2) GetEmbeddingSourceColumns(ctx context.Context) ([]EmbeddingSourceColumn_SdkV2, bool) {
	if o.EmbeddingSourceColumns.IsNull() || o.EmbeddingSourceColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingSourceColumn_SdkV2
	d := o.EmbeddingSourceColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingSourceColumns sets the value of the EmbeddingSourceColumns field in DeltaSyncVectorIndexSpecRequest_SdkV2.
func (o *DeltaSyncVectorIndexSpecRequest_SdkV2) SetEmbeddingSourceColumns(ctx context.Context, v []EmbeddingSourceColumn_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_source_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EmbeddingSourceColumns = types.ListValueMust(t, vs)
}

// GetEmbeddingVectorColumns returns the value of the EmbeddingVectorColumns field in DeltaSyncVectorIndexSpecRequest_SdkV2 as
// a slice of EmbeddingVectorColumn_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *DeltaSyncVectorIndexSpecRequest_SdkV2) GetEmbeddingVectorColumns(ctx context.Context) ([]EmbeddingVectorColumn_SdkV2, bool) {
	if o.EmbeddingVectorColumns.IsNull() || o.EmbeddingVectorColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingVectorColumn_SdkV2
	d := o.EmbeddingVectorColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingVectorColumns sets the value of the EmbeddingVectorColumns field in DeltaSyncVectorIndexSpecRequest_SdkV2.
func (o *DeltaSyncVectorIndexSpecRequest_SdkV2) SetEmbeddingVectorColumns(ctx context.Context, v []EmbeddingVectorColumn_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_vector_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EmbeddingVectorColumns = types.ListValueMust(t, vs)
}

type DeltaSyncVectorIndexSpecResponse_SdkV2 struct {
	// The columns that contain the embedding source.
	EmbeddingSourceColumns types.List `tfsdk:"embedding_source_columns"`
	// The columns that contain the embedding vectors.
	EmbeddingVectorColumns types.List `tfsdk:"embedding_vector_columns"`
	// [Optional] Name of the Delta table to sync the vector index contents and
	// computed embeddings to.
	EmbeddingWritebackTable types.String `tfsdk:"embedding_writeback_table"`
	// The ID of the pipeline that is used to sync the index.
	PipelineId types.String `tfsdk:"pipeline_id"`
	// Pipeline execution mode.
	//
	// - `TRIGGERED`: If the pipeline uses the triggered execution mode, the
	// system stops processing after successfully refreshing the source table in
	// the pipeline once, ensuring the table is updated based on the data
	// available when the update started. - `CONTINUOUS`: If the pipeline uses
	// continuous execution, the pipeline processes new data as it arrives in
	// the source table to keep vector index fresh.
	PipelineType types.String `tfsdk:"pipeline_type"`
	// The name of the source table.
	SourceTable types.String `tfsdk:"source_table"`
}

func (newState *DeltaSyncVectorIndexSpecResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeltaSyncVectorIndexSpecResponse_SdkV2) {
}

func (newState *DeltaSyncVectorIndexSpecResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeltaSyncVectorIndexSpecResponse_SdkV2) {
}

func (c DeltaSyncVectorIndexSpecResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeltaSyncVectorIndexSpecResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"embedding_source_columns": reflect.TypeOf(EmbeddingSourceColumn_SdkV2{}),
		"embedding_vector_columns": reflect.TypeOf(EmbeddingVectorColumn_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaSyncVectorIndexSpecResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeltaSyncVectorIndexSpecResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding_source_columns":  o.EmbeddingSourceColumns,
			"embedding_vector_columns":  o.EmbeddingVectorColumns,
			"embedding_writeback_table": o.EmbeddingWritebackTable,
			"pipeline_id":               o.PipelineId,
			"pipeline_type":             o.PipelineType,
			"source_table":              o.SourceTable,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeltaSyncVectorIndexSpecResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
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
func (o *DeltaSyncVectorIndexSpecResponse_SdkV2) GetEmbeddingSourceColumns(ctx context.Context) ([]EmbeddingSourceColumn_SdkV2, bool) {
	if o.EmbeddingSourceColumns.IsNull() || o.EmbeddingSourceColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingSourceColumn_SdkV2
	d := o.EmbeddingSourceColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingSourceColumns sets the value of the EmbeddingSourceColumns field in DeltaSyncVectorIndexSpecResponse_SdkV2.
func (o *DeltaSyncVectorIndexSpecResponse_SdkV2) SetEmbeddingSourceColumns(ctx context.Context, v []EmbeddingSourceColumn_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_source_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EmbeddingSourceColumns = types.ListValueMust(t, vs)
}

// GetEmbeddingVectorColumns returns the value of the EmbeddingVectorColumns field in DeltaSyncVectorIndexSpecResponse_SdkV2 as
// a slice of EmbeddingVectorColumn_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *DeltaSyncVectorIndexSpecResponse_SdkV2) GetEmbeddingVectorColumns(ctx context.Context) ([]EmbeddingVectorColumn_SdkV2, bool) {
	if o.EmbeddingVectorColumns.IsNull() || o.EmbeddingVectorColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingVectorColumn_SdkV2
	d := o.EmbeddingVectorColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingVectorColumns sets the value of the EmbeddingVectorColumns field in DeltaSyncVectorIndexSpecResponse_SdkV2.
func (o *DeltaSyncVectorIndexSpecResponse_SdkV2) SetEmbeddingVectorColumns(ctx context.Context, v []EmbeddingVectorColumn_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_vector_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EmbeddingVectorColumns = types.ListValueMust(t, vs)
}

type DirectAccessVectorIndexSpec_SdkV2 struct {
	// Contains the optional model endpoint to use during query time.
	EmbeddingSourceColumns types.List `tfsdk:"embedding_source_columns"`

	EmbeddingVectorColumns types.List `tfsdk:"embedding_vector_columns"`
	// The schema of the index in JSON format.
	//
	// Supported types are `integer`, `long`, `float`, `double`, `boolean`,
	// `string`, `date`, `timestamp`.
	//
	// Supported types for vector column: `array<float>`, `array<double>`,`.
	SchemaJson types.String `tfsdk:"schema_json"`
}

func (newState *DirectAccessVectorIndexSpec_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DirectAccessVectorIndexSpec_SdkV2) {
}

func (newState *DirectAccessVectorIndexSpec_SdkV2) SyncEffectiveFieldsDuringRead(existingState DirectAccessVectorIndexSpec_SdkV2) {
}

func (c DirectAccessVectorIndexSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DirectAccessVectorIndexSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"embedding_source_columns": reflect.TypeOf(EmbeddingSourceColumn_SdkV2{}),
		"embedding_vector_columns": reflect.TypeOf(EmbeddingVectorColumn_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DirectAccessVectorIndexSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o DirectAccessVectorIndexSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding_source_columns": o.EmbeddingSourceColumns,
			"embedding_vector_columns": o.EmbeddingVectorColumns,
			"schema_json":              o.SchemaJson,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DirectAccessVectorIndexSpec_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *DirectAccessVectorIndexSpec_SdkV2) GetEmbeddingSourceColumns(ctx context.Context) ([]EmbeddingSourceColumn_SdkV2, bool) {
	if o.EmbeddingSourceColumns.IsNull() || o.EmbeddingSourceColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingSourceColumn_SdkV2
	d := o.EmbeddingSourceColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingSourceColumns sets the value of the EmbeddingSourceColumns field in DirectAccessVectorIndexSpec_SdkV2.
func (o *DirectAccessVectorIndexSpec_SdkV2) SetEmbeddingSourceColumns(ctx context.Context, v []EmbeddingSourceColumn_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_source_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EmbeddingSourceColumns = types.ListValueMust(t, vs)
}

// GetEmbeddingVectorColumns returns the value of the EmbeddingVectorColumns field in DirectAccessVectorIndexSpec_SdkV2 as
// a slice of EmbeddingVectorColumn_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *DirectAccessVectorIndexSpec_SdkV2) GetEmbeddingVectorColumns(ctx context.Context) ([]EmbeddingVectorColumn_SdkV2, bool) {
	if o.EmbeddingVectorColumns.IsNull() || o.EmbeddingVectorColumns.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingVectorColumn_SdkV2
	d := o.EmbeddingVectorColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbeddingVectorColumns sets the value of the EmbeddingVectorColumns field in DirectAccessVectorIndexSpec_SdkV2.
func (o *DirectAccessVectorIndexSpec_SdkV2) SetEmbeddingVectorColumns(ctx context.Context, v []EmbeddingVectorColumn_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding_vector_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EmbeddingVectorColumns = types.ListValueMust(t, vs)
}

type EmbeddingSourceColumn_SdkV2 struct {
	// Name of the embedding model endpoint
	EmbeddingModelEndpointName types.String `tfsdk:"embedding_model_endpoint_name"`
	// Name of the column
	Name types.String `tfsdk:"name"`
}

func (newState *EmbeddingSourceColumn_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EmbeddingSourceColumn_SdkV2) {
}

func (newState *EmbeddingSourceColumn_SdkV2) SyncEffectiveFieldsDuringRead(existingState EmbeddingSourceColumn_SdkV2) {
}

func (c EmbeddingSourceColumn_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["embedding_model_endpoint_name"] = attrs["embedding_model_endpoint_name"].SetOptional()
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
func (a EmbeddingSourceColumn_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EmbeddingSourceColumn_SdkV2
// only implements ToObjectValue() and Type().
func (o EmbeddingSourceColumn_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding_model_endpoint_name": o.EmbeddingModelEndpointName,
			"name":                          o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EmbeddingSourceColumn_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"embedding_model_endpoint_name": types.StringType,
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

func (newState *EmbeddingVectorColumn_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EmbeddingVectorColumn_SdkV2) {
}

func (newState *EmbeddingVectorColumn_SdkV2) SyncEffectiveFieldsDuringRead(existingState EmbeddingVectorColumn_SdkV2) {
}

func (c EmbeddingVectorColumn_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EmbeddingVectorColumn_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EmbeddingVectorColumn_SdkV2
// only implements ToObjectValue() and Type().
func (o EmbeddingVectorColumn_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding_dimension": o.EmbeddingDimension,
			"name":                o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EmbeddingVectorColumn_SdkV2) Type(ctx context.Context) attr.Type {
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
	// Current status of the endpoint
	EndpointStatus types.List `tfsdk:"endpoint_status"`
	// Type of endpoint.
	EndpointType types.String `tfsdk:"endpoint_type"`
	// Unique identifier of the endpoint
	Id types.String `tfsdk:"id"`
	// Timestamp of last update to the endpoint
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// User who last updated the endpoint
	LastUpdatedUser types.String `tfsdk:"last_updated_user"`
	// Name of endpoint
	Name types.String `tfsdk:"name"`
	// Number of indexes on the endpoint
	NumIndexes types.Int64 `tfsdk:"num_indexes"`
}

func (newState *EndpointInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointInfo_SdkV2) {
}

func (newState *EndpointInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState EndpointInfo_SdkV2) {
}

func (c EndpointInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creation_timestamp"] = attrs["creation_timestamp"].SetOptional()
	attrs["creator"] = attrs["creator"].SetOptional()
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
func (a EndpointInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoint_status": reflect.TypeOf(EndpointStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o EndpointInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_timestamp":     o.CreationTimestamp,
			"creator":                o.Creator,
			"endpoint_status":        o.EndpointStatus,
			"endpoint_type":          o.EndpointType,
			"id":                     o.Id,
			"last_updated_timestamp": o.LastUpdatedTimestamp,
			"last_updated_user":      o.LastUpdatedUser,
			"name":                   o.Name,
			"num_indexes":            o.NumIndexes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_timestamp": types.Int64Type,
			"creator":            types.StringType,
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

// GetEndpointStatus returns the value of the EndpointStatus field in EndpointInfo_SdkV2 as
// a EndpointStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointInfo_SdkV2) GetEndpointStatus(ctx context.Context) (EndpointStatus_SdkV2, bool) {
	var e EndpointStatus_SdkV2
	if o.EndpointStatus.IsNull() || o.EndpointStatus.IsUnknown() {
		return e, false
	}
	var v []EndpointStatus_SdkV2
	d := o.EndpointStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEndpointStatus sets the value of the EndpointStatus field in EndpointInfo_SdkV2.
func (o *EndpointInfo_SdkV2) SetEndpointStatus(ctx context.Context, v EndpointStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["endpoint_status"]
	o.EndpointStatus = types.ListValueMust(t, vs)
}

// Status information of an endpoint
type EndpointStatus_SdkV2 struct {
	// Additional status message
	Message types.String `tfsdk:"message"`
	// Current state of the endpoint
	State types.String `tfsdk:"state"`
}

func (newState *EndpointStatus_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointStatus_SdkV2) {
}

func (newState *EndpointStatus_SdkV2) SyncEffectiveFieldsDuringRead(existingState EndpointStatus_SdkV2) {
}

func (c EndpointStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EndpointStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o EndpointStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": o.Message,
			"state":   o.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"state":   types.StringType,
		},
	}
}

// Get an endpoint
type GetEndpointRequest_SdkV2 struct {
	// Name of the endpoint
	EndpointName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_name": o.EndpointName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_name": types.StringType,
		},
	}
}

// Get an index
type GetIndexRequest_SdkV2 struct {
	// Name of the index
	IndexName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name": o.IndexName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name": types.StringType,
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

func (newState *ListEndpointResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListEndpointResponse_SdkV2) {
}

func (newState *ListEndpointResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListEndpointResponse_SdkV2) {
}

func (c ListEndpointResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListEndpointResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoints": reflect.TypeOf(EndpointInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListEndpointResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoints":       o.Endpoints,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListEndpointResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListEndpointResponse_SdkV2) GetEndpoints(ctx context.Context) ([]EndpointInfo_SdkV2, bool) {
	if o.Endpoints.IsNull() || o.Endpoints.IsUnknown() {
		return nil, false
	}
	var v []EndpointInfo_SdkV2
	d := o.Endpoints.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEndpoints sets the value of the Endpoints field in ListEndpointResponse_SdkV2.
func (o *ListEndpointResponse_SdkV2) SetEndpoints(ctx context.Context, v []EndpointInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["endpoints"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Endpoints = types.ListValueMust(t, vs)
}

// List all endpoints
type ListEndpointsRequest_SdkV2 struct {
	// Token for pagination
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListEndpointsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListEndpointsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListEndpointsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListEndpointsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_token": types.StringType,
		},
	}
}

// List indexes
type ListIndexesRequest_SdkV2 struct {
	// Name of the endpoint
	EndpointName types.String `tfsdk:"-"`
	// Token for pagination
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListIndexesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListIndexesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListIndexesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListIndexesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_name": o.EndpointName,
			"page_token":    o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListIndexesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_name": types.StringType,
			"page_token":    types.StringType,
		},
	}
}

type ListValue_SdkV2 struct {
	Values types.List `tfsdk:"values"`
}

func (newState *ListValue_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListValue_SdkV2) {
}

func (newState *ListValue_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListValue_SdkV2) {
}

func (c ListValue_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListValue_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"values": reflect.TypeOf(Value_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListValue_SdkV2
// only implements ToObjectValue() and Type().
func (o ListValue_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"values": o.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListValue_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListValue_SdkV2) GetValues(ctx context.Context) ([]Value_SdkV2, bool) {
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

// SetValues sets the value of the Values field in ListValue_SdkV2.
func (o *ListValue_SdkV2) SetValues(ctx context.Context, v []Value_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Values = types.ListValueMust(t, vs)
}

type ListVectorIndexesResponse_SdkV2 struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`

	VectorIndexes types.List `tfsdk:"vector_indexes"`
}

func (newState *ListVectorIndexesResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListVectorIndexesResponse_SdkV2) {
}

func (newState *ListVectorIndexesResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListVectorIndexesResponse_SdkV2) {
}

func (c ListVectorIndexesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListVectorIndexesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"vector_indexes": reflect.TypeOf(MiniVectorIndex_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListVectorIndexesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListVectorIndexesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"vector_indexes":  o.VectorIndexes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListVectorIndexesResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListVectorIndexesResponse_SdkV2) GetVectorIndexes(ctx context.Context) ([]MiniVectorIndex_SdkV2, bool) {
	if o.VectorIndexes.IsNull() || o.VectorIndexes.IsUnknown() {
		return nil, false
	}
	var v []MiniVectorIndex_SdkV2
	d := o.VectorIndexes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVectorIndexes sets the value of the VectorIndexes field in ListVectorIndexesResponse_SdkV2.
func (o *ListVectorIndexesResponse_SdkV2) SetVectorIndexes(ctx context.Context, v []MiniVectorIndex_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["vector_indexes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.VectorIndexes = types.ListValueMust(t, vs)
}

// Key-value pair.
type MapStringValueEntry_SdkV2 struct {
	// Column name.
	Key types.String `tfsdk:"key"`
	// Column value, nullable.
	Value types.List `tfsdk:"value"`
}

func (newState *MapStringValueEntry_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MapStringValueEntry_SdkV2) {
}

func (newState *MapStringValueEntry_SdkV2) SyncEffectiveFieldsDuringRead(existingState MapStringValueEntry_SdkV2) {
}

func (c MapStringValueEntry_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a MapStringValueEntry_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"value": reflect.TypeOf(Value_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MapStringValueEntry_SdkV2
// only implements ToObjectValue() and Type().
func (o MapStringValueEntry_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MapStringValueEntry_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *MapStringValueEntry_SdkV2) GetValue(ctx context.Context) (Value_SdkV2, bool) {
	var e Value_SdkV2
	if o.Value.IsNull() || o.Value.IsUnknown() {
		return e, false
	}
	var v []Value_SdkV2
	d := o.Value.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetValue sets the value of the Value field in MapStringValueEntry_SdkV2.
func (o *MapStringValueEntry_SdkV2) SetValue(ctx context.Context, v Value_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["value"]
	o.Value = types.ListValueMust(t, vs)
}

type MiniVectorIndex_SdkV2 struct {
	// The user who created the index.
	Creator types.String `tfsdk:"creator"`
	// Name of the endpoint associated with the index
	EndpointName types.String `tfsdk:"endpoint_name"`
	// There are 2 types of Vector Search indexes:
	//
	// - `DELTA_SYNC`: An index that automatically syncs with a source Delta
	// Table, automatically and incrementally updating the index as the
	// underlying data in the Delta Table changes. - `DIRECT_ACCESS`: An index
	// that supports direct read and write of vectors and metadata through our
	// REST and SDK APIs. With this model, the user manages index updates.
	IndexType types.String `tfsdk:"index_type"`
	// Name of the index
	Name types.String `tfsdk:"name"`
	// Primary key of the index
	PrimaryKey types.String `tfsdk:"primary_key"`
}

func (newState *MiniVectorIndex_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MiniVectorIndex_SdkV2) {
}

func (newState *MiniVectorIndex_SdkV2) SyncEffectiveFieldsDuringRead(existingState MiniVectorIndex_SdkV2) {
}

func (c MiniVectorIndex_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a MiniVectorIndex_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MiniVectorIndex_SdkV2
// only implements ToObjectValue() and Type().
func (o MiniVectorIndex_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creator":       o.Creator,
			"endpoint_name": o.EndpointName,
			"index_type":    o.IndexType,
			"name":          o.Name,
			"primary_key":   o.PrimaryKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MiniVectorIndex_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *QueryVectorIndexNextPageRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryVectorIndexNextPageRequest_SdkV2) {
}

func (newState *QueryVectorIndexNextPageRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState QueryVectorIndexNextPageRequest_SdkV2) {
}

func (c QueryVectorIndexNextPageRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint_name"] = attrs["endpoint_name"].SetOptional()
	attrs["index_name"] = attrs["index_name"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryVectorIndexNextPageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryVectorIndexNextPageRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryVectorIndexNextPageRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o QueryVectorIndexNextPageRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_name": o.EndpointName,
			"index_name":    o.IndexName,
			"page_token":    o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryVectorIndexNextPageRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
	// JSON string representing query filters.
	//
	// Example filters: - `{"id <": 5}`: Filter for id less than 5. - `{"id >":
	// 5}`: Filter for id greater than 5. - `{"id <=": 5}`: Filter for id less
	// than equal to 5. - `{"id >=": 5}`: Filter for id greater than equal to 5.
	// - `{"id": 5}`: Filter for id equal to 5.
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
	// Threshold for the approximate nearest neighbor search. Defaults to 0.0.
	ScoreThreshold types.Float64 `tfsdk:"score_threshold"`
}

func (newState *QueryVectorIndexRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryVectorIndexRequest_SdkV2) {
}

func (newState *QueryVectorIndexRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState QueryVectorIndexRequest_SdkV2) {
}

func (c QueryVectorIndexRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["columns"] = attrs["columns"].SetRequired()
	attrs["filters_json"] = attrs["filters_json"].SetOptional()
	attrs["index_name"] = attrs["index_name"].SetRequired()
	attrs["num_results"] = attrs["num_results"].SetOptional()
	attrs["query_text"] = attrs["query_text"].SetOptional()
	attrs["query_type"] = attrs["query_type"].SetOptional()
	attrs["query_vector"] = attrs["query_vector"].SetOptional()
	attrs["score_threshold"] = attrs["score_threshold"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryVectorIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryVectorIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns":      reflect.TypeOf(types.String{}),
		"query_vector": reflect.TypeOf(types.Float64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryVectorIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o QueryVectorIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns":         o.Columns,
			"filters_json":    o.FiltersJson,
			"index_name":      o.IndexName,
			"num_results":     o.NumResults,
			"query_text":      o.QueryText,
			"query_type":      o.QueryType,
			"query_vector":    o.QueryVector,
			"score_threshold": o.ScoreThreshold,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryVectorIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns": basetypes.ListType{
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
			"score_threshold": types.Float64Type,
		},
	}
}

// GetColumns returns the value of the Columns field in QueryVectorIndexRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryVectorIndexRequest_SdkV2) GetColumns(ctx context.Context) ([]types.String, bool) {
	if o.Columns.IsNull() || o.Columns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in QueryVectorIndexRequest_SdkV2.
func (o *QueryVectorIndexRequest_SdkV2) SetColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Columns = types.ListValueMust(t, vs)
}

// GetQueryVector returns the value of the QueryVector field in QueryVectorIndexRequest_SdkV2 as
// a slice of types.Float64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryVectorIndexRequest_SdkV2) GetQueryVector(ctx context.Context) ([]types.Float64, bool) {
	if o.QueryVector.IsNull() || o.QueryVector.IsUnknown() {
		return nil, false
	}
	var v []types.Float64
	d := o.QueryVector.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQueryVector sets the value of the QueryVector field in QueryVectorIndexRequest_SdkV2.
func (o *QueryVectorIndexRequest_SdkV2) SetQueryVector(ctx context.Context, v []types.Float64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query_vector"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.QueryVector = types.ListValueMust(t, vs)
}

type QueryVectorIndexResponse_SdkV2 struct {
	// Metadata about the result set.
	Manifest types.List `tfsdk:"manifest"`
	// [Optional] Token that can be used in `QueryVectorIndexNextPage` API to
	// get next page of results. If more than 1000 results satisfy the query,
	// they are returned in groups of 1000. Empty value means no more results.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Data returned in the query result.
	Result types.List `tfsdk:"result"`
}

func (newState *QueryVectorIndexResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryVectorIndexResponse_SdkV2) {
}

func (newState *QueryVectorIndexResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState QueryVectorIndexResponse_SdkV2) {
}

func (c QueryVectorIndexResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a QueryVectorIndexResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"manifest": reflect.TypeOf(ResultManifest_SdkV2{}),
		"result":   reflect.TypeOf(ResultData_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryVectorIndexResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o QueryVectorIndexResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"manifest":        o.Manifest,
			"next_page_token": o.NextPageToken,
			"result":          o.Result,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryVectorIndexResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *QueryVectorIndexResponse_SdkV2) GetManifest(ctx context.Context) (ResultManifest_SdkV2, bool) {
	var e ResultManifest_SdkV2
	if o.Manifest.IsNull() || o.Manifest.IsUnknown() {
		return e, false
	}
	var v []ResultManifest_SdkV2
	d := o.Manifest.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetManifest sets the value of the Manifest field in QueryVectorIndexResponse_SdkV2.
func (o *QueryVectorIndexResponse_SdkV2) SetManifest(ctx context.Context, v ResultManifest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["manifest"]
	o.Manifest = types.ListValueMust(t, vs)
}

// GetResult returns the value of the Result field in QueryVectorIndexResponse_SdkV2 as
// a ResultData_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryVectorIndexResponse_SdkV2) GetResult(ctx context.Context) (ResultData_SdkV2, bool) {
	var e ResultData_SdkV2
	if o.Result.IsNull() || o.Result.IsUnknown() {
		return e, false
	}
	var v []ResultData_SdkV2
	d := o.Result.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetResult sets the value of the Result field in QueryVectorIndexResponse_SdkV2.
func (o *QueryVectorIndexResponse_SdkV2) SetResult(ctx context.Context, v ResultData_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["result"]
	o.Result = types.ListValueMust(t, vs)
}

// Data returned in the query result.
type ResultData_SdkV2 struct {
	// Data rows returned in the query.
	DataArray types.List `tfsdk:"data_array"`
	// Number of rows in the result set.
	RowCount types.Int64 `tfsdk:"row_count"`
}

func (newState *ResultData_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResultData_SdkV2) {
}

func (newState *ResultData_SdkV2) SyncEffectiveFieldsDuringRead(existingState ResultData_SdkV2) {
}

func (c ResultData_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ResultData_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_array": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultData_SdkV2
// only implements ToObjectValue() and Type().
func (o ResultData_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_array": o.DataArray,
			"row_count":  o.RowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResultData_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ResultData_SdkV2) GetDataArray(ctx context.Context) ([]types.String, bool) {
	if o.DataArray.IsNull() || o.DataArray.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.DataArray.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataArray sets the value of the DataArray field in ResultData_SdkV2.
func (o *ResultData_SdkV2) SetDataArray(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data_array"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DataArray = types.ListValueMust(t, vs)
}

// Metadata about the result set.
type ResultManifest_SdkV2 struct {
	// Number of columns in the result set.
	ColumnCount types.Int64 `tfsdk:"column_count"`
	// Information about each column in the result set.
	Columns types.List `tfsdk:"columns"`
}

func (newState *ResultManifest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResultManifest_SdkV2) {
}

func (newState *ResultManifest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ResultManifest_SdkV2) {
}

func (c ResultManifest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ResultManifest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(ColumnInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultManifest_SdkV2
// only implements ToObjectValue() and Type().
func (o ResultManifest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column_count": o.ColumnCount,
			"columns":      o.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ResultManifest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ResultManifest_SdkV2) GetColumns(ctx context.Context) ([]ColumnInfo_SdkV2, bool) {
	if o.Columns.IsNull() || o.Columns.IsUnknown() {
		return nil, false
	}
	var v []ColumnInfo_SdkV2
	d := o.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in ResultManifest_SdkV2.
func (o *ResultManifest_SdkV2) SetColumns(ctx context.Context, v []ColumnInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Columns = types.ListValueMust(t, vs)
}

// Request payload for scanning data from a vector index.
type ScanVectorIndexRequest_SdkV2 struct {
	// Name of the vector index to scan.
	IndexName types.String `tfsdk:"-"`
	// Primary key of the last entry returned in the previous scan.
	LastPrimaryKey types.String `tfsdk:"last_primary_key"`
	// Number of results to return. Defaults to 10.
	NumResults types.Int64 `tfsdk:"num_results"`
}

func (newState *ScanVectorIndexRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ScanVectorIndexRequest_SdkV2) {
}

func (newState *ScanVectorIndexRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ScanVectorIndexRequest_SdkV2) {
}

func (c ScanVectorIndexRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["index_name"] = attrs["index_name"].SetRequired()
	attrs["last_primary_key"] = attrs["last_primary_key"].SetOptional()
	attrs["num_results"] = attrs["num_results"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ScanVectorIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ScanVectorIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ScanVectorIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ScanVectorIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name":       o.IndexName,
			"last_primary_key": o.LastPrimaryKey,
			"num_results":      o.NumResults,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ScanVectorIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *ScanVectorIndexResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ScanVectorIndexResponse_SdkV2) {
}

func (newState *ScanVectorIndexResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ScanVectorIndexResponse_SdkV2) {
}

func (c ScanVectorIndexResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ScanVectorIndexResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data": reflect.TypeOf(Struct_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ScanVectorIndexResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ScanVectorIndexResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data":             o.Data,
			"last_primary_key": o.LastPrimaryKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ScanVectorIndexResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ScanVectorIndexResponse_SdkV2) GetData(ctx context.Context) ([]Struct_SdkV2, bool) {
	if o.Data.IsNull() || o.Data.IsUnknown() {
		return nil, false
	}
	var v []Struct_SdkV2
	d := o.Data.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetData sets the value of the Data field in ScanVectorIndexResponse_SdkV2.
func (o *ScanVectorIndexResponse_SdkV2) SetData(ctx context.Context, v []Struct_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Data = types.ListValueMust(t, vs)
}

type Struct_SdkV2 struct {
	// Data entry, corresponding to a row in a vector index.
	Fields types.List `tfsdk:"fields"`
}

func (newState *Struct_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Struct_SdkV2) {
}

func (newState *Struct_SdkV2) SyncEffectiveFieldsDuringRead(existingState Struct_SdkV2) {
}

func (c Struct_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Struct_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"fields": reflect.TypeOf(MapStringValueEntry_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Struct_SdkV2
// only implements ToObjectValue() and Type().
func (o Struct_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"fields": o.Fields,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Struct_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *Struct_SdkV2) GetFields(ctx context.Context) ([]MapStringValueEntry_SdkV2, bool) {
	if o.Fields.IsNull() || o.Fields.IsUnknown() {
		return nil, false
	}
	var v []MapStringValueEntry_SdkV2
	d := o.Fields.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFields sets the value of the Fields field in Struct_SdkV2.
func (o *Struct_SdkV2) SetFields(ctx context.Context, v []MapStringValueEntry_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["fields"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Fields = types.ListValueMust(t, vs)
}

// Synchronize an index
type SyncIndexRequest_SdkV2 struct {
	// Name of the vector index to synchronize. Must be a Delta Sync Index.
	IndexName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SyncIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o SyncIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name": o.IndexName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SyncIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name": types.StringType,
		},
	}
}

type SyncIndexResponse_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncIndexResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SyncIndexResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncIndexResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o SyncIndexResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o SyncIndexResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Result of the upsert or delete operation.
type UpsertDataResult_SdkV2 struct {
	// List of primary keys for rows that failed to process.
	FailedPrimaryKeys types.List `tfsdk:"failed_primary_keys"`
	// Count of successfully processed rows.
	SuccessRowCount types.Int64 `tfsdk:"success_row_count"`
}

func (newState *UpsertDataResult_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpsertDataResult_SdkV2) {
}

func (newState *UpsertDataResult_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpsertDataResult_SdkV2) {
}

func (c UpsertDataResult_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpsertDataResult_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"failed_primary_keys": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpsertDataResult_SdkV2
// only implements ToObjectValue() and Type().
func (o UpsertDataResult_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"failed_primary_keys": o.FailedPrimaryKeys,
			"success_row_count":   o.SuccessRowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpsertDataResult_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpsertDataResult_SdkV2) GetFailedPrimaryKeys(ctx context.Context) ([]types.String, bool) {
	if o.FailedPrimaryKeys.IsNull() || o.FailedPrimaryKeys.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.FailedPrimaryKeys.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFailedPrimaryKeys sets the value of the FailedPrimaryKeys field in UpsertDataResult_SdkV2.
func (o *UpsertDataResult_SdkV2) SetFailedPrimaryKeys(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["failed_primary_keys"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.FailedPrimaryKeys = types.ListValueMust(t, vs)
}

// Request payload for upserting data into a vector index.
type UpsertDataVectorIndexRequest_SdkV2 struct {
	// Name of the vector index where data is to be upserted. Must be a Direct
	// Vector Access Index.
	IndexName types.String `tfsdk:"-"`
	// JSON string representing the data to be upserted.
	InputsJson types.String `tfsdk:"inputs_json"`
}

func (newState *UpsertDataVectorIndexRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpsertDataVectorIndexRequest_SdkV2) {
}

func (newState *UpsertDataVectorIndexRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpsertDataVectorIndexRequest_SdkV2) {
}

func (c UpsertDataVectorIndexRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["index_name"] = attrs["index_name"].SetRequired()
	attrs["inputs_json"] = attrs["inputs_json"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpsertDataVectorIndexRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpsertDataVectorIndexRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpsertDataVectorIndexRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpsertDataVectorIndexRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_name":  o.IndexName,
			"inputs_json": o.InputsJson,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpsertDataVectorIndexRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_name":  types.StringType,
			"inputs_json": types.StringType,
		},
	}
}

// Response to an upsert data vector index request.
type UpsertDataVectorIndexResponse_SdkV2 struct {
	// Result of the upsert or delete operation.
	Result types.List `tfsdk:"result"`
	// Status of the upsert operation.
	Status types.String `tfsdk:"status"`
}

func (newState *UpsertDataVectorIndexResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpsertDataVectorIndexResponse_SdkV2) {
}

func (newState *UpsertDataVectorIndexResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpsertDataVectorIndexResponse_SdkV2) {
}

func (c UpsertDataVectorIndexResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpsertDataVectorIndexResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"result": reflect.TypeOf(UpsertDataResult_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpsertDataVectorIndexResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpsertDataVectorIndexResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"result": o.Result,
			"status": o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpsertDataVectorIndexResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpsertDataVectorIndexResponse_SdkV2) GetResult(ctx context.Context) (UpsertDataResult_SdkV2, bool) {
	var e UpsertDataResult_SdkV2
	if o.Result.IsNull() || o.Result.IsUnknown() {
		return e, false
	}
	var v []UpsertDataResult_SdkV2
	d := o.Result.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetResult sets the value of the Result field in UpsertDataVectorIndexResponse_SdkV2.
func (o *UpsertDataVectorIndexResponse_SdkV2) SetResult(ctx context.Context, v UpsertDataResult_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["result"]
	o.Result = types.ListValueMust(t, vs)
}

type Value_SdkV2 struct {
	BoolValue types.Bool `tfsdk:"bool_value"`

	ListValue types.List `tfsdk:"list_value"`

	NullValue types.String `tfsdk:"null_value"`

	NumberValue types.Float64 `tfsdk:"number_value"`

	StringValue types.String `tfsdk:"string_value"`

	StructValue types.List `tfsdk:"struct_value"`
}

func (newState *Value_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Value_SdkV2) {
}

func (newState *Value_SdkV2) SyncEffectiveFieldsDuringRead(existingState Value_SdkV2) {
}

func (c Value_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["bool_value"] = attrs["bool_value"].SetOptional()
	attrs["list_value"] = attrs["list_value"].SetOptional()
	attrs["list_value"] = attrs["list_value"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["null_value"] = attrs["null_value"].SetOptional()
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
func (a Value_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"list_value":   reflect.TypeOf(ListValue_SdkV2{}),
		"struct_value": reflect.TypeOf(Struct_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Value_SdkV2
// only implements ToObjectValue() and Type().
func (o Value_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bool_value":   o.BoolValue,
			"list_value":   o.ListValue,
			"null_value":   o.NullValue,
			"number_value": o.NumberValue,
			"string_value": o.StringValue,
			"struct_value": o.StructValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Value_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bool_value": types.BoolType,
			"list_value": basetypes.ListType{
				ElemType: ListValue_SdkV2{}.Type(ctx),
			},
			"null_value":   types.StringType,
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
func (o *Value_SdkV2) GetListValue(ctx context.Context) (ListValue_SdkV2, bool) {
	var e ListValue_SdkV2
	if o.ListValue.IsNull() || o.ListValue.IsUnknown() {
		return e, false
	}
	var v []ListValue_SdkV2
	d := o.ListValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetListValue sets the value of the ListValue field in Value_SdkV2.
func (o *Value_SdkV2) SetListValue(ctx context.Context, v ListValue_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["list_value"]
	o.ListValue = types.ListValueMust(t, vs)
}

// GetStructValue returns the value of the StructValue field in Value_SdkV2 as
// a Struct_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Value_SdkV2) GetStructValue(ctx context.Context) (Struct_SdkV2, bool) {
	var e Struct_SdkV2
	if o.StructValue.IsNull() || o.StructValue.IsUnknown() {
		return e, false
	}
	var v []Struct_SdkV2
	d := o.StructValue.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStructValue sets the value of the StructValue field in Value_SdkV2.
func (o *Value_SdkV2) SetStructValue(ctx context.Context, v Struct_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["struct_value"]
	o.StructValue = types.ListValueMust(t, vs)
}

type VectorIndex_SdkV2 struct {
	// The user who created the index.
	Creator types.String `tfsdk:"creator"`

	DeltaSyncIndexSpec types.List `tfsdk:"delta_sync_index_spec"`

	DirectAccessIndexSpec types.List `tfsdk:"direct_access_index_spec"`
	// Name of the endpoint associated with the index
	EndpointName types.String `tfsdk:"endpoint_name"`
	// There are 2 types of Vector Search indexes:
	//
	// - `DELTA_SYNC`: An index that automatically syncs with a source Delta
	// Table, automatically and incrementally updating the index as the
	// underlying data in the Delta Table changes. - `DIRECT_ACCESS`: An index
	// that supports direct read and write of vectors and metadata through our
	// REST and SDK APIs. With this model, the user manages index updates.
	IndexType types.String `tfsdk:"index_type"`
	// Name of the index
	Name types.String `tfsdk:"name"`
	// Primary key of the index
	PrimaryKey types.String `tfsdk:"primary_key"`

	Status types.List `tfsdk:"status"`
}

func (newState *VectorIndex_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan VectorIndex_SdkV2) {
}

func (newState *VectorIndex_SdkV2) SyncEffectiveFieldsDuringRead(existingState VectorIndex_SdkV2) {
}

func (c VectorIndex_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a VectorIndex_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"delta_sync_index_spec":    reflect.TypeOf(DeltaSyncVectorIndexSpecResponse_SdkV2{}),
		"direct_access_index_spec": reflect.TypeOf(DirectAccessVectorIndexSpec_SdkV2{}),
		"status":                   reflect.TypeOf(VectorIndexStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VectorIndex_SdkV2
// only implements ToObjectValue() and Type().
func (o VectorIndex_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creator":                  o.Creator,
			"delta_sync_index_spec":    o.DeltaSyncIndexSpec,
			"direct_access_index_spec": o.DirectAccessIndexSpec,
			"endpoint_name":            o.EndpointName,
			"index_type":               o.IndexType,
			"name":                     o.Name,
			"primary_key":              o.PrimaryKey,
			"status":                   o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o VectorIndex_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *VectorIndex_SdkV2) GetDeltaSyncIndexSpec(ctx context.Context) (DeltaSyncVectorIndexSpecResponse_SdkV2, bool) {
	var e DeltaSyncVectorIndexSpecResponse_SdkV2
	if o.DeltaSyncIndexSpec.IsNull() || o.DeltaSyncIndexSpec.IsUnknown() {
		return e, false
	}
	var v []DeltaSyncVectorIndexSpecResponse_SdkV2
	d := o.DeltaSyncIndexSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDeltaSyncIndexSpec sets the value of the DeltaSyncIndexSpec field in VectorIndex_SdkV2.
func (o *VectorIndex_SdkV2) SetDeltaSyncIndexSpec(ctx context.Context, v DeltaSyncVectorIndexSpecResponse_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["delta_sync_index_spec"]
	o.DeltaSyncIndexSpec = types.ListValueMust(t, vs)
}

// GetDirectAccessIndexSpec returns the value of the DirectAccessIndexSpec field in VectorIndex_SdkV2 as
// a DirectAccessVectorIndexSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *VectorIndex_SdkV2) GetDirectAccessIndexSpec(ctx context.Context) (DirectAccessVectorIndexSpec_SdkV2, bool) {
	var e DirectAccessVectorIndexSpec_SdkV2
	if o.DirectAccessIndexSpec.IsNull() || o.DirectAccessIndexSpec.IsUnknown() {
		return e, false
	}
	var v []DirectAccessVectorIndexSpec_SdkV2
	d := o.DirectAccessIndexSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDirectAccessIndexSpec sets the value of the DirectAccessIndexSpec field in VectorIndex_SdkV2.
func (o *VectorIndex_SdkV2) SetDirectAccessIndexSpec(ctx context.Context, v DirectAccessVectorIndexSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["direct_access_index_spec"]
	o.DirectAccessIndexSpec = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in VectorIndex_SdkV2 as
// a VectorIndexStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *VectorIndex_SdkV2) GetStatus(ctx context.Context) (VectorIndexStatus_SdkV2, bool) {
	var e VectorIndexStatus_SdkV2
	if o.Status.IsNull() || o.Status.IsUnknown() {
		return e, false
	}
	var v []VectorIndexStatus_SdkV2
	d := o.Status.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStatus sets the value of the Status field in VectorIndex_SdkV2.
func (o *VectorIndex_SdkV2) SetStatus(ctx context.Context, v VectorIndexStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["status"]
	o.Status = types.ListValueMust(t, vs)
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

func (newState *VectorIndexStatus_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan VectorIndexStatus_SdkV2) {
}

func (newState *VectorIndexStatus_SdkV2) SyncEffectiveFieldsDuringRead(existingState VectorIndexStatus_SdkV2) {
}

func (c VectorIndexStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a VectorIndexStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VectorIndexStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o VectorIndexStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"index_url":         o.IndexUrl,
			"indexed_row_count": o.IndexedRowCount,
			"message":           o.Message,
			"ready":             o.Ready,
		})
}

// Type implements basetypes.ObjectValuable.
func (o VectorIndexStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"index_url":         types.StringType,
			"indexed_row_count": types.Int64Type,
			"message":           types.StringType,
			"ready":             types.BoolType,
		},
	}
}
