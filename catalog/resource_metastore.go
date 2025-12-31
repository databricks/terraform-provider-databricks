package catalog

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"golang.org/x/exp/slices"
)

const maxDeltaSharingRecipientTokenLifetimeInSeconds = int64(365 * 24 * time.Hour / time.Second) // 1 year

// This and the next function should be updated together to keep them in sync.
func updateForceSendFieldsWorkspaceLevel(req *catalog.UpdateMetastore) {
	if req.DeltaSharingScope != "" && !slices.Contains(req.ForceSendFields, "DeltaSharingRecipientTokenLifetimeInSeconds") {
		req.ForceSendFields = append(req.ForceSendFields, "DeltaSharingRecipientTokenLifetimeInSeconds")
	}
}

// This and the previous function should be updated together to keep them in sync.
func updateForceSendFieldsAccountLevel(req *catalog.UpdateAccountsMetastore) {
	if req.DeltaSharingScope != "" && !slices.Contains(req.ForceSendFields, "DeltaSharingRecipientTokenLifetimeInSeconds") {
		req.ForceSendFields = append(req.ForceSendFields, "DeltaSharingRecipientTokenLifetimeInSeconds")
	}
}

// Cannot set lifetime to 0 (unlimited), so set to 1 year (maximum allowed)
func setDefaultDeltaSharingRecipientTokenLifetimeInSecondsWorkspaceLevel(req *catalog.UpdateMetastore) {
	if req.DeltaSharingRecipientTokenLifetimeInSeconds == 0 {
		log.Printf("[DEBUG] Setting delta sharing recipient token lifetime to 1 year")
		req.DeltaSharingRecipientTokenLifetimeInSeconds = maxDeltaSharingRecipientTokenLifetimeInSeconds
	}
}

func setDefaultDeltaSharingRecipientTokenLifetimeInSecondsAccountLevel(req *catalog.UpdateAccountsMetastore) {
	if req.DeltaSharingRecipientTokenLifetimeInSeconds == 0 {
		log.Printf("[DEBUG] Setting delta sharing recipient token lifetime to 1 year")
		req.DeltaSharingRecipientTokenLifetimeInSeconds = maxDeltaSharingRecipientTokenLifetimeInSeconds
	}
}

func ResourceMetastore() common.Resource {
	s := common.StructToSchema(catalog.MetastoreInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			// Add custom field
			m["force_destroy"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			}

			// Mark fields as force_new
			common.CustomizeSchemaPath(m, "storage_root").SetForceNew()

			// Mark optional and computed fields (user can set, but has default from backend)
			for _, v := range []string{"owner", "privilege_model_version", "region"} {
				common.CustomizeSchemaPath(m, v).SetComputed()
			}

			// Set read-only fields (returned by backend, not settable by user)
			// Note: SetReadOnly() implies SetComputed(), so no need to call both
			for _, v := range []string{"metastore_id", "cloud", "global_metastore_id",
				"created_at", "created_by", "updated_at", "updated_by"} {
				common.CustomizeSchemaPath(m, v).SetReadOnly()
			}

			// Custom diff suppressions
			common.CustomizeSchemaPath(m, "default_data_access_config_id").SetSuppressDiff()
			common.CustomizeSchemaPath(m, "storage_root_credential_id").SetSuppressDiff()
			common.CustomizeSchemaPath(m, "delta_sharing_scope").SetSuppressDiff()

			common.CustomizeSchemaPath(m, "name").SetCustomSuppressDiff(common.EqualFoldDiffSuppress)

			// Custom storage_root diff suppression
			m["storage_root"].DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
				if strings.HasPrefix(old, new) {
					log.Printf("[DEBUG] Ignoring configuration drift from %s to %s", old, new)
					return true
				}
				return false
			}

			// Field dependencies and validation
			m["delta_sharing_scope"].RequiredWith = []string{"delta_sharing_recipient_token_lifetime_in_seconds"}
			m["delta_sharing_recipient_token_lifetime_in_seconds"].RequiredWith = []string{"delta_sharing_scope"}
			common.CustomizeSchemaPath(m, "delta_sharing_scope").SetValidateFunc(
				validation.StringInSlice([]string{"INTERNAL", "INTERNAL_AND_EXTERNAL"}, false),
			)

			return m
		})

	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				var create catalog.CreateAccountsMetastore
				var update catalog.UpdateAccountsMetastore
				common.DataToStructPointer(d, s, &create)
				common.DataToStructPointer(d, s, &update)
				updateForceSendFieldsAccountLevel(&update)
				emptyRequest, err := common.IsRequestEmpty(update)
				if err != nil {
					return err
				}
				mi, err := acc.Metastores.Create(ctx,
					catalog.AccountsCreateMetastore{
						MetastoreInfo: &create,
					})
				if err != nil {
					return err
				}
				d.SetId(mi.MetastoreInfo.MetastoreId)
				if emptyRequest {
					return nil
				}
				_, err = acc.Metastores.Update(ctx, catalog.AccountsUpdateMetastore{
					MetastoreId:   mi.MetastoreInfo.MetastoreId,
					MetastoreInfo: &update,
				})
				if err != nil {
					return err
				}
				return nil
			}, func(w *databricks.WorkspaceClient) error {
				var create catalog.CreateMetastore
				var update catalog.UpdateMetastore
				common.DataToStructPointer(d, s, &create)
				common.DataToStructPointer(d, s, &update)
				updateForceSendFieldsWorkspaceLevel(&update)
				emptyRequest, err := common.IsRequestEmpty(update)
				if err != nil {
					return err
				}
				mi, err := w.Metastores.Create(ctx, create)
				if err != nil {
					return err
				}
				d.SetId(mi.MetastoreId)
				if emptyRequest {
					return nil
				}
				update.Id = mi.MetastoreId
				_, err = w.Metastores.Update(ctx, update)
				if err != nil {
					return err
				}
				return nil
			})
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				mi, err := acc.Metastores.GetByMetastoreId(ctx, d.Id())
				if err != nil {
					return err
				}
				return common.StructToData(mi.MetastoreInfo, s, d)
			}, func(w *databricks.WorkspaceClient) error {
				mi, err := w.Metastores.GetById(ctx, d.Id())
				if err != nil {
					return err
				}
				return common.StructToData(mi, s, d)
			})
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {

			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				var update catalog.UpdateAccountsMetastore
				common.DataToStructPointer(d, s, &update)
				updateForceSendFieldsAccountLevel(&update)
				setDefaultDeltaSharingRecipientTokenLifetimeInSecondsAccountLevel(&update)
				if d.HasChange("owner") {
					ownerUpdate := catalog.UpdateAccountsMetastore{
						Owner: update.Owner,
					}
					_, err := acc.Metastores.Update(ctx, catalog.AccountsUpdateMetastore{
						MetastoreId:   d.Id(),
						MetastoreInfo: &ownerUpdate,
					})
					if err != nil {
						return err
					}
				}

				if !d.HasChangeExcept("owner") {
					return nil
				}

				update.Owner = ""
				_, err := acc.Metastores.Update(ctx, catalog.AccountsUpdateMetastore{
					MetastoreId:   d.Id(),
					MetastoreInfo: &update,
				})
				if err != nil {
					if d.HasChange("owner") {
						// Rollback
						old, new := d.GetChange("owner")
						rollbackUpdate := catalog.UpdateAccountsMetastore{
							Owner: old.(string),
						}
						_, rollbackErr := acc.Metastores.Update(ctx, catalog.AccountsUpdateMetastore{
							MetastoreId:   d.Id(),
							MetastoreInfo: &rollbackUpdate,
						})
						if rollbackErr != nil {
							return common.OwnerRollbackError(err, rollbackErr, old.(string), new.(string))
						}
					}
					return err
				}
				return nil
			}, func(w *databricks.WorkspaceClient) error {
				var update catalog.UpdateMetastore
				common.DataToStructPointer(d, s, &update)
				update.Id = d.Id()
				updateForceSendFieldsWorkspaceLevel(&update)
				setDefaultDeltaSharingRecipientTokenLifetimeInSecondsWorkspaceLevel(&update)
				if d.HasChange("owner") {
					_, err := w.Metastores.Update(ctx, catalog.UpdateMetastore{
						Id:    update.Id,
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
				_, err := w.Metastores.Update(ctx, update)
				if err != nil {
					if d.HasChange("owner") {
						// Rollback
						old, new := d.GetChange("owner")
						_, rollbackErr := w.Metastores.Update(ctx, catalog.UpdateMetastore{
							Id:    update.Id,
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
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			force := d.Get("force_destroy").(bool)
			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				_, err := acc.Metastores.Delete(ctx, catalog.DeleteAccountMetastoreRequest{Force: force, MetastoreId: d.Id()})
				return err
			}, func(w *databricks.WorkspaceClient) error {
				return w.Metastores.Delete(ctx, catalog.DeleteMetastoreRequest{Force: force, Id: d.Id()})
			})
		},
	}
}
