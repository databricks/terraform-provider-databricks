package service

import (
	"github.com/databrickslabs/databricks-terraform/client/model"
)

// LibrariesAPI exposes the Library API
type LibrariesAPI struct {
	client *DatabricksClient
}

// Install library list on cluster
func (a LibrariesAPI) Install(req model.ClusterLibraryList) error {
	return a.client.post("/libraries/install", req, nil)
}

// Uninstall library list form cluster
func (a LibrariesAPI) Uninstall(req model.ClusterLibraryList) error {
	return a.client.post("/libraries/uninstall", req, nil)
}

// ClusterStatus returns library status in cluster
func (a LibrariesAPI) ClusterStatus(clusterID string) (model.ClusterLibraryStatuses, error) {
	var clusterLibraryStatuses model.ClusterLibraryStatuses
	err := a.client.get("/libraries/cluster-status", model.ClusterID{clusterID}, &clusterLibraryStatuses)
	return clusterLibraryStatuses, err
}

// Create [DEPRECATED] installs the list of libraries given a cluster id
func (a LibrariesAPI) Create(clusterID string, libraries []model.Library) error {
	return a.Install(model.ClusterLibraryList{
		ClusterID: clusterID,
		Libraries: libraries,
	})
}

// Delete [DEPRECATED] deletes the list of given libraries from the cluster given the cluster id
func (a LibrariesAPI) Delete(clusterID string, libraries []model.Library) error {
	return a.Uninstall(model.ClusterLibraryList{
		ClusterID: clusterID,
		Libraries: libraries,
	})
}

// List [DEPRECATED] lists all the libraries given a cluster id
func (a LibrariesAPI) List(clusterID string) ([]model.LibraryStatus, error) {
	cll, err := a.ClusterStatus(clusterID)
	return cll.LibraryStatuses, err // check error scenario...
}
