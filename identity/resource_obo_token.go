package identity

import (
	"context"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"
)

type OboToken struct {
	ApplicationID   string `json:"application_id"`
	LifetimeSeconds int32  `json:"lifetime_seconds"`
	Comment         string `json:"comment"`
}

func NewTokenManagementAPI(ctx context.Context, m interface{}) TokenManagementAPI {
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
	return a.client.Delete(a.context, fmt.Sprintf("/token-management/tokens/%s", tokenID), map[string]interface{}{})
}

func (a TokenManagementAPI) Read(tokenID string) (ti TokenResponse, err error) {
	err = a.client.Get(a.context, fmt.Sprintf("/token-management/tokens/%s", tokenID), nil, &ti)
	return
}
