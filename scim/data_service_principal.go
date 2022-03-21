package scim

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceUser returns information about user specified by user name
func DataSourceServicePrincipal() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"application_id": {
				Type:     schema.TypeString,
				Computed: true
			},
			"home": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"repos": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			}
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			spnAPI := NewServicePrincipalsAPI(ctx, m)
			spn, err := spnAPI.read(d.Get("application_id").(string))
			if err != nil {
				return diag.FromErr(err)
			}
			d.Set("display_name", spn.DisplayName)
			d.Set("home", fmt.Sprintf("/Users/%s", spn.UserName))
			d.Set("repos", fmt.Sprintf("/Repos/%s", spn.UserName))
			d.Set("external_id", spn.ExternalID)
			d.Set("application_id", spn.ApplicationID)
			d.SetId(spn.ID)
			return nil
		},
	}
}