package service

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var (
	oncePool           sync.Once
	onceClient         sync.Once
	commonInstancePool *model.InstancePoolAndStats
	commonClient       *DatabricksClient
)

// NewClientFromEnvironment makes very good client for testing purposes
func NewClientFromEnvironment() *DatabricksClient {
	client := DatabricksClient{
		Host:       os.Getenv("DATABRICKS_HOST"),
		Token:      os.Getenv("DATABRICKS_TOKEN"),
		Username:   os.Getenv("DATABRICKS_USERNAME"),
		Password:   os.Getenv("DATABRICKS_PASSWORD"),
		ConfigFile: os.Getenv("DATABRICKS_CONFIG_FILE"),
		Profile:    os.Getenv("DATABRICKS_CONFIG_PROFILE"),
		AzureAuth: AzureAuth{
			ResourceID:     os.Getenv("DATABRICKS_AZURE_WORKSPACE_RESOURCE_ID"),
			WorkspaceName:  os.Getenv("DATABRICKS_AZURE_WORKSPACE_NAME"),
			ResourceGroup:  os.Getenv("DATABRICKS_AZURE_RESOURCE_GROUP"),
			SubscriptionID: os.Getenv("ARM_SUBSCRIPTION_ID"),
			ClientID:       os.Getenv("ARM_CLIENT_ID"),
			ClientSecret:   os.Getenv("ARM_CLIENT_SECRET"),
			TenantID:       os.Getenv("ARM_TENANT_ID"),
		},
	}
	err := client.Configure("dev")
	if err != nil {
		panic(err)
	}
	return &client
}

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

// CommonEnvironmentClient configured once per run of application
func CommonEnvironmentClient() *DatabricksClient {
	if commonClient != nil {
		return commonClient
	}
	onceClient.Do(func() {
		commonClient = NewClientFromEnvironment()
	})
	return commonClient
}

// CommonInstancePoolID returns common instance pool that is supposed to be used for internal testing purposes
func CommonInstancePoolID() string {
	if commonInstancePool != nil {
		return commonInstancePool.InstancePoolID
	}
	oncePool.Do(func() { // atomic
		log.Printf("[INFO] Initializing common instance pool")
		client := CommonEnvironmentClient()
		instancePools := client.InstancePools()
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
					commonClient.Host, existingPool.InstancePoolID)
				commonInstancePool = &existingPool
				return
			}
		}
		instancePool := model.InstancePool{
			PreloadedSparkVersions:             []string{CommonRuntimeVersion()},
			NodeTypeID:                         CommonInstanceType(),
			InstancePoolName:                   currentUserPool,
			MaxCapacity:                        10,
			IdleInstanceAutoTerminationMinutes: 15,
		}
		if !client.UsingAzureAuth() {
			instancePool.DiskSpec = &model.InstancePoolDiskSpec{
				DiskType: &model.InstancePoolDiskType{
					EbsVolumeType: model.EbsVolumeTypeGeneralPurposeSsd,
				},
				DiskCount: 1,
				DiskSize:  32,
			}
			instancePool.AwsAttributes = &model.InstancePoolAwsAttributes{
				Availability: model.AwsAvailabilitySpot,
			}
		}
		newPool, err := instancePools.Create(instancePool)
		if err != nil {
			log.Printf("[ERROR] Cannot create instance pool: %v", err)
			panic(err)
		}
		log.Printf("[INFO] Created common instance pool: %s/#setting/clusters/instance-pools/view/%s",
			commonClient.Host, newPool.InstancePoolID)
		commonInstancePool = &newPool
	})
	return commonInstancePool.InstancePoolID
}

// NewTinyClusterInCommonPool creates new cluster for short-lived purposes
func NewTinyClusterInCommonPool() (c model.ClusterInfo, err error) {
	randomName := randomName()
	c, err = CommonEnvironmentClient().Clusters().Create(model.Cluster{
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
func NewTinyClusterInCommonPoolPossiblyReused() (c model.ClusterInfo) {
	randomName := randomName()
	client := CommonEnvironmentClient()
	currentCluster := "TerraformIntegrationTest"
	c, err := client.Clusters().GetOrCreateRunningCluster(currentCluster, model.Cluster{
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

func randomName() string {
	randLen := 12
	b := make([]byte, randLen)
	for i := range b {
		b[i] = charset[rand.Intn(randLen)]
	}
	return string(b)
}
