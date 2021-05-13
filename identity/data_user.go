package identity

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceUser returns information about user specified by user name
func DataSourceUser() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"user_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"home": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"alphanumeric": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			usersAPI := NewUsersAPI(ctx, m)
			userList, err := usersAPI.Filter(fmt.Sprintf("userName eq '%s'", d.Get("user_name")))
			if err != nil {
				return diag.FromErr(err)
			}
			if len(userList) == 0 {
				return diag.FromErr(fmt.Errorf("cannot find user %s", d.Get("user_name")))
			}
			d.Set("user_name", userList[0].UserName)
			d.Set("home", fmt.Sprintf("/Users/%s", userList[0].UserName))
			splits := strings.Split(userList[0].UserName, "@")
			norm := nonAlphanumeric.ReplaceAllLiteralString(splits[0], "_")
			norm = strings.ToLower(norm)
			d.Set("alphanumeric", norm)
			d.SetId(userList[0].ID)
			return nil
		},
	}
}
