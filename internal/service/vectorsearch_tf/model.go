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
	// Data type of the column (e.g., "string", "int", "array<float>")
	TypeText types.String `tfsdk:"type_text"`
}

func (to *ColumnInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ColumnInfo) {
}

func (to *ColumnInfo) SyncFieldsDuringRead(ctx context.Context, from ColumnInfo) {
}

func (m ColumnInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetOptional()
	attrs["type_text"] = attrs["type_text"].SetOptional()

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
			"name":      m.Name,
			"type_text": m.TypeText,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ColumnInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":      types.StringType,
			"type_text": types.StringType,
		},
	}
}

type CreateEndpoint struct {
	// The budget policy id to be applied
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// Type of endpoint
	EndpointType types.String `tfsdk:"endpoint_type"`
	// Name of the AI Search endpoint
	Name types.String `tfsdk:"name"`
	// Target QPS for the endpoint. Mutually exclusive with num_replicas. The
	// actual replica count is calculated at index creation/sync time based on
	// this value. Best-effort target; the system does not guarantee this QPS
	// will be achieved.
	TargetQps types.Int64 `tfsdk:"target_qps"`
	// The usage policy id to be applied once we've migrated to usage policies
	UsagePolicyId types.String `tfsdk:"usage_policy_id"`
}

func (to *CreateEndpoint) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateEndpoint) {
}

func (to *CreateEndpoint) SyncFieldsDuringRead(ctx context.Context, from CreateEndpoint) {
}

func (m CreateEndpoint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["endpoint_type"] = attrs["endpoint_type"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["target_qps"] = attrs["target_qps"].SetOptional()
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
			"target_qps":       m.TargetQps,
			"usage_policy_id":  m.UsagePolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateEndpoint) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_policy_id": types.StringType,
			"endpoint_type":    types.StringType,
			"name":             types.StringType,
			"target_qps":       types.Int64Type,
			"usage_policy_id":  types.StringType,
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
	// The subtype of the index. Use `HYBRID` or `FULL_TEXT`. `VECTOR` is not
	// supported.
	IndexSubtype types.String `tfsdk:"index_subtype"`

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
	attrs["index_subtype"] = attrs["index_subtype"].SetOptional()
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
			"index_subtype":            m.IndexSubtype,
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
			"index_subtype":            types.StringType,
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
	// Key field for an AI Search endpoint tag.
	Key types.String `tfsdk:"key"`
	// [Optional] Value field for an AI Search endpoint tag.
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
	// Name of the AI Search endpoint
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
	// [Optional] Alias for columns_to_sync. Select the columns to include in
	// the vector index. If you leave this field blank, all columns from the
	// source table are included. The primary key column and embedding source
	// column or embedding vector column are always included. Only one of
	// columns_to_sync or columns_to_index may be specified.
	ColumnsToIndex types.List `tfsdk:"columns_to_index"`
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
	if !from.ColumnsToIndex.IsNull() && !from.ColumnsToIndex.IsUnknown() && to.ColumnsToIndex.IsNull() && len(from.ColumnsToIndex.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnsToIndex, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnsToIndex = from.ColumnsToIndex
	}
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
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() {
		if toEmbeddingSourceColumns, ok := to.GetEmbeddingSourceColumns(ctx); ok {
			if fromEmbeddingSourceColumns, ok := from.GetEmbeddingSourceColumns(ctx); ok {
				// Recursively sync the fields of each EmbeddingSourceColumns element by position.
				for i := range toEmbeddingSourceColumns {
					if i < len(fromEmbeddingSourceColumns) {
						toEmbeddingSourceColumns[i].SyncFieldsDuringCreateOrUpdate(ctx, fromEmbeddingSourceColumns[i])
					}
				}
				to.SetEmbeddingSourceColumns(ctx, toEmbeddingSourceColumns)
			}
		}
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() && to.EmbeddingVectorColumns.IsNull() && len(from.EmbeddingVectorColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingVectorColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingVectorColumns = from.EmbeddingVectorColumns
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() {
		if toEmbeddingVectorColumns, ok := to.GetEmbeddingVectorColumns(ctx); ok {
			if fromEmbeddingVectorColumns, ok := from.GetEmbeddingVectorColumns(ctx); ok {
				// Recursively sync the fields of each EmbeddingVectorColumns element by position.
				for i := range toEmbeddingVectorColumns {
					if i < len(fromEmbeddingVectorColumns) {
						toEmbeddingVectorColumns[i].SyncFieldsDuringCreateOrUpdate(ctx, fromEmbeddingVectorColumns[i])
					}
				}
				to.SetEmbeddingVectorColumns(ctx, toEmbeddingVectorColumns)
			}
		}
	}
}

func (to *DeltaSyncVectorIndexSpecRequest) SyncFieldsDuringRead(ctx context.Context, from DeltaSyncVectorIndexSpecRequest) {
	if !from.ColumnsToIndex.IsNull() && !from.ColumnsToIndex.IsUnknown() && to.ColumnsToIndex.IsNull() && len(from.ColumnsToIndex.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnsToIndex, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnsToIndex = from.ColumnsToIndex
	}
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
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() {
		if toEmbeddingSourceColumns, ok := to.GetEmbeddingSourceColumns(ctx); ok {
			if fromEmbeddingSourceColumns, ok := from.GetEmbeddingSourceColumns(ctx); ok {
				for i := range toEmbeddingSourceColumns {
					if i < len(fromEmbeddingSourceColumns) {
						toEmbeddingSourceColumns[i].SyncFieldsDuringRead(ctx, fromEmbeddingSourceColumns[i])
					}
				}
				to.SetEmbeddingSourceColumns(ctx, toEmbeddingSourceColumns)
			}
		}
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() && to.EmbeddingVectorColumns.IsNull() && len(from.EmbeddingVectorColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingVectorColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingVectorColumns = from.EmbeddingVectorColumns
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() {
		if toEmbeddingVectorColumns, ok := to.GetEmbeddingVectorColumns(ctx); ok {
			if fromEmbeddingVectorColumns, ok := from.GetEmbeddingVectorColumns(ctx); ok {
				for i := range toEmbeddingVectorColumns {
					if i < len(fromEmbeddingVectorColumns) {
						toEmbeddingVectorColumns[i].SyncFieldsDuringRead(ctx, fromEmbeddingVectorColumns[i])
					}
				}
				to.SetEmbeddingVectorColumns(ctx, toEmbeddingVectorColumns)
			}
		}
	}
}

func (m DeltaSyncVectorIndexSpecRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["columns_to_index"] = attrs["columns_to_index"].SetOptional()
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
		"columns_to_index":         reflect.TypeOf(types.String{}),
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
			"columns_to_index":          m.ColumnsToIndex,
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
			"columns_to_index": basetypes.ListType{
				ElemType: types.StringType,
			},
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

// GetColumnsToIndex returns the value of the ColumnsToIndex field in DeltaSyncVectorIndexSpecRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSyncVectorIndexSpecRequest) GetColumnsToIndex(ctx context.Context) ([]types.String, bool) {
	if m.ColumnsToIndex.IsNull() || m.ColumnsToIndex.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ColumnsToIndex.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumnsToIndex sets the value of the ColumnsToIndex field in DeltaSyncVectorIndexSpecRequest.
func (m *DeltaSyncVectorIndexSpecRequest) SetColumnsToIndex(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns_to_index"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ColumnsToIndex = types.ListValueMust(t, vs)
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
	// [Optional] Alias for columns_to_sync. Select the columns to include in
	// the vector index. If you leave this field blank, all columns from the
	// source table are included. The primary key column and embedding source
	// column or embedding vector column are always included. Only one of
	// columns_to_sync or columns_to_index may be specified.
	ColumnsToIndex types.List `tfsdk:"columns_to_index"`
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
	if !from.ColumnsToIndex.IsNull() && !from.ColumnsToIndex.IsUnknown() && to.ColumnsToIndex.IsNull() && len(from.ColumnsToIndex.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnsToIndex, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnsToIndex = from.ColumnsToIndex
	}
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
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() {
		if toEmbeddingSourceColumns, ok := to.GetEmbeddingSourceColumns(ctx); ok {
			if fromEmbeddingSourceColumns, ok := from.GetEmbeddingSourceColumns(ctx); ok {
				// Recursively sync the fields of each EmbeddingSourceColumns element by position.
				for i := range toEmbeddingSourceColumns {
					if i < len(fromEmbeddingSourceColumns) {
						toEmbeddingSourceColumns[i].SyncFieldsDuringCreateOrUpdate(ctx, fromEmbeddingSourceColumns[i])
					}
				}
				to.SetEmbeddingSourceColumns(ctx, toEmbeddingSourceColumns)
			}
		}
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() && to.EmbeddingVectorColumns.IsNull() && len(from.EmbeddingVectorColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingVectorColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingVectorColumns = from.EmbeddingVectorColumns
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() {
		if toEmbeddingVectorColumns, ok := to.GetEmbeddingVectorColumns(ctx); ok {
			if fromEmbeddingVectorColumns, ok := from.GetEmbeddingVectorColumns(ctx); ok {
				// Recursively sync the fields of each EmbeddingVectorColumns element by position.
				for i := range toEmbeddingVectorColumns {
					if i < len(fromEmbeddingVectorColumns) {
						toEmbeddingVectorColumns[i].SyncFieldsDuringCreateOrUpdate(ctx, fromEmbeddingVectorColumns[i])
					}
				}
				to.SetEmbeddingVectorColumns(ctx, toEmbeddingVectorColumns)
			}
		}
	}
}

func (to *DeltaSyncVectorIndexSpecResponse) SyncFieldsDuringRead(ctx context.Context, from DeltaSyncVectorIndexSpecResponse) {
	if !from.ColumnsToIndex.IsNull() && !from.ColumnsToIndex.IsUnknown() && to.ColumnsToIndex.IsNull() && len(from.ColumnsToIndex.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnsToIndex, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnsToIndex = from.ColumnsToIndex
	}
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
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() {
		if toEmbeddingSourceColumns, ok := to.GetEmbeddingSourceColumns(ctx); ok {
			if fromEmbeddingSourceColumns, ok := from.GetEmbeddingSourceColumns(ctx); ok {
				for i := range toEmbeddingSourceColumns {
					if i < len(fromEmbeddingSourceColumns) {
						toEmbeddingSourceColumns[i].SyncFieldsDuringRead(ctx, fromEmbeddingSourceColumns[i])
					}
				}
				to.SetEmbeddingSourceColumns(ctx, toEmbeddingSourceColumns)
			}
		}
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() && to.EmbeddingVectorColumns.IsNull() && len(from.EmbeddingVectorColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingVectorColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingVectorColumns = from.EmbeddingVectorColumns
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() {
		if toEmbeddingVectorColumns, ok := to.GetEmbeddingVectorColumns(ctx); ok {
			if fromEmbeddingVectorColumns, ok := from.GetEmbeddingVectorColumns(ctx); ok {
				for i := range toEmbeddingVectorColumns {
					if i < len(fromEmbeddingVectorColumns) {
						toEmbeddingVectorColumns[i].SyncFieldsDuringRead(ctx, fromEmbeddingVectorColumns[i])
					}
				}
				to.SetEmbeddingVectorColumns(ctx, toEmbeddingVectorColumns)
			}
		}
	}
}

func (m DeltaSyncVectorIndexSpecResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["columns_to_index"] = attrs["columns_to_index"].SetOptional()
	attrs["columns_to_sync"] = attrs["columns_to_sync"].SetOptional()
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
		"columns_to_index":         reflect.TypeOf(types.String{}),
		"columns_to_sync":          reflect.TypeOf(types.String{}),
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
			"columns_to_index":          m.ColumnsToIndex,
			"columns_to_sync":           m.ColumnsToSync,
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
			"columns_to_index": basetypes.ListType{
				ElemType: types.StringType,
			},
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
			"pipeline_id":               types.StringType,
			"pipeline_type":             types.StringType,
			"source_table":              types.StringType,
		},
	}
}

// GetColumnsToIndex returns the value of the ColumnsToIndex field in DeltaSyncVectorIndexSpecResponse as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSyncVectorIndexSpecResponse) GetColumnsToIndex(ctx context.Context) ([]types.String, bool) {
	if m.ColumnsToIndex.IsNull() || m.ColumnsToIndex.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ColumnsToIndex.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumnsToIndex sets the value of the ColumnsToIndex field in DeltaSyncVectorIndexSpecResponse.
func (m *DeltaSyncVectorIndexSpecResponse) SetColumnsToIndex(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns_to_index"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ColumnsToIndex = types.ListValueMust(t, vs)
}

// GetColumnsToSync returns the value of the ColumnsToSync field in DeltaSyncVectorIndexSpecResponse as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaSyncVectorIndexSpecResponse) GetColumnsToSync(ctx context.Context) ([]types.String, bool) {
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

// SetColumnsToSync sets the value of the ColumnsToSync field in DeltaSyncVectorIndexSpecResponse.
func (m *DeltaSyncVectorIndexSpecResponse) SetColumnsToSync(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns_to_sync"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ColumnsToSync = types.ListValueMust(t, vs)
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
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() {
		if toEmbeddingSourceColumns, ok := to.GetEmbeddingSourceColumns(ctx); ok {
			if fromEmbeddingSourceColumns, ok := from.GetEmbeddingSourceColumns(ctx); ok {
				// Recursively sync the fields of each EmbeddingSourceColumns element by position.
				for i := range toEmbeddingSourceColumns {
					if i < len(fromEmbeddingSourceColumns) {
						toEmbeddingSourceColumns[i].SyncFieldsDuringCreateOrUpdate(ctx, fromEmbeddingSourceColumns[i])
					}
				}
				to.SetEmbeddingSourceColumns(ctx, toEmbeddingSourceColumns)
			}
		}
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() && to.EmbeddingVectorColumns.IsNull() && len(from.EmbeddingVectorColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingVectorColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingVectorColumns = from.EmbeddingVectorColumns
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() {
		if toEmbeddingVectorColumns, ok := to.GetEmbeddingVectorColumns(ctx); ok {
			if fromEmbeddingVectorColumns, ok := from.GetEmbeddingVectorColumns(ctx); ok {
				// Recursively sync the fields of each EmbeddingVectorColumns element by position.
				for i := range toEmbeddingVectorColumns {
					if i < len(fromEmbeddingVectorColumns) {
						toEmbeddingVectorColumns[i].SyncFieldsDuringCreateOrUpdate(ctx, fromEmbeddingVectorColumns[i])
					}
				}
				to.SetEmbeddingVectorColumns(ctx, toEmbeddingVectorColumns)
			}
		}
	}
}

func (to *DirectAccessVectorIndexSpec) SyncFieldsDuringRead(ctx context.Context, from DirectAccessVectorIndexSpec) {
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() && to.EmbeddingSourceColumns.IsNull() && len(from.EmbeddingSourceColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingSourceColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingSourceColumns = from.EmbeddingSourceColumns
	}
	if !from.EmbeddingSourceColumns.IsNull() && !from.EmbeddingSourceColumns.IsUnknown() {
		if toEmbeddingSourceColumns, ok := to.GetEmbeddingSourceColumns(ctx); ok {
			if fromEmbeddingSourceColumns, ok := from.GetEmbeddingSourceColumns(ctx); ok {
				for i := range toEmbeddingSourceColumns {
					if i < len(fromEmbeddingSourceColumns) {
						toEmbeddingSourceColumns[i].SyncFieldsDuringRead(ctx, fromEmbeddingSourceColumns[i])
					}
				}
				to.SetEmbeddingSourceColumns(ctx, toEmbeddingSourceColumns)
			}
		}
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() && to.EmbeddingVectorColumns.IsNull() && len(from.EmbeddingVectorColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EmbeddingVectorColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EmbeddingVectorColumns = from.EmbeddingVectorColumns
	}
	if !from.EmbeddingVectorColumns.IsNull() && !from.EmbeddingVectorColumns.IsUnknown() {
		if toEmbeddingVectorColumns, ok := to.GetEmbeddingVectorColumns(ctx); ok {
			if fromEmbeddingVectorColumns, ok := from.GetEmbeddingVectorColumns(ctx); ok {
				for i := range toEmbeddingVectorColumns {
					if i < len(fromEmbeddingVectorColumns) {
						toEmbeddingVectorColumns[i].SyncFieldsDuringRead(ctx, fromEmbeddingVectorColumns[i])
					}
				}
				to.SetEmbeddingVectorColumns(ctx, toEmbeddingVectorColumns)
			}
		}
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
	// The user-selected budget policy id for the endpoint.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
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
	// Name of the AI Search endpoint
	Name types.String `tfsdk:"name"`
	// Number of indexes on the endpoint
	NumIndexes types.Int64 `tfsdk:"num_indexes"`
	// Scaling information for the endpoint
	ScalingInfo types.Object `tfsdk:"scaling_info"`
}

func (to *EndpointInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointInfo) {
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() {
		if toCustomTags, ok := to.GetCustomTags(ctx); ok {
			if fromCustomTags, ok := from.GetCustomTags(ctx); ok {
				// Recursively sync the fields of each CustomTags element by position.
				for i := range toCustomTags {
					if i < len(fromCustomTags) {
						toCustomTags[i].SyncFieldsDuringCreateOrUpdate(ctx, fromCustomTags[i])
					}
				}
				to.SetCustomTags(ctx, toCustomTags)
			}
		}
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
	if !from.ScalingInfo.IsNull() && !from.ScalingInfo.IsUnknown() {
		if toScalingInfo, ok := to.GetScalingInfo(ctx); ok {
			if fromScalingInfo, ok := from.GetScalingInfo(ctx); ok {
				// Recursively sync the fields of ScalingInfo
				toScalingInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromScalingInfo)
				to.SetScalingInfo(ctx, toScalingInfo)
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
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() {
		if toCustomTags, ok := to.GetCustomTags(ctx); ok {
			if fromCustomTags, ok := from.GetCustomTags(ctx); ok {
				for i := range toCustomTags {
					if i < len(fromCustomTags) {
						toCustomTags[i].SyncFieldsDuringRead(ctx, fromCustomTags[i])
					}
				}
				to.SetCustomTags(ctx, toCustomTags)
			}
		}
	}
	if !from.EndpointStatus.IsNull() && !from.EndpointStatus.IsUnknown() {
		if toEndpointStatus, ok := to.GetEndpointStatus(ctx); ok {
			if fromEndpointStatus, ok := from.GetEndpointStatus(ctx); ok {
				toEndpointStatus.SyncFieldsDuringRead(ctx, fromEndpointStatus)
				to.SetEndpointStatus(ctx, toEndpointStatus)
			}
		}
	}
	if !from.ScalingInfo.IsNull() && !from.ScalingInfo.IsUnknown() {
		if toScalingInfo, ok := to.GetScalingInfo(ctx); ok {
			if fromScalingInfo, ok := from.GetScalingInfo(ctx); ok {
				toScalingInfo.SyncFieldsDuringRead(ctx, fromScalingInfo)
				to.SetScalingInfo(ctx, toScalingInfo)
			}
		}
	}
}

func (m EndpointInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
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
	attrs["scaling_info"] = attrs["scaling_info"].SetOptional()

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
		"scaling_info":    reflect.TypeOf(EndpointScalingInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointInfo
// only implements ToObjectValue() and Type().
func (m EndpointInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget_policy_id":           m.BudgetPolicyId,
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
			"scaling_info":               m.ScalingInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_policy_id":   types.StringType,
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
			"scaling_info":               EndpointScalingInfo{}.Type(ctx),
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

// GetScalingInfo returns the value of the ScalingInfo field in EndpointInfo as
// a EndpointScalingInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointInfo) GetScalingInfo(ctx context.Context) (EndpointScalingInfo, bool) {
	var e EndpointScalingInfo
	if m.ScalingInfo.IsNull() || m.ScalingInfo.IsUnknown() {
		return e, false
	}
	var v EndpointScalingInfo
	d := m.ScalingInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetScalingInfo sets the value of the ScalingInfo field in EndpointInfo.
func (m *EndpointInfo) SetScalingInfo(ctx context.Context, v EndpointScalingInfo) {
	vs := v.ToObjectValue(ctx)
	m.ScalingInfo = vs
}

type EndpointScalingInfo struct {
	// The requested QPS target for the endpoint. Best-effort; the system does
	// not guarantee this QPS will be achieved.
	RequestedTargetQps types.Int64 `tfsdk:"requested_target_qps"`
	// The current state of the scaling change request.
	State types.String `tfsdk:"state"`
}

func (to *EndpointScalingInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointScalingInfo) {
}

func (to *EndpointScalingInfo) SyncFieldsDuringRead(ctx context.Context, from EndpointScalingInfo) {
}

func (m EndpointScalingInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["requested_target_qps"] = attrs["requested_target_qps"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointScalingInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointScalingInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointScalingInfo
// only implements ToObjectValue() and Type().
func (m EndpointScalingInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"requested_target_qps": m.RequestedTargetQps,
			"state":                m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointScalingInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"requested_target_qps": types.Int64Type,
			"state":                types.StringType,
		},
	}
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

// Facet aggregation rows returned by a query.
type FacetResultData struct {
	// Facet rows. Each row is `[facet_column_name, value_or_range, count]`.
	FacetArray types.List `tfsdk:"facet_array"`
	// Number of facet rows returned.
	FacetRowCount types.Int64 `tfsdk:"facet_row_count"`
}

func (to *FacetResultData) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FacetResultData) {
	if !from.FacetArray.IsNull() && !from.FacetArray.IsUnknown() && to.FacetArray.IsNull() && len(from.FacetArray.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FacetArray, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FacetArray = from.FacetArray
	}
}

func (to *FacetResultData) SyncFieldsDuringRead(ctx context.Context, from FacetResultData) {
	if !from.FacetArray.IsNull() && !from.FacetArray.IsUnknown() && to.FacetArray.IsNull() && len(from.FacetArray.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FacetArray, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FacetArray = from.FacetArray
	}
}

func (m FacetResultData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["facet_array"] = attrs["facet_array"].SetOptional()
	attrs["facet_row_count"] = attrs["facet_row_count"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FacetResultData.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FacetResultData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"facet_array": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FacetResultData
// only implements ToObjectValue() and Type().
func (m FacetResultData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"facet_array":     m.FacetArray,
			"facet_row_count": m.FacetRowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FacetResultData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"facet_array": basetypes.ListType{
				ElemType: basetypes.ListType{
					ElemType: types.StringType,
				},
			},
			"facet_row_count": types.Int64Type,
		},
	}
}

// GetFacetArray returns the value of the FacetArray field in FacetResultData as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *FacetResultData) GetFacetArray(ctx context.Context) ([]types.String, bool) {
	if m.FacetArray.IsNull() || m.FacetArray.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.FacetArray.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFacetArray sets the value of the FacetArray field in FacetResultData.
func (m *FacetResultData) SetFacetArray(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["facet_array"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FacetArray = types.ListValueMust(t, vs)
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

type GetVectorSearchEndpointPermissionLevelsRequest struct {
	// The vector search endpoint for which to get or manage permissions.
	EndpointId types.String `tfsdk:"-"`
}

func (to *GetVectorSearchEndpointPermissionLevelsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetVectorSearchEndpointPermissionLevelsRequest) {
}

func (to *GetVectorSearchEndpointPermissionLevelsRequest) SyncFieldsDuringRead(ctx context.Context, from GetVectorSearchEndpointPermissionLevelsRequest) {
}

func (m GetVectorSearchEndpointPermissionLevelsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint_id"] = attrs["endpoint_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetVectorSearchEndpointPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetVectorSearchEndpointPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetVectorSearchEndpointPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (m GetVectorSearchEndpointPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_id": m.EndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetVectorSearchEndpointPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_id": types.StringType,
		},
	}
}

type GetVectorSearchEndpointPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (to *GetVectorSearchEndpointPermissionLevelsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetVectorSearchEndpointPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() {
		if toPermissionLevels, ok := to.GetPermissionLevels(ctx); ok {
			if fromPermissionLevels, ok := from.GetPermissionLevels(ctx); ok {
				// Recursively sync the fields of each PermissionLevels element by position.
				for i := range toPermissionLevels {
					if i < len(fromPermissionLevels) {
						toPermissionLevels[i].SyncFieldsDuringCreateOrUpdate(ctx, fromPermissionLevels[i])
					}
				}
				to.SetPermissionLevels(ctx, toPermissionLevels)
			}
		}
	}
}

func (to *GetVectorSearchEndpointPermissionLevelsResponse) SyncFieldsDuringRead(ctx context.Context, from GetVectorSearchEndpointPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() {
		if toPermissionLevels, ok := to.GetPermissionLevels(ctx); ok {
			if fromPermissionLevels, ok := from.GetPermissionLevels(ctx); ok {
				for i := range toPermissionLevels {
					if i < len(fromPermissionLevels) {
						toPermissionLevels[i].SyncFieldsDuringRead(ctx, fromPermissionLevels[i])
					}
				}
				to.SetPermissionLevels(ctx, toPermissionLevels)
			}
		}
	}
}

func (m GetVectorSearchEndpointPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_levels"] = attrs["permission_levels"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetVectorSearchEndpointPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetVectorSearchEndpointPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(VectorSearchEndpointPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetVectorSearchEndpointPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (m GetVectorSearchEndpointPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetVectorSearchEndpointPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: VectorSearchEndpointPermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetVectorSearchEndpointPermissionLevelsResponse as
// a slice of VectorSearchEndpointPermissionsDescription values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetVectorSearchEndpointPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]VectorSearchEndpointPermissionsDescription, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []VectorSearchEndpointPermissionsDescription
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetVectorSearchEndpointPermissionLevelsResponse.
func (m *GetVectorSearchEndpointPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []VectorSearchEndpointPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetVectorSearchEndpointPermissionsRequest struct {
	// The vector search endpoint for which to get or manage permissions.
	EndpointId types.String `tfsdk:"-"`
}

func (to *GetVectorSearchEndpointPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetVectorSearchEndpointPermissionsRequest) {
}

func (to *GetVectorSearchEndpointPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from GetVectorSearchEndpointPermissionsRequest) {
}

func (m GetVectorSearchEndpointPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoint_id"] = attrs["endpoint_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetVectorSearchEndpointPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetVectorSearchEndpointPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetVectorSearchEndpointPermissionsRequest
// only implements ToObjectValue() and Type().
func (m GetVectorSearchEndpointPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_id": m.EndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetVectorSearchEndpointPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_id": types.StringType,
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
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() {
		if toEndpoints, ok := to.GetEndpoints(ctx); ok {
			if fromEndpoints, ok := from.GetEndpoints(ctx); ok {
				// Recursively sync the fields of each Endpoints element by position.
				for i := range toEndpoints {
					if i < len(fromEndpoints) {
						toEndpoints[i].SyncFieldsDuringCreateOrUpdate(ctx, fromEndpoints[i])
					}
				}
				to.SetEndpoints(ctx, toEndpoints)
			}
		}
	}
}

func (to *ListEndpointResponse) SyncFieldsDuringRead(ctx context.Context, from ListEndpointResponse) {
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() && to.Endpoints.IsNull() && len(from.Endpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Endpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Endpoints = from.Endpoints
	}
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() {
		if toEndpoints, ok := to.GetEndpoints(ctx); ok {
			if fromEndpoints, ok := from.GetEndpoints(ctx); ok {
				for i := range toEndpoints {
					if i < len(fromEndpoints) {
						toEndpoints[i].SyncFieldsDuringRead(ctx, fromEndpoints[i])
					}
				}
				to.SetEndpoints(ctx, toEndpoints)
			}
		}
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
	if !from.Values.IsNull() && !from.Values.IsUnknown() {
		if toValues, ok := to.GetValues(ctx); ok {
			if fromValues, ok := from.GetValues(ctx); ok {
				// Recursively sync the fields of each Values element by position.
				for i := range toValues {
					if i < len(fromValues) {
						toValues[i].SyncFieldsDuringCreateOrUpdate(ctx, fromValues[i])
					}
				}
				to.SetValues(ctx, toValues)
			}
		}
	}
}

func (to *ListValue) SyncFieldsDuringRead(ctx context.Context, from ListValue) {
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
	if !from.Values.IsNull() && !from.Values.IsUnknown() {
		if toValues, ok := to.GetValues(ctx); ok {
			if fromValues, ok := from.GetValues(ctx); ok {
				for i := range toValues {
					if i < len(fromValues) {
						toValues[i].SyncFieldsDuringRead(ctx, fromValues[i])
					}
				}
				to.SetValues(ctx, toValues)
			}
		}
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
	if !from.VectorIndexes.IsNull() && !from.VectorIndexes.IsUnknown() {
		if toVectorIndexes, ok := to.GetVectorIndexes(ctx); ok {
			if fromVectorIndexes, ok := from.GetVectorIndexes(ctx); ok {
				// Recursively sync the fields of each VectorIndexes element by position.
				for i := range toVectorIndexes {
					if i < len(fromVectorIndexes) {
						toVectorIndexes[i].SyncFieldsDuringCreateOrUpdate(ctx, fromVectorIndexes[i])
					}
				}
				to.SetVectorIndexes(ctx, toVectorIndexes)
			}
		}
	}
}

func (to *ListVectorIndexesResponse) SyncFieldsDuringRead(ctx context.Context, from ListVectorIndexesResponse) {
	if !from.VectorIndexes.IsNull() && !from.VectorIndexes.IsUnknown() && to.VectorIndexes.IsNull() && len(from.VectorIndexes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for VectorIndexes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.VectorIndexes = from.VectorIndexes
	}
	if !from.VectorIndexes.IsNull() && !from.VectorIndexes.IsUnknown() {
		if toVectorIndexes, ok := to.GetVectorIndexes(ctx); ok {
			if fromVectorIndexes, ok := from.GetVectorIndexes(ctx); ok {
				for i := range toVectorIndexes {
					if i < len(fromVectorIndexes) {
						toVectorIndexes[i].SyncFieldsDuringRead(ctx, fromVectorIndexes[i])
					}
				}
				to.SetVectorIndexes(ctx, toVectorIndexes)
			}
		}
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

// Metric specification
type Metric struct {
	// Metric labels
	Labels types.List `tfsdk:"labels"`
	// Metric name
	Name types.String `tfsdk:"name"`
	// Percentile for the metric
	Percentile types.Float64 `tfsdk:"percentile"`
}

func (to *Metric) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Metric) {
	if !from.Labels.IsNull() && !from.Labels.IsUnknown() && to.Labels.IsNull() && len(from.Labels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Labels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Labels = from.Labels
	}
	if !from.Labels.IsNull() && !from.Labels.IsUnknown() {
		if toLabels, ok := to.GetLabels(ctx); ok {
			if fromLabels, ok := from.GetLabels(ctx); ok {
				// Recursively sync the fields of each Labels element by position.
				for i := range toLabels {
					if i < len(fromLabels) {
						toLabels[i].SyncFieldsDuringCreateOrUpdate(ctx, fromLabels[i])
					}
				}
				to.SetLabels(ctx, toLabels)
			}
		}
	}
}

func (to *Metric) SyncFieldsDuringRead(ctx context.Context, from Metric) {
	if !from.Labels.IsNull() && !from.Labels.IsUnknown() && to.Labels.IsNull() && len(from.Labels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Labels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Labels = from.Labels
	}
	if !from.Labels.IsNull() && !from.Labels.IsUnknown() {
		if toLabels, ok := to.GetLabels(ctx); ok {
			if fromLabels, ok := from.GetLabels(ctx); ok {
				for i := range toLabels {
					if i < len(fromLabels) {
						toLabels[i].SyncFieldsDuringRead(ctx, fromLabels[i])
					}
				}
				to.SetLabels(ctx, toLabels)
			}
		}
	}
}

func (m Metric) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["labels"] = attrs["labels"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["percentile"] = attrs["percentile"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Metric.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Metric) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"labels": reflect.TypeOf(MetricLabel{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Metric
// only implements ToObjectValue() and Type().
func (m Metric) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"labels":     m.Labels,
			"name":       m.Name,
			"percentile": m.Percentile,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Metric) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"labels": basetypes.ListType{
				ElemType: MetricLabel{}.Type(ctx),
			},
			"name":       types.StringType,
			"percentile": types.Float64Type,
		},
	}
}

// GetLabels returns the value of the Labels field in Metric as
// a slice of MetricLabel values.
// If the field is unknown or null, the boolean return value is false.
func (m *Metric) GetLabels(ctx context.Context) ([]MetricLabel, bool) {
	if m.Labels.IsNull() || m.Labels.IsUnknown() {
		return nil, false
	}
	var v []MetricLabel
	d := m.Labels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLabels sets the value of the Labels field in Metric.
func (m *Metric) SetLabels(ctx context.Context, v []MetricLabel) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["labels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Labels = types.ListValueMust(t, vs)
}

// Label for a metric
type MetricLabel struct {
	// Label name
	Name types.String `tfsdk:"name"`
	// Label value
	Value types.String `tfsdk:"value"`
}

func (to *MetricLabel) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MetricLabel) {
}

func (to *MetricLabel) SyncFieldsDuringRead(ctx context.Context, from MetricLabel) {
}

func (m MetricLabel) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MetricLabel.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m MetricLabel) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MetricLabel
// only implements ToObjectValue() and Type().
func (m MetricLabel) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":  m.Name,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m MetricLabel) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":  types.StringType,
			"value": types.StringType,
		},
	}
}

// Single metric value at a specific timestamp
type MetricValue struct {
	// Timestamp of the metric value (milliseconds since epoch)
	Timestamp types.Int64 `tfsdk:"timestamp"`
	// Metric value
	Value types.Float64 `tfsdk:"value"`
}

func (to *MetricValue) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MetricValue) {
}

func (to *MetricValue) SyncFieldsDuringRead(ctx context.Context, from MetricValue) {
}

func (m MetricValue) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["timestamp"] = attrs["timestamp"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MetricValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m MetricValue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MetricValue
// only implements ToObjectValue() and Type().
func (m MetricValue) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"timestamp": m.Timestamp,
			"value":     m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m MetricValue) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"timestamp": types.Int64Type,
			"value":     types.Float64Type,
		},
	}
}

// Collection of metric values for a specific metric
type MetricValues struct {
	// Metric specification
	Metric types.Object `tfsdk:"metric"`
	// Time series of metric values
	Values types.List `tfsdk:"values"`
}

func (to *MetricValues) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MetricValues) {
	if !from.Metric.IsNull() && !from.Metric.IsUnknown() {
		if toMetric, ok := to.GetMetric(ctx); ok {
			if fromMetric, ok := from.GetMetric(ctx); ok {
				// Recursively sync the fields of Metric
				toMetric.SyncFieldsDuringCreateOrUpdate(ctx, fromMetric)
				to.SetMetric(ctx, toMetric)
			}
		}
	}
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
	if !from.Values.IsNull() && !from.Values.IsUnknown() {
		if toValues, ok := to.GetValues(ctx); ok {
			if fromValues, ok := from.GetValues(ctx); ok {
				// Recursively sync the fields of each Values element by position.
				for i := range toValues {
					if i < len(fromValues) {
						toValues[i].SyncFieldsDuringCreateOrUpdate(ctx, fromValues[i])
					}
				}
				to.SetValues(ctx, toValues)
			}
		}
	}
}

func (to *MetricValues) SyncFieldsDuringRead(ctx context.Context, from MetricValues) {
	if !from.Metric.IsNull() && !from.Metric.IsUnknown() {
		if toMetric, ok := to.GetMetric(ctx); ok {
			if fromMetric, ok := from.GetMetric(ctx); ok {
				toMetric.SyncFieldsDuringRead(ctx, fromMetric)
				to.SetMetric(ctx, toMetric)
			}
		}
	}
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
	if !from.Values.IsNull() && !from.Values.IsUnknown() {
		if toValues, ok := to.GetValues(ctx); ok {
			if fromValues, ok := from.GetValues(ctx); ok {
				for i := range toValues {
					if i < len(fromValues) {
						toValues[i].SyncFieldsDuringRead(ctx, fromValues[i])
					}
				}
				to.SetValues(ctx, toValues)
			}
		}
	}
}

func (m MetricValues) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["metric"] = attrs["metric"].SetOptional()
	attrs["values"] = attrs["values"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MetricValues.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m MetricValues) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metric": reflect.TypeOf(Metric{}),
		"values": reflect.TypeOf(MetricValue{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MetricValues
// only implements ToObjectValue() and Type().
func (m MetricValues) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metric": m.Metric,
			"values": m.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (m MetricValues) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metric": Metric{}.Type(ctx),
			"values": basetypes.ListType{
				ElemType: MetricValue{}.Type(ctx),
			},
		},
	}
}

// GetMetric returns the value of the Metric field in MetricValues as
// a Metric value.
// If the field is unknown or null, the boolean return value is false.
func (m *MetricValues) GetMetric(ctx context.Context) (Metric, bool) {
	var e Metric
	if m.Metric.IsNull() || m.Metric.IsUnknown() {
		return e, false
	}
	var v Metric
	d := m.Metric.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetric sets the value of the Metric field in MetricValues.
func (m *MetricValues) SetMetric(ctx context.Context, v Metric) {
	vs := v.ToObjectValue(ctx)
	m.Metric = vs
}

// GetValues returns the value of the Values field in MetricValues as
// a slice of MetricValue values.
// If the field is unknown or null, the boolean return value is false.
func (m *MetricValues) GetValues(ctx context.Context) ([]MetricValue, bool) {
	if m.Values.IsNull() || m.Values.IsUnknown() {
		return nil, false
	}
	var v []MetricValue
	d := m.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in MetricValues.
func (m *MetricValues) SetValues(ctx context.Context, v []MetricValue) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Values = types.ListValueMust(t, vs)
}

type MiniVectorIndex struct {
	// The user who created the index.
	Creator types.String `tfsdk:"creator"`
	// ID of the endpoint associated with the index.
	EndpointId types.String `tfsdk:"endpoint_id"`
	// Name of the endpoint associated with the index
	EndpointName types.String `tfsdk:"endpoint_name"`
	// The subtype of the index.
	IndexSubtype types.String `tfsdk:"index_subtype"`

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
	attrs["endpoint_id"] = attrs["endpoint_id"].SetOptional()
	attrs["endpoint_name"] = attrs["endpoint_name"].SetOptional()
	attrs["index_subtype"] = attrs["index_subtype"].SetOptional()
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
			"endpoint_id":   m.EndpointId,
			"endpoint_name": m.EndpointName,
			"index_subtype": m.IndexSubtype,
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
			"endpoint_id":   types.StringType,
			"endpoint_name": types.StringType,
			"index_subtype": types.StringType,
			"index_type":    types.StringType,
			"name":          types.StringType,
			"primary_key":   types.StringType,
		},
	}
}

type PatchEndpointBudgetPolicyRequest struct {
	// The budget policy id to be applied
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// Name of the AI Search endpoint
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
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// The budget policy applied to the AI Search endpoint.
	EffectiveBudgetPolicyId types.String `tfsdk:"effective_budget_policy_id"`
}

func (to *PatchEndpointBudgetPolicyResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PatchEndpointBudgetPolicyResponse) {
}

func (to *PatchEndpointBudgetPolicyResponse) SyncFieldsDuringRead(ctx context.Context, from PatchEndpointBudgetPolicyResponse) {
}

func (m PatchEndpointBudgetPolicyResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
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
			"budget_policy_id":           m.BudgetPolicyId,
			"effective_budget_policy_id": m.EffectiveBudgetPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PatchEndpointBudgetPolicyResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_policy_id":           types.StringType,
			"effective_budget_policy_id": types.StringType,
		},
	}
}

type PatchEndpointRequest struct {
	// Name of the AI Search endpoint
	EndpointName types.String `tfsdk:"-"`
	// Target QPS for the endpoint. Best-effort; the system does not guarantee
	// this QPS will be achieved.
	TargetQps types.Int64 `tfsdk:"target_qps"`
}

func (to *PatchEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PatchEndpointRequest) {
}

func (to *PatchEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from PatchEndpointRequest) {
}

func (m PatchEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["target_qps"] = attrs["target_qps"].SetOptional()
	attrs["endpoint_name"] = attrs["endpoint_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PatchEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PatchEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchEndpointRequest
// only implements ToObjectValue() and Type().
func (m PatchEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoint_name": m.EndpointName,
			"target_qps":    m.TargetQps,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PatchEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoint_name": types.StringType,
			"target_qps":    types.Int64Type,
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
	// Facets to compute over the matched results. Each entry has one of these
	// forms: `"<column>"` - top 10 distinct values by count `"<column> TOP
	// <n>"` - top n distinct values, where n > 0 `"<column> BUCKETS
	// [[from,to],...]"` - inclusive numeric ranges `TOP` and `BUCKETS` are
	// case-insensitive. A column may appear at most once.
	Facets types.List `tfsdk:"facets"`
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
	// Text columns to search for `query_text`. When empty, all text columns are
	// searched.
	QueryColumns types.List `tfsdk:"query_columns"`
	// Query text. Required for Delta Sync Index using model endpoint.
	QueryText types.String `tfsdk:"query_text"`
	// The query type to use. Choices are `ANN` and `HYBRID` and `FULL_TEXT`.
	// Defaults to `ANN`.
	QueryType types.String `tfsdk:"query_type"`
	// Query vector. Required for Direct Vector Access Index and Delta Sync
	// Index using self-managed vectors.
	QueryVector types.List `tfsdk:"query_vector"`
	// If set, the top 50 results are reranked with the Databricks Reranker
	// model before returning the `num_results` results to the user. The setting
	// `columns_to_rerank` selects which columns are used for reranking. For
	// each datapoint, the columns selected are concatenated before being sent
	// to the reranking model. See
	// https://docs.databricks.com/aws/en/vector-search/query-vector-search#rerank
	// for more information.
	Reranker types.Object `tfsdk:"reranker"`
	// Threshold for the approximate nearest neighbor search. Defaults to 0.0.
	ScoreThreshold types.Float64 `tfsdk:"score_threshold"`
	// Sort results by column values instead of the default relevance ordering.
	// Each clause has the form `"<column> ASC"` or `"<column> DESC"`, for
	// example `["rating DESC", "price ASC"]`.
	SortColumns types.List `tfsdk:"sort_columns"`
}

func (to *QueryVectorIndexRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryVectorIndexRequest) {
	if !from.ColumnsToRerank.IsNull() && !from.ColumnsToRerank.IsUnknown() && to.ColumnsToRerank.IsNull() && len(from.ColumnsToRerank.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnsToRerank, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnsToRerank = from.ColumnsToRerank
	}
	if !from.Facets.IsNull() && !from.Facets.IsUnknown() && to.Facets.IsNull() && len(from.Facets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Facets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Facets = from.Facets
	}
	if !from.QueryColumns.IsNull() && !from.QueryColumns.IsUnknown() && to.QueryColumns.IsNull() && len(from.QueryColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for QueryColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.QueryColumns = from.QueryColumns
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
	if !from.SortColumns.IsNull() && !from.SortColumns.IsUnknown() && to.SortColumns.IsNull() && len(from.SortColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SortColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SortColumns = from.SortColumns
	}
}

func (to *QueryVectorIndexRequest) SyncFieldsDuringRead(ctx context.Context, from QueryVectorIndexRequest) {
	if !from.ColumnsToRerank.IsNull() && !from.ColumnsToRerank.IsUnknown() && to.ColumnsToRerank.IsNull() && len(from.ColumnsToRerank.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ColumnsToRerank, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ColumnsToRerank = from.ColumnsToRerank
	}
	if !from.Facets.IsNull() && !from.Facets.IsUnknown() && to.Facets.IsNull() && len(from.Facets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Facets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Facets = from.Facets
	}
	if !from.QueryColumns.IsNull() && !from.QueryColumns.IsUnknown() && to.QueryColumns.IsNull() && len(from.QueryColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for QueryColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.QueryColumns = from.QueryColumns
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
	if !from.SortColumns.IsNull() && !from.SortColumns.IsUnknown() && to.SortColumns.IsNull() && len(from.SortColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SortColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SortColumns = from.SortColumns
	}
}

func (m QueryVectorIndexRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["columns"] = attrs["columns"].SetRequired()
	attrs["columns_to_rerank"] = attrs["columns_to_rerank"].SetOptional()
	attrs["facets"] = attrs["facets"].SetOptional()
	attrs["filters_json"] = attrs["filters_json"].SetOptional()
	attrs["num_results"] = attrs["num_results"].SetOptional()
	attrs["query_columns"] = attrs["query_columns"].SetOptional()
	attrs["query_text"] = attrs["query_text"].SetOptional()
	attrs["query_type"] = attrs["query_type"].SetOptional()
	attrs["query_vector"] = attrs["query_vector"].SetOptional()
	attrs["reranker"] = attrs["reranker"].SetOptional()
	attrs["score_threshold"] = attrs["score_threshold"].SetOptional()
	attrs["sort_columns"] = attrs["sort_columns"].SetOptional()
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
		"facets":            reflect.TypeOf(types.String{}),
		"query_columns":     reflect.TypeOf(types.String{}),
		"query_vector":      reflect.TypeOf(types.Float64{}),
		"reranker":          reflect.TypeOf(RerankerConfig{}),
		"sort_columns":      reflect.TypeOf(types.String{}),
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
			"facets":            m.Facets,
			"filters_json":      m.FiltersJson,
			"index_name":        m.IndexName,
			"num_results":       m.NumResults,
			"query_columns":     m.QueryColumns,
			"query_text":        m.QueryText,
			"query_type":        m.QueryType,
			"query_vector":      m.QueryVector,
			"reranker":          m.Reranker,
			"score_threshold":   m.ScoreThreshold,
			"sort_columns":      m.SortColumns,
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
			"facets": basetypes.ListType{
				ElemType: types.StringType,
			},
			"filters_json": types.StringType,
			"index_name":   types.StringType,
			"num_results":  types.Int64Type,
			"query_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"query_text": types.StringType,
			"query_type": types.StringType,
			"query_vector": basetypes.ListType{
				ElemType: types.Float64Type,
			},
			"reranker":        RerankerConfig{}.Type(ctx),
			"score_threshold": types.Float64Type,
			"sort_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
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

// GetFacets returns the value of the Facets field in QueryVectorIndexRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryVectorIndexRequest) GetFacets(ctx context.Context) ([]types.String, bool) {
	if m.Facets.IsNull() || m.Facets.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Facets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFacets sets the value of the Facets field in QueryVectorIndexRequest.
func (m *QueryVectorIndexRequest) SetFacets(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["facets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Facets = types.ListValueMust(t, vs)
}

// GetQueryColumns returns the value of the QueryColumns field in QueryVectorIndexRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryVectorIndexRequest) GetQueryColumns(ctx context.Context) ([]types.String, bool) {
	if m.QueryColumns.IsNull() || m.QueryColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.QueryColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQueryColumns sets the value of the QueryColumns field in QueryVectorIndexRequest.
func (m *QueryVectorIndexRequest) SetQueryColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["query_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.QueryColumns = types.ListValueMust(t, vs)
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

// GetSortColumns returns the value of the SortColumns field in QueryVectorIndexRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryVectorIndexRequest) GetSortColumns(ctx context.Context) ([]types.String, bool) {
	if m.SortColumns.IsNull() || m.SortColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.SortColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSortColumns sets the value of the SortColumns field in QueryVectorIndexRequest.
func (m *QueryVectorIndexRequest) SetSortColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["sort_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SortColumns = types.ListValueMust(t, vs)
}

type QueryVectorIndexResponse struct {
	// Facet aggregation rows returned by a query.
	FacetResult types.Object `tfsdk:"facet_result"`
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
	if !from.FacetResult.IsNull() && !from.FacetResult.IsUnknown() {
		if toFacetResult, ok := to.GetFacetResult(ctx); ok {
			if fromFacetResult, ok := from.GetFacetResult(ctx); ok {
				// Recursively sync the fields of FacetResult
				toFacetResult.SyncFieldsDuringCreateOrUpdate(ctx, fromFacetResult)
				to.SetFacetResult(ctx, toFacetResult)
			}
		}
	}
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
	if !from.FacetResult.IsNull() && !from.FacetResult.IsUnknown() {
		if toFacetResult, ok := to.GetFacetResult(ctx); ok {
			if fromFacetResult, ok := from.GetFacetResult(ctx); ok {
				toFacetResult.SyncFieldsDuringRead(ctx, fromFacetResult)
				to.SetFacetResult(ctx, toFacetResult)
			}
		}
	}
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
	attrs["facet_result"] = attrs["facet_result"].SetOptional()
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
		"facet_result": reflect.TypeOf(FacetResultData{}),
		"manifest":     reflect.TypeOf(ResultManifest{}),
		"result":       reflect.TypeOf(ResultData{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryVectorIndexResponse
// only implements ToObjectValue() and Type().
func (m QueryVectorIndexResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"facet_result":    m.FacetResult,
			"manifest":        m.Manifest,
			"next_page_token": m.NextPageToken,
			"result":          m.Result,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryVectorIndexResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"facet_result":    FacetResultData{}.Type(ctx),
			"manifest":        ResultManifest{}.Type(ctx),
			"next_page_token": types.StringType,
			"result":          ResultData{}.Type(ctx),
		},
	}
}

// GetFacetResult returns the value of the FacetResult field in QueryVectorIndexResponse as
// a FacetResultData value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryVectorIndexResponse) GetFacetResult(ctx context.Context) (FacetResultData, bool) {
	var e FacetResultData
	if m.FacetResult.IsNull() || m.FacetResult.IsUnknown() {
		return e, false
	}
	var v FacetResultData
	d := m.FacetResult.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFacetResult sets the value of the FacetResult field in QueryVectorIndexResponse.
func (m *QueryVectorIndexResponse) SetFacetResult(ctx context.Context, v FacetResultData) {
	vs := v.ToObjectValue(ctx)
	m.FacetResult = vs
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
	// Reranker identifier: - When model_type=BASE/UNSPECIFIED: must be
	// "databricks_reranker". - When model_type=FINETUNED: the Model Serving
	// endpoint name hosting a finetuned reranker.
	Model types.String `tfsdk:"model"`
	// Parameters that control how the reranker processes the query results.
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
	// Number of columns in `facet_result`.
	FacetColumnCount types.Int64 `tfsdk:"facet_column_count"`
	// Information about each column in `facet_result`.
	FacetColumns types.List `tfsdk:"facet_columns"`
}

func (to *ResultManifest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ResultManifest) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() {
		if toColumns, ok := to.GetColumns(ctx); ok {
			if fromColumns, ok := from.GetColumns(ctx); ok {
				// Recursively sync the fields of each Columns element by position.
				for i := range toColumns {
					if i < len(fromColumns) {
						toColumns[i].SyncFieldsDuringCreateOrUpdate(ctx, fromColumns[i])
					}
				}
				to.SetColumns(ctx, toColumns)
			}
		}
	}
	if !from.FacetColumns.IsNull() && !from.FacetColumns.IsUnknown() && to.FacetColumns.IsNull() && len(from.FacetColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FacetColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FacetColumns = from.FacetColumns
	}
	if !from.FacetColumns.IsNull() && !from.FacetColumns.IsUnknown() {
		if toFacetColumns, ok := to.GetFacetColumns(ctx); ok {
			if fromFacetColumns, ok := from.GetFacetColumns(ctx); ok {
				// Recursively sync the fields of each FacetColumns element by position.
				for i := range toFacetColumns {
					if i < len(fromFacetColumns) {
						toFacetColumns[i].SyncFieldsDuringCreateOrUpdate(ctx, fromFacetColumns[i])
					}
				}
				to.SetFacetColumns(ctx, toFacetColumns)
			}
		}
	}
}

func (to *ResultManifest) SyncFieldsDuringRead(ctx context.Context, from ResultManifest) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() {
		if toColumns, ok := to.GetColumns(ctx); ok {
			if fromColumns, ok := from.GetColumns(ctx); ok {
				for i := range toColumns {
					if i < len(fromColumns) {
						toColumns[i].SyncFieldsDuringRead(ctx, fromColumns[i])
					}
				}
				to.SetColumns(ctx, toColumns)
			}
		}
	}
	if !from.FacetColumns.IsNull() && !from.FacetColumns.IsUnknown() && to.FacetColumns.IsNull() && len(from.FacetColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FacetColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FacetColumns = from.FacetColumns
	}
	if !from.FacetColumns.IsNull() && !from.FacetColumns.IsUnknown() {
		if toFacetColumns, ok := to.GetFacetColumns(ctx); ok {
			if fromFacetColumns, ok := from.GetFacetColumns(ctx); ok {
				for i := range toFacetColumns {
					if i < len(fromFacetColumns) {
						toFacetColumns[i].SyncFieldsDuringRead(ctx, fromFacetColumns[i])
					}
				}
				to.SetFacetColumns(ctx, toFacetColumns)
			}
		}
	}
}

func (m ResultManifest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["column_count"] = attrs["column_count"].SetOptional()
	attrs["columns"] = attrs["columns"].SetOptional()
	attrs["facet_column_count"] = attrs["facet_column_count"].SetOptional()
	attrs["facet_columns"] = attrs["facet_columns"].SetOptional()

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
		"columns":       reflect.TypeOf(ColumnInfo{}),
		"facet_columns": reflect.TypeOf(ColumnInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ResultManifest
// only implements ToObjectValue() and Type().
func (m ResultManifest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"column_count":       m.ColumnCount,
			"columns":            m.Columns,
			"facet_column_count": m.FacetColumnCount,
			"facet_columns":      m.FacetColumns,
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
			"facet_column_count": types.Int64Type,
			"facet_columns": basetypes.ListType{
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

// GetFacetColumns returns the value of the FacetColumns field in ResultManifest as
// a slice of ColumnInfo values.
// If the field is unknown or null, the boolean return value is false.
func (m *ResultManifest) GetFacetColumns(ctx context.Context) ([]ColumnInfo, bool) {
	if m.FacetColumns.IsNull() || m.FacetColumns.IsUnknown() {
		return nil, false
	}
	var v []ColumnInfo
	d := m.FacetColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFacetColumns sets the value of the FacetColumns field in ResultManifest.
func (m *ResultManifest) SetFacetColumns(ctx context.Context, v []ColumnInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["facet_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FacetColumns = types.ListValueMust(t, vs)
}

// Request to retrieve user-visible metrics
type RetrieveUserVisibleMetricsRequest struct {
	// End time for metrics query
	EndTime types.String `tfsdk:"end_time"`
	// Granularity in seconds
	GranularityInSeconds types.Int64 `tfsdk:"granularity_in_seconds"`
	// List of metrics to retrieve
	Metrics types.List `tfsdk:"metrics"`
	// AI Search endpoint name
	Name types.String `tfsdk:"-"`
	// Token for pagination
	PageToken types.String `tfsdk:"page_token"`
	// Start time for metrics query
	StartTime types.String `tfsdk:"start_time"`
}

func (to *RetrieveUserVisibleMetricsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RetrieveUserVisibleMetricsRequest) {
	if !from.Metrics.IsNull() && !from.Metrics.IsUnknown() && to.Metrics.IsNull() && len(from.Metrics.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Metrics, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Metrics = from.Metrics
	}
	if !from.Metrics.IsNull() && !from.Metrics.IsUnknown() {
		if toMetrics, ok := to.GetMetrics(ctx); ok {
			if fromMetrics, ok := from.GetMetrics(ctx); ok {
				// Recursively sync the fields of each Metrics element by position.
				for i := range toMetrics {
					if i < len(fromMetrics) {
						toMetrics[i].SyncFieldsDuringCreateOrUpdate(ctx, fromMetrics[i])
					}
				}
				to.SetMetrics(ctx, toMetrics)
			}
		}
	}
}

func (to *RetrieveUserVisibleMetricsRequest) SyncFieldsDuringRead(ctx context.Context, from RetrieveUserVisibleMetricsRequest) {
	if !from.Metrics.IsNull() && !from.Metrics.IsUnknown() && to.Metrics.IsNull() && len(from.Metrics.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Metrics, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Metrics = from.Metrics
	}
	if !from.Metrics.IsNull() && !from.Metrics.IsUnknown() {
		if toMetrics, ok := to.GetMetrics(ctx); ok {
			if fromMetrics, ok := from.GetMetrics(ctx); ok {
				for i := range toMetrics {
					if i < len(fromMetrics) {
						toMetrics[i].SyncFieldsDuringRead(ctx, fromMetrics[i])
					}
				}
				to.SetMetrics(ctx, toMetrics)
			}
		}
	}
}

func (m RetrieveUserVisibleMetricsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["end_time"] = attrs["end_time"].SetOptional()
	attrs["granularity_in_seconds"] = attrs["granularity_in_seconds"].SetOptional()
	attrs["metrics"] = attrs["metrics"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["start_time"] = attrs["start_time"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RetrieveUserVisibleMetricsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RetrieveUserVisibleMetricsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metrics": reflect.TypeOf(Metric{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RetrieveUserVisibleMetricsRequest
// only implements ToObjectValue() and Type().
func (m RetrieveUserVisibleMetricsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"end_time":               m.EndTime,
			"granularity_in_seconds": m.GranularityInSeconds,
			"metrics":                m.Metrics,
			"name":                   m.Name,
			"page_token":             m.PageToken,
			"start_time":             m.StartTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RetrieveUserVisibleMetricsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"end_time":               types.StringType,
			"granularity_in_seconds": types.Int64Type,
			"metrics": basetypes.ListType{
				ElemType: Metric{}.Type(ctx),
			},
			"name":       types.StringType,
			"page_token": types.StringType,
			"start_time": types.StringType,
		},
	}
}

// GetMetrics returns the value of the Metrics field in RetrieveUserVisibleMetricsRequest as
// a slice of Metric values.
// If the field is unknown or null, the boolean return value is false.
func (m *RetrieveUserVisibleMetricsRequest) GetMetrics(ctx context.Context) ([]Metric, bool) {
	if m.Metrics.IsNull() || m.Metrics.IsUnknown() {
		return nil, false
	}
	var v []Metric
	d := m.Metrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetrics sets the value of the Metrics field in RetrieveUserVisibleMetricsRequest.
func (m *RetrieveUserVisibleMetricsRequest) SetMetrics(ctx context.Context, v []Metric) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Metrics = types.ListValueMust(t, vs)
}

// Response containing user-visible metrics
type RetrieveUserVisibleMetricsResponse struct {
	// Collection of metric values
	MetricValues types.List `tfsdk:"metric_values"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *RetrieveUserVisibleMetricsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RetrieveUserVisibleMetricsResponse) {
	if !from.MetricValues.IsNull() && !from.MetricValues.IsUnknown() && to.MetricValues.IsNull() && len(from.MetricValues.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for MetricValues, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.MetricValues = from.MetricValues
	}
	if !from.MetricValues.IsNull() && !from.MetricValues.IsUnknown() {
		if toMetricValues, ok := to.GetMetricValues(ctx); ok {
			if fromMetricValues, ok := from.GetMetricValues(ctx); ok {
				// Recursively sync the fields of each MetricValues element by position.
				for i := range toMetricValues {
					if i < len(fromMetricValues) {
						toMetricValues[i].SyncFieldsDuringCreateOrUpdate(ctx, fromMetricValues[i])
					}
				}
				to.SetMetricValues(ctx, toMetricValues)
			}
		}
	}
}

func (to *RetrieveUserVisibleMetricsResponse) SyncFieldsDuringRead(ctx context.Context, from RetrieveUserVisibleMetricsResponse) {
	if !from.MetricValues.IsNull() && !from.MetricValues.IsUnknown() && to.MetricValues.IsNull() && len(from.MetricValues.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for MetricValues, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.MetricValues = from.MetricValues
	}
	if !from.MetricValues.IsNull() && !from.MetricValues.IsUnknown() {
		if toMetricValues, ok := to.GetMetricValues(ctx); ok {
			if fromMetricValues, ok := from.GetMetricValues(ctx); ok {
				for i := range toMetricValues {
					if i < len(fromMetricValues) {
						toMetricValues[i].SyncFieldsDuringRead(ctx, fromMetricValues[i])
					}
				}
				to.SetMetricValues(ctx, toMetricValues)
			}
		}
	}
}

func (m RetrieveUserVisibleMetricsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["metric_values"] = attrs["metric_values"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RetrieveUserVisibleMetricsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RetrieveUserVisibleMetricsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metric_values": reflect.TypeOf(MetricValues{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RetrieveUserVisibleMetricsResponse
// only implements ToObjectValue() and Type().
func (m RetrieveUserVisibleMetricsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metric_values":   m.MetricValues,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RetrieveUserVisibleMetricsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metric_values": basetypes.ListType{
				ElemType: MetricValues{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetMetricValues returns the value of the MetricValues field in RetrieveUserVisibleMetricsResponse as
// a slice of MetricValues values.
// If the field is unknown or null, the boolean return value is false.
func (m *RetrieveUserVisibleMetricsResponse) GetMetricValues(ctx context.Context) ([]MetricValues, bool) {
	if m.MetricValues.IsNull() || m.MetricValues.IsUnknown() {
		return nil, false
	}
	var v []MetricValues
	d := m.MetricValues.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetricValues sets the value of the MetricValues field in RetrieveUserVisibleMetricsResponse.
func (m *RetrieveUserVisibleMetricsResponse) SetMetricValues(ctx context.Context, v []MetricValues) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["metric_values"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.MetricValues = types.ListValueMust(t, vs)
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
	if !from.Data.IsNull() && !from.Data.IsUnknown() {
		if toData, ok := to.GetData(ctx); ok {
			if fromData, ok := from.GetData(ctx); ok {
				// Recursively sync the fields of each Data element by position.
				for i := range toData {
					if i < len(fromData) {
						toData[i].SyncFieldsDuringCreateOrUpdate(ctx, fromData[i])
					}
				}
				to.SetData(ctx, toData)
			}
		}
	}
}

func (to *ScanVectorIndexResponse) SyncFieldsDuringRead(ctx context.Context, from ScanVectorIndexResponse) {
	if !from.Data.IsNull() && !from.Data.IsUnknown() && to.Data.IsNull() && len(from.Data.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Data, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Data = from.Data
	}
	if !from.Data.IsNull() && !from.Data.IsUnknown() {
		if toData, ok := to.GetData(ctx); ok {
			if fromData, ok := from.GetData(ctx); ok {
				for i := range toData {
					if i < len(fromData) {
						toData[i].SyncFieldsDuringRead(ctx, fromData[i])
					}
				}
				to.SetData(ctx, toData)
			}
		}
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
	if !from.Fields.IsNull() && !from.Fields.IsUnknown() {
		if toFields, ok := to.GetFields(ctx); ok {
			if fromFields, ok := from.GetFields(ctx); ok {
				// Recursively sync the fields of each Fields element by position.
				for i := range toFields {
					if i < len(fromFields) {
						toFields[i].SyncFieldsDuringCreateOrUpdate(ctx, fromFields[i])
					}
				}
				to.SetFields(ctx, toFields)
			}
		}
	}
}

func (to *Struct) SyncFieldsDuringRead(ctx context.Context, from Struct) {
	if !from.Fields.IsNull() && !from.Fields.IsUnknown() && to.Fields.IsNull() && len(from.Fields.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Fields, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Fields = from.Fields
	}
	if !from.Fields.IsNull() && !from.Fields.IsUnknown() {
		if toFields, ok := to.GetFields(ctx); ok {
			if fromFields, ok := from.GetFields(ctx); ok {
				for i := range toFields {
					if i < len(fromFields) {
						toFields[i].SyncFieldsDuringRead(ctx, fromFields[i])
					}
				}
				to.SetFields(ctx, toFields)
			}
		}
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
	// The new custom tags for the AI Search endpoint
	CustomTags types.List `tfsdk:"custom_tags"`
	// Name of the AI Search endpoint
	EndpointName types.String `tfsdk:"-"`
}

func (to *UpdateEndpointCustomTagsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEndpointCustomTagsRequest) {
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() {
		if toCustomTags, ok := to.GetCustomTags(ctx); ok {
			if fromCustomTags, ok := from.GetCustomTags(ctx); ok {
				// Recursively sync the fields of each CustomTags element by position.
				for i := range toCustomTags {
					if i < len(fromCustomTags) {
						toCustomTags[i].SyncFieldsDuringCreateOrUpdate(ctx, fromCustomTags[i])
					}
				}
				to.SetCustomTags(ctx, toCustomTags)
			}
		}
	}
}

func (to *UpdateEndpointCustomTagsRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateEndpointCustomTagsRequest) {
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() {
		if toCustomTags, ok := to.GetCustomTags(ctx); ok {
			if fromCustomTags, ok := from.GetCustomTags(ctx); ok {
				for i := range toCustomTags {
					if i < len(fromCustomTags) {
						toCustomTags[i].SyncFieldsDuringRead(ctx, fromCustomTags[i])
					}
				}
				to.SetCustomTags(ctx, toCustomTags)
			}
		}
	}
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
	// All the custom tags that are applied to the AI Search endpoint.
	CustomTags types.List `tfsdk:"custom_tags"`
	// The name of the AI Search endpoint whose custom tags were updated.
	Name types.String `tfsdk:"name"`
}

func (to *UpdateEndpointCustomTagsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEndpointCustomTagsResponse) {
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() {
		if toCustomTags, ok := to.GetCustomTags(ctx); ok {
			if fromCustomTags, ok := from.GetCustomTags(ctx); ok {
				// Recursively sync the fields of each CustomTags element by position.
				for i := range toCustomTags {
					if i < len(fromCustomTags) {
						toCustomTags[i].SyncFieldsDuringCreateOrUpdate(ctx, fromCustomTags[i])
					}
				}
				to.SetCustomTags(ctx, toCustomTags)
			}
		}
	}
}

func (to *UpdateEndpointCustomTagsResponse) SyncFieldsDuringRead(ctx context.Context, from UpdateEndpointCustomTagsResponse) {
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() {
		if toCustomTags, ok := to.GetCustomTags(ctx); ok {
			if fromCustomTags, ok := from.GetCustomTags(ctx); ok {
				for i := range toCustomTags {
					if i < len(fromCustomTags) {
						toCustomTags[i].SyncFieldsDuringRead(ctx, fromCustomTags[i])
					}
				}
				to.SetCustomTags(ctx, toCustomTags)
			}
		}
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
	// ID of the endpoint associated with the index.
	EndpointId types.String `tfsdk:"endpoint_id"`
	// Name of the endpoint associated with the index
	EndpointName types.String `tfsdk:"endpoint_name"`
	// The subtype of the index.
	IndexSubtype types.String `tfsdk:"index_subtype"`

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
	attrs["endpoint_id"] = attrs["endpoint_id"].SetOptional()
	attrs["endpoint_name"] = attrs["endpoint_name"].SetOptional()
	attrs["index_subtype"] = attrs["index_subtype"].SetOptional()
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
			"endpoint_id":              m.EndpointId,
			"endpoint_name":            m.EndpointName,
			"index_subtype":            m.IndexSubtype,
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
			"endpoint_id":              types.StringType,
			"endpoint_name":            types.StringType,
			"index_subtype":            types.StringType,
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

type VectorSearchEndpointAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *VectorSearchEndpointAccessControlRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from VectorSearchEndpointAccessControlRequest) {
}

func (to *VectorSearchEndpointAccessControlRequest) SyncFieldsDuringRead(ctx context.Context, from VectorSearchEndpointAccessControlRequest) {
}

func (m VectorSearchEndpointAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in VectorSearchEndpointAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m VectorSearchEndpointAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VectorSearchEndpointAccessControlRequest
// only implements ToObjectValue() and Type().
func (m VectorSearchEndpointAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             m.GroupName,
			"permission_level":       m.PermissionLevel,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m VectorSearchEndpointAccessControlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type VectorSearchEndpointAccessControlResponse struct {
	// All permissions.
	AllPermissions types.List `tfsdk:"all_permissions"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name"`
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *VectorSearchEndpointAccessControlResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from VectorSearchEndpointAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() {
		if toAllPermissions, ok := to.GetAllPermissions(ctx); ok {
			if fromAllPermissions, ok := from.GetAllPermissions(ctx); ok {
				// Recursively sync the fields of each AllPermissions element by position.
				for i := range toAllPermissions {
					if i < len(fromAllPermissions) {
						toAllPermissions[i].SyncFieldsDuringCreateOrUpdate(ctx, fromAllPermissions[i])
					}
				}
				to.SetAllPermissions(ctx, toAllPermissions)
			}
		}
	}
}

func (to *VectorSearchEndpointAccessControlResponse) SyncFieldsDuringRead(ctx context.Context, from VectorSearchEndpointAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() {
		if toAllPermissions, ok := to.GetAllPermissions(ctx); ok {
			if fromAllPermissions, ok := from.GetAllPermissions(ctx); ok {
				for i := range toAllPermissions {
					if i < len(fromAllPermissions) {
						toAllPermissions[i].SyncFieldsDuringRead(ctx, fromAllPermissions[i])
					}
				}
				to.SetAllPermissions(ctx, toAllPermissions)
			}
		}
	}
}

func (m VectorSearchEndpointAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["all_permissions"] = attrs["all_permissions"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in VectorSearchEndpointAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m VectorSearchEndpointAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(VectorSearchEndpointPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VectorSearchEndpointAccessControlResponse
// only implements ToObjectValue() and Type().
func (m VectorSearchEndpointAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        m.AllPermissions,
			"display_name":           m.DisplayName,
			"group_name":             m.GroupName,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m VectorSearchEndpointAccessControlResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: VectorSearchEndpointPermission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in VectorSearchEndpointAccessControlResponse as
// a slice of VectorSearchEndpointPermission values.
// If the field is unknown or null, the boolean return value is false.
func (m *VectorSearchEndpointAccessControlResponse) GetAllPermissions(ctx context.Context) ([]VectorSearchEndpointPermission, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []VectorSearchEndpointPermission
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in VectorSearchEndpointAccessControlResponse.
func (m *VectorSearchEndpointAccessControlResponse) SetAllPermissions(ctx context.Context, v []VectorSearchEndpointPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
}

type VectorSearchEndpointPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *VectorSearchEndpointPermission) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from VectorSearchEndpointPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *VectorSearchEndpointPermission) SyncFieldsDuringRead(ctx context.Context, from VectorSearchEndpointPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m VectorSearchEndpointPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited"] = attrs["inherited"].SetOptional()
	attrs["inherited_from_object"] = attrs["inherited_from_object"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in VectorSearchEndpointPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m VectorSearchEndpointPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VectorSearchEndpointPermission
// only implements ToObjectValue() and Type().
func (m VectorSearchEndpointPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m VectorSearchEndpointPermission) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inherited": types.BoolType,
			"inherited_from_object": basetypes.ListType{
				ElemType: types.StringType,
			},
			"permission_level": types.StringType,
		},
	}
}

// GetInheritedFromObject returns the value of the InheritedFromObject field in VectorSearchEndpointPermission as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *VectorSearchEndpointPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if m.InheritedFromObject.IsNull() || m.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in VectorSearchEndpointPermission.
func (m *VectorSearchEndpointPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type VectorSearchEndpointPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *VectorSearchEndpointPermissions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from VectorSearchEndpointPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() {
		if toAccessControlList, ok := to.GetAccessControlList(ctx); ok {
			if fromAccessControlList, ok := from.GetAccessControlList(ctx); ok {
				// Recursively sync the fields of each AccessControlList element by position.
				for i := range toAccessControlList {
					if i < len(fromAccessControlList) {
						toAccessControlList[i].SyncFieldsDuringCreateOrUpdate(ctx, fromAccessControlList[i])
					}
				}
				to.SetAccessControlList(ctx, toAccessControlList)
			}
		}
	}
}

func (to *VectorSearchEndpointPermissions) SyncFieldsDuringRead(ctx context.Context, from VectorSearchEndpointPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() {
		if toAccessControlList, ok := to.GetAccessControlList(ctx); ok {
			if fromAccessControlList, ok := from.GetAccessControlList(ctx); ok {
				for i := range toAccessControlList {
					if i < len(fromAccessControlList) {
						toAccessControlList[i].SyncFieldsDuringRead(ctx, fromAccessControlList[i])
					}
				}
				to.SetAccessControlList(ctx, toAccessControlList)
			}
		}
	}
}

func (m VectorSearchEndpointPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in VectorSearchEndpointPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m VectorSearchEndpointPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(VectorSearchEndpointAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VectorSearchEndpointPermissions
// only implements ToObjectValue() and Type().
func (m VectorSearchEndpointPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m VectorSearchEndpointPermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: VectorSearchEndpointAccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in VectorSearchEndpointPermissions as
// a slice of VectorSearchEndpointAccessControlResponse values.
// If the field is unknown or null, the boolean return value is false.
func (m *VectorSearchEndpointPermissions) GetAccessControlList(ctx context.Context) ([]VectorSearchEndpointAccessControlResponse, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []VectorSearchEndpointAccessControlResponse
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in VectorSearchEndpointPermissions.
func (m *VectorSearchEndpointPermissions) SetAccessControlList(ctx context.Context, v []VectorSearchEndpointAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type VectorSearchEndpointPermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *VectorSearchEndpointPermissionsDescription) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from VectorSearchEndpointPermissionsDescription) {
}

func (to *VectorSearchEndpointPermissionsDescription) SyncFieldsDuringRead(ctx context.Context, from VectorSearchEndpointPermissionsDescription) {
}

func (m VectorSearchEndpointPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in VectorSearchEndpointPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m VectorSearchEndpointPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VectorSearchEndpointPermissionsDescription
// only implements ToObjectValue() and Type().
func (m VectorSearchEndpointPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m VectorSearchEndpointPermissionsDescription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type VectorSearchEndpointPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The vector search endpoint for which to get or manage permissions.
	EndpointId types.String `tfsdk:"-"`
}

func (to *VectorSearchEndpointPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from VectorSearchEndpointPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() {
		if toAccessControlList, ok := to.GetAccessControlList(ctx); ok {
			if fromAccessControlList, ok := from.GetAccessControlList(ctx); ok {
				// Recursively sync the fields of each AccessControlList element by position.
				for i := range toAccessControlList {
					if i < len(fromAccessControlList) {
						toAccessControlList[i].SyncFieldsDuringCreateOrUpdate(ctx, fromAccessControlList[i])
					}
				}
				to.SetAccessControlList(ctx, toAccessControlList)
			}
		}
	}
}

func (to *VectorSearchEndpointPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from VectorSearchEndpointPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() {
		if toAccessControlList, ok := to.GetAccessControlList(ctx); ok {
			if fromAccessControlList, ok := from.GetAccessControlList(ctx); ok {
				for i := range toAccessControlList {
					if i < len(fromAccessControlList) {
						toAccessControlList[i].SyncFieldsDuringRead(ctx, fromAccessControlList[i])
					}
				}
				to.SetAccessControlList(ctx, toAccessControlList)
			}
		}
	}
}

func (m VectorSearchEndpointPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["endpoint_id"] = attrs["endpoint_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in VectorSearchEndpointPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m VectorSearchEndpointPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(VectorSearchEndpointAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, VectorSearchEndpointPermissionsRequest
// only implements ToObjectValue() and Type().
func (m VectorSearchEndpointPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"endpoint_id":         m.EndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m VectorSearchEndpointPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: VectorSearchEndpointAccessControlRequest{}.Type(ctx),
			},
			"endpoint_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in VectorSearchEndpointPermissionsRequest as
// a slice of VectorSearchEndpointAccessControlRequest values.
// If the field is unknown or null, the boolean return value is false.
func (m *VectorSearchEndpointPermissionsRequest) GetAccessControlList(ctx context.Context) ([]VectorSearchEndpointAccessControlRequest, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []VectorSearchEndpointAccessControlRequest
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in VectorSearchEndpointPermissionsRequest.
func (m *VectorSearchEndpointPermissionsRequest) SetAccessControlList(ctx context.Context, v []VectorSearchEndpointAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}
