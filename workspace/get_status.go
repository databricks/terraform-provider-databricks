package workspace

import (
	"context"
	"log"
	"regexp"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/workspace"
)

var timeoutRegex = regexp.MustCompile(`request timed out after .* of inactivity`)

func robustGetStatus(ctx context.Context, w *databricks.WorkspaceClient, path string) (status *workspace.ObjectInfo, err error) {
	r := retries.New[workspace.ObjectInfo](retries.WithRetryFunc(func(err error) bool {
		msg := err.Error()
		isTimeout := timeoutRegex.MatchString(msg)
		if isTimeout {
			log.Printf("[DEBUG] Retrying due to timeout: %s", msg)
		}
		return isTimeout
	}))
	return r.Run(ctx, func(ctx context.Context) (*workspace.ObjectInfo, error) {
		return w.Workspace.GetStatusByPath(ctx, path)
	})
}
