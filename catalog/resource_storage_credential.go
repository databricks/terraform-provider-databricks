package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var storageCredentialSchema = common.StructToSchema(StorageCredentialInfo{},
	func(m map[string]*schema.Schema) map[string]*schema.Schema {
		m["storage_credential_id"] = &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		}
		return adjustDataAccessSchema(m)
	})

func ResourceStorageCredential() common.Resource {
	return common.Resource{
		Schema: storageCredentialSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			metastoreId := d.Get("metastore_id").(string)
			tmpSchema := removeGcpSaField(storageCredentialSchema)

			var create catalog.CreateStorageCredential
			var update catalog.UpdateStorageCredential
			common.DataToStructPointer(d, tmpSchema, &create)
			common.DataToStructPointer(d, tmpSchema, &update)
			update.Name = d.Get("name").(string)

			//manually add empty struct back for databricks_gcp_service_account
			if _, ok := d.GetOk("databricks_gcp_service_account"); ok {
				create.DatabricksGcpServiceAccount = &catalog.DatabricksGcpServiceAccountRequest{}
			}

			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				storageCredential, err := acc.StorageCredentials.Create(ctx,
					catalog.AccountsCreateStorageCredential{
						MetastoreId:    metastoreId,
						CredentialInfo: &create,
					})
				if err != nil {
					return err
				}
				d.SetId(storageCredential.CredentialInfo.Name)

				// Don't update owner if it is not provided
				if d.Get("owner") == "" {
					return nil
				}
				update.Name = d.Id()
				_, err = acc.StorageCredentials.Update(ctx, catalog.AccountsUpdateStorageCredential{
					CredentialInfo:        &update,
					MetastoreId:           metastoreId,
					StorageCredentialName: storageCredential.CredentialInfo.Name,
				})
				if err != nil {
					return err
				}
				return nil
			}, func(w *databricks.WorkspaceClient) error {
				err := validateMetastoreId(ctx, w, d.Get("metastore_id").(string))
				if err != nil {
					return err
				}
				storageCredential, err := w.StorageCredentials.Create(ctx, create)
				if err != nil {
					return err
				}
				d.SetId(storageCredential.Name)

				// Don't update owner if it is not provided
				if d.Get("owner") == "" {
					return nil
				}

				update.Name = d.Id()
				_, err = w.StorageCredentials.Update(ctx, update)
				if err != nil {
					return err
				}
				return nil
			})
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				storageCredential, err := acc.StorageCredentials.Get(ctx, catalog.GetAccountStorageCredentialRequest{
					MetastoreId:           d.Get("metastore_id").(string),
					StorageCredentialName: d.Id(),
				})
				if err != nil {
					return err
				}
				// azure client secret is sensitive, so we need to preserve it
				var scOrig catalog.CreateStorageCredential
				common.DataToStructPointer(d, storageCredentialSchema, &scOrig)
				if scOrig.AzureServicePrincipal != nil {
					if scOrig.AzureServicePrincipal.ClientSecret != "" {
						storageCredential.CredentialInfo.AzureServicePrincipal.ClientSecret = scOrig.AzureServicePrincipal.ClientSecret
					}
				}
				err = common.StructToData(storageCredential.CredentialInfo, storageCredentialSchema, d)
				if err != nil {
					return err
				}
				d.Set("storage_credential_id", storageCredential.CredentialInfo.Id)
				return nil
			}, func(w *databricks.WorkspaceClient) error {
				storageCredential, err := w.StorageCredentials.GetByName(ctx, d.Id())
				if err != nil {
					return err
				}
				// azure client secret is sensitive, so we need to preserve it
				var scOrig catalog.CreateStorageCredential
				common.DataToStructPointer(d, storageCredentialSchema, &scOrig)
				if scOrig.AzureServicePrincipal != nil {
					if scOrig.AzureServicePrincipal.ClientSecret != "" {
						storageCredential.AzureServicePrincipal.ClientSecret = scOrig.AzureServicePrincipal.ClientSecret
					}
				}
				err = common.StructToData(storageCredential, storageCredentialSchema, d)
				if err != nil {
					return err
				}
				d.Set("storage_credential_id", storageCredential.Id)
				return nil
			})
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var update catalog.UpdateStorageCredential
			force := d.Get("force_update").(bool)
			common.DataToStructPointer(d, storageCredentialSchema, &update)
			update.Name = d.Id()
			update.Force = force
			if _, ok := d.GetOk("azure_managed_identity"); ok {
				update.AzureManagedIdentity.CredentialId = ""
			}
			return updateStorageCredential(ctx, d, c, d.Get("metastore_id").(string), d.Id(), update)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return deleteStorageCredential(ctx, c, d.Get("metastore_id").(string), d.Id(), d.Get("force_destroy").(bool))
		},
	}
}
