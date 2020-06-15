package databricks

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/arn"
	"strings"
	"time"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
)

func changeClusterIntoRunningState(clusterID string, client *service.DBApiClient) error {
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

func isClusterMissing(errorMsg, resourceID string) bool {
	return strings.Contains(errorMsg, "INVALID_PARAMETER_VALUE") &&
		strings.Contains(errorMsg, fmt.Sprintf("Cluster %s does not exist", resourceID))
}

// PackagedMWSIds is a struct that contains both the MWS acct id and the ResourceId (resources are networks, creds, etc.)
type PackagedMWSIds struct {
	MwsAcctId  string
	ResourceId string
}

// Helps package up MWSAccountId with another id such as credentials id or network id
// uses format mwsAcctId/otherId
func packMWSAccountId(idsToPackage PackagedMWSIds) string {
	return fmt.Sprintf("%s/%s", idsToPackage.MwsAcctId, idsToPackage.ResourceId)
}

// Helps unpackage MWSAccountId from another id such as credentials id or network id
func unpackMWSAccountId(combined string) (PackagedMWSIds, error) {
	var packagedMWSIds PackagedMWSIds
	parts := strings.Split(combined, "/")
	if len(parts) != 2 {
		return packagedMWSIds, fmt.Errorf("unpacked account has more than or less than two parts, combined id: %s", combined)
	}
	packagedMWSIds.MwsAcctId = parts[0]
	packagedMWSIds.ResourceId = parts[1]
	return packagedMWSIds, nil
}

// ValidateAWSARN is a ValidateFunc that ensures the role id is a valid aws iam instance profile arn
func ValidateInstanceProfileARN(val interface{}, key string) (warns []string, errs []error) {
	v := val.(string)

	if v == "" {
		return nil, []error{fmt.Errorf("%s is empty got: %s, must be an aws instance profile arn", key, v)}
	}

	// Parse and verify instance profiles
	instanceProfileArn, err := arn.Parse(v)
	if err != nil {
		return nil, []error{fmt.Errorf("%s is invalid got: %s recieved error: %w", key, v, err)}
	}
	// Verify instance profile resource type, Resource gets parsed as instance-profile/<profile-name>
	if !strings.HasPrefix(instanceProfileArn.Resource, "instance-profile") {
		return nil, []error{fmt.Errorf("%s must be an instance profile resource, got: %s in %s",
			key, instanceProfileArn.Resource, v)}
	}
	return nil, nil
}
