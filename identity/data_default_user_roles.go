package identity

import (
	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceDefaultUserRoles() *schema.Resource {
	return &schema.Resource{
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
