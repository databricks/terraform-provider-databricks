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

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CancelCustomLlmOptimizationRunRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (toState *CancelCustomLlmOptimizationRunRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CancelCustomLlmOptimizationRunRequest_SdkV2) {
}

func (toState *CancelCustomLlmOptimizationRunRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CancelCustomLlmOptimizationRunRequest_SdkV2) {
}

func (c CancelCustomLlmOptimizationRunRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelCustomLlmOptimizationRunRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CancelCustomLlmOptimizationRunRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelCustomLlmOptimizationRunRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CancelCustomLlmOptimizationRunRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CancelCustomLlmOptimizationRunRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type CreateCustomLlmRequest_SdkV2 struct {
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

func (toState *CreateCustomLlmRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateCustomLlmRequest_SdkV2) {
}

func (toState *CreateCustomLlmRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateCustomLlmRequest_SdkV2) {
}

func (c CreateCustomLlmRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["agent_artifact_path"] = attrs["agent_artifact_path"].SetOptional()
	attrs["datasets"] = attrs["datasets"].SetOptional()
	attrs["guidelines"] = attrs["guidelines"].SetOptional()
	attrs["instructions"] = attrs["instructions"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCustomLlmRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCustomLlmRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"datasets":   reflect.TypeOf(Dataset_SdkV2{}),
		"guidelines": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCustomLlmRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCustomLlmRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateCustomLlmRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"agent_artifact_path": types.StringType,
			"datasets": basetypes.ListType{
				ElemType: Dataset_SdkV2{}.Type(ctx),
			},
			"guidelines": basetypes.ListType{
				ElemType: types.StringType,
			},
			"instructions": types.StringType,
			"name":         types.StringType,
		},
	}
}

// GetDatasets returns the value of the Datasets field in CreateCustomLlmRequest_SdkV2 as
// a slice of Dataset_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomLlmRequest_SdkV2) GetDatasets(ctx context.Context) ([]Dataset_SdkV2, bool) {
	if o.Datasets.IsNull() || o.Datasets.IsUnknown() {
		return nil, false
	}
	var v []Dataset_SdkV2
	d := o.Datasets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatasets sets the value of the Datasets field in CreateCustomLlmRequest_SdkV2.
func (o *CreateCustomLlmRequest_SdkV2) SetDatasets(ctx context.Context, v []Dataset_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["datasets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Datasets = types.ListValueMust(t, vs)
}

// GetGuidelines returns the value of the Guidelines field in CreateCustomLlmRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomLlmRequest_SdkV2) GetGuidelines(ctx context.Context) ([]types.String, bool) {
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

// SetGuidelines sets the value of the Guidelines field in CreateCustomLlmRequest_SdkV2.
func (o *CreateCustomLlmRequest_SdkV2) SetGuidelines(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["guidelines"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Guidelines = types.ListValueMust(t, vs)
}

type CustomLlm_SdkV2 struct {
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

func (toState *CustomLlm_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CustomLlm_SdkV2) {
}

func (toState *CustomLlm_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CustomLlm_SdkV2) {
}

func (c CustomLlm_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CustomLlm_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"datasets":   reflect.TypeOf(Dataset_SdkV2{}),
		"guidelines": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomLlm_SdkV2
// only implements ToObjectValue() and Type().
func (o CustomLlm_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CustomLlm_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"agent_artifact_path": types.StringType,
			"creation_time":       types.StringType,
			"creator":             types.StringType,
			"datasets": basetypes.ListType{
				ElemType: Dataset_SdkV2{}.Type(ctx),
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

// GetDatasets returns the value of the Datasets field in CustomLlm_SdkV2 as
// a slice of Dataset_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomLlm_SdkV2) GetDatasets(ctx context.Context) ([]Dataset_SdkV2, bool) {
	if o.Datasets.IsNull() || o.Datasets.IsUnknown() {
		return nil, false
	}
	var v []Dataset_SdkV2
	d := o.Datasets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatasets sets the value of the Datasets field in CustomLlm_SdkV2.
func (o *CustomLlm_SdkV2) SetDatasets(ctx context.Context, v []Dataset_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["datasets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Datasets = types.ListValueMust(t, vs)
}

// GetGuidelines returns the value of the Guidelines field in CustomLlm_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomLlm_SdkV2) GetGuidelines(ctx context.Context) ([]types.String, bool) {
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

// SetGuidelines sets the value of the Guidelines field in CustomLlm_SdkV2.
func (o *CustomLlm_SdkV2) SetGuidelines(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["guidelines"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Guidelines = types.ListValueMust(t, vs)
}

type Dataset_SdkV2 struct {
	Table types.List `tfsdk:"table"`
}

func (toState *Dataset_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Dataset_SdkV2) {
	if !fromPlan.Table.IsNull() && !fromPlan.Table.IsUnknown() {
		if toStateTable, ok := toState.GetTable(ctx); ok {
			if fromPlanTable, ok := fromPlan.GetTable(ctx); ok {
				toStateTable.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanTable)
				toState.SetTable(ctx, toStateTable)
			}
		}
	}
}

func (toState *Dataset_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Dataset_SdkV2) {
	if !fromState.Table.IsNull() && !fromState.Table.IsUnknown() {
		if toStateTable, ok := toState.GetTable(ctx); ok {
			if fromStateTable, ok := fromState.GetTable(ctx); ok {
				toStateTable.SyncFieldsDuringRead(ctx, fromStateTable)
				toState.SetTable(ctx, toStateTable)
			}
		}
	}
}

func (c Dataset_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["table"] = attrs["table"].SetRequired()
	attrs["table"] = attrs["table"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Dataset.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Dataset_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table": reflect.TypeOf(Table_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Dataset_SdkV2
// only implements ToObjectValue() and Type().
func (o Dataset_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table": o.Table,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Dataset_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table": basetypes.ListType{
				ElemType: Table_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTable returns the value of the Table field in Dataset_SdkV2 as
// a Table_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Dataset_SdkV2) GetTable(ctx context.Context) (Table_SdkV2, bool) {
	var e Table_SdkV2
	if o.Table.IsNull() || o.Table.IsUnknown() {
		return e, false
	}
	var v []Table_SdkV2
	d := o.Table.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTable sets the value of the Table field in Dataset_SdkV2.
func (o *Dataset_SdkV2) SetTable(ctx context.Context, v Table_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["table"]
	o.Table = types.ListValueMust(t, vs)
}

type DeleteCustomLlmRequest_SdkV2 struct {
	// The id of the custom llm
	Id types.String `tfsdk:"-"`
}

func (toState *DeleteCustomLlmRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteCustomLlmRequest_SdkV2) {
}

func (toState *DeleteCustomLlmRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteCustomLlmRequest_SdkV2) {
}

func (c DeleteCustomLlmRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCustomLlmRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCustomLlmRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCustomLlmRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteCustomLlmRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCustomLlmRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetCustomLlmRequest_SdkV2 struct {
	// The id of the custom llm
	Id types.String `tfsdk:"-"`
}

func (toState *GetCustomLlmRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetCustomLlmRequest_SdkV2) {
}

func (toState *GetCustomLlmRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetCustomLlmRequest_SdkV2) {
}

func (c GetCustomLlmRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCustomLlmRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCustomLlmRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCustomLlmRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetCustomLlmRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCustomLlmRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type StartCustomLlmOptimizationRunRequest_SdkV2 struct {
	// The Id of the tile.
	Id types.String `tfsdk:"-"`
}

func (toState *StartCustomLlmOptimizationRunRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan StartCustomLlmOptimizationRunRequest_SdkV2) {
}

func (toState *StartCustomLlmOptimizationRunRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState StartCustomLlmOptimizationRunRequest_SdkV2) {
}

func (c StartCustomLlmOptimizationRunRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StartCustomLlmOptimizationRunRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a StartCustomLlmOptimizationRunRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartCustomLlmOptimizationRunRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o StartCustomLlmOptimizationRunRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StartCustomLlmOptimizationRunRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type Table_SdkV2 struct {
	// Name of the request column
	RequestCol types.String `tfsdk:"request_col"`
	// Optional: Name of the response column if the data is labeled
	ResponseCol types.String `tfsdk:"response_col"`
	// Full UC table path in catalog.schema.table_name format
	TablePath types.String `tfsdk:"table_path"`
}

func (toState *Table_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Table_SdkV2) {
}

func (toState *Table_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Table_SdkV2) {
}

func (c Table_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Table_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Table_SdkV2
// only implements ToObjectValue() and Type().
func (o Table_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"request_col":  o.RequestCol,
			"response_col": o.ResponseCol,
			"table_path":   o.TablePath,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Table_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"request_col":  types.StringType,
			"response_col": types.StringType,
			"table_path":   types.StringType,
		},
	}
}

type UpdateCustomLlmRequest_SdkV2 struct {
	// The CustomLlm containing the fields which should be updated.
	CustomLlm types.List `tfsdk:"custom_llm"`
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

func (toState *UpdateCustomLlmRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateCustomLlmRequest_SdkV2) {
	if !fromPlan.CustomLlm.IsNull() && !fromPlan.CustomLlm.IsUnknown() {
		if toStateCustomLlm, ok := toState.GetCustomLlm(ctx); ok {
			if fromPlanCustomLlm, ok := fromPlan.GetCustomLlm(ctx); ok {
				toStateCustomLlm.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanCustomLlm)
				toState.SetCustomLlm(ctx, toStateCustomLlm)
			}
		}
	}
}

func (toState *UpdateCustomLlmRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UpdateCustomLlmRequest_SdkV2) {
	if !fromState.CustomLlm.IsNull() && !fromState.CustomLlm.IsUnknown() {
		if toStateCustomLlm, ok := toState.GetCustomLlm(ctx); ok {
			if fromStateCustomLlm, ok := fromState.GetCustomLlm(ctx); ok {
				toStateCustomLlm.SyncFieldsDuringRead(ctx, fromStateCustomLlm)
				toState.SetCustomLlm(ctx, toStateCustomLlm)
			}
		}
	}
}

func (c UpdateCustomLlmRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["custom_llm"] = attrs["custom_llm"].SetRequired()
	attrs["custom_llm"] = attrs["custom_llm"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["update_mask"] = attrs["update_mask"].SetRequired()
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCustomLlmRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCustomLlmRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_llm": reflect.TypeOf(CustomLlm_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCustomLlmRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateCustomLlmRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"custom_llm":  o.CustomLlm,
			"id":          o.Id,
			"update_mask": o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCustomLlmRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"custom_llm": basetypes.ListType{
				ElemType: CustomLlm_SdkV2{}.Type(ctx),
			},
			"id":          types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetCustomLlm returns the value of the CustomLlm field in UpdateCustomLlmRequest_SdkV2 as
// a CustomLlm_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCustomLlmRequest_SdkV2) GetCustomLlm(ctx context.Context) (CustomLlm_SdkV2, bool) {
	var e CustomLlm_SdkV2
	if o.CustomLlm.IsNull() || o.CustomLlm.IsUnknown() {
		return e, false
	}
	var v []CustomLlm_SdkV2
	d := o.CustomLlm.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCustomLlm sets the value of the CustomLlm field in UpdateCustomLlmRequest_SdkV2.
func (o *UpdateCustomLlmRequest_SdkV2) SetCustomLlm(ctx context.Context, v CustomLlm_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_llm"]
	o.CustomLlm = types.ListValueMust(t, vs)
}
