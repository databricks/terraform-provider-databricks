package service

import (
	"net/http"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

func TestLibrariesAPI_Create(t *testing.T) {
	type args struct {
		ClusterID string          `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
		Libraries []model.Library `json:"libraries,omitempty" url:"libraries,omitempty"`
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
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/libraries/install", &input, tt.response, tt.responseStatus, nil, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Libraries().Create(tt.args.ClusterID, tt.args.Libraries)
			})
		})
	}
}

func TestLibrariesAPI_Delete(t *testing.T) {
	type args struct {
		ClusterID string          `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
		Libraries []model.Library `json:"libraries,omitempty" url:"libraries,omitempty"`
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
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/libraries/uninstall", &input, tt.response, tt.responseStatus, nil, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Libraries().Delete(tt.args.ClusterID, tt.args.Libraries)
			})
		})
	}
}

func TestLibrariesAPI_List(t *testing.T) {
	type args struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}
	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		wantURI        string
		want           []model.LibraryStatus
		wantErr        bool
	}{
		{
			name: "List non recursive test",
			response: `{
						  "cluster_id": "11203-my-cluster",
						  "library_statuses": [
							{
							  "library": {
								"jar": "dbfs:/mnt/libraries/library.jar"
							  },
							  "status": "INSTALLED",
							  "messages": [],
							  "is_library_for_all_clusters": false
							},
							{
							  "library": {
								"pypi": {
								  "package": "beautifulsoup4"
								}
							  },
							  "status": "INSTALLING",
							  "messages": ["Successfully resolved package from PyPI"],
							  "is_library_for_all_clusters": false
							},
							{
							  "library": {
								"cran": {
								  "package": "ada",
								  "repo": "https://cran.us.r-project.org"
								}
							  },
							  "status": "FAILED",
							  "messages": ["R package installation is not supported on this spark version.\nPlease upgrade to Runtime 3.2 or higher"],
							  "is_library_for_all_clusters": false
							}
						  ]
						}`,
			responseStatus: http.StatusOK,
			args: args{
				ClusterID: "11203-my-cluster",
			},
			wantURI: "/api/2.0/libraries/cluster-status?cluster_id=11203-my-cluster",
			want: []model.LibraryStatus{
				{
					Library: &model.Library{
						Jar: "dbfs:/mnt/libraries/library.jar",
					},
					Status:                          "INSTALLED",
					IsLibraryInstalledOnAllClusters: false,
					Messages:                        []string{},
				},
				{
					Library: &model.Library{
						Pypi: &model.PyPi{
							Package: "beautifulsoup4",
						},
					},
					Status:                          "INSTALLING",
					IsLibraryInstalledOnAllClusters: false,
					Messages:                        []string{"Successfully resolved package from PyPI"},
				},
				{
					Library: &model.Library{
						Cran: &model.Cran{
							Package: "ada",
							Repo:    "https://cran.us.r-project.org",
						},
					},
					Status:                          "FAILED",
					IsLibraryInstalledOnAllClusters: false,
					Messages:                        []string{"R package installation is not supported on this spark version.\nPlease upgrade to Runtime 3.2 or higher"},
				},
			},
			wantErr: false,
		},
		{
			name:           "List non recursive test",
			response:       ``,
			responseStatus: http.StatusBadRequest,
			args: args{
				ClusterID: "11203-my-cluster",
			},
			wantURI: "/api/2.0/libraries/cluster-status?cluster_id=11203-my-cluster",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, tt.args, http.MethodGet, tt.wantURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.Libraries().List(tt.args.ClusterID)
			})
		})
	}
}
