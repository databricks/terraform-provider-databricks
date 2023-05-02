package catalog

import (
	"context"
	"reflect"
	"sort"

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
	Updates []ShareDataChange `json:"updates"`
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

func (a SharesAPI) create(si *ShareInfo) error {
	si.sortSharesByName()
	return a.client.Post(a.context, "/unity-catalog/shares", si, si)
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
	return a.client.Patch(a.context, "/unity-catalog/shares/"+name, su)
}

func (a SharesAPI) delete(name string) error {
	return a.client.Delete(a.context, "/unity-catalog/shares/"+name, nil)
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

func Equal(a, b SharedDataObject) bool {
	if b.SharedAs == "" {
		b.SharedAs = a.SharedAs
	}
	return reflect.DeepEqual(a, b)
}

func (si ShareInfo) Diff(other ShareInfo) []ShareDataChange {
	beforeMap := si.resourceShareMap()
	afterMap := other.resourceShareMap()
	changes := []ShareDataChange{}
	// not in after so remove
	for _, beforeSdo := range si.Objects {
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
	for _, afterSdo := range other.Objects {
		beforeSdo, exists := beforeMap[afterSdo.Name]
		if exists {
			if !Equal(beforeSdo, afterSdo) {
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

func ResourceShare() *schema.Resource {
	shareSchema := common.StructToSchema(ShareInfo{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		return m
	})
	return common.Resource{
		Schema: shareSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var si ShareInfo
			common.DataToStructPointer(d, shareSchema, &si)
			if err := NewSharesAPI(ctx, c).create(&si); err != nil {
				return err
			}
			//can only create empty share, objects have to be added using update API
			shareChanges := si.shareChanges(ShareAdd)
			if err := NewSharesAPI(ctx, c).update(si.Name, shareChanges); err != nil {
				//delete orphaned share if update fails
				if d_err := NewSharesAPI(ctx, c).delete(si.Name); d_err != nil {
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
			si, err := NewSharesAPI(ctx, c).get(d.Id())
			if err != nil {
				return err
			}
			var shareInfo ShareInfo
			common.DataToStructPointer(d, shareSchema, &shareInfo)
			changes := si.Diff(shareInfo)
			return NewSharesAPI(ctx, c).update(d.Id(), ShareUpdates{
				Updates: changes,
			})
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewSharesAPI(ctx, c).delete(d.Id())
		},
	}.ToResource()
}
