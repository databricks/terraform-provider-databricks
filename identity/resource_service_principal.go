package identity

import (
	"context"
	"log"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceServicePrincipal manages service principals within workspace
func ResourceServicePrincipal() *schema.Resource {
	servicePrincipalSchema := internal.StructToSchema(ServicePrincipalEntity{}, func(
		s map[string]*schema.Schema) map[string]*schema.Schema {
		s["application_id"].ForceNew = true
		s["active"].Default = true
		return s
	})
	readContext := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		servicePrincipal, err := NewServicePrincipalsAPI(m).ReadR(d.Id())
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		if err != nil {
			return diag.FromErr(err)
		}
		err = internal.StructToData(servicePrincipal, servicePrincipalSchema, d)
		if err != nil {
			return diag.FromErr(err)
		}
		return nil
	}
	return &schema.Resource{
		Schema:      servicePrincipalSchema,
		ReadContext: readContext,
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			var ru ServicePrincipalEntity
			err := internal.DataToStructPointer(d, servicePrincipalSchema, &ru)
			if err != nil {
				return diag.FromErr(err)
			}
			servicePrincipal, err := NewServicePrincipalsAPI(m).CreateR(ru)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(servicePrincipal.ID)
			return readContext(ctx, d, m)
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			var ru ServicePrincipalEntity
			err := internal.DataToStructPointer(d, servicePrincipalSchema, &ru)
			if err != nil {
				return diag.FromErr(err)
			}
			err = NewServicePrincipalsAPI(m).UpdateR(d.Id(), ru)
			if err != nil {
				return diag.FromErr(err)
			}
			return readContext(ctx, d, m)
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			err := NewServicePrincipalsAPI(m).Delete(d.Id())
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
