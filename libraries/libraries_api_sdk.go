package libraries

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Given a compute.Wait struct, returns library statuses based on the input parameter.
// If wait.IsRunning is set to true, this function will wait until all of the libraries are installed to return. Otherwise, it will directly return the list of libraries.
func WaitForLibrariesInstalledSdk(ctx context.Context, w *databricks.WorkspaceClient, wait compute.Wait, timeout time.Duration) (result *compute.ClusterLibraryStatuses, err error) {
	err = resource.RetryContext(ctx, timeout, func() *resource.RetryError {
		libsClusterStatus, err := w.Libraries.ClusterStatusByClusterId(ctx, wait.ClusterID)
		if err != nil {
			var apiErr *apierr.APIError
			if !errors.As(err, &apiErr) {
				return resource.NonRetryableError(err)
			}
			if apiErr.StatusCode != 404 && strings.Contains(apiErr.Message,
				fmt.Sprintf("Cluster %s does not exist", wait.ClusterID)) {
				apiErr.ErrorCode = "NOT_FOUND"
				apiErr.StatusCode = 404
			}
			return resource.NonRetryableError(apiErr)
		}
		if !wait.IsRunning {
			log.Printf("[INFO] We don't have to wait until the libraries are installed, so just returning list of %d libraries",
				len(libsClusterStatus.LibraryStatuses))
			result = libsClusterStatus
			return nil
		}
		retry, err := libsClusterStatus.IsRetryNeeded(wait)
		if retry {
			return resource.RetryableError(err)
		}
		if err != nil {
			return resource.NonRetryableError(err)
		}
		result = libsClusterStatus
		return nil
	})
	if err != nil {
		return
	}
	if wait.IsRunning {
		installed := []compute.LibraryFullStatus{}
		cleanup := compute.UninstallLibraries{
			ClusterId: wait.ClusterID,
			Libraries: []compute.Library{},
		}
		// cleanup libraries that failed to install
		for _, v := range result.LibraryStatuses {
			if v.Status == "FAILED" {
				log.Printf("[WARN] Removing failed library %s from %s", v.Library, wait.ClusterID)
				cleanup.Libraries = append(cleanup.Libraries, *v.Library)
				continue
			}
			installed = append(installed, v)
		}
		// and result contains only the libraries that were successfully installed
		result.LibraryStatuses = installed
		if len(cleanup.Libraries) > 0 {
			err = w.Libraries.Uninstall(ctx, cleanup)
			if err != nil {
				err = fmt.Errorf("cannot cleanup libraries: %w", err)
			}
		}
	}
	return
}
