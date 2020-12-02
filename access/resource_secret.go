package access

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// SecretsRequest ...
type SecretsRequest struct {
	StringValue string `json:"string_value,omitempty" mask:"true"`
	Scope       string `json:"scope,omitempty"`
	Key         string `json:"key,omitempty"`
}

// SecretsList ...
type SecretsList struct {
	Secrets []SecretMetadata `json:"secrets,omitempty"`
}

// SecretMetadata is a struct that encapsulates the metadata for a secret object in a scope
type SecretMetadata struct {
	Key                  string `json:"key,omitempty"`
	LastUpdatedTimestamp int64  `json:"last_updated_timestamp,omitempty"`
}

// ACLPermission is a custom type for acl permissions
type ACLPermission string

// List of possible ACL Permissions on Databricks
const (
	ACLPermissionRead   ACLPermission = "READ"
	ACLPermissionWrite  ACLPermission = "WRITE"
	ACLPermissionManage ACLPermission = "MANAGE"
)

// ACLItem is a struct that contains information about a secret scope acl
type ACLItem struct {
	Principal  string        `json:"principal,omitempty"`
	Permission ACLPermission `json:"permission,omitempty"`
}

// SecretACLRequest generic request for secret acls
type SecretACLRequest struct {
	Scope      string        `json:"scope,omitempty" url:"scope,omitempty"`
	Principal  string        `json:"principal,omitempty" url:"principal,omitempty"`
	Permission ACLPermission `json:"permission,omitempty" url:"permission,omitempty"`
}

// NewSecretsAPI creates SecretsAPI instance from provider meta
func NewSecretsAPI(ctx context.Context, m interface{}) SecretsAPI {
	return SecretsAPI{
		client:  m.(*common.DatabricksClient),
		context: ctx,
	}
}

// SecretsAPI exposes the Secrets API
type SecretsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates or modifies a string secret depends on the type of scope backend
func (a SecretsAPI) Create(stringValue, scope, key string) error {
	return a.client.Post(a.context, "/secrets/put", SecretsRequest{
		StringValue: stringValue,
		Scope:       scope,
		Key:         key,
	}, nil)
}

// Delete deletes a secret depends on the type of scope backend
func (a SecretsAPI) Delete(scope, key string) error {
	return a.client.Post(a.context, "/secrets/delete", SecretsRequest{
		Scope: scope,
		Key:   key,
	}, nil)
}

// List lists the secret keys that are stored at this scope
func (a SecretsAPI) List(scope string) ([]SecretMetadata, error) {
	var secretsList SecretsList
	err := a.client.Get(a.context, "/secrets/list", map[string]string{
		"scope": scope,
	}, &secretsList)
	return secretsList.Secrets, err
}

// Read returns the metadata for the secret and not the contents of the secret
func (a SecretsAPI) Read(scope string, key string) (SecretMetadata, error) {
	var secretMeta SecretMetadata
	secrets, err := a.List(scope)
	if err != nil {
		return secretMeta, err
	}
	for _, secret := range secrets {
		if secret.Key == key {
			return secret, nil
		}
	}
	return secretMeta, common.APIError{
		ErrorCode:  "NOT_FOUND",
		Message:    fmt.Sprintf("no secret Scope found with secret metadata scope name: %s and key: %s", scope, key),
		Resource:   "/api/2.0/secrets/scopes/list",
		StatusCode: http.StatusNotFound,
	}
}

// ResourceSecret manages secrets
func ResourceSecret() *schema.Resource {
	p := util.NewPairID("scope", "key")
	return util.CommonResource{
		Schema: map[string]*schema.Schema{
			"string_value": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"scope": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"last_updated_timestamp": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			if err := NewSecretsAPI(ctx, c).Create(d.Get("string_value").(string), d.Get("scope").(string),
				d.Get("key").(string)); err != nil {
				return err
			}
			p.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			scope, key, err := p.Unpack(d)
			if err != nil {
				return err
			}
			m, err := NewSecretsAPI(ctx, c).Read(scope, key)
			if err != nil {
				return err
			}
			return d.Set("last_updated_timestamp", m.LastUpdatedTimestamp)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			scope, key, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return NewSecretsAPI(ctx, c).Delete(scope, key)
		},
	}.ToResource()
}
