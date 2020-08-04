package databricks

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceMWSNetworks() *schema.Resource {
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
				Type: schema.TypeList,
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
	client := m.(*service.DatabricksClient)
	networkName := d.Get("network_name").(string)
	mwsAcctID := d.Get("account_id").(string)
	VPCID := d.Get("vpc_id").(string)
	subnetIds := convertListInterfaceToString(d.Get("subnet_ids").(*schema.Set).List())
	securityGroupIds := convertListInterfaceToString(d.Get("security_group_ids").(*schema.Set).List())

	network, err := client.MWSNetworks().Create(mwsAcctID, networkName, VPCID, subnetIds, securityGroupIds)
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
	client := m.(*service.DatabricksClient)
	packagedMwsID, err := unpackMWSAccountID(id)
	if err != nil {
		return err
	}
	network, err := client.MWSNetworks().Read(packagedMwsID.MwsAcctID, packagedMwsID.ResourceID)
	if err != nil {
		if e, ok := err.(service.APIError); ok && e.IsMissing() {
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
	client := m.(*service.DatabricksClient)
	packagedMwsID, err := unpackMWSAccountID(id)
	if err != nil {
		return err
	}
	err = client.MWSNetworks().Delete(packagedMwsID.MwsAcctID, packagedMwsID.ResourceID)
	if err != nil {
		return err
	}
	return resource.Retry(60*time.Second, func() *resource.RetryError {
		network, err := client.MWSNetworks().Read(packagedMwsID.MwsAcctID, packagedMwsID.ResourceID)
		if e, ok := err.(service.APIError); ok && e.IsMissing() {
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

func convertErrorMessagesToListOfMaps(errorMsgs []model.NetworkHealth) []map[string]string {
	var resp []map[string]string
	for _, errorMsg := range errorMsgs {
		errorMap := map[string]string{}
		errorMap["error_type"] = errorMsg.ErrorType
		errorMap["error_message"] = errorMsg.ErrorMessage
		resp = append(resp, errorMap)
	}
	return resp
}
