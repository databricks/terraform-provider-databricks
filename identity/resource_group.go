package identity

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceGroup manages user groups
func ResourceGroup() *schema.Resource {
	groupSchema := map[string]*schema.Schema{
		"display_name": {
			Type:     schema.TypeString,
			ForceNew: true,
			Required: true,
		},
		"url": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
	addEntitlementsToSchema(&groupSchema)
	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			groupName := d.Get("display_name").(string)
			entitlements := CreateEntitlements(d)
			group, err := NewGroupsAPI(ctx, c).Create(groupName, nil, nil, entitlements)
			if err != nil {
				return err
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
			d.Set("url", c.FormatURL("#setting/accounts/groups/", d.Id()))
			return group.Entitlements.Read(d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			// Handle entitlements update
			added, removed := UpdateEntitlements(d)
			// TODO: not currently possible to update group display name
			return NewGroupsAPI(ctx, c).Patch(d.Id(), added, removed, GroupEntitlementsPath)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewGroupsAPI(ctx, c).Delete(d.Id())
		},
		Schema: groupSchema,
	}.ToResource()
}
