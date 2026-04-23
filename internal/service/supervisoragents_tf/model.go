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

// Databricks connection. Supported connection: external mcp server.
type Connection struct {
	Name types.String `tfsdk:"name"`
}

func (to *Connection) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Connection) {
}

func (to *Connection) SyncFieldsDuringRead(ctx context.Context, from Connection) {
}

func (m Connection) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Connection) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Connection
// only implements ToObjectValue() and Type().
func (m Connection) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Connection) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
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

type GenieSpace struct {
	// The ID of the genie space.
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

type Tool struct {
	App types.Object `tfsdk:"app"`

	UcConnection types.Object `tfsdk:"uc_connection"`
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
	// "uc_function", "connection", "app", "volume", "lakeview_dashboard",
	// "serving_endpoint", "uc_table", "vector_search_index".
	ToolType types.String `tfsdk:"tool_type"`

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

func (to *Tool) SyncFieldsDuringRead(ctx context.Context, from Tool) {
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

func (m Tool) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app"] = attrs["app"].SetOptional()
	attrs["uc_connection"] = attrs["uc_connection"].SetOptional()
	attrs["description"] = attrs["description"].SetRequired()
	attrs["genie_space"] = attrs["genie_space"].SetOptional()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["knowledge_assistant"] = attrs["knowledge_assistant"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["tool_id"] = attrs["tool_id"].SetComputed()
	attrs["tool_type"] = attrs["tool_type"].SetRequired()
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
		"uc_connection":       reflect.TypeOf(Connection{}),
		"genie_space":         reflect.TypeOf(GenieSpace{}),
		"knowledge_assistant": reflect.TypeOf(KnowledgeAssistant{}),
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
func (m Tool) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app":                 App{}.Type(ctx),
			"uc_connection":       Connection{}.Type(ctx),
			"description":         types.StringType,
			"genie_space":         GenieSpace{}.Type(ctx),
			"id":                  types.StringType,
			"knowledge_assistant": KnowledgeAssistant{}.Type(ctx),
			"name":                types.StringType,
			"tool_id":             types.StringType,
			"tool_type":           types.StringType,
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

// GetUcConnection returns the value of the UcConnection field in Tool as
// a Connection value.
// If the field is unknown or null, the boolean return value is false.
func (m *Tool) GetUcConnection(ctx context.Context) (Connection, bool) {
	var e Connection
	if m.UcConnection.IsNull() || m.UcConnection.IsUnknown() {
		return e, false
	}
	var v Connection
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
func (m *Tool) SetUcConnection(ctx context.Context, v Connection) {
	vs := v.ToObjectValue(ctx)
	m.UcConnection = vs
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
