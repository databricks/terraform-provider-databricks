package databricks

import (
	"log"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupCreate,
		Update: resourceGroupUpdate,
		Read:   resourceGroupRead,
		Delete: resourceGroupDelete,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"allow_cluster_create": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_instance_pool_create": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceGroupCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	groupName := d.Get("display_name").(string)
	allowClusterCreate := d.Get("allow_cluster_create").(bool)
	allowInstancePoolCreate := d.Get("allow_instance_pool_create").(bool)

	// If entitlement flags are set to be true
	var entitlementsList []string
	if allowClusterCreate {
		entitlementsList = append(entitlementsList, string(model.AllowClusterCreateEntitlement))
	}
	if allowInstancePoolCreate {
		entitlementsList = append(entitlementsList, string(model.AllowInstancePoolCreateEntitlement))
	}

	group, err := client.Groups().Create(groupName, nil, nil, entitlementsList)
	if err != nil {
		return err
	}
	d.SetId(group.ID)
	return resourceGroupRead(d, m)
}

func resourceGroupRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	group, err := client.Groups().Read(id)
	if err != nil {
		if isScimGroupMissing(err.Error(), id) {
			log.Printf("Missing scim group with id: %s.", id)
			d.SetId("")
			return nil
		}
		return err
	}

	err = d.Set("display_name", group.DisplayName)
	if err != nil {
		return err
	}

	err = d.Set("allow_cluster_create", isGroupClusterCreateEntitled(&group))
	if err != nil {
		return err
	}

	err = d.Set("allow_instance_pool_create", isGroupInstancePoolCreateEntitled(&group))
	return err
}

func resourceGroupUpdate(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)

	// Handle entitlements update
	var entitlementsAddList []string
	var entitlementsRemoveList []string
	// If allow_cluster_create has changed
	if d.HasChange("allow_cluster_create") {
		allowClusterCreate := d.Get("allow_cluster_create").(bool)
		// Changed to true
		if allowClusterCreate {
			entitlementsAddList = append(entitlementsAddList, string(model.AllowClusterCreateEntitlement))
		}
		// Changed to false
		entitlementsRemoveList = append(entitlementsRemoveList, string(model.AllowClusterCreateEntitlement))
	}
	// If allow_instance_pool_create has changed
	if d.HasChange("allow_instance_pool_create") {
		allowClusterCreate := d.Get("allow_instance_pool_create").(bool)
		// Changed to true
		if allowClusterCreate {
			entitlementsAddList = append(entitlementsAddList, string(model.AllowClusterCreateEntitlement))
		}
		// Changed to false
		entitlementsRemoveList = append(entitlementsRemoveList, string(model.AllowClusterCreateEntitlement))
	}

	if entitlementsAddList != nil || entitlementsRemoveList != nil {
		err := client.Groups().Patch(id, entitlementsAddList, entitlementsRemoveList, model.GroupEntitlementsPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func resourceGroupDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	err := client.Groups().Delete(id)
	return err
}

func isGroupClusterCreateEntitled(group *model.Group) bool {
	for _, entitlement := range group.Entitlements {
		if entitlement.Value == model.AllowClusterCreateEntitlement {
			return true
		}
	}
	return false
}

func isGroupInstancePoolCreateEntitled(group *model.Group) bool {
	for _, entitlement := range group.Entitlements {
		if entitlement.Value == model.AllowClusterCreateEntitlement {
			return true
		}
	}
	return false
}
