// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package agentbricks_tf

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

type CancelCustomLlmOptimizationRunRequest struct {
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelCustomLlmOptimizationRunRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelCustomLlmOptimizationRunRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelCustomLlmOptimizationRunRequest
// only implements ToObjectValue() and Type().
func (o CancelCustomLlmOptimizationRunRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CancelCustomLlmOptimizationRunRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type CreateCustomLlmRequest struct {
	// This will soon be deprecated!! Optional: UC path for agent artifacts. If
	// you are using a dataset that you only have read permissions, please
	// provide a destination path where you have write permissions. Please
	// provide this in catalog.schema format.
	AgentArtifactPath types.String `tfsdk:"agent_artifact_path"`
	// Datasets used for training and evaluating the model, not for inference.
	// Currently, only 1 dataset is accepted.
	Datasets types.List `tfsdk:"datasets"`
	// Guidelines for the custom LLM to adhere to
	Guidelines types.List `tfsdk:"guidelines"`
	// Instructions for the custom LLM to follow
	Instructions types.String `tfsdk:"instructions"`
	// Name of the custom LLM. Only alphanumeric characters and dashes allowed.
	Name types.String `tfsdk:"name"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCustomLlmRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCustomLlmRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"datasets":   reflect.TypeOf(Dataset{}),
		"guidelines": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCustomLlmRequest
// only implements ToObjectValue() and Type().
func (o CreateCustomLlmRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"agent_artifact_path": o.AgentArtifactPath,
			"datasets":            o.Datasets,
			"guidelines":          o.Guidelines,
			"instructions":        o.Instructions,
			"name":                o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCustomLlmRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"agent_artifact_path": types.StringType,
			"datasets": basetypes.ListType{
				ElemType: Dataset{}.Type(ctx),
			},
			"guidelines": basetypes.ListType{
				ElemType: types.StringType,
			},
			"instructions": types.StringType,
			"name":         types.StringType,
		},
	}
}

// GetDatasets returns the value of the Datasets field in CreateCustomLlmRequest as
// a slice of Dataset values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomLlmRequest) GetDatasets(ctx context.Context) ([]Dataset, bool) {
	if o.Datasets.IsNull() || o.Datasets.IsUnknown() {
		return nil, false
	}
	var v []Dataset
	d := o.Datasets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatasets sets the value of the Datasets field in CreateCustomLlmRequest.
func (o *CreateCustomLlmRequest) SetDatasets(ctx context.Context, v []Dataset) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["datasets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Datasets = types.ListValueMust(t, vs)
}

// GetGuidelines returns the value of the Guidelines field in CreateCustomLlmRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomLlmRequest) GetGuidelines(ctx context.Context) ([]types.String, bool) {
	if o.Guidelines.IsNull() || o.Guidelines.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Guidelines.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGuidelines sets the value of the Guidelines field in CreateCustomLlmRequest.
func (o *CreateCustomLlmRequest) SetGuidelines(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["guidelines"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Guidelines = types.ListValueMust(t, vs)
}

type CustomLlm struct {
	AgentArtifactPath types.String `tfsdk:"agent_artifact_path"`
	// Creation timestamp of the custom LLM
	CreationTime types.String `tfsdk:"creation_time"`
	// Creator of the custom LLM
	Creator types.String `tfsdk:"creator"`
	// Datasets used for training and evaluating the model, not for inference
	Datasets types.List `tfsdk:"datasets"`
	// Name of the endpoint that will be used to serve the custom LLM
	EndpointName types.String `tfsdk:"endpoint_name"`
	// Guidelines for the custom LLM to adhere to
	Guidelines types.List `tfsdk:"guidelines"`

	Id types.String `tfsdk:"id"`
	// Instructions for the custom LLM to follow
	Instructions types.String `tfsdk:"instructions"`
	// Name of the custom LLM
	Name types.String `tfsdk:"name"`
	// If optimization is kicked off, tracks the state of the custom LLM
	OptimizationState types.String `tfsdk:"optimization_state"`
}

func (toState *CustomLlm) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CustomLlm) {
}

func (toState *CustomLlm) SyncFieldsDuringRead(ctx context.Context, fromState CustomLlm) {
}

func (c CustomLlm) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["agent_artifact_path"] = attrs["agent_artifact_path"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["creation_time"] = attrs["creation_time"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["creator"] = attrs["creator"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["datasets"] = attrs["datasets"].SetOptional()
	attrs["endpoint_name"] = attrs["endpoint_name"].SetComputed()
	attrs["guidelines"] = attrs["guidelines"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["instructions"] = attrs["instructions"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["optimization_state"] = attrs["optimization_state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CustomLlm.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CustomLlm) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"datasets":   reflect.TypeOf(Dataset{}),
		"guidelines": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomLlm
// only implements ToObjectValue() and Type().
func (o CustomLlm) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"agent_artifact_path": o.AgentArtifactPath,
			"creation_time":       o.CreationTime,
			"creator":             o.Creator,
			"datasets":            o.Datasets,
			"endpoint_name":       o.EndpointName,
			"guidelines":          o.Guidelines,
			"id":                  o.Id,
			"instructions":        o.Instructions,
			"name":                o.Name,
			"optimization_state":  o.OptimizationState,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CustomLlm) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"agent_artifact_path": types.StringType,
			"creation_time":       types.StringType,
			"creator":             types.StringType,
			"datasets": basetypes.ListType{
				ElemType: Dataset{}.Type(ctx),
			},
			"endpoint_name": types.StringType,
			"guidelines": basetypes.ListType{
				ElemType: types.StringType,
			},
			"id":                 types.StringType,
			"instructions":       types.StringType,
			"name":               types.StringType,
			"optimization_state": types.StringType,
		},
	}
}

// GetDatasets returns the value of the Datasets field in CustomLlm as
// a slice of Dataset values.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomLlm) GetDatasets(ctx context.Context) ([]Dataset, bool) {
	if o.Datasets.IsNull() || o.Datasets.IsUnknown() {
		return nil, false
	}
	var v []Dataset
	d := o.Datasets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatasets sets the value of the Datasets field in CustomLlm.
func (o *CustomLlm) SetDatasets(ctx context.Context, v []Dataset) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["datasets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Datasets = types.ListValueMust(t, vs)
}

// GetGuidelines returns the value of the Guidelines field in CustomLlm as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomLlm) GetGuidelines(ctx context.Context) ([]types.String, bool) {
	if o.Guidelines.IsNull() || o.Guidelines.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Guidelines.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGuidelines sets the value of the Guidelines field in CustomLlm.
func (o *CustomLlm) SetGuidelines(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["guidelines"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Guidelines = types.ListValueMust(t, vs)
}

type Dataset struct {
	Table types.Object `tfsdk:"table"`
}

func (toState *Dataset) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Dataset) {
	if !fromPlan.Table.IsNull() && !fromPlan.Table.IsUnknown() {
		if toStateTable, ok := toState.GetTable(ctx); ok {
			if fromPlanTable, ok := fromPlan.GetTable(ctx); ok {
				toStateTable.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanTable)
				toState.SetTable(ctx, toStateTable)
			}
		}
	}
}

func (toState *Dataset) SyncFieldsDuringRead(ctx context.Context, fromState Dataset) {
	if !fromState.Table.IsNull() && !fromState.Table.IsUnknown() {
		if toStateTable, ok := toState.GetTable(ctx); ok {
			if fromStateTable, ok := fromState.GetTable(ctx); ok {
				toStateTable.SyncFieldsDuringRead(ctx, fromStateTable)
				toState.SetTable(ctx, toStateTable)
			}
		}
	}
}

func (c Dataset) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["table"] = attrs["table"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Dataset.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Dataset) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table": reflect.TypeOf(Table{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Dataset
// only implements ToObjectValue() and Type().
func (o Dataset) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table": o.Table,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Dataset) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table": Table{}.Type(ctx),
		},
	}
}

// GetTable returns the value of the Table field in Dataset as
// a Table value.
// If the field is unknown or null, the boolean return value is false.
func (o *Dataset) GetTable(ctx context.Context) (Table, bool) {
	var e Table
	if o.Table.IsNull() || o.Table.IsUnknown() {
		return e, false
	}
	var v Table
	d := o.Table.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTable sets the value of the Table field in Dataset.
func (o *Dataset) SetTable(ctx context.Context, v Table) {
	vs := v.ToObjectValue(ctx)
	o.Table = vs
}

type DeleteCustomLlmRequest struct {
	// The id of the custom llm
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCustomLlmRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCustomLlmRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCustomLlmRequest
// only implements ToObjectValue() and Type().
func (o DeleteCustomLlmRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCustomLlmRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetCustomLlmRequest struct {
	// The id of the custom llm
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCustomLlmRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCustomLlmRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCustomLlmRequest
// only implements ToObjectValue() and Type().
func (o GetCustomLlmRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCustomLlmRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type StartCustomLlmOptimizationRunRequest struct {
	// The Id of the tile.
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StartCustomLlmOptimizationRunRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StartCustomLlmOptimizationRunRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartCustomLlmOptimizationRunRequest
// only implements ToObjectValue() and Type().
func (o StartCustomLlmOptimizationRunRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StartCustomLlmOptimizationRunRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type Table struct {
	// Name of the request column
	RequestCol types.String `tfsdk:"request_col"`
	// Optional: Name of the response column if the data is labeled
	ResponseCol types.String `tfsdk:"response_col"`
	// Full UC table path in catalog.schema.table_name format
	TablePath types.String `tfsdk:"table_path"`
}

func (toState *Table) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Table) {
}

func (toState *Table) SyncFieldsDuringRead(ctx context.Context, fromState Table) {
}

func (c Table) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["request_col"] = attrs["request_col"].SetRequired()
	attrs["response_col"] = attrs["response_col"].SetOptional()
	attrs["table_path"] = attrs["table_path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Table.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Table) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Table
// only implements ToObjectValue() and Type().
func (o Table) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"request_col":  o.RequestCol,
			"response_col": o.ResponseCol,
			"table_path":   o.TablePath,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Table) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"request_col":  types.StringType,
			"response_col": types.StringType,
			"table_path":   types.StringType,
		},
	}
}

type UpdateCustomLlmRequest struct {
	// The CustomLlm containing the fields which should be updated.
	CustomLlm types.Object `tfsdk:"custom_llm"`
	// The id of the custom llm
	Id types.String `tfsdk:"-"`
	// The list of the CustomLlm fields to update. These should correspond to
	// the values (or lack thereof) present in `custom_llm`.
	//
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
	UpdateMask types.String `tfsdk:"update_mask"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCustomLlmRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCustomLlmRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_llm": reflect.TypeOf(CustomLlm{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCustomLlmRequest
// only implements ToObjectValue() and Type().
func (o UpdateCustomLlmRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"custom_llm":  o.CustomLlm,
			"id":          o.Id,
			"update_mask": o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCustomLlmRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"custom_llm":  CustomLlm{}.Type(ctx),
			"id":          types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetCustomLlm returns the value of the CustomLlm field in UpdateCustomLlmRequest as
// a CustomLlm value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCustomLlmRequest) GetCustomLlm(ctx context.Context) (CustomLlm, bool) {
	var e CustomLlm
	if o.CustomLlm.IsNull() || o.CustomLlm.IsUnknown() {
		return e, false
	}
	var v CustomLlm
	d := o.CustomLlm.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomLlm sets the value of the CustomLlm field in UpdateCustomLlmRequest.
func (o *UpdateCustomLlmRequest) SetCustomLlm(ctx context.Context, v CustomLlm) {
	vs := v.ToObjectValue(ctx)
	o.CustomLlm = vs
}
