package identity

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceMe returns information about caller identity
func DataSourceMe() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"home": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			me, err := NewUsersAPI(ctx, m).Me()
			if err != nil {
				return diag.FromErr(err)
			}
			d.Set("user_name", me.UserName)
			d.Set("home", fmt.Sprintf("/Users/%s", me.UserName))
			d.SetId(me.ID)
			return nil
		},
	}
}
