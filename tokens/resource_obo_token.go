package tokens

import (
	"context"
	"log"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// OboTokenFields defines the schema fields for OBO token creation.
type OboTokenFields struct {
	ApplicationID   string   `json:"application_id" tf:"force_new"`
	LifetimeSeconds int64    `json:"lifetime_seconds,omitempty" tf:"force_new"`
	Comment         string   `json:"comment,omitempty" tf:"force_new"`
	Scopes          []string `json:"scopes,omitempty" tf:"force_new"`
}

func ResourceOboToken() common.Resource {
	oboTokenSchema := common.StructToSchema(OboTokenFields{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["token_value"] = &schema.Schema{
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			}
			return m
		})
	common.AddNamespaceInSchema(oboTokenSchema)
	common.NamespaceCustomizeSchemaMap(oboTokenSchema)
	return common.Resource{
		Schema: oboTokenSchema,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, c *common.DatabricksClient) error {
			return common.NamespaceCustomizeDiff(ctx, d, c)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			request := settings.CreateOboTokenRequest{
				ApplicationId:   d.Get("application_id").(string),
				LifetimeSeconds: int64(d.Get("lifetime_seconds").(int)),
				Comment:         d.Get("comment").(string),
			}
			scopesRaw := d.Get("scopes").([]any)
			if len(scopesRaw) > 0 {
				scopes := make([]string, len(scopesRaw))
				for i, v := range scopesRaw {
					scopes[i] = v.(string)
				}
				request.Scopes = scopes
			}
			ot, err := w.TokenManagement.CreateOboToken(ctx, request)
			if err != nil {
				return err
			}
			d.SetId(ot.TokenInfo.TokenId)
			return d.Set("token_value", ot.TokenValue)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			ot, err := w.TokenManagement.Get(ctx, settings.GetTokenManagementRequest{
				TokenId: d.Id(),
			})
			if err != nil {
				if apierr.IsMissing(err) {
					log.Printf("[INFO] OBO token with id %s not found, recreating it", d.Id())
					d.SetId("")
					return nil
				}
				return err
			}
			if ot.TokenInfo.ExpiryTime > 0 && time.Now().UnixMilli() > ot.TokenInfo.ExpiryTime {
				log.Printf("[INFO] OBO token with id %s is expired, recreating it", d.Id())
				d.SetId("")
				return nil
			}
			d.Set("comment", ot.TokenInfo.Comment)
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			err = w.TokenManagement.Delete(ctx, settings.DeleteTokenManagementRequest{
				TokenId: d.Id(),
			})
			if apierr.IsMissing(err) {
				return nil
			}
			return err
		},
	}
}
