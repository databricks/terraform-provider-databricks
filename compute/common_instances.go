package compute

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

var (
	oncePool           sync.Once
	commonInstancePool *InstancePoolAndStats
)

// CommonRuntimeVersion presents recommended Spark Version
func CommonRuntimeVersion() string {
	return "6.6.x-scala2.11"
}

// CommonInstanceType presents smallest recommended instance type
func CommonInstanceType() string {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if strings.ToLower(cloudEnv) == "azure" {
		return "Standard_DS3_v2"
	}
	// TODO: create a method on ClustersAPI to give
	// cloud specific delta-cache enabled instance by default.
	return "m4.large"
}

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
			PreloadedSparkVersions: []string{CommonRuntimeVersion()},
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
		SparkVersion:           CommonRuntimeVersion(),
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
		SparkVersion:           CommonRuntimeVersion(),
		InstancePoolID:         CommonInstancePoolID(),
		IdempotencyToken:       "tf-" + randomName,
		AutoterminationMinutes: 20,
	})
	if err != nil {
		panic(err)
	}
	return
}
