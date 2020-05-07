package databricks

import (
	"fmt"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"strings"
	"time"
)

func changeClusterIntoRunningState(clusterID string, client service.DBApiClient) error {
	//return nil
	clusterInfo, err := client.Clusters().Get(clusterID)
	if err != nil {
		return err
	}
	currentState := clusterInfo.State

	if model.ContainsClusterState([]model.ClusterState{model.ClusterStateRunning}, currentState) {
		time.Sleep(5 * time.Second)
		return nil
	}

	if model.ContainsClusterState([]model.ClusterState{model.ClusterStatePending, model.ClusterStateResizing, model.ClusterStateRestarting}, currentState) {
		err := client.Clusters().WaitForClusterRunning(clusterID, 5, 180)
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
		return nil
	}

	if model.ContainsClusterState([]model.ClusterState{model.ClusterStateTerminating}, currentState) {
		err := client.Clusters().WaitForClusterTerminated(clusterID, 5, 180)
		if err != nil {
			return err
		}
		err = client.Clusters().Start(clusterID)
		if err != nil {
			if !strings.Contains(err.Error(), fmt.Sprintf("Cluster %s is in unexpected state Pending.", clusterID)) {
				return err
			}
		}
		err = client.Clusters().WaitForClusterRunning(clusterID, 5, 180)
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
		return nil
	}

	if model.ContainsClusterState([]model.ClusterState{model.ClusterStateTerminated}, currentState) {
		err = client.Clusters().Start(clusterID)
		if err != nil {
			if !strings.Contains(err.Error(), fmt.Sprintf("Cluster %s is in unexpected state Pending.", clusterID)) {
				return err
			}
		}

		err = client.Clusters().WaitForClusterRunning(clusterID, 5, 180)
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
		return nil
	}

	return fmt.Errorf("cluster is in a non recoverable state: %s", currentState)

}
