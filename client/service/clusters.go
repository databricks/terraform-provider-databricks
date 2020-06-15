package service

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// ClustersAPI is a struct that contains the Databricks api client to perform queries
type ClustersAPI struct {
	Client *DBApiClient
}

// Create creates a new Spark cluster
func (a ClustersAPI) Create(cluster model.Cluster) (model.ClusterInfo, error) {
	var clusterInfo model.ClusterInfo

	resp, err := a.Client.performQuery(http.MethodPost, "/clusters/create", "2.0", nil, cluster, nil)
	if err != nil {
		return clusterInfo, err
	}

	err = json.Unmarshal(resp, &clusterInfo)
	return clusterInfo, err
}

// Edit edits the configuration of a cluster to match the provided attributes and size
func (a ClustersAPI) Edit(clusterInfo model.Cluster) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/edit", "2.0", nil, clusterInfo, nil)
	return err
}

// ListZones returns the zones info sent by the cloud service provider
func (a ClustersAPI) ListZones() (model.ZonesInfo, error) {
	var zonesInfo model.ZonesInfo
	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/list-zones", "2.0", nil, nil, nil)
	if err != nil {
		return zonesInfo, err
	}
	err = json.Unmarshal(resp, &zonesInfo)
	return zonesInfo, err
}

// Start starts a terminated Spark cluster given its ID
func (a ClustersAPI) Start(clusterID string) error {
	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		clusterID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/start", "2.0", nil, data, nil)
	return err
}

// Restart restart a Spark cluster given its ID. If the cluster is not in a RUNNING state, nothing will happen.
func (a ClustersAPI) Restart(clusterID string) error {
	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		clusterID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/restart", "2.0", nil, data, nil)
	return err
}

// WaitForClusterRunning will block main thread and wait till cluster is in a RUNNING state
func (a ClustersAPI) WaitForClusterRunning(clusterID string, sleepDurationSeconds time.Duration, timeoutDurationMinutes time.Duration) error {
	errChan := make(chan error, 1)
	go func() {
		for {
			clusterInfo, err := a.Get(clusterID)
			if err != nil {
				errChan <- err
				return
			}
			if clusterInfo.State == model.ClusterStateRunning {
				errChan <- nil
				return
			} else if model.ContainsClusterState(model.ClusterStateNonRunnable, clusterInfo.State) {
				errChan <- errors.New("Cluster is in a non runnable state will not be able to transition to running, needs " +
					"to be started again. Current state: " + string(clusterInfo.State))
				return
			}
			log.Println("Waiting for cluster to go to running, current state is: " + string(clusterInfo.State))
			time.Sleep(sleepDurationSeconds * time.Second)
		}
	}()
	select {
	case err := <-errChan:
		return err
	case <-time.After(timeoutDurationMinutes * time.Minute):
		return errors.New("Timed out cluster has not reached running state")
	}
}

// WaitForClusterTerminated will block main thread and wait till cluster is in a TERMINATED state
func (a ClustersAPI) WaitForClusterTerminated(clusterID string, sleepDurationSeconds time.Duration, timeoutDurationMinutes time.Duration) error {
	errChan := make(chan error, 1)
	go func() {
		for {
			clusterInfo, err := a.Get(clusterID)
			if err != nil {
				errChan <- err
				return
			}
			if clusterInfo.State == model.ClusterStateTerminated {
				errChan <- nil
				return
			} else if model.ContainsClusterState(model.ClusterStateNonTerminating, clusterInfo.State) {
				errChan <- errors.New("Cluster is in a non runnable state will not be able to transition to terminated, needs " +
					"to be terminated again. Current state: " + string(clusterInfo.State))
				return
			}
			log.Println("Waiting for cluster to go to terminate, current state is: " + string(clusterInfo.State))
			time.Sleep(sleepDurationSeconds * time.Second)
		}
	}()
	select {
	case err := <-errChan:
		return err
	case <-time.After(timeoutDurationMinutes * time.Minute):
		return errors.New("Timed out cluster has not reached terminated state")
	}
}

// Terminate terminates a Spark cluster given its ID
func (a ClustersAPI) Terminate(clusterID string) error {
	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		clusterID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/delete", "2.0", nil, data, nil)
	return err
}

// Delete is an alias of Terminate
func (a ClustersAPI) Delete(clusterID string) error {
	return a.Terminate(clusterID)
}

// PermanentDelete permanently delete a cluster
func (a ClustersAPI) PermanentDelete(clusterID string) error {
	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		clusterID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/permanent-delete", "2.0", nil, data, nil)
	return err
}

// Get retrieves the information for a cluster given its identifier
func (a ClustersAPI) Get(clusterID string) (model.ClusterInfo, error) {
	var clusterInfo model.ClusterInfo

	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		clusterID,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/get", "2.0", nil, data, nil)
	if err != nil {
		return clusterInfo, err
	}

	err = json.Unmarshal(resp, &clusterInfo)
	return clusterInfo, err
}

// Pin ensure that an interactive cluster configuration is retained even after a cluster has been terminated for more than 30 days
func (a ClustersAPI) Pin(clusterID string) error {
	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		clusterID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/pin", "2.0", nil, data, nil)
	return err
}

// Unpin allows the cluster to eventually be removed from the list returned by the List API
func (a ClustersAPI) Unpin(clusterID string) error {
	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		clusterID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/unpin", "2.0", nil, data, nil)
	return err
}

// List return information about all pinned clusters, currently active clusters,
// up to 70 of the most recently terminated interactive clusters in the past 30 days,
// and up to 30 of the most recently terminated job clusters in the past 30 days
func (a ClustersAPI) List() ([]model.ClusterInfo, error) {
	var clusterList = struct {
		Clusters []model.ClusterInfo `json:"clusters,omitempty" url:"clusters,omitempty"`
	}{}

	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/list", "2.0", nil, nil, nil)
	if err != nil {
		return clusterList.Clusters, err
	}

	err = json.Unmarshal(resp, &clusterList)
	return clusterList.Clusters, err
}

// ListNodeTypes returns a list of supported Spark node types
func (a ClustersAPI) ListNodeTypes() ([]model.NodeType, error) {
	var nodeTypeList = struct {
		NodeTypes []model.NodeType `json:"node_types,omitempty" url:"node_types,omitempty"`
	}{}

	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/list-node-types", "2.0", nil, nil, nil)
	if err != nil {
		return nodeTypeList.NodeTypes, err
	}

	err = json.Unmarshal(resp, &nodeTypeList)
	return nodeTypeList.NodeTypes, err
}
