package clusters

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestClusterDataByID(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockClustersAPI().EXPECT()
			e.GetByClusterId(mock.Anything, "abc").Return(&compute.ClusterDetails{
				ClusterId:              "abc",
				NumWorkers:             100,
				ClusterName:            "Shared Autoscaling",
				SparkVersion:           "7.1-scala12",
				NodeTypeId:             "i3.xlarge",
				AutoterminationMinutes: 15,
				State:                  ClusterStateRunning,
				Autoscale: &compute.AutoScale{
					MaxWorkers: 4,
				},
			}, nil)
		},
		Resource:    DataSourceCluster(),
		HCL:         `cluster_id = "abc"`,
		Read:        true,
		NonWritable: true,
		ID:          "abc",
	}.ApplyAndExpectData(t, map[string]any{
		"cluster_info.0.autotermination_minutes": 15,
		"cluster_info.0.cluster_name":            "Shared Autoscaling",
		"cluster_info.0.node_type_id":            "i3.xlarge",
		"cluster_info.0.autoscale.0.max_workers": 4,
		"cluster_info.0.state":                   "RUNNING",
	})
}

func TestClusterDataByName(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockClustersAPI().EXPECT()
			e.ListAll(mock.Anything, compute.ListClustersRequest{}).Return([]compute.ClusterDetails{{
				ClusterId:              "abc",
				NumWorkers:             100,
				ClusterName:            "Shared Autoscaling",
				SparkVersion:           "7.1-scala12",
				NodeTypeId:             "i3.xlarge",
				AutoterminationMinutes: 15,
				State:                  ClusterStateRunning,
				Autoscale: &compute.AutoScale{
					MaxWorkers: 4,
				},
			}}, nil)
		},
		Resource:    DataSourceCluster(),
		HCL:         `cluster_name = "Shared Autoscaling"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"cluster_info.0.autotermination_minutes": 15,
		"cluster_info.0.cluster_name":            "Shared Autoscaling",
		"cluster_info.0.node_type_id":            "i3.xlarge",
		"cluster_info.0.autoscale.0.max_workers": 4,
		"cluster_info.0.state":                   "RUNNING",
	})
}

func TestClusterDataByName_NotFound(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockClustersAPI().EXPECT()
			e.ListAll(mock.Anything, compute.ListClustersRequest{}).Return([]compute.ClusterDetails{}, nil)
		},
		Resource:    DataSourceCluster(),
		HCL:         `cluster_name = "Unknown"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "there is no cluster with name 'Unknown'")
}

func TestClusterDataByName_DuplicateNames(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockClustersAPI().EXPECT()
			e.ListAll(mock.Anything, compute.ListClustersRequest{}).Return([]compute.ClusterDetails{
				{
					ClusterId:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  ClusterStateRunning,
					Autoscale: &compute.AutoScale{
						MaxWorkers: 4,
					},
				},
				{
					ClusterId:              "def",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  ClusterStateRunning,
					Autoscale: &compute.AutoScale{
						MaxWorkers: 4,
					},
				},
			}, nil)
		},
		Resource:    DataSourceCluster(),
		HCL:         `cluster_name = "Shared Autoscaling"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "there is more than one cluster with name 'Shared Autoscaling'")
}

func TestClusterDataByName_ListError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceCluster(),
		HCL:         `cluster_name = "Unknown"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}

func TestClusterData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceCluster(),
		Read:        true,
		NonWritable: true,
		HCL:         `cluster_id = "abc"`,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}

func TestClusterData_ErrorNoParams(t *testing.T) {
	qa.ResourceFixture{
		Resource:    DataSourceCluster(),
		Read:        true,
		NonWritable: true,
		HCL:         "",
		ID:          "_",
	}.ExpectError(t, "you need to specify either `cluster_name` or `cluster_id`")
}
