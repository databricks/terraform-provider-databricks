package clusters

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestLibraryCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceLibrary())
}

func TestLibraryCreate(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceLibrary(),
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				ReuseRequest: true,
				Response: ClusterInfo{
					State: ClusterStateRunning,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/libraries/install",
				ExpectedRequest: compute.InstallLibraries{
					Libraries: []compute.Library{
						{
							Whl: "foo.whl",
						},
					},
					ClusterId: "abc",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: compute.ClusterLibraryStatuses{
					LibraryStatuses: []compute.LibraryFullStatus{
						{
							Library: &compute.Library{
								Whl: "foo.whl",
							},
							Status: "INSTALLED",
						},
					},
				},
			},
		},
		Create: true,
		HCL: `
		cluster_id = "abc"
		whl = "foo.whl"
		`,
	}.ApplyNoError(t)
}

func TestLibraryDelete(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceLibrary(),
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				ReuseRequest: true,
				Response: ClusterInfo{
					State: ClusterStateRunning,
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: compute.ClusterLibraryStatuses{
					LibraryStatuses: []compute.LibraryFullStatus{
						{
							Library: &compute.Library{
								Whl: "foo.whl",
							},
							Status: "INSTALLED",
						},
					},
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/libraries/uninstall",
				ExpectedRequest: compute.UninstallLibraries{
					Libraries: []compute.Library{
						{
							Whl: "foo.whl",
						},
					},
					ClusterId: "abc",
				},
			},
		},
		Delete: true,
		ID:     "abc/whl:foo.whl",
	}.ApplyNoError(t)
}

func TestLibraryDeleteClusterNotFound(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceLibrary(),
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				ReuseRequest: true,
				Response: apierr.APIError{
					ErrorCode: "NOT_FOUND",
					Message:   "Cluster does not exist",
				},
				Status: 404,
			},
		},
		Delete:  true,
		Removed: true,
		ID:      "abc/whl:foo.whl",
	}.ApplyNoError(t)
}
