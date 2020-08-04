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

// Uninstall library list from cluster
func (a LibrariesAPI) Uninstall(req model.ClusterLibraryList) error {
	return a.client.post("/libraries/uninstall", req, nil)
}

// ClusterStatus returns library status in cluster
func (a LibrariesAPI) ClusterStatus(clusterID string) (cls model.ClusterLibraryStatuses, err error) {
	err = a.client.get("/libraries/cluster-status", model.ClusterID{ClusterID: clusterID}, &cls)
	return
}
