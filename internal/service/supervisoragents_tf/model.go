// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package supervisoragents_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Databricks app. Supported app: custom mcp, custom agent.
type App struct {
	// App name
	Name types.String `tfsdk:"name"`
}

func (to *App) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from App) {
}

func (to *App) SyncFieldsDuringRead(ctx context.Context, from App) {
}

func (m App) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in App.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m App) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, App
// only implements ToObjectValue() and Type().
func (m App) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m App) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type CreateExampleRequest struct {
	// The example to create under the parent Supervisor Agent.
	Example types.Object `tfsdk:"example"`
	// Parent resource where this example will be created. Format:
	// supervisor-agents/{supervisor_agent_id}
	Parent types.String `tfsdk:"-"`
}

func (to *CreateExampleRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateExampleRequest) {
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

func (to *CreateExampleRequest) SyncFieldsDuringRead(ctx context.Context, from CreateExampleRequest) {
	if !from.Example.IsNull() && !from.Example.IsUnknown() {
		if toExample, ok := to.GetExample(ctx); ok {
			if fromExample, ok := from.GetExample(ctx); ok {
				toExample.SyncFieldsDuringRead(ctx, fromExample)
				to.SetExample(ctx, toExample)
			}
		}
	}
}

func (m CreateExampleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["example"] = attrs["example"].SetRequired()
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
func (m CreateExampleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"example": reflect.TypeOf(Example{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExampleRequest
// only implements ToObjectValue() and Type().
func (m CreateExampleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"example": m.Example,
			"parent":  m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateExampleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"example": Example{}.Type(ctx),
			"parent":  types.StringType,
		},
	}
}

// GetExample returns the value of the Example field in CreateExampleRequest as
// a Example value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateExampleRequest) GetExample(ctx context.Context) (Example, bool) {
	var e Example
	if m.Example.IsNull() || m.Example.IsUnknown() {
		return e, false
	}
	var v Example
	d := m.Example.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExample sets the value of the Example field in CreateExampleRequest.
func (m *CreateExampleRequest) SetExample(ctx context.Context, v Example) {
	vs := v.ToObjectValue(ctx)
	m.Example = vs
}

type CreateSupervisorAgentRequest struct {
	// The Supervisor Agent to create.
	SupervisorAgent types.Object `tfsdk:"supervisor_agent"`
}

func (to *CreateSupervisorAgentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateSupervisorAgentRequest) {
	if !from.SupervisorAgent.IsNull() && !from.SupervisorAgent.IsUnknown() {
		if toSupervisorAgent, ok := to.GetSupervisorAgent(ctx); ok {
			if fromSupervisorAgent, ok := from.GetSupervisorAgent(ctx); ok {
				// Recursively sync the fields of SupervisorAgent
				toSupervisorAgent.SyncFieldsDuringCreateOrUpdate(ctx, fromSupervisorAgent)
				to.SetSupervisorAgent(ctx, toSupervisorAgent)
			}
		}
	}
}

func (to *CreateSupervisorAgentRequest) SyncFieldsDuringRead(ctx context.Context, from CreateSupervisorAgentRequest) {
	if !from.SupervisorAgent.IsNull() && !from.SupervisorAgent.IsUnknown() {
		if toSupervisorAgent, ok := to.GetSupervisorAgent(ctx); ok {
			if fromSupervisorAgent, ok := from.GetSupervisorAgent(ctx); ok {
				toSupervisorAgent.SyncFieldsDuringRead(ctx, fromSupervisorAgent)
				to.SetSupervisorAgent(ctx, toSupervisorAgent)
			}
		}
	}
}

func (m CreateSupervisorAgentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["supervisor_agent"] = attrs["supervisor_agent"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateSupervisorAgentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateSupervisorAgentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"supervisor_agent": reflect.TypeOf(SupervisorAgent{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateSupervisorAgentRequest
// only implements ToObjectValue() and Type().
func (m CreateSupervisorAgentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"supervisor_agent": m.SupervisorAgent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateSupervisorAgentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"supervisor_agent": SupervisorAgent{}.Type(ctx),
		},
	}
}

// GetSupervisorAgent returns the value of the SupervisorAgent field in CreateSupervisorAgentRequest as
// a SupervisorAgent value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateSupervisorAgentRequest) GetSupervisorAgent(ctx context.Context) (SupervisorAgent, bool) {
	var e SupervisorAgent
	if m.SupervisorAgent.IsNull() || m.SupervisorAgent.IsUnknown() {
		return e, false
	}
	var v SupervisorAgent
	d := m.SupervisorAgent.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSupervisorAgent sets the value of the SupervisorAgent field in CreateSupervisorAgentRequest.
func (m *CreateSupervisorAgentRequest) SetSupervisorAgent(ctx context.Context, v SupervisorAgent) {
	vs := v.ToObjectValue(ctx)
	m.SupervisorAgent = vs
}

type CreateToolRequest struct {
	// Parent resource where this tool will be created. Format:
	// supervisor-agents/{supervisor_agent_id}
	Parent types.String `tfsdk:"-"`

	Tool types.Object `tfsdk:"tool"`
	// The ID to use for the tool, which will become the final component of the
	// tool's resource name.
	ToolId types.String `tfsdk:"-"`
}

func (to *CreateToolRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateToolRequest) {
	if !from.Tool.IsNull() && !from.Tool.IsUnknown() {
		if toTool, ok := to.GetTool(ctx); ok {
			if fromTool, ok := from.GetTool(ctx); ok {
				// Recursively sync the fields of Tool
				toTool.SyncFieldsDuringCreateOrUpdate(ctx, fromTool)
				to.SetTool(ctx, toTool)
			}
		}
	}
}

func (to *CreateToolRequest) SyncFieldsDuringRead(ctx context.Context, from CreateToolRequest) {
	if !from.Tool.IsNull() && !from.Tool.IsUnknown() {
		if toTool, ok := to.GetTool(ctx); ok {
			if fromTool, ok := from.GetTool(ctx); ok {
				toTool.SyncFieldsDuringRead(ctx, fromTool)
				to.SetTool(ctx, toTool)
			}
		}
	}
}

func (m CreateToolRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["tool"] = attrs["tool"].SetRequired()
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["tool_id"] = attrs["tool_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateToolRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateToolRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tool": reflect.TypeOf(Tool{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateToolRequest
// only implements ToObjectValue() and Type().
func (m CreateToolRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parent":  m.Parent,
			"tool":    m.Tool,
			"tool_id": m.ToolId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateToolRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent":  types.StringType,
			"tool":    Tool{}.Type(ctx),
			"tool_id": types.StringType,
		},
	}
}

// GetTool returns the value of the Tool field in CreateToolRequest as
// a Tool value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateToolRequest) GetTool(ctx context.Context) (Tool, bool) {
	var e Tool
	if m.Tool.IsNull() || m.Tool.IsUnknown() {
		return e, false
	}
	var v Tool
	d := m.Tool.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTool sets the value of the Tool field in CreateToolRequest.
func (m *CreateToolRequest) SetTool(ctx context.Context, v Tool) {
	vs := v.ToObjectValue(ctx)
	m.Tool = vs
}

type DeleteExampleRequest struct {
	// The resource name of the example to delete. Format:
	// supervisor-agents/{supervisor_agent_id}/examples/{example_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteExampleRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteExampleRequest) {
}

func (to *DeleteExampleRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteExampleRequest) {
}

func (m DeleteExampleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteExampleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExampleRequest
// only implements ToObjectValue() and Type().
func (m DeleteExampleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteExampleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteSupervisorAgentRequest struct {
	// The resource name of the Supervisor Agent. Format:
	// supervisor-agents/{supervisor_agent_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteSupervisorAgentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteSupervisorAgentRequest) {
}

func (to *DeleteSupervisorAgentRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteSupervisorAgentRequest) {
}

func (m DeleteSupervisorAgentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSupervisorAgentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteSupervisorAgentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSupervisorAgentRequest
// only implements ToObjectValue() and Type().
func (m DeleteSupervisorAgentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteSupervisorAgentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteToolRequest struct {
	// The resource name of the Tool. Format:
	// supervisor-agents/{supervisor_agent_id}/tools/{tool_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteToolRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteToolRequest) {
}

func (to *DeleteToolRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteToolRequest) {
}

func (m DeleteToolRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteToolRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteToolRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteToolRequest
// only implements ToObjectValue() and Type().
func (m DeleteToolRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteToolRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// An example associated with a Supervisor Agent. Contains a question and
// guidelines for how the agent should respond.
type Example struct {
	// The universally unique identifier (UUID) of the example.
	ExampleId types.String `tfsdk:"example_id"`
	// Guidelines for answering the question.
	Guidelines types.List `tfsdk:"guidelines"`
	// Full resource name:
	// supervisor-agents/{supervisor_agent_id}/examples/{example_id}
	Name types.String `tfsdk:"name"`
	// The example question.
	Question types.String `tfsdk:"question"`
}

func (to *Example) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Example) {
}

func (to *Example) SyncFieldsDuringRead(ctx context.Context, from Example) {
}

func (m Example) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["example_id"] = attrs["example_id"].SetComputed()
	attrs["guidelines"] = attrs["guidelines"].SetRequired()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["question"] = attrs["question"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Example.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Example) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"guidelines": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Example
// only implements ToObjectValue() and Type().
func (m Example) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"example_id": m.ExampleId,
			"guidelines": m.Guidelines,
			"name":       m.Name,
			"question":   m.Question,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Example) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"example_id": types.StringType,
			"guidelines": basetypes.ListType{
				ElemType: types.StringType,
			},
			"name":     types.StringType,
			"question": types.StringType,
		},
	}
}

// GetGuidelines returns the value of the Guidelines field in Example as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Example) GetGuidelines(ctx context.Context) ([]types.String, bool) {
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

// SetGuidelines sets the value of the Guidelines field in Example.
func (m *Example) SetGuidelines(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["guidelines"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Guidelines = types.ListValueMust(t, vs)
}

type GenieSpace struct {
	// Deprecated: use space_id instead. Still REQUIRED for backward
	// compatibility until a future API version removes it.
	Id types.String `tfsdk:"id"`
}

func (to *GenieSpace) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieSpace) {
}

func (to *GenieSpace) SyncFieldsDuringRead(ctx context.Context, from GenieSpace) {
}

func (m GenieSpace) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieSpace.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenieSpace) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieSpace
// only implements ToObjectValue() and Type().
func (m GenieSpace) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieSpace) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetExampleRequest struct {
	// The resource name of the example. Format:
	// supervisor-agents/{supervisor_agent_id}/examples/{example_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetExampleRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExampleRequest) {
}

func (to *GetExampleRequest) SyncFieldsDuringRead(ctx context.Context, from GetExampleRequest) {
}

func (m GetExampleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetExampleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExampleRequest
// only implements ToObjectValue() and Type().
func (m GetExampleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExampleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetSupervisorAgentPermissionLevelsRequest struct {
	// The supervisor agent for which to get or manage permissions.
	SupervisorAgentId types.String `tfsdk:"-"`
}

func (to *GetSupervisorAgentPermissionLevelsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSupervisorAgentPermissionLevelsRequest) {
}

func (to *GetSupervisorAgentPermissionLevelsRequest) SyncFieldsDuringRead(ctx context.Context, from GetSupervisorAgentPermissionLevelsRequest) {
}

func (m GetSupervisorAgentPermissionLevelsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["supervisor_agent_id"] = attrs["supervisor_agent_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSupervisorAgentPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetSupervisorAgentPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSupervisorAgentPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (m GetSupervisorAgentPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"supervisor_agent_id": m.SupervisorAgentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSupervisorAgentPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"supervisor_agent_id": types.StringType,
		},
	}
}

type GetSupervisorAgentPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (to *GetSupervisorAgentPermissionLevelsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSupervisorAgentPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetSupervisorAgentPermissionLevelsResponse) SyncFieldsDuringRead(ctx context.Context, from GetSupervisorAgentPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (m GetSupervisorAgentPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_levels"] = attrs["permission_levels"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSupervisorAgentPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetSupervisorAgentPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(SupervisorAgentPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSupervisorAgentPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (m GetSupervisorAgentPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSupervisorAgentPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: SupervisorAgentPermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetSupervisorAgentPermissionLevelsResponse as
// a slice of SupervisorAgentPermissionsDescription values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetSupervisorAgentPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]SupervisorAgentPermissionsDescription, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []SupervisorAgentPermissionsDescription
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetSupervisorAgentPermissionLevelsResponse.
func (m *GetSupervisorAgentPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []SupervisorAgentPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetSupervisorAgentPermissionsRequest struct {
	// The supervisor agent for which to get or manage permissions.
	SupervisorAgentId types.String `tfsdk:"-"`
}

func (to *GetSupervisorAgentPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSupervisorAgentPermissionsRequest) {
}

func (to *GetSupervisorAgentPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from GetSupervisorAgentPermissionsRequest) {
}

func (m GetSupervisorAgentPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["supervisor_agent_id"] = attrs["supervisor_agent_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSupervisorAgentPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetSupervisorAgentPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSupervisorAgentPermissionsRequest
// only implements ToObjectValue() and Type().
func (m GetSupervisorAgentPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"supervisor_agent_id": m.SupervisorAgentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSupervisorAgentPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"supervisor_agent_id": types.StringType,
		},
	}
}

type GetSupervisorAgentRequest struct {
	// The resource name of the Supervisor Agent. Format:
	// supervisor-agents/{supervisor_agent_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetSupervisorAgentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSupervisorAgentRequest) {
}

func (to *GetSupervisorAgentRequest) SyncFieldsDuringRead(ctx context.Context, from GetSupervisorAgentRequest) {
}

func (m GetSupervisorAgentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSupervisorAgentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetSupervisorAgentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSupervisorAgentRequest
// only implements ToObjectValue() and Type().
func (m GetSupervisorAgentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSupervisorAgentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetToolRequest struct {
	// The resource name of the Tool. Format:
	// supervisor-agents/{supervisor_agent_id}/tools/{tool_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetToolRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetToolRequest) {
}

func (to *GetToolRequest) SyncFieldsDuringRead(ctx context.Context, from GetToolRequest) {
}

func (m GetToolRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetToolRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetToolRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetToolRequest
// only implements ToObjectValue() and Type().
func (m GetToolRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetToolRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type KnowledgeAssistant struct {
	// The ID of the knowledge assistant.
	KnowledgeAssistantId types.String `tfsdk:"knowledge_assistant_id"`
	// Deprecated: use knowledge_assistant_id instead.
	ServingEndpointName types.String `tfsdk:"serving_endpoint_name"`
}

func (to *KnowledgeAssistant) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KnowledgeAssistant) {
}

func (to *KnowledgeAssistant) SyncFieldsDuringRead(ctx context.Context, from KnowledgeAssistant) {
}

func (m KnowledgeAssistant) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["knowledge_assistant_id"] = attrs["knowledge_assistant_id"].SetRequired()
	attrs["serving_endpoint_name"] = attrs["serving_endpoint_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in KnowledgeAssistant.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m KnowledgeAssistant) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, KnowledgeAssistant
// only implements ToObjectValue() and Type().
func (m KnowledgeAssistant) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"knowledge_assistant_id": m.KnowledgeAssistantId,
			"serving_endpoint_name":  m.ServingEndpointName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m KnowledgeAssistant) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"knowledge_assistant_id": types.StringType,
			"serving_endpoint_name":  types.StringType,
		},
	}
}

type ListExamplesRequest struct {
	// The maximum number of examples to return. If unspecified, at most 100
	// examples will be returned. The maximum value is 100; values above 100
	// will be coerced to 100.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListExamples` call. Provide this
	// to retrieve the subsequent page. If unspecified, the first page will be
	// returned.
	PageToken types.String `tfsdk:"-"`
	// Parent resource to list from. Format:
	// supervisor-agents/{supervisor_agent_id}
	Parent types.String `tfsdk:"-"`
}

func (to *ListExamplesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListExamplesRequest) {
}

func (to *ListExamplesRequest) SyncFieldsDuringRead(ctx context.Context, from ListExamplesRequest) {
}

func (m ListExamplesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListExamplesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExamplesRequest
// only implements ToObjectValue() and Type().
func (m ListExamplesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListExamplesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// A list of Supervisor Agent examples.
type ListExamplesResponse struct {
	Examples types.List `tfsdk:"examples"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListExamplesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListExamplesResponse) {
	if !from.Examples.IsNull() && !from.Examples.IsUnknown() && to.Examples.IsNull() && len(from.Examples.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Examples, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Examples = from.Examples
	}
}

func (to *ListExamplesResponse) SyncFieldsDuringRead(ctx context.Context, from ListExamplesResponse) {
	if !from.Examples.IsNull() && !from.Examples.IsUnknown() && to.Examples.IsNull() && len(from.Examples.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Examples, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Examples = from.Examples
	}
}

func (m ListExamplesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListExamplesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"examples": reflect.TypeOf(Example{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExamplesResponse
// only implements ToObjectValue() and Type().
func (m ListExamplesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"examples":        m.Examples,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListExamplesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"examples": basetypes.ListType{
				ElemType: Example{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetExamples returns the value of the Examples field in ListExamplesResponse as
// a slice of Example values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListExamplesResponse) GetExamples(ctx context.Context) ([]Example, bool) {
	if m.Examples.IsNull() || m.Examples.IsUnknown() {
		return nil, false
	}
	var v []Example
	d := m.Examples.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExamples sets the value of the Examples field in ListExamplesResponse.
func (m *ListExamplesResponse) SetExamples(ctx context.Context, v []Example) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["examples"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Examples = types.ListValueMust(t, vs)
}

type ListSupervisorAgentsRequest struct {
	// The maximum number of supervisor agents to return. If unspecified, at
	// most 100 supervisor agents will be returned. The maximum value is 100;
	// values above 100 will be coerced to 100.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListSupervisorAgents` call.
	// Provide this to retrieve the subsequent page. If unspecified, the first
	// page will be returned.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListSupervisorAgentsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSupervisorAgentsRequest) {
}

func (to *ListSupervisorAgentsRequest) SyncFieldsDuringRead(ctx context.Context, from ListSupervisorAgentsRequest) {
}

func (m ListSupervisorAgentsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSupervisorAgentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSupervisorAgentsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSupervisorAgentsRequest
// only implements ToObjectValue() and Type().
func (m ListSupervisorAgentsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSupervisorAgentsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListSupervisorAgentsResponse struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`

	SupervisorAgents types.List `tfsdk:"supervisor_agents"`
}

func (to *ListSupervisorAgentsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSupervisorAgentsResponse) {
	if !from.SupervisorAgents.IsNull() && !from.SupervisorAgents.IsUnknown() && to.SupervisorAgents.IsNull() && len(from.SupervisorAgents.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SupervisorAgents, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SupervisorAgents = from.SupervisorAgents
	}
}

func (to *ListSupervisorAgentsResponse) SyncFieldsDuringRead(ctx context.Context, from ListSupervisorAgentsResponse) {
	if !from.SupervisorAgents.IsNull() && !from.SupervisorAgents.IsUnknown() && to.SupervisorAgents.IsNull() && len(from.SupervisorAgents.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SupervisorAgents, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SupervisorAgents = from.SupervisorAgents
	}
}

func (m ListSupervisorAgentsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["supervisor_agents"] = attrs["supervisor_agents"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSupervisorAgentsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSupervisorAgentsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"supervisor_agents": reflect.TypeOf(SupervisorAgent{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSupervisorAgentsResponse
// only implements ToObjectValue() and Type().
func (m ListSupervisorAgentsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":   m.NextPageToken,
			"supervisor_agents": m.SupervisorAgents,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSupervisorAgentsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"supervisor_agents": basetypes.ListType{
				ElemType: SupervisorAgent{}.Type(ctx),
			},
		},
	}
}

// GetSupervisorAgents returns the value of the SupervisorAgents field in ListSupervisorAgentsResponse as
// a slice of SupervisorAgent values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListSupervisorAgentsResponse) GetSupervisorAgents(ctx context.Context) ([]SupervisorAgent, bool) {
	if m.SupervisorAgents.IsNull() || m.SupervisorAgents.IsUnknown() {
		return nil, false
	}
	var v []SupervisorAgent
	d := m.SupervisorAgents.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSupervisorAgents sets the value of the SupervisorAgents field in ListSupervisorAgentsResponse.
func (m *ListSupervisorAgentsResponse) SetSupervisorAgents(ctx context.Context, v []SupervisorAgent) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["supervisor_agents"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SupervisorAgents = types.ListValueMust(t, vs)
}

type ListToolsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
	// Parent resource to list from. Format:
	// supervisor-agents/{supervisor_agent_id}
	Parent types.String `tfsdk:"-"`
}

func (to *ListToolsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListToolsRequest) {
}

func (to *ListToolsRequest) SyncFieldsDuringRead(ctx context.Context, from ListToolsRequest) {
}

func (m ListToolsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListToolsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListToolsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListToolsRequest
// only implements ToObjectValue() and Type().
func (m ListToolsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListToolsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListToolsResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Tools types.List `tfsdk:"tools"`
}

func (to *ListToolsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListToolsResponse) {
	if !from.Tools.IsNull() && !from.Tools.IsUnknown() && to.Tools.IsNull() && len(from.Tools.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tools, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tools = from.Tools
	}
}

func (to *ListToolsResponse) SyncFieldsDuringRead(ctx context.Context, from ListToolsResponse) {
	if !from.Tools.IsNull() && !from.Tools.IsUnknown() && to.Tools.IsNull() && len(from.Tools.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tools, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tools = from.Tools
	}
}

func (m ListToolsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["tools"] = attrs["tools"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListToolsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListToolsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tools": reflect.TypeOf(Tool{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListToolsResponse
// only implements ToObjectValue() and Type().
func (m ListToolsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"tools":           m.Tools,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListToolsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"tools": basetypes.ListType{
				ElemType: Tool{}.Type(ctx),
			},
		},
	}
}

// GetTools returns the value of the Tools field in ListToolsResponse as
// a slice of Tool values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListToolsResponse) GetTools(ctx context.Context) ([]Tool, bool) {
	if m.Tools.IsNull() || m.Tools.IsUnknown() {
		return nil, false
	}
	var v []Tool
	d := m.Tools.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTools sets the value of the Tools field in ListToolsResponse.
func (m *ListToolsResponse) SetTools(ctx context.Context, v []Tool) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tools"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tools = types.ListValueMust(t, vs)
}

type SupervisorAgent struct {
	// Creation timestamp.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The creator of the Supervisor Agent.
	Creator types.String `tfsdk:"creator"`
	// Description of what this agent can do (user-facing).
	Description types.String `tfsdk:"description"`
	// The display name of the Supervisor Agent, unique at workspace level.
	DisplayName types.String `tfsdk:"display_name"`
	// The name of the supervisor agent's serving endpoint.
	EndpointName types.String `tfsdk:"endpoint_name"`
	// The MLflow experiment ID.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// Deprecated: Use supervisor_agent_id instead.
	Id types.String `tfsdk:"id"`
	// Optional natural-language instructions for the supervisor agent.
	Instructions types.String `tfsdk:"instructions"`
	// The resource name of the SupervisorAgent. Format:
	// supervisor-agents/{supervisor_agent_id}
	Name types.String `tfsdk:"name"`
	// The universally unique identifier (UUID) of the Supervisor Agent.
	SupervisorAgentId types.String `tfsdk:"supervisor_agent_id"`
}

func (to *SupervisorAgent) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SupervisorAgent) {
}

func (to *SupervisorAgent) SyncFieldsDuringRead(ctx context.Context, from SupervisorAgent) {
}

func (m SupervisorAgent) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetRequired()
	attrs["endpoint_name"] = attrs["endpoint_name"].SetComputed()
	attrs["experiment_id"] = attrs["experiment_id"].SetComputed()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["instructions"] = attrs["instructions"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["supervisor_agent_id"] = attrs["supervisor_agent_id"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SupervisorAgent.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SupervisorAgent) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SupervisorAgent
// only implements ToObjectValue() and Type().
func (m SupervisorAgent) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":         m.CreateTime,
			"creator":             m.Creator,
			"description":         m.Description,
			"display_name":        m.DisplayName,
			"endpoint_name":       m.EndpointName,
			"experiment_id":       m.ExperimentId,
			"id":                  m.Id,
			"instructions":        m.Instructions,
			"name":                m.Name,
			"supervisor_agent_id": m.SupervisorAgentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SupervisorAgent) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":         timetypes.RFC3339{}.Type(ctx),
			"creator":             types.StringType,
			"description":         types.StringType,
			"display_name":        types.StringType,
			"endpoint_name":       types.StringType,
			"experiment_id":       types.StringType,
			"id":                  types.StringType,
			"instructions":        types.StringType,
			"name":                types.StringType,
			"supervisor_agent_id": types.StringType,
		},
	}
}

type SupervisorAgentAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *SupervisorAgentAccessControlRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SupervisorAgentAccessControlRequest) {
}

func (to *SupervisorAgentAccessControlRequest) SyncFieldsDuringRead(ctx context.Context, from SupervisorAgentAccessControlRequest) {
}

func (m SupervisorAgentAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SupervisorAgentAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SupervisorAgentAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SupervisorAgentAccessControlRequest
// only implements ToObjectValue() and Type().
func (m SupervisorAgentAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m SupervisorAgentAccessControlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type SupervisorAgentAccessControlResponse struct {
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

func (to *SupervisorAgentAccessControlResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SupervisorAgentAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *SupervisorAgentAccessControlResponse) SyncFieldsDuringRead(ctx context.Context, from SupervisorAgentAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (m SupervisorAgentAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["all_permissions"] = attrs["all_permissions"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SupervisorAgentAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SupervisorAgentAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(SupervisorAgentPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SupervisorAgentAccessControlResponse
// only implements ToObjectValue() and Type().
func (m SupervisorAgentAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m SupervisorAgentAccessControlResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: SupervisorAgentPermission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in SupervisorAgentAccessControlResponse as
// a slice of SupervisorAgentPermission values.
// If the field is unknown or null, the boolean return value is false.
func (m *SupervisorAgentAccessControlResponse) GetAllPermissions(ctx context.Context) ([]SupervisorAgentPermission, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []SupervisorAgentPermission
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in SupervisorAgentAccessControlResponse.
func (m *SupervisorAgentAccessControlResponse) SetAllPermissions(ctx context.Context, v []SupervisorAgentPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
}

type SupervisorAgentPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *SupervisorAgentPermission) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SupervisorAgentPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *SupervisorAgentPermission) SyncFieldsDuringRead(ctx context.Context, from SupervisorAgentPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m SupervisorAgentPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited"] = attrs["inherited"].SetOptional()
	attrs["inherited_from_object"] = attrs["inherited_from_object"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SupervisorAgentPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SupervisorAgentPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SupervisorAgentPermission
// only implements ToObjectValue() and Type().
func (m SupervisorAgentPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SupervisorAgentPermission) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in SupervisorAgentPermission as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SupervisorAgentPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in SupervisorAgentPermission.
func (m *SupervisorAgentPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type SupervisorAgentPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *SupervisorAgentPermissions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SupervisorAgentPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *SupervisorAgentPermissions) SyncFieldsDuringRead(ctx context.Context, from SupervisorAgentPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m SupervisorAgentPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SupervisorAgentPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SupervisorAgentPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(SupervisorAgentAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SupervisorAgentPermissions
// only implements ToObjectValue() and Type().
func (m SupervisorAgentPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SupervisorAgentPermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: SupervisorAgentAccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in SupervisorAgentPermissions as
// a slice of SupervisorAgentAccessControlResponse values.
// If the field is unknown or null, the boolean return value is false.
func (m *SupervisorAgentPermissions) GetAccessControlList(ctx context.Context) ([]SupervisorAgentAccessControlResponse, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []SupervisorAgentAccessControlResponse
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in SupervisorAgentPermissions.
func (m *SupervisorAgentPermissions) SetAccessControlList(ctx context.Context, v []SupervisorAgentAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type SupervisorAgentPermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *SupervisorAgentPermissionsDescription) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SupervisorAgentPermissionsDescription) {
}

func (to *SupervisorAgentPermissionsDescription) SyncFieldsDuringRead(ctx context.Context, from SupervisorAgentPermissionsDescription) {
}

func (m SupervisorAgentPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SupervisorAgentPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SupervisorAgentPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SupervisorAgentPermissionsDescription
// only implements ToObjectValue() and Type().
func (m SupervisorAgentPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SupervisorAgentPermissionsDescription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type SupervisorAgentPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The supervisor agent for which to get or manage permissions.
	SupervisorAgentId types.String `tfsdk:"-"`
}

func (to *SupervisorAgentPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SupervisorAgentPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *SupervisorAgentPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from SupervisorAgentPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m SupervisorAgentPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["supervisor_agent_id"] = attrs["supervisor_agent_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SupervisorAgentPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SupervisorAgentPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(SupervisorAgentAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SupervisorAgentPermissionsRequest
// only implements ToObjectValue() and Type().
func (m SupervisorAgentPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"supervisor_agent_id": m.SupervisorAgentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SupervisorAgentPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: SupervisorAgentAccessControlRequest{}.Type(ctx),
			},
			"supervisor_agent_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in SupervisorAgentPermissionsRequest as
// a slice of SupervisorAgentAccessControlRequest values.
// If the field is unknown or null, the boolean return value is false.
func (m *SupervisorAgentPermissionsRequest) GetAccessControlList(ctx context.Context) ([]SupervisorAgentAccessControlRequest, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []SupervisorAgentAccessControlRequest
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in SupervisorAgentPermissionsRequest.
func (m *SupervisorAgentPermissionsRequest) SetAccessControlList(ctx context.Context, v []SupervisorAgentAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type Tool struct {
	App types.Object `tfsdk:"app"`
	// Description of what this tool does (user-facing).
	Description types.String `tfsdk:"description"`

	GenieSpace types.Object `tfsdk:"genie_space"`
	// Deprecated: Use tool_id instead.
	Id types.String `tfsdk:"id"`

	KnowledgeAssistant types.Object `tfsdk:"knowledge_assistant"`
	// Full resource name:
	// supervisor-agents/{supervisor_agent_id}/tools/{tool_id}
	Name types.String `tfsdk:"name"`
	// User specified id of the Tool.
	ToolId types.String `tfsdk:"tool_id"`
	// Tool type. Must be one of: "genie_space", "knowledge_assistant",
	// "uc_function", "uc_connection", "uc_mcp", "app", "volume", "dashboard",
	// "serving_endpoint", "table", "vector_search_index", "catalog", "schema",
	// "supervisor_agent", "databricks_web_search", "skill". The legacy values
	// "lakeview_dashboard", "uc_table", and "web_search" are also accepted and
	// remain equivalent to "dashboard", "table", and "databricks_web_search"
	// respectively. The "databricks_web_search" tool_type maps to the
	// `web_search` spec field.
	ToolType types.String `tfsdk:"tool_type"`

	UcConnection types.Object `tfsdk:"uc_connection"`

	UcFunction types.Object `tfsdk:"uc_function"`

	Volume types.Object `tfsdk:"volume"`
}

func (to *Tool) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Tool) {
	if !from.App.IsNull() && !from.App.IsUnknown() {
		if toApp, ok := to.GetApp(ctx); ok {
			if fromApp, ok := from.GetApp(ctx); ok {
				// Recursively sync the fields of App
				toApp.SyncFieldsDuringCreateOrUpdate(ctx, fromApp)
				to.SetApp(ctx, toApp)
			}
		}
	}
	if !from.GenieSpace.IsNull() && !from.GenieSpace.IsUnknown() {
		if toGenieSpace, ok := to.GetGenieSpace(ctx); ok {
			if fromGenieSpace, ok := from.GetGenieSpace(ctx); ok {
				// Recursively sync the fields of GenieSpace
				toGenieSpace.SyncFieldsDuringCreateOrUpdate(ctx, fromGenieSpace)
				to.SetGenieSpace(ctx, toGenieSpace)
			}
		}
	}
	if !from.KnowledgeAssistant.IsNull() && !from.KnowledgeAssistant.IsUnknown() {
		if toKnowledgeAssistant, ok := to.GetKnowledgeAssistant(ctx); ok {
			if fromKnowledgeAssistant, ok := from.GetKnowledgeAssistant(ctx); ok {
				// Recursively sync the fields of KnowledgeAssistant
				toKnowledgeAssistant.SyncFieldsDuringCreateOrUpdate(ctx, fromKnowledgeAssistant)
				to.SetKnowledgeAssistant(ctx, toKnowledgeAssistant)
			}
		}
	}
	if !from.UcConnection.IsNull() && !from.UcConnection.IsUnknown() {
		if toUcConnection, ok := to.GetUcConnection(ctx); ok {
			if fromUcConnection, ok := from.GetUcConnection(ctx); ok {
				// Recursively sync the fields of UcConnection
				toUcConnection.SyncFieldsDuringCreateOrUpdate(ctx, fromUcConnection)
				to.SetUcConnection(ctx, toUcConnection)
			}
		}
	}
	if !from.UcFunction.IsNull() && !from.UcFunction.IsUnknown() {
		if toUcFunction, ok := to.GetUcFunction(ctx); ok {
			if fromUcFunction, ok := from.GetUcFunction(ctx); ok {
				// Recursively sync the fields of UcFunction
				toUcFunction.SyncFieldsDuringCreateOrUpdate(ctx, fromUcFunction)
				to.SetUcFunction(ctx, toUcFunction)
			}
		}
	}
	if !from.Volume.IsNull() && !from.Volume.IsUnknown() {
		if toVolume, ok := to.GetVolume(ctx); ok {
			if fromVolume, ok := from.GetVolume(ctx); ok {
				// Recursively sync the fields of Volume
				toVolume.SyncFieldsDuringCreateOrUpdate(ctx, fromVolume)
				to.SetVolume(ctx, toVolume)
			}
		}
	}
}

func (to *Tool) SyncFieldsDuringRead(ctx context.Context, from Tool) {
	if !from.App.IsNull() && !from.App.IsUnknown() {
		if toApp, ok := to.GetApp(ctx); ok {
			if fromApp, ok := from.GetApp(ctx); ok {
				toApp.SyncFieldsDuringRead(ctx, fromApp)
				to.SetApp(ctx, toApp)
			}
		}
	}
	if !from.GenieSpace.IsNull() && !from.GenieSpace.IsUnknown() {
		if toGenieSpace, ok := to.GetGenieSpace(ctx); ok {
			if fromGenieSpace, ok := from.GetGenieSpace(ctx); ok {
				toGenieSpace.SyncFieldsDuringRead(ctx, fromGenieSpace)
				to.SetGenieSpace(ctx, toGenieSpace)
			}
		}
	}
	if !from.KnowledgeAssistant.IsNull() && !from.KnowledgeAssistant.IsUnknown() {
		if toKnowledgeAssistant, ok := to.GetKnowledgeAssistant(ctx); ok {
			if fromKnowledgeAssistant, ok := from.GetKnowledgeAssistant(ctx); ok {
				toKnowledgeAssistant.SyncFieldsDuringRead(ctx, fromKnowledgeAssistant)
				to.SetKnowledgeAssistant(ctx, toKnowledgeAssistant)
			}
		}
	}
	if !from.UcConnection.IsNull() && !from.UcConnection.IsUnknown() {
		if toUcConnection, ok := to.GetUcConnection(ctx); ok {
			if fromUcConnection, ok := from.GetUcConnection(ctx); ok {
				toUcConnection.SyncFieldsDuringRead(ctx, fromUcConnection)
				to.SetUcConnection(ctx, toUcConnection)
			}
		}
	}
	if !from.UcFunction.IsNull() && !from.UcFunction.IsUnknown() {
		if toUcFunction, ok := to.GetUcFunction(ctx); ok {
			if fromUcFunction, ok := from.GetUcFunction(ctx); ok {
				toUcFunction.SyncFieldsDuringRead(ctx, fromUcFunction)
				to.SetUcFunction(ctx, toUcFunction)
			}
		}
	}
	if !from.Volume.IsNull() && !from.Volume.IsUnknown() {
		if toVolume, ok := to.GetVolume(ctx); ok {
			if fromVolume, ok := from.GetVolume(ctx); ok {
				toVolume.SyncFieldsDuringRead(ctx, fromVolume)
				to.SetVolume(ctx, toVolume)
			}
		}
	}
}

func (m Tool) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app"] = attrs["app"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["genie_space"] = attrs["genie_space"].SetOptional()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["knowledge_assistant"] = attrs["knowledge_assistant"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["tool_id"] = attrs["tool_id"].SetComputed()
	attrs["tool_type"] = attrs["tool_type"].SetRequired()
	attrs["uc_connection"] = attrs["uc_connection"].SetOptional()
	attrs["uc_function"] = attrs["uc_function"].SetOptional()
	attrs["volume"] = attrs["volume"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Tool.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Tool) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"app":                 reflect.TypeOf(App{}),
		"genie_space":         reflect.TypeOf(GenieSpace{}),
		"knowledge_assistant": reflect.TypeOf(KnowledgeAssistant{}),
		"uc_connection":       reflect.TypeOf(UcConnection{}),
		"uc_function":         reflect.TypeOf(UcFunction{}),
		"volume":              reflect.TypeOf(Volume{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Tool
// only implements ToObjectValue() and Type().
func (m Tool) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app":                 m.App,
			"description":         m.Description,
			"genie_space":         m.GenieSpace,
			"id":                  m.Id,
			"knowledge_assistant": m.KnowledgeAssistant,
			"name":                m.Name,
			"tool_id":             m.ToolId,
			"tool_type":           m.ToolType,
			"uc_connection":       m.UcConnection,
			"uc_function":         m.UcFunction,
			"volume":              m.Volume,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Tool) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app":                 App{}.Type(ctx),
			"description":         types.StringType,
			"genie_space":         GenieSpace{}.Type(ctx),
			"id":                  types.StringType,
			"knowledge_assistant": KnowledgeAssistant{}.Type(ctx),
			"name":                types.StringType,
			"tool_id":             types.StringType,
			"tool_type":           types.StringType,
			"uc_connection":       UcConnection{}.Type(ctx),
			"uc_function":         UcFunction{}.Type(ctx),
			"volume":              Volume{}.Type(ctx),
		},
	}
}

// GetApp returns the value of the App field in Tool as
// a App value.
// If the field is unknown or null, the boolean return value is false.
func (m *Tool) GetApp(ctx context.Context) (App, bool) {
	var e App
	if m.App.IsNull() || m.App.IsUnknown() {
		return e, false
	}
	var v App
	d := m.App.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApp sets the value of the App field in Tool.
func (m *Tool) SetApp(ctx context.Context, v App) {
	vs := v.ToObjectValue(ctx)
	m.App = vs
}

// GetGenieSpace returns the value of the GenieSpace field in Tool as
// a GenieSpace value.
// If the field is unknown or null, the boolean return value is false.
func (m *Tool) GetGenieSpace(ctx context.Context) (GenieSpace, bool) {
	var e GenieSpace
	if m.GenieSpace.IsNull() || m.GenieSpace.IsUnknown() {
		return e, false
	}
	var v GenieSpace
	d := m.GenieSpace.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGenieSpace sets the value of the GenieSpace field in Tool.
func (m *Tool) SetGenieSpace(ctx context.Context, v GenieSpace) {
	vs := v.ToObjectValue(ctx)
	m.GenieSpace = vs
}

// GetKnowledgeAssistant returns the value of the KnowledgeAssistant field in Tool as
// a KnowledgeAssistant value.
// If the field is unknown or null, the boolean return value is false.
func (m *Tool) GetKnowledgeAssistant(ctx context.Context) (KnowledgeAssistant, bool) {
	var e KnowledgeAssistant
	if m.KnowledgeAssistant.IsNull() || m.KnowledgeAssistant.IsUnknown() {
		return e, false
	}
	var v KnowledgeAssistant
	d := m.KnowledgeAssistant.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetKnowledgeAssistant sets the value of the KnowledgeAssistant field in Tool.
func (m *Tool) SetKnowledgeAssistant(ctx context.Context, v KnowledgeAssistant) {
	vs := v.ToObjectValue(ctx)
	m.KnowledgeAssistant = vs
}

// GetUcConnection returns the value of the UcConnection field in Tool as
// a UcConnection value.
// If the field is unknown or null, the boolean return value is false.
func (m *Tool) GetUcConnection(ctx context.Context) (UcConnection, bool) {
	var e UcConnection
	if m.UcConnection.IsNull() || m.UcConnection.IsUnknown() {
		return e, false
	}
	var v UcConnection
	d := m.UcConnection.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUcConnection sets the value of the UcConnection field in Tool.
func (m *Tool) SetUcConnection(ctx context.Context, v UcConnection) {
	vs := v.ToObjectValue(ctx)
	m.UcConnection = vs
}

// GetUcFunction returns the value of the UcFunction field in Tool as
// a UcFunction value.
// If the field is unknown or null, the boolean return value is false.
func (m *Tool) GetUcFunction(ctx context.Context) (UcFunction, bool) {
	var e UcFunction
	if m.UcFunction.IsNull() || m.UcFunction.IsUnknown() {
		return e, false
	}
	var v UcFunction
	d := m.UcFunction.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUcFunction sets the value of the UcFunction field in Tool.
func (m *Tool) SetUcFunction(ctx context.Context, v UcFunction) {
	vs := v.ToObjectValue(ctx)
	m.UcFunction = vs
}

// GetVolume returns the value of the Volume field in Tool as
// a Volume value.
// If the field is unknown or null, the boolean return value is false.
func (m *Tool) GetVolume(ctx context.Context) (Volume, bool) {
	var e Volume
	if m.Volume.IsNull() || m.Volume.IsUnknown() {
		return e, false
	}
	var v Volume
	d := m.Volume.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVolume sets the value of the Volume field in Tool.
func (m *Tool) SetVolume(ctx context.Context, v Volume) {
	vs := v.ToObjectValue(ctx)
	m.Volume = vs
}

// Databricks UC connection. Supported connection: external mcp server.
type UcConnection struct {
	Name types.String `tfsdk:"name"`
}

func (to *UcConnection) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UcConnection) {
}

func (to *UcConnection) SyncFieldsDuringRead(ctx context.Context, from UcConnection) {
}

func (m UcConnection) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UcConnection.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UcConnection) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UcConnection
// only implements ToObjectValue() and Type().
func (m UcConnection) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UcConnection) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type UcFunction struct {
	// Full uc function name
	Name types.String `tfsdk:"name"`
}

func (to *UcFunction) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UcFunction) {
}

func (to *UcFunction) SyncFieldsDuringRead(ctx context.Context, from UcFunction) {
}

func (m UcFunction) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UcFunction.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UcFunction) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UcFunction
// only implements ToObjectValue() and Type().
func (m UcFunction) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UcFunction) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type UpdateExampleRequest struct {
	Example types.Object `tfsdk:"example"`
	// The resource name of the example to update. Format:
	// supervisor-agents/{supervisor_agent_id}/examples/{example_id}
	Name types.String `tfsdk:"-"`
	// Comma-delimited list of fields to update on the example. Allowed values:
	// `question`, `guidelines`. Examples: - `question` - `question,guidelines`
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateExampleRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateExampleRequest) {
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

func (to *UpdateExampleRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateExampleRequest) {
	if !from.Example.IsNull() && !from.Example.IsUnknown() {
		if toExample, ok := to.GetExample(ctx); ok {
			if fromExample, ok := from.GetExample(ctx); ok {
				toExample.SyncFieldsDuringRead(ctx, fromExample)
				to.SetExample(ctx, toExample)
			}
		}
	}
}

func (m UpdateExampleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["example"] = attrs["example"].SetRequired()
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
func (m UpdateExampleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"example": reflect.TypeOf(Example{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExampleRequest
// only implements ToObjectValue() and Type().
func (m UpdateExampleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"example":     m.Example,
			"name":        m.Name,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateExampleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"example":     Example{}.Type(ctx),
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetExample returns the value of the Example field in UpdateExampleRequest as
// a Example value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateExampleRequest) GetExample(ctx context.Context) (Example, bool) {
	var e Example
	if m.Example.IsNull() || m.Example.IsUnknown() {
		return e, false
	}
	var v Example
	d := m.Example.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExample sets the value of the Example field in UpdateExampleRequest.
func (m *UpdateExampleRequest) SetExample(ctx context.Context, v Example) {
	vs := v.ToObjectValue(ctx)
	m.Example = vs
}

type UpdateSupervisorAgentRequest struct {
	// The resource name of the SupervisorAgent. Format:
	// supervisor-agents/{supervisor_agent_id}
	Name types.String `tfsdk:"-"`
	// The SupervisorAgent to update.
	SupervisorAgent types.Object `tfsdk:"supervisor_agent"`
	// Field mask for fields to be updated.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateSupervisorAgentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateSupervisorAgentRequest) {
	if !from.SupervisorAgent.IsNull() && !from.SupervisorAgent.IsUnknown() {
		if toSupervisorAgent, ok := to.GetSupervisorAgent(ctx); ok {
			if fromSupervisorAgent, ok := from.GetSupervisorAgent(ctx); ok {
				// Recursively sync the fields of SupervisorAgent
				toSupervisorAgent.SyncFieldsDuringCreateOrUpdate(ctx, fromSupervisorAgent)
				to.SetSupervisorAgent(ctx, toSupervisorAgent)
			}
		}
	}
}

func (to *UpdateSupervisorAgentRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateSupervisorAgentRequest) {
	if !from.SupervisorAgent.IsNull() && !from.SupervisorAgent.IsUnknown() {
		if toSupervisorAgent, ok := to.GetSupervisorAgent(ctx); ok {
			if fromSupervisorAgent, ok := from.GetSupervisorAgent(ctx); ok {
				toSupervisorAgent.SyncFieldsDuringRead(ctx, fromSupervisorAgent)
				to.SetSupervisorAgent(ctx, toSupervisorAgent)
			}
		}
	}
}

func (m UpdateSupervisorAgentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["supervisor_agent"] = attrs["supervisor_agent"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateSupervisorAgentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateSupervisorAgentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"supervisor_agent": reflect.TypeOf(SupervisorAgent{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSupervisorAgentRequest
// only implements ToObjectValue() and Type().
func (m UpdateSupervisorAgentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":             m.Name,
			"supervisor_agent": m.SupervisorAgent,
			"update_mask":      m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateSupervisorAgentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":             types.StringType,
			"supervisor_agent": SupervisorAgent{}.Type(ctx),
			"update_mask":      types.StringType,
		},
	}
}

// GetSupervisorAgent returns the value of the SupervisorAgent field in UpdateSupervisorAgentRequest as
// a SupervisorAgent value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateSupervisorAgentRequest) GetSupervisorAgent(ctx context.Context) (SupervisorAgent, bool) {
	var e SupervisorAgent
	if m.SupervisorAgent.IsNull() || m.SupervisorAgent.IsUnknown() {
		return e, false
	}
	var v SupervisorAgent
	d := m.SupervisorAgent.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSupervisorAgent sets the value of the SupervisorAgent field in UpdateSupervisorAgentRequest.
func (m *UpdateSupervisorAgentRequest) SetSupervisorAgent(ctx context.Context, v SupervisorAgent) {
	vs := v.ToObjectValue(ctx)
	m.SupervisorAgent = vs
}

type UpdateToolRequest struct {
	// Full resource name:
	// supervisor-agents/{supervisor_agent_id}/tools/{tool_id}
	Name types.String `tfsdk:"-"`
	// The Tool to update.
	Tool types.Object `tfsdk:"tool"`
	// Field mask for fields to be updated.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateToolRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateToolRequest) {
	if !from.Tool.IsNull() && !from.Tool.IsUnknown() {
		if toTool, ok := to.GetTool(ctx); ok {
			if fromTool, ok := from.GetTool(ctx); ok {
				// Recursively sync the fields of Tool
				toTool.SyncFieldsDuringCreateOrUpdate(ctx, fromTool)
				to.SetTool(ctx, toTool)
			}
		}
	}
}

func (to *UpdateToolRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateToolRequest) {
	if !from.Tool.IsNull() && !from.Tool.IsUnknown() {
		if toTool, ok := to.GetTool(ctx); ok {
			if fromTool, ok := from.GetTool(ctx); ok {
				toTool.SyncFieldsDuringRead(ctx, fromTool)
				to.SetTool(ctx, toTool)
			}
		}
	}
}

func (m UpdateToolRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["tool"] = attrs["tool"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateToolRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateToolRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tool": reflect.TypeOf(Tool{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateToolRequest
// only implements ToObjectValue() and Type().
func (m UpdateToolRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":        m.Name,
			"tool":        m.Tool,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateToolRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":        types.StringType,
			"tool":        Tool{}.Type(ctx),
			"update_mask": types.StringType,
		},
	}
}

// GetTool returns the value of the Tool field in UpdateToolRequest as
// a Tool value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateToolRequest) GetTool(ctx context.Context) (Tool, bool) {
	var e Tool
	if m.Tool.IsNull() || m.Tool.IsUnknown() {
		return e, false
	}
	var v Tool
	d := m.Tool.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTool sets the value of the Tool field in UpdateToolRequest.
func (m *UpdateToolRequest) SetTool(ctx context.Context, v Tool) {
	vs := v.ToObjectValue(ctx)
	m.Tool = vs
}

type Volume struct {
	// Full uc volume name
	Name types.String `tfsdk:"name"`
}

func (to *Volume) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Volume) {
}

func (to *Volume) SyncFieldsDuringRead(ctx context.Context, from Volume) {
}

func (m Volume) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Volume.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Volume) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Volume
// only implements ToObjectValue() and Type().
func (m Volume) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Volume) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}
