package compute

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/commands"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/pools"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

var (
	oncePool           sync.Once
	commonInstancePool *pools.InstancePoolAndStats
)

// CommonInstancePoolID returns common instance pool that is supposed to be used for internal testing purposes
func CommonInstancePoolID() string {
	if commonInstancePool != nil {
		return commonInstancePool.InstancePoolID
	}
	configured := os.Getenv("TEST_INSTANCE_POOL_ID")
	if configured != "" {
		return configured
	}
	client := common.CommonEnvironmentClient()
	oncePool.Do(func() { // atomic
		log.Printf("[INFO] Initializing common instance pool")
		ctx := context.Background()
		instancePoolsAPI := pools.NewInstancePoolsAPI(ctx, client)
		clustersAPI := clusters.NewClustersAPI(ctx, client)
		currentUserPool := fmt.Sprintf("Terraform Integration Test by %s", os.Getenv("USER"))
		poolList, err := instancePoolsAPI.List()
		if err != nil {
			log.Printf("[ERROR] Cannot list instance pools: %v", err)
			panic(err)
		}
		for _, existingPool := range poolList.InstancePools {
			if existingPool.InstancePoolName == currentUserPool {
				log.Printf(
					"[INFO] Using existing instance pool: %s/#setting/clusters/instance-pools/view/%s",
					client.Host, existingPool.InstancePoolID)
				commonInstancePool = &existingPool
				return
			}
		}
		instancePool := pools.InstancePool{
			PreloadedSparkVersions: []string{
				clustersAPI.LatestSparkVersionOrDefault(clusters.SparkVersionRequest{Latest: true, LongTermSupport: true})},
			NodeTypeID: clustersAPI.GetSmallestNodeType(clusters.NodeTypeRequest{
				LocalDisk: true,
			}),
			InstancePoolName: currentUserPool,
			MaxCapacity:      10,

			IdleInstanceAutoTerminationMinutes: 15,
		}
		if client.IsAws() {
			instancePool.AwsAttributes = &pools.InstancePoolAwsAttributes{
				Availability: clusters.AwsAvailabilitySpot,
			}
		}
		newPool, err := instancePoolsAPI.Create(instancePool)
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
		return commands.NewCommandsAPI(ctx, client)
	})
	return client
}

// NewTinyClusterInCommonPool creates new cluster for short-lived purposes
func NewTinyClusterInCommonPool() (c clusters.ClusterInfo, err error) {
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	ctx := context.Background()
	clustersAPI := clusters.NewClustersAPI(ctx, CommonEnvironmentClientWithRealCommandExecutor())
	c, err = clustersAPI.Create(clusters.Cluster{
		NumWorkers:  1,
		ClusterName: "Terraform " + randomName,
		SparkVersion: clustersAPI.LatestSparkVersionOrDefault(clusters.SparkVersionRequest{
			Latest:          true,
		}),
		InstancePoolID:         CommonInstancePoolID(),
		IdempotencyToken:       "tf-" + randomName,
		AutoterminationMinutes: 20,
	})
	return
}

// NewTinyClusterInCommonPoolPossiblyReused is recommended to be used for testing only
func NewTinyClusterInCommonPoolPossiblyReused() (c clusters.ClusterInfo) {
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	currentCluster := "TerraformIntegrationTest"
	ctx := context.Background()
	clustersAPI := clusters.NewClustersAPI(ctx, CommonEnvironmentClientWithRealCommandExecutor())
	c, err := clustersAPI.GetOrCreateRunningCluster(currentCluster, clusters.Cluster{
		NumWorkers:  1,
		ClusterName: currentCluster,
		SparkVersion: clustersAPI.LatestSparkVersionOrDefault(clusters.SparkVersionRequest{
			Latest:          true,
		}),
		InstancePoolID:         CommonInstancePoolID(),
		IdempotencyToken:       "tf-" + randomName,
		AutoterminationMinutes: 20,
	})
	if err != nil {
		panic(err)
	}
	return
}
