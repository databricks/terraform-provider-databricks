package sharing

import (
	"context"

	"reflect"
	"sort"

	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SharesAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewSharesAPI(ctx context.Context, m any) SharesAPI {
	return SharesAPI{m.(*common.DatabricksClient), context.WithValue(ctx, common.Api, common.API_2_1)}
}

const (
	ShareAdd    = "ADD"
	ShareRemove = "REMOVE"
	ShareUpdate = "UPDATE"
)

type ShareInfo struct {
	Name      string             `json:"name" tf:"force_new"`
	Owner     string             `json:"owner,omitempty" tf:"suppress_diff"`
	Objects   []SharedDataObject `json:"objects,omitempty" tf:"alias:object"`
	CreatedAt int64              `json:"created_at,omitempty" tf:"computed"`
	CreatedBy string             `json:"created_by,omitempty" tf:"computed"`
}

type SharedDataObject struct {
	Name                     string      `json:"name"`
	DataObjectType           string      `json:"data_object_type"`
	Comment                  string      `json:"comment,omitempty"`
	SharedAs                 string      `json:"shared_as,omitempty" tf:"suppress_diff"`
	CDFEnabled               bool        `json:"cdf_enabled,omitempty" tf:"suppress_diff"`
	StartVersion             int64       `json:"start_version,omitempty" tf:"suppress_diff"`
	HistoryDataSharingStatus string      `json:"history_data_sharing_status,omitempty" tf:"suppress_diff"`
	Partitions               []Partition `json:"partitions,omitempty" tf:"alias:partition"`
	Status                   string      `json:"status,omitempty" tf:"computed"`
	AddedAt                  int64       `json:"added_at,omitempty" tf:"computed"`
	AddedBy                  string      `json:"added_by,omitempty" tf:"computed"`
}

type ShareDataChange struct {
	Action     string           `json:"action"`
	DataObject SharedDataObject `json:"data_object"`
}

type ShareUpdates struct {
	Owner   string            `json:"owner,omitempty"`
	Updates []ShareDataChange `json:"updates"`
}

func (su *ShareUpdates) sortSharesByName() {
	sort.Slice(su.Updates, func(i, j int) bool {
		return su.Updates[i].DataObject.Name < su.Updates[j].DataObject.Name
	})
}

type Shares struct {
	Shares []ShareInfo `json:"shares"`
}

type Partition struct {
	Values []PartitionValue `json:"values" tf:"alias:value"`
}

type PartitionValue struct {
	Name                 string `json:"name"`
	Op                   string `json:"op"`
	RecipientPropertyKey string `json:"recipient_property_key,omitempty"`
	Value                string `json:"value,omitempty"`
}

func (si *ShareInfo) sortSharesByName() {
	sort.Slice(si.Objects, func(i, j int) bool {
		return si.Objects[i].Name < si.Objects[j].Name
	})
}

func (si *ShareInfo) suppressCDFEnabledDiff() {
	//suppress diff for CDF Enabled if HistoryDataSharingStatus is enabled , as API does not accept both fields to be set
	for i := range si.Objects {
		if si.Objects[i].HistoryDataSharingStatus == "ENABLED" {
			si.Objects[i].CDFEnabled = false
		}
	}
}

func (a SharesAPI) get(name string) (si ShareInfo, err error) {
	err = a.client.Get(a.context, "/unity-catalog/shares/"+name+"?include_shared_data=true", nil, &si)
	si.sortSharesByName()
	si.suppressCDFEnabledDiff()
	return
}

func (a SharesAPI) update(name string, su ShareUpdates) error {
	if len(su.Updates) == 0 {
		return nil
	}
	su.sortSharesByName()
	err := a.client.Patch(a.context, "/unity-catalog/shares/"+name, su)
	return err
}

func (si ShareInfo) shareChanges(action string) ShareUpdates {
	var changes []ShareDataChange
	for _, obj := range si.Objects {
		changes = append(changes, ShareDataChange{
			Action:     action,
			DataObject: obj,
		})
	}
	return ShareUpdates{
		Updates: changes,
	}
}

func (si ShareInfo) resourceShareMap() map[string]SharedDataObject {
	m := make(map[string]SharedDataObject, len(si.Objects))
	for _, sdo := range si.Objects {
		m[sdo.Name] = sdo
	}
	return m
}

func (sdo SharedDataObject) Equal(other SharedDataObject) bool {
	if other.SharedAs == "" {
		other.SharedAs = sdo.SharedAs
	}
	//don't compare computed fields
	other.AddedAt = sdo.AddedAt
	other.AddedBy = sdo.AddedBy
	other.Status = sdo.Status
	return reflect.DeepEqual(sdo, other)
}

func (beforeSi ShareInfo) Diff(afterSi ShareInfo) []ShareDataChange {
	beforeMap := beforeSi.resourceShareMap()
	afterMap := afterSi.resourceShareMap()
	changes := []ShareDataChange{}
	// not in after so remove
	for _, beforeSdo := range beforeSi.Objects {
		_, exists := afterMap[beforeSdo.Name]
		if exists {
			continue
		}
		changes = append(changes, ShareDataChange{
			Action:     ShareRemove,
			DataObject: beforeSdo,
		})
	}

	// not in before so add
	// if in before but diff then update
	for _, afterSdo := range afterSi.Objects {
		beforeSdo, exists := beforeMap[afterSdo.Name]
		if exists {
			if !beforeSdo.Equal(afterSdo) {
				// do not send SharedAs
				afterSdo.SharedAs = ""
				changes = append(changes, ShareDataChange{
					Action:     ShareUpdate,
					DataObject: afterSdo,
				})
			}
			continue
		}
		changes = append(changes, ShareDataChange{
			Action:     ShareAdd,
			DataObject: afterSdo,
		})
	}
	return changes
}

func ResourceShare() common.Resource {
	shareSchema := common.StructToSchema(ShareInfo{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		m["name"].DiffSuppressFunc = common.EqualFoldDiffSuppress
		return m
	})
	return common.Resource{
		Schema: shareSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}

			var createRequest sharing.CreateShare
			common.DataToStructPointer(d, shareSchema, &createRequest)
			if _, err := w.Shares.Create(ctx, createRequest); err != nil {
				return err
			}

			//can only create empty share, objects & owners have to be added using update API
			var si ShareInfo
			common.DataToStructPointer(d, shareSchema, &si)
			shareChanges := si.shareChanges(ShareAdd)
			shareChanges.Owner = si.Owner
			if err := NewSharesAPI(ctx, c).update(si.Name, shareChanges); err != nil {
				//delete orphaned share if update fails
				if d_err := w.Shares.DeleteByName(ctx, si.Name); d_err != nil {
					return d_err
				}
				return err
			}
			d.SetId(si.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			si, err := NewSharesAPI(ctx, c).get(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(si, shareSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			beforeSi, err := NewSharesAPI(ctx, c).get(d.Id())
			if err != nil {
				return err
			}
			var afterSi ShareInfo
			common.DataToStructPointer(d, shareSchema, &afterSi)
			changes := beforeSi.Diff(afterSi)

			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}

			if d.HasChange("owner") {
				_, err = w.Shares.Update(ctx, sharing.UpdateShare{
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

			err = NewSharesAPI(ctx, c).update(d.Id(), ShareUpdates{
				Updates: changes,
			})
			if err != nil {
				if d.HasChange("owner") {
					// Rollback
					old, new := d.GetChange("owner")
					_, rollbackErr := w.Shares.Update(ctx, sharing.UpdateShare{
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
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.Shares.DeleteByName(ctx, d.Id())
		},
	}
}
