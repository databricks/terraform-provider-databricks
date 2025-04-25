package oauth

import (
	"context"
	"github.com/databricks/databricks-sdk-go/service/oauth2"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func ResourceServicePrincipalFederationPolicy() common.Resource {
	s := map[string]*schema.Schema{
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"service_principal_id": {
			Type:     schema.TypeInt,
			Required: true,
			ForceNew: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
			//Computed: true,
		},
		"uid": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"oidc_policy": {
			Type:     schema.TypeSet,
			Required: true,
			MinItems: 1,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"audiences": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"issuer": {
						Type:     schema.TypeString,
						Required: true,
					},
					"jwks_json": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"jwks_uri": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"subject": {
						Type:     schema.TypeString,
						Required: true,
					},
					"subject_claim": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
	}
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var federationPolicy oauth2.FederationPolicy
			common.DataToStructPointer(d, s, &federationPolicy)
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			spfp, err := acc.ServicePrincipalFederationPolicy.Create(ctx,
				oauth2.CreateServicePrincipalFederationPolicyRequest{
					Policy:             &federationPolicy,
					ServicePrincipalId: int64(d.Get("service_principal_id").(int)),
				})
			if err != nil {
				return err
			}
			name := spfp.Name
			d.SetId(name[strings.LastIndex(name, "/")+1:])
			return common.StructToData(spfp, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			return acc.ServicePrincipalFederationPolicy.DeleteByServicePrincipalIdAndPolicyId(
				ctx,
				int64(d.Get("service_principal_id").(int)),
				d.Id(),
			)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			spfp, err := acc.ServicePrincipalFederationPolicy.GetByServicePrincipalIdAndPolicyId(ctx,
				int64(d.Get("service_principal_id").(int)), d.Id())
			if err != nil {
				err = common.IgnoreNotFoundError(err)
				if err != nil {
					return err
				}
				log.Printf("[INFO] service principal federation policy with id %s not found, recreating it", d.Id())
				d.SetId("")
				return nil
			}
			return common.StructToData(spfp, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			var federationPolicy oauth2.FederationPolicy
			common.DataToStructPointer(d, s, &federationPolicy)
			spfp, err := acc.ServicePrincipalFederationPolicy.Update(ctx,
				oauth2.UpdateServicePrincipalFederationPolicyRequest{
					Policy:             &federationPolicy,
					PolicyId:           d.Id(),
					ServicePrincipalId: int64(d.Get("service_principal_id").(int)),
				},
			)
			return common.StructToData(spfp, s, d)
		},
	}

}
