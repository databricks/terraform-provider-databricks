// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package knowledgeassistants_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateExampleRequest_SdkV2 struct {
	// The example to create under the parent Knowledge Assistant.
	Example types.List `tfsdk:"example"`
	// Parent resource where this example will be created. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Parent types.String `tfsdk:"-"`
}

func (to *CreateExampleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateExampleRequest_SdkV2) {
	if !from.Example.IsNull() && !from.Example.IsUnknown() {
		if toExample, ok := to.GetExample(ctx); ok {
			if fromExample, ok := from.GetExample(ctx); ok {
				// Recursively sync the fields of Example
				toExample.SyncFieldsDuringCreateOrUpdate(ctx, fromExample)
				to.SetExample(ctx, toExample)
			}
		}
	}
}

func (to *CreateExampleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateExampleRequest_SdkV2) {
	if !from.Example.IsNull() && !from.Example.IsUnknown() {
		if toExample, ok := to.GetExample(ctx); ok {
			if fromExample, ok := from.GetExample(ctx); ok {
				toExample.SyncFieldsDuringRead(ctx, fromExample)
				to.SetExample(ctx, toExample)
			}
		}
	}
}

func (m CreateExampleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["example"] = attrs["example"].SetRequired()
	attrs["example"] = attrs["example"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateExampleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateExampleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"example": reflect.TypeOf(Example_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExampleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateExampleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"example": m.Example,
			"parent":  m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateExampleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"example": basetypes.ListType{
				ElemType: Example_SdkV2{}.Type(ctx),
			},
			"parent": types.StringType,
		},
	}
}

// GetExample returns the value of the Example field in CreateExampleRequest_SdkV2 as
// a Example_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateExampleRequest_SdkV2) GetExample(ctx context.Context) (Example_SdkV2, bool) {
	var e Example_SdkV2
	if m.Example.IsNull() || m.Example.IsUnknown() {
		return e, false
	}
	var v []Example_SdkV2
	d := m.Example.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExample sets the value of the Example field in CreateExampleRequest_SdkV2.
func (m *CreateExampleRequest_SdkV2) SetExample(ctx context.Context, v Example_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["example"]
	m.Example = types.ListValueMust(t, vs)
}

type CreateKnowledgeAssistantRequest_SdkV2 struct {
	// The Knowledge Assistant to create.
	KnowledgeAssistant types.List `tfsdk:"knowledge_assistant"`
}

func (to *CreateKnowledgeAssistantRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateKnowledgeAssistantRequest_SdkV2) {
	if !from.KnowledgeAssistant.IsNull() && !from.KnowledgeAssistant.IsUnknown() {
		if toKnowledgeAssistant, ok := to.GetKnowledgeAssistant(ctx); ok {
			if fromKnowledgeAssistant, ok := from.GetKnowledgeAssistant(ctx); ok {
				// Recursively sync the fields of KnowledgeAssistant
				toKnowledgeAssistant.SyncFieldsDuringCreateOrUpdate(ctx, fromKnowledgeAssistant)
				to.SetKnowledgeAssistant(ctx, toKnowledgeAssistant)
			}
		}
	}
}

func (to *CreateKnowledgeAssistantRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateKnowledgeAssistantRequest_SdkV2) {
	if !from.KnowledgeAssistant.IsNull() && !from.KnowledgeAssistant.IsUnknown() {
		if toKnowledgeAssistant, ok := to.GetKnowledgeAssistant(ctx); ok {
			if fromKnowledgeAssistant, ok := from.GetKnowledgeAssistant(ctx); ok {
				toKnowledgeAssistant.SyncFieldsDuringRead(ctx, fromKnowledgeAssistant)
				to.SetKnowledgeAssistant(ctx, toKnowledgeAssistant)
			}
		}
	}
}

func (m CreateKnowledgeAssistantRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["knowledge_assistant"] = attrs["knowledge_assistant"].SetRequired()
	attrs["knowledge_assistant"] = attrs["knowledge_assistant"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateKnowledgeAssistantRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateKnowledgeAssistantRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"knowledge_assistant": reflect.TypeOf(KnowledgeAssistant_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateKnowledgeAssistantRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateKnowledgeAssistantRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"knowledge_assistant": m.KnowledgeAssistant,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateKnowledgeAssistantRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"knowledge_assistant": basetypes.ListType{
				ElemType: KnowledgeAssistant_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetKnowledgeAssistant returns the value of the KnowledgeAssistant field in CreateKnowledgeAssistantRequest_SdkV2 as
// a KnowledgeAssistant_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateKnowledgeAssistantRequest_SdkV2) GetKnowledgeAssistant(ctx context.Context) (KnowledgeAssistant_SdkV2, bool) {
	var e KnowledgeAssistant_SdkV2
	if m.KnowledgeAssistant.IsNull() || m.KnowledgeAssistant.IsUnknown() {
		return e, false
	}
	var v []KnowledgeAssistant_SdkV2
	d := m.KnowledgeAssistant.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetKnowledgeAssistant sets the value of the KnowledgeAssistant field in CreateKnowledgeAssistantRequest_SdkV2.
func (m *CreateKnowledgeAssistantRequest_SdkV2) SetKnowledgeAssistant(ctx context.Context, v KnowledgeAssistant_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["knowledge_assistant"]
	m.KnowledgeAssistant = types.ListValueMust(t, vs)
}

type CreateKnowledgeSourceRequest_SdkV2 struct {
	KnowledgeSource types.List `tfsdk:"knowledge_source"`
	// Parent resource where this source will be created. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Parent types.String `tfsdk:"-"`
}

func (to *CreateKnowledgeSourceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateKnowledgeSourceRequest_SdkV2) {
	if !from.KnowledgeSource.IsNull() && !from.KnowledgeSource.IsUnknown() {
		if toKnowledgeSource, ok := to.GetKnowledgeSource(ctx); ok {
			if fromKnowledgeSource, ok := from.GetKnowledgeSource(ctx); ok {
				// Recursively sync the fields of KnowledgeSource
				toKnowledgeSource.SyncFieldsDuringCreateOrUpdate(ctx, fromKnowledgeSource)
				to.SetKnowledgeSource(ctx, toKnowledgeSource)
			}
		}
	}
}

func (to *CreateKnowledgeSourceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateKnowledgeSourceRequest_SdkV2) {
	if !from.KnowledgeSource.IsNull() && !from.KnowledgeSource.IsUnknown() {
		if toKnowledgeSource, ok := to.GetKnowledgeSource(ctx); ok {
			if fromKnowledgeSource, ok := from.GetKnowledgeSource(ctx); ok {
				toKnowledgeSource.SyncFieldsDuringRead(ctx, fromKnowledgeSource)
				to.SetKnowledgeSource(ctx, toKnowledgeSource)
			}
		}
	}
}

func (m CreateKnowledgeSourceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["knowledge_source"] = attrs["knowledge_source"].SetRequired()
	attrs["knowledge_source"] = attrs["knowledge_source"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateKnowledgeSourceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateKnowledgeSourceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"knowledge_source": reflect.TypeOf(KnowledgeSource_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateKnowledgeSourceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateKnowledgeSourceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"knowledge_source": m.KnowledgeSource,
			"parent":           m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateKnowledgeSourceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"knowledge_source": basetypes.ListType{
				ElemType: KnowledgeSource_SdkV2{}.Type(ctx),
			},
			"parent": types.StringType,
		},
	}
}

// GetKnowledgeSource returns the value of the KnowledgeSource field in CreateKnowledgeSourceRequest_SdkV2 as
// a KnowledgeSource_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateKnowledgeSourceRequest_SdkV2) GetKnowledgeSource(ctx context.Context) (KnowledgeSource_SdkV2, bool) {
	var e KnowledgeSource_SdkV2
	if m.KnowledgeSource.IsNull() || m.KnowledgeSource.IsUnknown() {
		return e, false
	}
	var v []KnowledgeSource_SdkV2
	d := m.KnowledgeSource.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetKnowledgeSource sets the value of the KnowledgeSource field in CreateKnowledgeSourceRequest_SdkV2.
func (m *CreateKnowledgeSourceRequest_SdkV2) SetKnowledgeSource(ctx context.Context, v KnowledgeSource_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["knowledge_source"]
	m.KnowledgeSource = types.ListValueMust(t, vs)
}

type DeleteExampleRequest_SdkV2 struct {
	// The resource name of the example to delete. Format:
	// knowledge-assistants/{knowledge_assistant_id}/examples/{example_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteExampleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteExampleRequest_SdkV2) {
}

func (to *DeleteExampleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteExampleRequest_SdkV2) {
}

func (m DeleteExampleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteExampleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteExampleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExampleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteExampleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteExampleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteKnowledgeAssistantRequest_SdkV2 struct {
	// The resource name of the knowledge assistant to be deleted. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteKnowledgeAssistantRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteKnowledgeAssistantRequest_SdkV2) {
}

func (to *DeleteKnowledgeAssistantRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteKnowledgeAssistantRequest_SdkV2) {
}

func (m DeleteKnowledgeAssistantRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteKnowledgeAssistantRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteKnowledgeAssistantRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteKnowledgeAssistantRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteKnowledgeAssistantRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteKnowledgeAssistantRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteKnowledgeSourceRequest_SdkV2 struct {
	// The resource name of the Knowledge Source to delete. Format:
	// knowledge-assistants/{knowledge_assistant_id}/knowledge-sources/{knowledge_source_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteKnowledgeSourceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteKnowledgeSourceRequest_SdkV2) {
}

func (to *DeleteKnowledgeSourceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteKnowledgeSourceRequest_SdkV2) {
}

func (m DeleteKnowledgeSourceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteKnowledgeSourceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteKnowledgeSourceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteKnowledgeSourceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteKnowledgeSourceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteKnowledgeSourceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// An example associated with a Knowledge Assistant. Contains a question and
// guidelines for how the assistant should respond.
type Example_SdkV2 struct {
	// Timestamp when this example was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The universally unique identifier (UUID) of the example.
	ExampleId types.String `tfsdk:"example_id"`
	// Guidelines for answering the question.
	Guidelines types.List `tfsdk:"guidelines"`
	// Full resource name:
	// knowledge-assistants/{knowledge_assistant_id}/examples/{example_id}
	Name types.String `tfsdk:"name"`
	// The example question.
	Question types.String `tfsdk:"question"`
	// Timestamp when this example was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
}

func (to *Example_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Example_SdkV2) {
}

func (to *Example_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Example_SdkV2) {
}

func (m Example_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["example_id"] = attrs["example_id"].SetComputed()
	attrs["guidelines"] = attrs["guidelines"].SetRequired()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["question"] = attrs["question"].SetRequired()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Example.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Example_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"guidelines": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Example_SdkV2
// only implements ToObjectValue() and Type().
func (m Example_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time": m.CreateTime,
			"example_id":  m.ExampleId,
			"guidelines":  m.Guidelines,
			"name":        m.Name,
			"question":    m.Question,
			"update_time": m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Example_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": timetypes.RFC3339{}.Type(ctx),
			"example_id":  types.StringType,
			"guidelines": basetypes.ListType{
				ElemType: types.StringType,
			},
			"name":        types.StringType,
			"question":    types.StringType,
			"update_time": timetypes.RFC3339{}.Type(ctx),
		},
	}
}

// GetGuidelines returns the value of the Guidelines field in Example_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Example_SdkV2) GetGuidelines(ctx context.Context) ([]types.String, bool) {
	if m.Guidelines.IsNull() || m.Guidelines.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Guidelines.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGuidelines sets the value of the Guidelines field in Example_SdkV2.
func (m *Example_SdkV2) SetGuidelines(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["guidelines"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Guidelines = types.ListValueMust(t, vs)
}

// FileTableSpec specifies a file table source configuration.
type FileTableSpec_SdkV2 struct {
	// The name of the column containing BINARY file content to be indexed.
	FileCol types.String `tfsdk:"file_col"`
	// Full UC name of the table, in the format of
	// {CATALOG}.{SCHEMA}.{TABLE_NAME}.
	TableName types.String `tfsdk:"table_name"`
}

func (to *FileTableSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FileTableSpec_SdkV2) {
}

func (to *FileTableSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FileTableSpec_SdkV2) {
}

func (m FileTableSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file_col"] = attrs["file_col"].SetRequired()
	attrs["table_name"] = attrs["table_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FileTableSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FileTableSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FileTableSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m FileTableSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_col":   m.FileCol,
			"table_name": m.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FileTableSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_col":   types.StringType,
			"table_name": types.StringType,
		},
	}
}

// FilesSpec specifies a files source configuration.
type FilesSpec_SdkV2 struct {
	// A UC volume path that includes a list of files.
	Path types.String `tfsdk:"path"`
}

func (to *FilesSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FilesSpec_SdkV2) {
}

func (to *FilesSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FilesSpec_SdkV2) {
}

func (m FilesSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path"] = attrs["path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FilesSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FilesSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FilesSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m FilesSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path": m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FilesSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path": types.StringType,
		},
	}
}

type GetExampleRequest_SdkV2 struct {
	// The resource name of the example. Format:
	// knowledge-assistants/{knowledge_assistant_id}/examples/{example_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetExampleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExampleRequest_SdkV2) {
}

func (to *GetExampleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetExampleRequest_SdkV2) {
}

func (m GetExampleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExampleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetExampleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExampleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetExampleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExampleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetKnowledgeAssistantPermissionLevelsRequest_SdkV2 struct {
	// The knowledge assistant for which to get or manage permissions.
	KnowledgeAssistantId types.String `tfsdk:"-"`
}

func (to *GetKnowledgeAssistantPermissionLevelsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetKnowledgeAssistantPermissionLevelsRequest_SdkV2) {
}

func (to *GetKnowledgeAssistantPermissionLevelsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetKnowledgeAssistantPermissionLevelsRequest_SdkV2) {
}

func (m GetKnowledgeAssistantPermissionLevelsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["knowledge_assistant_id"] = attrs["knowledge_assistant_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetKnowledgeAssistantPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetKnowledgeAssistantPermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetKnowledgeAssistantPermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetKnowledgeAssistantPermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"knowledge_assistant_id": m.KnowledgeAssistantId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetKnowledgeAssistantPermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"knowledge_assistant_id": types.StringType,
		},
	}
}

type GetKnowledgeAssistantPermissionLevelsResponse_SdkV2 struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (to *GetKnowledgeAssistantPermissionLevelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetKnowledgeAssistantPermissionLevelsResponse_SdkV2) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetKnowledgeAssistantPermissionLevelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetKnowledgeAssistantPermissionLevelsResponse_SdkV2) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (m GetKnowledgeAssistantPermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_levels"] = attrs["permission_levels"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetKnowledgeAssistantPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetKnowledgeAssistantPermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(KnowledgeAssistantPermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetKnowledgeAssistantPermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetKnowledgeAssistantPermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetKnowledgeAssistantPermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: KnowledgeAssistantPermissionsDescription_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetKnowledgeAssistantPermissionLevelsResponse_SdkV2 as
// a slice of KnowledgeAssistantPermissionsDescription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetKnowledgeAssistantPermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]KnowledgeAssistantPermissionsDescription_SdkV2, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []KnowledgeAssistantPermissionsDescription_SdkV2
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetKnowledgeAssistantPermissionLevelsResponse_SdkV2.
func (m *GetKnowledgeAssistantPermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []KnowledgeAssistantPermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetKnowledgeAssistantPermissionsRequest_SdkV2 struct {
	// The knowledge assistant for which to get or manage permissions.
	KnowledgeAssistantId types.String `tfsdk:"-"`
}

func (to *GetKnowledgeAssistantPermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetKnowledgeAssistantPermissionsRequest_SdkV2) {
}

func (to *GetKnowledgeAssistantPermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetKnowledgeAssistantPermissionsRequest_SdkV2) {
}

func (m GetKnowledgeAssistantPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["knowledge_assistant_id"] = attrs["knowledge_assistant_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetKnowledgeAssistantPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetKnowledgeAssistantPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetKnowledgeAssistantPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetKnowledgeAssistantPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"knowledge_assistant_id": m.KnowledgeAssistantId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetKnowledgeAssistantPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"knowledge_assistant_id": types.StringType,
		},
	}
}

type GetKnowledgeAssistantRequest_SdkV2 struct {
	// The resource name of the knowledge assistant. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetKnowledgeAssistantRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetKnowledgeAssistantRequest_SdkV2) {
}

func (to *GetKnowledgeAssistantRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetKnowledgeAssistantRequest_SdkV2) {
}

func (m GetKnowledgeAssistantRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetKnowledgeAssistantRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetKnowledgeAssistantRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetKnowledgeAssistantRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetKnowledgeAssistantRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetKnowledgeAssistantRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetKnowledgeSourceRequest_SdkV2 struct {
	// The resource name of the Knowledge Source. Format:
	// knowledge-assistants/{knowledge_assistant_id}/knowledge-sources/{knowledge_source_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetKnowledgeSourceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetKnowledgeSourceRequest_SdkV2) {
}

func (to *GetKnowledgeSourceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetKnowledgeSourceRequest_SdkV2) {
}

func (m GetKnowledgeSourceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetKnowledgeSourceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetKnowledgeSourceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetKnowledgeSourceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetKnowledgeSourceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetKnowledgeSourceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// IndexSpec specifies a vector search index source configuration.
type IndexSpec_SdkV2 struct {
	// The column that specifies a link or reference to where the information
	// came from.
	DocUriCol types.String `tfsdk:"doc_uri_col"`
	// Full UC name of the vector search index, in the format of
	// {CATALOG}.{SCHEMA}.{INDEX_NAME}.
	IndexName types.String `tfsdk:"index_name"`
	// The column that includes the document text for retrieval.
	TextCol types.String `tfsdk:"text_col"`
}

func (to *IndexSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from IndexSpec_SdkV2) {
}

func (to *IndexSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from IndexSpec_SdkV2) {
}

func (m IndexSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["doc_uri_col"] = attrs["doc_uri_col"].SetRequired()
	attrs["index_name"] = attrs["index_name"].SetRequired()
	attrs["text_col"] = attrs["text_col"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in IndexSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m IndexSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IndexSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m IndexSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"doc_uri_col": m.DocUriCol,
			"index_name":  m.IndexName,
			"text_col":    m.TextCol,
		})
}

// Type implements basetypes.ObjectValuable.
func (m IndexSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"doc_uri_col": types.StringType,
			"index_name":  types.StringType,
			"text_col":    types.StringType,
		},
	}
}

// Entity message that represents a knowledge assistant. Note: REQUIRED
// annotations below represent create-time requirements. For updates, required
// fields are determined by the update mask.
type KnowledgeAssistant_SdkV2 struct {
	// Creation timestamp.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The creator of the Knowledge Assistant.
	Creator types.String `tfsdk:"creator"`
	// Description of what this agent can do (user-facing). Required when
	// creating a Knowledge Assistant. When updating a Knowledge Assistant,
	// optional unless included in update_mask.
	Description types.String `tfsdk:"description"`
	// The display name of the Knowledge Assistant, unique at workspace level.
	// Required when creating a Knowledge Assistant. When updating a Knowledge
	// Assistant, optional unless included in update_mask.
	DisplayName types.String `tfsdk:"display_name"`
	// The name of the knowledge assistant agent endpoint.
	EndpointName types.String `tfsdk:"endpoint_name"`
	// Error details when the Knowledge Assistant is in FAILED state.
	ErrorInfo types.String `tfsdk:"error_info"`
	// The MLflow experiment ID.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// The universally unique identifier (UUID) of the Knowledge Assistant.
	Id types.String `tfsdk:"id"`
	// Additional global instructions on how the agent should generate answers.
	// Optional on create and update. When updating a Knowledge Assistant,
	// include this field in update_mask to modify it.
	Instructions types.String `tfsdk:"instructions"`
	// The resource name of the Knowledge Assistant. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name types.String `tfsdk:"name"`
	// State of the Knowledge Assistant. Not returned in List responses.
	State types.String `tfsdk:"state"`
}

func (to *KnowledgeAssistant_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KnowledgeAssistant_SdkV2) {
}

func (to *KnowledgeAssistant_SdkV2) SyncFieldsDuringRead(ctx context.Context, from KnowledgeAssistant_SdkV2) {
}

func (m KnowledgeAssistant_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["description"] = attrs["description"].SetRequired()
	attrs["display_name"] = attrs["display_name"].SetRequired()
	attrs["endpoint_name"] = attrs["endpoint_name"].SetComputed()
	attrs["error_info"] = attrs["error_info"].SetComputed()
	attrs["experiment_id"] = attrs["experiment_id"].SetComputed()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["instructions"] = attrs["instructions"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["state"] = attrs["state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in KnowledgeAssistant.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m KnowledgeAssistant_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KnowledgeAssistant_SdkV2
// only implements ToObjectValue() and Type().
func (m KnowledgeAssistant_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":   m.CreateTime,
			"creator":       m.Creator,
			"description":   m.Description,
			"display_name":  m.DisplayName,
			"endpoint_name": m.EndpointName,
			"error_info":    m.ErrorInfo,
			"experiment_id": m.ExperimentId,
			"id":            m.Id,
			"instructions":  m.Instructions,
			"name":          m.Name,
			"state":         m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m KnowledgeAssistant_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":   timetypes.RFC3339{}.Type(ctx),
			"creator":       types.StringType,
			"description":   types.StringType,
			"display_name":  types.StringType,
			"endpoint_name": types.StringType,
			"error_info":    types.StringType,
			"experiment_id": types.StringType,
			"id":            types.StringType,
			"instructions":  types.StringType,
			"name":          types.StringType,
			"state":         types.StringType,
		},
	}
}

type KnowledgeAssistantAccessControlRequest_SdkV2 struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *KnowledgeAssistantAccessControlRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KnowledgeAssistantAccessControlRequest_SdkV2) {
}

func (to *KnowledgeAssistantAccessControlRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from KnowledgeAssistantAccessControlRequest_SdkV2) {
}

func (m KnowledgeAssistantAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in KnowledgeAssistantAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m KnowledgeAssistantAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KnowledgeAssistantAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m KnowledgeAssistantAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m KnowledgeAssistantAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type KnowledgeAssistantAccessControlResponse_SdkV2 struct {
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

func (to *KnowledgeAssistantAccessControlResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KnowledgeAssistantAccessControlResponse_SdkV2) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *KnowledgeAssistantAccessControlResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from KnowledgeAssistantAccessControlResponse_SdkV2) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (m KnowledgeAssistantAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["all_permissions"] = attrs["all_permissions"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in KnowledgeAssistantAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m KnowledgeAssistantAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(KnowledgeAssistantPermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KnowledgeAssistantAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m KnowledgeAssistantAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m KnowledgeAssistantAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: KnowledgeAssistantPermission_SdkV2{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in KnowledgeAssistantAccessControlResponse_SdkV2 as
// a slice of KnowledgeAssistantPermission_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *KnowledgeAssistantAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]KnowledgeAssistantPermission_SdkV2, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []KnowledgeAssistantPermission_SdkV2
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in KnowledgeAssistantAccessControlResponse_SdkV2.
func (m *KnowledgeAssistantAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []KnowledgeAssistantPermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
}

type KnowledgeAssistantPermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *KnowledgeAssistantPermission_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KnowledgeAssistantPermission_SdkV2) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *KnowledgeAssistantPermission_SdkV2) SyncFieldsDuringRead(ctx context.Context, from KnowledgeAssistantPermission_SdkV2) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m KnowledgeAssistantPermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited"] = attrs["inherited"].SetOptional()
	attrs["inherited_from_object"] = attrs["inherited_from_object"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in KnowledgeAssistantPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m KnowledgeAssistantPermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KnowledgeAssistantPermission_SdkV2
// only implements ToObjectValue() and Type().
func (m KnowledgeAssistantPermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m KnowledgeAssistantPermission_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in KnowledgeAssistantPermission_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *KnowledgeAssistantPermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in KnowledgeAssistantPermission_SdkV2.
func (m *KnowledgeAssistantPermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type KnowledgeAssistantPermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *KnowledgeAssistantPermissions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KnowledgeAssistantPermissions_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *KnowledgeAssistantPermissions_SdkV2) SyncFieldsDuringRead(ctx context.Context, from KnowledgeAssistantPermissions_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m KnowledgeAssistantPermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in KnowledgeAssistantPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m KnowledgeAssistantPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(KnowledgeAssistantAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KnowledgeAssistantPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (m KnowledgeAssistantPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m KnowledgeAssistantPermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: KnowledgeAssistantAccessControlResponse_SdkV2{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in KnowledgeAssistantPermissions_SdkV2 as
// a slice of KnowledgeAssistantAccessControlResponse_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *KnowledgeAssistantPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]KnowledgeAssistantAccessControlResponse_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []KnowledgeAssistantAccessControlResponse_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in KnowledgeAssistantPermissions_SdkV2.
func (m *KnowledgeAssistantPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []KnowledgeAssistantAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type KnowledgeAssistantPermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *KnowledgeAssistantPermissionsDescription_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KnowledgeAssistantPermissionsDescription_SdkV2) {
}

func (to *KnowledgeAssistantPermissionsDescription_SdkV2) SyncFieldsDuringRead(ctx context.Context, from KnowledgeAssistantPermissionsDescription_SdkV2) {
}

func (m KnowledgeAssistantPermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in KnowledgeAssistantPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m KnowledgeAssistantPermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KnowledgeAssistantPermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (m KnowledgeAssistantPermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m KnowledgeAssistantPermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type KnowledgeAssistantPermissionsRequest_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The knowledge assistant for which to get or manage permissions.
	KnowledgeAssistantId types.String `tfsdk:"-"`
}

func (to *KnowledgeAssistantPermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KnowledgeAssistantPermissionsRequest_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *KnowledgeAssistantPermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from KnowledgeAssistantPermissionsRequest_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m KnowledgeAssistantPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["knowledge_assistant_id"] = attrs["knowledge_assistant_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in KnowledgeAssistantPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m KnowledgeAssistantPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(KnowledgeAssistantAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KnowledgeAssistantPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m KnowledgeAssistantPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list":    m.AccessControlList,
			"knowledge_assistant_id": m.KnowledgeAssistantId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m KnowledgeAssistantPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: KnowledgeAssistantAccessControlRequest_SdkV2{}.Type(ctx),
			},
			"knowledge_assistant_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in KnowledgeAssistantPermissionsRequest_SdkV2 as
// a slice of KnowledgeAssistantAccessControlRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *KnowledgeAssistantPermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]KnowledgeAssistantAccessControlRequest_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []KnowledgeAssistantAccessControlRequest_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in KnowledgeAssistantPermissionsRequest_SdkV2.
func (m *KnowledgeAssistantPermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []KnowledgeAssistantAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

// KnowledgeSource represents a source of knowledge for the KnowledgeAssistant.
// Used in create/update requests and returned in Get/List responses. Note:
// REQUIRED annotations below represent create-time requirements. For updates,
// required fields are determined by the update mask.
type KnowledgeSource_SdkV2 struct {
	// Timestamp when this knowledge source was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Description of the knowledge source. Required when creating a Knowledge
	// Source. When updating a Knowledge Source, optional unless included in
	// update_mask.
	Description types.String `tfsdk:"description"`
	// Human-readable display name of the knowledge source. Required when
	// creating a Knowledge Source. When updating a Knowledge Source, optional
	// unless included in update_mask.
	DisplayName types.String `tfsdk:"display_name"`

	FileTable types.List `tfsdk:"file_table"`

	Files types.List `tfsdk:"files"`

	Id types.String `tfsdk:"id"`

	Index types.List `tfsdk:"index"`
	// Timestamp representing the cutoff before which content in this knowledge
	// source is being ingested.
	KnowledgeCutoffTime timetypes.RFC3339 `tfsdk:"knowledge_cutoff_time"`
	// Full resource name:
	// knowledge-assistants/{knowledge_assistant_id}/knowledge-sources/{knowledge_source_id}
	Name types.String `tfsdk:"name"`
	// The type of the source: "index", "files", or "file_table". Required when
	// creating a Knowledge Source. When updating a Knowledge Source, this field
	// is ignored.
	SourceType types.String `tfsdk:"source_type"`

	State types.String `tfsdk:"state"`
}

func (to *KnowledgeSource_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KnowledgeSource_SdkV2) {
	if !from.FileTable.IsNull() && !from.FileTable.IsUnknown() {
		if toFileTable, ok := to.GetFileTable(ctx); ok {
			if fromFileTable, ok := from.GetFileTable(ctx); ok {
				// Recursively sync the fields of FileTable
				toFileTable.SyncFieldsDuringCreateOrUpdate(ctx, fromFileTable)
				to.SetFileTable(ctx, toFileTable)
			}
		}
	}
	if !from.Files.IsNull() && !from.Files.IsUnknown() {
		if toFiles, ok := to.GetFiles(ctx); ok {
			if fromFiles, ok := from.GetFiles(ctx); ok {
				// Recursively sync the fields of Files
				toFiles.SyncFieldsDuringCreateOrUpdate(ctx, fromFiles)
				to.SetFiles(ctx, toFiles)
			}
		}
	}
	if !from.Index.IsNull() && !from.Index.IsUnknown() {
		if toIndex, ok := to.GetIndex(ctx); ok {
			if fromIndex, ok := from.GetIndex(ctx); ok {
				// Recursively sync the fields of Index
				toIndex.SyncFieldsDuringCreateOrUpdate(ctx, fromIndex)
				to.SetIndex(ctx, toIndex)
			}
		}
	}
}

func (to *KnowledgeSource_SdkV2) SyncFieldsDuringRead(ctx context.Context, from KnowledgeSource_SdkV2) {
	if !from.FileTable.IsNull() && !from.FileTable.IsUnknown() {
		if toFileTable, ok := to.GetFileTable(ctx); ok {
			if fromFileTable, ok := from.GetFileTable(ctx); ok {
				toFileTable.SyncFieldsDuringRead(ctx, fromFileTable)
				to.SetFileTable(ctx, toFileTable)
			}
		}
	}
	if !from.Files.IsNull() && !from.Files.IsUnknown() {
		if toFiles, ok := to.GetFiles(ctx); ok {
			if fromFiles, ok := from.GetFiles(ctx); ok {
				toFiles.SyncFieldsDuringRead(ctx, fromFiles)
				to.SetFiles(ctx, toFiles)
			}
		}
	}
	if !from.Index.IsNull() && !from.Index.IsUnknown() {
		if toIndex, ok := to.GetIndex(ctx); ok {
			if fromIndex, ok := from.GetIndex(ctx); ok {
				toIndex.SyncFieldsDuringRead(ctx, fromIndex)
				to.SetIndex(ctx, toIndex)
			}
		}
	}
}

func (m KnowledgeSource_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["description"] = attrs["description"].SetRequired()
	attrs["display_name"] = attrs["display_name"].SetRequired()
	attrs["file_table"] = attrs["file_table"].SetOptional()
	attrs["file_table"] = attrs["file_table"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["files"] = attrs["files"].SetOptional()
	attrs["files"] = attrs["files"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetComputed()
	attrs["index"] = attrs["index"].SetOptional()
	attrs["index"] = attrs["index"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["knowledge_cutoff_time"] = attrs["knowledge_cutoff_time"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["source_type"] = attrs["source_type"].SetRequired()
	attrs["state"] = attrs["state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in KnowledgeSource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m KnowledgeSource_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"file_table": reflect.TypeOf(FileTableSpec_SdkV2{}),
		"files":      reflect.TypeOf(FilesSpec_SdkV2{}),
		"index":      reflect.TypeOf(IndexSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KnowledgeSource_SdkV2
// only implements ToObjectValue() and Type().
func (m KnowledgeSource_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":           m.CreateTime,
			"description":           m.Description,
			"display_name":          m.DisplayName,
			"file_table":            m.FileTable,
			"files":                 m.Files,
			"id":                    m.Id,
			"index":                 m.Index,
			"knowledge_cutoff_time": m.KnowledgeCutoffTime,
			"name":                  m.Name,
			"source_type":           m.SourceType,
			"state":                 m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m KnowledgeSource_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":  timetypes.RFC3339{}.Type(ctx),
			"description":  types.StringType,
			"display_name": types.StringType,
			"file_table": basetypes.ListType{
				ElemType: FileTableSpec_SdkV2{}.Type(ctx),
			},
			"files": basetypes.ListType{
				ElemType: FilesSpec_SdkV2{}.Type(ctx),
			},
			"id": types.StringType,
			"index": basetypes.ListType{
				ElemType: IndexSpec_SdkV2{}.Type(ctx),
			},
			"knowledge_cutoff_time": timetypes.RFC3339{}.Type(ctx),
			"name":                  types.StringType,
			"source_type":           types.StringType,
			"state":                 types.StringType,
		},
	}
}

// GetFileTable returns the value of the FileTable field in KnowledgeSource_SdkV2 as
// a FileTableSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *KnowledgeSource_SdkV2) GetFileTable(ctx context.Context) (FileTableSpec_SdkV2, bool) {
	var e FileTableSpec_SdkV2
	if m.FileTable.IsNull() || m.FileTable.IsUnknown() {
		return e, false
	}
	var v []FileTableSpec_SdkV2
	d := m.FileTable.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFileTable sets the value of the FileTable field in KnowledgeSource_SdkV2.
func (m *KnowledgeSource_SdkV2) SetFileTable(ctx context.Context, v FileTableSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["file_table"]
	m.FileTable = types.ListValueMust(t, vs)
}

// GetFiles returns the value of the Files field in KnowledgeSource_SdkV2 as
// a FilesSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *KnowledgeSource_SdkV2) GetFiles(ctx context.Context) (FilesSpec_SdkV2, bool) {
	var e FilesSpec_SdkV2
	if m.Files.IsNull() || m.Files.IsUnknown() {
		return e, false
	}
	var v []FilesSpec_SdkV2
	d := m.Files.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFiles sets the value of the Files field in KnowledgeSource_SdkV2.
func (m *KnowledgeSource_SdkV2) SetFiles(ctx context.Context, v FilesSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["files"]
	m.Files = types.ListValueMust(t, vs)
}

// GetIndex returns the value of the Index field in KnowledgeSource_SdkV2 as
// a IndexSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *KnowledgeSource_SdkV2) GetIndex(ctx context.Context) (IndexSpec_SdkV2, bool) {
	var e IndexSpec_SdkV2
	if m.Index.IsNull() || m.Index.IsUnknown() {
		return e, false
	}
	var v []IndexSpec_SdkV2
	d := m.Index.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIndex sets the value of the Index field in KnowledgeSource_SdkV2.
func (m *KnowledgeSource_SdkV2) SetIndex(ctx context.Context, v IndexSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["index"]
	m.Index = types.ListValueMust(t, vs)
}

type ListExamplesRequest_SdkV2 struct {
	// The maximum number of examples to return. If unspecified, at most 100
	// examples will be returned. The maximum value is 100; values above 100
	// will be coerced to 100.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListExamples` call. Provide this
	// to retrieve the subsequent page. If unspecified, the first page will be
	// returned.
	PageToken types.String `tfsdk:"-"`
	// Parent resource to list from. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Parent types.String `tfsdk:"-"`
}

func (to *ListExamplesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListExamplesRequest_SdkV2) {
}

func (to *ListExamplesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListExamplesRequest_SdkV2) {
}

func (m ListExamplesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExamplesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListExamplesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExamplesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListExamplesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListExamplesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// A list of Knowledge Assistant examples.
type ListExamplesResponse_SdkV2 struct {
	Examples types.List `tfsdk:"examples"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListExamplesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListExamplesResponse_SdkV2) {
	if !from.Examples.IsNull() && !from.Examples.IsUnknown() && to.Examples.IsNull() && len(from.Examples.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Examples, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Examples = from.Examples
	}
}

func (to *ListExamplesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListExamplesResponse_SdkV2) {
	if !from.Examples.IsNull() && !from.Examples.IsUnknown() && to.Examples.IsNull() && len(from.Examples.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Examples, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Examples = from.Examples
	}
}

func (m ListExamplesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["examples"] = attrs["examples"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExamplesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListExamplesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"examples": reflect.TypeOf(Example_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExamplesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListExamplesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"examples":        m.Examples,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListExamplesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"examples": basetypes.ListType{
				ElemType: Example_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetExamples returns the value of the Examples field in ListExamplesResponse_SdkV2 as
// a slice of Example_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListExamplesResponse_SdkV2) GetExamples(ctx context.Context) ([]Example_SdkV2, bool) {
	if m.Examples.IsNull() || m.Examples.IsUnknown() {
		return nil, false
	}
	var v []Example_SdkV2
	d := m.Examples.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExamples sets the value of the Examples field in ListExamplesResponse_SdkV2.
func (m *ListExamplesResponse_SdkV2) SetExamples(ctx context.Context, v []Example_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["examples"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Examples = types.ListValueMust(t, vs)
}

type ListKnowledgeAssistantsRequest_SdkV2 struct {
	// The maximum number of knowledge assistants to return. If unspecified, at
	// most 100 knowledge assistants will be returned. The maximum value is 100;
	// values above 100 will be coerced to 100.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListKnowledgeAssistants` call.
	// Provide this to retrieve the subsequent page. If unspecified, the first
	// page will be returned.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListKnowledgeAssistantsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListKnowledgeAssistantsRequest_SdkV2) {
}

func (to *ListKnowledgeAssistantsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListKnowledgeAssistantsRequest_SdkV2) {
}

func (m ListKnowledgeAssistantsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListKnowledgeAssistantsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListKnowledgeAssistantsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListKnowledgeAssistantsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListKnowledgeAssistantsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListKnowledgeAssistantsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// A list of Knowledge Assistants.
type ListKnowledgeAssistantsResponse_SdkV2 struct {
	KnowledgeAssistants types.List `tfsdk:"knowledge_assistants"`
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListKnowledgeAssistantsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListKnowledgeAssistantsResponse_SdkV2) {
	if !from.KnowledgeAssistants.IsNull() && !from.KnowledgeAssistants.IsUnknown() && to.KnowledgeAssistants.IsNull() && len(from.KnowledgeAssistants.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for KnowledgeAssistants, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.KnowledgeAssistants = from.KnowledgeAssistants
	}
}

func (to *ListKnowledgeAssistantsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListKnowledgeAssistantsResponse_SdkV2) {
	if !from.KnowledgeAssistants.IsNull() && !from.KnowledgeAssistants.IsUnknown() && to.KnowledgeAssistants.IsNull() && len(from.KnowledgeAssistants.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for KnowledgeAssistants, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.KnowledgeAssistants = from.KnowledgeAssistants
	}
}

func (m ListKnowledgeAssistantsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["knowledge_assistants"] = attrs["knowledge_assistants"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListKnowledgeAssistantsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListKnowledgeAssistantsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"knowledge_assistants": reflect.TypeOf(KnowledgeAssistant_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListKnowledgeAssistantsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListKnowledgeAssistantsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"knowledge_assistants": m.KnowledgeAssistants,
			"next_page_token":      m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListKnowledgeAssistantsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"knowledge_assistants": basetypes.ListType{
				ElemType: KnowledgeAssistant_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetKnowledgeAssistants returns the value of the KnowledgeAssistants field in ListKnowledgeAssistantsResponse_SdkV2 as
// a slice of KnowledgeAssistant_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListKnowledgeAssistantsResponse_SdkV2) GetKnowledgeAssistants(ctx context.Context) ([]KnowledgeAssistant_SdkV2, bool) {
	if m.KnowledgeAssistants.IsNull() || m.KnowledgeAssistants.IsUnknown() {
		return nil, false
	}
	var v []KnowledgeAssistant_SdkV2
	d := m.KnowledgeAssistants.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetKnowledgeAssistants sets the value of the KnowledgeAssistants field in ListKnowledgeAssistantsResponse_SdkV2.
func (m *ListKnowledgeAssistantsResponse_SdkV2) SetKnowledgeAssistants(ctx context.Context, v []KnowledgeAssistant_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["knowledge_assistants"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.KnowledgeAssistants = types.ListValueMust(t, vs)
}

type ListKnowledgeSourcesRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
	// Parent resource to list from. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Parent types.String `tfsdk:"-"`
}

func (to *ListKnowledgeSourcesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListKnowledgeSourcesRequest_SdkV2) {
}

func (to *ListKnowledgeSourcesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListKnowledgeSourcesRequest_SdkV2) {
}

func (m ListKnowledgeSourcesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListKnowledgeSourcesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListKnowledgeSourcesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListKnowledgeSourcesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListKnowledgeSourcesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListKnowledgeSourcesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListKnowledgeSourcesResponse_SdkV2 struct {
	KnowledgeSources types.List `tfsdk:"knowledge_sources"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListKnowledgeSourcesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListKnowledgeSourcesResponse_SdkV2) {
	if !from.KnowledgeSources.IsNull() && !from.KnowledgeSources.IsUnknown() && to.KnowledgeSources.IsNull() && len(from.KnowledgeSources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for KnowledgeSources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.KnowledgeSources = from.KnowledgeSources
	}
}

func (to *ListKnowledgeSourcesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListKnowledgeSourcesResponse_SdkV2) {
	if !from.KnowledgeSources.IsNull() && !from.KnowledgeSources.IsUnknown() && to.KnowledgeSources.IsNull() && len(from.KnowledgeSources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for KnowledgeSources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.KnowledgeSources = from.KnowledgeSources
	}
}

func (m ListKnowledgeSourcesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["knowledge_sources"] = attrs["knowledge_sources"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListKnowledgeSourcesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListKnowledgeSourcesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"knowledge_sources": reflect.TypeOf(KnowledgeSource_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListKnowledgeSourcesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListKnowledgeSourcesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"knowledge_sources": m.KnowledgeSources,
			"next_page_token":   m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListKnowledgeSourcesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"knowledge_sources": basetypes.ListType{
				ElemType: KnowledgeSource_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetKnowledgeSources returns the value of the KnowledgeSources field in ListKnowledgeSourcesResponse_SdkV2 as
// a slice of KnowledgeSource_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListKnowledgeSourcesResponse_SdkV2) GetKnowledgeSources(ctx context.Context) ([]KnowledgeSource_SdkV2, bool) {
	if m.KnowledgeSources.IsNull() || m.KnowledgeSources.IsUnknown() {
		return nil, false
	}
	var v []KnowledgeSource_SdkV2
	d := m.KnowledgeSources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetKnowledgeSources sets the value of the KnowledgeSources field in ListKnowledgeSourcesResponse_SdkV2.
func (m *ListKnowledgeSourcesResponse_SdkV2) SetKnowledgeSources(ctx context.Context, v []KnowledgeSource_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["knowledge_sources"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.KnowledgeSources = types.ListValueMust(t, vs)
}

type SyncKnowledgeSourcesRequest_SdkV2 struct {
	// The resource name of the Knowledge Assistant. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name types.String `tfsdk:"-"`
}

func (to *SyncKnowledgeSourcesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncKnowledgeSourcesRequest_SdkV2) {
}

func (to *SyncKnowledgeSourcesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SyncKnowledgeSourcesRequest_SdkV2) {
}

func (m SyncKnowledgeSourcesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncKnowledgeSourcesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SyncKnowledgeSourcesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncKnowledgeSourcesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m SyncKnowledgeSourcesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SyncKnowledgeSourcesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type UpdateExampleRequest_SdkV2 struct {
	Example types.List `tfsdk:"example"`
	// The resource name of the example to update. Format:
	// knowledge-assistants/{knowledge_assistant_id}/examples/{example_id}
	Name types.String `tfsdk:"-"`
	// Comma-delimited list of fields to update on the example. Allowed values:
	// `question`, `guidelines`. Examples: - `question` - `question,guidelines`
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateExampleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateExampleRequest_SdkV2) {
	if !from.Example.IsNull() && !from.Example.IsUnknown() {
		if toExample, ok := to.GetExample(ctx); ok {
			if fromExample, ok := from.GetExample(ctx); ok {
				// Recursively sync the fields of Example
				toExample.SyncFieldsDuringCreateOrUpdate(ctx, fromExample)
				to.SetExample(ctx, toExample)
			}
		}
	}
}

func (to *UpdateExampleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateExampleRequest_SdkV2) {
	if !from.Example.IsNull() && !from.Example.IsUnknown() {
		if toExample, ok := to.GetExample(ctx); ok {
			if fromExample, ok := from.GetExample(ctx); ok {
				toExample.SyncFieldsDuringRead(ctx, fromExample)
				to.SetExample(ctx, toExample)
			}
		}
	}
}

func (m UpdateExampleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["example"] = attrs["example"].SetRequired()
	attrs["example"] = attrs["example"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateExampleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateExampleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"example": reflect.TypeOf(Example_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExampleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateExampleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"example":     m.Example,
			"name":        m.Name,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateExampleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"example": basetypes.ListType{
				ElemType: Example_SdkV2{}.Type(ctx),
			},
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetExample returns the value of the Example field in UpdateExampleRequest_SdkV2 as
// a Example_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateExampleRequest_SdkV2) GetExample(ctx context.Context) (Example_SdkV2, bool) {
	var e Example_SdkV2
	if m.Example.IsNull() || m.Example.IsUnknown() {
		return e, false
	}
	var v []Example_SdkV2
	d := m.Example.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExample sets the value of the Example field in UpdateExampleRequest_SdkV2.
func (m *UpdateExampleRequest_SdkV2) SetExample(ctx context.Context, v Example_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["example"]
	m.Example = types.ListValueMust(t, vs)
}

type UpdateKnowledgeAssistantRequest_SdkV2 struct {
	// The Knowledge Assistant update payload. Only fields listed in update_mask
	// are updated. REQUIRED annotations on Knowledge Assistant fields describe
	// create-time requirements and do not mean all those fields are required
	// for update.
	KnowledgeAssistant types.List `tfsdk:"knowledge_assistant"`
	// The resource name of the Knowledge Assistant. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name types.String `tfsdk:"-"`
	// Comma-delimited list of fields to update on the Knowledge Assistant.
	// Allowed values: `display_name`, `description`, `instructions`. Examples:
	// - `display_name` - `description,instructions`
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateKnowledgeAssistantRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateKnowledgeAssistantRequest_SdkV2) {
	if !from.KnowledgeAssistant.IsNull() && !from.KnowledgeAssistant.IsUnknown() {
		if toKnowledgeAssistant, ok := to.GetKnowledgeAssistant(ctx); ok {
			if fromKnowledgeAssistant, ok := from.GetKnowledgeAssistant(ctx); ok {
				// Recursively sync the fields of KnowledgeAssistant
				toKnowledgeAssistant.SyncFieldsDuringCreateOrUpdate(ctx, fromKnowledgeAssistant)
				to.SetKnowledgeAssistant(ctx, toKnowledgeAssistant)
			}
		}
	}
}

func (to *UpdateKnowledgeAssistantRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateKnowledgeAssistantRequest_SdkV2) {
	if !from.KnowledgeAssistant.IsNull() && !from.KnowledgeAssistant.IsUnknown() {
		if toKnowledgeAssistant, ok := to.GetKnowledgeAssistant(ctx); ok {
			if fromKnowledgeAssistant, ok := from.GetKnowledgeAssistant(ctx); ok {
				toKnowledgeAssistant.SyncFieldsDuringRead(ctx, fromKnowledgeAssistant)
				to.SetKnowledgeAssistant(ctx, toKnowledgeAssistant)
			}
		}
	}
}

func (m UpdateKnowledgeAssistantRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["knowledge_assistant"] = attrs["knowledge_assistant"].SetRequired()
	attrs["knowledge_assistant"] = attrs["knowledge_assistant"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateKnowledgeAssistantRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateKnowledgeAssistantRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"knowledge_assistant": reflect.TypeOf(KnowledgeAssistant_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateKnowledgeAssistantRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateKnowledgeAssistantRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"knowledge_assistant": m.KnowledgeAssistant,
			"name":                m.Name,
			"update_mask":         m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateKnowledgeAssistantRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"knowledge_assistant": basetypes.ListType{
				ElemType: KnowledgeAssistant_SdkV2{}.Type(ctx),
			},
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetKnowledgeAssistant returns the value of the KnowledgeAssistant field in UpdateKnowledgeAssistantRequest_SdkV2 as
// a KnowledgeAssistant_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateKnowledgeAssistantRequest_SdkV2) GetKnowledgeAssistant(ctx context.Context) (KnowledgeAssistant_SdkV2, bool) {
	var e KnowledgeAssistant_SdkV2
	if m.KnowledgeAssistant.IsNull() || m.KnowledgeAssistant.IsUnknown() {
		return e, false
	}
	var v []KnowledgeAssistant_SdkV2
	d := m.KnowledgeAssistant.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetKnowledgeAssistant sets the value of the KnowledgeAssistant field in UpdateKnowledgeAssistantRequest_SdkV2.
func (m *UpdateKnowledgeAssistantRequest_SdkV2) SetKnowledgeAssistant(ctx context.Context, v KnowledgeAssistant_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["knowledge_assistant"]
	m.KnowledgeAssistant = types.ListValueMust(t, vs)
}

type UpdateKnowledgeSourceRequest_SdkV2 struct {
	// The Knowledge Source update payload. Only fields listed in update_mask
	// are updated. REQUIRED annotations on Knowledge Source fields describe
	// create-time requirements and do not mean all those fields are required
	// for update.
	KnowledgeSource types.List `tfsdk:"knowledge_source"`
	// The resource name of the Knowledge Source to update. Format:
	// knowledge-assistants/{knowledge_assistant_id}/knowledge-sources/{knowledge_source_id}
	Name types.String `tfsdk:"-"`
	// Comma-delimited list of fields to update on the Knowledge Source. Allowed
	// values: `display_name`, `description`. Examples: - `display_name` -
	// `display_name,description`
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateKnowledgeSourceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateKnowledgeSourceRequest_SdkV2) {
	if !from.KnowledgeSource.IsNull() && !from.KnowledgeSource.IsUnknown() {
		if toKnowledgeSource, ok := to.GetKnowledgeSource(ctx); ok {
			if fromKnowledgeSource, ok := from.GetKnowledgeSource(ctx); ok {
				// Recursively sync the fields of KnowledgeSource
				toKnowledgeSource.SyncFieldsDuringCreateOrUpdate(ctx, fromKnowledgeSource)
				to.SetKnowledgeSource(ctx, toKnowledgeSource)
			}
		}
	}
}

func (to *UpdateKnowledgeSourceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateKnowledgeSourceRequest_SdkV2) {
	if !from.KnowledgeSource.IsNull() && !from.KnowledgeSource.IsUnknown() {
		if toKnowledgeSource, ok := to.GetKnowledgeSource(ctx); ok {
			if fromKnowledgeSource, ok := from.GetKnowledgeSource(ctx); ok {
				toKnowledgeSource.SyncFieldsDuringRead(ctx, fromKnowledgeSource)
				to.SetKnowledgeSource(ctx, toKnowledgeSource)
			}
		}
	}
}

func (m UpdateKnowledgeSourceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["knowledge_source"] = attrs["knowledge_source"].SetRequired()
	attrs["knowledge_source"] = attrs["knowledge_source"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateKnowledgeSourceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateKnowledgeSourceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"knowledge_source": reflect.TypeOf(KnowledgeSource_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateKnowledgeSourceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateKnowledgeSourceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"knowledge_source": m.KnowledgeSource,
			"name":             m.Name,
			"update_mask":      m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateKnowledgeSourceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"knowledge_source": basetypes.ListType{
				ElemType: KnowledgeSource_SdkV2{}.Type(ctx),
			},
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetKnowledgeSource returns the value of the KnowledgeSource field in UpdateKnowledgeSourceRequest_SdkV2 as
// a KnowledgeSource_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateKnowledgeSourceRequest_SdkV2) GetKnowledgeSource(ctx context.Context) (KnowledgeSource_SdkV2, bool) {
	var e KnowledgeSource_SdkV2
	if m.KnowledgeSource.IsNull() || m.KnowledgeSource.IsUnknown() {
		return e, false
	}
	var v []KnowledgeSource_SdkV2
	d := m.KnowledgeSource.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetKnowledgeSource sets the value of the KnowledgeSource field in UpdateKnowledgeSourceRequest_SdkV2.
func (m *UpdateKnowledgeSourceRequest_SdkV2) SetKnowledgeSource(ctx context.Context, v KnowledgeSource_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["knowledge_source"]
	m.KnowledgeSource = types.ListValueMust(t, vs)
}
