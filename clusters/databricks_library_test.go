package clusters

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/libraries"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
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
				Resource:     "/api/2.0/clusters/get?cluster_id=abc",
				ReuseRequest: true,
				Response: ClusterInfo{
					State: ClusterStateRunning,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/libraries/install",
				ExpectedRequest: libraries.ClusterLibraryList{
					Libraries: []libraries.Library{
						{
							Whl: "foo.whl",
						},
					},
					ClusterID: "abc",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: libraries.ClusterLibraryStatuses{
					LibraryStatuses: []libraries.LibraryStatus{
						{
							Library: &libraries.Library{
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
				Resource:     "/api/2.0/clusters/get?cluster_id=abc",
				ReuseRequest: true,
				Response: ClusterInfo{
					State: ClusterStateRunning,
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: libraries.ClusterLibraryStatuses{
					LibraryStatuses: []libraries.LibraryStatus{
						{
							Library: &libraries.Library{
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
				ExpectedRequest: libraries.ClusterLibraryList{
					Libraries: []libraries.Library{
						{
							Whl: "foo.whl",
						},
					},
					ClusterID: "abc",
				},
			},
		},
		Delete: true,
		ID:     "abc/whl:foo.whl",
	}.ApplyNoError(t)
}
