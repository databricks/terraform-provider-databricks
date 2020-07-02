package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// ClustersAPI is a struct that contains the Databricks api client to perform queries
type ClustersAPI struct {
	Client *DBApiClient
}

func (a ClustersAPI) defaultTimeout() time.Duration {
	return 5 * time.Minute
}

// Create creates a new Spark cluster
func (a ClustersAPI) Create(cluster model.Cluster) (model.ClusterInfo, error) {
	var clusterInfo model.ClusterInfo
	resp, err := a.Client.performQuery(http.MethodPost, "/clusters/create", "2.0", nil, cluster, nil)
	if err != nil {
		return clusterInfo, err
	}
	err = json.Unmarshal(resp, &clusterInfo)
	if err != nil {
		return clusterInfo, APIError{
			ErrorCode:  "INVALID_RESPONSE",
			Message:    fmt.Sprintf("Error reading JSON from [%v]: %v", resp, err),
			Resource:   "/clusters/create",
			StatusCode: 400,
		}
	}
	return clusterInfo, nil
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

// Start a terminated Spark cluster given its ID and wait till it's running
func (a ClustersAPI) Start(clusterID string) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/start", "2.0", nil, model.ClusterIDRequest{
		ClusterID: clusterID,
	}, nil)
	if err != nil {
		if !strings.Contains(err.Error(), fmt.Sprintf("Cluster %s is in unexpected state Pending.", clusterID)) {
			return err
		}
	}
	return a.waitForClusterStatus(clusterID, model.ClusterStateRunning)
}

// Restart restart a Spark cluster given its ID. If the cluster is not in a RUNNING state, nothing will happen.
func (a ClustersAPI) Restart(clusterID string) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/restart", "2.0", nil, model.ClusterIDRequest{
		ClusterID: clusterID,
	}, nil)
	return err
}

func isClusterMissing(err error, resourceID string) bool {
	if apiErr, ok := err.(APIError); ok {
		return apiErr.IsMissing() || strings.Contains(err.Error(),
			fmt.Sprintf("Cluster %s does not exist", resourceID))
	}
	return false
}

func (a ClustersAPI) waitForClusterStatus(clusterID string, desired model.ClusterState) error {
	// this tangles client with terraform more, which is inevitable
	return resource.Retry(a.defaultTimeout(), func() *resource.RetryError {
		clusterInfo, err := a.Get(clusterID)
		if isClusterMissing(err, clusterID) {
			log.Printf("[INFO] Cluster %s not found. Retrying", clusterID)
			return resource.RetryableError(err)
		}
		if err != nil {
			return resource.NonRetryableError(err)
		}
		log.Printf("[DEBUG] Cluster %s is %s", clusterID, clusterInfo.State)
		if clusterInfo.State == desired {
			return nil
		}
		if !clusterInfo.State.CanReach(desired) {
			docLink := "https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterclusterstate"
			return resource.NonRetryableError(fmt.Errorf(
				"%s is not able to transition from %s to %s. Please see %s for more details",
				clusterID, clusterInfo.State, desired, docLink))
		}
		return resource.RetryableError(
			fmt.Errorf("%s is %s, but has to be %s",
				clusterID, clusterInfo.State, desired))
	})
}

// WaitForClusterRunning will block main thread and wait till cluster is in a RUNNING state
func (a ClustersAPI) WaitForClusterRunning(clusterID string) error {
	return a.waitForClusterStatus(clusterID, model.ClusterStateRunning)
}

// WaitForClusterTerminated will block main thread and wait till cluster is in a TERMINATED state
func (a ClustersAPI) WaitForClusterTerminated(clusterID string) error {
	return a.waitForClusterStatus(clusterID, model.ClusterStateTerminated)
}

// Terminate terminates a Spark cluster given its ID
func (a ClustersAPI) Terminate(clusterID string) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/delete", "2.0", nil, model.ClusterIDRequest{
		ClusterID: clusterID,
	}, nil)
	if err != nil {
		return err
	}
	return a.waitForClusterStatus(clusterID, model.ClusterStateTerminated)
}

// Delete is an alias of Terminate
func (a ClustersAPI) Delete(clusterID string) error {
	return a.Terminate(clusterID)
}

// PermanentDelete will gracefully terminate and  permanently delete a cluster
func (a ClustersAPI) PermanentDelete(clusterID string) error {
	err := a.Terminate(clusterID)
	if err != nil {
		return err
	}
	_, err = a.Client.performQuery(http.MethodPost, "/clusters/permanent-delete", "2.0", nil, model.ClusterIDRequest{
		ClusterID: clusterID,
	}, nil)
	return err
}

// Get retrieves the information for a cluster given its identifier
func (a ClustersAPI) Get(clusterID string) (model.ClusterInfo, error) {
	var clusterInfo model.ClusterInfo
	resp, err := a.Client.performQuery(http.MethodGet, "/clusters/get", "2.0", nil, model.ClusterIDRequest{
		ClusterID: clusterID,
	}, nil)
	if err != nil {
		return clusterInfo, err
	}

	err = json.Unmarshal(resp, &clusterInfo)
	return clusterInfo, err
}

// Pin ensure that an interactive cluster configuration is retained even after a cluster has been terminated for more than 30 days
func (a ClustersAPI) Pin(clusterID string) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/pin", "2.0", nil, model.ClusterIDRequest{
		ClusterID: clusterID,
	}, nil)
	return err
}

// Unpin allows the cluster to eventually be removed from the list returned by the List API
func (a ClustersAPI) Unpin(clusterID string) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/unpin", "2.0", nil, model.ClusterIDRequest{
		ClusterID: clusterID,
	}, nil)
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
