package catalog

import (
	"context"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SharesAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewSharesAPI(ctx context.Context, m interface{}) SharesAPI {
	return SharesAPI{m.(*common.DatabricksClient), ctx}
}

const (
	ShareAdd    = "ADD"
	ShareRemove = "REMOVE"
)

type ShareInfo struct {
	Name      string             `json:"name" tf:"force_new"`
	Objects   []SharedDataObject `json:"objects,omitempty" tf:"slice_set"`
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

func (a SharesAPI) createShare(ci *ShareInfo) error {
	return a.client.Post(a.context, "/unity-catalog/shares", ci, ci)
}

func (a SharesAPI) getShare(name string) (ci ShareInfo, err error) {
	err = a.client.Get(a.context, "/unity-catalog/shares/"+name+"?include_shared_data=true", nil, &ci)
	return
}

func (a SharesAPI) updateShare(name string, sdc []ShareDataChange) error {
	if sdc == nil {
		return nil
	}
	return a.client.Patch(a.context, "/unity-catalog/shares/"+name, ShareUpdates{
		Updates: sdc,
	})
}

func (a SharesAPI) deleteShare(name string) error {
	return a.client.Delete(a.context, "/unity-catalog/shares/"+name, nil)
}

func shareObjectsToShareChange(objects []SharedDataObject, action string) []ShareDataChange {
	var changes []ShareDataChange
	for _, obj := range objects {
		changes = append(changes, ShareDataChange{
			Action:     action,
			DataObject: obj,
		})
	}
	return changes
}

func getResourceShareMap(sdos []SharedDataObject) map[string]SharedDataObject {
	m := make(map[string]SharedDataObject, len(sdos))
	for _, sdo := range sdos {
		m[sdo.Name] = sdo
	}
	return m
}

func getResourceShareChanges(before []SharedDataObject, after []SharedDataObject) []ShareDataChange {
	beforeMap := getResourceShareMap(before)
	afterMap := getResourceShareMap(after)
	var changes []ShareDataChange
	// not in after so remove
	for _, beforeSdo := range before {
		if _, ok := afterMap[beforeSdo.Name]; !ok {
			changes = append(changes, ShareDataChange{
				Action:     ShareRemove,
				DataObject: beforeSdo,
			})
		}
	}

	// not in before so add
	for _, afterSdo := range after {
		if _, ok := beforeMap[afterSdo.Name]; !ok {
			changes = append(changes, ShareDataChange{
				Action:     ShareAdd,
				DataObject: afterSdo,
			})
		}
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
			if err := NewSharesAPI(ctx, c).createShare(&si); err != nil {
				return err
			}
			shareChanges := shareObjectsToShareChange(si.Objects, ShareAdd)
			if err := NewSharesAPI(ctx, c).updateShare(si.Name, shareChanges); err != nil {
				return err
			}
			d.SetId(si.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			si, err := NewSharesAPI(ctx, c).getShare(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(si, shareSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			si, err := NewSharesAPI(ctx, c).getShare(d.Id())
			if err != nil {
				return err
			}
			before := si.Objects
			var shareInfo ShareInfo
			common.DataToStructPointer(d, shareSchema, &shareInfo)
			after := shareInfo.Objects
			changes := getResourceShareChanges(before, after)
			return NewSharesAPI(ctx, c).updateShare(d.Id(), changes)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewSharesAPI(ctx, c).deleteShare(d.Id())
		},
	}.ToResource()
}
