package clusters

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestClustersDataSource(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockClustersAPI().EXPECT()
			e.ListAll(mock.Anything, compute.ListClustersRequest{}).Return([]compute.ClusterDetails{
				{
					ClusterId: "b",
				},
				{
					ClusterId: "a",
				},
			}, nil)
		},
		Resource:    DataSourceClusters(),
		NonWritable: true,
		Read:        true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": []string{"a", "b"},
	})
}

func TestClustersDataSourceContainsName(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockClustersAPI().EXPECT()
			e.ListAll(mock.Anything, compute.ListClustersRequest{}).Return([]compute.ClusterDetails{
				{
					ClusterId:   "b",
					ClusterName: "THIS NAME",
				},
				{
					ClusterId:   "a",
					ClusterName: "that name",
				},
			}, nil)
		},
		Resource:    DataSourceClusters(),
		NonWritable: true,
		Read:        true,
		ID:          "_",
		HCL:         `cluster_name_contains = "this"`,
	}.ApplyAndExpectData(t, map[string]any{
		"ids": []string{"b"},
	})
}

func TestClustersDataSourceErrorsOut(t *testing.T) {
	client, _ := client.New(&config.Config{
		Host:                ".",
		Token:               ".",
		RetryTimeoutSeconds: 1,
		HTTPTimeoutSeconds:  1,
	})
	diag := DataSourceClusters().ToResource().ReadContext(context.Background(), nil, &common.DatabricksClient{
		DatabricksClient: client,
	})
	assert.NotNil(t, diag)
	assert.True(t, diag.HasError())
}
