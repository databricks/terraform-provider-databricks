// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package genieagents

import (
	"context"
)

// Genie provides a no-code experience for business users, powered by AI/BI.
// Analysts set up Genie Agents that business users can use to ask questions
// using natural language. Genie uses data registered to Unity Catalog and
// requires at least CAN USE permission on a Pro or Serverless SQL warehouse.
// Also, Databricks Assistant must be enabled.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type GenieAgentsService interface {

	// Creates a Genie Agent.
	CreateGenieAgent(ctx context.Context, request CreateGenieAgentRequest) (*GenieAgent, error)

	// Deletes a Genie Agent, moving it to trash.
	DeleteGenieAgent(ctx context.Context, request DeleteGenieAgentRequest) error

	// Gets a Genie Agent.
	GetGenieAgent(ctx context.Context, request GetGenieAgentRequest) (*GenieAgent, error)

	// Lists the Genie Agents in the workspace. Genie Agents created from an
	// integration (for example, from a dashboard) are not listed.
	ListGenieAgents(ctx context.Context, request ListGenieAgentsRequest) (*ListGenieAgentsResponse, error)

	// Updates a Genie Agent. Only the fields named in update_mask are applied.
	UpdateGenieAgent(ctx context.Context, request UpdateGenieAgentRequest) (*GenieAgent, error)
}
