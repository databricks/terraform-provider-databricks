// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// Test service for Idempotency of Operations
package idempotencytesting

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
)

type IdempotencyTestingInterface interface {
	CreateTestResource(ctx context.Context, request CreateTestResourceRequest) (*TestResource, error)
}

func NewIdempotencyTesting(client *client.DatabricksClient) *IdempotencyTestingAPI {
	return &IdempotencyTestingAPI{
		idempotencyTestingImpl: idempotencyTestingImpl{
			client: client,
		},
	}
}

// Test service for Idempotency of Operations
type IdempotencyTestingAPI struct {
	idempotencyTestingImpl
}
