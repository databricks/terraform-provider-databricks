package tokens

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type OboToken struct {
	ApplicationID   string `json:"application_id"`
	LifetimeSeconds int32  `json:"lifetime_seconds,omitempty"`
	Comment         string `json:"comment,omitempty"`
}

func NewTokenManagementAPI(ctx context.Context, m any) TokenManagementAPI {
	return TokenManagementAPI{m.(*common.DatabricksClient), ctx}
}

type TokenManagementAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func (a TokenManagementAPI) CreateTokenOnBehalfOfServicePrincipal(request OboToken) (token TokenResponse, err error) {
	err = a.client.Post(a.context, "/token-management/on-behalf-of/tokens", request, &token)
	return
}

func (a TokenManagementAPI) Delete(tokenID string) error {
	err := a.client.Delete(a.context, fmt.Sprintf("/token-management/tokens/%s", tokenID), map[string]any{})
	return common.IgnoreNotFoundError(err) // ignore not found error on delete, as it is idempotent
}

func (a TokenManagementAPI) Read(tokenID string) (ti TokenResponse, err error) {
	err = a.client.Get(a.context, fmt.Sprintf("/token-management/tokens/%s", tokenID), nil, &ti)
	return
}

func ResourceOboToken() common.Resource {
	oboTokenSchema := common.StructToSchema(OboToken{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["token_value"] = &schema.Schema{
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			}
			return m
		})
	return common.Resource{
		Schema: oboTokenSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var request OboToken
			common.DataToStructPointer(d, oboTokenSchema, &request)
			ot, err := NewTokenManagementAPI(ctx, c).CreateTokenOnBehalfOfServicePrincipal(request)
			if err != nil {
				return err
			}
			d.SetId(ot.TokenInfo.TokenID)
			return d.Set("token_value", ot.TokenValue)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			ot, err := NewTokenManagementAPI(ctx, c).Read(d.Id())
			if err != nil {
				err = common.IgnoreNotFoundError(err)
				if err != nil {
					return err
				}
				log.Printf("[INFO] OBO token with id %s not found, recreating it", d.Id())
				d.SetId("")
			} else { // check if token is expired
				if ot.TokenInfo.ExpiryTime > 0 && time.Now().UnixMilli() > ot.TokenInfo.ExpiryTime {
					log.Printf("[INFO] OBO token with id %s is expired, recreating it", d.Id())
					d.SetId("")
				}
			}
			if d.Id() != "" {
				// set comment only if token exists
				d.Set("comment", ot.TokenInfo.Comment)
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewTokenManagementAPI(ctx, c).Delete(d.Id())
		},
	}
}
