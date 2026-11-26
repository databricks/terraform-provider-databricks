// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package agentbricks

import (
	"context"
)

// The Custom LLMs service manages state and powers the UI for the Custom LLM
// product.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AgentBricksService interface {

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
