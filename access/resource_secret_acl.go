package access

import (
	"context"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// NewSecretAclsAPI creates SecretAclsAPI instance from provider meta
func NewSecretAclsAPI(ctx context.Context, m interface{}) SecretAclsAPI {
	return SecretAclsAPI{m.(*common.DatabricksClient), ctx}
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
	p := util.NewPairSeparatedID("scope", "principal", "|||")
	return util.CommonResource{
		Schema: map[string]*schema.Schema{
			"scope": {
				Type:         schema.TypeString,
				ValidateFunc: validScope,
				Required:     true,
				ForceNew:     true,
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
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			if err := NewSecretAclsAPI(ctx, c).Create(
				d.Get("scope").(string), d.Get("principal").(string),
				ACLPermission(d.Get("permission").(string))); err != nil {
				return err
			}
			// TODO: check what happens if ID is set before error happens in create
			p.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			scope, principal, err := p.Unpack(d)
			if err != nil {
				return err
			}
			secretACL, err := NewSecretAclsAPI(ctx, c).Read(scope, principal)
			if err != nil {
				return err
			}
			return d.Set("permission", secretACL.Permission)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			scope, principal, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return NewSecretAclsAPI(ctx, c).Delete(scope, principal)
		},
	}.ToResource()
}
