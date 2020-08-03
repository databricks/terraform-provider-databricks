package service

import (
	"net/http"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/stretchr/testify/assert"
)

func TestLibrariesAPI_Create(t *testing.T) {
	type args struct {
		ClusterID string          `json:"cluster_id,omitempty"`
		Libraries []model.Library `json:"libraries,omitempty"`
	}

	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		wantErr        bool
	}{
		{
			name:           "Create test",
			response:       "",
			responseStatus: http.StatusOK,
			args: args{
				ClusterID: "my-cluster-id",
				Libraries: []model.Library{
					{
						Whl: "dbfs:/my/dbfs/wheel.whl",
					},
				},
			},
			wantErr: false,
		},
		{
			name:           "Create faulure test",
			response:       "",
			responseStatus: http.StatusBadRequest,
			args: args{
				ClusterID: "my-cluster-id",
				Libraries: []model.Library{
					{
						Whl: "dbfs:/my/dbfs/wheel.whl",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/libraries/install",
				&input, tt.response, tt.responseStatus, nil, tt.wantErr, func(client DatabricksClient) (interface{}, error) {
					return nil, client.Libraries().Install(model.ClusterLibraryList{
						ClusterID: tt.args.ClusterID,
						Libraries: tt.args.Libraries,
					})
				})
		})
	}
}

func TestLibrariesAPI_Delete(t *testing.T) {
	type args struct {
		ClusterID string          `json:"cluster_id,omitempty"`
		Libraries []model.Library `json:"libraries,omitempty"`
	}

	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		wantErr        bool
	}{
		{
			name:           "Delete Test",
			response:       "",
			responseStatus: http.StatusOK,
			args: args{
				ClusterID: "my-cluster-id",
				Libraries: []model.Library{
					{
						Whl: "dbfs:/my/dbfs/wheel.whl",
					},
				},
			},
			wantErr: false,
		},
		{
			name:           "Delete failure Test",
			response:       "",
			responseStatus: http.StatusBadRequest,
			args: args{
				ClusterID: "my-cluster-id",
				Libraries: []model.Library{
					{
						Whl: "dbfs:/my/dbfs/wheel.whl",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/libraries/uninstall",
				&input, tt.response, tt.responseStatus, nil, tt.wantErr, func(client DatabricksClient) (interface{}, error) {
					return nil, client.Libraries().Uninstall(model.ClusterLibraryList{
						ClusterID: tt.args.ClusterID,
						Libraries: tt.args.Libraries,
					})
				})
		})
	}
}

func TestAccLibraryCreate(t *testing.T) {
	cloud := os.Getenv("CLOUD_ENV")
	if cloud == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := CommonEnvironmentClient()
	clusterInfo, err := NewTinyClusterInCommonPool()
	assert.NoError(t, err, err)
	defer func() {
		err := client.Clusters().PermanentDelete(clusterInfo.ClusterID)
		assert.NoError(t, err, err)
	}()

	clusterID := clusterInfo.ClusterID
	libraries := []model.Library{
		{
			Pypi: &model.PyPi{
				Package: "networkx",
			},
		},
		{
			Maven: &model.Maven{
				Coordinates: "com.crealytics:spark-excel_2.12:0.13.1",
			},
		},
	}

	err = client.Libraries().Install(model.ClusterLibraryList{
		ClusterID: clusterID,
		Libraries: libraries,
	})
	assert.NoError(t, err, err)

	defer func() {
		err = client.Libraries().Uninstall(model.ClusterLibraryList{
			ClusterID: clusterID,
			Libraries: libraries,
		})
		assert.NoError(t, err, err)
	}()

	libraryStatusList, err := client.Libraries().ClusterStatus(clusterID)
	assert.NoError(t, err, err)
	assert.Equal(t, len(libraryStatusList.LibraryStatuses), len(libraries))
}
