package databricks

import (
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/stretchr/testify/assert"
)

func TestAzureAccADLSv1Mount(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := service.CommonEnvironmentClient()
	if !client.UsingAzureAuth() {
		t.Skip("Test is meant only for Azure")
	}
	if !client.AzureAuth.IsClientSecretSet() {
		t.Skip("Test is meant only for client-secret conf Azure")
	}
	gen1Name := os.Getenv("TEST_GEN1_NAME")
	if gen1Name == "" {
		t.Skip("No ADLS account given")
	}
	clusterInfo := service.NewTinyClusterInCommonPoolPossiblyReused()

	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	mp := MountPoint{
		exec:      client.Commands(),
		clusterID: clusterInfo.ClusterID,
		name:      randomName,
	}
	err := mp.Delete()
	assertErrorStartsWith(t, err, "Directory not mounted: /mnt/"+randomName)

	source, err := mp.Source()
	assert.Equal(t, "", source)
	assertErrorStartsWith(t, err, "Mount not found")

	source, err = mp.Mount(AzureADLSGen1Mount{
		ClientID:        client.AzureAuth.ClientID,
		TenantID:        client.AzureAuth.TenantID,
		PrefixType:      "dfs.adls",
		StorageResource: gen1Name,
		Directory:       "/",
		SecretKey:       "key",
		SecretScope:     "y",
	})
	assert.Equal(t, "", source)
	assertErrorStartsWith(t, err, "Secret does not exist with scope: y and key: key")

	randomScope := "test" + randomName
	err = client.SecretScopes().Create(randomScope, "users")
	assert.NoError(t, err)

	err = client.Secrets().Create(client.AzureAuth.ClientSecret, randomScope, "key")
	assert.NoError(t, err)
	defer func() {
		err = client.SecretScopes().Delete(randomScope)
		assert.NoError(t, err)
	}()

	m := AzureADLSGen1Mount{
		ClientID:        client.AzureAuth.ClientID,
		TenantID:        client.AzureAuth.TenantID,
		PrefixType:      "dfs.adls",
		StorageResource: gen1Name,
		Directory:       "/",
		SecretKey:       "key",
		SecretScope:     "test" + randomName,
	}

	source, err = mp.Mount(m)
	assert.Equal(t, m.Source(), source)
	assert.NoError(t, err)
	defer func() {
		err = mp.Delete()
		assert.NoError(t, err)
	}()

	source, err = mp.Source()
	assert.Equal(t, m.Source(), source)
	assert.NoError(t, err)
}
