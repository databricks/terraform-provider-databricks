package identity

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/databrickslabs/databricks-terraform/common"
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
func NewTokensAPI(m interface{}) TokensAPI {
	return TokensAPI{m.(*common.DatabricksClient), context.TODO()}
}

// TokensAPI exposes the Secrets API
type TokensAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates a api token given a expiration duration and a comment
func (a TokensAPI) Create(tokenLifetime time.Duration, comment string) (r TokenResponse, err error) {
	err = a.client.Post(a.context, "/token/create", TokenRequest{
		LifetimeSeconds: int32(tokenLifetime.Seconds()),
		Comment:         comment,
	}, &r)
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
	return tokenInfo, common.APIError{
		ErrorCode:  "NOT_FOUND",
		Message:    fmt.Sprintf("Unable to locate token: %s", tokenID),
		Resource:   "/api/2.0/token/list",
		StatusCode: http.StatusNotFound,
	}
}

// Delete will delete the token given a token id
func (a TokensAPI) Delete(tokenID string) error {
	return a.client.Post(a.context, "/token/delete", map[string]string{
		"token_id": tokenID,
	}, nil)
}

func ResourceToken() *schema.Resource {
	return &schema.Resource{
		Create: resourceTokenCreate,
		Read:   resourceTokenRead,
		Delete: resourceTokenDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"lifetime_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Default:  0,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "",
			},
			"creation_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"token_value": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"expiry_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceTokenCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	lifeTimeSeconds := d.Get("lifetime_seconds").(int)
	comment := d.Get("comment").(string)
	tokenDuration := time.Duration(lifeTimeSeconds) * time.Second
	tokenResp, err := NewTokensAPI(client).Create(tokenDuration, comment)
	if err != nil {
		return err
	}
	d.SetId(tokenResp.TokenInfo.TokenID)
	err = d.Set("token_value", tokenResp.TokenValue)
	if err != nil {
		return err
	}
	// TODO: (loprio) check if we can omit read
	return resourceTokenRead(d, m)
}

func resourceTokenRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)
	token, err := NewTokensAPI(client).Read(id)
	if err != nil {
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		return err
	}
	err = d.Set("creation_time", token.CreationTime)
	if err != nil {
		return err
	}
	err = d.Set("comment", token.Comment)
	if err != nil {
		return err
	}
	err = d.Set("expiry_time", token.ExpiryTime)
	return err
}

func resourceTokenDelete(d *schema.ResourceData, m interface{}) error {
	tokenID := d.Id()
	client := m.(*common.DatabricksClient)
	err := NewTokensAPI(client).Delete(tokenID)
	return err
}
