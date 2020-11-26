package access

import (
	"context"
	"log"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// NewSecretAclsAPI creates SecretAclsAPI instance from provider meta
func NewSecretAclsAPI(m interface{}) SecretAclsAPI {
	return SecretAclsAPI{m.(*common.DatabricksClient), context.TODO()}
}

// SecretAclsAPI exposes the Secret ACL API
type SecretAclsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates or overwrites the ACL associated with the given principal (user or group) on the specified scope point
func (a SecretAclsAPI) Create(scope string, principal string, permission ACLPermission) error {
	return a.client.Post(a.context, "/secrets/acls/put", SecretACLRequest{
		Scope:      scope,
		Principal:  principal,
		Permission: permission,
	}, nil)
}

// Delete deletes the given ACL on the given scope
func (a SecretAclsAPI) Delete(scope string, principal string) error {
	return a.client.Post(a.context, "/secrets/acls/delete", SecretACLRequest{
		Scope:     scope,
		Principal: principal,
	}, nil)
}

// Read describe the details about the given ACL, such as the group and permission
func (a SecretAclsAPI) Read(scope string, principal string) (ACLItem, error) {
	var aclItem ACLItem
	err := a.client.Get(a.context, "/secrets/acls/get", SecretACLRequest{
		Scope:     scope,
		Principal: principal,
	}, &aclItem)
	return aclItem, err
}

// List lists the ACLs set on the given scope
func (a SecretAclsAPI) List(scope string) ([]ACLItem, error) {
	var aclItem struct {
		Items []ACLItem `json:"items,omitempty"`
	}
	err := a.client.Get(a.context, "/secrets/acls/list", map[string]string{
		"scope": scope,
	}, &aclItem)
	return aclItem.Items, err
}

// ResourceSecretACL manages access to secret scopes
func ResourceSecretACL() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretACLCreate,
		Read:   resourceSecretACLRead,
		Delete: resourceSecretACLDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"scope": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"principal": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"permission": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func getSecretACLID(scope string, key string) (string, error) {
	return scope + "|||" + key, nil
}

func getScopeAndKeyFromSecretACLID(secretACLIDString string) (string, string, error) {
	return getScopeAndKeyFromSecretID(secretACLIDString)
}

func resourceSecretACLCreate(d *schema.ResourceData, m interface{}) error {
	scopeName := d.Get("scope").(string)
	principal := d.Get("principal").(string)
	permission := ACLPermission(d.Get("permission").(string))
	err := NewSecretAclsAPI(m).Create(scopeName, principal, permission)
	if err != nil {
		return err
	}
	id, err := getSecretACLID(scopeName, principal)
	if err != nil {
		return err
	}
	d.SetId(id)
	return resourceSecretACLRead(d, m)
}

func resourceSecretACLRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	scope, principal, err := getScopeAndKeyFromSecretACLID(id)
	if err != nil {
		return err
	}
	secretACL, err := NewSecretAclsAPI(m).Read(scope, principal)
	if err != nil {
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		return err
	}
	err = d.Set("scope", scope)
	if err != nil {
		return err
	}
	err = d.Set("principal", principal)
	if err != nil {
		return err
	}
	err = d.Set("permission", secretACL.Permission)
	return err
}

func resourceSecretACLDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	scope, key, err := getScopeAndKeyFromSecretACLID(id)
	if err != nil {
		return err
	}
	return NewSecretAclsAPI(m).Delete(scope, key)
}
