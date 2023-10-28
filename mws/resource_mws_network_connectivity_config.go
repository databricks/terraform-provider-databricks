package mws

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// NewNetworkConnectivityConfigAPI creates NetworkConnectivityConfigAPI instance from provider meta
func NewNetworkConnectivityConfigAPI(ctx context.Context, m any) NetworkConnectivityConfigAPI {
	return NetworkConnectivityConfigAPI{m.(*common.DatabricksClient), ctx}
}

// NetworkConnectivityConfigAPI exposes the mws network connectivity config API
type NetworkConnectivityConfigAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates network connectivity config objects.
func (api NetworkConnectivityConfigAPI) Create(ncc *NetworkConnectivityConfig) error {
	nccAPIPath := fmt.Sprintf("/accounts/%s/network-connectivity-configs", ncc.AccountID)
	return api.client.Post(api.context, nccAPIPath, ncc, &ncc)
}

// Read returns the network connectivity config objects with egress rules
func (api NetworkConnectivityConfigAPI) Read(mwsAcctID, nccID string) (NetworkConnectivityConfig, error) {
	var ncc NetworkConnectivityConfig
	nccAPIPath := fmt.Sprintf("/accounts/%s/network-connectivity-configs/%s", mwsAcctID, nccID)
	err := api.client.Get(api.context, nccAPIPath, nil, &ncc)
	return ncc, err
}

// Delete deletes the NCC object by account ID and NCC ID
func (api NetworkConnectivityConfigAPI) Delete(mwsAcctID, nccID string) error {
	nccAPIPath := fmt.Sprintf("/accounts/%s/network-connectivity-configs/%s", mwsAcctID, nccID)
	if err := api.client.Delete(api.context, nccAPIPath, nil); err != nil {
		return err
	}
	return nil
}

// List lists all the available NCC objects in the mws account
func (api NetworkConnectivityConfigAPI) List(mwsAcctID string) ([]NetworkConnectivityConfig, error) {
	var nccList []NetworkConnectivityConfig
	nccAPIPath := fmt.Sprintf("/accounts/%s/network-connectivity-configs", mwsAcctID)
	err := api.client.Get(api.context, nccAPIPath, nil, &nccList)
	return nccList, err
}

func ResourceMwsNetworkConnectivityConfig() *schema.Resource {
	s := common.StructToSchema(NetworkConnectivityConfig{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s["name"].ValidateFunc = validation.StringLenBetween(3, 30)
		return s
	})
	p := common.NewPairSeparatedID("account_id", "network_connectivity_config_id", "/")
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ncc NetworkConnectivityConfig
			common.DataToStructPointer(d, s, &ncc)
			if err := NewNetworkConnectivityConfigAPI(ctx, c).Create(&ncc); err != nil {
				return err
			}
			d.Set("network_connectivity_config_id", ncc.NetworkConnectivityConfigID)
			p.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, nccID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			ncc, err := NewNetworkConnectivityConfigAPI(ctx, c).Read(accountID, nccID)
			if err != nil {
				return err
			}
			return common.StructToData(ncc, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, nccID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return NewNetworkConnectivityConfigAPI(ctx, c).Delete(accountID, nccID)
		},
	}.ToResource()
}
