package scim

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getUser(usersAPI UsersAPI, id, name string) (user User, err error) {
	if id != "" {
		return usersAPI.Read(id, "userName,displayName,externalId,applicationId")
	}
	userList, err := usersAPI.Filter(fmt.Sprintf(`userName eq "%s"`, name), true)
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
func DataSourceUser() common.Resource {
	return common.Resource{
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
			"repos": {
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
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"application_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"acl_principal_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Read: func(ctx context.Context, d *schema.ResourceData, m *common.DatabricksClient) error {
			usersAPI := NewUsersAPI(ctx, m)
			user, err := getUser(usersAPI, d.Get("user_id").(string), d.Get("user_name").(string))
			if err != nil {
				return err
			}
			d.Set("user_name", user.UserName)
			d.Set("display_name", user.DisplayName)
			d.Set("home", fmt.Sprintf("/Users/%s", user.UserName))
			d.Set("repos", fmt.Sprintf("/Repos/%s", user.UserName))
			d.Set("acl_principal_id", fmt.Sprintf("users/%s", user.UserName))
			d.Set("external_id", user.ExternalID)
			d.Set("application_id", user.ApplicationID)
			splits := strings.Split(user.UserName, "@")
			norm := nonAlphanumeric.ReplaceAllLiteralString(splits[0], "_")
			norm = strings.ToLower(norm)
			d.Set("alphanumeric", norm)
			d.SetId(user.ID)
			return nil
		},
	}
}
