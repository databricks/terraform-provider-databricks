package secrets

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/databricks/databricks-sdk-go/service/workspace"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var req workspace.PutAcl
			common.DataToStructPointer(d, s, &req)
			
			// Default timeout: 5 minutes (allows for ~3 retries with 90-second intervals)
			timeout := 5 * time.Minute
			
			// Allow tests to override timeout for faster execution
			if testInterval := os.Getenv("DATABRICKS_SECRET_ACL_TEST_RETRY_INTERVAL_MS"); testInterval != "" {
				if _, err := strconv.Atoi(testInterval); err == nil {
					// For tests, use a timeout that accommodates retry package's backoff (minimum ~500ms per retry)
					timeout = 2 * time.Second
				}
			}
			
			err = retry.RetryContext(ctx, timeout, func() *retry.RetryError {
				// Attempt to create the ACL
				err = w.Secrets.PutAcl(ctx, req)
				if err != nil {
					return retry.NonRetryableError(fmt.Errorf("failed to create Secret ACL: %w", err))
				}
				
				// Verify the ACL was created by reading it back
				secretACL, readErr := w.Secrets.GetAcl(ctx, workspace.GetAclRequest{
					Scope:     req.Scope,
					Principal: req.Principal,
				})
				
				if readErr != nil {
					return retry.RetryableError(fmt.Errorf("secret ACL creation could not be verified: %w", readErr))
				}
				
				// Verify the permission matches what was requested
				if secretACL.Permission.String() != req.Permission.String() {
					return retry.RetryableError(fmt.Errorf("secret ACL permission mismatch: expected %s, got %s", req.Permission.String(), secretACL.Permission.String()))
				}
				
				// Success!
				return nil
			})
			
			if err != nil {
				return err
			}
			
			// Set the resource ID only after successful creation and verification
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
			err = w.Secrets.DeleteAcl(ctx, workspace.DeleteAcl{
				Scope:     scope,
				Principal: principal,
			})
			return common.IgnoreNotFoundError(err)
		},
	}
}
