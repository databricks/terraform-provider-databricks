package catalog

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type StorageCredentialsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewStorageCredentialsAPI(ctx context.Context, m interface{}) StorageCredentialsAPI {
	return StorageCredentialsAPI{m.(*common.DatabricksClient), ctx}
}

type StorageCredentialConfig struct {
	// name is actually the id of the storage credentials, it is used for GET and DELETE ops
	// In a given metastore credentials name is unique
	Name        string                 `json:"name" tf:"force_new"`
	Aws         *AwsIamRole            `json:"aws_iam_role,omitempty" tf:"group:access"`
	Azure       *AzureServicePrincipal `json:"azure_service_principal,omitempty" tf:"group:access"`
	Comment     string                 `json:"comment,omitempty" tf:"force_new"`
	Owner       string                 `json:"owner,omitempty" tf:"computed"`
	MetastoreID string                 `json:"metastore_id,omitempty" tf:"computed"`
	CreatedBy   string                 `json:"created_by,omitempty" tf:"computed"`
}

//func (a StorageCredentialsAPI) listStorageCredentials() (scis []StorageCredentialConfig, err error) {
//	err = a.client.Get(a.context, "/unity-catalog/storage-credentials", nil, &scis)
//	return
//}

func (a StorageCredentialsAPI) create(cm StorageCredentialConfig) (mi StorageCredentialConfig, err error) {
	err = a.client.Post(a.context, "/unity-catalog/storage-credentials", cm, &mi)
	return
}

func (a StorageCredentialsAPI) get(id string) (sci StorageCredentialConfig, err error) {
	err = a.client.Get(a.context, "/unity-catalog/storage-credentials/"+id, nil, &sci)
	return
}

func (a StorageCredentialsAPI) update(storageCredentialsName string, info StorageCredentialConfig) error {
	return a.client.Patch(a.context, "/unity-catalog/storage-credentials/"+storageCredentialsName, info)
}

func (a StorageCredentialsAPI) delete(id string) error {
	return a.client.Delete(a.context, "/unity-catalog/storage-credentials/"+id, nil)
}

func ResourceStorageCredential() *schema.Resource {
	s := common.StructToSchema(StorageCredentialConfig{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			alof := []string{"aws_iam_role", "azure_service_principal"}
			m["aws_iam_role"].AtLeastOneOf = alof
			m["azure_service_principal"].AtLeastOneOf = alof
			return m
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var create StorageCredentialConfig
			common.DataToStructPointer(d, s, &create)
			sc, err := NewStorageCredentialsAPI(ctx, c).create(create)
			if err != nil {
				return err
			}
			d.SetId(sc.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			sci, err := NewStorageCredentialsAPI(ctx, c).get(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(sci, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var sci StorageCredentialConfig
			common.DataToStructPointer(d, s, &sci)
			// only aws or azure options and name can be provided for update, everything else is either computed
			// or force_new
			update := StorageCredentialConfig{
				Name:  d.Id(),
				Aws:   sci.Aws,
				Azure: sci.Azure,
			}
			return NewStorageCredentialsAPI(ctx, c).update(d.Id(), update)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewStorageCredentialsAPI(ctx, c).delete(d.Id())
		},
	}.ToResource()
}
