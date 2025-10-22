package scim

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var nonAlphanumeric = regexp.MustCompile(`\W`)

// DataSourceCurrentUser returns information about caller identity
func DataSourceCurrentUser() common.Resource {
	return common.Resource{
		Schema: map[string]*schema.Schema{
			"user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"home": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"repos": {
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
			"workspace_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"acl_principal_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := common.WorkspaceClientUnifiedProvider(ctx, d, c)
			if err != nil {
				return err
			}
			me, err := w.CurrentUser.Me(ctx)
			if err != nil {
				return err
			}
			d.Set("user_name", me.UserName)
			d.Set("home", fmt.Sprintf("/Users/%s", me.UserName))
			d.Set("repos", fmt.Sprintf("/Repos/%s", me.UserName))
			if common.StringIsUUID(me.UserName) {
				d.Set("acl_principal_id", fmt.Sprintf("servicePrincipals/%s", me.UserName))
			} else {
				d.Set("acl_principal_id", fmt.Sprintf("users/%s", me.UserName))
			}
			d.Set("external_id", me.ExternalId)
			splits := strings.Split(me.UserName, "@")
			norm := nonAlphanumeric.ReplaceAllLiteralString(splits[0], "_")
			norm = strings.ToLower(norm)
			d.Set("alphanumeric", norm)
			d.Set("workspace_url", w.Config.Host)
			d.SetId(me.Id)
			return nil
		},
	}
}
