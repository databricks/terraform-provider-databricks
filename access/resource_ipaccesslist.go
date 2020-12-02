package access

import (
	"context"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/databrickslabs/databricks-terraform/internal/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type listIPAccessListsResponse struct {
	ListIPAccessListsResponse []ipAccessListStatus `json:"ip_access_lists,omitempty"`
}

type createIPAccessListRequest struct {
	Label       string   `json:"label"`
	ListType    string   `json:"list_type"`
	IPAddresses []string `json:"ip_addresses"`
}

type ipAccessListStatus struct {
	ListID        string   `json:"list_id,omitempty"`
	Label         string   `json:"label,omitempty"`
	ListType      string   `json:"list_type,omitempty"`
	IPAddresses   []string `json:"ip_addresses,omitempty"`
	AddressCount  int      `json:"address_count,omitempty"`
	CreatedAt     int64    `json:"created_at,omitempty"`
	CreatorUserID int64    `json:"creator_user_id,omitempty"`
	UpdatedAt     int64    `json:"updated_at,omitempty"`
	UpdatorUserID int64    `json:"updator_user_id,omitempty"`
	Enabled       bool     `json:"enabled,omitempty"`
}

type ipAccessListStatusWrapper struct {
	IPAccessList ipAccessListStatus `json:"ip_access_list,omitempty"`
}

type ipAccessListUpdateRequest struct {
	Label       string   `json:"label"`
	ListType    string   `json:"list_type"`
	IPAddresses []string `json:"ip_addresses"`
	Enabled     bool     `json:"enabled,omitempty"`
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
func (a ipAccessListsAPI) Create(cr createIPAccessListRequest) (status ipAccessListStatus, err error) {
	wrapper := ipAccessListStatusWrapper{}
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
	err = a.client.Delete(a.context, "/ip-access-lists/"+objectID, nil)
	return
}

func (a ipAccessListsAPI) Read(objectID string) (status ipAccessListStatus, err error) {
	wrapper := ipAccessListStatusWrapper{}
	err = a.client.Get(a.context, "/ip-access-lists/"+objectID, nil, &wrapper)
	status = wrapper.IPAccessList
	return
}

func (a ipAccessListsAPI) List() (listResponse listIPAccessListsResponse, err error) {
	listResponse = listIPAccessListsResponse{}
	err = a.client.Get(a.context, "/ip-access-lists", &listResponse, nil)
	return
}

// ResourceIPAccessList manages IP access lists
func ResourceIPAccessList() *schema.Resource {
	s := internal.StructToSchema(ipAccessListUpdateRequest{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		// nolint
		s["list_type"].ValidateFunc = validation.StringInSlice([]string{"ALLOW", "BLOCK"}, false)
		s["ip_addresses"].Elem = &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.Any(validation.IsIPv4Address, validation.IsCIDR),
		}
		s["enabled"].Default = true
		return s
	})
	return util.CommonResource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var iacl createIPAccessListRequest
			if err := internal.DataToStructPointer(d, s, &iacl); err != nil {
				return err
			}
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
			return internal.StructToData(status, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var iacl ipAccessListUpdateRequest
			if err := internal.DataToStructPointer(d, s, &iacl); err != nil {
				return err
			}
			return NewIPAccessListsAPI(ctx, c).Update(d.Id(), iacl)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewIPAccessListsAPI(ctx, c).Delete(d.Id())
		},
	}.ToResource()
}
