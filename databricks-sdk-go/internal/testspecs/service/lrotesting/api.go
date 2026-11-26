// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// Test service for Long Running Operations
package lrotesting

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/experimental/api"
	"github.com/databricks/databricks-sdk-go/useragent"
)

type LroTestingInterface interface {
	CancelOperation(ctx context.Context, request CancelOperationRequest) error

	CreateTestResource(ctx context.Context, request CreateTestResourceRequest) (CreateTestResourceOperationInterface, error)

	DeleteTestResource(ctx context.Context, request DeleteTestResourceRequest) (DeleteTestResourceOperationInterface, error)

	GetOperation(ctx context.Context, request GetOperationRequest) (*Operation, error)

	// Simple method to get test resource
	GetTestResource(ctx context.Context, request GetTestResourceRequest) (*TestResource, error)
}

func NewLroTesting(client *client.DatabricksClient) *LroTestingAPI {
	return &LroTestingAPI{
		lroTestingImpl: lroTestingImpl{
			client: client,
		},
	}
}

// Test service for Long Running Operations
type LroTestingAPI struct {
	lroTestingImpl
}

func (a *LroTestingAPI) CreateTestResource(ctx context.Context, request CreateTestResourceRequest) (CreateTestResourceOperationInterface, error) {
	operation, err := a.lroTestingImpl.CreateTestResource(ctx, request)
	if err != nil {
		return nil, err
	}
	return &createTestResourceOperation{
		impl:      &a.lroTestingImpl,
		operation: operation,
	}, nil
}

type CreateTestResourceOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) (*TestResource, error)

	// Starts asynchronous cancellation on a long-running operation. The server
	// makes a best effort to cancel the operation, but success is not guaranteed.
	Cancel(ctx context.Context) error

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*TestResourceOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type createTestResourceOperation struct {
	impl      *lroTestingImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *createTestResourceOperation) Wait(ctx context.Context, opts ...api.Option) (*TestResource, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")
	var result *TestResource
	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
			Name: a.operation.Name,
		})
		if err != nil {
			return err
		}

		// Update local operation state
		a.operation = operation

		if !operation.Done {
			return errOperationInProgress
		}

		if operation.Error != nil {
			var errorMsg string
			if operation.Error.Message != "" {
				errorMsg = operation.Error.Message
			} else {
				errorMsg = "unknown error"
			}

			if operation.Error.ErrorCode != "" {
				errorMsg = fmt.Sprintf("[%s] %s", operation.Error.ErrorCode, errorMsg)
			}

			return fmt.Errorf("operation failed: %s", errorMsg)
		}

		// Operation completed successfully, unmarshal response
		if operation.Response == nil {
			return fmt.Errorf("operation completed but no response available")
		}

		var testResource TestResource
		err = json.Unmarshal(operation.Response, &testResource)
		if err != nil {
			return fmt.Errorf("failed to unmarshal testResource response: %w", err)
		}

		result = &testResource

		return nil
	}

	// Create a retrier that retries on errOperationInProgress with exponential backoff.
	retrier := api.RetryOn(api.BackoffPolicy{}, func(err error) bool {
		return errors.Is(err, errOperationInProgress)
	})

	// Add default retrier.
	defaultOpts := []api.Option{
		api.WithRetrier(func() api.Retrier { return retrier }),
	}
	allOpts := append(defaultOpts, opts...)

	err := api.Execute(ctx, call, allOpts...)

	if err != nil {
		return nil, err
	}
	return result, nil

}

// Starts asynchronous cancellation on a long-running operation. The server
// makes a best effort to cancel the operation, but success is not guaranteed.
func (a *createTestResourceOperation) Cancel(ctx context.Context) error {
	return a.impl.CancelOperation(ctx, CancelOperationRequest{
		Name: a.operation.Name,
	})
}

// Name returns the name of the long-running operation. The name is assigned
// by the server and is unique within the service from which the operation is created.
func (a *createTestResourceOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *createTestResourceOperation) Metadata() (*TestResourceOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata TestResourceOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *createTestResourceOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

func (a *LroTestingAPI) DeleteTestResource(ctx context.Context, request DeleteTestResourceRequest) (DeleteTestResourceOperationInterface, error) {
	operation, err := a.lroTestingImpl.DeleteTestResource(ctx, request)
	if err != nil {
		return nil, err
	}
	return &deleteTestResourceOperation{
		impl:      &a.lroTestingImpl,
		operation: operation,
	}, nil
}

type DeleteTestResourceOperationInterface interface {

	// Wait blocks until the long-running operation is completed. If no timeout is
	// specified, this will poll indefinitely. If a timeout is provided and the operation
	// didn't finish within the timeout, this function will return an error, otherwise
	// returns successful response and any errors encountered.
	Wait(ctx context.Context, opts ...api.Option) error

	// Starts asynchronous cancellation on a long-running operation. The server
	// makes a best effort to cancel the operation, but success is not guaranteed.
	Cancel(ctx context.Context) error

	// Name returns the name of the long-running operation. The name is assigned
	// by the server and is unique within the service from which the operation is created.
	Name() string

	// Metadata returns metadata associated with the long-running operation.
	// If the metadata is not available, the returned metadata and error are both nil.
	Metadata() (*TestResourceOperationMetadata, error)

	// Done reports whether the long-running operation has completed.
	Done() (bool, error)
}

type deleteTestResourceOperation struct {
	impl      *lroTestingImpl
	operation *Operation
}

// Wait blocks until the long-running operation is completed. If no timeout is
// specified, this will poll indefinitely. If a timeout is provided and the operation
// didn't finish within the timeout, this function will return an error, otherwise
// returns successful response and any errors encountered.
func (a *deleteTestResourceOperation) Wait(ctx context.Context, opts ...api.Option) error {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")

	errOperationInProgress := errors.New("operation still in progress")

	call := func(ctx context.Context) error {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
			Name: a.operation.Name,
		})
		if err != nil {
			return err
		}

		// Update local operation state
		a.operation = operation

		if !operation.Done {
			return errOperationInProgress
		}

		if operation.Error != nil {
			var errorMsg string
			if operation.Error.Message != "" {
				errorMsg = operation.Error.Message
			} else {
				errorMsg = "unknown error"
			}

			if operation.Error.ErrorCode != "" {
				errorMsg = fmt.Sprintf("[%s] %s", operation.Error.ErrorCode, errorMsg)
			}

			return fmt.Errorf("operation failed: %s", errorMsg)
		}

		// Operation completed successfully, unmarshal response
		if operation.Response == nil {
			return fmt.Errorf("operation completed but no response available")
		}

		return nil
	}

	// Create a retrier that retries on errOperationInProgress with exponential backoff.
	retrier := api.RetryOn(api.BackoffPolicy{}, func(err error) bool {
		return errors.Is(err, errOperationInProgress)
	})

	// Add default retrier.
	defaultOpts := []api.Option{
		api.WithRetrier(func() api.Retrier { return retrier }),
	}
	allOpts := append(defaultOpts, opts...)

	err := api.Execute(ctx, call, allOpts...)

	return err

}

// Starts asynchronous cancellation on a long-running operation. The server
// makes a best effort to cancel the operation, but success is not guaranteed.
func (a *deleteTestResourceOperation) Cancel(ctx context.Context) error {
	return a.impl.CancelOperation(ctx, CancelOperationRequest{
		Name: a.operation.Name,
	})
}

// Name returns the name of the long-running operation. The name is assigned
// by the server and is unique within the service from which the operation is created.
func (a *deleteTestResourceOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *deleteTestResourceOperation) Metadata() (*TestResourceOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata TestResourceOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *deleteTestResourceOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}
