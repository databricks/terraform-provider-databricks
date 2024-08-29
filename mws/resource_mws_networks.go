package mws

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// NewNetworksAPI creates MWSNetworksAPI instance from provider meta
func NewNetworksAPI(ctx context.Context, m any) NetworksAPI {
	return NetworksAPI{m.(*common.DatabricksClient), ctx}
}

// NetworksAPI exposes the mws networks API
type NetworksAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates a set of MWS Networks for the BYOVPC
func (a NetworksAPI) Create(network *Network) error {
	networksAPIPath := fmt.Sprintf("/accounts/%s/networks", network.AccountID)
	return a.client.Post(a.context, networksAPIPath, network, &network)
}

// Read returns the network object along with metadata and any additional errors when attaching to workspace
func (a NetworksAPI) Read(mwsAcctID, networksID string) (Network, error) {
	var mwsNetwork Network
	networksAPIPath := fmt.Sprintf("/accounts/%s/networks/%s", mwsAcctID, networksID)
	err := a.client.Get(a.context, networksAPIPath, nil, &mwsNetwork)
	return mwsNetwork, err
}

// Delete deletes the network object given a network id
func (a NetworksAPI) Delete(mwsAcctID, networksID string) error {
	networksAPIPath := fmt.Sprintf("/accounts/%s/networks/%s", mwsAcctID, networksID)
	if err := a.client.Delete(a.context, networksAPIPath, nil); err != nil {
		return err
	}
	return resource.RetryContext(a.context, 60*time.Second, func() *resource.RetryError {
		network, err := a.Read(mwsAcctID, networksID)
		if apierr.IsMissing(err) {
			log.Printf("[INFO] Network %s/%s is removed.", mwsAcctID, networksID)
			return nil
		}
		if err != nil {
			return resource.NonRetryableError(err)
		}
		msg := fmt.Errorf("Network %s is not removed yet. VPC Status: %s", network.NetworkName, network.VPCStatus)
		log.Printf("[INFO] %s", msg)
		return resource.RetryableError(msg)
	})
}

// List lists all the available network objects in the mws account
func (a NetworksAPI) List(mwsAcctID string) ([]Network, error) {
	var mwsNetworkList []Network
	networksAPIPath := fmt.Sprintf("/accounts/%s/networks", mwsAcctID)
	err := a.client.Get(a.context, networksAPIPath, nil, &mwsNetworkList)
	return mwsNetworkList, err
}

func ResourceMwsNetworks() common.Resource {
	s := common.StructToSchema(Network{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s["account_id"].Sensitive = true
		// nolint
		s["network_name"].ValidateFunc = validation.StringLenBetween(4, 256)
		s["subnet_ids"].MinItems = 2
		s["security_group_ids"].MinItems = 1
		s["security_group_ids"].MaxItems = 5

		s["vpc_id"].ExactlyOneOf = []string{"vpc_id", "gcp_network_info"}
		s["subnet_ids"].ExactlyOneOf = []string{"subnet_ids", "gcp_network_info"}
		s["security_group_ids"].ExactlyOneOf = []string{"security_group_ids", "gcp_network_info"}
		s["gcp_network_info"].ConflictsWith = []string{"vpc_id", "subnet_ids", "security_group_ids"}

		return s
	})
	p := common.NewPairSeparatedID("account_id", "network_id", "/")
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var network Network
			common.DataToStructPointer(d, s, &network)
			if err := NewNetworksAPI(ctx, c).Create(&network); err != nil {
				return err
			}
			d.Set("network_id", network.NetworkID)
			p.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, networkID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			network, err := NewNetworksAPI(ctx, c).Read(accountID, networkID)
			if err != nil {
				return err
			}
			return common.StructToData(network, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, networkID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return NewNetworksAPI(ctx, c).Delete(accountID, networkID)
		},
	}
}
