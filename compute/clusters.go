package compute

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func (a ClustersAPI) defaultTimeout() time.Duration {
	return 30 * time.Minute
}

// NewClustersAPI creates ClustersAPI instance from provider meta
func NewClustersAPI(ctx context.Context, m interface{}) ClustersAPI {
	return ClustersAPI{
		client:  m.(*common.DatabricksClient),
		context: ctx,
	}
}

// ClustersAPI is a struct that contains the Databricks api client to perform queries
type ClustersAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates a new Spark cluster and waits till it's running
func (a ClustersAPI) Create(cluster Cluster) (info ClusterInfo, err error) {
	var ci ClusterID
	err = a.client.Post(a.context, "/clusters/create", cluster, &ci)
	if err != nil {
		return
	}
	info, err = a.waitForClusterStatus(ci.ClusterID, ClusterStateRunning)
	if err != nil {
		// https://github.com/databrickslabs/terraform-provider-databricks/issues/383
		log.Printf("[ERROR] Cleaning up created cluster, that failed to start: %s", err.Error())
		deleteErr := a.PermanentDelete(ci.ClusterID)
		if deleteErr != nil {
			log.Printf("[ERROR] Failed : %s", deleteErr.Error())
			err = deleteErr
		}
	}
	return
}

// Edit edits the configuration of a cluster to match the provided attributes and size
func (a ClustersAPI) Edit(cluster Cluster) (info ClusterInfo, err error) {
	info, err = a.Get(cluster.ClusterID)
	if err != nil {
		return info, err
	}
	switch info.State {
	case ClusterStateRunning, ClusterStateTerminated:
		// it's already running or terminated, so we're safe to edit
		break
	case ClusterStatePending, ClusterStateResizing, ClusterStateRestarting:
		// let's wait tiny bit, so we return RUNNING cluster info
		info, err = a.waitForClusterStatus(info.ClusterID, ClusterStateRunning)
		if err != nil {
			return info, err
		}
	case ClusterStateTerminating:
		// let it finish terminating, so it's safe to edit.
		// TERMINATED cluster info will be returned this way
		info, err = a.waitForClusterStatus(info.ClusterID, ClusterStateTerminated)
		if err != nil {
			return info, err
		}
	case ClusterStateError, ClusterStateUnknown:
		// we don't know what to do, so return error
		return info, fmt.Errorf("Unexpected state: %#v", info.StateMessage)
	}
	err = a.client.Post(a.context, "/clusters/edit", cluster, nil)
	if err != nil {
		return info, err
	}
	if info.IsRunningOrResizing() {
		// so if cluster was running, we'll start and wait again
		return a.StartAndGetInfo(info.ClusterID)
	}
	// only State / ClusterID properties will be valid in this return
	return info, err
}

// ListZones returns the zones info sent by the cloud service provider
func (a ClustersAPI) ListZones() (ZonesInfo, error) {
	var zonesInfo ZonesInfo
	err := a.client.Get(a.context, "/clusters/list-zones", nil, &zonesInfo)
	return zonesInfo, err
}

// Start a terminated Spark cluster given its ID and wait till it's running
func (a ClustersAPI) Start(clusterID string) error {
	_, err := a.StartAndGetInfo(clusterID)
	return err
}

// StartAndGetInfo starts cluster and returns info
func (a ClustersAPI) StartAndGetInfo(clusterID string) (ClusterInfo, error) {
	info, err := a.Get(clusterID)
	if err != nil {
		return info, err
	}
	switch info.State {
	case ClusterStateRunning:
		// it's already running, so we're good to return
		return info, nil
	case ClusterStatePending, ClusterStateResizing, ClusterStateRestarting:
		// let's wait tiny bit, so we return RUNNING cluster info
		return a.waitForClusterStatus(info.ClusterID, ClusterStateRunning)
	case ClusterStateTerminating:
		// let it finish terminating, so it's safe to start again.
		// TERMINATED cluster info will be returned this way
		info, err = a.waitForClusterStatus(info.ClusterID, ClusterStateTerminated)
		if err != nil {
			return info, err
		}
	case ClusterStateError, ClusterStateUnknown:
		// most likely we can start error'ed cluster again...
		log.Printf("[ERROR] Cluster %s: %s", info.State, info.StateMessage)
	}
	err = a.client.Post(a.context, "/clusters/start", ClusterID{ClusterID: clusterID}, nil)
	if err != nil {
		if !strings.Contains(err.Error(),
			fmt.Sprintf("Cluster %s is in unexpected state Pending.", clusterID)) {
			return info, err
		}
	}
	return a.waitForClusterStatus(clusterID, ClusterStateRunning)
}

// Restart restart a Spark cluster given its ID. If the cluster is not in a RUNNING state, nothing will happen.
func (a ClustersAPI) Restart(clusterID string) error {
	return a.client.Post(a.context, "/clusters/restart", ClusterID{ClusterID: clusterID}, nil)
}

func wrapMissingClusterError(err error, id string) error {
	if err == nil {
		return nil
	}
	apiErr, ok := err.(common.APIError)
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

func (a ClustersAPI) waitForClusterStatus(clusterID string, desired ClusterState) (result ClusterInfo, err error) {
	// this tangles client with terraform more, which is inevitable
	// nolint should be a bigger context-aware refactor
	return result, resource.RetryContext(a.context, a.defaultTimeout(), func() *resource.RetryError {
		clusterInfo, err := a.Get(clusterID)
		if ae, ok := err.(common.APIError); ok && ae.IsMissing() {
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
			if clusterInfo.TerminationReason != nil {
				log.Printf("[DEBUG] Cluster %s termination info: code: %s, type: %s, parameters: %v",
					clusterID, clusterInfo.TerminationReason.Code, clusterInfo.TerminationReason.Type,
					clusterInfo.TerminationReason.Parameters)
			}
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
	err := a.client.Post(a.context, "/clusters/delete", ClusterID{ClusterID: clusterID}, nil)
	if err != nil {
		return err
	}
	_, err = a.waitForClusterStatus(clusterID, ClusterStateTerminated)
	return err
}

// PermanentDelete permanently delete a cluster
func (a ClustersAPI) PermanentDelete(clusterID string) error {
	err := a.Terminate(clusterID)
	if err != nil {
		return err
	}
	r := ClusterID{ClusterID: clusterID}
	err = a.client.Post(a.context, "/clusters/permanent-delete", r, nil)
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
	return a.client.Post(a.context, "/clusters/permanent-delete", r, nil)
}

// Get retrieves the information for a cluster given its identifier
func (a ClustersAPI) Get(clusterID string) (ci ClusterInfo, err error) {
	err = wrapMissingClusterError(a.client.Get(a.context, "/clusters/get",
		ClusterID{ClusterID: clusterID}, &ci), clusterID)
	return
}

// Pin ensure that an interactive cluster configuration is retained even after a cluster has been terminated for more than 30 days
func (a ClustersAPI) Pin(clusterID string) error {
	return a.client.Post(a.context, "/clusters/pin", ClusterID{ClusterID: clusterID}, nil)
}

// Unpin allows the cluster to eventually be removed from the list returned by the List API
func (a ClustersAPI) Unpin(clusterID string) error {
	return a.client.Post(a.context, "/clusters/unpin", ClusterID{ClusterID: clusterID}, nil)
}

// Events - only using Cluster ID string to get all events
// https://docs.databricks.com/dev-tools/api/latest/clusters.html#events
func (a ClustersAPI) Events(eventsRequest EventsRequest) ([]ClusterEvent, error) {
	var eventsResponse EventsResponse
	err := a.client.Post(a.context, "/clusters/events", eventsRequest, &eventsResponse)
	if err != nil {
		return nil, err
	}

	totalCount := int(eventsResponse.TotalCount)
	if (eventsRequest.MaxItems) > 0 && (eventsRequest.MaxItems < uint(totalCount)) {
		totalCount = int(eventsRequest.MaxItems)
	}
	events := make([]ClusterEvent, totalCount)
	if totalCount == 0 {
		return events, nil
	}
	startPos := 0
	curPos := len(eventsResponse.Events)
	copy(events[startPos:curPos], eventsResponse.Events)
	for curPos < totalCount && eventsResponse.NextPage != nil {
		err := a.client.Post(a.context, "/clusters/events", eventsResponse.NextPage, &eventsResponse)
		if err != nil {
			return nil, err
		}
		startPos = curPos
		curLen := len(eventsResponse.Events)
		restItems := totalCount - startPos
		if restItems < curLen {
			curLen = restItems
		}
		curPos += curLen
		copy(events[startPos:curPos], eventsResponse.Events[0:curLen])
	}

	return events[0:curPos], err
}

// List return information about all pinned clusters, currently active clusters,
// up to 70 of the most recently terminated interactive clusters in the past 30 days,
// and up to 30 of the most recently terminated job clusters in the past 30 days
func (a ClustersAPI) List() ([]ClusterInfo, error) {
	var clusterList ClusterList
	err := a.client.Get(a.context, "/clusters/list", nil, &clusterList)
	return clusterList.Clusters, err
}

// ListNodeTypes returns a sorted list of supported Spark node types
func (a ClustersAPI) ListNodeTypes() (l NodeTypeList, err error) {
	err = a.client.Get(a.context, "/clusters/list-node-types", nil, &l)
	return
}

// getOrCreateClusterMutex guards "mounting" cluster creation to prevent multiple
// redundant instances created at the same name. Compute package private property.
// https://github.com/databrickslabs/terraform-provider-databricks/issues/445
var getOrCreateClusterMutex sync.Mutex

// GetOrCreateRunningCluster creates an autoterminating cluster if it doesn't exist
func (a ClustersAPI) GetOrCreateRunningCluster(name string, custom ...Cluster) (c ClusterInfo, err error) {
	getOrCreateClusterMutex.Lock()
	defer getOrCreateClusterMutex.Unlock()

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
	smallestNodeType := a.GetSmallestNodeType(NodeTypeRequest{
		LocalDisk: true,
	})
	log.Printf("[INFO] Creating an autoterminating cluster with node type %s", smallestNodeType)
	r := Cluster{
		NumWorkers:  1,
		ClusterName: name,
		SparkVersion: a.LatestSparkVersionOrDefault(SparkVersionRequest{
			Latest:          true,
			LongTermSupport: true,
		}),
		NodeTypeID:             smallestNodeType,
		AutoterminationMinutes: 10,
	}
	if a.client.IsAws() {
		r.AwsAttributes = &AwsAttributes{
			Availability: "SPOT",
		}
	}
	if len(custom) == 1 {
		r = custom[0]
	}
	return a.Create(r)
}

// NodeTypeRequest is a wrapper for local filtering of node types
type NodeTypeRequest struct {
	MinMemoryGB int32  `json:"min_memory_gb,omitempty"`
	GBPerCore   int32  `json:"gb_per_core,omitempty"`
	MinCores    int32  `json:"min_cores,omitempty"`
	MinGPUs     int32  `json:"min_gpus,omitempty"`
	LocalDisk   bool   `json:"local_disk,omitempty"`
	Category    string `json:"category,omitempty"`
}

func defaultSmallestNodeType(a ClustersAPI) string {
	if a.client.IsAzure() {
		return "Standard_D3_v2"
	} else if a.client.IsGcp() {
		return "n1-standard-4"
	}
	return "i3.xlarge"
}

// GetSmallestNodeType returns smallest (or default) node type id given the criteria
func (a ClustersAPI) GetSmallestNodeType(r NodeTypeRequest) string {
	list, _ := a.ListNodeTypes()
	// error is explicitly ingored here, because Azure returns
	// apparently too big of a JSON for Go to parse
	if len(list.NodeTypes) == 0 {
		return defaultSmallestNodeType(a)
	}
	list.Sort()
	for _, nt := range list.NodeTypes {
		gbs := (nt.MemoryMB / 1024)
		if r.MinMemoryGB > 0 && gbs < r.MinMemoryGB {
			continue
		}
		if r.GBPerCore > 0 && (gbs/int32(nt.NumCores)) < r.GBPerCore {
			continue
		}
		if r.MinCores > 0 && int32(nt.NumCores) < r.MinCores {
			continue
		}
		if r.MinGPUs > 0 && nt.NumGPUs < r.MinGPUs {
			continue
		}
		if r.LocalDisk && nt.NodeInstanceType != nil &&
			(nt.NodeInstanceType.LocalDisks < 1 &&
				nt.NodeInstanceType.LocalNVMeDisks < 1) {
			continue
		}
		if r.Category != "" && nt.Category != r.Category {
			continue
		}
		return nt.NodeTypeID
	}
	return defaultSmallestNodeType(a)
}

// ListSparkVersions returns smallest (or default) node type id given the criteria
func (a ClustersAPI) ListSparkVersions() (SparkVersionsList, error) {
	var sparkVersions SparkVersionsList
	err := a.client.Get(a.context, "/clusters/spark-versions", nil, &sparkVersions)
	return sparkVersions, err
}

// LatestSparkVersion returns latest version matching the request parameters
func (sparkVersions SparkVersionsList) LatestSparkVersion(req SparkVersionRequest) (string, error) {
	var versions []string

	for _, version := range sparkVersions.SparkVersions {
		if strings.Contains(version.Version, "-scala"+req.Scala) {
			matches := ((!strings.Contains(version.Version, "apache-spark-")) &&
				(strings.Contains(version.Version, "-ml-") == req.ML) &&
				(strings.Contains(version.Version, "-hls-") == req.Genomics) &&
				(strings.Contains(version.Version, "-gpu-") == req.GPU) &&
				(strings.Contains(version.Description, "Beta") == req.Beta))
			if matches && req.LongTermSupport {
				matches = (matches && strings.Contains(version.Description, "LTS"))
			}
			if matches && len(req.SparkVersion) > 0 {
				matches = (matches && strings.Contains(version.Description, "Apache Spark "+req.SparkVersion))
			}
			if matches {
				versions = append(versions, version.Version)
			}
		}
	}
	if len(versions) < 1 {
		return "", fmt.Errorf("Spark versions query returned no results. Please change your search criteria and try again")
	} else if len(versions) > 1 {
		if req.Latest {
			sort.Sort(sort.Reverse(sort.StringSlice(versions)))
		} else {
			return "", fmt.Errorf("Spark versions query returned multiple results. Please change your search criteria and try again")
		}
	}

	return versions[0], nil
}

// LatestSparkVersion returns latest version matching the request parameters
func (a ClustersAPI) LatestSparkVersion(svr SparkVersionRequest) (string, error) {
	sparkVersions, err := a.ListSparkVersions()
	if err != nil {
		return "", err
	}
	return sparkVersions.LatestSparkVersion(svr)
}

// LatestSparkVersionOrDefault returns Spark version matching the definition, or default in case of error
func (a ClustersAPI) LatestSparkVersionOrDefault(svr SparkVersionRequest) string {
	version, err := a.LatestSparkVersion(svr)
	if err != nil {
		return "7.3.x-scala2.12"
	}
	return version
}
