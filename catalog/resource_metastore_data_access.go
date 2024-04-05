package catalog

import (
	"context"
	"fmt"
	"log"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var dacSchema = common.StructToSchema(StorageCredentialInfo{},
	func(m map[string]*schema.Schema) map[string]*schema.Schema {
		m["is_default"] = &schema.Schema{
			// having more than one default DAC per metastore will lead
			// to Terraform re-assigning default_data_access_config_id
			// on every apply.
			Type:     schema.TypeBool,
			Optional: true,
			ForceNew: true,
		}

		return adjustDataAccessSchema(m)
	})

func ResourceMetastoreDataAccess() common.Resource {
	p := common.NewPairID("metastore_id", "name")
	return common.Resource{
		Schema:        dacSchema,
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Version: 0,
				Type:    dacSchemaV0(),
				Upgrade: dacMigrateV0,
			},
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			metastoreId := d.Get("metastore_id").(string)

			tmpSchema := removeGcpSaField(dacSchema)
			var create catalog.CreateStorageCredential
			common.DataToStructPointer(d, tmpSchema, &create)
			//manually add empty struct back for databricks_gcp_service_account
			if _, ok := d.GetOk("databricks_gcp_service_account"); ok {
				create.DatabricksGcpServiceAccount = &catalog.DatabricksGcpServiceAccountRequest{}
			}

			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
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
							StorageRootCredentialId: dac.CredentialInfo.Id,
						},
					})
					if err != nil {
						// Rollback
						rollback_err := acc.StorageCredentials.DeleteByMetastoreIdAndStorageCredentialName(ctx, metastoreId, dac.CredentialInfo.Name)
						if rollback_err != nil {
							return fmt.Errorf("failed to rollback: %v", rollback_err)
						}
						return err
					}
				}
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
					if err != nil {
						// Rollback
						rollback_err := w.StorageCredentials.DeleteByName(ctx, dac.Name)
						if rollback_err != nil {
							return fmt.Errorf("failed to rollback: %v", rollback_err)
						}
						return err
					}
				}
				p.Pack(d)
				return nil
			})
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			metastoreId, dacName, err := p.Unpack(d)
			if err != nil {
				return err
			}
			var metastore *catalog.MetastoreInfo

			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				var storageCredential *catalog.AccountsStorageCredentialInfo
				storageCredential, err = acc.StorageCredentials.GetByMetastoreIdAndStorageCredentialName(ctx, metastoreId, dacName)
				if err != nil {
					return err
				}
				m, err := acc.Metastores.GetByMetastoreId(ctx, metastoreId)
				metastore = m.MetastoreInfo
				if err != nil {
					return err
				}
				isDefault := metastore.StorageRootCredentialName == dacName
				d.Set("is_default", isDefault)
				return common.StructToData(storageCredential.CredentialInfo, dacSchema, d)
			}, func(w *databricks.WorkspaceClient) error {
				var storageCredential *catalog.StorageCredentialInfo
				storageCredential, err = w.StorageCredentials.GetByName(ctx, dacName)
				if err != nil {
					return err
				}
				m, err := w.Metastores.Summary(ctx)
				if err != nil {
					return err
				}
				isDefault := m.StorageRootCredentialName == dacName
				d.Set("is_default", isDefault)
				return common.StructToData(storageCredential, dacSchema, d)
			})
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var update catalog.UpdateStorageCredential
			force := d.Get("force_update").(bool)
			metastoreId, dacName, err := p.Unpack(d)
			if err != nil {
				return err
			}
			common.DataToStructPointer(d, storageCredentialSchema, &update)
			update.Name = dacName
			update.Force = force
			if _, ok := d.GetOk("azure_managed_identity"); ok {
				update.AzureManagedIdentity.CredentialId = ""
			}
			return updateStorageCredential(ctx, d, c, metastoreId, dacName, update)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			metastoreId, dacName, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return deleteStorageCredential(ctx, c, metastoreId, dacName, d.Get("force_destroy").(bool))
		},
	}
}

// migrate to v1 state, as the id is now changed
func dacMigrateV0(ctx context.Context,
	rawState map[string]any,
	meta any) (map[string]any, error) {
	newState := map[string]any{}
	for k, v := range rawState {
		switch k {
		case "id":
			log.Printf("[INFO] Upgrade id")
			newState[k] = fmt.Sprintf("%v%s%v", rawState["metastore_id"], "|", rawState["name"])
			continue
		default:
			newState[k] = v
		}
	}
	return newState, nil
}

func dacSchemaV0() cty.Type {
	return (&schema.Resource{
		Schema: dacSchema}).CoreConfigSchema().ImpliedType()
}
