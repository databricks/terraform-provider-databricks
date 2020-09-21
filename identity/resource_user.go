package identity

import (
	"context"
	"log"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Bind entity to data
func Bind(entity interface{}, err error) func(d *schema.ResourceData, s map[string]*schema.Schema) diag.Diagnostics {
	return func(d *schema.ResourceData, s map[string]*schema.Schema) diag.Diagnostics {
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		if err != nil {
			return diag.FromErr(err)
		}
		err = internal.StructToData(entity, s, d)
		if err != nil {
			return diag.FromErr(err)
		}
		return nil
	}
}

// ResourceUser manages users within workspace
func ResourceUser() *schema.Resource {
	userSchema := internal.StructToSchema(UserEntity{}, func(
		s map[string]*schema.Schema) map[string]*schema.Schema {
		s["user_name"].ForceNew = true
		s["active"].Default = true
		return s
	})
	readContext := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		return Bind(NewUsersAPI(m).ReadR(d.Id()))(d, userSchema)
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
			user, err := NewUsersAPI(m).CreateR(ru)
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
			err = NewUsersAPI(m).UpdateR(d.Id(), ru)
			if err != nil {
				return diag.FromErr(err)
			}
			return readContext(ctx, d, m)
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			err := NewUsersAPI(m).Delete(d.Id())
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
