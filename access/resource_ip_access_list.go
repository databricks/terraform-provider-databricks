package access

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"

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

// Preview feature: https://docs.databricks.com/security/network/ip-access-list.html
// REST API: https://docs.databricks.com/dev-tools/api/latest/ip-access-list.html#operation/create-list
type ipAccessListsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// NewIPAccessListsAPI ...
func NewIPAccessListsAPI(ctx context.Context, m interface{}) ipAccessListsAPI {
	return ipAccessListsAPI{
		client:  m.(*common.DatabricksClient),
		context: ctx,
	}
}

// Create creates the IP Access List to given the instance pool configuration
func (a ipAccessListsAPI) Create(cr createIPAccessListRequest) (status IpAccessListStatus, err error) {
	wrapper := IpAccessListStatusWrapper{}
	err = a.client.Post(a.context, "/ip-access-lists", cr, &wrapper)
	if err != nil {
		return
	}
	status = wrapper.IPAccessList
	return
}

func (a ipAccessListsAPI) Update(objectID string, ur ipAccessListUpdateRequest) error {
	return a.client.Put(a.context, "/ip-access-lists/"+objectID, ur)
}

func (a ipAccessListsAPI) Delete(objectID string) (err error) {
	err = a.client.Delete(a.context, "/ip-access-lists/"+objectID, map[string]interface{}{})
	return
}

func (a ipAccessListsAPI) Read(objectID string) (status IpAccessListStatus, err error) {
	wrapper := IpAccessListStatusWrapper{}
	err = a.client.Get(a.context, "/ip-access-lists/"+objectID, nil, &wrapper)
	status = wrapper.IPAccessList
	return
}

func (a ipAccessListsAPI) List() (listResponse ListIPAccessListsResponse, err error) {
	listResponse = ListIPAccessListsResponse{}
	err = a.client.Get(a.context, "/ip-access-lists", nil, &listResponse)
	return
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
			var iacl createIPAccessListRequest
			common.DataToStructPointer(d, s, &iacl)
			status, err := NewIPAccessListsAPI(ctx, c).Create(iacl)
			if err != nil {
				return err
			}
			d.SetId(status.ListID)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			status, err := NewIPAccessListsAPI(ctx, c).Read(d.Id())
			if err != nil {
				return err
			}
			common.StructToData(status, s, d)
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var iacl ipAccessListUpdateRequest
			common.DataToStructPointer(d, s, &iacl)
			return NewIPAccessListsAPI(ctx, c).Update(d.Id(), iacl)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewIPAccessListsAPI(ctx, c).Delete(d.Id())
		},
	}.ToResource()
}
