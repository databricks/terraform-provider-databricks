package compute

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/qa"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/google/go-cmp/cmp"
)

func TestClustersAPI_CreateAndWait_timesOutAndCleansUp(t *testing.T) {
	const clusterID = "test-cluster"
	createRequest := CreateCluster{
		ClusterName:            "test-cluster",
		SparkVersion:           "15.4.x-scala2.12",
		InstancePoolId:         "test-pool",
		AutoterminationMinutes: 10,
		NumWorkers:             1,
	}
	pending := ClusterDetails{
		ClusterId:    clusterID,
		State:        StatePending,
		StateMessage: "still provisioning",
	}

	qa.HTTPFixtures{
		{
			Method:          http.MethodPost,
			Resource:        "/api/2.1/clusters/create",
			ExpectedRequest: createRequest,
			Response: CreateClusterResponse{
				ClusterId: clusterID,
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.1/clusters/get?cluster_id=" + clusterID,
			Response: pending,
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.1/clusters/get?cluster_id=" + clusterID,
			Response: pending,
		},
		{
			Method:   http.MethodPost,
			Resource: "/api/2.1/clusters/delete",
			ExpectedRequest: DeleteCluster{
				ClusterId: clusterID,
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.1/clusters/get?cluster_id=" + clusterID,
			Response: ClusterDetails{
				ClusterId: clusterID,
				State:     StateTerminated,
			},
		},
	}.ApplyClient(t, func(ctx context.Context, apiClient *client.DatabricksClient) {
		clusters := NewClusters(apiClient)
		createCtx, cancelCreate := context.WithCancel(ctx)
		defer cancelCreate()

		var gotStates []State
		_, err := clusters.CreateAndWait(
			createCtx,
			createRequest,
			retries.Timeout[ClusterDetails](time.Minute),
			retries.OnPoll(func(details *ClusterDetails) {
				gotStates = append(gotStates, details.State)
				if len(gotStates) == 2 {
					cancelCreate()
				}
			}),
		)
		if want := "timed out: still provisioning"; err == nil || err.Error() != want {
			t.Errorf("create error: want %q, got %v", want, err)
		}
		if diff := cmp.Diff([]State{StatePending, StatePending}, gotStates); diff != "" {
			t.Errorf("polled states mismatch (-want +got):\n%s", diff)
		}

		deleted, err := clusters.DeleteByClusterIdAndWait(ctx, clusterID)
		if err != nil {
			t.Fatal(err)
		}
		if deleted.ClusterId != clusterID || deleted.State != StateTerminated {
			t.Errorf("deleted cluster: want %s %s, got %s %s", clusterID, StateTerminated, deleted.ClusterId, deleted.State)
		}
	})
}

func TestLibrariesAPI_UpdateAndWait_installsThenUninstalls(t *testing.T) {
	const clusterID = "test-cluster"
	library := Library{
		Pypi: &PythonPyPiLibrary{
			Package: "test-package",
		},
	}
	installRequest := InstallLibraries{
		ClusterId: clusterID,
		Libraries: []Library{library},
	}
	uninstallRequest := UninstallLibraries{
		ClusterId: clusterID,
		Libraries: []Library{library},
	}

	qa.HTTPFixtures{
		{
			Method:          http.MethodPost,
			Resource:        "/api/2.0/libraries/install",
			ExpectedRequest: installRequest,
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/libraries/cluster-status?cluster_id=" + clusterID,
			Response: ClusterLibraryStatuses{
				ClusterId: clusterID,
				LibraryStatuses: []LibraryFullStatus{
					{
						Library: &library,
						Status:  LibraryInstallStatusPending,
					},
				},
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/libraries/cluster-status?cluster_id=" + clusterID,
			Response: ClusterLibraryStatuses{
				ClusterId: clusterID,
				LibraryStatuses: []LibraryFullStatus{
					{
						Library: &library,
						Status:  LibraryInstallStatusInstalled,
					},
				},
			},
		},
		{
			Method:          http.MethodPost,
			Resource:        "/api/2.0/libraries/uninstall",
			ExpectedRequest: uninstallRequest,
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/libraries/cluster-status?cluster_id=" + clusterID,
			Response: ClusterLibraryStatuses{
				ClusterId: clusterID,
				LibraryStatuses: []LibraryFullStatus{
					{
						Library: &library,
						Status:  LibraryInstallStatusUninstallOnRestart,
					},
				},
			},
		},
	}.ApplyClient(t, func(ctx context.Context, apiClient *client.DatabricksClient) {
		libraries := NewLibraries(apiClient)
		timeout := retries.Timeout[ClusterLibraryStatuses](5 * time.Second)

		err := libraries.UpdateAndWait(ctx, Update{
			ClusterId: clusterID,
			Install:   []Library{library},
		}, timeout)
		if err != nil {
			t.Fatalf("install: %s", err)
		}

		err = libraries.UpdateAndWait(ctx, Update{
			ClusterId: clusterID,
			Uninstall: []Library{library},
		}, timeout)
		if err != nil {
			t.Fatalf("uninstall: %s", err)
		}
	})
}
