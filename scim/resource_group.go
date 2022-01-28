package scim

import (
	"context"
	"fmt"
	"strings"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceGroup manages user groups
func ResourceGroup() *schema.Resource {
	groupSchema := map[string]*schema.Schema{
		"display_name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"url": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"external_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"force": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
	addEntitlementsToSchema(&groupSchema)
	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			groupName := d.Get("display_name").(string)
			g := Group{
				DisplayName:  groupName,
				Entitlements: readEntitlementsFromData(d),
				ExternalID:   d.Get("external_id").(string),
			}
			groupsAPI := NewGroupsAPI(ctx, c)
			group, err := groupsAPI.Create(g)
			if err != nil {
				return createForceOverridesManuallyAddedGroup(err, d, groupsAPI, g)
			}
			d.SetId(group.ID)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			group, err := NewGroupsAPI(ctx, c).Read(d.Id())
			if err != nil {
				return err
			}
			d.Set("display_name", group.DisplayName)
			d.Set("external_id", group.ExternalID)
			d.Set("url", c.FormatURL("#setting/accounts/groups/", d.Id()))
			return group.Entitlements.readIntoData(d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			groupName := d.Get("display_name").(string)
			return NewGroupsAPI(ctx, c).UpdateNameAndEntitlements(d.Id(), groupName,
				d.Get("external_id").(string), readEntitlementsFromData(d))
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewGroupsAPI(ctx, c).Delete(d.Id())
		},
		Schema: groupSchema,
	}.ToResource()
}

func createForceOverridesManuallyAddedGroup(err error, d *schema.ResourceData, groupsAPI GroupsAPI, g Group) error {
	forceCreate := d.Get("force").(bool)
	if !forceCreate {
		return err
	}
	// corner-case for overriding manually provisioned groups
	groupName := strings.ReplaceAll(g.DisplayName, "'", "")
	force := fmt.Sprintf("Group with name %s already exists.", groupName)
	if err.Error() != force {
		return err
	}
	group, err := groupsAPI.ReadByDisplayName(groupName)
	if err != nil {
		return err
	}
	d.SetId(group.ID)
	return groupsAPI.UpdateNameAndEntitlements(d.Id(), g.DisplayName, g.ExternalID, g.Entitlements)
}
