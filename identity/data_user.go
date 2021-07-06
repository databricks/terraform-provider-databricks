package identity

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getUser(usersAPI UsersAPI, id, name string) (user ScimUser, err error) {
	if id != "" {
		return usersAPI.read(id)
	}
	userList, err := usersAPI.Filter(fmt.Sprintf("userName eq '%s'", name))
	if err != nil {
		return
	}
	if len(userList) == 0 {
		err = fmt.Errorf("cannot find user %s", name)
		return
	}
	user = userList[0]
	return
}

// DataSourceUser returns information about user specified by user name
func DataSourceUser() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"user_name": {
				Type:         schema.TypeString,
				ExactlyOneOf: []string{"user_name", "user_id"},
				Optional:     true,
			},
			"user_id": {
				Type:         schema.TypeString,
				ExactlyOneOf: []string{"user_name", "user_id"},
				Optional:     true,
			},
			"home": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
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
			user, err := getUser(usersAPI, d.Get("user_id").(string), d.Get("user_name").(string))
			if err != nil {
				return diag.FromErr(err)
			}
			d.Set("user_name", user.UserName)
			d.Set("display_name", user.DisplayName)
			d.Set("home", fmt.Sprintf("/Users/%s", user.UserName))
			splits := strings.Split(user.UserName, "@")
			norm := nonAlphanumeric.ReplaceAllLiteralString(splits[0], "_")
			norm = strings.ToLower(norm)
			d.Set("alphanumeric", norm)
			d.SetId(user.ID)
			return nil
		},
	}
}
