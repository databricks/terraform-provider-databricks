package identity

import (
	"context"
	"log"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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
	readContext := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		user, err := NewUsersAPI(ctx, m).ReadR(d.Id())
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		if err != nil {
			return diag.FromErr(err)
		}
		err = internal.StructToData(user, userSchema, d)
		if err != nil {
			return diag.FromErr(err)
		}
		return nil
	}
	return &schema.Resource{
		Schema:      userSchema,
		ReadContext: readContext,
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			var ru UserEntity
			err := internal.DataToStructPointer(d, userSchema, &ru)
			if err != nil {
				return diag.FromErr(err)
			}
			user, err := NewUsersAPI(ctx, m).CreateR(ru)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(user.ID)
			return readContext(ctx, d, m)
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			var ru UserEntity
			err := internal.DataToStructPointer(d, userSchema, &ru)
			if err != nil {
				return diag.FromErr(err)
			}
			err = NewUsersAPI(ctx, m).UpdateR(d.Id(), ru)
			if err != nil {
				return diag.FromErr(err)
			}
			return readContext(ctx, d, m)
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			err := NewUsersAPI(ctx, m).Delete(d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
			return nil
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
