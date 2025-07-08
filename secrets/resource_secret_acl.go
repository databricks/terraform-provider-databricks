package secrets

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/databricks/databricks-sdk-go/service/workspace"

	"github.com/databricks/terraform-provider-databricks/common"

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
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var req workspace.PutAcl
			common.DataToStructPointer(d, s, &req)
			
			// Retry logic: up to 3 attempts with 90-second intervals
			const maxRetries = 3
			retryInterval := 90 * time.Second
			
			// Allow tests to override retry interval
			if testInterval := os.Getenv("DATABRICKS_SECRET_ACL_TEST_RETRY_INTERVAL_MS"); testInterval != "" {
				if ms, err := strconv.Atoi(testInterval); err == nil {
					retryInterval = time.Duration(ms) * time.Millisecond
				}
			}
			
			for attempt := 1; attempt <= maxRetries; attempt++ {
				// Attempt to create the ACL
				err = w.Secrets.PutAcl(ctx, req)
				if err != nil {
					if attempt == maxRetries {
						return fmt.Errorf("failed to create Secret ACL after %d attempts: %w", maxRetries, err)
					}
					time.Sleep(retryInterval)
					continue
				}
				
				// Verify the ACL was created by reading it back
				secretACL, readErr := w.Secrets.GetAcl(ctx, workspace.GetAclRequest{
					Scope:     req.Scope,
					Principal: req.Principal,
				})
				
				if readErr != nil {
					if attempt == maxRetries {
						return fmt.Errorf("secret ACL creation could not be verified after %d attempts: %w", maxRetries, readErr)
					}
					time.Sleep(retryInterval)
					continue
				}
				
				// Verify the permission matches what was requested
				if secretACL.Permission.String() != req.Permission.String() {
					if attempt == maxRetries {
						return fmt.Errorf("secret ACL permission mismatch after %d attempts: expected %s, got %s", maxRetries, req.Permission.String(), secretACL.Permission.String())
					}
					time.Sleep(retryInterval)
					continue
				}
				
				// Success! Set the permission in the state and resource ID
				p.Pack(d)
				return nil
			}
			
			// This should never be reached, but just in case
			return fmt.Errorf("unexpected error: exhausted all retry attempts")
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
