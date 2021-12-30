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

type StorageCredentialInfo struct {
	Name        string                 `json:"name" tf:"force_new"`
	Comment     string                 `json:"comment,omitempty" tf:"force_new"`
	Aws         *AwsIamRole            `json:"aws_iam_role,omitempty" tf:"group:access"`
	Azure       *AzureServicePrincipal `json:"azure_service_principal,omitempty" tf:"group:access"`
	MetastoreID string                 `json:"metastore_id,omitempty" tf:"computed"`
}

func (a StorageCredentialsAPI) create(sci *StorageCredentialInfo) error {
	return a.client.Post(a.context, "/unity-catalog/storage-credentials", sci, &sci)
}

func (a StorageCredentialsAPI) get(id string) (sci StorageCredentialInfo, err error) {
	err = a.client.Get(a.context, "/unity-catalog/storage-credentials/"+id, nil, &sci)
	return
}

func (a StorageCredentialsAPI) update(name string, sci StorageCredentialInfo) error {
	return a.client.Patch(a.context, "/unity-catalog/storage-credentials/"+name, sci)
}

func (a StorageCredentialsAPI) delete(id string) error {
	return a.client.Delete(a.context, "/unity-catalog/storage-credentials/"+id, nil)
}

func ResourceStorageCredential() *schema.Resource {
	s := common.StructToSchema(StorageCredentialInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			alof := []string{"aws_iam_role", "azure_service_principal"}
			m["aws_iam_role"].AtLeastOneOf = alof
			m["azure_service_principal"].AtLeastOneOf = alof
			return m
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var sci StorageCredentialInfo
			common.DataToStructPointer(d, s, &sci)
			err := NewStorageCredentialsAPI(ctx, c).create(&sci)
			if err != nil {
				return err
			}
			d.SetId(sci.Name)
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
			var sci StorageCredentialInfo
			common.DataToStructPointer(d, s, &sci)
			return NewStorageCredentialsAPI(ctx, c).update(d.Id(), StorageCredentialInfo{
				Name:  d.Id(),
				Aws:   sci.Aws,
				Azure: sci.Azure,
			})
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewStorageCredentialsAPI(ctx, c).delete(d.Id())
		},
	}.ToResource()
}
