package secrets

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// defaultTimeout is the default timeout for the robustPutACL function.
var defaultTimeout = 5 * time.Minute

// ResourceSecretACL manages access to secret scopes
func ResourceSecretACL() common.Resource {
	p := common.NewPairSeparatedID("scope", "principal", "|||")
	s := map[string]*schema.Schema{
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
	}
	return common.Resource{
		Schema: s,
		CanSkipReadAfterCreateAndUpdate: func(_ *schema.ResourceData) bool {
			return true
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := common.WorkspaceClientUnifiedProvider(ctx, d, c)
			if err != nil {
				return err
			}
			var req workspace.PutAcl
			common.DataToStructPointer(d, s, &req)
			if err := robustPutACL(w.Secrets, ctx, req, defaultTimeout); err != nil {
				return err
			}
			p.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			scope, principal, err := p.Unpack(d)
			if err != nil {
				return err
			}
			w, err := common.WorkspaceClientUnifiedProvider(ctx, d, c)
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
			w, err := common.WorkspaceClientUnifiedProvider(ctx, d, c)
			if err != nil {
				return err
			}
			err = w.Secrets.DeleteAcl(ctx, workspace.DeleteAcl{
				Scope:     scope,
				Principal: principal,
			})
			return common.IgnoreNotFoundError(err)
		},
	}
}

// secretsClient is an interface for testing.
type secretsClient interface {
	PutAcl(context.Context, workspace.PutAcl) error
	GetAcl(context.Context, workspace.GetAclRequest) (*workspace.AclItem, error)
}

// robustPutACL creates or overwrites the ACL associated with the given
// principal.
//
// The function is retried until the ACL is applied with the right
// permission. This is necessary to workaround current limitations due to
// an internal caching mechanism.
//
// See [issue-4195] for reference.
//
// [issue-4195]: https://github.com/databricks/terraform-provider-databricks/issues/4195
func robustPutACL(sc secretsClient, ctx context.Context, req workspace.PutAcl, timeout time.Duration) error {
	return retry.RetryContext(ctx, timeout, func() *retry.RetryError {
		if err := sc.PutAcl(ctx, req); err != nil {
			return retry.NonRetryableError(fmt.Errorf("failed to create Secret ACL: %w", err))
		}

		// Verify that the ACL was properly applied with the right permissions.
		// If not, retry the operation until the ACL is applied correctly.
		secretACL, err := sc.GetAcl(ctx, workspace.GetAclRequest{
			Scope:     req.Scope,
			Principal: req.Principal,
		})
		if err != nil {
			return retry.RetryableError(fmt.Errorf("secret ACL creation could not be verified: %w", err))
		}
		if secretACL.Permission.String() != req.Permission.String() {
			return retry.RetryableError(fmt.Errorf("secret ACL permission mismatch: expected %s, got %s", req.Permission.String(), secretACL.Permission.String()))
		}

		return nil
	})
}
