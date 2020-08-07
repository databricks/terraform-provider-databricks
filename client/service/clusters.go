package service

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func (a ClustersAPI) defaultTimeout() time.Duration {
	return 5 * time.Minute
}

// ClustersAPI is a struct that contains the Databricks api client to perform queries
type ClustersAPI struct {
	client *DatabricksClient
}

// Create creates a new Spark cluster and waits till it's running
func (a ClustersAPI) Create(cluster model.Cluster) (info model.ClusterInfo, err error) {
	var ci model.ClusterID
	err = a.client.post("/clusters/create", cluster, &ci)
	if err != nil {
		return
	}
	info, err = a.waitForClusterStatus(ci.ClusterID, model.ClusterStateRunning)
	return
}

// Edit edits the configuration of a cluster to match the provided attributes and size
func (a ClustersAPI) Edit(cluster model.Cluster) (info model.ClusterInfo, err error) {
	info, err = a.Get(cluster.ClusterID)
	if err != nil {
		return info, err
	}
	switch info.State {
	case model.ClusterStateRunning, model.ClusterStateTerminated:
		// it's already running or terminated, so we're safe to edit
		break
	case model.ClusterStatePending, model.ClusterStateResizing, model.ClusterStateRestarting:
		// let's wait tiny bit, so we return RUNNING cluster info
		info, err = a.waitForClusterStatus(info.ClusterID, model.ClusterStateRunning)
		if err != nil {
			return info, err
		}
	case model.ClusterStateTerminating:
		// let it finish terminating, so it's safe to edit.
		// TERMINATED cluster info will be returned this way
		info, err = a.waitForClusterStatus(info.ClusterID, model.ClusterStateTerminated)
		if err != nil {
			return info, err
		}
	case model.ClusterStateError, model.ClusterStateUnknown:
		// we don't know what to do, so return error
		return info, fmt.Errorf("Unexpected state: %#v", info.StateMessage)
	}
	err = a.client.post("/clusters/edit", cluster, nil)
	if err != nil {
		return info, err
	}
	if info.IsRunningOrResizing() {
		// so if cluster was running, we'll start and wait again
		err = a.Start(info.ClusterID)
		if err != nil {
			return info, err
		}
	}
	// only State / ClusterID properties will be valid in this return
	return info, err
}

// ListZones returns the zones info sent by the cloud service provider
func (a ClustersAPI) ListZones() (model.ZonesInfo, error) {
	var zonesInfo model.ZonesInfo
	err := a.client.get("/clusters/list-zones", nil, &zonesInfo)
	return zonesInfo, err
}

// Start a terminated Spark cluster given its ID and wait till it's running
func (a ClustersAPI) Start(clusterID string) error {
	err := a.client.post("/clusters/start", model.ClusterID{ClusterID: clusterID}, nil)
	if err != nil {
		if !strings.Contains(err.Error(), fmt.Sprintf("Cluster %s is in unexpected state Pending.", clusterID)) {
			return err
		}
	}
	_, err = a.waitForClusterStatus(clusterID, model.ClusterStateRunning)
	return err
}

// Restart restart a Spark cluster given its ID. If the cluster is not in a RUNNING state, nothing will happen.
func (a ClustersAPI) Restart(clusterID string) error {
	return a.client.post("/clusters/restart", model.ClusterID{ClusterID: clusterID}, nil)
}

func wrapMissingClusterError(err error, id string) error {
	if err == nil {
		return nil
	}
	apiErr, ok := err.(APIError)
	if !ok {
		return err
	}
	if apiErr.IsMissing() {
		return err
	}
	// fix non-compliant error code
	if strings.Contains(apiErr.Message,
		fmt.Sprintf("Cluster %s does not exist", id)) {
		apiErr.StatusCode = 404
		return apiErr
	}
	return err
}

func (a ClustersAPI) waitForClusterStatus(clusterID string, desired model.ClusterState) (result model.ClusterInfo, err error) {
	// this tangles client with terraform more, which is inevitable
	return result, resource.Retry(a.defaultTimeout(), func() *resource.RetryError {
		clusterInfo, err := a.Get(clusterID)
		if ae, ok := err.(APIError); ok && ae.IsMissing() {
			log.Printf("[INFO] Cluster %s not found. Retrying", clusterID)
			return resource.RetryableError(err)
		}
		if err != nil {
			return resource.NonRetryableError(err)
		}
		result = clusterInfo
		log.Printf("[DEBUG] Cluster %s is %s: %s", clusterID, clusterInfo.State, clusterInfo.StateMessage)
		if clusterInfo.State == desired {
			return nil
		}
		if !clusterInfo.State.CanReach(desired) {
			docLink := "https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterclusterstate"
			return resource.NonRetryableError(fmt.Errorf(
				"%s is not able to transition from %s to %s: %s. Please see %s for more details",
				clusterID, clusterInfo.State, desired, clusterInfo.StateMessage, docLink))
		}
		return resource.RetryableError(
			fmt.Errorf("%s is %s, but has to be %s",
				clusterID, clusterInfo.State, desired))
	})
}

// Terminate terminates a Spark cluster given its ID
func (a ClustersAPI) Terminate(clusterID string) error {
	err := a.client.post("/clusters/delete", model.ClusterID{ClusterID: clusterID}, nil)
	if err != nil {
		return err
	}
	_, err = a.waitForClusterStatus(clusterID, model.ClusterStateTerminated)
	return err
}

// PermanentDelete permanently delete a cluster
func (a ClustersAPI) PermanentDelete(clusterID string) error {
	err := a.Terminate(clusterID)
	if err != nil {
		return err
	}
	r := model.ClusterID{ClusterID: clusterID}
	err = a.client.post("/clusters/permanent-delete", r, nil)
	if err == nil {
		return nil
	}
	if !strings.Contains(err.Error(), "unpin the cluster first") {
		return err
	}
	// unpin cluster if it's pinned
	err = a.Unpin(clusterID)
	if err != nil {
		return err
	}
	// and try removing it again
	return a.client.post("/clusters/permanent-delete", r, nil)
}

// Get retrieves the information for a cluster given its identifier
func (a ClustersAPI) Get(clusterID string) (ci model.ClusterInfo, err error) {
	err = wrapMissingClusterError(a.client.get("/clusters/get",
		model.ClusterID{ClusterID: clusterID}, &ci), clusterID)
	return
}

// Pin ensure that an interactive cluster configuration is retained even after a cluster has been terminated for more than 30 days
func (a ClustersAPI) Pin(clusterID string) error {
	return a.client.post("/clusters/pin", model.ClusterID{ClusterID: clusterID}, nil)
}

// Unpin allows the cluster to eventually be removed from the list returned by the List API
func (a ClustersAPI) Unpin(clusterID string) error {
	return a.client.post("/clusters/unpin", model.ClusterID{ClusterID: clusterID}, nil)
}

// List return information about all pinned clusters, currently active clusters,
// up to 70 of the most recently terminated interactive clusters in the past 30 days,
// and up to 30 of the most recently terminated job clusters in the past 30 days
func (a ClustersAPI) List() ([]model.ClusterInfo, error) {
	var clusterList = struct {
		Clusters []model.ClusterInfo `json:"clusters,omitempty" url:"clusters,omitempty"`
	}{}
	err := a.client.get("/clusters/list", nil, &clusterList)
	return clusterList.Clusters, err
}

// ListNodeTypes returns a list of supported Spark node types
func (a ClustersAPI) ListNodeTypes() ([]model.NodeType, error) {
	var nodeTypeList = struct {
		NodeTypes []model.NodeType `json:"node_types,omitempty" url:"node_types,omitempty"`
	}{}
	err := a.client.get("/clusters/list-node-types", nil, &nodeTypeList)
	return nodeTypeList.NodeTypes, err
}

// GetOrCreateRunningCluster creates an autoterminating cluster if it doesn't exist
func (a ClustersAPI) GetOrCreateRunningCluster(name string, custom ...model.Cluster) (c model.ClusterInfo, err error) {
	if len(custom) > 1 {
		err = fmt.Errorf("You can only specify 1 custom cluster conf, not %d", len(custom))
		return
	}

	clusters, err := a.List()
	if err != nil {
		return
	}
	for _, cl := range clusters {
		if cl.ClusterName == name {
			log.Printf("[INFO] Found reusable cluster '%s'", name)

			clusterAvailable := true
			if !cl.IsRunningOrResizing() {
				err = a.Start(cl.ClusterID)
				if err != nil {
					clusterAvailable = false
					log.Printf("[INFO] Cluster %s cannot be started, creating an autoterminating cluster", name)
				}
			}
			if clusterAvailable {
				return cl, nil
			}
		}
	}

	nodeTypes, err := a.ListNodeTypes()
	if err != nil {
		return
	}

	nodeType := GetSmallestNodeType(nodeTypes)

	log.Printf("[INFO] Creating an autoterminating cluster with node type %s", nodeType.NodeTypeID)

	r := model.Cluster{
		NumWorkers:             1,
		ClusterName:            name,
		SparkVersion:           CommonRuntimeVersion(),
		NodeTypeID:             nodeType.NodeTypeID,
		IdempotencyToken:       name,
		AutoterminationMinutes: 10,
	}
	if len(custom) == 1 {
		r = custom[0]
	}
	return a.Create(r)
}

// GetSmallestNodeType returns the smallest node type in a list of node types
func GetSmallestNodeType(nodeTypes []model.NodeType) model.NodeType {
	sortedNodeTypes := nodeTypes

	sort.Slice(sortedNodeTypes, func(i, j int) bool {
		if sortedNodeTypes[i].IsDeprecated != sortedNodeTypes[j].IsDeprecated {
			return !sortedNodeTypes[i].IsDeprecated
		}

		if sortedNodeTypes[i].MemoryMB != sortedNodeTypes[j].MemoryMB {
			return sortedNodeTypes[i].MemoryMB < sortedNodeTypes[j].MemoryMB
		}

		if sortedNodeTypes[i].NumCores != sortedNodeTypes[j].NumCores {
			return sortedNodeTypes[i].NumCores < sortedNodeTypes[j].NumCores
		}

		if sortedNodeTypes[i].NumGPUs != sortedNodeTypes[j].NumGPUs {
			return sortedNodeTypes[i].NumGPUs < sortedNodeTypes[j].NumGPUs
		}

		if sortedNodeTypes[i].NodeInstanceType != nil && sortedNodeTypes[j].NodeInstanceType != nil {
			if sortedNodeTypes[i].NodeInstanceType.LocalNVMeDisks != sortedNodeTypes[j].NodeInstanceType.LocalNVMeDisks {
				return sortedNodeTypes[i].NodeInstanceType.LocalNVMeDisks < sortedNodeTypes[j].NodeInstanceType.LocalNVMeDisks
			}

			if sortedNodeTypes[i].NodeInstanceType.LocalDisks != sortedNodeTypes[j].NodeInstanceType.LocalDisks {
				return sortedNodeTypes[i].NodeInstanceType.LocalDisks < sortedNodeTypes[j].NodeInstanceType.LocalDisks
			}

			if sortedNodeTypes[i].NodeInstanceType.LocalDiskSizeGB != sortedNodeTypes[j].NodeInstanceType.LocalDiskSizeGB {
				return sortedNodeTypes[i].NodeInstanceType.LocalDiskSizeGB < sortedNodeTypes[j].NodeInstanceType.LocalDiskSizeGB
			}
		}

		if sortedNodeTypes[i].NodeTypeID != sortedNodeTypes[j].NodeTypeID {
			return sortedNodeTypes[i].NodeTypeID < sortedNodeTypes[j].NodeTypeID
		}

		return sortedNodeTypes[i].InstanceTypeID < sortedNodeTypes[j].InstanceTypeID
	})

	nodeType := sortedNodeTypes[0]
	return nodeType
}
