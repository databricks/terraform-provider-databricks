package app

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/apps"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWaitForAppDeleted_AppDoesNotExist(t *testing.T) {
	ctx := context.Background()
	mockClient := mocks.NewMockWorkspaceClient(t)
	mockAppsAPI := mockClient.GetMockAppsAPI()

	mockAppsAPI.EXPECT().GetByName(mock.Anything, "test-app").Return(nil, &apierr.APIError{
		StatusCode: 404,
		Message:    "App not found",
	})

	err := waitForAppDeleted(ctx, mockClient.WorkspaceClient, "test-app")
	assert.NoError(t, err)
}

func TestWaitForAppDeleted_AppInDeletingState(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	mockClient := mocks.NewMockWorkspaceClient(t)
	mockAppsAPI := mockClient.GetMockAppsAPI()

	mockAppsAPI.EXPECT().GetByName(mock.Anything, "test-app").Return(&apps.App{
		Name: "test-app",
		ComputeStatus: &apps.ComputeStatus{
			State:   apps.ComputeStateDeleting,
			Message: "App is being deleted",
		},
	}, nil).Once()

	mockAppsAPI.EXPECT().GetByName(mock.Anything, "test-app").Return(nil, &apierr.APIError{
		StatusCode: 404,
		Message:    "App not found",
	}).Run(func(_ context.Context, _ string) {
		cancel()
	}).Once()

	err := waitForAppDeleted(ctx, mockClient.WorkspaceClient, "test-app")
	assert.NoError(t, err)
}

func TestWaitForAppDeleted_AppNotInDeletingState_ReturnsImmediately(t *testing.T) {
	ctx := context.Background()
	mockClient := mocks.NewMockWorkspaceClient(t)
	mockAppsAPI := mockClient.GetMockAppsAPI()

	// If app exists but is not in DELETING state, return immediately.
	// The subsequent Create() call will fail with a proper API error.
	mockAppsAPI.EXPECT().GetByName(mock.Anything, "test-app").Return(&apps.App{
		Name: "test-app",
		ComputeStatus: &apps.ComputeStatus{
			State:   apps.ComputeStateActive,
			Message: "App is active",
		},
	}, nil).Once()

	err := waitForAppDeleted(ctx, mockClient.WorkspaceClient, "test-app")
	assert.NoError(t, err)
}

func TestWaitForAppDeleted_APIError(t *testing.T) {
	ctx := context.Background()
	mockClient := mocks.NewMockWorkspaceClient(t)
	mockAppsAPI := mockClient.GetMockAppsAPI()
	mockAppsAPI.EXPECT().GetByName(mock.Anything, "test-app").Return(nil, &apierr.APIError{
		StatusCode: 500,
		Message:    "Internal server error",
	})

	err := waitForAppDeleted(ctx, mockClient.WorkspaceClient, "test-app")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Internal server error")
}
