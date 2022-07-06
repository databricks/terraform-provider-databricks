package catalog

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type StorageCredentialsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewStorageCredentialsAPI(ctx context.Context, m interface{}) StorageCredentialsAPI {
	return StorageCredentialsAPI{m.(*common.DatabricksClient), context.WithValue(ctx, common.Api, common.API_2_1)}
}

type StorageCredentialInfo struct {
	Name        string                 `json:"name" tf:"force_new"`
	Owner       string                 `json:"owner,omitempty" tf:"computed"`
	Comment     string                 `json:"comment,omitempty"`
	Aws         *AwsIamRole            `json:"aws_iam_role,omitempty" tf:"group:access"`
	Azure       *AzureServicePrincipal `json:"azure_service_principal,omitempty" tf:"group:access"`
	AzMI        *AzureManagedIdentity  `json:"azure_managed_identity,omitempty" tf:"group:access"`
	MetastoreID string                 `json:"metastore_id,omitempty" tf:"computed"`
}

func (a StorageCredentialsAPI) create(sci *StorageCredentialInfo) error {
	return a.client.Post(a.context, "/unity-catalog/storage-credentials", sci, &sci)
}

func (a StorageCredentialsAPI) get(id string) (sci StorageCredentialInfo, err error) {
	err = a.client.Get(a.context, "/unity-catalog/storage-credentials/"+id, nil, &sci)
	return
}

func (a StorageCredentialsAPI) delete(id string) error {
	return a.client.Delete(a.context, "/unity-catalog/storage-credentials/"+id, nil)
}

func ResourceStorageCredential() *schema.Resource {
	s := common.StructToSchema(StorageCredentialInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			alof := []string{"aws_iam_role", "azure_service_principal", "azure_managed_identity"}
			m["aws_iam_role"].AtLeastOneOf = alof
			m["azure_service_principal"].AtLeastOneOf = alof
			m["azure_managed_identity"].AtLeastOneOf = alof
			return m
		})
	update := updateFunctionFactory("/unity-catalog/storage-credentials", []string{
		"owner", "comment", "aws_iam_role", "azure_service_principal"})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var sci StorageCredentialInfo
			common.DataToStructPointer(d, s, &sci)
			sci.Owner = ""
			err := NewStorageCredentialsAPI(ctx, c).create(&sci)
			if err != nil {
				return err
			}
			d.SetId(sci.Name)
			return update(ctx, d, c)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			sci, err := NewStorageCredentialsAPI(ctx, c).get(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(sci, s, d)
		},
		Update: update,
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewStorageCredentialsAPI(ctx, c).delete(d.Id())
		},
	}.ToResource()
}
