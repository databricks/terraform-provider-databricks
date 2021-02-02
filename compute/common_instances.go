package compute

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

var (
	oncePool           sync.Once
	commonInstancePool *InstancePoolAndStats
)

// CommonInstancePoolID returns common instance pool that is supposed to be used for internal testing purposes
func CommonInstancePoolID() string {
	if commonInstancePool != nil {
		return commonInstancePool.InstancePoolID
	}
	client := common.CommonEnvironmentClient()
	oncePool.Do(func() { // atomic
		log.Printf("[INFO] Initializing common instance pool")
		ctx := context.Background()
		instancePools := NewInstancePoolsAPI(ctx, client)
		clusters := NewClustersAPI(ctx, client)
		currentUserPool := fmt.Sprintf("Terraform Integration Test by %s", os.Getenv("USER"))
		pools, err := instancePools.List()
		if err != nil {
			log.Printf("[ERROR] Cannot list instance pools: %v", err)
			panic(err)
		}
		for _, existingPool := range pools.InstancePools {
			if existingPool.InstancePoolName == currentUserPool {
				log.Printf(
					"[INFO] Using existing instance pool: %s/#setting/clusters/instance-pools/view/%s",
					client.Host, existingPool.InstancePoolID)
				commonInstancePool = &existingPool
				return
			}
		}
		instancePool := InstancePool{
			PreloadedSparkVersions: []string{
				clusters.LatestSparkVersionOrDefault(SparkVersionRequest{Latest: true, LongTermSupport: true})},
			NodeTypeID: clusters.GetSmallestNodeType(NodeTypeRequest{
				LocalDisk: true,
			}),
			InstancePoolName: currentUserPool,
			MaxCapacity:      10,

			IdleInstanceAutoTerminationMinutes: 15,
		}
		if !client.IsAzure() {
			instancePool.AwsAttributes = &InstancePoolAwsAttributes{
				Availability: AwsAvailabilitySpot,
			}
		}
		newPool, err := instancePools.Create(instancePool)
		if err != nil {
			log.Printf("[ERROR] Cannot create instance pool: %v", err)
			panic(err)
		}
		log.Printf("[INFO] Created common instance pool: %s/#setting/clusters/instance-pools/view/%s",
			client.Host, newPool.InstancePoolID)
		commonInstancePool = &newPool
	})
	return commonInstancePool.InstancePoolID
}

// CommonEnvironmentClientWithRealCommandExecutor is good for internal tests
func CommonEnvironmentClientWithRealCommandExecutor() *common.DatabricksClient {
	client := common.CommonEnvironmentClient()
	client.WithCommandExecutor(func(ctx context.Context, _ *common.DatabricksClient) common.CommandExecutor {
		return NewCommandsAPI(ctx, client)
	})
	return client
}

// NewTinyClusterInCommonPool creates new cluster for short-lived purposes
func NewTinyClusterInCommonPool() (c ClusterInfo, err error) {
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	ctx := context.Background()
	clusters := NewClustersAPI(ctx, CommonEnvironmentClientWithRealCommandExecutor())
	c, err = clusters.Create(Cluster{
		NumWorkers:             1,
		ClusterName:            "Terraform " + randomName,
		SparkVersion:           clusters.LatestSparkVersionOrDefault(SparkVersionRequest{Latest: true, LongTermSupport: true}),
		InstancePoolID:         CommonInstancePoolID(),
		IdempotencyToken:       "tf-" + randomName,
		AutoterminationMinutes: 20,
	})
	return
}

// NewTinyClusterInCommonPoolPossiblyReused is recommended to be used for testing only
func NewTinyClusterInCommonPoolPossiblyReused() (c ClusterInfo) {
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	currentCluster := "TerraformIntegrationTest"
	ctx := context.Background()
	clusters := NewClustersAPI(ctx, CommonEnvironmentClientWithRealCommandExecutor())
	c, err := clusters.GetOrCreateRunningCluster(currentCluster, Cluster{
		NumWorkers:             1,
		ClusterName:            currentCluster,
		SparkVersion:           clusters.LatestSparkVersionOrDefault(SparkVersionRequest{Latest: true, LongTermSupport: true}),
		InstancePoolID:         CommonInstancePoolID(),
		IdempotencyToken:       "tf-" + randomName,
		AutoterminationMinutes: 20,
	})
	if err != nil {
		panic(err)
	}
	return
}
