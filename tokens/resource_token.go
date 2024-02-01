package tokens

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TokenRequest asks for a token
type TokenRequest struct {
	LifetimeSeconds int32  `json:"lifetime_seconds,omitempty"`
	Comment         string `json:"comment,omitempty"`
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

// Create creates a api token given a expiration duration and a comment
func (a TokensAPI) Create(tokenLifetime time.Duration, comment string) (r TokenResponse, err error) {
	request := TokenRequest{}
	seconds := int32(tokenLifetime.Seconds())
	if seconds > 0 {
		request.LifetimeSeconds = seconds
	}
	if comment != "" {
		request.Comment = comment
	}

	err = a.client.Post(a.context, "/token/create", request, &r)
	return
}

// List will list all the token metadata and not the content of the tokens in the workspace
func (a TokensAPI) List() ([]TokenInfo, error) {
	var tokenListResult TokenList
	err := a.client.Get(a.context, "/token/list", nil, &tokenListResult)
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
	return tokenInfo, apierr.NotFound(fmt.Sprintf("Unable to locate token: %s", tokenID))
}

// Delete will delete the token given a token id
func (a TokensAPI) Delete(tokenID string) error {
	err := a.client.Post(a.context, "/token/delete", map[string]string{
		"token_id": tokenID,
	}, nil)
	if apierr.IsMissing(err) {
		return nil
	}
	return err
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
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			comment := d.Get("comment").(string)
			lifeTimeSeconds := d.Get("lifetime_seconds").(int)
			tokenDuration := time.Duration(lifeTimeSeconds) * time.Second
			tokenResp, err := NewTokensAPI(ctx, c).Create(tokenDuration, comment)
			if err != nil {
				return err
			}
			d.SetId(tokenResp.TokenInfo.TokenID)
			return d.Set("token_value", tokenResp.TokenValue)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			tokenInfo, err := NewTokensAPI(ctx, c).Read(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(tokenInfo, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewTokensAPI(ctx, c).Delete(d.Id())
		},
	}
}
