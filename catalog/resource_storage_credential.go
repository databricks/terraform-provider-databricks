package catalog

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/catalog/bindings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type StorageCredentialInfo struct {
	common.Namespace
	Name                        string                                       `json:"name" tf:"force_new"`
	Owner                       string                                       `json:"owner,omitempty" tf:"computed"`
	Comment                     string                                       `json:"comment,omitempty"`
	Aws                         *catalog.AwsIamRoleResponse                  `json:"aws_iam_role,omitempty" tf:"group:access"`
	Azure                       *catalog.AzureServicePrincipal               `json:"azure_service_principal,omitempty" tf:"group:access"`
	AzMI                        *catalog.AzureManagedIdentityResponse        `json:"azure_managed_identity,omitempty" tf:"group:access"`
	GcpSAKey                    *GcpServiceAccountKey                        `json:"gcp_service_account_key,omitempty" tf:"group:access"`
	DatabricksGcpServiceAccount *catalog.DatabricksGcpServiceAccountResponse `json:"databricks_gcp_service_account,omitempty" tf:"computed"`
	CloudflareApiToken          *catalog.CloudflareApiToken                  `json:"cloudflare_api_token,omitempty" tf:"group:access"`
	MetastoreID                 string                                       `json:"metastore_id,omitempty" tf:"computed"`
	ReadOnly                    bool                                         `json:"read_only,omitempty"`
	SkipValidation              bool                                         `json:"skip_validation,omitempty"`
	IsolationMode               string                                       `json:"isolation_mode,omitempty" tf:"computed"`
}

var storageCredentialSchema = common.StructToSchema(StorageCredentialInfo{},
	func(m map[string]*schema.Schema) map[string]*schema.Schema {
		m["storage_credential_id"] = &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		}
		common.MustSchemaPath(m, "databricks_gcp_service_account", "email").Computed = true
		common.MustSchemaPath(m, "databricks_gcp_service_account", "credential_id").Computed = true
		common.NamespaceCustomizeSchemaMap(m)
		return adjustDataAccessSchema(m)
	})

// parseStorageCredentialId parses the resource ID to extract metastore_id and storage_credential_name
// for composite IDs in the format "metastore_id|storage_credential_name" (account-level)
// or just returns the ID as is for workspace-level resources
func parseStorageCredentialId(d *schema.ResourceData) (metastoreId, storageCredentialName string, err error) {
	id := d.Id()
	parts := strings.Split(id, "|")

	if len(parts) == 2 {
		// Account-level format: metastore_id|storage_credential_name
		metastoreId = parts[0]
		storageCredentialName = parts[1]

		// Set the metastore_id in the state if not already set
		if d.Get("metastore_id").(string) == "" {
			if err := d.Set("metastore_id", metastoreId); err != nil {
				return "", "", fmt.Errorf("failed to set metastore_id: %w", err)
			}
		}

		// Update the resource ID to just the storage credential name
		d.SetId(storageCredentialName)
		return metastoreId, storageCredentialName, nil
	} else if len(parts) == 1 {
		// Workspace-level format: just the storage credential name
		// Get metastore_id from the existing state
		metastoreId = d.Get("metastore_id").(string)
		storageCredentialName = id
		return metastoreId, storageCredentialName, nil
	} else {
		return "", "", fmt.Errorf("invalid storage credential ID format: expected 'name' or 'metastore_id|name', got '%s'", id)
	}
}

func ResourceStorageCredential() common.Resource {
	return common.Resource{
		Schema: storageCredentialSchema,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			return common.NamespaceCustomizeDiff(d)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			metastoreId := d.Get("metastore_id").(string)

			var create catalog.CreateStorageCredential
			var update catalog.UpdateStorageCredential
			common.DataToStructPointer(d, storageCredentialSchema, &create)
			common.DataToStructPointer(d, storageCredentialSchema, &update)
			update.Name = d.Get("name").(string)
			if update.DatabricksGcpServiceAccount != nil { // we can't update it at all
				update.DatabricksGcpServiceAccount = nil
			}

			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				storageCredential, err := acc.StorageCredentials.Create(ctx,
					catalog.AccountsCreateStorageCredential{
						MetastoreId:    metastoreId,
						CredentialInfo: toCreateAccountsStorageCredential(&create),
					})
				if err != nil {
					return err
				}
				d.SetId(storageCredential.CredentialInfo.Name)

				// Update owner or isolation mode if it is provided
				if !updateRequired(d, []string{"owner", "isolation_mode"}) {
					return nil
				}

				update.Name = d.Id()
				_, err = acc.StorageCredentials.Update(ctx, catalog.AccountsUpdateStorageCredential{
					CredentialInfo:        toUpdateAccountsStorageCredential(&update),
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

				// Update owner or isolation mode if it is provided
				if !updateRequired(d, []string{"owner", "isolation_mode"}) {
					return nil
				}

				update.Name = d.Id()
				_, err = w.StorageCredentials.Update(ctx, update)
				if err != nil {
					return err
				}
				// Bind the current workspace if the storage credential is isolated, otherwise the read will fail
				return bindings.AddCurrentWorkspaceBindings(ctx, d, w, storageCredential.Name, bindings.BindingsSecurableTypeStorageCredential)
			})
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			// Parse the ID to handle both composite and simple formats
			metastoreId, storageCredentialName, err := parseStorageCredentialId(d)
			if err != nil {
				return err
			}

			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				storageCredential, err := acc.StorageCredentials.Get(ctx, catalog.GetAccountStorageCredentialRequest{
					MetastoreId:           metastoreId,
					StorageCredentialName: storageCredentialName,
				})
				if err != nil {
					return err
				}
				// azure client secret, & r2 secret access key are sensitive, so we need to preserve them
				var scOrig catalog.CreateStorageCredential
				common.DataToStructPointer(d, storageCredentialSchema, &scOrig)
				if scOrig.AzureServicePrincipal != nil {
					if scOrig.AzureServicePrincipal.ClientSecret != "" {
						storageCredential.CredentialInfo.AzureServicePrincipal.ClientSecret = scOrig.AzureServicePrincipal.ClientSecret
					}
				}
				if scOrig.CloudflareApiToken != nil {
					if scOrig.CloudflareApiToken.SecretAccessKey != "" {
						storageCredential.CredentialInfo.CloudflareApiToken.SecretAccessKey = scOrig.CloudflareApiToken.SecretAccessKey
					}
				}
				err = common.StructToData(storageCredential.CredentialInfo, storageCredentialSchema, d)
				if err != nil {
					return err
				}
				d.Set("storage_credential_id", storageCredential.CredentialInfo.Id)
				return nil
			}, func(w *databricks.WorkspaceClient) error {
				storageCredential, err := w.StorageCredentials.GetByName(ctx, storageCredentialName)
				if err != nil {
					return err
				}
				// azure client secret, & r2 secret access key are sensitive, so we need to preserve them
				var scOrig catalog.CreateStorageCredential
				common.DataToStructPointer(d, storageCredentialSchema, &scOrig)
				if scOrig.AzureServicePrincipal != nil {
					if scOrig.AzureServicePrincipal.ClientSecret != "" {
						storageCredential.AzureServicePrincipal.ClientSecret = scOrig.AzureServicePrincipal.ClientSecret
					}
				}
				if scOrig.CloudflareApiToken != nil {
					if scOrig.CloudflareApiToken.SecretAccessKey != "" {
						storageCredential.CloudflareApiToken.SecretAccessKey = scOrig.CloudflareApiToken.SecretAccessKey
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
			// Parse the ID to handle both composite and simple formats
			metastoreId, storageCredentialName, err := parseStorageCredentialId(d)
			if err != nil {
				return err
			}

			var update catalog.UpdateStorageCredential
			force := d.Get("force_update").(bool)
			common.DataToStructPointer(d, storageCredentialSchema, &update)
			update.Name = storageCredentialName
			update.Force = force
			if _, ok := d.GetOk("azure_managed_identity"); ok {
				update.AzureManagedIdentity.CredentialId = ""
			}
			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				if d.HasChange("owner") {
					ownerUpdate := catalog.UpdateStorageCredential{
						Name:  update.Name,
						Owner: update.Owner,
					}
					_, err := acc.StorageCredentials.Update(ctx, catalog.AccountsUpdateStorageCredential{
						CredentialInfo:        toUpdateAccountsStorageCredential(&ownerUpdate),
						MetastoreId:           metastoreId,
						StorageCredentialName: storageCredentialName,
					})
					if err != nil {
						return err
					}
				}

				if !d.HasChangeExcept("owner") {
					return nil
				}

				if d.HasChange("read_only") {
					update.ForceSendFields = append(update.ForceSendFields, "ReadOnly")
				}
				if update.DatabricksGcpServiceAccount != nil { // we can't update it at all
					update.DatabricksGcpServiceAccount = nil
				}
				update.Owner = ""
				_, err := acc.StorageCredentials.Update(ctx, catalog.AccountsUpdateStorageCredential{
					CredentialInfo:        toUpdateAccountsStorageCredential(&update),
					MetastoreId:           metastoreId,
					StorageCredentialName: storageCredentialName,
				})
				if err != nil {
					if d.HasChange("owner") {
						// Rollback
						old, new := d.GetChange("owner")
						rollbackUpdate := catalog.UpdateStorageCredential{
							Name:  update.Name,
							Owner: old.(string),
						}
						_, rollbackErr := acc.StorageCredentials.Update(ctx, catalog.AccountsUpdateStorageCredential{
							CredentialInfo:        toUpdateAccountsStorageCredential(&rollbackUpdate),
							MetastoreId:           metastoreId,
							StorageCredentialName: storageCredentialName,
						})
						if rollbackErr != nil {
							return common.OwnerRollbackError(err, rollbackErr, old.(string), new.(string))
						}
					}
					return err
				}
				return nil
			}, func(w *databricks.WorkspaceClient) error {
				err := validateMetastoreId(ctx, w, metastoreId)
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

				if d.HasChange("read_only") {
					update.ForceSendFields = append(update.ForceSendFields, "ReadOnly")
				}
				if update.DatabricksGcpServiceAccount != nil { // we can't update it at all
					update.DatabricksGcpServiceAccount = nil
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
				return bindings.AddCurrentWorkspaceBindings(ctx, d, w, update.Name, bindings.BindingsSecurableTypeStorageCredential)
			})
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			// Parse the ID to handle both composite and simple formats
			metastoreId, storageCredentialName, err := parseStorageCredentialId(d)
			if err != nil {
				return err
			}

			force := d.Get("force_destroy").(bool)
			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				_, err := acc.StorageCredentials.Delete(ctx, catalog.DeleteAccountStorageCredentialRequest{
					Force:                 force,
					StorageCredentialName: storageCredentialName,
					MetastoreId:           metastoreId,
				})
				return err
			}, func(w *databricks.WorkspaceClient) error {
				err := validateMetastoreId(ctx, w, metastoreId)
				if err != nil {
					return err
				}
				return w.StorageCredentials.Delete(ctx, catalog.DeleteStorageCredentialRequest{
					Force: force,
					Name:  storageCredentialName,
				})
			})
		},
	}
}
