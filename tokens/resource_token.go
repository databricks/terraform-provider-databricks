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

// ResourceToken refreshes token in case it's expired
func ResourceToken() common.Resource {
	s := map[string]*schema.Schema{
		"lifetime_seconds": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},
		"comment": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"scopes": {
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"token_value": {
			Type:      schema.TypeString,
			Computed:  true,
			Sensitive: true,
		},
		"creation_time": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"expiry_time": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"token_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
	common.AddNamespaceInSchema(s)
	common.NamespaceCustomizeSchemaMap(s)
	return common.Resource{
		Schema: s,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, c *common.DatabricksClient) error {
			return common.NamespaceCustomizeDiff(ctx, d, c)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			request := settings.CreateTokenRequest{
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
			tokenResp, err := w.Tokens.Create(ctx, request)
			if err != nil {
				return err
			}
			d.SetId(tokenResp.TokenInfo.TokenId)
			return d.Set("token_value", tokenResp.TokenValue)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			tokenList, err := w.Tokens.ListAll(ctx)
			if err != nil {
				return err
			}
			var tokenInfo *settings.PublicTokenInfo
			for _, ti := range tokenList {
				if ti.TokenId == d.Id() {
					tokenInfo = &ti
					break
				}
			}
			if tokenInfo == nil {
				log.Printf("[INFO] token with id %s not found, recreating it", d.Id())
				d.SetId("")
				return nil
			}
			// we need to set the scopes back to the resource data
			scopes := d.Get("scopes")
			d.Set("token_id", tokenInfo.TokenId)
			d.Set("comment", tokenInfo.Comment)
			d.Set("creation_time", tokenInfo.CreationTime)
			d.Set("expiry_time", tokenInfo.ExpiryTime)
			d.Set("scopes", scopes)
			if tokenInfo.ExpiryTime > 0 && time.Now().UnixMilli() > tokenInfo.ExpiryTime {
				log.Printf("[INFO] token with id %s is expired, recreating it", d.Id())
				d.SetId("")
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			err = w.Tokens.Delete(ctx, settings.RevokeTokenRequest{
				TokenId: d.Id(),
			})
			if apierr.IsMissing(err) {
				return nil
			}
			return err
		},
	}
}
