package mocks

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMockWorkspaceClient(t *testing.T) {
	ctx := context.Background()
	w := NewMockWorkspaceClient(t)
	mockClusters := w.GetMockClustersAPI()
	mockClusters.EXPECT().ListAll(
		ctx,
		mock.AnythingOfType("compute.ListClustersRequest"),
	).Return(
		[]compute.ClusterDetails{
			{ClusterName: "test-cluster-1"},
			{ClusterName: "test-cluster-2"},
		}, nil)

	clusters, err := listClusters(ctx, w.WorkspaceClient)

	assert.NoError(t, err)
	assert.NotEmpty(t, clusters)
	assert.Equal(t, clusters[0].ClusterName, "test-cluster-1")
	assert.Equal(t, clusters[1].ClusterName, "test-cluster-2")
}

func TestAccountWorkspaceClient(t *testing.T) {
	ctx := context.Background()
	w := NewMockAccountClient(t)
	mockUsers := w.GetMockAccountUsersV2API()
	mockUsers.EXPECT().ListAll(
		ctx,
		mock.AnythingOfType("iam.ListAccountUsersRequest"),
	).Return(
		[]iam.AccountUser{
			{DisplayName: "test-user-1"},
			{DisplayName: "test-user-2"},
		}, nil)

	users, err := listAccountUsers(ctx, w.AccountClient)

	assert.NoError(t, err)
	assert.NotEmpty(t, users)
	assert.Equal(t, users[0].DisplayName, "test-user-1")
	assert.Equal(t, users[1].DisplayName, "test-user-2")
}

func listClusters(ctx context.Context, w *databricks.WorkspaceClient) ([]compute.ClusterDetails, error) {
	return w.Clusters.ListAll(ctx, compute.ListClustersRequest{})
}

func listAccountUsers(ctx context.Context, a *databricks.AccountClient) ([]iam.AccountUser, error) {
	return a.UsersV2.ListAll(ctx, iam.ListAccountUsersRequest{})
}
