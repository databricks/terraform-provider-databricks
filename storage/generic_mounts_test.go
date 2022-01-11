package storage

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNames(t *testing.T) {
	mount_name := "abc"
	gm := GenericMount{MountName: mount_name}
	assert.Equal(t, gm.Name(), mount_name)
	assert.Equal(t, GenericMount{}.Name(), "")
	gm = GenericMount{Abfs: &AzureADLSGen2MountGeneric{ContainerName: mount_name}}
	assert.Equal(t, gm.Name(), mount_name)
	gm = GenericMount{Wasb: &AzureBlobMountGeneric{ContainerName: mount_name}}
	assert.Equal(t, gm.Name(), mount_name)
	gm = GenericMount{Adl: &AzureADLSGen1MountGeneric{StorageResource: mount_name}}
	assert.Equal(t, gm.Name(), mount_name)
}

func TestGenericMountDefaults(t *testing.T) {
	client := &common.DatabricksClient{}

	gm := GenericMount{URI: "foo://bar", Options: map[string]string{"foo": "bar"}}
	assert.Equal(t, "foo://bar", gm.Source())
	assert.Len(t, gm.Config(client), 1)
	assert.Equal(t, "bar", gm.Config(client)["foo"])

	gm = GenericMount{MountName: "test"}
	d := ResourceMount().TestResourceData()
	err := gm.ValidateAndApplyDefaults(d, client)
	assert.EqualError(t, err, "value of name is not specified or empty")

	d.Set("name", "X")
	err = gm.ValidateAndApplyDefaults(d, client)
	assert.EqualError(t, err, "value of uri is not specified or empty")

	d.Set("uri", "s3://abc/")
	err = gm.ValidateAndApplyDefaults(d, client)
	assert.NoError(t, err)

	gm = GenericMount{Abfs: &AzureADLSGen2MountGeneric{}}
	err = gm.ValidateAndApplyDefaults(d, client)
	assert.EqualError(t, err, "container_name or storage_account_name are empty, and resource_id or uri aren't specified")
}

const containerResourceID = "/subscriptions/5363c143-2af7-4fb5-8a9d-ab1b2c8e756e/resourceGroups/test-rg/providers/Microsoft.Storage/storageAccounts/lrs-acc/blobServices/default/containers/test"
const adlResourceID = "/subscriptions/6369c148-f8a9-4fb5-8a9d-ac1b2c8e756e/resourceGroups/alexott-rg/providers/Microsoft.DataLakeStore/accounts/aottgen1"

func TestARMParsing(t *testing.T) {
	acc, container, err := parseStorageContainerId(containerResourceID)
	require.NoError(t, err, err)
	assert.Equal(t, acc, "lrs-acc")
	assert.Equal(t, container, "test")
}

func TestARMParsingError(t *testing.T) {
	_, _, err := parseStorageContainerId("abc")
	qa.AssertErrorStartsWith(t, err, "parsing failed for ")
}

func TestARMParsing2(t *testing.T) {
	res, err := azure.ParseResourceID(adlResourceID)
	require.NoError(t, err, err)
	assert.Equal(t, res.ResourceName, "aottgen1")
}

func TestGetContainerDefaults(t *testing.T) {
	d := ResourceMount().TestResourceData()
	d.Set("resource_id", "..")
	_, _, err := getContainerDefaults(d, []string{}, "")
	assert.EqualError(t, err, "parsing failed for ... Invalid container resource Id format")
}

func TestAzureADLSGen2MountGeneric_Resource_NoTenant(t *testing.T) {
	d := ResourceMount().TestResourceData()
	d.Set("resource_id", containerResourceID)
	x := &AzureADLSGen2MountGeneric{}
	err := x.ValidateAndApplyDefaults(d, &common.DatabricksClient{})
	assert.EqualError(t, err, "tenant_id is not defined, and we can't extract it: can't get Azure JWT token in non-Azure environment")
}

func newTestJwt(t *testing.T, claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	result, err := token.SignedString([]byte("_"))
	assert.NoError(t, err)
	return result
}

func setupJwtTestClient() (*httptest.Server, *common.DatabricksClient) {
	client := &common.DatabricksClient{InsecureSkipVerify: true,
		Host: "https://adb-1232.azuredatabricks.net",
	}
	server := httptest.NewUnstartedServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
		}))
	server.StartTLS()
	// resource management endpoints end with a trailing slash in url
	client.AzureEnvironment = &azure.Environment{
		ResourceManagerEndpoint: fmt.Sprintf("%s/", server.URL),
		ActiveDirectoryEndpoint: "http://mock-aad/",
	}
	ctx := context.Background()
	err := client.Authenticate(ctx)
	if err != nil {
		log.Printf("[ERROR] Can't auth: %s", err)
		return nil, nil
	}
	return server, client
}

func TestGetTenantIDCannotBeDetected(t *testing.T) {
	defer common.CleanupEnvironment()()
	os.Setenv("PATH", "../common/testdata:/bin")
	os.Setenv("TF_AAD_TOKEN", newTestJwt(t, jwt.MapClaims{
		"tid": "",
	}))

	srv, client := setupJwtTestClient()
	require.NotNil(t, srv)
	defer srv.Close()

	_, err := getTenantID(client)
	assert.EqualError(t, err, "tenant_id isn't provided & we can't detect it")
}

func TestAzureADLSGen2MountGeneric_NameAndTenant(t *testing.T) {
	defer common.CleanupEnvironment()()
	os.Setenv("PATH", "../common/testdata:/bin")
	os.Setenv("TF_AAD_TOKEN", newTestJwt(t, jwt.MapClaims{
		"tid": "some-tenant",
	}))

	srv, client := setupJwtTestClient()
	require.NotNil(t, srv)
	defer srv.Close()

	d := ResourceMount().TestResourceData()
	d.Set("resource_id", containerResourceID)
	x := &AzureADLSGen2MountGeneric{
		SecretScope: "a",
		SecretKey:   "b",
		ClientID:    "c",
	}

	err := x.ValidateAndApplyDefaults(d, client)
	assert.NoError(t, err)
	assert.Equal(t, "abfss://test@lrs-acc.dfs.core.windows.net", x.Source())

	config := x.Config(client)
	assert.Len(t, config, 6)
	assert.Equal(t, "OAuth", config["fs.azure.account.auth.type"])
	assert.Equal(t, "c", config["fs.azure.account.oauth2.client.id"])
	assert.Equal(t, "{{secrets/a/b}}", config["fs.azure.account.oauth2.client.secret"])
	assert.Equal(t, "http://mock-aad/some-tenant/oauth2/token", config["fs.azure.account.oauth2.client.endpoint"])
	assert.Equal(t, "false", config["fs.azure.createRemoteFileSystemDuringInitialization"])
}

func TestAzureADLSGen1MountGeneric(t *testing.T) {
	defer common.CleanupEnvironment()()
	os.Setenv("PATH", "../common/testdata:/bin")
	os.Setenv("TF_AAD_TOKEN", newTestJwt(t, jwt.MapClaims{
		"tid": "some-tenant",
	}))

	srv, client := setupJwtTestClient()
	require.NotNil(t, srv)
	defer srv.Close()

	d := ResourceMount().TestResourceData()
	x := &AzureADLSGen1MountGeneric{
		PrefixType:  "X",
		SecretScope: "a",
		SecretKey:   "b",
		ClientID:    "c",
	}

	d.Set("resource_id", "..")
	err := x.ValidateAndApplyDefaults(d, client)
	assert.EqualError(t, err, "parsing failed for ... Invalid resource Id format")

	d.Set("resource_id", containerResourceID)
	err = x.ValidateAndApplyDefaults(d, client)
	assert.EqualError(t, err, "incorrect resource type or provider in resource_id: "+containerResourceID)

	d.Set("resource_id", adlResourceID)
	err = x.ValidateAndApplyDefaults(d, &common.DatabricksClient{})
	assert.EqualError(t, err, "tenant_id is not defined, and we can't extract it: can't get Azure JWT token in non-Azure environment")

	d.Set("resource_id", adlResourceID)
	err = x.ValidateAndApplyDefaults(d, client)
	assert.NoError(t, err)
	assert.Equal(t, "adl://aottgen1.azuredatalakestore.net", x.Source())

	config := x.Config(client)
	assert.Len(t, config, 4)
	assert.Equal(t, "ClientCredential", config["X.oauth2.access.token.provider.type"])
	assert.Equal(t, "c", config["X.oauth2.client.id"])
	assert.Equal(t, "{{secrets/a/b}}", config["X.oauth2.credential"])
	assert.Equal(t, "http://mock-aad/some-tenant/oauth2/token", config["X.oauth2.refresh.url"])
}

func TestAzureBlobMount(t *testing.T) {
	d := ResourceMount().TestResourceData()
	d.Set("resource_id", containerResourceID)
	x := &AzureBlobMountGeneric{
		SecretScope: "a",
		SecretKey:   "b",
	}

	client := &common.DatabricksClient{}
	err := x.ValidateAndApplyDefaults(d, client)
	assert.NoError(t, err)
	assert.Equal(t, "wasbs://test@lrs-acc.blob.core.windows.net", x.Source())

	config := x.Config(client)
	assert.Len(t, config, 1)
	assert.Equal(t, "{{secrets/a/b}}", config["fs.azure.account.key.lrs-acc.blob.core.windows.net"])

	x.AuthType = "SAS"
	config = x.Config(client)
	assert.Len(t, config, 1)
	assert.Equal(t, "{{secrets/a/b}}", config["fs.azure.sas.test.lrs-acc.blob.core.windows.net"])
}
