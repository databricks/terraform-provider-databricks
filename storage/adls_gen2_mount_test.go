package storage

import (
	"os"
	"testing"

	. "github.com/databrickslabs/databricks-terraform/access"
	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/stretchr/testify/assert"
)

func TestAzureAccADLSv2Mount(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := common.CommonEnvironmentClient()
	if !client.UsingAzureAuth() {
		t.Skip("Test is meant only for Azure")
	}
	if !client.AzureAuth.IsClientSecretSet() {
		t.Skip("Test is meant only for client-secret conf Azure")
	}
	gen2AdalName := os.Getenv("TEST_GEN2_ADAL_NAME")
	if gen2AdalName == "" {
		t.Skip("No ADLS account given")
	}
	clusterInfo := compute.NewTinyClusterInCommonPoolPossiblyReused()

	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	mp := MountPoint{
		exec:      client.CommandExecutor(),
		clusterID: clusterInfo.ClusterID,
		name:      randomName,
	}
	err := mp.Delete()
	qa.AssertErrorStartsWith(t, err, "Directory not mounted: /mnt/"+randomName)

	source, err := mp.Source()
	assert.Equal(t, "", source)
	qa.AssertErrorStartsWith(t, err, "Mount not found")

	source, err = mp.Mount(AzureADLSGen2Mount{
		ClientID:             client.AzureAuth.ClientID,
		TenantID:             client.AzureAuth.TenantID,
		ContainerName:        "dev",
		Directory:            "/",
		InitializeFileSystem: true,
		SecretKey:            "e",
		SecretScope:          "f",
		StorageAccountName:   gen2AdalName,
	})
	assert.Equal(t, "", source)
	qa.AssertErrorStartsWith(t, err, "Secret does not exist with scope: f and key: e")

	randomScope := "test" + randomName
	err = NewSecretScopesAPI(client).Create(randomScope, "users")
	assert.NoError(t, err)
	defer func() {
		err = NewSecretScopesAPI(client).Delete(randomScope)
		assert.NoError(t, err)
	}()

	err = NewSecretsAPI(client).Create(client.AzureAuth.ClientSecret, randomScope, "key")
	assert.NoError(t, err)

	m := AzureADLSGen2Mount{
		ClientID:             client.AzureAuth.ClientID,
		TenantID:             client.AzureAuth.TenantID,
		ContainerName:        "dev",
		Directory:            "/",
		InitializeFileSystem: true,
		SecretKey:            "key",
		SecretScope:          randomScope,
		StorageAccountName:   gen2AdalName,
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
