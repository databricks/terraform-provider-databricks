package common

import (
	"context"
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRetryOnTimeout_NoError(t *testing.T) {
	w := mocks.NewMockWorkspaceClient(t)
	expected := &workspace.ObjectInfo{}
	api := w.GetMockWorkspaceAPI().EXPECT()
	api.GetStatusByPath(mock.Anything, mock.Anything).Return(expected, nil)
	res, err := RetryOnTimeout(context.Background(), func(ctx context.Context) (*workspace.ObjectInfo, error) {
		return w.WorkspaceClient.Workspace.GetStatusByPath(ctx, "path")
	})
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestRetryOnTimeout_OneError(t *testing.T) {
	w := mocks.NewMockWorkspaceClient(t)
	expected := &workspace.ObjectInfo{}
	api := w.GetMockWorkspaceAPI().EXPECT()
	call1 := api.GetStatusByPath(mock.Anything, mock.Anything).Return(nil, errors.New("request failed: request timed out after 1m0s of inactivity"))
	call1.Repeatability = 1
	api.GetStatusByPath(mock.Anything, mock.Anything).Return(expected, nil)
	res, err := RetryOnTimeout(context.Background(), func(ctx context.Context) (*workspace.ObjectInfo, error) {
		return w.WorkspaceClient.Workspace.GetStatusByPath(ctx, "path")
	})
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestRetryOnTimeout_NonRetriableError(t *testing.T) {
	w := mocks.NewMockWorkspaceClient(t)
	expected := errors.New("request failed: non-retriable error")
	api := w.GetMockWorkspaceAPI().EXPECT()
	api.GetStatusByPath(mock.Anything, mock.Anything).Return(nil, expected)
	_, err := RetryOnTimeout(context.Background(), func(ctx context.Context) (*workspace.ObjectInfo, error) {
		return w.WorkspaceClient.Workspace.GetStatusByPath(ctx, "path")
	})
	assert.ErrorIs(t, err, expected)
}

func TestRetryOn504_noError(t *testing.T) {
	wantErr := error(nil)
	wantRes := (*workspace.ObjectInfo)(nil)
	wantCalls := 1

	w := mocks.NewMockWorkspaceClient(t)
	api := w.GetMockWorkspaceAPI().EXPECT()
	api.GetStatusByPath(mock.Anything, mock.Anything).Return(wantRes, wantErr)

	gotCalls := 0
	gotRes, gotErr := RetryOn504(context.Background(), func(ctx context.Context) (*workspace.ObjectInfo, error) {
		gotCalls += 1
		return w.WorkspaceClient.Workspace.GetStatusByPath(ctx, "path")
	})

	assert.ErrorIs(t, gotErr, wantErr)
	assert.Equal(t, gotRes, wantRes)
	assert.Equal(t, gotCalls, wantCalls)
}

func TestRetryOn504_errorNot504(t *testing.T) {
	wantErr := errors.New("test error")
	wantRes := (*workspace.ObjectInfo)(nil)
	wantCalls := 1

	w := mocks.NewMockWorkspaceClient(t)
	api := w.GetMockWorkspaceAPI().EXPECT()
	api.GetStatusByPath(mock.Anything, mock.Anything).Return(wantRes, wantErr)

	gotCalls := 0
	gotRes, gotErr := RetryOn504(context.Background(), func(ctx context.Context) (*workspace.ObjectInfo, error) {
		gotCalls += 1
		return w.WorkspaceClient.Workspace.GetStatusByPath(ctx, "path")
	})

	assert.ErrorIs(t, gotErr, wantErr)
	assert.Equal(t, gotRes, wantRes)
	assert.Equal(t, gotCalls, wantCalls)
}

func TestRetryOn504_error504ThenFail(t *testing.T) {
	wantErr := errors.New("test error")
	wantRes := (*workspace.ObjectInfo)(nil)
	wantCalls := 2

	w := mocks.NewMockWorkspaceClient(t)
	api := w.GetMockWorkspaceAPI().EXPECT()
	call := api.GetStatusByPath(mock.Anything, mock.Anything).Return(nil, apierr.ErrDeadlineExceeded)
	call.Repeatability = 1
	api.GetStatusByPath(mock.Anything, mock.Anything).Return(wantRes, wantErr)

	gotCalls := 0
	gotRes, gotErr := RetryOn504(context.Background(), func(ctx context.Context) (*workspace.ObjectInfo, error) {
		gotCalls++
		return w.WorkspaceClient.Workspace.GetStatusByPath(ctx, "path")
	})

	assert.ErrorIs(t, gotErr, wantErr)
	assert.Equal(t, gotRes, wantRes)
	assert.Equal(t, gotCalls, wantCalls)
}

func TestRetryOn504_error504ThenSuccess(t *testing.T) {
	wantErr := error(nil)
	wantRes := &workspace.ObjectInfo{}
	wantCalls := 2

	w := mocks.NewMockWorkspaceClient(t)
	api := w.GetMockWorkspaceAPI().EXPECT()
	call := api.GetStatusByPath(mock.Anything, mock.Anything).Return(nil, apierr.ErrDeadlineExceeded)
	call.Repeatability = 1
	api.GetStatusByPath(mock.Anything, mock.Anything).Return(wantRes, wantErr)

	gotCalls := 0
	gotRes, gotErr := RetryOn504(context.Background(), func(ctx context.Context) (*workspace.ObjectInfo, error) {
		gotCalls++
		return w.WorkspaceClient.Workspace.GetStatusByPath(ctx, "path")
	})

	assert.ErrorIs(t, gotErr, wantErr)
	assert.Equal(t, gotRes, wantRes)
	assert.Equal(t, gotCalls, wantCalls)
}
