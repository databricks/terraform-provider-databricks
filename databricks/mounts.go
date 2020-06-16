package databricks

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/databrickslabs/databricks-terraform/client/service"
)

// Mount interface describes the functionality of any mount which is create, read and delete
type Mount interface {
	Create(client *service.DBApiClient, clusterID string) error
	Delete(client *service.DBApiClient, clusterID string) error
	Read(client *service.DBApiClient, clusterID string) (string, error)
}

// AWSIamMount describes the object for a aws mount using iam role
type AWSIamMount struct {
	Mount
	S3BucketName string
	MountName    string
}

// NewAWSIamMount constructor for AWSIamMount
func NewAWSIamMount(s3BucketName string, mountName string) *AWSIamMount {
	return &AWSIamMount{S3BucketName: s3BucketName, MountName: mountName}
}

// Create creates an aws iam mount given a cluster ID
func (m AWSIamMount) Create(exec service.CommandExecutor, clusterID string) error {
	iamMountCommand := fmt.Sprintf(`
dbutils.fs.mount("s3a://%s", "/mnt/%s")
dbutils.fs.ls("/mnt/%s")
dbutils.notebook.exit("success")
`, m.S3BucketName, m.MountName, m.MountName)
	resp, err := exec.Execute(clusterID, "python", iamMountCommand)
	if err != nil {
		return err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return errors.New(resp.Results.Summary)
	}
	return nil
}

// Delete deletes an aws iam mount given a cluster ID
func (m AWSIamMount) Delete(exec service.CommandExecutor, clusterID string) error {
	iamMountCommand := fmt.Sprintf(`
dbutils.fs.unmount("/mnt/%s")
dbutils.fs.refreshMounts()
dbutils.notebook.exit("success")
`, m.MountName)
	resp, err := exec.Execute(clusterID, "python", iamMountCommand)
	if err != nil {
		return err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return errors.New(resp.Results.Summary)
	}
	return nil
}

// Read verifies an aws iam mount given a cluster ID
func (m AWSIamMount) Read(exec service.CommandExecutor, clusterID string) (string, error) {
	iamMountCommand := fmt.Sprintf(`
dbutils.fs.ls("/mnt/%s")
for mount in dbutils.fs.mounts():
  if mount.mountPoint == "/mnt/%s":
    dbutils.notebook.exit(mount.source)
`, m.MountName, m.MountName)
	resp, err := exec.Execute(clusterID, "python", iamMountCommand)
	if err != nil {
		return "", err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return "", errors.New(resp.Results.Summary)
	}
	if resp.Results.ResultType == "text" && resp.Results.Data.(string) == "" {
		return "", errors.New("unable to find mount point")
	}
	respBucket := strings.Replace(resp.Results.Data.(string), "s3a://", "", 1)
	if resp.Results.ResultType == "text" && respBucket != m.S3BucketName {
		return "", fmt.Errorf("does not match bucket value! %s != %s", resp.Results.Data.(string), m.S3BucketName)
	}
	return respBucket, nil
}

// AzureBlobMount describes the object for a azure blob storage mount
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

// NewAzureBlobMount constructor for AzureBlobMount
func NewAzureBlobMount(containerName string, storageAccountName string, directory string, mountName string, authType string, secretScope string, secretKey string) *AzureBlobMount {
	return &AzureBlobMount{ContainerName: containerName, StorageAccountName: storageAccountName, Directory: directory, MountName: mountName, AuthType: authType, SecretScope: secretScope, SecretKey: secretKey}
}

// Create creates a azure blob storage mount given a cluster id
func (m AzureBlobMount) Create(exec service.CommandExecutor, clusterID string) error {
	var confKey string

	if m.AuthType == "SAS" {
		confKey = fmt.Sprintf("fs.azure.sas.%s.%s.blob.core.windows.net", m.ContainerName, m.StorageAccountName)
	} else {
		confKey = fmt.Sprintf("fs.azure.account.key.%s.blob.core.windows.net", m.StorageAccountName)
	}
	iamMountCommand := fmt.Sprintf(`
for mount in dbutils.fs.mounts():
  if mount.mountPoint == "/mnt/%[4]s" and mount.source=="wasbs://%[1]s@%[2]s.blob.core.windows.net%[3]s":
    print ("Mount already exists")
    dbutils.notebook.exit("success")

try:
  dbutils.fs.mount(
    source = "wasbs://%[1]s@%[2]s.blob.core.windows.net%[3]s",
    mount_point = "/mnt/%[4]s",
    extra_configs = {"%[5]s":dbutils.secrets.get(scope = "%[6]s", key = "%[7]s")})
except Exception as e:
  dbutils.fs.unmount("/mnt/%[4]s")
  raise e
dbutils.notebook.exit("success")
`, m.ContainerName, m.StorageAccountName, m.Directory, m.MountName, confKey, m.SecretScope, m.SecretKey)
	resp, err := exec.Execute(clusterID, "python", iamMountCommand)
	if err != nil {
		return err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return errors.New(resp.Results.Summary)
	}
	return nil
}

// Delete deletes a azure blob storage mount given a cluster id
func (m AzureBlobMount) Delete(exec service.CommandExecutor, clusterID string) error {
	iamMountCommand := fmt.Sprintf(`
dbutils.fs.unmount("/mnt/%s")
dbutils.fs.refreshMounts()
dbutils.notebook.exit("success")
`, m.MountName)
	resp, err := exec.Execute(clusterID, "python", iamMountCommand)
	if err != nil {
		return err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return errors.New(resp.Results.Summary)
	}
	return nil
}

// Read verifies a azure blob storage mount given a cluster id
func (m AzureBlobMount) Read(exec service.CommandExecutor, clusterID string) (string, error) {
	iamMountCommand := fmt.Sprintf(`
for mount in dbutils.fs.mounts():
  if mount.mountPoint == "/mnt/%s":
    dbutils.notebook.exit(mount.source)
`, m.MountName)
	resp, err := exec.Execute(clusterID, "python", iamMountCommand)
	if err != nil {
		return "", err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return "", errors.New(resp.Results.Summary)
	}
	if resp.Results.ResultType == "text" && resp.Results.Data.(string) == "" {
		return "", errors.New("unable to find mount point")
	}

	container, storageAccount, directory, err := ProcessAzureWasbAbfssUris(resp.Results.Data.(string))
	if err != nil {
		return "", err
	}
	if resp.Results.ResultType == "text" && container != m.ContainerName &&
		storageAccount != m.StorageAccountName &&
		m.Directory != directory {
		return "", fmt.Errorf("does not match uri with storage account and container values"+
			" %s@%s != %s!", m.ContainerName, m.StorageAccountName, resp.Results.Data.(string))
	}
	return resp.Results.Data.(string), nil
}

// AzureADLSGen1Mount describes the object for a azure datalake gen 1 storage mount
type AzureADLSGen1Mount struct {
	Mount
	StorageResource string
	Directory       string
	MountName       string
	PrefixType      string
	ClientID        string
	TenantID        string
	SecretScope     string
	SecretKey       string
}

// NewAzureADLSGen1Mount constructor for AzureADLSGen1Mount
func NewAzureADLSGen1Mount(storageResource string, directory string, mountName string, prefixType string, clientID string, tenantID string, secretScope string, secretKey string) *AzureADLSGen1Mount {
	return &AzureADLSGen1Mount{StorageResource: storageResource, Directory: directory, MountName: mountName, PrefixType: prefixType, ClientID: clientID, TenantID: tenantID, SecretScope: secretScope, SecretKey: secretKey}
}

// Create creates a azure datalake gen 1 storage mount given a cluster id
func (m AzureADLSGen1Mount) Create(exec service.CommandExecutor, clusterID string) error {
	iamMountCommand := fmt.Sprintf(`
for mount in dbutils.fs.mounts():
  if mount.mountPoint == "/mnt/%[8]s" and mount.source=="adl://%[6]s.azuredatalakestore.net%[7]s":
    print ("Mount already exists")
    dbutils.notebook.exit("success")

try:
  configs = {"%[1]s.oauth2.access.token.provider.type": "ClientCredential",
          "%[1]s.oauth2.client.id": "%[2]s",
          "%[1]s.oauth2.credential": dbutils.secrets.get(scope = "%[3]s", key = "%[4]s"),
          "%[1]s.oauth2.refresh.url": "https://login.microsoftonline.com/%[5]s/oauth2/token"}
  dbutils.fs.mount(
   source = "adl://%[6]s.azuredatalakestore.net%[7]s",
   mount_point = "/mnt/%[8]s",
   extra_configs = configs)
except Exception as e:
  dbutils.fs.unmount("/mnt/%[8]s")
  raise e
dbutils.notebook.exit("success")
`, m.PrefixType, m.ClientID, m.SecretScope, m.SecretKey, m.TenantID, m.StorageResource, m.Directory, m.MountName, m.MountName)
	resp, err := exec.Execute(clusterID, "python", iamMountCommand)
	if err != nil {
		return err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return errors.New(resp.Results.Summary)
	}
	return nil
}

// Delete deletes a azure datalake gen 1 storage mount given a cluster id
func (m AzureADLSGen1Mount) Delete(exec service.CommandExecutor, clusterID string) error {
	iamMountCommand := fmt.Sprintf(`
dbutils.fs.unmount("/mnt/%s")
dbutils.fs.refreshMounts()
dbutils.notebook.exit("success")
`, m.MountName)
	resp, err := exec.Execute(clusterID, "python", iamMountCommand)
	if err != nil {
		return err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return errors.New(resp.Results.Summary)
	}
	return nil
}

// Read verifies the azure datalake gen 1 storage mount given a cluster id
func (m AzureADLSGen1Mount) Read(exec service.CommandExecutor, clusterID string) (string, error) {
	iamMountCommand := fmt.Sprintf(`
for mount in dbutils.fs.mounts():
  if mount.mountPoint == "/mnt/%s":
    dbutils.notebook.exit(mount.source)
`, m.MountName)
	resp, err := exec.Execute(clusterID, "python", iamMountCommand)
	if err != nil {
		return "", err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return "", errors.New(resp.Results.Summary)
	}
	if resp.Results.ResultType == "text" && resp.Results.Data.(string) == "" {
		return "", errors.New("unable to find mount point")
	}

	storageResource, directory, err := ProcessAzureAdlsGen1Uri(resp.Results.Data.(string))
	if err != nil {
		return "", err
	}
	if resp.Results.ResultType == "text" && storageResource != m.StorageResource &&
		m.Directory != directory {
		return "", fmt.Errorf("does not match uri with storage account and container values"+
			" %s@%s != %s!", m.StorageResource, m.Directory, resp.Results.Data.(string))
	}
	return resp.Results.Data.(string), nil
}

// AzureADLSGen2Mount describes the object for a azure datalake gen 2 storage mount
type AzureADLSGen2Mount struct {
	Mount
	ContainerName        string
	StorageAccountName   string
	Directory            string
	MountName            string
	ClientID             string
	TenantID             string
	SecretScope          string
	SecretKey            string
	InitializeFileSystem bool
}

// NewAzureADLSGen2Mount is a constructor for AzureADLSGen2Mount
func NewAzureADLSGen2Mount(containerName string, storageAccountName string, directory string, mountName string, clientID string, tenantID string, secretScope string, secretKey string, initializeFileSystem bool) *AzureADLSGen2Mount {
	return &AzureADLSGen2Mount{ContainerName: containerName, StorageAccountName: storageAccountName, Directory: directory, MountName: mountName, ClientID: clientID, TenantID: tenantID, SecretScope: secretScope, SecretKey: secretKey, InitializeFileSystem: initializeFileSystem}
}

// Create creates a azure datalake gen 2 storage mount
func (m AzureADLSGen2Mount) Create(exec service.CommandExecutor, clusterID string) error {
	iamMountCommand := fmt.Sprintf(`
for mount in dbutils.fs.mounts():
  if mount.mountPoint == "/mnt/%[9]s" and mount.source=="abfss://%[6]s@%[7]s.dfs.core.windows.net%[8]s":
    print ("Mount already exists")
    dbutils.notebook.exit("success")

try:
  configs = {"fs.azure.account.auth.type": "OAuth",
           "fs.azure.account.oauth.provider.type": "org.apache.hadoop.fs.azurebfs.oauth2.ClientCredsTokenProvider",
           "fs.azure.account.oauth2.client.id": "%[1]s",
           "fs.azure.account.oauth2.client.secret": dbutils.secrets.get(scope = "%[2]s", key = "%[3]s"),
		   "fs.azure.account.oauth2.client.endpoint": "https://login.microsoftonline.com/%[4]s/oauth2/token",
		   "fs.azure.createRemoteFileSystemDuringInitialization": "%[5]t"}
  dbutils.fs.mount(
   source = "abfss://%[6]s@%[7]s.dfs.core.windows.net%[8]s",
   mount_point = "/mnt/%[9]s",
   extra_configs = configs)
except Exception as e:
  try:
    dbutils.fs.unmount("/mnt/%[9]s")
  except Exception as e2:
    print ("Failed to unmount", e2)
  raise e
dbutils.notebook.exit("success")
`, m.ClientID, m.SecretScope, m.SecretKey, m.TenantID, m.InitializeFileSystem, m.ContainerName, m.StorageAccountName, m.Directory, m.MountName)
	resp, err := exec.Execute(clusterID, "python", iamMountCommand)
	if err != nil {
		return err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return errors.New(resp.Results.Summary)
	}
	return nil
}

// Delete deletes a azure datalake gen 2 storage mount
func (m AzureADLSGen2Mount) Delete(exec service.CommandExecutor, clusterID string) error {
	iamMountCommand := fmt.Sprintf(`
dbutils.fs.unmount("/mnt/%s")
dbutils.fs.refreshMounts()
dbutils.notebook.exit("success")
`, m.MountName)
	resp, err := exec.Execute(clusterID, "python", iamMountCommand)
	if err != nil {
		return err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return errors.New(resp.Results.Summary)
	}
	return nil
}

// Read verifies the azure datalake gen 2 storage mount
func (m AzureADLSGen2Mount) Read(exec service.CommandExecutor, clusterID string) (string, error) {
	iamMountCommand := fmt.Sprintf(`
for mount in dbutils.fs.mounts():
  if mount.mountPoint == "/mnt/%s":
    dbutils.notebook.exit(mount.source)
`, m.MountName)
	resp, err := exec.Execute(clusterID, "python", iamMountCommand)
	if err != nil {
		return "", err
	}
	if resp.Results.ResultType == "error" {
		log.Println(fmt.Sprintf("[ERROR] [CAUSED BY] %s", resp.Results.Cause))
		return "", errors.New(resp.Results.Summary)
	}
	if resp.Results.ResultType == "text" && resp.Results.Data.(string) == "" {
		return "", errors.New("unable to find mount point")
	}

	containerName, storageAccountName, directory, err := ProcessAzureWasbAbfssUris(resp.Results.Data.(string))
	if err != nil {
		return "", err
	}
	if resp.Results.ResultType == "text" && containerName != m.ContainerName &&
		m.StorageAccountName != storageAccountName && m.Directory != directory {
		return "", fmt.Errorf("does not match uri with storage account and container values"+
			" %s@%s != %s!", m.ContainerName, m.StorageAccountName, resp.Results.Data.(string))
	}
	return resp.Results.Data.(string), nil
}

// ProcessAzureAdlsGen1Uri will return given a adls gen 1 URI the storage account name and the directory if it exists
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

// ProcessAzureWasbAbfssUris will return given a WASBS or ABFSS URI the
// containerName, storageAccount and the directory if it exists
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

// ValidateMountDirectory is a ValidateFunc that ensures the mount directory starts with a '/'
func ValidateMountDirectory(val interface{}, key string) (warns []string, errs []error) {
	v := val.(string)
	if v != "" && !strings.HasPrefix(v, "/") {
		return nil, []error{fmt.Errorf("%s must start with /, got: %s", key, v)}
	}
	return nil, nil
}
