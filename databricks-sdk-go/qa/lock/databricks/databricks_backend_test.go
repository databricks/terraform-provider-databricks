package databricks

import (
	"bytes"
	"context"
	"errors"
	"io"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/qa/lock/core"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAcquireLock_NoExistingLock(t *testing.T) {
	w := mocks.NewMockWorkspaceClient(t)
	w.GetMockWorkspaceAPI().EXPECT().Download(mock.Anything, "/Shared/locks/my-lock.lock", mock.Anything).Return(nil, apierr.ErrNotFound)
	w.GetMockWorkspaceAPI().EXPECT().Upload(mock.Anything, "/Shared/locks/my-lock.lock", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	backend := &Backend{
		lockClient: w.WorkspaceClient,
		lockDir:    "/Shared/locks",
		lockName:   "my-lock",
	}

	err := backend.AcquireLock(context.Background(), &core.LockState{})
	assert.NoError(t, err)
}

func TestAcquireLock_ExistingExpiredLock(t *testing.T) {
	w := mocks.NewMockWorkspaceClient(t)
	resp := io.NopCloser(bytes.NewReader([]byte(`{"Expiry": "2021-01-01T00:00:00Z"}`)))
	w.GetMockWorkspaceAPI().EXPECT().Download(mock.Anything, "/Shared/locks/my-lock.lock", mock.Anything).Return(resp, nil)
	w.GetMockWorkspaceAPI().EXPECT().Delete(mock.Anything, workspace.Delete{Path: "/Shared/locks/my-lock.lock"}).Return(nil)
	w.GetMockWorkspaceAPI().EXPECT().Upload(mock.Anything, "/Shared/locks/my-lock.lock", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	backend := &Backend{
		lockClient: w.WorkspaceClient,
		lockDir:    "/Shared/locks",
		lockName:   "my-lock",
	}

	err := backend.AcquireLock(context.Background(), &core.LockState{})
	assert.NoError(t, err)
}

func TestAcquireLock_ExistingValidLock(t *testing.T) {
	w := mocks.NewMockWorkspaceClient(t)
	resp := io.NopCloser(bytes.NewReader([]byte(`{"Expiry": "3021-01-01T00:00:00Z", "LeaseId": "old-lease-id"}`)))
	w.GetMockWorkspaceAPI().EXPECT().Download(mock.Anything, "/Shared/locks/my-lock.lock", mock.Anything).Return(resp, nil)

	backend := &Backend{
		lockClient: w.WorkspaceClient,
		lockDir:    "/Shared/locks",
		lockName:   "my-lock",
	}

	err := backend.AcquireLock(context.Background(), &core.LockState{LeaseId: "new-lease-id"})
	assert.ErrorContains(t, err, "lock my-lock is held by another lease, current lock state: ")
}

func TestAcquireLock_ExistingValidLockHeldBySelf(t *testing.T) {
	w := mocks.NewMockWorkspaceClient(t)
	resp := io.NopCloser(bytes.NewReader([]byte(`{"Expiry": "3021-01-01T00:00:00Z", "LeaseId": "new-lease-id"}`)))
	w.GetMockWorkspaceAPI().EXPECT().Download(mock.Anything, "/Shared/locks/my-lock.lock", mock.Anything).Return(resp, nil)

	backend := &Backend{
		lockClient: w.WorkspaceClient,
		lockDir:    "/Shared/locks",
		lockName:   "my-lock",
	}

	err := backend.AcquireLock(context.Background(), &core.LockState{LeaseId: "new-lease-id"})
	assert.NoError(t, err)
}

func TestAcquireLock_Fails(t *testing.T) {
	w := mocks.NewMockWorkspaceClient(t)
	w.GetMockWorkspaceAPI().EXPECT().Download(mock.Anything, "/Shared/locks/my-lock.lock", mock.Anything).Return(nil, apierr.ErrNotFound)
	testErr := errors.New("test error")
	w.GetMockWorkspaceAPI().EXPECT().Upload(mock.Anything, "/Shared/locks/my-lock.lock", mock.Anything, mock.Anything, mock.Anything).Return(testErr)

	backend := &Backend{
		lockClient: w.WorkspaceClient,
		lockDir:    "/Shared/locks",
		lockName:   "my-lock",
	}

	err := backend.AcquireLock(context.Background(), &core.LockState{})
	assert.ErrorIs(t, err, testErr)
}

func TestRenewLock(t *testing.T) {
	w := mocks.NewMockWorkspaceClient(t)
	resp := io.NopCloser(bytes.NewReader([]byte(`{"Expiry": "3021-01-01T00:00:00Z", "LeaseId": "lease-id"}`)))
	w.GetMockWorkspaceAPI().EXPECT().Download(mock.Anything, "/Shared/locks/my-lock.lock", mock.Anything).Return(resp, nil)
	w.GetMockWorkspaceAPI().EXPECT().Upload(mock.Anything, "/Shared/locks/my-lock.lock", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	backend := &Backend{
		lockClient: w.WorkspaceClient,
		lockDir:    "/Shared/locks",
		lockName:   "my-lock",
	}

	err := backend.RenewLock(context.Background(), "lease-id")
	assert.NoError(t, err)
}

func TestRenewLock_FailWhenNoLock(t *testing.T) {
	w := mocks.NewMockWorkspaceClient(t)
	w.GetMockWorkspaceAPI().EXPECT().Download(mock.Anything, "/Shared/locks/my-lock.lock", mock.Anything).Return(nil, apierr.ErrNotFound)

	backend := &Backend{
		lockClient: w.WorkspaceClient,
		lockDir:    "/Shared/locks",
		lockName:   "my-lock",
	}

	err := backend.RenewLock(context.Background(), "lease-id")
	assert.EqualError(t, err, "lock my-lock does not exist")
}

func TestRenewLock_FailWhenInvalid(t *testing.T) {
	w := mocks.NewMockWorkspaceClient(t)
	resp := io.NopCloser(bytes.NewReader([]byte(`{"Expiry": "2021-01-01T00:00:00Z", "LeaseId": "lease-id"}`)))
	w.GetMockWorkspaceAPI().EXPECT().Download(mock.Anything, "/Shared/locks/my-lock.lock", mock.Anything).Return(resp, nil)

	backend := &Backend{
		lockClient: w.WorkspaceClient,
		lockDir:    "/Shared/locks",
		lockName:   "my-lock",
	}

	err := backend.RenewLock(context.Background(), "lease-id")
	assert.EqualError(t, err, "lease lease-id on lock my-lock expired at 2021-01-01 00:00:00 +0000 UTC")
}

func TestRenewLock_FailWhenHeldByOther(t *testing.T) {
	w := mocks.NewMockWorkspaceClient(t)
	resp := io.NopCloser(bytes.NewReader([]byte(`{"Expiry": "3021-01-01T00:00:00Z", "LeaseId": "other-lease-id"}`)))
	w.GetMockWorkspaceAPI().EXPECT().Download(mock.Anything, "/Shared/locks/my-lock.lock", mock.Anything).Return(resp, nil)

	backend := &Backend{
		lockClient: w.WorkspaceClient,
		lockDir:    "/Shared/locks",
		lockName:   "my-lock",
	}

	err := backend.RenewLock(context.Background(), "lease-id")
	assert.EqualError(t, err, "lock my-lock is held by other-lease-id, not lease-id")
}

func TestReleaseLock(t *testing.T) {
	w := mocks.NewMockWorkspaceClient(t)
	resp := io.NopCloser(bytes.NewReader([]byte(`{"Expiry": "3021-01-01T00:00:00Z", "LeaseId": "lease-id"}`)))
	w.GetMockWorkspaceAPI().EXPECT().Download(mock.Anything, "/Shared/locks/my-lock.lock", mock.Anything).Return(resp, nil)
	w.GetMockWorkspaceAPI().EXPECT().Delete(mock.Anything, workspace.Delete{Path: "/Shared/locks/my-lock.lock"}).Return(nil)

	backend := &Backend{
		lockClient: w.WorkspaceClient,
		lockDir:    "/Shared/locks",
		lockName:   "my-lock",
	}

	err := backend.ReleaseLock(context.Background(), "lease-id")
	assert.NoError(t, err)
}

func TestReleaseLock_NoLock(t *testing.T) {
	w := mocks.NewMockWorkspaceClient(t)
	w.GetMockWorkspaceAPI().EXPECT().Download(mock.Anything, "/Shared/locks/my-lock.lock", mock.Anything).Return(nil, apierr.ErrNotFound)

	backend := &Backend{
		lockClient: w.WorkspaceClient,
		lockDir:    "/Shared/locks",
		lockName:   "my-lock",
	}

	err := backend.ReleaseLock(context.Background(), "lease-id")
	assert.NoError(t, err)
}

func TestReleaseLock_FailWhenHeldByOther(t *testing.T) {
	w := mocks.NewMockWorkspaceClient(t)
	resp := io.NopCloser(bytes.NewReader([]byte(`{"Expiry": "3021-01-01T00:00:00Z", "LeaseId": "other-lease-id"}`)))
	w.GetMockWorkspaceAPI().EXPECT().Download(mock.Anything, "/Shared/locks/my-lock.lock", mock.Anything).Return(resp, nil)

	backend := &Backend{
		lockClient: w.WorkspaceClient,
		lockDir:    "/Shared/locks",
		lockName:   "my-lock",
	}

	err := backend.ReleaseLock(context.Background(), "lease-id")
	assert.EqualError(t, err, "lock my-lock is held by other-lease-id, not lease-id")
}
