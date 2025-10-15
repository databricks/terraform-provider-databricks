package sharing

import (
	"context"
	"reflect"
	"sort"

	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ShareInfo struct {
	sharing.ShareInfo
	common.Namespace
}

func (ShareInfo) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	common.NamespaceCustomizeSchema(s)
	s.SchemaPath("name").SetRequired()
	s.SchemaPath("name").SetForceNew()
	s.SchemaPath("name").SetCustomSuppressDiff(common.EqualFoldDiffSuppress)
	s.SchemaPath("comment").SetSuppressDiff()
	s.SchemaPath("owner").SetSuppressDiff()
	s.SchemaPath("created_at").SetComputed()
	s.SchemaPath("created_by").SetComputed()
	s.SchemaPath("updated_at").SetComputed()
	s.SchemaPath("updated_by").SetComputed()

	s.SchemaPath("object").SetMinItems(1)
	s.SchemaPath("object", "data_object_type").SetRequired()
	s.SchemaPath("object", "shared_as").SetSuppressDiff()
	s.SchemaPath("object", "string_shared_as").SetSuppressDiff()
	s.SchemaPath("object", "cdf_enabled").SetSuppressDiff()
	s.SchemaPath("object", "start_version").SetSuppressDiff()
	s.SchemaPath("object", "history_data_sharing_status").SetSuppressDiff()
	s.SchemaPath("object", "status").SetComputed()
	s.SchemaPath("object", "added_at").SetComputed()
	s.SchemaPath("object", "added_by").SetComputed()
	s.SchemaPath("object", "partition", "value", "op").SetRequired()
	s.SchemaPath("object", "partition", "value", "name").SetRequired()

	return s
}

func (ShareInfo) Aliases() map[string]map[string]string {
	return map[string]map[string]string{
		"sharing.ShareInfo": {
			"objects": "object",
		},
		"sharing.SharedDataObject": {
			"partitions": "partition",
		},
		"sharing.Partition": {
			"values": "value",
		},
	}
}

type Shares struct {
	Shares []ShareInfo `json:"shares"`
}

func (si *ShareInfo) sortSharesByName() {
	sort.Slice(si.Objects, func(i, j int) bool {
		return si.Objects[i].Name < si.Objects[j].Name
	})
}

func (si *ShareInfo) suppressCDFEnabledDiff() {
	// suppress diff for CDF Enabled if HistoryDataSharingStatus is enabled , as API does not accept both fields to be set
	for i := range si.Objects {
		if si.Objects[i].HistoryDataSharingStatus == "ENABLED" {
			si.Objects[i].CdfEnabled = false
		}
	}
}

func (si ShareInfo) shareChanges(action string) sharing.UpdateShare {
	var changes []sharing.SharedDataObjectUpdate
	for _, obj := range si.Objects {
		changes = append(changes, sharing.SharedDataObjectUpdate{
			Action:     sharing.SharedDataObjectUpdateAction(action),
			DataObject: &obj,
		},
		)
	}
	return sharing.UpdateShare{
		Updates: changes,
	}
}

func (si ShareInfo) resourceShareMap() map[string]sharing.SharedDataObject {
	m := make(map[string]sharing.SharedDataObject, len(si.Objects))
	for _, sdo := range si.Objects {
		m[sdo.Name] = sdo
	}
	return m
}

func Equal(this sharing.SharedDataObject, other sharing.SharedDataObject) bool {
	if other.SharedAs == "" {
		other.SharedAs = this.SharedAs
	}
	if other.StringSharedAs == "" {
		other.StringSharedAs = this.StringSharedAs
	}
	// don't compare computed fields
	other.AddedAt = this.AddedAt
	other.AddedBy = this.AddedBy
	other.Status = this.Status
	other.ForceSendFields = this.ForceSendFields // TODO: is this the right thing to do?
	return reflect.DeepEqual(this, other)
}

func (beforeSi ShareInfo) Diff(afterSi ShareInfo) []sharing.SharedDataObjectUpdate {
	beforeMap := beforeSi.resourceShareMap()
	afterMap := afterSi.resourceShareMap()
	changes := []sharing.SharedDataObjectUpdate{}
	// not in after so remove
	for _, beforeSdo := range beforeSi.Objects {
		_, exists := afterMap[beforeSdo.Name]
		if exists {
			continue
		}
		changes = append(changes, sharing.SharedDataObjectUpdate{
			Action:     sharing.SharedDataObjectUpdateActionRemove,
			DataObject: &beforeSdo,
		})
	}

	// not in before so add
	// if in before but diff then update
	for _, afterSdo := range afterSi.Objects {
		beforeSdo, exists := beforeMap[afterSdo.Name]
		if exists {
			if !Equal(beforeSdo, afterSdo) {
				// do not send SharedAs
				afterSdo.SharedAs = ""
				changes = append(changes, sharing.SharedDataObjectUpdate{
					Action:     sharing.SharedDataObjectUpdateActionUpdate,
					DataObject: &afterSdo,
				})
			}
			continue
		}
		changes = append(changes, sharing.SharedDataObjectUpdate{
			Action:     sharing.SharedDataObjectUpdateActionAdd,
			DataObject: &afterSdo,
		})
	}
	return changes
}

func ResourceShare() common.Resource {
	shareSchema := common.StructToSchema(ShareInfo{}, nil)
	return common.Resource{
		Schema: shareSchema,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			return common.NamespaceCustomizeDiff(d)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}

			var createRequest sharing.CreateShare
			common.DataToStructPointer(d, shareSchema, &createRequest)
			if _, err := w.Shares.Create(ctx, createRequest); err != nil {
				return err
			}

			// can only create empty share, objects & owners have to be added using update API
			var si ShareInfo
			common.DataToStructPointer(d, shareSchema, &si)
			shareChanges := si.shareChanges(string(sharing.SharedDataObjectUpdateActionAdd))
			shareChanges.Name = si.Name
			shareChanges.Comment = si.Comment
			shareChanges.Owner = si.Owner
			if _, err := w.Shares.Update(ctx, shareChanges); err != nil {
				// delete orphaned share if update fails
				if d_err := w.Shares.DeleteByName(ctx, si.Name); d_err != nil {
					return d_err
				}
				return err
			}
			d.SetId(si.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			client, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}

			shareInfo, err := client.Shares.Get(ctx, sharing.GetShareRequest{
				Name:              d.Id(),
				IncludeSharedData: true,
			})
			si := ShareInfo{ShareInfo: *shareInfo}
			si.sortSharesByName()
			si.suppressCDFEnabledDiff()
			if err != nil {
				return err
			}

			return common.StructToData(si, shareSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			client, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}

			si, err := client.Shares.Get(ctx, sharing.GetShareRequest{
				Name:              d.Id(),
				IncludeSharedData: true,
			})
			if err != nil {
				return err
			}

			beforeSi := ShareInfo{ShareInfo: *si}
			beforeSi.sortSharesByName()
			beforeSi.suppressCDFEnabledDiff()
			var afterSi ShareInfo
			common.DataToStructPointer(d, shareSchema, &afterSi)
			changes := beforeSi.Diff(afterSi)

			if d.HasChange("owner") {
				_, err = client.Shares.Update(ctx, sharing.UpdateShare{
					Name:  afterSi.Name,
					Owner: afterSi.Owner,
				})
				if err != nil {
					return err
				}
			}

			if !d.HasChangeExcept("owner") {
				return nil
			}

			if len(changes) == 0 {
				return nil
			}

			_, err = client.Shares.Update(ctx, sharing.UpdateShare{
				Name:    d.Id(),
				Comment: afterSi.Comment,
				Updates: changes,
			})
			if err != nil {
				if d.HasChange("owner") {
					// Rollback
					old, new := d.GetChange("owner")
					_, rollbackErr := client.Shares.Update(ctx, sharing.UpdateShare{
						Name:  beforeSi.Name,
						Owner: old.(string),
					})
					if rollbackErr != nil {
						return common.OwnerRollbackError(err, rollbackErr, old.(string), new.(string))
					}
				}
				return err
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return w.Shares.DeleteByName(ctx, d.Id())
		},
	}
}
