package access

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/databrickslabs/databricks-terraform/common"
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
func NewSecretsAPI(m interface{}) SecretsAPI {
	return SecretsAPI{
		client:  m.(*common.DatabricksClient),
		context: context.TODO(),
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
	return &schema.Resource{
		Create: resourceSecretCreate,
		Read:   resourceSecretRead,
		Delete: resourceSecretDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
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
	}
}

func getSecretID(scope string, key string) (string, error) {
	return scope + "|||" + key, nil
}

func getScopeAndKeyFromSecretID(secretIDString string) (string, string, error) {
	split := strings.Split(secretIDString, "|||")
	if len(split) != 2 {
		return "", "", fmt.Errorf("Malformed secret id: %s", secretIDString)
	}
	return strings.Split(secretIDString, "|||")[0], strings.Split(secretIDString, "|||")[1], nil
}

func resourceSecretCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	scopeName := d.Get("scope").(string)
	key := d.Get("key").(string)
	secretValue := d.Get("string_value").(string)
	err := NewSecretsAPI(client).Create(secretValue, scopeName, key)
	if err != nil {
		return err
	}
	id, err := getSecretID(scopeName, key)
	if err != nil {
		return err
	}
	d.SetId(id)
	return resourceSecretRead(d, m)
}

func resourceSecretRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)
	scope, key, err := getScopeAndKeyFromSecretID(id)
	if err != nil {
		return err
	}
	secretMetaData, err := NewSecretsAPI(client).Read(scope, key)
	if err != nil {
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		return err
	}

	err = d.Set("last_updated_timestamp", secretMetaData.LastUpdatedTimestamp)
	if err != nil {
		return err
	}

	err = d.Set("scope", scope)
	if err != nil {
		return err
	}
	err = d.Set("key", secretMetaData.Key)
	if err != nil {
		return err
	}
	d.SetId(id)
	return nil
}

func resourceSecretDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	id := d.Id()
	scope, key, err := getScopeAndKeyFromSecretID(id)
	if err != nil {
		return err
	}
	err = NewSecretsAPI(client).Delete(scope, key)
	return err
}
