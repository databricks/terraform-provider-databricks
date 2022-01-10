package acceptance

import (
	"context"
	"os"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/internal/compute"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/databrickslabs/terraform-provider-databricks/secrets"
	"github.com/databrickslabs/terraform-provider-databricks/storage"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testMounting(t *testing.T, mp storage.MountPoint, m storage.Mount) {
	client := common.CommonEnvironmentClient()
	info, err := mp.Mount(m, client)
	assert.Equal(t, m.Source(), info.Source)
	assert.NoError(t, err)
	defer func() {
		err = mp.Delete()
		assert.NoError(t, err)
	}()
	source, err := mp.Source(m, client)
	require.Equalf(t, m.Source(), source, "Error: %v", err)
}

func mountPointThroughReusedCluster(t *testing.T) (*common.DatabricksClient, storage.MountPoint) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	ctx := context.Background()
	client := common.CommonEnvironmentClient()
	clusterInfo := compute.NewTinyClusterInCommonPoolPossiblyReused()
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	return client, storage.MountPoint{
		Exec:      client.CommandExecutor(ctx),
		ClusterID: clusterInfo.ClusterID,
		Name:      randomName,
	}
}

func testWithNewSecretScope(t *testing.T, callback func(string, string),
	client *common.DatabricksClient, suffix, secret string) {
	randomScope := "test" + suffix
	randomKey := "key" + suffix

	ctx := context.Background()
	secretScopes := secrets.NewSecretScopesAPI(ctx, client)
	err := secretScopes.Create(secrets.SecretScope{
		Name:                   randomScope,
		InitialManagePrincipal: "users",
	})
	require.NoError(t, err)
	defer func() {
		err = secretScopes.Delete(randomScope)
		assert.NoError(t, err)
	}()

	secrets := secrets.NewSecretsAPI(ctx, client)
	err = secrets.Create(secret, randomScope, randomKey)
	require.NoError(t, err)

	callback(randomScope, randomKey)
}

func TestAccSourceOnInvalidMountFails(t *testing.T) {
	client, mp := mountPointThroughReusedCluster(t)
	source, err := mp.Source(&storage.AzureADLSGen2MountGeneric{
		ContainerName:      "a",
		StorageAccountName: "b",
	}, client)
	assert.Equal(t, "", source)
	qa.AssertErrorStartsWith(t, err, "Mount not found")
}

func TestAccInvalidSecretScopeFails(t *testing.T) {
	_, mp := mountPointThroughReusedCluster(t)
	client := common.CommonEnvironmentClient()
	info, err := mp.Mount(storage.AzureADLSGen1Mount{
		ClientID:        "abc",
		TenantID:        "bcd",
		PrefixType:      "dfs.adls",
		StorageResource: "def",
		Directory:       "/",
		SecretKey:       "key",
		SecretScope:     "y",
	}, client)
	assert.Equal(t, "", info.Source)
	qa.AssertErrorStartsWith(t, err, "Secret does not exist with scope: y and key: key")
}
