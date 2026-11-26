// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// The Custom LLMs service manages state and powers the UI for the Custom LLM
// product.
package agentbricks

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
)

type AgentBricksInterface interface {

	// Cancel a Custom LLM Optimization Run.
	CancelOptimize(ctx context.Context, request CancelCustomLlmOptimizationRunRequest) error

	// Create a Custom LLM.
	CreateCustomLlm(ctx context.Context, request CreateCustomLlmRequest) (*CustomLlm, error)

	// Delete a Custom LLM.
	DeleteCustomLlm(ctx context.Context, request DeleteCustomLlmRequest) error

	// Get a Custom LLM.
	GetCustomLlm(ctx context.Context, request GetCustomLlmRequest) (*CustomLlm, error)

	// Start a Custom LLM Optimization Run.
	StartOptimize(ctx context.Context, request StartCustomLlmOptimizationRunRequest) (*CustomLlm, error)

	// Update a Custom LLM.
	UpdateCustomLlm(ctx context.Context, request UpdateCustomLlmRequest) (*CustomLlm, error)
}

func NewAgentBricks(client *client.DatabricksClient) *AgentBricksAPI {
	return &AgentBricksAPI{
		agentBricksImpl: agentBricksImpl{
			client: client,
		},
	}
}

// The Custom LLMs service manages state and powers the UI for the Custom LLM
// product.
type AgentBricksAPI struct {
	agentBricksImpl
}
