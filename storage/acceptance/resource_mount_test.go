package acceptance

import (
	"context"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/clusters"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/identity"
	"github.com/databrickslabs/terraform-provider-databricks/internal/compute"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/databrickslabs/terraform-provider-databricks/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAwsAccS3MountGeneric(t *testing.T) {
	client := common.NewClientFromEnvironment()
	instanceProfile := qa.GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	ctx := context.WithValue(context.Background(), common.Current, t.Name())
	instanceProfilesAPI := identity.NewInstanceProfilesAPI(ctx, client)
	instanceProfilesAPI.Synchronized(instanceProfile, func() bool {
		if err := instanceProfilesAPI.Create(identity.InstanceProfileInfo{
			InstanceProfileArn: instanceProfile,
		}); err != nil {
			return false
		}
		bucket := qa.GetEnvOrSkipTest(t, "TEST_S3_BUCKET")
		client := compute.CommonEnvironmentClientWithRealCommandExecutor()
		clustersAPI := clusters.NewClustersAPI(ctx, client)
		clusterInfo, err := storage.GetOrCreateMountingClusterWithInstanceProfile(
			clustersAPI, instanceProfile)
		require.NoError(t, err)
		defer func() {
			err = clustersAPI.PermanentDelete(clusterInfo.ClusterID)
			assert.NoError(t, err)
			err = instanceProfilesAPI.Delete(instanceProfile)
			assert.NoError(t, err)
		}()
		testMounting(t, storage.MountPoint{
			Exec:      client.CommandExecutor(ctx),
			ClusterID: clusterInfo.ClusterID,
			Name:      qa.RandomName("t"),
		}, storage.GenericMount{
			S3: &storage.S3IamMount{
				BucketName: bucket,
			},
		})
		return true
	})
}

func TestAzureAccADLSv1MountGeneric(t *testing.T) {
	client, mp := mountPointThroughReusedCluster(t)
	if !client.IsAzureClientSecretSet() {
		t.Skip("Test is meant only for client-secret conf Azure")
	}
	storageResource := qa.GetEnvOrSkipTest(t, "TEST_DATA_LAKE_STORE_NAME")
	testWithNewSecretScope(t, func(scope, key string) {
		testMounting(t, mp,
			storage.GenericMount{
				Adl: &storage.AzureADLSGen1MountGeneric{
					ClientID:        client.AzureClientID,
					TenantID:        client.AzureTenantID,
					PrefixType:      "dfs.adls",
					StorageResource: storageResource,
					Directory:       "/",
					SecretScope:     scope,
					SecretKey:       key,
				}})
	}, client, mp.Name, client.AzureClientSecret)
}

func TestAzureAccADLSv2MountGeneric(t *testing.T) {
	client, mp := mountPointThroughReusedCluster(t)
	if !client.IsAzureClientSecretSet() {
		t.Skip("Test is meant only for client-secret conf Azure")
	}
	storageAccountName := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_ACCOUNT")
	container := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_ABFSS")
	testWithNewSecretScope(t, func(scope, key string) {
		testMounting(t, mp, storage.GenericMount{
			Abfs: &storage.AzureADLSGen2MountGeneric{
				ClientID:             client.AzureClientID,
				TenantID:             client.AzureTenantID,
				StorageAccountName:   storageAccountName,
				ContainerName:        container,
				SecretScope:          scope,
				SecretKey:            key,
				InitializeFileSystem: true,
				Directory:            "/",
			},
		})
	}, client, mp.Name, client.AzureClientSecret)
}

func TestAccAzureBlobMountGeneric(t *testing.T) {
	client, mp := mountPointThroughReusedCluster(t)
	storageAccountName := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_ACCOUNT")
	accountKey := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_KEY")
	container := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_WASBS")
	testWithNewSecretScope(t, func(scope, key string) {
		testMounting(t, mp, storage.GenericMount{
			Wasb: &storage.AzureBlobMountGeneric{
				StorageAccountName: storageAccountName,
				ContainerName:      container,
				SecretScope:        scope,
				SecretKey:          key,
				Directory:          "/",
			}})
	}, client, mp.Name, accountKey)
}

// TODO: implement it
// func TestGcpAccGcsMount(t *testing.T) {
// 	client, mp := mountPointThroughReusedCluster(t)
// 	storageAccountName := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_ACCOUNT")
// 	accountKey := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_KEY")
// 	container := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_WASBS")
// 	testWithNewSecretScope(t, func(scope, key string) {
// 		testMounting(t, mp, GenericMount{Wasb: &AzureBlobMount{
// 			StorageAccountName: storageAccountName,
// 			ContainerName:      container,
// 			SecretScope:        scope,
// 			SecretKey:          key,
// 			Directory:          "/",
// 		}})
// 	}, client, mp.name, accountKey)
// }
