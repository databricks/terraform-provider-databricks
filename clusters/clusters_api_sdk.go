package clusters

import (
	"context"
	"log"
	"time"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
)

// Start a terminated Spark cluster given its ID and wait till it's running
func StartCluster(ctx context.Context, c *common.DatabricksClient, clusterID string) error {
	_, err := StartClusterAndGetInfo(ctx, c, clusterID)
	return err
}

// StartAndGetInfo starts cluster and returns info
func StartClusterAndGetInfo(ctx context.Context, c *common.DatabricksClient, clusterID string) (*compute.ClusterDetails, error) {
	w, err := c.WorkspaceClient()
	if err != nil {
		return nil, err
	}
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
