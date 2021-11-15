package acceptance

import (
	"context"
	"os"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/clusters"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/internal/compute"
	"github.com/databrickslabs/terraform-provider-databricks/libraries"
	"github.com/stretchr/testify/assert"
)

func TestAccLibraryCreate(t *testing.T) {
	cloud := os.Getenv("CLOUD_ENV")
	if cloud == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := common.CommonEnvironmentClient()
	clusterInfo, err := compute.NewTinyClusterInCommonPool()
	assert.NoError(t, err, err)
	defer func() {
		ctx := context.Background()
		err := clusters.NewClustersAPI(ctx, client).PermanentDelete(clusterInfo.ClusterID)
		assert.NoError(t, err, err)
	}()

	clusterID := clusterInfo.ClusterID
	libs := []libraries.Library{
		{
			Pypi: &libraries.PyPi{
				Package: "networkx",
			},
		},
		{
			Maven: &libraries.Maven{
				Coordinates: "com.crealytics:spark-excel_2.12:0.13.1",
			},
		},
	}

	ctx := context.Background()
	libsAPI := libraries.NewLibrariesAPI(ctx, client)
	err = libsAPI.Install(libraries.ClusterLibraryList{
		ClusterID: clusterID,
		Libraries: libs,
	})
	assert.NoError(t, err, err)

	defer func() {
		err = libsAPI.Uninstall(libraries.ClusterLibraryList{
			ClusterID: clusterID,
			Libraries: libs,
		})
		assert.NoError(t, err, err)
	}()

	libraryStatusList, err := libsAPI.ClusterStatus(clusterID)
	assert.NoError(t, err, err)
	assert.Equal(t, len(libraryStatusList.LibraryStatuses), len(libs))
}
