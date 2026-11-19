package common

import (
	"context"
	"fmt"
	"time"
)

// OwnerRollbackTimeout bounds the owner rollback, which ignores Update's cancellation
// and would otherwise retry forever under the default retry timeout.
const OwnerRollbackTimeout = time.Minute

// OwnerRollbackContext returns the context for the owner rollback after a failed update.
// It keeps ctx's values but not its cancellation, because a rollback on a cancelled
// context fails in the SDK's rate limiter before it is sent and leaves the new owner on
// the server. It ends after OwnerRollbackTimeout.
func OwnerRollbackContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.WithoutCancel(ctx), OwnerRollbackTimeout)
}

func OwnerRollbackError(err error, rollbackErr error, oldOwner string, newOwner string) error {
	help := fmt.Sprintf("Owner of this resource was updated but other fields couldn't be updated and owner couldn't be rolled back. \n As a result, the owner of this resource is updated to %s but other attributes aren't. To revert the owner change, please manually change the owner to %s. \n\n You can also use the databricks cli (https://docs.databricks.com/en/dev-tools/cli/install.html) to update the owner.", newOwner, oldOwner)
	return fmt.Errorf("%w. Owner rollback also failed: %w. %s", err, rollbackErr, help)
}
