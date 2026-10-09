package mws

import (
	"context"
	"log"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

// After a workspace is deleted, the account API can keep rejecting deletes of the
// network, credentials, storage configuration and customer-managed keys that were
// attached to it for a short time. Terraform destroys those right after the workspace,
// so without a retry `terraform destroy` fails and has to be run again.
var attachedToWorkspaceRegex = regexp.MustCompile(
	`while it is attached to a workspace|is being used by active workspace`)

var deleteWhileAttachedTimeout = 5 * time.Minute

func retryDeleteWhileAttached(ctx context.Context, deleteFunc func() error) error {
	return retry.RetryContext(ctx, deleteWhileAttachedTimeout, func() *retry.RetryError {
		err := deleteFunc()
		if err == nil {
			return nil
		}
		if attachedToWorkspaceRegex.MatchString(err.Error()) {
			log.Printf("[INFO] %s. Retrying until the workspace deletion is complete", err)
			return retry.RetryableError(err)
		}
		return retry.NonRetryableError(err)
	})
}
