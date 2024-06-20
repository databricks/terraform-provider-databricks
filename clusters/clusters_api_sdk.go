package clusters

import (
	"context"
	"log"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/compute"
)

// StartAndGetInfo starts cluster and returns info
func StartClusterAndGetInfo(ctx context.Context, w *databricks.WorkspaceClient, clusterID string) (*compute.ClusterDetails, error) {
	cluster, err := w.Clusters.GetByClusterId(ctx, clusterID)
	if err != nil {
		return cluster, err
	}
	switch cluster.State {
	case compute.StateRunning:
		// it's already running, so we're good to return
		return cluster, nil
	case compute.StatePending, compute.StateResizing, compute.StateRestarting:
		// let's wait tiny bit, so we return RUNNING cluster info
		return w.Clusters.WaitGetClusterRunning(ctx, clusterID, 20*time.Minute, nil)
	case compute.StateTerminating:
		// Let it finish terminating, so it's safe to start again.
		// TERMINATED cluster info will be returned this way
		info, err := w.Clusters.WaitGetClusterTerminated(ctx, clusterID, 20*time.Minute, nil)
		if err != nil {
			return info, err
		}
	case compute.StateError, compute.StateUnknown:
		// most likely we can start error'ed cluster again...
		log.Printf("[ERROR] Cluster %s: %s", cluster.State, cluster.StateMessage)
	}
	return w.Clusters.StartByClusterIdAndWait(ctx, clusterID)
}

// LatestSparkVersionOrDefault returns Spark version matching the definition, or default in case of error
func LatestSparkVersionOrDefault(ctx context.Context, w *databricks.WorkspaceClient, svr compute.SparkVersionRequest) string {
	version, err := w.Clusters.SelectSparkVersion(ctx, svr)
	if err != nil {
		return "7.3.x-scala2.12"
	}
	return version
}
