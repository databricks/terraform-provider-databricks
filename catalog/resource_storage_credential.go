package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/catalog/bindings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type StorageCredentialInfo struct {
	Name                        string                                       `json:"name" tf:"force_new"`
	Owner                       string                                       `json:"owner,omitempty" tf:"computed"`
	Comment                     string                                       `json:"comment,omitempty"`
	Aws                         *catalog.AwsIamRoleResponse                  `json:"aws_iam_role,omitempty" tf:"group:access"`
	Azure                       *catalog.AzureServicePrincipal               `json:"azure_service_principal,omitempty" tf:"group:access"`
	AzMI                        *catalog.AzureManagedIdentityResponse        `json:"azure_managed_identity,omitempty" tf:"group:access"`
	GcpSAKey                    *GcpServiceAccountKey                        `json:"gcp_service_account_key,omitempty" tf:"group:access"`
	DatabricksGcpServiceAccount *catalog.DatabricksGcpServiceAccountResponse `json:"databricks_gcp_service_account,omitempty" tf:"computed"`
	MetastoreID                 string                                       `json:"metastore_id,omitempty" tf:"computed"`
	ReadOnly                    bool                                         `json:"read_only,omitempty"`
	SkipValidation              bool                                         `json:"skip_validation,omitempty"`
	IsolationMode               string                                       `json:"isolation_mode,omitempty" tf:"computed"`
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
				// Bind the current workspace if the storage credential is isolated, otherwise the read will fail
				return bindings.AddCurrentWorkspaceBindings(ctx, d, w, storageCredential.Name, "storage-credentials")
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
			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				if d.HasChange("owner") {
					_, err := acc.StorageCredentials.Update(ctx, catalog.AccountsUpdateStorageCredential{
						CredentialInfo: &catalog.UpdateStorageCredential{
							Name:  update.Name,
							Owner: update.Owner,
						},
						MetastoreId:           d.Get("metastore_id").(string),
						StorageCredentialName: d.Id(),
					})
					if err != nil {
						return err
					}
				}

				if !d.HasChangeExcept("owner") {
					return nil
				}

				update.Owner = ""
				_, err := acc.StorageCredentials.Update(ctx, catalog.AccountsUpdateStorageCredential{
					CredentialInfo:        &update,
					MetastoreId:           d.Get("metastore_id").(string),
					StorageCredentialName: d.Id(),
				})
				if err != nil {
					if d.HasChange("owner") {
						// Rollback
						old, new := d.GetChange("owner")
						_, rollbackErr := acc.StorageCredentials.Update(ctx, catalog.AccountsUpdateStorageCredential{
							CredentialInfo: &catalog.UpdateStorageCredential{
								Name:  update.Name,
								Owner: old.(string),
							},
							MetastoreId:           d.Get("metastore_id").(string),
							StorageCredentialName: d.Id(),
						})
						if rollbackErr != nil {
							return common.OwnerRollbackError(err, rollbackErr, old.(string), new.(string))
						}
					}
					return err
				}
				return nil
			}, func(w *databricks.WorkspaceClient) error {
				err := validateMetastoreId(ctx, w, d.Get("metastore_id").(string))
				if err != nil {
					return err
				}
				if d.HasChange("owner") {
					_, err := w.StorageCredentials.Update(ctx, catalog.UpdateStorageCredential{
						Name:  update.Name,
						Owner: update.Owner,
					})
					if err != nil {
						return err
					}
				}

				if !d.HasChangeExcept("owner") {
					return nil
				}

				update.Owner = ""
				_, err = w.StorageCredentials.Update(ctx, update)
				if err != nil {
					if d.HasChange("owner") {
						// Rollback
						old, new := d.GetChange("owner")
						_, rollbackErr := w.StorageCredentials.Update(ctx, catalog.UpdateStorageCredential{
							Name:  update.Name,
							Owner: old.(string),
						})
						if rollbackErr != nil {
							return common.OwnerRollbackError(err, rollbackErr, old.(string), new.(string))
						}
					}
					return err
				}
				// Bind the current workspace if the storage credential is isolated, otherwise the read will fail
				return bindings.AddCurrentWorkspaceBindings(ctx, d, w, update.Name, "storage-credentials")
			})
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			force := d.Get("force_destroy").(bool)
			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				return acc.StorageCredentials.Delete(ctx, catalog.DeleteAccountStorageCredentialRequest{
					Force:                 force,
					StorageCredentialName: d.Id(),
					MetastoreId:           d.Get("metastore_id").(string),
				})
			}, func(w *databricks.WorkspaceClient) error {
				err := validateMetastoreId(ctx, w, d.Get("metastore_id").(string))
				if err != nil {
					return err
				}
				return w.StorageCredentials.Delete(ctx, catalog.DeleteStorageCredentialRequest{
					Force: force,
					Name:  d.Id(),
				})
			})
		},
	}
}
