package access

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type ListIPAccessListsResponse struct {
	ListIPAccessListsResponse []IpAccessListStatus `json:"ip_access_lists,omitempty"`
}

type createIPAccessListRequest struct {
	Label       string   `json:"label"`
	ListType    string   `json:"list_type"`
	IPAddresses []string `json:"ip_addresses"`
}

type IpAccessListStatus struct {
	ListID        string   `json:"list_id"`
	Label         string   `json:"label"`
	ListType      string   `json:"list_type"`
	IPAddresses   []string `json:"ip_addresses"`
	AddressCount  int      `json:"address_count,omitempty"`
	CreatedAt     int64    `json:"created_at,omitempty"`
	CreatorUserID int64    `json:"creator_user_id,omitempty"`
	UpdatedAt     int64    `json:"updated_at,omitempty"`
	UpdatorUserID int64    `json:"updator_user_id,omitempty"`
	Enabled       bool     `json:"enabled,omitempty"`
}

type IpAccessListStatusWrapper struct {
	IPAccessList IpAccessListStatus `json:"ip_access_list,omitempty"`
}

type ipAccessListUpdateRequest struct {
	Label       string   `json:"label"`
	ListType    string   `json:"list_type"`
	IPAddresses []string `json:"ip_addresses"`
	Enabled     bool     `json:"enabled,omitempty" tf:"default:true"`
}

// ResourceIPAccessList manages IP access lists
func ResourceIPAccessList() *schema.Resource {
	s := common.StructToSchema(ipAccessListUpdateRequest{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		// nolint
		s["list_type"].ValidateFunc = validation.StringInSlice([]string{"ALLOW", "BLOCK"}, false)
		s["ip_addresses"].Elem = &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.Any(validation.IsIPv4Address, validation.IsCIDR),
		}
		return s
	})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var iacl settings.CreateIpAccessList
			common.DataToStructPointer(d, s, &iacl)
			status, err := w.IpAccessLists.Create(ctx, iacl)
			if err != nil {
				return err
			}
			d.SetId(status.IpAccessList.ListId)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			status, err := w.IpAccessLists.GetByIpAccessListId(ctx, d.Id())
			if err != nil {
				return err
			}
			common.StructToData(status, s, d)
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var iacl settings.UpdateIpAccessList
			common.DataToStructPointer(d, s, &iacl)
			return w.IpAccessLists.Update(ctx, iacl)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.IpAccessLists.DeleteByIpAccessListId(ctx, d.Id())
		},
	}.ToResource()
}
