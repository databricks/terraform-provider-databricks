package mws

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// NewNetworksAPI creates MWSNetworksAPI instance from provider meta
func NewNetworksAPI(m interface{}) NetworksAPI {
	return NetworksAPI{m.(*common.DatabricksClient), context.TODO()}
}

// NetworksAPI exposes the mws networks API
type NetworksAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates a set of MWS Networks for the BYOVPC
func (a NetworksAPI) Create(mwsAcctID, networkName string, vpcID string, subnetIds []string, securityGroupIds []string) (Network, error) {
	var mwsNetwork Network
	networksAPIPath := fmt.Sprintf("/accounts/%s/networks", mwsAcctID)
	err := a.client.Post(a.context, networksAPIPath, Network{
		NetworkName:      networkName,
		VPCID:            vpcID,
		SubnetIds:        subnetIds,
		SecurityGroupIds: securityGroupIds,
	}, &mwsNetwork)
	return mwsNetwork, err
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
	return a.client.Delete(a.context, networksAPIPath, nil)
}

// List lists all the available network objects in the mws account
func (a NetworksAPI) List(mwsAcctID string) ([]Network, error) {
	var mwsNetworkList []Network
	networksAPIPath := fmt.Sprintf("/accounts/%s/networks", mwsAcctID)
	err := a.client.Get(a.context, networksAPIPath, nil, &mwsNetworkList)
	return mwsNetworkList, err
}

func ResourceNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceMWSNetworksCreate,
		Read:   resourceMWSNetworksRead,
		Delete: resourceMWSNetworksDelete,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				ForceNew:  true,
			},
			"network_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(4, 256),
				Required:     true,
				ForceNew:     true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_ids": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				MinItems: 2,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"security_group_ids": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				MinItems: 1,
				MaxItems: 5,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vpc_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"error_messages": {
				Deprecated: "`error_messages` are deprecated and would be removed in 0.3",
				Type:       schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_message": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
				Optional: true,
				Computed: true,
			},
			"workspace_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"creation_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceMWSNetworksCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	networkName := d.Get("network_name").(string)
	mwsAcctID := d.Get("account_id").(string)
	VPCID := d.Get("vpc_id").(string)
	subnetIds := internal.ConvertListInterfaceToString(d.Get("subnet_ids").(*schema.Set).List())
	securityGroupIds := internal.ConvertListInterfaceToString(d.Get("security_group_ids").(*schema.Set).List())

	network, err := NewNetworksAPI(client).Create(mwsAcctID, networkName, VPCID, subnetIds, securityGroupIds)
	if err != nil {
		return err
	}
	networksResourceID := PackagedMWSIds{
		MwsAcctID:  mwsAcctID,
		ResourceID: network.NetworkID,
	}
	d.SetId(packMWSAccountID(networksResourceID))
	return resourceMWSNetworksRead(d, m)
}

func resourceMWSNetworksRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)
	packagedMwsID, err := UnpackMWSAccountID(id)
	if err != nil {
		return err
	}
	network, err := NewNetworksAPI(client).Read(packagedMwsID.MwsAcctID, packagedMwsID.ResourceID)
	if err != nil {
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		return err
	}
	err = d.Set("network_name", network.NetworkName)
	if err != nil {
		return err
	}
	err = d.Set("vpc_id", network.VPCID)
	if err != nil {
		return err
	}
	err = d.Set("subnet_ids", network.SubnetIds)
	if err != nil {
		return err
	}
	err = d.Set("security_group_ids", network.SecurityGroupIds)
	if err != nil {
		return err
	}
	err = d.Set("vpc_status", network.VPCStatus)
	if err != nil {
		return err
	}

	if !reflect.ValueOf(network.ErrorMessages).IsZero() {
		// TODO: should this really be a state or rather error return?
		err = d.Set("error_messages", convertErrorMessagesToListOfMaps(network.ErrorMessages))
		if err != nil {
			return err
		}
	}

	err = d.Set("workspace_id", network.WorkspaceID)
	if err != nil {
		return err
	}
	err = d.Set("account_id", network.AccountID)
	if err != nil {
		return err
	}
	err = d.Set("creation_time", network.CreationTime)
	if err != nil {
		return err
	}
	err = d.Set("network_id", network.NetworkID)
	if err != nil {
		return err
	}

	return nil
}

func resourceMWSNetworksDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)
	packagedMwsID, err := UnpackMWSAccountID(id)
	if err != nil {
		return err
	}
	networksAPI := NewNetworksAPI(client)
	err = networksAPI.Delete(packagedMwsID.MwsAcctID, packagedMwsID.ResourceID)
	if err != nil {
		return err
	}
	return resource.RetryContext(networksAPI.context, 60*time.Second, func() *resource.RetryError {
		network, err := networksAPI.Read(packagedMwsID.MwsAcctID, packagedMwsID.ResourceID)
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("[INFO] Network %s is removed.", packagedMwsID.ResourceID)
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

func convertErrorMessagesToListOfMaps(errorMsgs []NetworkHealth) []map[string]string {
	var resp []map[string]string
	for _, errorMsg := range errorMsgs {
		errorMap := map[string]string{}
		errorMap["error_type"] = errorMsg.ErrorType
		errorMap["error_message"] = errorMsg.ErrorMessage
		resp = append(resp, errorMap)
	}
	return resp
}
