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
			group, err := NewGroupsAPI(ctx, c).Create(ScimGroup{
				DisplayName:  groupName,
				Entitlements: readEntitlementsFromData(d),
			})
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
			return group.Entitlements.readIntoData(d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			groupName := d.Get("display_name").(string)
			return NewGroupsAPI(ctx, c).UpdateNameAndEntitlements(d.Id(), groupName, readEntitlementsFromData(d))
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewGroupsAPI(ctx, c).Delete(d.Id())
		},
		Schema: groupSchema,
	}.ToResource()
}
