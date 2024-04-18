package tokens

import (
	"context"
	"fmt"

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
	return a.client.Delete(a.context, fmt.Sprintf("/token-management/tokens/%s", tokenID), map[string]any{})
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
				return err
			}
			// this method is just a shim to check if token does still exist
			return d.Set("comment", ot.TokenInfo.Comment)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewTokenManagementAPI(ctx, c).Delete(d.Id())
		},
	}
}
