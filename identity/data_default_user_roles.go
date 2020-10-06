package identity

import (
	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceDefaultUserRoles() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "Data source `databricks_default_user_roles` is no longer supported and would be removed in version 0.3",
		Read: func(d *schema.ResourceData, m interface{}) error {
			client := m.(*common.DatabricksClient)

			defaultRolesUserName := d.Get("default_username").(string)
			metaUser, err := NewUsersAPI(client).GetOrCreateDefaultMetaUser(defaultRolesUserName, defaultRolesUserName, true)
			if err != nil {
				return err
			}
			d.SetId(metaUser.ID)
			err = d.Set("default_username", metaUser.UserName)
			if err != nil {
				return err
			}

			err = d.Set("roles", getListOfRoles(metaUser.Roles))
			return err
		},
		Schema: map[string]*schema.Schema{
			"default_username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"roles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				ForceNew: true,
			},
		},
	}
}
