package service

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
)

type Mount interface {
	Create(client DBApiClient, clusterId string) error
	Delete(client DBApiClient, clusterId string) error
	Read(client DBApiClient, clusterId string) (string, error)
}

type AWSIamMount struct {
	Mount
	S3BucketName string
	MountName    string
}

func NewAWSIamMount(s3BucketName string, mountName string) *AWSIamMount {
	return &AWSIamMount{S3BucketName: s3BucketName, MountName: mountName}
}

func (m AWSIamMount) Create(client DBApiClient, clusterId string) error {
	iamMountCommand := fmt.Sprintf(`
dbutils.fs.mount("s3a://%s", "/mnt/%s")
dbutils.fs.ls("/mnt/%s")
dbutils.notebook.exit("success")
`, m.S3BucketName, m.MountName, m.MountName)
	resp, err := client.Commands().Execute(clusterId, "python", iamMountCommand)
	if err != nil {
		return err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return errors.New(resp.Results.Summary)
	}
	return nil
}

func (m AWSIamMount) Delete(client DBApiClient, clusterId string) error {
	iamMountCommand := fmt.Sprintf(`
dbutils.fs.unmount("/mnt/%s")
dbutils.fs.refreshMounts()
dbutils.notebook.exit("success")
`, m.MountName)
	resp, err := client.Commands().Execute(clusterId, "python", iamMountCommand)
	if err != nil {
		return err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return errors.New(resp.Results.Summary)
	}
	return nil
}

func (m AWSIamMount) Read(client DBApiClient, clusterId string) (string, error) {
	iamMountCommand := fmt.Sprintf(`
dbutils.fs.ls("/mnt/%s")
for mount in dbutils.fs.mounts():
  if mount.mountPoint == "/mnt/%s":
    dbutils.notebook.exit(mount.source)
`, m.MountName, m.MountName)
	resp, err := client.Commands().Execute(clusterId, "python", iamMountCommand)
	if err != nil {
		return "", err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return "", errors.New(resp.Results.Summary)
	}
	if resp.Results.ResultType == "text" && resp.Results.Data.(string) == "" {
		return "", errors.New("Unable to find mount point!")
	}
	respBucket := strings.Replace(resp.Results.Data.(string), "s3a://", "", 1)
	if resp.Results.ResultType == "text" && respBucket != m.S3BucketName {
		return "", errors.New(fmt.Sprintf("Does not match bucket value! %s != %s!", resp.Results.Data.(string), m.S3BucketName))
	}
	return respBucket, nil
}

type AzureBlobMount struct {
	Mount
	ContainerName      string
	StorageAccountName string
	Directory          string
	MountName          string
	AuthType           string
	SecretScope        string
	SecretKey          string
}

func NewAzureBlobMount(containerName string, storageAccountName string, directory string, mountName string, authType string, secretScope string, secretKey string) *AzureBlobMount {
	return &AzureBlobMount{ContainerName: containerName, StorageAccountName: storageAccountName, Directory: directory, MountName: mountName, AuthType: authType, SecretScope: secretScope, SecretKey: secretKey}
}

func (m AzureBlobMount) Create(client DBApiClient, clusterId string) error {
	var confKey string

	if m.AuthType == "SAS" {
		confKey = fmt.Sprintf("fs.azure.sas.%s.%s.blob.core.windows.net", m.ContainerName, m.StorageAccountName)
	} else {
		confKey = fmt.Sprintf("fs.azure.account.key.%s.blob.core.windows.net", m.StorageAccountName)
	}
	iamMountCommand := fmt.Sprintf(`
try:
  dbutils.fs.mount(
    source = "wasbs://%s@%s.blob.core.windows.net%s",
    mount_point = "/mnt/%s",
    extra_configs = {"%s":dbutils.secrets.get(scope = "%s", key = "%s")})
except Exception as e:
  dbutils.fs.unmount("/mnt/%s")
  raise e
dbutils.notebook.exit("success")
`, m.ContainerName, m.StorageAccountName, m.Directory, m.MountName, confKey, m.SecretScope, m.SecretKey, m.MountName)
	resp, err := client.Commands().Execute(clusterId, "python", iamMountCommand)
	if err != nil {
		return err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return errors.New(resp.Results.Summary)
	}
	return nil
}

func (m AzureBlobMount) Delete(client DBApiClient, clusterId string) error {
	iamMountCommand := fmt.Sprintf(`
dbutils.fs.unmount("/mnt/%s")
dbutils.fs.refreshMounts()
dbutils.notebook.exit("success")
`, m.MountName)
	resp, err := client.Commands().Execute(clusterId, "python", iamMountCommand)
	if err != nil {
		return err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return errors.New(resp.Results.Summary)
	}
	return nil
}

func (m AzureBlobMount) Read(client DBApiClient, clusterId string) (string, error) {
	iamMountCommand := fmt.Sprintf(`
dbutils.fs.ls("/mnt/%s")
for mount in dbutils.fs.mounts():
 if mount.mountPoint == "/mnt/%s":
   dbutils.notebook.exit(mount.source)
`, m.MountName, m.MountName)
	resp, err := client.Commands().Execute(clusterId, "python", iamMountCommand)
	if err != nil {
		return "", err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return "", errors.New(resp.Results.Summary)
	}
	if resp.Results.ResultType == "text" && resp.Results.Data.(string) == "" {
		return "", errors.New("Unable to find mount point!")
	}

	container, storageAccount, directory, err := ProcessAzureWasbAbfssUris(resp.Results.Data.(string))
	if err != nil {
		return "", err
	}
	if resp.Results.ResultType == "text" && container != m.ContainerName &&
		storageAccount != m.StorageAccountName &&
		m.Directory != directory {
		return "", errors.New(fmt.Sprintf("Does not match uri with storage account and container values!"+
			" %s@%s != %s!", m.ContainerName, m.StorageAccountName, resp.Results.Data.(string)))
	}
	return resp.Results.Data.(string), nil
}

type AzureADLSGen1Mount struct {
	Mount
	StorageResource string
	Directory       string
	MountName       string
	PrefixType      string
	ClientId        string
	TenantId        string
	SecretScope     string
	SecretKey       string
}

func NewAzureADLSGen1Mount(storageResource string, directory string, mountName string, prefixType string, clientId string, tenantId string, secretScope string, secretKey string) *AzureADLSGen1Mount {
	return &AzureADLSGen1Mount{StorageResource: storageResource, Directory: directory, MountName: mountName, PrefixType: prefixType, ClientId: clientId, TenantId: tenantId, SecretScope: secretScope, SecretKey: secretKey}
}

func (m AzureADLSGen1Mount) Create(client DBApiClient, clusterId string) error {
	iamMountCommand := fmt.Sprintf(`
try:
  configs = {"%s.oauth2.access.token.provider.type": "ClientCredential",
          "%s.oauth2.client.id": "%s",
          "%s.oauth2.credential": dbutils.secrets.get(scope = "%s", key = "%s"),
          "%s.oauth2.refresh.url": "https://login.microsoftonline.com/%s/oauth2/token"}
  dbutils.fs.mount(
   source = "adl://%s.azuredatalakestore.net%s",
   mount_point = "/mnt/%s",
   extra_configs = configs)
except Exception as e:
  dbutils.fs.unmount("/mnt/%s")
  raise e
dbutils.notebook.exit("success")
`, m.PrefixType, m.PrefixType, m.ClientId, m.PrefixType, m.SecretScope, m.SecretKey, m.PrefixType, m.TenantId,
		m.StorageResource, m.Directory, m.MountName, m.MountName)
	resp, err := client.Commands().Execute(clusterId, "python", iamMountCommand)
	if err != nil {
		return err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return errors.New(resp.Results.Summary)
	}
	return nil
}

func (m AzureADLSGen1Mount) Delete(client DBApiClient, clusterId string) error {
	iamMountCommand := fmt.Sprintf(`
dbutils.fs.unmount("/mnt/%s")
dbutils.fs.refreshMounts()
dbutils.notebook.exit("success")
`, m.MountName)
	resp, err := client.Commands().Execute(clusterId, "python", iamMountCommand)
	if err != nil {
		return err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return errors.New(resp.Results.Summary)
	}
	return nil
}

func (m AzureADLSGen1Mount) Read(client DBApiClient, clusterId string) (string, error) {
	iamMountCommand := fmt.Sprintf(`
dbutils.fs.ls("/mnt/%s")
for mount in dbutils.fs.mounts():
 if mount.mountPoint == "/mnt/%s":
   dbutils.notebook.exit(mount.source)
`, m.MountName, m.MountName)
	resp, err := client.Commands().Execute(clusterId, "python", iamMountCommand)
	if err != nil {
		return "", err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return "", errors.New(resp.Results.Summary)
	}
	if resp.Results.ResultType == "text" && resp.Results.Data.(string) == "" {
		return "", errors.New("Unable to find mount point!")
	}

	storageResource, directory, err := ProcessAzureAdlsGen1Uri(resp.Results.Data.(string))
	if err != nil {
		return "", err
	}
	if resp.Results.ResultType == "text" && storageResource != m.StorageResource &&
		m.Directory != directory {
		return "", errors.New(fmt.Sprintf("Does not match uri with storage account and container values!"+
			" %s@%s != %s!", m.StorageResource, m.Directory, resp.Results.Data.(string)))
	}
	return resp.Results.Data.(string), nil
}

type AzureADLSGen2Mount struct {
	Mount
	ContainerName      string
	StorageAccountName string
	Directory          string
	MountName          string
	ClientId           string
	TenantId           string
	SecretScope        string
	SecretKey          string
}

func NewAzureADLSGen2Mount(containerName string, storageAccountName string, directory string, mountName string, clientId string, tenantId string, secretScope string, secretKey string) *AzureADLSGen2Mount {
	return &AzureADLSGen2Mount{ContainerName: containerName, StorageAccountName: storageAccountName, Directory: directory, MountName: mountName, ClientId: clientId, TenantId: tenantId, SecretScope: secretScope, SecretKey: secretKey}
}

func (m AzureADLSGen2Mount) Create(client DBApiClient, clusterId string) error {
	iamMountCommand := fmt.Sprintf(`
try:
  configs = {"fs.azure.account.auth.type": "OAuth",
           "fs.azure.account.oauth.provider.type": "org.apache.hadoop.fs.azurebfs.oauth2.ClientCredsTokenProvider",
           "fs.azure.account.oauth2.client.id": "%s",
           "fs.azure.account.oauth2.client.secret": dbutils.secrets.get(scope = "%s", key = "%s"),
           "fs.azure.account.oauth2.client.endpoint": "https://login.microsoftonline.com/%s/oauth2/token"}
  dbutils.fs.mount(
   source = "abfss://%s@%s.dfs.core.windows.net/%s",
   mount_point = "/mnt/%s",
   extra_configs = configs)
except Exception as e:
  dbutils.fs.unmount("/mnt/%s")
  raise e
dbutils.notebook.exit("success")
`, m.ClientId, m.SecretScope, m.SecretKey, m.TenantId, m.ContainerName, m.StorageAccountName, m.Directory, m.MountName, m.MountName)
	resp, err := client.Commands().Execute(clusterId, "python", iamMountCommand)
	if err != nil {
		return err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return errors.New(resp.Results.Summary)
	}
	return nil
}

func (m AzureADLSGen2Mount) Delete(client DBApiClient, clusterId string) error {
	iamMountCommand := fmt.Sprintf(`
dbutils.fs.unmount("/mnt/%s")
dbutils.fs.refreshMounts()
dbutils.notebook.exit("success")
`, m.MountName)
	resp, err := client.Commands().Execute(clusterId, "python", iamMountCommand)
	if err != nil {
		return err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return errors.New(resp.Results.Summary)
	}
	return nil
}

func (m AzureADLSGen2Mount) Read(client DBApiClient, clusterId string) (string, error) {
	iamMountCommand := fmt.Sprintf(`
dbutils.fs.ls("/mnt/%s")
for mount in dbutils.fs.mounts():
 if mount.mountPoint == "/mnt/%s":
   dbutils.notebook.exit(mount.source)
`, m.MountName, m.MountName)
	resp, err := client.Commands().Execute(clusterId, "python", iamMountCommand)
	if err != nil {
		return "", err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return "", errors.New(resp.Results.Summary)
	}
	if resp.Results.ResultType == "text" && resp.Results.Data.(string) == "" {
		return "", errors.New("Unable to find mount point!")
	}

	containerName, storageAccountName, directory, err := ProcessAzureWasbAbfssUris(resp.Results.Data.(string))
	if err != nil {
		return "", err
	}
	if resp.Results.ResultType == "text" && containerName != m.ContainerName &&
		m.StorageAccountName != storageAccountName && m.Directory != directory {
		return "", errors.New(fmt.Sprintf("Does not match uri with storage account and container values!"+
			" %s@%s != %s!", m.ContainerName, m.StorageAccountName, resp.Results.Data.(string)))
	}
	return resp.Results.Data.(string), nil
}

func ProcessAzureAdlsGen1Uri(uri string) (string, string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", "", err
	}
	storageAccount := strings.Split(u.Host, ".")[0]
	var directory string
	if len(u.Path) > 1 {
		directory = u.Path
	} else {
		directory = ""
	}
	return storageAccount, directory, nil
}

func ProcessAzureWasbAbfssUris(uri string) (string, string, string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", "", "", err
	}
	containerName := u.User.String()
	storageAccount := strings.Split(u.Host, ".")[0]
	var directory string
	if len(u.Path) > 1 {
		directory = u.Path
	} else {
		directory = ""
	}
	return containerName, storageAccount, directory, nil
}
