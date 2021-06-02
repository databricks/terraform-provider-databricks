package identity

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceUser manages users within workspace
func ResourceUser() *schema.Resource {
	userSchema := map[string]*schema.Schema{
		"user_name": {
			Type:     schema.TypeString,
			ForceNew: true,
		},
		"display_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"active": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
	}
	addEntitlementsToSchema(&userSchema)
	return common.Resource{
		Schema: userSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ru ScimUser
			if err := common.DataToStructPointer(d, userSchema, &ru); err != nil {
				return err
			}
			user, err := NewUsersAPI(ctx, c).Create(ru)
			if err != nil {
				return err
			}
			d.SetId(user.ID)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			user, err := NewUsersAPI(ctx, c).read(d.Id())
			if err != nil {
				return err
			}
			d.Set("user_name", user.UserName)
			d.Set("display_name", user.DisplayName)
			d.Set("active", user.Active)
			return user.Entitlements.Read(d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ru ScimUser
			if err := common.DataToStructPointer(d, userSchema, &ru); err != nil {
				return err
			}
			ru.Entitlements = CreateEntitlements(d)
			return NewUsersAPI(ctx, c).Update(d.Id(), ru)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewUsersAPI(ctx, c).Delete(d.Id())
		},
	}.ToResource()
}
