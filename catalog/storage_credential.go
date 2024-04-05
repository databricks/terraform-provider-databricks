package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type StorageCredentialInfo struct {
	Name                        string                                       `json:"name" tf:"force_new"`
	Owner                       string                                       `json:"owner,omitempty" tf:"computed"`
	Comment                     string                                       `json:"comment,omitempty"`
	Aws                         *catalog.AwsIamRoleResponse                  `json:"aws_iam_role,omitempty" tf:"group:access"`
	Azure                       *catalog.AzureServicePrincipal               `json:"azure_service_principal,omitempty" tf:"group:access"`
	AzMI                        *catalog.AzureManagedIdentity                `json:"azure_managed_identity,omitempty" tf:"group:access"`
	GcpSAKey                    *GcpServiceAccountKey                        `json:"gcp_service_account_key,omitempty" tf:"group:access"`
	DatabricksGcpServiceAccount *catalog.DatabricksGcpServiceAccountResponse `json:"databricks_gcp_service_account,omitempty" tf:"computed"`
	MetastoreID                 string                                       `json:"metastore_id,omitempty" tf:"computed"`
	ReadOnly                    bool                                         `json:"read_only,omitempty"`
	SkipValidation              bool                                         `json:"skip_validation,omitempty"`
}

type GcpServiceAccountKey struct {
	Email        string `json:"email"`
	PrivateKeyId string `json:"private_key_id"`
	PrivateKey   string `json:"private_key" tf:"sensitive"`
}

var alofCred = []string{"aws_iam_role", "azure_service_principal", "azure_managed_identity",
	"gcp_service_account_key", "databricks_gcp_service_account"}

func SuppressGcpSAKeyDiff(k, old, new string, d *schema.ResourceData) bool {
	//ignore changes in private_key
	return !d.HasChanges("gcp_service_account_key.0.email", "gcp_service_account_key.0.private_key_id")
}

// it's used by both ResourceMetastoreDataAccess & ResourceStorageCredential
func adjustDataAccessSchema(m map[string]*schema.Schema) map[string]*schema.Schema {
	m["aws_iam_role"].AtLeastOneOf = alofCred
	m["azure_service_principal"].AtLeastOneOf = alofCred
	m["azure_managed_identity"].AtLeastOneOf = alofCred
	m["gcp_service_account_key"].AtLeastOneOf = alofCred
	m["databricks_gcp_service_account"].AtLeastOneOf = alofCred

	// suppress changes for private_key
	m["gcp_service_account_key"].DiffSuppressFunc = SuppressGcpSAKeyDiff

	common.MustSchemaPath(m, "aws_iam_role", "external_id").Computed = true
	common.MustSchemaPath(m, "aws_iam_role", "unity_catalog_iam_arn").Computed = true
	common.MustSchemaPath(m, "azure_managed_identity", "credential_id").Computed = true
	common.MustSchemaPath(m, "databricks_gcp_service_account", "email").Computed = true
	common.MustSchemaPath(m, "databricks_gcp_service_account", "credential_id").Computed = true

	m["force_destroy"] = &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
	}

	m["force_update"] = &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
	}

	m["skip_validation"].DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
		return old == "false" && new == "true"
	}

	m["name"].DiffSuppressFunc = common.EqualFoldDiffSuppress

	return m
}

func updateStorageCredential(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient, metastoreId string, name string, update catalog.UpdateStorageCredential) error {
	return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
		if d.HasChange("owner") {
			_, err := acc.StorageCredentials.Update(ctx, catalog.AccountsUpdateStorageCredential{
				CredentialInfo: &catalog.UpdateStorageCredential{
					Name:  update.Name,
					Owner: update.Owner,
				},
				MetastoreId:           metastoreId,
				StorageCredentialName: name,
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
			MetastoreId:           metastoreId,
			StorageCredentialName: name,
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
					MetastoreId:           metastoreId,
					StorageCredentialName: name,
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
		return nil
	})
}

func deleteStorageCredential(ctx context.Context, c *common.DatabricksClient, metastoreId string, name string, force bool) error {
	return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
		return acc.StorageCredentials.Delete(ctx, catalog.DeleteAccountStorageCredentialRequest{
			MetastoreId:           metastoreId,
			StorageCredentialName: name,
			Force:                 force,
		})
	}, func(w *databricks.WorkspaceClient) error {
		err := validateMetastoreId(ctx, w, metastoreId)
		if err != nil {
			return err
		}
		return w.StorageCredentials.Delete(ctx, catalog.DeleteStorageCredentialRequest{
			Name:  name,
			Force: force,
		})
	})
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
