package catalog

import (
	"context"

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
)

type ShareInfo struct {
	Name      string             `json:"name" tf:"force_new"`
	Objects   []SharedDataObject `json:"objects,omitempty" tf:"slice_set,alias:object"`
	CreatedAt int64              `json:"created_at,omitempty" tf:"computed"`
	CreatedBy string             `json:"created_by,omitempty" tf:"computed"`
}

type SharedDataObject struct {
	Name           string `json:"name"`
	DataObjectType string `json:"data_object_type"`
	Comment        string `json:"comment,omitempty"`
	SharedAs       string `json:"shared_as,omitempty" tf:"computed"`
	AddedAt        int64  `json:"added_at,omitempty" tf:"computed"`
	AddedBy        string `json:"added_by,omitempty" tf:"computed"`
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

func (a SharesAPI) list() (shares Shares, err error) {
	err = a.client.Get(a.context, "/unity-catalog/shares", nil, &shares)
	return
}

func (a SharesAPI) create(ci *ShareInfo) error {
	return a.client.Post(a.context, "/unity-catalog/shares", ci, ci)
}

func (a SharesAPI) get(name string) (ci ShareInfo, err error) {
	err = a.client.Get(a.context, "/unity-catalog/shares/"+name+"?include_shared_data=true", nil, &ci)
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
	for _, afterSdo := range other.Objects {
		_, exists := beforeMap[afterSdo.Name]
		if exists {
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
				if u_err := NewSharesAPI(ctx, c).delete(si.Name); err != nil {
					return u_err
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
