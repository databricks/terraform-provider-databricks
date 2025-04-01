package libraries

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestWaitForLibrariesInstalledSdk(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:       "GET",
			Resource:     "/api/2.0/libraries/cluster-status?cluster_id=missing",
			ReuseRequest: true,
			Status:       404,
			Response: &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "missing",
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/libraries/cluster-status?cluster_id=error",
			ReuseRequest: true,
			Status:       500,
			Response: apierr.APIError{
				Message: "internal error",
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/libraries/cluster-status?cluster_id=1005-abcd",
			ReuseRequest: true,
			Status:       400,
			Response: apierr.APIError{
				Message: "Cluster 1005-abcd does not exist",
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/libraries/cluster-status?cluster_id=still-installing",
			ReuseRequest: true,
			Response: compute.ClusterLibraryStatuses{
				ClusterId: "still-installing",
				LibraryStatuses: []compute.LibraryFullStatus{
					{
						Status: "PENDING",
						Library: &compute.Library{
							Jar: "a.jar",
						},
					},
				},
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/libraries/cluster-status?cluster_id=failed-wheel",
			ReuseRequest: true,
			Response: compute.ClusterLibraryStatuses{
				ClusterId: "still-installing",
				LibraryStatuses: []compute.LibraryFullStatus{
					{
						Status:   "FAILED",
						Messages: []string{"does not compute"},
						Library: &compute.Library{
							Whl: "b.whl",
						},
					},
					{
						Status: "INSTALLED",
						Library: &compute.Library{
							Jar: "a.jar",
						},
					},
				},
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/libraries/uninstall",
			ExpectedRequest: compute.UninstallLibraries{
				ClusterId: "failed-wheel",
				Libraries: []compute.Library{
					{
						Whl: "b.whl",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		w, err := client.WorkspaceClient()
		if err != nil {
			panic(err)
		}
		_, err = WaitForLibrariesInstalledSdk(ctx, w, compute.Wait{
			ClusterID: "missing", Libraries: []compute.Library{}, IsRunning: true, IsRefresh: false,
		}, 50*time.Millisecond)
		assert.EqualError(t, err, "missing")

		_, err = WaitForLibrariesInstalledSdk(ctx, w, compute.Wait{
			ClusterID: "error", Libraries: []compute.Library{}, IsRunning: true, IsRefresh: false,
		}, 50*time.Millisecond)
		assert.EqualError(t, err, "internal error")

		// cluster is not running
		_, err = WaitForLibrariesInstalledSdk(ctx, w, compute.Wait{
			ClusterID: "still-installing", Libraries: []compute.Library{}, IsRunning: false, IsRefresh: false,
		}, 50*time.Millisecond)
		assert.NoError(t, err)

		// cluster is running
		_, err = WaitForLibrariesInstalledSdk(ctx, w, compute.Wait{
			ClusterID: "still-installing", Libraries: []compute.Library{}, IsRunning: true, IsRefresh: false,
		}, 50*time.Millisecond)
		assert.EqualError(t, err, "0 libraries are ready, but there are still 1 pending")

		_, err = WaitForLibrariesInstalledSdk(ctx, w, compute.Wait{
			ClusterID: "failed-wheel", Libraries: []compute.Library{}, IsRunning: true, IsRefresh: false,
		}, 50*time.Millisecond)
		assert.EqualError(t, err, "whl:b.whl failed: does not compute")

		// uninstall b.whl and continue executing
		_, err = WaitForLibrariesInstalledSdk(ctx, w, compute.Wait{
			ClusterID: "failed-wheel", Libraries: []compute.Library{}, IsRunning: true, IsRefresh: true,
		}, 50*time.Millisecond)
		assert.NoError(t, err, "library should have been uninstalled and work proceeded")

		// Cluster not available or doesn't exist
		_, err = WaitForLibrariesInstalledSdk(ctx, w, compute.Wait{
			ClusterID: "1005-abcd", Libraries: []compute.Library{}, IsRunning: false, IsRefresh: false,
		}, 50*time.Millisecond)

		var ae *apierr.APIError
		assert.True(t, errors.As(err, &ae))
		assert.Equal(t, 404, ae.StatusCode)
		assert.Equal(t, "Cluster 1005-abcd does not exist", ae.Message)
	})
}
