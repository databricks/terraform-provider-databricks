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
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Databricks app. Supported app: custom mcp, custom agent.
type App_SdkV2 struct {
	// App name
	Name types.String `tfsdk:"name"`
}

func (to *App_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from App_SdkV2) {
}

func (to *App_SdkV2) SyncFieldsDuringRead(ctx context.Context, from App_SdkV2) {
}

func (m App_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m App_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, App_SdkV2
// only implements ToObjectValue() and Type().
func (m App_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m App_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Databricks connection. Supported connection: external mcp server.
type Connection_SdkV2 struct {
	Name types.String `tfsdk:"name"`
}

func (to *Connection_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Connection_SdkV2) {
}

func (to *Connection_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Connection_SdkV2) {
}

func (m Connection_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Connection.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Connection_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Connection_SdkV2
// only implements ToObjectValue() and Type().
func (m Connection_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Connection_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type CreateSupervisorAgentRequest_SdkV2 struct {
	// The Supervisor Agent to create.
	SupervisorAgent types.List `tfsdk:"supervisor_agent"`
}

func (to *CreateSupervisorAgentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateSupervisorAgentRequest_SdkV2) {
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

func (to *CreateSupervisorAgentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateSupervisorAgentRequest_SdkV2) {
	if !from.SupervisorAgent.IsNull() && !from.SupervisorAgent.IsUnknown() {
		if toSupervisorAgent, ok := to.GetSupervisorAgent(ctx); ok {
			if fromSupervisorAgent, ok := from.GetSupervisorAgent(ctx); ok {
				toSupervisorAgent.SyncFieldsDuringRead(ctx, fromSupervisorAgent)
				to.SetSupervisorAgent(ctx, toSupervisorAgent)
			}
		}
	}
}

func (m CreateSupervisorAgentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["supervisor_agent"] = attrs["supervisor_agent"].SetRequired()
	attrs["supervisor_agent"] = attrs["supervisor_agent"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateSupervisorAgentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateSupervisorAgentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"supervisor_agent": reflect.TypeOf(SupervisorAgent_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateSupervisorAgentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateSupervisorAgentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"supervisor_agent": m.SupervisorAgent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateSupervisorAgentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"supervisor_agent": basetypes.ListType{
				ElemType: SupervisorAgent_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSupervisorAgent returns the value of the SupervisorAgent field in CreateSupervisorAgentRequest_SdkV2 as
// a SupervisorAgent_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateSupervisorAgentRequest_SdkV2) GetSupervisorAgent(ctx context.Context) (SupervisorAgent_SdkV2, bool) {
	var e SupervisorAgent_SdkV2
	if m.SupervisorAgent.IsNull() || m.SupervisorAgent.IsUnknown() {
		return e, false
	}
	var v []SupervisorAgent_SdkV2
	d := m.SupervisorAgent.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSupervisorAgent sets the value of the SupervisorAgent field in CreateSupervisorAgentRequest_SdkV2.
func (m *CreateSupervisorAgentRequest_SdkV2) SetSupervisorAgent(ctx context.Context, v SupervisorAgent_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["supervisor_agent"]
	m.SupervisorAgent = types.ListValueMust(t, vs)
}

type CreateToolRequest_SdkV2 struct {
	// Parent resource where this tool will be created. Format:
	// supervisor-agents/{supervisor_agent_id}
	Parent types.String `tfsdk:"-"`

	Tool types.List `tfsdk:"tool"`
	// The ID to use for the tool, which will become the final component of the
	// tool's resource name.
	ToolId types.String `tfsdk:"-"`
}

func (to *CreateToolRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateToolRequest_SdkV2) {
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

func (to *CreateToolRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateToolRequest_SdkV2) {
	if !from.Tool.IsNull() && !from.Tool.IsUnknown() {
		if toTool, ok := to.GetTool(ctx); ok {
			if fromTool, ok := from.GetTool(ctx); ok {
				toTool.SyncFieldsDuringRead(ctx, fromTool)
				to.SetTool(ctx, toTool)
			}
		}
	}
}

func (m CreateToolRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["tool"] = attrs["tool"].SetRequired()
	attrs["tool"] = attrs["tool"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateToolRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tool": reflect.TypeOf(Tool_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateToolRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateToolRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parent":  m.Parent,
			"tool":    m.Tool,
			"tool_id": m.ToolId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateToolRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent": types.StringType,
			"tool": basetypes.ListType{
				ElemType: Tool_SdkV2{}.Type(ctx),
			},
			"tool_id": types.StringType,
		},
	}
}

// GetTool returns the value of the Tool field in CreateToolRequest_SdkV2 as
// a Tool_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateToolRequest_SdkV2) GetTool(ctx context.Context) (Tool_SdkV2, bool) {
	var e Tool_SdkV2
	if m.Tool.IsNull() || m.Tool.IsUnknown() {
		return e, false
	}
	var v []Tool_SdkV2
	d := m.Tool.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTool sets the value of the Tool field in CreateToolRequest_SdkV2.
func (m *CreateToolRequest_SdkV2) SetTool(ctx context.Context, v Tool_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tool"]
	m.Tool = types.ListValueMust(t, vs)
}

type DeleteSupervisorAgentRequest_SdkV2 struct {
	// The resource name of the Supervisor Agent. Format:
	// supervisor-agents/{supervisor_agent_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteSupervisorAgentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteSupervisorAgentRequest_SdkV2) {
}

func (to *DeleteSupervisorAgentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteSupervisorAgentRequest_SdkV2) {
}

func (m DeleteSupervisorAgentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteSupervisorAgentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSupervisorAgentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteSupervisorAgentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteSupervisorAgentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteToolRequest_SdkV2 struct {
	// The resource name of the Tool. Format:
	// supervisor-agents/{supervisor_agent_id}/tools/{tool_id}
	Name types.String `tfsdk:"-"`
}

func (to *DeleteToolRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteToolRequest_SdkV2) {
}

func (to *DeleteToolRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteToolRequest_SdkV2) {
}

func (m DeleteToolRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteToolRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteToolRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteToolRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteToolRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GenieSpace_SdkV2 struct {
	// The ID of the genie space.
	Id types.String `tfsdk:"id"`
}

func (to *GenieSpace_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieSpace_SdkV2) {
}

func (to *GenieSpace_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GenieSpace_SdkV2) {
}

func (m GenieSpace_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GenieSpace_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenieSpace_SdkV2
// only implements ToObjectValue() and Type().
func (m GenieSpace_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenieSpace_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetSupervisorAgentRequest_SdkV2 struct {
	// The resource name of the Supervisor Agent. Format:
	// supervisor-agents/{supervisor_agent_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetSupervisorAgentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSupervisorAgentRequest_SdkV2) {
}

func (to *GetSupervisorAgentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetSupervisorAgentRequest_SdkV2) {
}

func (m GetSupervisorAgentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetSupervisorAgentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSupervisorAgentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetSupervisorAgentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSupervisorAgentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetToolRequest_SdkV2 struct {
	// The resource name of the Tool. Format:
	// supervisor-agents/{supervisor_agent_id}/tools/{tool_id}
	Name types.String `tfsdk:"-"`
}

func (to *GetToolRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetToolRequest_SdkV2) {
}

func (to *GetToolRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetToolRequest_SdkV2) {
}

func (m GetToolRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetToolRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetToolRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetToolRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetToolRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type KnowledgeAssistant_SdkV2 struct {
	// The ID of the knowledge assistant.
	KnowledgeAssistantId types.String `tfsdk:"knowledge_assistant_id"`
	// Deprecated: use knowledge_assistant_id instead.
	ServingEndpointName types.String `tfsdk:"serving_endpoint_name"`
}

func (to *KnowledgeAssistant_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from KnowledgeAssistant_SdkV2) {
}

func (to *KnowledgeAssistant_SdkV2) SyncFieldsDuringRead(ctx context.Context, from KnowledgeAssistant_SdkV2) {
}

func (m KnowledgeAssistant_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"knowledge_assistant_id": m.KnowledgeAssistantId,
			"serving_endpoint_name":  m.ServingEndpointName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m KnowledgeAssistant_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"knowledge_assistant_id": types.StringType,
			"serving_endpoint_name":  types.StringType,
		},
	}
}

type ListSupervisorAgentsRequest_SdkV2 struct {
	// The maximum number of supervisor agents to return. If unspecified, at
	// most 100 supervisor agents will be returned. The maximum value is 100;
	// values above 100 will be coerced to 100.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListSupervisorAgents` call.
	// Provide this to retrieve the subsequent page. If unspecified, the first
	// page will be returned.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListSupervisorAgentsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSupervisorAgentsRequest_SdkV2) {
}

func (to *ListSupervisorAgentsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSupervisorAgentsRequest_SdkV2) {
}

func (m ListSupervisorAgentsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListSupervisorAgentsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSupervisorAgentsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListSupervisorAgentsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSupervisorAgentsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListSupervisorAgentsResponse_SdkV2 struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`

	SupervisorAgents types.List `tfsdk:"supervisor_agents"`
}

func (to *ListSupervisorAgentsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSupervisorAgentsResponse_SdkV2) {
	if !from.SupervisorAgents.IsNull() && !from.SupervisorAgents.IsUnknown() && to.SupervisorAgents.IsNull() && len(from.SupervisorAgents.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SupervisorAgents, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SupervisorAgents = from.SupervisorAgents
	}
}

func (to *ListSupervisorAgentsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSupervisorAgentsResponse_SdkV2) {
	if !from.SupervisorAgents.IsNull() && !from.SupervisorAgents.IsUnknown() && to.SupervisorAgents.IsNull() && len(from.SupervisorAgents.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SupervisorAgents, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SupervisorAgents = from.SupervisorAgents
	}
}

func (m ListSupervisorAgentsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListSupervisorAgentsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"supervisor_agents": reflect.TypeOf(SupervisorAgent_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSupervisorAgentsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListSupervisorAgentsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":   m.NextPageToken,
			"supervisor_agents": m.SupervisorAgents,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSupervisorAgentsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"supervisor_agents": basetypes.ListType{
				ElemType: SupervisorAgent_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSupervisorAgents returns the value of the SupervisorAgents field in ListSupervisorAgentsResponse_SdkV2 as
// a slice of SupervisorAgent_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListSupervisorAgentsResponse_SdkV2) GetSupervisorAgents(ctx context.Context) ([]SupervisorAgent_SdkV2, bool) {
	if m.SupervisorAgents.IsNull() || m.SupervisorAgents.IsUnknown() {
		return nil, false
	}
	var v []SupervisorAgent_SdkV2
	d := m.SupervisorAgents.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSupervisorAgents sets the value of the SupervisorAgents field in ListSupervisorAgentsResponse_SdkV2.
func (m *ListSupervisorAgentsResponse_SdkV2) SetSupervisorAgents(ctx context.Context, v []SupervisorAgent_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["supervisor_agents"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SupervisorAgents = types.ListValueMust(t, vs)
}

type ListToolsRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
	// Parent resource to list from. Format:
	// supervisor-agents/{supervisor_agent_id}
	Parent types.String `tfsdk:"-"`
}

func (to *ListToolsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListToolsRequest_SdkV2) {
}

func (to *ListToolsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListToolsRequest_SdkV2) {
}

func (m ListToolsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListToolsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListToolsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListToolsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListToolsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

type ListToolsResponse_SdkV2 struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Tools types.List `tfsdk:"tools"`
}

func (to *ListToolsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListToolsResponse_SdkV2) {
	if !from.Tools.IsNull() && !from.Tools.IsUnknown() && to.Tools.IsNull() && len(from.Tools.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tools, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tools = from.Tools
	}
}

func (to *ListToolsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListToolsResponse_SdkV2) {
	if !from.Tools.IsNull() && !from.Tools.IsUnknown() && to.Tools.IsNull() && len(from.Tools.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tools, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tools = from.Tools
	}
}

func (m ListToolsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListToolsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tools": reflect.TypeOf(Tool_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListToolsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListToolsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"tools":           m.Tools,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListToolsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"tools": basetypes.ListType{
				ElemType: Tool_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTools returns the value of the Tools field in ListToolsResponse_SdkV2 as
// a slice of Tool_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListToolsResponse_SdkV2) GetTools(ctx context.Context) ([]Tool_SdkV2, bool) {
	if m.Tools.IsNull() || m.Tools.IsUnknown() {
		return nil, false
	}
	var v []Tool_SdkV2
	d := m.Tools.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTools sets the value of the Tools field in ListToolsResponse_SdkV2.
func (m *ListToolsResponse_SdkV2) SetTools(ctx context.Context, v []Tool_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tools"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tools = types.ListValueMust(t, vs)
}

type SupervisorAgent_SdkV2 struct {
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

func (to *SupervisorAgent_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SupervisorAgent_SdkV2) {
}

func (to *SupervisorAgent_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SupervisorAgent_SdkV2) {
}

func (m SupervisorAgent_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["description"] = attrs["description"].SetRequired()
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
func (m SupervisorAgent_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SupervisorAgent_SdkV2
// only implements ToObjectValue() and Type().
func (m SupervisorAgent_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m SupervisorAgent_SdkV2) Type(ctx context.Context) attr.Type {
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

type Tool_SdkV2 struct {
	App types.List `tfsdk:"app"`

	UcConnection types.List `tfsdk:"uc_connection"`
	// Description of what this tool does (user-facing).
	Description types.String `tfsdk:"description"`

	GenieSpace types.List `tfsdk:"genie_space"`
	// Deprecated: Use tool_id instead.
	Id types.String `tfsdk:"id"`

	KnowledgeAssistant types.List `tfsdk:"knowledge_assistant"`
	// Full resource name:
	// supervisor-agents/{supervisor_agent_id}/tools/{tool_id}
	Name types.String `tfsdk:"name"`
	// User specified id of the Tool.
	ToolId types.String `tfsdk:"tool_id"`
	// Tool type. Must be one of: "genie_space", "knowledge_assistant",
	// "uc_function", "connection", "app", "volume", "lakeview_dashboard",
	// "serving_endpoint", "uc_table", "vector_search_index".
	ToolType types.String `tfsdk:"tool_type"`

	UcFunction types.List `tfsdk:"uc_function"`

	Volume types.List `tfsdk:"volume"`
}

func (to *Tool_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Tool_SdkV2) {
	if !from.App.IsNull() && !from.App.IsUnknown() {
		if toApp, ok := to.GetApp(ctx); ok {
			if fromApp, ok := from.GetApp(ctx); ok {
				// Recursively sync the fields of App
				toApp.SyncFieldsDuringCreateOrUpdate(ctx, fromApp)
				to.SetApp(ctx, toApp)
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

func (to *Tool_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Tool_SdkV2) {
	if !from.App.IsNull() && !from.App.IsUnknown() {
		if toApp, ok := to.GetApp(ctx); ok {
			if fromApp, ok := from.GetApp(ctx); ok {
				toApp.SyncFieldsDuringRead(ctx, fromApp)
				to.SetApp(ctx, toApp)
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

func (m Tool_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app"] = attrs["app"].SetOptional()
	attrs["app"] = attrs["app"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["uc_connection"] = attrs["uc_connection"].SetOptional()
	attrs["uc_connection"] = attrs["uc_connection"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["description"] = attrs["description"].SetRequired()
	attrs["genie_space"] = attrs["genie_space"].SetOptional()
	attrs["genie_space"] = attrs["genie_space"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetComputed()
	attrs["knowledge_assistant"] = attrs["knowledge_assistant"].SetOptional()
	attrs["knowledge_assistant"] = attrs["knowledge_assistant"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetOptional()
	attrs["tool_id"] = attrs["tool_id"].SetComputed()
	attrs["tool_type"] = attrs["tool_type"].SetRequired()
	attrs["uc_function"] = attrs["uc_function"].SetOptional()
	attrs["uc_function"] = attrs["uc_function"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["volume"] = attrs["volume"].SetOptional()
	attrs["volume"] = attrs["volume"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Tool.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Tool_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"app":                 reflect.TypeOf(App_SdkV2{}),
		"uc_connection":       reflect.TypeOf(Connection_SdkV2{}),
		"genie_space":         reflect.TypeOf(GenieSpace_SdkV2{}),
		"knowledge_assistant": reflect.TypeOf(KnowledgeAssistant_SdkV2{}),
		"uc_function":         reflect.TypeOf(UcFunction_SdkV2{}),
		"volume":              reflect.TypeOf(Volume_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Tool_SdkV2
// only implements ToObjectValue() and Type().
func (m Tool_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app":                 m.App,
			"uc_connection":       m.UcConnection,
			"description":         m.Description,
			"genie_space":         m.GenieSpace,
			"id":                  m.Id,
			"knowledge_assistant": m.KnowledgeAssistant,
			"name":                m.Name,
			"tool_id":             m.ToolId,
			"tool_type":           m.ToolType,
			"uc_function":         m.UcFunction,
			"volume":              m.Volume,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Tool_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app": basetypes.ListType{
				ElemType: App_SdkV2{}.Type(ctx),
			},
			"uc_connection": basetypes.ListType{
				ElemType: Connection_SdkV2{}.Type(ctx),
			},
			"description": types.StringType,
			"genie_space": basetypes.ListType{
				ElemType: GenieSpace_SdkV2{}.Type(ctx),
			},
			"id": types.StringType,
			"knowledge_assistant": basetypes.ListType{
				ElemType: KnowledgeAssistant_SdkV2{}.Type(ctx),
			},
			"name":      types.StringType,
			"tool_id":   types.StringType,
			"tool_type": types.StringType,
			"uc_function": basetypes.ListType{
				ElemType: UcFunction_SdkV2{}.Type(ctx),
			},
			"volume": basetypes.ListType{
				ElemType: Volume_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetApp returns the value of the App field in Tool_SdkV2 as
// a App_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Tool_SdkV2) GetApp(ctx context.Context) (App_SdkV2, bool) {
	var e App_SdkV2
	if m.App.IsNull() || m.App.IsUnknown() {
		return e, false
	}
	var v []App_SdkV2
	d := m.App.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetApp sets the value of the App field in Tool_SdkV2.
func (m *Tool_SdkV2) SetApp(ctx context.Context, v App_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["app"]
	m.App = types.ListValueMust(t, vs)
}

// GetUcConnection returns the value of the UcConnection field in Tool_SdkV2 as
// a Connection_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Tool_SdkV2) GetUcConnection(ctx context.Context) (Connection_SdkV2, bool) {
	var e Connection_SdkV2
	if m.UcConnection.IsNull() || m.UcConnection.IsUnknown() {
		return e, false
	}
	var v []Connection_SdkV2
	d := m.UcConnection.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUcConnection sets the value of the UcConnection field in Tool_SdkV2.
func (m *Tool_SdkV2) SetUcConnection(ctx context.Context, v Connection_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["uc_connection"]
	m.UcConnection = types.ListValueMust(t, vs)
}

// GetGenieSpace returns the value of the GenieSpace field in Tool_SdkV2 as
// a GenieSpace_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Tool_SdkV2) GetGenieSpace(ctx context.Context) (GenieSpace_SdkV2, bool) {
	var e GenieSpace_SdkV2
	if m.GenieSpace.IsNull() || m.GenieSpace.IsUnknown() {
		return e, false
	}
	var v []GenieSpace_SdkV2
	d := m.GenieSpace.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGenieSpace sets the value of the GenieSpace field in Tool_SdkV2.
func (m *Tool_SdkV2) SetGenieSpace(ctx context.Context, v GenieSpace_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["genie_space"]
	m.GenieSpace = types.ListValueMust(t, vs)
}

// GetKnowledgeAssistant returns the value of the KnowledgeAssistant field in Tool_SdkV2 as
// a KnowledgeAssistant_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Tool_SdkV2) GetKnowledgeAssistant(ctx context.Context) (KnowledgeAssistant_SdkV2, bool) {
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

// SetKnowledgeAssistant sets the value of the KnowledgeAssistant field in Tool_SdkV2.
func (m *Tool_SdkV2) SetKnowledgeAssistant(ctx context.Context, v KnowledgeAssistant_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["knowledge_assistant"]
	m.KnowledgeAssistant = types.ListValueMust(t, vs)
}

// GetUcFunction returns the value of the UcFunction field in Tool_SdkV2 as
// a UcFunction_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Tool_SdkV2) GetUcFunction(ctx context.Context) (UcFunction_SdkV2, bool) {
	var e UcFunction_SdkV2
	if m.UcFunction.IsNull() || m.UcFunction.IsUnknown() {
		return e, false
	}
	var v []UcFunction_SdkV2
	d := m.UcFunction.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUcFunction sets the value of the UcFunction field in Tool_SdkV2.
func (m *Tool_SdkV2) SetUcFunction(ctx context.Context, v UcFunction_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["uc_function"]
	m.UcFunction = types.ListValueMust(t, vs)
}

// GetVolume returns the value of the Volume field in Tool_SdkV2 as
// a Volume_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Tool_SdkV2) GetVolume(ctx context.Context) (Volume_SdkV2, bool) {
	var e Volume_SdkV2
	if m.Volume.IsNull() || m.Volume.IsUnknown() {
		return e, false
	}
	var v []Volume_SdkV2
	d := m.Volume.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVolume sets the value of the Volume field in Tool_SdkV2.
func (m *Tool_SdkV2) SetVolume(ctx context.Context, v Volume_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["volume"]
	m.Volume = types.ListValueMust(t, vs)
}

type UcFunction_SdkV2 struct {
	// Full uc function name
	Name types.String `tfsdk:"name"`
}

func (to *UcFunction_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UcFunction_SdkV2) {
}

func (to *UcFunction_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UcFunction_SdkV2) {
}

func (m UcFunction_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UcFunction_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UcFunction_SdkV2
// only implements ToObjectValue() and Type().
func (m UcFunction_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UcFunction_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type UpdateSupervisorAgentRequest_SdkV2 struct {
	// The resource name of the SupervisorAgent. Format:
	// supervisor-agents/{supervisor_agent_id}
	Name types.String `tfsdk:"-"`
	// The SupervisorAgent to update.
	SupervisorAgent types.List `tfsdk:"supervisor_agent"`
	// Field mask for fields to be updated.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateSupervisorAgentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateSupervisorAgentRequest_SdkV2) {
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

func (to *UpdateSupervisorAgentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateSupervisorAgentRequest_SdkV2) {
	if !from.SupervisorAgent.IsNull() && !from.SupervisorAgent.IsUnknown() {
		if toSupervisorAgent, ok := to.GetSupervisorAgent(ctx); ok {
			if fromSupervisorAgent, ok := from.GetSupervisorAgent(ctx); ok {
				toSupervisorAgent.SyncFieldsDuringRead(ctx, fromSupervisorAgent)
				to.SetSupervisorAgent(ctx, toSupervisorAgent)
			}
		}
	}
}

func (m UpdateSupervisorAgentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["supervisor_agent"] = attrs["supervisor_agent"].SetRequired()
	attrs["supervisor_agent"] = attrs["supervisor_agent"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateSupervisorAgentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"supervisor_agent": reflect.TypeOf(SupervisorAgent_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSupervisorAgentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateSupervisorAgentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":             m.Name,
			"supervisor_agent": m.SupervisorAgent,
			"update_mask":      m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateSupervisorAgentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"supervisor_agent": basetypes.ListType{
				ElemType: SupervisorAgent_SdkV2{}.Type(ctx),
			},
			"update_mask": types.StringType,
		},
	}
}

// GetSupervisorAgent returns the value of the SupervisorAgent field in UpdateSupervisorAgentRequest_SdkV2 as
// a SupervisorAgent_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateSupervisorAgentRequest_SdkV2) GetSupervisorAgent(ctx context.Context) (SupervisorAgent_SdkV2, bool) {
	var e SupervisorAgent_SdkV2
	if m.SupervisorAgent.IsNull() || m.SupervisorAgent.IsUnknown() {
		return e, false
	}
	var v []SupervisorAgent_SdkV2
	d := m.SupervisorAgent.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSupervisorAgent sets the value of the SupervisorAgent field in UpdateSupervisorAgentRequest_SdkV2.
func (m *UpdateSupervisorAgentRequest_SdkV2) SetSupervisorAgent(ctx context.Context, v SupervisorAgent_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["supervisor_agent"]
	m.SupervisorAgent = types.ListValueMust(t, vs)
}

type UpdateToolRequest_SdkV2 struct {
	// Full resource name:
	// supervisor-agents/{supervisor_agent_id}/tools/{tool_id}
	Name types.String `tfsdk:"-"`
	// The Tool to update.
	Tool types.List `tfsdk:"tool"`
	// Field mask for fields to be updated.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateToolRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateToolRequest_SdkV2) {
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

func (to *UpdateToolRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateToolRequest_SdkV2) {
	if !from.Tool.IsNull() && !from.Tool.IsUnknown() {
		if toTool, ok := to.GetTool(ctx); ok {
			if fromTool, ok := from.GetTool(ctx); ok {
				toTool.SyncFieldsDuringRead(ctx, fromTool)
				to.SetTool(ctx, toTool)
			}
		}
	}
}

func (m UpdateToolRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["tool"] = attrs["tool"].SetRequired()
	attrs["tool"] = attrs["tool"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateToolRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tool": reflect.TypeOf(Tool_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateToolRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateToolRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":        m.Name,
			"tool":        m.Tool,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateToolRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"tool": basetypes.ListType{
				ElemType: Tool_SdkV2{}.Type(ctx),
			},
			"update_mask": types.StringType,
		},
	}
}

// GetTool returns the value of the Tool field in UpdateToolRequest_SdkV2 as
// a Tool_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateToolRequest_SdkV2) GetTool(ctx context.Context) (Tool_SdkV2, bool) {
	var e Tool_SdkV2
	if m.Tool.IsNull() || m.Tool.IsUnknown() {
		return e, false
	}
	var v []Tool_SdkV2
	d := m.Tool.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTool sets the value of the Tool field in UpdateToolRequest_SdkV2.
func (m *UpdateToolRequest_SdkV2) SetTool(ctx context.Context, v Tool_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tool"]
	m.Tool = types.ListValueMust(t, vs)
}

type Volume_SdkV2 struct {
	// Full uc volume name
	Name types.String `tfsdk:"name"`
}

func (to *Volume_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Volume_SdkV2) {
}

func (to *Volume_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Volume_SdkV2) {
}

func (m Volume_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Volume_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Volume_SdkV2
// only implements ToObjectValue() and Type().
func (m Volume_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Volume_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}
