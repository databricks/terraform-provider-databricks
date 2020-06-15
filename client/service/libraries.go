package service

import (
	"encoding/json"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// LibrariesAPI exposes the Library API
type LibrariesAPI struct {
	Client *DBApiClient
}

// Create installs the list of libraries given a cluster id
func (a LibrariesAPI) Create(clusterID string, libraries []model.Library) error {
	var libraryInstallRequest = struct {
		ClusterID string          `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
		Libraries []model.Library `json:"libraries,omitempty" url:"libraries,omitempty"`
	}{
		ClusterID: clusterID,
		Libraries: libraries,
	}

	_, err := a.Client.performQuery(http.MethodPost, "/libraries/install", "2.0", nil, libraryInstallRequest, nil)

	return err
}

// Delete deletes the list of given libraries from the cluster given the cluster id
func (a LibrariesAPI) Delete(clusterID string, libraries []model.Library) error {
	var libraryInstallRequest = struct {
		ClusterID string          `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
		Libraries []model.Library `json:"libraries,omitempty" url:"libraries,omitempty"`
	}{
		ClusterID: clusterID,
		Libraries: libraries,
	}

	_, err := a.Client.performQuery(http.MethodPost, "/libraries/uninstall", "2.0", nil, libraryInstallRequest, nil)

	return err
}

// List lists all the libraries given a cluster id
func (a LibrariesAPI) List(clusterID string) ([]model.LibraryStatus, error) {
	var libraryStatusListResp struct {
		ClusterID       string                `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
		LibraryStatuses []model.LibraryStatus `json:"library_statuses,omitempty" url:"libraries,omitempty"`
	}
	var libraryInstallRequest = struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		ClusterID: clusterID,
	}

	resp, err := a.Client.performQuery(http.MethodGet, "/libraries/cluster-status", "2.0", nil, libraryInstallRequest, nil)
	if err != nil {
		return libraryStatusListResp.LibraryStatuses, err
	}

	err = json.Unmarshal(resp, &libraryStatusListResp)

	return libraryStatusListResp.LibraryStatuses, err
}
