package compute

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/retries"
)

type clustersAPIUtilities interface {
	EnsureClusterIsRunning(ctx context.Context, clusterId string) error
	GetOrCreateRunningCluster(ctx context.Context, name string, custom ...CreateCluster) (c *ClusterDetails, err error)
	SelectSparkVersion(ctx context.Context, r SparkVersionRequest) (string, error)
	SelectNodeType(ctx context.Context, r NodeTypeRequest) (string, error)
}

// getOrCreateClusterMutex guards "mounting" cluster creation to prevent multiple
// redundant instances created at the same name. Compute package private property.
// https://github.com/databricks/terraform-provider-databricks/issues/445
var getOrCreateClusterMutex sync.Mutex

func (c *ClusterDetails) IsRunningOrResizing() bool {
	return c.State == StateRunning || c.State == StateResizing
}

// use mutex around starting a cluster by ID
var mu sync.Mutex

func (a *ClustersAPI) isErrFailedToReach(err error) bool {
	if err == nil {
		return false
	}
	// TODO: get a bit better handling of these
	return strings.HasPrefix(err.Error(), "failed to reach")
}

func (a *ClustersAPI) startClusterIfNeeded(ctx context.Context, clusterId string, timeout time.Duration) error {
	info, err := a.GetByClusterId(ctx, clusterId)
	if err != nil {
		return fmt.Errorf("get cluster info: %w", err)
	}
	switch info.State {
	case StateRunning:
		return nil
	case StateTerminating:
		_, err = a.WaitGetClusterTerminated(ctx, clusterId, timeout, nil)
		if err != nil {
			return fmt.Errorf("terminating -> terminated: %w", err)
		}
		_, err = a.StartByClusterIdAndWait(ctx, clusterId)
		if err != nil {
			return fmt.Errorf("terminated -> terminating -> start: %w", err)
		}
		return nil
	case StateTerminated:
		_, err = a.StartByClusterIdAndWait(ctx, clusterId)
		if err != nil {
			return fmt.Errorf("terminated -> start: %w", err)
		}
		return nil
	case StatePending, StateResizing, StateRestarting:
		_, err = a.WaitGetClusterRunning(ctx, clusterId, timeout, nil)
		if err != nil {
			return fmt.Errorf("wait: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("cluster %s is in %s state: %s", info.ClusterName, info.State, info.StateMessage)
	}
}

func (a *ClustersAPI) EnsureClusterIsRunning(ctx context.Context, clusterId string) error {
	mu.Lock()
	defer mu.Unlock()
	timeout := 20 * time.Minute
	return retries.Wait(ctx, timeout, func() *retries.Err {
		err := a.startClusterIfNeeded(ctx, clusterId, timeout)
		var apiErr *apierr.APIError
		if errors.As(err, &apiErr) && apiErr.ErrorCode == "INVALID_STATE" {
			logger.Debugf(ctx, "Cluster was started by other process: %s Retrying.", apiErr.Message)
			return retries.Continue(apiErr)
		}
		if a.isErrFailedToReach(err) {
			return retries.Continue(err)
		}
		if err != nil {
			return retries.Halt(err)
		}
		return nil
	})
}

// GetOrCreateRunningCluster creates an autoterminating cluster if it doesn't exist
func (a *ClustersAPI) GetOrCreateRunningCluster(ctx context.Context, name string, custom ...CreateCluster) (c *ClusterDetails, err error) {
	getOrCreateClusterMutex.Lock()
	defer getOrCreateClusterMutex.Unlock()
	if len(custom) > 1 {
		err = fmt.Errorf("you can only specify 1 custom cluster conf, not %d", len(custom))
		return
	}
	clusters, err := a.ListAll(ctx, ListClustersRequest{})
	if err != nil {
		return
	}
	for _, cl := range clusters {
		if cl.ClusterName != name {
			continue
		}
		logger.Infof(ctx, "Found reusable cluster '%s'", name)
		if cl.IsRunningOrResizing() {
			return &cl, nil
		}
		started, err := a.StartByClusterIdAndWait(ctx, cl.ClusterId)
		if err != nil {
			logger.Infof(ctx, "Cluster %s cannot be started, creating an autoterminating cluster: %s", name, err)
			break
		}
		return started, nil
	}
	nodeTypes, err := a.ListNodeTypes(ctx)
	if err != nil {
		return nil, fmt.Errorf("list node types: %w", err)
	}
	smallestNodeType, err := nodeTypes.Smallest(NodeTypeRequest{
		LocalDisk: true,
	})
	if err != nil {
		return nil, err
	}
	logger.Infof(ctx, "Creating an autoterminating cluster with node type %s", smallestNodeType)
	versions, err := a.SparkVersions(ctx)
	if err != nil {
		return nil, fmt.Errorf("list spark versions: %w", err)
	}
	version, err := versions.Select(SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})
	if err != nil {
		return nil, err
	}
	r := CreateCluster{
		NumWorkers:             1,
		ClusterName:            name,
		SparkVersion:           version,
		NodeTypeId:             smallestNodeType,
		AutoterminationMinutes: 10,
	}
	if len(custom) == 1 {
		r = custom[0]
	}
	return a.CreateAndWait(ctx, r)
}
