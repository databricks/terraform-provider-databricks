package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AwsIamRole struct {
	RoleARN string `json:"role_arn"`
}

type AzureServicePrincipal struct {
	DirectoryID   string `json:"directory_id"`
	ApplicationID string `json:"application_id"`
	ClientSecret  string `json:"client_secret" tf:"sensitive"`
}

type AzureManagedIdentity struct {
	AccessConnectorID string `json:"access_connector_id"`
}

type GcpServiceAccountKey struct {
	Email        string `json:"email"`
	PrivateKeyId string `json:"private_key_id"`
	PrivateKey   string `json:"private_key" tf:"sensitive"`
}

type DbGcpServiceAccount struct {
	Email string `json:"email,omitempty" tf:"computed"`
}

type DataAccessConfiguration struct {
	ID                string                 `json:"id,omitempty" tf:"computed"`
	Name              string                 `json:"name"`
	ConfigurationType string                 `json:"configuration_type,omitempty" tf:"computed"`
	Aws               *AwsIamRole            `json:"aws_iam_role,omitempty" tf:"group:access"`
	Azure             *AzureServicePrincipal `json:"azure_service_principal,omitempty" tf:"group:access"`
	AzMI              *AzureManagedIdentity  `json:"azure_managed_identity,omitempty" tf:"group:access"`
	GcpSAKey          *GcpServiceAccountKey  `json:"gcp_service_account_key,omitempty" tf:"group:access"`
	DBGcpSA           *DbGcpServiceAccount   `json:"databricks_gcp_service_account,omitempty" tf:"group:access"`
}

var alofCred = []string{"aws_iam_role", "azure_service_principal", "azure_managed_identity", "gcp_service_account_key", "databricks_gcp_service_account"}

func SuppressGcpSAKeyDiff(k, old, new string, d *schema.ResourceData) bool {
	//ignore changes in private_key
	return !d.HasChanges("gcp_service_account_key.0.email", "gcp_service_account_key.0.private_key_id")
}

func ResourceMetastoreDataAccess() *schema.Resource {
	s := common.StructToSchema(DataAccessConfiguration{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["metastore_id"] = &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			}
			m["is_default"] = &schema.Schema{
				// having more than one default DAC per metastore will lead
				// to Terraform re-assigning default_data_access_config_id
				// on every apply.
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
	p := common.NewPairID("metastore_id", "id")
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			metastoreId := d.Get("metastore_id").(string)

			tmpSchema := removeGcpSaField(s)
			var create catalog.CreateStorageCredential
			common.DataToStructPointer(d, tmpSchema, &create)

			return c.WorkspaceOrAccountRequest(func(acc *databricks.AccountClient) error {
				dac, err := acc.StorageCredentials.Create(ctx,
					catalog.AccountsCreateStorageCredential{
						MetastoreId:    metastoreId,
						CredentialInfo: &create,
					})
				if err != nil {
					return err
				}
				if d.Get("is_default").(bool) {
					_, err = acc.Metastores.Update(ctx, catalog.AccountsUpdateMetastore{
						MetastoreId: metastoreId,
						MetastoreInfo: &catalog.UpdateMetastore{
							StorageRootCredentialId: dac.Id,
						},
					})
					if err != nil {
						return err
					}
				}
				d.Set("id", dac.Id)
				p.Pack(d)
				return nil
			}, func(w *databricks.WorkspaceClient) error {
				dac, err := w.StorageCredentials.Create(ctx, create)
				if err != nil {
					return err
				}
				if d.Get("is_default").(bool) {
					_, err = w.Metastores.Update(ctx, catalog.UpdateMetastore{
						Id:                      metastoreId,
						StorageRootCredentialId: dac.Id,
					})
				}
				if err != nil {
					return err
				}
				d.Set("id", dac.Id)
				p.Pack(d)
				return nil
			})
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			metastoreId, dacId, err := p.Unpack(d)
			if err != nil {
				return err
			}
			var storageCredential *catalog.StorageCredentialInfo
			var metastore *catalog.MetastoreInfo

			return c.WorkspaceOrAccountRequest(func(acc *databricks.AccountClient) error {
				storageCredential, err = acc.StorageCredentials.Get(ctx, catalog.GetAccountStorageCredentialRequest{
					MetastoreId: metastoreId,
					Name:        dacId,
				})
				if err != nil {
					return err
				}
				m, err := acc.Metastores.GetByMetastoreId(ctx, metastoreId)
				metastore = m.MetastoreInfo
				if err != nil {
					return err
				}
				isDefault := metastore.StorageRootCredentialId == dacId
				d.Set("is_default", isDefault)
				return common.StructToData(storageCredential, s, d)
			}, func(w *databricks.WorkspaceClient) error {
				storageCredential, err = w.StorageCredentials.GetByName(ctx, dacId)
				if err != nil {
					return err
				}
				metastore, err = w.Metastores.GetById(ctx, metastoreId)
				if err != nil {
					return err
				}
				isDefault := metastore.StorageRootCredentialId == dacId
				d.Set("is_default", isDefault)
				return common.StructToData(storageCredential, s, d)
			})
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			metastoreId, dacId, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return c.WorkspaceOrAccountRequest(func(acc *databricks.AccountClient) error {
				return acc.StorageCredentials.Delete(ctx, catalog.DeleteAccountStorageCredentialRequest{
					MetastoreId: metastoreId,
					Name:        dacId,
				})
			}, func(w *databricks.WorkspaceClient) error {
				return w.StorageCredentials.DeleteByName(ctx, dacId)
			})
		},
	}.ToResource()
}
