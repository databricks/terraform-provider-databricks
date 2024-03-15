package secrets

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/workspace"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceSecretACL manages access to secret scopes
func ResourceSecretACL() common.Resource {
	p := common.NewPairSeparatedID("scope", "principal", "|||")
	return common.Resource{
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
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			err = w.Secrets.PutAcl(ctx, workspace.PutAcl{
				Scope:      d.Get("scope").(string),
				Principal:  d.Get("principal").(string),
				Permission: workspace.AclPermission(d.Get("permission").(string)),
			})
			if err != nil {
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
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			secretACL, err := w.Secrets.GetAcl(ctx, workspace.GetAclRequest{
				Scope:     scope,
				Principal: principal,
			})
			if err != nil {
				return err
			}
			return d.Set("permission", secretACL.Permission.String())
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			scope, principal, err := p.Unpack(d)
			if err != nil {
				return err
			}
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.Secrets.DeleteAcl(ctx, workspace.DeleteAcl{
				Scope:     scope,
				Principal: principal,
			})
		},
	}
}
