package tokens

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TokenRequest asks for a token
type TokenRequest struct {
	LifetimeSeconds int32    `json:"lifetime_seconds,omitempty"`
	Comment         string   `json:"comment,omitempty"`
	Scopes          []string `json:"scopes,omitempty"`
}

// TokenResponse is a struct that contains information about token that is created from the create tokens api
type TokenResponse struct {
	TokenValue string     `json:"token_value,omitempty"`
	TokenInfo  *TokenInfo `json:"token_info,omitempty"`
}

// TokenInfo is a struct that contains metadata about a given token
type TokenInfo struct {
	TokenID      string `json:"token_id,omitempty"`
	CreationTime int64  `json:"creation_time,omitempty"`
	ExpiryTime   int64  `json:"expiry_time,omitempty"`
	Comment      string `json:"comment,omitempty"`
}

// TokenList ...
type TokenList struct {
	TokenInfos []TokenInfo `json:"token_infos,omitempty"`
}

// NewTokensAPI creates TokensAPI instance from provider meta
func NewTokensAPI(ctx context.Context, m any) TokensAPI {
	return TokensAPI{m.(*common.DatabricksClient), ctx}
}

// TokensAPI exposes the Secrets API
type TokensAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates a api token given a token request
func (a TokensAPI) Create(request TokenRequest) (r TokenResponse, err error) {
	err = a.client.Post(a.context, "/token/create", request, &r, a.client.AddWorkspaceIdHeader)
	return
}

// List will list all the token metadata and not the content of the tokens in the workspace
func (a TokensAPI) List() ([]TokenInfo, error) {
	var tokenListResult TokenList
	err := a.client.Get(a.context, "/token/list", nil, &tokenListResult, a.client.AddWorkspaceIdHeader)
	return tokenListResult.TokenInfos, err
}

// Read will return the token metadata and not the content of the token
func (a TokensAPI) Read(tokenID string) (TokenInfo, error) {
	var tokenInfo TokenInfo
	tokenList, err := a.List()
	if err != nil {
		return tokenInfo, err
	}
	for _, tokenInfoRecord := range tokenList {
		if tokenInfoRecord.TokenID == tokenID {
			return tokenInfoRecord, nil
		}
	}
	return tokenInfo, &apierr.APIError{
		ErrorCode:  "NOT_FOUND",
		StatusCode: 404,
		Message:    fmt.Sprintf("Unable to locate token: %s", tokenID),
	}
}

// Delete will delete the token given a token id
func (a TokensAPI) Delete(tokenID string) error {
	err := a.client.Post(a.context, "/token/delete", map[string]string{
		"token_id": tokenID,
	}, nil, a.client.AddWorkspaceIdHeader)
	return common.IgnoreNotFoundError(err) // ignore not found error on delete, as it is idempotent
}

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
			newClient, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var request TokenRequest
			common.DataToStructPointer(d, s, &request)
			tokenResp, err := NewTokensAPI(ctx, newClient).Create(request)
			if err != nil {
				return err
			}
			d.SetId(tokenResp.TokenInfo.TokenID)
			return d.Set("token_value", tokenResp.TokenValue)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			newClient, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			tokenInfo, err := NewTokensAPI(ctx, newClient).Read(d.Id())
			if err != nil {
				err = common.IgnoreNotFoundError(err)
				if err != nil {
					return err
				}
				log.Printf("[INFO] token with id %s not found, recreating it", d.Id())
				d.SetId("")
				return nil
			}
			// we need to set the scopes back to the resource data
			scopes := d.Get("scopes")
			err = common.StructToData(tokenInfo, s, d)
			if err != nil {
				return err
			}
			d.Set("scopes", scopes)
			if tokenInfo.ExpiryTime > 0 && time.Now().UnixMilli() > tokenInfo.ExpiryTime {
				log.Printf("[INFO] token with id %s is expired, recreating it", d.Id())
				d.SetId("")
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			newClient, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return NewTokensAPI(ctx, newClient).Delete(d.Id())
		},
	}
}
