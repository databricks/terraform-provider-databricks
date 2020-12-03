package identity

import (
	"context"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/databrickslabs/databricks-terraform/internal/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceUser manages users within workspace
func ResourceUser() *schema.Resource {
	userSchema := internal.StructToSchema(UserEntity{}, func(
		s map[string]*schema.Schema) map[string]*schema.Schema {
		s["user_name"].ForceNew = true
		s["active"].Default = true
		return s
	})
	return util.CommonResource{
		Schema: userSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ru UserEntity
			if err := internal.DataToStructPointer(d, userSchema, &ru); err != nil {
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
			user, err := NewUsersAPI(ctx, c).Read(d.Id())
			if err != nil {
				return err
			}
			return internal.StructToData(user, userSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ru UserEntity
			if err := internal.DataToStructPointer(d, userSchema, &ru); err != nil {
				return err
			}
			return NewUsersAPI(ctx, c).Update(d.Id(), ru)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewUsersAPI(ctx, c).Delete(d.Id())
		},
	}.ToResource()
}
