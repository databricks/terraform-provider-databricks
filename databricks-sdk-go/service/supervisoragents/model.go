// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package supervisoragents

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

// Databricks app. Supported app: custom mcp, custom agent.
type App struct {
	// App name
	Name string `json:"name"`
}

func (s *App) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateExampleRequest struct {
	// The example to create under the parent Supervisor Agent.
	Example Example `json:"example"`
	// Parent resource where this example will be created. Format:
	// supervisor-agents/{supervisor_agent_id}
	Parent string `json:"-" url:"-"`
}

func (s *CreateExampleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateSupervisorAgentRequest struct {
	// The Supervisor Agent to create.
	SupervisorAgent SupervisorAgent `json:"supervisor_agent"`
}

func (s *CreateSupervisorAgentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateToolRequest struct {
	// Parent resource where this tool will be created. Format:
	// supervisor-agents/{supervisor_agent_id}
	Parent string `json:"-" url:"-"`

	Tool Tool `json:"tool"`
	// The ID to use for the tool, which will become the final component of the
	// tool's resource name.
	ToolId string `json:"-" url:"tool_id"`
}

func (s *CreateToolRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteExampleRequest struct {
	// The resource name of the example to delete. Format:
	// supervisor-agents/{supervisor_agent_id}/examples/{example_id}
	Name string `json:"-" url:"-"`
}

func (s *DeleteExampleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteSupervisorAgentRequest struct {
	// The resource name of the Supervisor Agent. Format:
	// supervisor-agents/{supervisor_agent_id}
	Name string `json:"-" url:"-"`
}

func (s *DeleteSupervisorAgentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteToolRequest struct {
	// The resource name of the Tool. Format:
	// supervisor-agents/{supervisor_agent_id}/tools/{tool_id}
	Name string `json:"-" url:"-"`
}

func (s *DeleteToolRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// An example associated with a Supervisor Agent. Contains a question and
// guidelines for how the agent should respond.
type Example struct {
	// The universally unique identifier (UUID) of the example.
	ExampleId string `json:"example_id,omitempty"`
	// Guidelines for answering the question.
	Guidelines []string `json:"guidelines"`
	// Full resource name:
	// supervisor-agents/{supervisor_agent_id}/examples/{example_id}
	Name string `json:"name,omitempty"`
	// The example question.
	Question string `json:"question"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Example) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Example) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieSpace struct {
	// Deprecated: use space_id instead. Still REQUIRED for backward
	// compatibility until a future API version removes it.
	Id string `json:"id"`
}

func (s *GenieSpace) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetExampleRequest struct {
	// The resource name of the example. Format:
	// supervisor-agents/{supervisor_agent_id}/examples/{example_id}
	Name string `json:"-" url:"-"`
}

func (s *GetExampleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetSupervisorAgentPermissionLevelsRequest struct {
	// The supervisor agent for which to get or manage permissions.
	SupervisorAgentId string `json:"-" url:"-"`
}

func (s *GetSupervisorAgentPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetSupervisorAgentPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []SupervisorAgentPermissionsDescription `json:"permission_levels,omitempty"`
}

func (s *GetSupervisorAgentPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetSupervisorAgentPermissionsRequest struct {
	// The supervisor agent for which to get or manage permissions.
	SupervisorAgentId string `json:"-" url:"-"`
}

func (s *GetSupervisorAgentPermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetSupervisorAgentRequest struct {
	// The resource name of the Supervisor Agent. Format:
	// supervisor-agents/{supervisor_agent_id}
	Name string `json:"-" url:"-"`
}

func (s *GetSupervisorAgentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetToolRequest struct {
	// The resource name of the Tool. Format:
	// supervisor-agents/{supervisor_agent_id}/tools/{tool_id}
	Name string `json:"-" url:"-"`
}

func (s *GetToolRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type KnowledgeAssistant struct {
	// The ID of the knowledge assistant.
	KnowledgeAssistantId string `json:"knowledge_assistant_id"`
	// Deprecated: use knowledge_assistant_id instead.
	ServingEndpointName string `json:"serving_endpoint_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *KnowledgeAssistant) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s KnowledgeAssistant) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExamplesRequest struct {
	// The maximum number of examples to return. If unspecified, at most 100
	// examples will be returned. The maximum value is 100; values above 100
	// will be coerced to 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListExamples` call. Provide this
	// to retrieve the subsequent page. If unspecified, the first page will be
	// returned.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Parent resource to list from. Format:
	// supervisor-agents/{supervisor_agent_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExamplesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExamplesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A list of Supervisor Agent examples.
type ListExamplesResponse struct {
	Examples []Example `json:"examples,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExamplesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExamplesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSupervisorAgentsRequest struct {
	// The maximum number of supervisor agents to return. If unspecified, at
	// most 100 supervisor agents will be returned. The maximum value is 100;
	// values above 100 will be coerced to 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListSupervisorAgents` call.
	// Provide this to retrieve the subsequent page. If unspecified, the first
	// page will be returned.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSupervisorAgentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSupervisorAgentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSupervisorAgentsResponse struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`

	SupervisorAgents []SupervisorAgent `json:"supervisor_agents,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSupervisorAgentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSupervisorAgentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListToolsRequest struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`
	// Parent resource to list from. Format:
	// supervisor-agents/{supervisor_agent_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListToolsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListToolsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListToolsResponse struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	Tools []Tool `json:"tools,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListToolsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListToolsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A Supervisor Agent coordinates tools and subagents, such as Genie Agents,
// Knowledge Assistants, Unity Catalog functions, MCP servers and agent
// endpoints, to answer requests that span several domains. It is served through
// its own serving endpoint. Tools and subagents are managed through the Tool
// resource. See
// https://docs.databricks.com/aws/en/generative-ai/agent-bricks/multi-agent-supervisor
// for the product documentation.
type SupervisorAgent struct {
	// The time at which the Supervisor Agent was created, as an RFC 3339
	// timestamp.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The email address of the user who created the Supervisor Agent.
	Creator string `json:"creator,omitempty"`
	// A summary of what the agent does. It is shown to users of the agent.
	Description string `json:"description,omitempty"`
	// The display name of the Supervisor Agent. It must be unique among the
	// agents in the workspace. Creating an agent with a display name that is
	// already in use fails with ALREADY_EXISTS.
	DisplayName string `json:"display_name"`
	// The name of the Model Serving endpoint that serves the Supervisor Agent.
	// Query this endpoint to chat with the agent.
	EndpointName string `json:"endpoint_name,omitempty"`
	// The ID of the MLflow experiment that is created together with the
	// Supervisor Agent.
	ExperimentId string `json:"experiment_id,omitempty"`
	// Deprecated: Use supervisor_agent_id instead. This field is no longer
	// populated.
	Id string `json:"id,omitempty"`
	// Natural-language guidelines for how the supervisor should respond and how
	// it should delegate requests to its tools and subagents.
	Instructions string `json:"instructions,omitempty"`
	// The resource name of the Supervisor Agent. Format:
	// supervisor-agents/{supervisor_agent_id} The server assigns the name when
	// the agent is created.
	Name string `json:"name,omitempty"`
	// The universally unique identifier (UUID) of the Supervisor Agent. It is
	// the last segment of `name`.
	SupervisorAgentId string `json:"supervisor_agent_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SupervisorAgent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SupervisorAgent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SupervisorAgentAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel SupervisorAgentPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SupervisorAgentAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SupervisorAgentAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SupervisorAgentAccessControlResponse struct {
	// All permissions.
	AllPermissions []SupervisorAgentPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SupervisorAgentAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SupervisorAgentAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SupervisorAgentPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel SupervisorAgentPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SupervisorAgentPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SupervisorAgentPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type SupervisorAgentPermissionLevel string

const SupervisorAgentPermissionLevelCanManage SupervisorAgentPermissionLevel = `CAN_MANAGE`

const SupervisorAgentPermissionLevelCanQuery SupervisorAgentPermissionLevel = `CAN_QUERY`

// String representation for [fmt.Print]
func (f *SupervisorAgentPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SupervisorAgentPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_QUERY`:
		*f = SupervisorAgentPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_QUERY"`, v)
	}
}

// Values returns all possible values for SupervisorAgentPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *SupervisorAgentPermissionLevel) Values() []SupervisorAgentPermissionLevel {
	return []SupervisorAgentPermissionLevel{
		SupervisorAgentPermissionLevelCanManage,
		SupervisorAgentPermissionLevelCanQuery,
	}
}

// Type always returns SupervisorAgentPermissionLevel to satisfy [pflag.Value] interface
func (f *SupervisorAgentPermissionLevel) Type() string {
	return "SupervisorAgentPermissionLevel"
}

type SupervisorAgentPermissions struct {
	AccessControlList []SupervisorAgentAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SupervisorAgentPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SupervisorAgentPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SupervisorAgentPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel SupervisorAgentPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SupervisorAgentPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SupervisorAgentPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SupervisorAgentPermissionsRequest struct {
	AccessControlList []SupervisorAgentAccessControlRequest `json:"access_control_list,omitempty"`
	// The supervisor agent for which to get or manage permissions.
	SupervisorAgentId string `json:"-" url:"-"`
}

func (s *SupervisorAgentPermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type Tool struct {
	App *App `json:"app,omitempty"`
	// Description of what this tool does (user-facing).
	Description string `json:"description,omitempty"`

	GenieSpace *GenieSpace `json:"genie_space,omitempty"`
	// Deprecated: Use tool_id instead.
	Id string `json:"id,omitempty"`

	KnowledgeAssistant *KnowledgeAssistant `json:"knowledge_assistant,omitempty"`
	// Full resource name:
	// supervisor-agents/{supervisor_agent_id}/tools/{tool_id}
	Name string `json:"name,omitempty"`
	// User specified id of the Tool.
	ToolId string `json:"tool_id,omitempty"`
	// Tool type. Must be one of: "genie_space", "knowledge_assistant",
	// "uc_function", "uc_connection", "uc_mcp", "app", "volume", "dashboard",
	// "serving_endpoint", "table", "vector_search_index", "catalog", "schema",
	// "supervisor_agent", "databricks_web_search", "skill". The legacy values
	// "lakeview_dashboard", "uc_table", and "web_search" are also accepted and
	// remain equivalent to "dashboard", "table", and "databricks_web_search"
	// respectively. The "databricks_web_search" tool_type maps to the
	// `web_search` spec field.
	ToolType string `json:"tool_type"`

	UcConnection *UcConnection `json:"uc_connection,omitempty"`

	UcFunction *UcFunction `json:"uc_function,omitempty"`

	Volume *Volume `json:"volume,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Tool) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Tool) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Databricks UC connection. Supported connection: external mcp server.
type UcConnection struct {
	Name string `json:"name"`
}

func (s *UcConnection) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UcFunction struct {
	// Full uc function name
	Name string `json:"name"`
}

func (s *UcFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateExampleRequest struct {
	Example Example `json:"example"`
	// The resource name of the example to update. Format:
	// supervisor-agents/{supervisor_agent_id}/examples/{example_id}
	Name string `json:"-" url:"-"`
	// Comma-delimited list of fields to update on the example. Allowed values:
	// `question`, `guidelines`. Examples: - `question` - `question,guidelines`
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateExampleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateSupervisorAgentRequest struct {
	// The resource name of the Supervisor Agent. Format:
	// supervisor-agents/{supervisor_agent_id} The server assigns the name when
	// the agent is created.
	Name string `json:"-" url:"-"`
	// The Supervisor Agent to update. `name` identifies the agent. The other
	// fields carry the new values for the paths listed in `update_mask`.
	SupervisorAgent SupervisorAgent `json:"supervisor_agent"`
	// Comma-separated list of the fields to update. Supported values are
	// `display_name`, `description` and `instructions`. The request fails if
	// the mask is empty or names any other field.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateSupervisorAgentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateToolRequest struct {
	// Full resource name:
	// supervisor-agents/{supervisor_agent_id}/tools/{tool_id}
	Name string `json:"-" url:"-"`
	// The Tool to update.
	Tool Tool `json:"tool"`
	// Field mask for fields to be updated.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateToolRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type Volume struct {
	// Full uc volume name
	Name string `json:"name"`
}

func (s *Volume) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}
