package service

import (
	"encoding/json"
	"errors"
	"github.com/stikkireddy/databricks-tf-provider/client/model"
	"log"
	"net/http"
	"time"
)

type ClustersAPI struct {
	Client DBApiClient
}

func (a ClustersAPI) init(client DBApiClient) ClustersAPI {
	a.Client = client
	return a
}

// Create creates a new Spark cluster
func (a ClustersAPI) Create(cluster model.Cluster) (model.ClusterInfo, error) {
	var clusterInfo model.ClusterInfo

	resp, err := a.Client.performQuery(http.MethodPost, "/clusters/create", "2.0", nil, cluster)
	if err != nil {
		return clusterInfo, err
	}

	err = json.Unmarshal(resp, &clusterInfo)
	return clusterInfo, err
}

// Update edits the configuration of a cluster to match the provided attributes and size
func (a ClustersAPI) Edit(clusterInfo model.ClusterInfo) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/edit", "2.0", nil, clusterInfo)
	return err
}

// Start starts a terminated Spark cluster given its ID
func (a ClustersAPI) Start(clusterID string) error {
	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		clusterID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/start", "2.0", nil, data)
	return err
}

// Restart restart a Spark cluster given its ID. If the cluster is not in a RUNNING state, nothing will happen.
func (a ClustersAPI) Restart(clusterID string) error {
	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		clusterID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/restart", "2.0", nil, data)
	return err
}

func (a ClustersAPI) WaitForClusterRunning(clusterID string, sleepDurationSeconds time.Duration, timeoutDurationMinutes time.Duration) error {
	errChan := make(chan error, 1)
	go func() {
		for {
			clusterInfo, err := a.Get(clusterID)
			if err != nil {
				errChan <- err
			}
			if clusterInfo.State == model.ClusterStateRunning {
				errChan <- nil
			} else if model.ContainsClusterState(model.ClusterStateNonRunnable, clusterInfo.State) {
				errChan <- errors.New("Cluster is in a non runnable state will not be able to transition to running, needs " +
					"to be started again. Current state: " + string(clusterInfo.State))
			}
			log.Println("Waiting for cluster to go to running, current state is: " + string(clusterInfo.State))
			time.Sleep(sleepDurationSeconds * time.Second)
		}
	}()
	select {
	case err := <-errChan:
		return err
	case <-time.After(timeoutDurationMinutes * time.Minute):
		defer a.Start(clusterID)
		return errors.New("Timed out cluster has not reached running state")
	}
}

func (a ClustersAPI) WaitForClusterTerminated(clusterID string, sleepDurationSeconds time.Duration, timeoutDurationMinutes time.Duration) error {
	errChan := make(chan error, 1)
	go func() {
		for {
			clusterInfo, err := a.Get(clusterID)
			if err != nil {
				errChan <- err
			}
			if clusterInfo.State == model.ClusterStateTerminated {
				errChan <- nil
			} else if model.ContainsClusterState(model.ClusterStateNonTerminating, clusterInfo.State) {
				errChan <- errors.New("Cluster is in a non runnable state will not be able to transition to terminated, needs " +
					"to be terminated again. Current state: " + string(clusterInfo.State))
			}
			log.Println("Waiting for cluster to go to terminate, current state is: " + string(clusterInfo.State))
			time.Sleep(sleepDurationSeconds * time.Second)
		}
	}()
	select {
	case err := <-errChan:
		return err
	case <-time.After(timeoutDurationMinutes * time.Minute):
		defer a.Delete(clusterID)
		return errors.New("Timed out cluster has not reached terminated state")
	}
}

//// Resize resizes a cluster to have a desired number of workers. This will fail unless the cluster is in a RUNNING state.
//func (a ClustersAPI) Resize(clusterID string, clusterSize models.ClusterSize) error {
//	data := struct {
//		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
//		models.ClusterSize
//	}{
//		clusterID,
//		clusterSize,
//	}
//	_, err := a.Client.performQuery(http.MethodPost, "/clusters/resize", data, nil)
//	return err
//}

// Terminate terminates a Spark cluster given its ID
func (a ClustersAPI) Terminate(clusterID string) error {
	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		clusterID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/delete", "2.0", nil, data)
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
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/permanent-delete", "2.0", nil, data)
	return err
}

// Read retrieves the information for a cluster given its identifier
func (a ClustersAPI) Get(clusterID string) (model.ClusterInfo, error) {
	var clusterInfo model.ClusterInfo

	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		clusterID,
	}
	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/get", "2.0", nil, data)
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
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/pin", "2.0", nil, data)
	return err
}

// Unpin allows the cluster to eventually be removed from the list returned by the List API
func (a ClustersAPI) Unpin(clusterID string) error {
	data := struct {
		ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	}{
		clusterID,
	}
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/unpin", "2.0", nil, data)
	return err
}

// List return information about all pinned clusters, currently active clusters,
// up to 70 of the most recently terminated interactive clusters in the past 30 days,
// and up to 30 of the most recently terminated job clusters in the past 30 days
func (a ClustersAPI) List() ([]model.ClusterInfo, error) {
	var clusterList = struct {
		Clusters []model.ClusterInfo `json:"clusters,omitempty" url:"clusters,omitempty"`
	}{}

	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/list", "2.0", nil, nil)
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

	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/list-node-types", "2.0", nil, nil)
	if err != nil {
		return nodeTypeList.NodeTypes, err
	}

	err = json.Unmarshal(resp, &nodeTypeList)
	return nodeTypeList.NodeTypes, err
}

//// SparkVersions return the list of available Spark versions
//func (a ClustersAPI) SparkVersions() ([]model.SparkVersion, error) {
//	var versionsList = struct {
//		Versions []models.SparkVersion `json:"versions,omitempty" url:"versions,omitempty"`
//	}{}
//
//	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/spark-versions", nil, nil)
//	if err != nil {
//		return versionsList.Versions, err
//	}
//
//	err = json.Unmarshal(resp, &versionsList)
//	return versionsList.Versions, err
//}

// ClustersListZonesResponse is the response from ListZones
//type ClustersListZonesResponse struct {
//	Zones       []string `json:"zones,omitempty" url:"zones,omitempty"`
//	DefaultZone string   `json:"default_zone,omitempty" url:"default_zone,omitempty"`
//}
//
//// ListZones returns a list of availability zones where clusters can be created in (ex: us-west-2a)
//func (a ClustersAPI) ListZones() (ClustersListZonesResponse, error) {
//	var zonesList ClustersListZonesResponse
//
//	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/list-zones", nil, nil)
//	if err != nil {
//		return zonesList, err
//	}
//
//	err = json.Unmarshal(resp, &zonesList)
//	return zonesList, err
//}
