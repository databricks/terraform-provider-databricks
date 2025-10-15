package catalog

import (
	"context"
	"log"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"golang.org/x/exp/slices"
)

type MetastoreInfo struct {
	common.Namespace
	Name                                        string `json:"name"`
	StorageRoot                                 string `json:"storage_root,omitempty" tf:"force_new"`
	DefaultDacID                                string `json:"default_data_access_config_id,omitempty" tf:"suppress_diff"`
	StorageRootCredentialId                     string `json:"storage_root_credential_id,omitempty" tf:"suppress_diff"`
	Owner                                       string `json:"owner,omitempty" tf:"computed"`
	MetastoreID                                 string `json:"metastore_id,omitempty" tf:"computed"`
	Region                                      string `json:"region,omitempty" tf:"computed"`
	Cloud                                       string `json:"cloud,omitempty" tf:"computed"`
	GlobalMetastoreId                           string `json:"global_metastore_id,omitempty" tf:"computed"`
	CreatedAt                                   int64  `json:"created_at,omitempty" tf:"computed"`
	CreatedBy                                   string `json:"created_by,omitempty" tf:"computed"`
	UpdatedAt                                   int64  `json:"updated_at,omitempty" tf:"computed"`
	UpdatedBy                                   string `json:"updated_by,omitempty" tf:"computed"`
	DeltaSharingScope                           string `json:"delta_sharing_scope,omitempty" tf:"suppress_diff"`
	DeltaSharingRecipientTokenLifetimeInSeconds int64  `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	DeltaSharingOrganizationName                string `json:"delta_sharing_organization_name,omitempty"`
}

func updateForceSendFields(req *catalog.UpdateMetastore) {
	if req.DeltaSharingScope != "" && !slices.Contains(req.ForceSendFields, "DeltaSharingRecipientTokenLifetimeInSeconds") {
		req.ForceSendFields = append(req.ForceSendFields, "DeltaSharingRecipientTokenLifetimeInSeconds")
	}
}

func ResourceMetastore() common.Resource {
	s := common.StructToSchema(MetastoreInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["force_destroy"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			}
			m["delta_sharing_scope"].RequiredWith = []string{"delta_sharing_recipient_token_lifetime_in_seconds"}
			m["delta_sharing_scope"].ValidateFunc = validation.StringInSlice([]string{"INTERNAL", "INTERNAL_AND_EXTERNAL"}, false)
			m["delta_sharing_recipient_token_lifetime_in_seconds"].RequiredWith = []string{"delta_sharing_scope"}
			m["storage_root"].DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
				if strings.HasPrefix(old, new) {
					log.Printf("[DEBUG] Ignoring configuration drift from %s to %s", old, new)
					return true
				}
				return false
			}
			m["name"].DiffSuppressFunc = common.EqualFoldDiffSuppress
			common.NamespaceCustomizeSchemaMap(m)
			return m
		})

	return common.Resource{
		Schema: s,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			return common.NamespaceCustomizeDiff(d)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var create catalog.CreateMetastore
			var update catalog.UpdateMetastore
			common.DataToStructPointer(d, s, &create)
			common.DataToStructPointer(d, s, &update)
			updateForceSendFields(&update)
			emptyRequest, err := common.IsRequestEmpty(update)
			if err != nil {
				return err
			}
			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
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
			var update catalog.UpdateMetastore
			common.DataToStructPointer(d, s, &update)
			update.Id = d.Id()
			updateForceSendFields(&update)

			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				if d.HasChange("owner") {
					_, err := acc.Metastores.Update(ctx, catalog.AccountsUpdateMetastore{
						MetastoreId: d.Id(),
						MetastoreInfo: &catalog.UpdateMetastore{
							Id:    update.Id,
							Owner: update.Owner,
						},
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
						_, rollbackErr := acc.Metastores.Update(ctx, catalog.AccountsUpdateMetastore{
							MetastoreId: d.Id(),
							MetastoreInfo: &catalog.UpdateMetastore{
								Id:    update.Id,
								Owner: old.(string),
							},
						})
						if rollbackErr != nil {
							return common.OwnerRollbackError(err, rollbackErr, old.(string), new.(string))
						}
					}
					return err
				}
				return nil
			}, func(w *databricks.WorkspaceClient) error {
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
				return acc.Metastores.Delete(ctx, catalog.DeleteAccountMetastoreRequest{Force: force, MetastoreId: d.Id()})
			}, func(w *databricks.WorkspaceClient) error {
				return w.Metastores.Delete(ctx, catalog.DeleteMetastoreRequest{Force: force, Id: d.Id()})
			})
		},
	}
}
