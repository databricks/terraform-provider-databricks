// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package supervisoragents

import (
	"context"
)

// Manage legacy Supervisor Agents and related resources.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type SupervisorAgentsService interface {

	// Creates an example for a Supervisor Agent.
	CreateExample(ctx context.Context, request CreateExampleRequest) (*Example, error)

	// Creates a Supervisor Agent together with its serving endpoint and MLflow
	// experiment. Tools and subagents are added to the agent afterwards with
	// the Tool methods.
	CreateSupervisorAgent(ctx context.Context, request CreateSupervisorAgentRequest) (*SupervisorAgent, error)

	// Creates a Tool under a Supervisor Agent. Specify one of "genie_space",
	// "knowledge_assistant", "uc_function", "uc_connection", "app", "volume",
	// "dashboard", "table", "vector_search_index", "catalog", "schema",
	// "supervisor_agent", "databricks_web_search", "skill" in the request body.
	// The legacy values "lakeview_dashboard", "uc_table", and "web_search" are
	// also accepted and remain equivalent to "dashboard", "table", and
	// "databricks_web_search" respectively. The "databricks_web_search"
	// tool_type maps to the `web_search` spec field.
	CreateTool(ctx context.Context, request CreateToolRequest) (*Tool, error)

	// Deletes an example from a Supervisor Agent.
	DeleteExample(ctx context.Context, request DeleteExampleRequest) error

	// Deletes a Supervisor Agent together with the resources created for it,
	// including its tools, examples, serving endpoint and MLflow experiment.
	DeleteSupervisorAgent(ctx context.Context, request DeleteSupervisorAgentRequest) error

	// Deletes a Tool.
	DeleteTool(ctx context.Context, request DeleteToolRequest) error

	// Gets an example from a Supervisor Agent.
	GetExample(ctx context.Context, request GetExampleRequest) (*Example, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetSupervisorAgentPermissionLevelsRequest) (*GetSupervisorAgentPermissionLevelsResponse, error)

	// Gets the permissions of a supervisor agent. Supervisor agents can inherit
	// permissions from their root object.
	GetPermissions(ctx context.Context, request GetSupervisorAgentPermissionsRequest) (*SupervisorAgentPermissions, error)

	// Gets a Supervisor Agent.
	GetSupervisorAgent(ctx context.Context, request GetSupervisorAgentRequest) (*SupervisorAgent, error)

	// Gets a Tool.
	GetTool(ctx context.Context, request GetToolRequest) (*Tool, error)

	// Lists examples under a Supervisor Agent.
	ListExamples(ctx context.Context, request ListExamplesRequest) (*ListExamplesResponse, error)

	// Lists the Supervisor Agents in the workspace that the caller has
	// permission to view.
	ListSupervisorAgents(ctx context.Context, request ListSupervisorAgentsRequest) (*ListSupervisorAgentsResponse, error)

	// Lists Tools under a Supervisor Agent.
	ListTools(ctx context.Context, request ListToolsRequest) (*ListToolsResponse, error)

	// Sets permissions on an object, replacing existing permissions if they
	// exist. Deletes all direct permissions if none are specified. Objects can
	// inherit permissions from their root object.
	SetPermissions(ctx context.Context, request SupervisorAgentPermissionsRequest) (*SupervisorAgentPermissions, error)

	// Updates an example in a Supervisor Agent.
	UpdateExample(ctx context.Context, request UpdateExampleRequest) (*Example, error)

	// Updates the permissions on a supervisor agent. Supervisor agents can
	// inherit permissions from their root object.
	UpdatePermissions(ctx context.Context, request SupervisorAgentPermissionsRequest) (*SupervisorAgentPermissions, error)

	// Updates a Supervisor Agent. Only the fields listed in `update_mask` are
	// changed. All other fields keep their current values.
	UpdateSupervisorAgent(ctx context.Context, request UpdateSupervisorAgentRequest) (*SupervisorAgent, error)

	// Updates a Tool. Only the `description` field can be updated. To change
	// immutable fields such as tool type, spec, or tool ID, delete the tool and
	// recreate it.
	UpdateTool(ctx context.Context, request UpdateToolRequest) (*Tool, error)
}
