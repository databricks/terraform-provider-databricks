package identity

import (
	"log"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupCreate,
		Update: resourceGroupUpdate,
		Read:   resourceGroupRead,
		Delete: resourceGroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
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
	}
}

func resourceGroupCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	groupName := d.Get("display_name").(string)
	allowClusterCreate := d.Get("allow_cluster_create").(bool)
	allowInstancePoolCreate := d.Get("allow_instance_pool_create").(bool)

	// If entitlement flags are set to be true
	var entitlementsList []string
	if allowClusterCreate {
		entitlementsList = append(entitlementsList, string(AllowClusterCreateEntitlement))
	}
	if allowInstancePoolCreate {
		entitlementsList = append(entitlementsList, string(AllowInstancePoolCreateEntitlement))
	}

	group, err := NewGroupsAPI(client).Create(groupName, nil, nil, entitlementsList)
	if err != nil {
		return err
	}
	d.SetId(group.ID)
	return resourceGroupRead(d, m)
}

func resourceGroupRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)
	group, err := NewGroupsAPI(client).Read(id)
	if err != nil {
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
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
	client := m.(*common.DatabricksClient)

	// Handle entitlements update
	var entitlementsAddList []string
	var entitlementsRemoveList []string
	// If allow_cluster_create has changed
	if d.HasChange("allow_cluster_create") {
		allowClusterCreate := d.Get("allow_cluster_create").(bool)
		// Changed to true
		if allowClusterCreate {
			entitlementsAddList = append(entitlementsAddList, string(AllowClusterCreateEntitlement))
		}
		// Changed to false
		entitlementsRemoveList = append(entitlementsRemoveList, string(AllowClusterCreateEntitlement))
	}
	// If allow_instance_pool_create has changed
	if d.HasChange("allow_instance_pool_create") {
		allowClusterCreate := d.Get("allow_instance_pool_create").(bool)
		// Changed to true
		if allowClusterCreate {
			entitlementsAddList = append(entitlementsAddList, string(AllowClusterCreateEntitlement))
		}
		// Changed to false
		entitlementsRemoveList = append(entitlementsRemoveList, string(AllowClusterCreateEntitlement))
	}

	// TODO: not currently possible to update group display name
	if entitlementsAddList != nil || entitlementsRemoveList != nil {
		err := NewGroupsAPI(client).Patch(id, entitlementsAddList, entitlementsRemoveList, GroupEntitlementsPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func resourceGroupDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)
	err := NewGroupsAPI(client).Delete(id)
	return err
}

func isGroupClusterCreateEntitled(group *Group) bool {
	for _, entitlement := range group.Entitlements {
		if entitlement.Value == AllowClusterCreateEntitlement {
			return true
		}
	}
	return false
}

func isGroupInstancePoolCreateEntitled(group *Group) bool {
	for _, entitlement := range group.Entitlements {
		if entitlement.Value == AllowClusterCreateEntitlement {
			return true
		}
	}
	return false
}
