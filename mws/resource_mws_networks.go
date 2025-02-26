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
		common.CustomizeSchemaPath(s, "gcp_network_info", "pod_ip_range_name").SetDeprecated(getGkeDeprecationMessage("gcp_network_info.pod_ip_range_name"))
		common.CustomizeSchemaPath(s, "gcp_network_info", "service_ip_range_name").SetDeprecated(getGkeDeprecationMessage("gcp_network_info.pod_ip_range_name"))

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
			// If gcp_network_info.0.pod_ip_range_name or gcp_network_info.0.service_ip_range_name are
			// unset in the plan, remove them from the returned network so that they are not persisted
			// in state.
			if v, ok := d.Get("gcp_network_info.0.pod_ip_range_name").(string); ok && v == "" && network.GcpNetworkInfo != nil {
				network.GcpNetworkInfo.PodIpRangeName = ""
			}
			if v, ok := d.Get("gcp_network_info.0.service_ip_range_name").(string); ok && v == "" && network.GcpNetworkInfo != nil {
				network.GcpNetworkInfo.ServiceIpRangeName = ""
			}
			return common.StructToData(network, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			// This will only be called when removing gcp_network_info.0.pod_ip_range_name or
			// gcp_network_info.0.service_ip_range_name. This is a no-op and doesn't require
			// making any API call.
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, networkID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return NewNetworksAPI(ctx, c).Delete(accountID, networkID)
		},
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			// For `gcp_network_info.0.pod_ip_range_name` or `gcp_network_info.0.service_ip_range_name`,
			// users should be able to remove these keys without recreating their networks as part of the
			// GKE deprecation process.
			//
			// Otherwise, any change for these keys or any change for any other key will cause the
			// network resource to be recreated.
			//
			// This should only run on update, thus we skip this check if the ID is not known.
			if d.Id() != "" {
				for _, key := range d.GetChangedKeysPrefix("") {
					v, ok := d.Get(key).(string)
					if ok && v == "" && (key == "gcp_network_info.0.pod_ip_range_name" || key == "gcp_network_info.0.service_ip_range_name") {
						continue
					}
					if err := d.ForceNew(key); err != nil {
						return err
					}
					break
				}
			}
			return nil
		},
	}
}
