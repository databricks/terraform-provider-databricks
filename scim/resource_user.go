package scim

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceUser manages users within workspace
func ResourceUser() *schema.Resource {
	type entity struct {
		UserName    string `json:"user_name" tf:"force_new"`
		DisplayName string `json:"display_name,omitempty" tf:"computed"`
		Active      bool   `json:"active,omitempty"`
		ExternalID  string `json:"external_id,omitempty"`
	}
	userSchema := common.StructToSchema(entity{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			addEntitlementsToSchema(&m)
			m["active"].Default = true
			return m
		})
	scimUserFromData := func(d *schema.ResourceData) (user User, err error) {
		var u entity
		if err = common.DataToStructPointer(d, userSchema, &u); err != nil {
			return
		}
		return User{
			UserName:     u.UserName,
			DisplayName:  u.DisplayName,
			Active:       u.Active,
			Entitlements: readEntitlementsFromData(d),
			ExternalID:   u.ExternalID,
		}, nil
	}
	return common.Resource{
		Schema: userSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			u, err := scimUserFromData(d)
			if err != nil {
				return err
			}
			user, err := NewUsersAPI(ctx, c).Create(u)
			if err != nil {
				return err
			}
			d.SetId(user.ID)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			user, err := NewUsersAPI(ctx, c).Read(d.Id())
			if err != nil {
				return err
			}
			d.Set("user_name", user.UserName)
			d.Set("display_name", user.DisplayName)
			d.Set("active", user.Active)
			d.Set("external_id", user.ExternalID)
			return user.Entitlements.readIntoData(d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			u, err := scimUserFromData(d)
			if err != nil {
				return err
			}
			return NewUsersAPI(ctx, c).Update(d.Id(), u)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewUsersAPI(ctx, c).Delete(d.Id())
		},
	}.ToResource()
}
