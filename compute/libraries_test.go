package compute

import (
	"context"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClusterLibraryStatuses_NoNeedAllClusters(t *testing.T) {
	need, err := ClusterLibraryStatuses{
		ClusterID: "abc",
		LibraryStatuses: []LibraryStatus{
			{
				IsLibraryInstalledOnAllClusters: true,
				Status:                          "INSTALLING",
			},
		},
	}.IsRetryNeeded()
	require.NoError(t, err)
	assert.False(t, need)
}

func TestClusterLibraryStatuses_RetryingCodes(t *testing.T) {
	need, err := ClusterLibraryStatuses{
		ClusterID: "abc",
		LibraryStatuses: []LibraryStatus{
			{
				Status: "PENDING",
			},
			{
				Status: "RESOLVING",
			},
			{
				Status: "INSTALLING",
			},
			{
				Status: "INSTALLING",
			},
		},
	}.IsRetryNeeded()
	require.Error(t, err)
	assert.Equal(t, "0 libraries are ready, but there are still 4 pending", err.Error())
	assert.True(t, need)
}

func TestClusterLibraryStatuses_ReadyStatuses(t *testing.T) {
	need, err := ClusterLibraryStatuses{
		ClusterID: "abc",
		LibraryStatuses: []LibraryStatus{
			{
				Status: "INSTALLED",
			},
			{
				Status: "SKIPPED",
			},
			{
				Status: "UNINSTALL_ON_RESTART",
			},
		},
	}.IsRetryNeeded()
	require.NoError(t, err)
	assert.False(t, need)
}

func TestClusterLibraryStatuses_Errors(t *testing.T) {
	need, err := ClusterLibraryStatuses{
		ClusterID: "abc",
		LibraryStatuses: []LibraryStatus{
			{
				Status: "FAILED",
				Library: &Library{
					Whl: "a",
				},
				Messages: []string{"b"},
			},
			{
				Status: "FAILED",
				Library: &Library{
					Maven: &Maven{
						Coordinates: "a.b.c",
					},
				},
				Messages: []string{"b"},
			},
			{
				Status: "FAILED",
				Library: &Library{
					Cran: &Cran{
						Package: "a",
					},
				},
				Messages: []string{"b"},
			},
		},
	}.IsRetryNeeded()
	require.Error(t, err)
	assert.Equal(t, "library_whl[a] failed: b\nlibrary_maven[a.b.c] failed: b\nlibrary_cran[a] failed: b", err.Error())
	assert.False(t, need)
}

func TestAccLibraryCreate(t *testing.T) {
	cloud := os.Getenv("CLOUD_ENV")
	if cloud == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := common.CommonEnvironmentClient()
	clusterInfo, err := NewTinyClusterInCommonPool()
	assert.NoError(t, err, err)
	defer func() {
		ctx := context.Background()
		err := NewClustersAPI(ctx, client).PermanentDelete(clusterInfo.ClusterID)
		assert.NoError(t, err, err)
	}()

	clusterID := clusterInfo.ClusterID
	libraries := []Library{
		{
			Pypi: &PyPi{
				Package: "networkx",
			},
		},
		{
			Maven: &Maven{
				Coordinates: "com.crealytics:spark-excel_2.12:0.13.1",
			},
		},
	}

	err = NewLibrariesAPI(client).Install(ClusterLibraryList{
		ClusterID: clusterID,
		Libraries: libraries,
	})
	assert.NoError(t, err, err)

	defer func() {
		err = NewLibrariesAPI(client).Uninstall(ClusterLibraryList{
			ClusterID: clusterID,
			Libraries: libraries,
		})
		assert.NoError(t, err, err)
	}()

	libraryStatusList, err := NewLibrariesAPI(client).ClusterStatus(clusterID)
	assert.NoError(t, err, err)
	assert.Equal(t, len(libraryStatusList.LibraryStatuses), len(libraries))
}
