package acceptance

import (
	"context"
	"os"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/access"
	"github.com/databrickslabs/terraform-provider-databricks/commands"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
	"github.com/databrickslabs/terraform-provider-databricks/internal/compute"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/databrickslabs/terraform-provider-databricks/storage"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func mountResourceCheck(name string,
	cb func(*common.DatabricksClient, storage.MountPoint) error) resource.TestCheckFunc {
	return acceptance.ResourceCheck(name,
		func(ctx context.Context, client *common.DatabricksClient, id string) error {
			client.WithCommandExecutor(func(ctx context.Context, client *common.DatabricksClient) common.CommandExecutor {
				return commands.NewCommandsAPI(ctx, client)
			})
			clusterInfo := compute.NewTinyClusterInCommonPoolPossiblyReused()
			mp := storage.NewMountPoint(client.CommandExecutor(context.Background()), id, clusterInfo.ClusterID)
			return cb(client, mp)
		})
}

func testMounting(t *testing.T, mp storage.MountPoint, m storage.Mount) {
	client := common.CommonEnvironmentClient()
	source, err := mp.Mount(m, client)
	assert.Equal(t, m.Source(), source)
	assert.NoError(t, err)
	defer func() {
		err = mp.Delete()
		assert.NoError(t, err)
	}()
	source, err = mp.Source()
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
	secretScopes := access.NewSecretScopesAPI(ctx, client)
	err := secretScopes.Create(access.SecretScope{
		Name:                   randomScope,
		InitialManagePrincipal: "users",
	})
	require.NoError(t, err)
	defer func() {
		err = secretScopes.Delete(randomScope)
		assert.NoError(t, err)
	}()

	secrets := access.NewSecretsAPI(ctx, client)
	err = secrets.Create(secret, randomScope, randomKey)
	require.NoError(t, err)

	callback(randomScope, randomKey)
}

func TestAccSourceOnInvalidMountFails(t *testing.T) {
	_, mp := mountPointThroughReusedCluster(t)
	source, err := mp.Source()
	assert.Equal(t, "", source)
	qa.AssertErrorStartsWith(t, err, "Mount not found")
}

func TestAccInvalidSecretScopeFails(t *testing.T) {
	_, mp := mountPointThroughReusedCluster(t)
	client := common.CommonEnvironmentClient()
	source, err := mp.Mount(storage.AzureADLSGen1Mount{
		ClientID:        "abc",
		TenantID:        "bcd",
		PrefixType:      "dfs.adls",
		StorageResource: "def",
		Directory:       "/",
		SecretKey:       "key",
		SecretScope:     "y",
	}, client)
	assert.Equal(t, "", source)
	qa.AssertErrorStartsWith(t, err, "Secret does not exist with scope: y and key: key")
}
