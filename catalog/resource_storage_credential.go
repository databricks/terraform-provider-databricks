package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type StorageCredentialInfo struct {
	Name        string                 `json:"name" tf:"force_new"`
	Owner       string                 `json:"owner,omitempty" tf:"computed"`
	Comment     string                 `json:"comment,omitempty"`
	Aws         *AwsIamRole            `json:"aws_iam_role,omitempty" tf:"group:access"`
	Azure       *AzureServicePrincipal `json:"azure_service_principal,omitempty" tf:"group:access"`
	AzMI        *AzureManagedIdentity  `json:"azure_managed_identity,omitempty" tf:"group:access"`
	GcpSAKey    *GcpServiceAccountKey  `json:"gcp_service_account_key,omitempty" tf:"group:access"`
	DBGcpSA     *DbGcpServiceAccount   `json:"databricks_gcp_service_account,omitempty" tf:"computed"`
	MetastoreID string                 `json:"metastore_id,omitempty" tf:"computed"`
	ReadOnly    bool                   `json:"read_only,omitempty"`
}

func removeGcpSaField(originalSchema map[string]*schema.Schema) map[string]*schema.Schema {
	//common.DataToStructPointer(d, s, &create) will error out because of DatabricksGcpServiceAccount any
	tmpSchema := make(map[string]*schema.Schema)
	for k, v := range originalSchema {
		tmpSchema[k] = v
	}
	delete(tmpSchema, "databricks_gcp_service_account")
	return tmpSchema
}

func ResourceStorageCredential() *schema.Resource {
	s := common.StructToSchema(StorageCredentialInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["force_destroy"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			}
			m["aws_iam_role"].AtLeastOneOf = alofCred
			m["azure_service_principal"].AtLeastOneOf = alofCred
			m["azure_managed_identity"].AtLeastOneOf = alofCred
			m["gcp_service_account_key"].AtLeastOneOf = alofCred
			m["databricks_gcp_service_account"].AtLeastOneOf = alofCred

			// suppress changes for private_key
			m["gcp_service_account_key"].DiffSuppressFunc = SuppressGcpSAKeyDiff

			return m
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			metastoreId := d.Get("metastore_id").(string)
			tmpSchema := removeGcpSaField(s)

			var create catalog.CreateStorageCredential
			var update catalog.UpdateStorageCredential
			common.DataToStructPointer(d, tmpSchema, &create)
			common.DataToStructPointer(d, tmpSchema, &update)

			var storageCredential *catalog.StorageCredentialInfo
			if c.Config.IsAccountClient() {
				acc, err := c.AccountClient()
				if err != nil {
					return err
				}
				storageCredential, err = acc.StorageCredentials.Create(ctx,
					catalog.AccountsCreateStorageCredential{
						MetastoreId:    metastoreId,
						CredentialInfo: &create,
					})
				if err != nil {
					return err
				}
				d.SetId(storageCredential.Name)
				_, err = acc.StorageCredentials.Update(ctx, catalog.AccountsUpdateStorageCredential{
					CredentialInfo: &update,
					MetastoreId:    metastoreId,
					Name:           storageCredential.Id,
				})
				if err != nil {
					return err
				}
			} else {
				w, err := c.WorkspaceClient()
				if err != nil {
					return err
				}
				storageCredential, err = w.StorageCredentials.Create(ctx, create)
				if err != nil {
					return err
				}
				d.SetId(storageCredential.Name)
				_, err = w.StorageCredentials.Update(ctx, update)
				if err != nil {
					return err
				}
			}
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			if c.Config.IsAccountClient() {
				acc, err := c.AccountClient()
				if err != nil {
					return err
				}
				storageCredential, err := acc.StorageCredentials.Get(ctx, catalog.GetAccountStorageCredentialRequest{
					MetastoreId: d.Get("metastore_id").(string),
					Name:        d.Id(),
				})
				if err != nil {
					return err
				}
				return common.StructToData(storageCredential, s, d)
			} else {
				w, err := c.WorkspaceClient()
				if err != nil {
					return err
				}

				storageCredential, err := w.StorageCredentials.GetByName(ctx, d.Id())
				if err != nil {
					return err
				}
				return common.StructToData(storageCredential, s, d)
			}
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var update catalog.UpdateStorageCredential
			common.DataToStructPointer(d, s, &update)

			if c.Config.IsAccountClient() {
				acc, err := c.AccountClient()
				if err != nil {
					return err
				}
				_, err = acc.StorageCredentials.Update(ctx, catalog.AccountsUpdateStorageCredential{
					CredentialInfo: &update,
					MetastoreId:    d.Get("metastore_id").(string),
					Name:           d.Id(),
				})
				if err != nil {
					return err
				}
			} else {
				w, err := c.WorkspaceClient()
				if err != nil {
					return err
				}
				_, err = w.StorageCredentials.Update(ctx, update)
				if err != nil {
					return err
				}
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			force := d.Get("force_destroy").(bool)
			if c.Config.IsAccountClient() {
				acc, err := c.AccountClient()
				if err != nil {
					return err
				}
				return acc.StorageCredentials.Delete(ctx, catalog.DeleteAccountStorageCredentialRequest{
					Name:        d.Id(),
					MetastoreId: d.Get("metastore_id").(string),
				})
			} else {
				w, err := c.WorkspaceClient()
				if err != nil {
					return err
				}
				return w.StorageCredentials.Delete(ctx, catalog.DeleteStorageCredentialRequest{
					Force: force,
					Name:  d.Id(),
				})
			}
		},
	}.ToResource()
}
