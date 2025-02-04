package clusters

import (
	"fmt"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var nothingPinned = qa.HTTPFixture{
	Method:   "GET",
	Resource: "/api/2.1/clusters/list?filter_by.is_pinned=true&page_size=100",
	Response: compute.ListClustersResponse{
		Clusters: []compute.ClusterDetails{},
	},
}

func TestResourceClusterCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			nothingPinned,
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/create",
				ExpectedRequest: compute.ClusterSpec{
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
				},
				Response: compute.ClusterDetails{
					ClusterId: "abc",
					State:     compute.StateRunning,
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				Response: compute.ClusterDetails{
					ClusterId:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  compute.StateRunning,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: compute.ClusterLibraryStatuses{
					LibraryStatuses: []compute.LibraryFullStatus{},
				},
			},
		},
		Create:   true,
		Resource: ResourceCluster(),
		State: map[string]any{
			"autotermination_minutes": 15,
			"cluster_name":            "Shared Autoscaling",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             100,
			"is_pinned":               false,
		},
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, false, d.Get("is_pinned"))
}

func TestResourceClusterCreatePinned(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/create",
				ExpectedRequest: compute.CreateCluster{
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
				},
				Response: compute.ClusterDetails{
					ClusterId: "abc",
					State:     compute.StateRunning,
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				Response: compute.ClusterDetails{
					ClusterId:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  compute.StateRunning,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/pin",
				ExpectedRequest: compute.PinCluster{
					ClusterId: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: compute.ClusterLibraryStatuses{
					LibraryStatuses: []compute.LibraryFullStatus{},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/clusters/list?filter_by.is_pinned=true&page_size=100",
				Response: compute.ListClustersResponse{
					Clusters: []compute.ClusterDetails{{
						ClusterId:              "abc",
						NumWorkers:             100,
						ClusterName:            "Shared Autoscaling",
						SparkVersion:           "7.1-scala12",
						NodeTypeId:             "i3.xlarge",
						AutoterminationMinutes: 15,
						State:                  compute.StateRunning,
					}},
				},
			},
		},
		Create:   true,
		Resource: ResourceCluster(),
		State: map[string]any{
			"autotermination_minutes": 15,
			"cluster_name":            "Shared Autoscaling",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             100,
			"is_pinned":               true,
		},
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, true, d.Get("is_pinned"))
}

func TestResourceClusterCreateErrorFollowedByDeletion(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/create",
				ExpectedRequest: compute.CreateCluster{
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
				},
				Response: compute.ClusterDetails{
					ClusterId: "abc",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				Response: compute.ClusterDetails{
					ClusterId:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  compute.StateTerminated,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/permanent-delete",
				ExpectedRequest: compute.PermanentDeleteCluster{
					ClusterId: "abc",
				},
			},
		},
		Create:   true,
		Resource: ResourceCluster(),
		State: map[string]any{
			"autotermination_minutes": 15,
			"cluster_name":            "Shared Autoscaling",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             100,
		},
	}.Apply(t)
	assert.ErrorContains(t, err, "failed to reach RUNNING, got TERMINATED")
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterCreateErrorFollowedByDeletionError(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/create",
				ExpectedRequest: compute.CreateCluster{
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
				},
				Response: compute.ClusterDetails{
					ClusterId: "abc",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				Response: compute.ClusterDetails{
					ClusterId:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  compute.StateTerminated,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/permanent-delete",
				ExpectedRequest: compute.PermanentDeleteCluster{
					ClusterId: "abc",
				},
				Status: 500,
				Response: common.APIErrorBody{
					ErrorCode: "INTERNAL_ERROR",
					Message:   "Internal error happened",
				},
			},
		},
		Create:   true,
		Resource: ResourceCluster(),
		State: map[string]any{
			"autotermination_minutes": 15,
			"cluster_name":            "Shared Autoscaling",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             100,
		},
	}.Apply(t)
	assert.ErrorContains(t, err, "failed to create cluster: failed to reach RUNNING, got TERMINATED:  and failed to delete it during cleanup: Internal error happened")
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterCreate_WithLibraries(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			nothingPinned,
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/create",
				ExpectedRequest: compute.CreateCluster{
					NumWorkers:             100,
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 60,
				},
				Response: compute.ClusterDetails{
					ClusterId: "abc",
					State:     compute.StateRunning,
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				Response: compute.ClusterDetails{
					ClusterId:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  compute.StateRunning,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/libraries/install",
				ExpectedRequest: compute.InstallLibraries{
					ClusterId: "abc",
					Libraries: []compute.Library{
						{
							Jar: "dbfs://foo.jar",
						},
						{
							Cran: &compute.RCranLibrary{
								Package: "rkeops",
								Repo:    "internal",
							},
						},
						{
							Egg: "dbfs://bar.egg",
						},
						{
							Maven: &compute.MavenLibrary{
								Coordinates: "foo:bar:baz:0.1.0",
								Exclusions:  []string{"org.apache:flink:base"},
								Repo:        "s3://maven-repo-in-s3/release",
							},
						},
						{
							Pypi: &compute.PythonPyPiLibrary{
								Package: "seaborn==1.2.4",
							},
						},
						{
							Whl: "dbfs://baz.whl",
						},
					},
				},
			},
			{
				Method: "GET",
				// 1 of 3 requests
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: compute.ClusterLibraryStatuses{
					LibraryStatuses: []compute.LibraryFullStatus{
						{
							Library: &compute.Library{
								Pypi: &compute.PythonPyPiLibrary{
									Package: "seaborn==1.2.4",
								},
							},
							Status: compute.LibraryInstallStatusPending,
						},
						{
							Library: &compute.Library{
								Whl: "dbfs://baz.whl",
							},
							Status: compute.LibraryInstallStatusInstalled,
						},
					},
				},
			},
			{
				Method: "GET",
				// 2 of 3 requests
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: compute.ClusterLibraryStatuses{
					LibraryStatuses: []compute.LibraryFullStatus{
						{
							Library: &compute.Library{
								Pypi: &compute.PythonPyPiLibrary{
									Package: "seaborn==1.2.4",
								},
							},
							Status: compute.LibraryInstallStatusInstalled,
						},
						{
							Library: &compute.Library{
								Whl: "dbfs://baz.whl",
							},
							Status: compute.LibraryInstallStatusInstalled,
						},
					},
				},
			},
			{
				Method: "GET",
				// 3 of 3 requests
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: compute.ClusterLibraryStatuses{
					LibraryStatuses: []compute.LibraryFullStatus{
						{
							Library: &compute.Library{
								Pypi: &compute.PythonPyPiLibrary{
									Package: "seaborn==1.2.4",
								},
							},
							Status: compute.LibraryInstallStatusInstalled,
						},
						{
							Library: &compute.Library{
								Whl: "dbfs://baz.whl",
							},
							Status: compute.LibraryInstallStatusInstalled,
						},
					},
				},
			},
		},
		Create:   true,
		Resource: ResourceCluster(),
		HCL: `num_workers = 100
		spark_version = "7.1-scala12"
		node_type_id = "i3.xlarge"

		library {
			jar = "dbfs://foo.jar"
		}

		library {
			egg = "dbfs://bar.egg"
		}

		library {
			whl = "dbfs://baz.whl"
		}

		library {
			pypi {
				package = "seaborn==1.2.4"
			}
		}

		library {
			maven {
				coordinates = "foo:bar:baz:0.1.0"
				repo = "s3://maven-repo-in-s3/release"
				exclusions = [
					"org.apache:flink:base"
				]
			}
		}

		library {
			cran {
				package = "rkeops"
				repo = "internal"
			}
		}`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterCreatePhoton(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			nothingPinned,
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/create",
				ExpectedRequest: compute.CreateCluster{
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					RuntimeEngine:          "PHOTON",
				},
				Response: compute.ClusterDetails{
					ClusterId: "abc",
					State:     compute.StateRunning,
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				Response: compute.ClusterDetails{
					ClusterId:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  compute.StateRunning,
					RuntimeEngine:          "PHOTON",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: compute.ClusterLibraryStatuses{
					LibraryStatuses: []compute.LibraryFullStatus{},
				},
			},
		},
		Create:   true,
		Resource: ResourceCluster(),
		State: map[string]any{
			"autotermination_minutes": 15,
			"cluster_name":            "Shared Autoscaling",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             100,
			"is_pinned":               false,
			"runtime_engine":          "PHOTON",
		},
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterCreateNoWait_WithLibraries(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			nothingPinned,
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/create",
				ExpectedRequest: compute.ClusterSpec{
					NumWorkers:             100,
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 60,
				},
				Response: compute.ClusterDetails{
					ClusterId: "abc",
					State:     compute.StateUnknown,
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				Response: compute.ClusterDetails{
					ClusterId:              "abc",
					NumWorkers:             100,
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  compute.StateUnknown,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/libraries/install",
				ExpectedRequest: compute.InstallLibraries{
					ClusterId: "abc",
					Libraries: []compute.Library{
						{
							Pypi: &compute.PythonPyPiLibrary{
								Package: "seaborn==1.2.4",
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: compute.ClusterLibraryStatuses{
					LibraryStatuses: []compute.LibraryFullStatus{
						{
							Library: &compute.Library{
								Pypi: &compute.PythonPyPiLibrary{
									Package: "seaborn==1.2.4",
								},
							},
							Status: compute.LibraryInstallStatusPending,
						},
					},
				},
			},
		},
		Create:   true,
		Resource: ResourceCluster(),
		HCL: `num_workers = 100
		spark_version = "7.1-scala12"
		node_type_id = "i3.xlarge"
		no_wait = true

		library {
			pypi {
				package = "seaborn==1.2.4"
			}
		}`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterCreateNoWait(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			nothingPinned,
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/create",
				ExpectedRequest: compute.ClusterSpec{
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
				},
				Response: compute.ClusterDetails{
					ClusterId: "abc",
					State:     compute.StateUnknown,
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				Response: compute.ClusterDetails{
					ClusterId:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  compute.StateUnknown,
				},
			},
		},
		Create:   true,
		Resource: ResourceCluster(),
		State: map[string]any{
			"autotermination_minutes": 15,
			"cluster_name":            "Shared Autoscaling",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             100,
			"is_pinned":               false,
			"no_wait":                 true,
		},
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/create",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Create:   true,
		Resource: ResourceCluster(),
		State: map[string]any{
			"autotermination_minutes": 15,
			"cluster_name":            "Shared Autoscaling",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             100,
		},
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceClusterRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			nothingPinned,
			{
				Method:   "GET",
				Resource: "/api/2.1/clusters/get?cluster_id=abc",
				Response: compute.ClusterDetails{
					ClusterId:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  compute.StateRunning,
					Autoscale: &compute.AutoScale{
						MaxWorkers: 4,
					},
				},
			},
		},
		Resource: ResourceCluster(),
		Read:     true,
		ID:       "abc",
		New:      true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, 15, d.Get("autotermination_minutes"))
	assert.Equal(t, "Shared Autoscaling", d.Get("cluster_name"))
	assert.Equal(t, "i3.xlarge", d.Get("node_type_id"))
	assert.Equal(t, 4, d.Get("autoscale.0.max_workers"))
	assert.Equal(t, "RUNNING", d.Get("state"))
	assert.Equal(t, false, d.Get("is_pinned"))

	for k, v := range d.State().Attributes {
		fmt.Printf("assert.Equal(t, %#v, d.Get(%#v))\n", v, k)
	}
}

func TestResourceClusterRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/clusters/get?cluster_id=abc",
				Response: common.APIErrorBody{
					// clusters API is not fully restful, so let's test for that
					// TODO: https://github.com/databricks/terraform-provider-databricks/issues/2021
					ErrorCode: "INVALID_STATE",
					Message:   "Cluster abc does not exist",
				},
				Status: 400,
			},
		},
		Resource: ResourceCluster(),
		Read:     true,
		Removed:  true,
		ID:       "abc",
	}.ApplyNoError(t)
}

func TestResourceClusterRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/clusters/get?cluster_id=abc",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceCluster(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id(), "Id should not be empty for error reads")
}

// resize api should be called when autoscaling cluster is converted to a non autoscaling one
func TestResourceClusterUpdate_ResizeForAutoscalingToNumWorkersCluster(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			nothingPinned,
			{
				Method:       "GET",
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				ReuseRequest: true,
				Response: compute.ClusterDetails{
					ClusterId: "abc",
					Autoscale: &compute.AutoScale{
						MinWorkers: 1,
						MaxWorkers: 4,
					},
					ClusterName:            "Non Autoscaling Cluster",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  compute.StateRunning,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/resize",
				ExpectedRequest: compute.ResizeCluster{
					ClusterId:  "abc",
					NumWorkers: 3,
				},
			},
		},
		ID:       "abc",
		Update:   true,
		Resource: ResourceCluster(),
		HCL: `
		autotermination_minutes = 15,
		cluster_name =            "Non Autoscaling Cluster"
		spark_version =           "7.1-scala12"
		node_type_id =            "i3.xlarge"
		num_workers = 3
		`,
		InstanceState: map[string]string{
			"autotermination_minutes": "15",
			"cluster_name":            "Non Autoscaling Cluster",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"autoscale": `"{
				min_workers = 1
				max_workers = 4
			}"`,
		},
	}.ApplyNoError(t)
}

// resize api should be called when non autoscaling cluster is converted to a autoscaling one
func TestResourceClusterUpdate_ResizeForNumWorkersToAutoscalingCluster(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			nothingPinned,
			{
				Method:       "GET",
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				ReuseRequest: true,
				Response: compute.ClusterDetails{
					ClusterId:              "abc",
					NumWorkers:             150,
					ClusterName:            "Non Autoscaling Cluster",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  compute.StateRunning,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/resize",
				ExpectedRequest: compute.ResizeCluster{
					ClusterId: "abc",
					Autoscale: &compute.AutoScale{
						MinWorkers: 4,
						MaxWorkers: 10,
					},
				},
			},
		},
		ID:       "abc",
		Update:   true,
		Resource: ResourceCluster(),
		HCL: `
		autotermination_minutes = 15,
		cluster_name =            "Non Autoscaling Cluster"
		spark_version =           "7.1-scala12"
		node_type_id =            "i3.xlarge"
		autoscale = {
			min_workers = 4
			max_workers = 10
		},
		`,
		InstanceState: map[string]string{
			"autotermination_minutes": "15",
			"cluster_name":            "Non Autoscaling Cluster",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             "150",
		},
	}.ApplyNoError(t)
}

// provider should call the edit api and not the resize api when the cluster is not running
func TestResourceClusterUpdate_EditNumWorkersWhenClusterTerminated(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			nothingPinned,
			{
				Method:       "GET",
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				ReuseRequest: true,
				Response: compute.ClusterDetails{
					ClusterId:              "abc",
					NumWorkers:             150,
					ClusterName:            "Non Autoscaling Cluster",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  compute.StateTerminated,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/edit",
				ExpectedRequest: compute.ClusterDetails{
					AutoterminationMinutes: 15,
					ClusterId:              "abc",
					NumWorkers:             100,
					ClusterName:            "Non Autoscaling Cluster",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
				},
			},
		},
		ID:       "abc",
		Update:   true,
		Resource: ResourceCluster(),
		State: map[string]any{
			"autotermination_minutes": 15,
			"cluster_name":            "Non Autoscaling Cluster",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             100,
		},
		InstanceState: map[string]string{
			"autotermination_minutes": "15",
			"cluster_name":            "Non Autoscaling Cluster",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             "150",
		},
	}.ApplyNoError(t)
}

func TestResourceClusterUpdate_ResizeAutoscale(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			nothingPinned,
			{
				Method:       "GET",
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				ReuseRequest: true,
				Response: compute.ClusterDetails{
					ClusterId: "abc",
					Autoscale: &compute.AutoScale{
						MaxWorkers: 4,
					},
					ClusterName:  "Shared Autoscaling",
					SparkVersion: "7.1-scala12",
					NodeTypeId:   "i3.xlarge",
					State:        compute.StateRunning,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/resize",
				ExpectedRequest: compute.ResizeCluster{
					ClusterId: "abc",
					Autoscale: &compute.AutoScale{
						MinWorkers: 4,
						MaxWorkers: 10,
					},
				},
			},
		},
		ID:       "abc",
		Update:   true,
		Resource: ResourceCluster(),
		InstanceState: map[string]string{
			"autotermination_minutes": "15",
			"cluster_name":            "Autoscaling Cluster",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"autoscale": `
			{
				min_workers = 0
				max_workers = 4
			}`,
		},
		HCL: `
		autotermination_minutes = 15,
		cluster_name =            "Autoscaling Cluster"
		spark_version =           "7.1-scala12"
		node_type_id =            "i3.xlarge"
		autoscale = {
			min_workers = 4
			max_workers = 10
		}
		`,
	}.ApplyNoError(t)
}

func TestResourceClusterUpdate_ResizeNumWorkers(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			nothingPinned,
			{
				Method:       "GET",
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				ReuseRequest: true,
				Response: compute.ClusterDetails{
					ClusterId:              "abc",
					NumWorkers:             150,
					ClusterName:            "Non Autoscaling Cluster",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  compute.StateRunning,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/resize",
				ExpectedRequest: compute.ResizeCluster{
					ClusterId:  "abc",
					NumWorkers: 100,
				},
			},
		},
		ID:       "abc",
		Update:   true,
		Resource: ResourceCluster(),
		State: map[string]any{
			"autotermination_minutes": 15,
			"cluster_name":            "Non Autoscaling Cluster",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             100,
		},
		InstanceState: map[string]string{
			"autotermination_minutes": "15",
			"cluster_name":            "Non Autoscaling Cluster",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             "150",
		},
	}.ApplyNoError(t)
}

func TestResourceClusterUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			nothingPinned,
			{
				Method:       "GET",
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				ReuseRequest: true,
				Response: compute.ClusterDetails{
					ClusterId:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  compute.StateRunning,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/start",
				ExpectedRequest: compute.StartCluster{
					ClusterId: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: compute.ClusterLibraryStatuses{
					LibraryStatuses: []compute.LibraryFullStatus{},
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/edit",
				ExpectedRequest: compute.ClusterDetails{
					AutoterminationMinutes: 15,
					ClusterId:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: compute.ClusterLibraryStatuses{
					LibraryStatuses: []compute.LibraryFullStatus{},
				},
			},
		},
		ID:       "abc",
		Update:   true,
		Resource: ResourceCluster(),
		State: map[string]any{
			"autotermination_minutes": 15,
			"cluster_name":            "Shared Autoscaling",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             100,
		},
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id(), "Id should be the same as in reading")
}

func TestResourceClusterUpdate_WhileScaling(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			nothingPinned,
			{
				Method:       "GET",
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				ReuseRequest: true,
				Response: compute.ClusterDetails{
					ClusterId:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  compute.StateRunning,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/start",
				ExpectedRequest: compute.StartCluster{
					ClusterId: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: compute.ClusterLibraryStatuses{
					LibraryStatuses: []compute.LibraryFullStatus{},
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/edit",
				ExpectedRequest: compute.ClusterDetails{
					AutoterminationMinutes: 15,
					ClusterId:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
				},
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_STATE",
				},
				Status: 404,
			},
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/edit",
				ExpectedRequest: compute.ClusterDetails{
					AutoterminationMinutes: 15,
					ClusterId:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: compute.ClusterLibraryStatuses{
					LibraryStatuses: []compute.LibraryFullStatus{},
				},
			},
		},
		ID:       "abc",
		Update:   true,
		Resource: ResourceCluster(),
		State: map[string]any{
			"autotermination_minutes": 15,
			"cluster_name":            "Shared Autoscaling",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             100,
		},
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id(), "Id should be the same as in reading")
}

func TestResourceClusterUpdateWithPinned(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			nothingPinned,
			{
				Method:       "GET",
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				ReuseRequest: true,
				Response: compute.ClusterDetails{
					ClusterId:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  compute.StateRunning,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/start",
				ExpectedRequest: compute.StartCluster{
					ClusterId: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: compute.ClusterLibraryStatuses{
					LibraryStatuses: []compute.LibraryFullStatus{},
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/pin",
				ExpectedRequest: compute.PinCluster{
					ClusterId: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: compute.ClusterLibraryStatuses{
					LibraryStatuses: []compute.LibraryFullStatus{},
				},
			},
		},
		ID:       "abc",
		Update:   true,
		Resource: ResourceCluster(),
		InstanceState: map[string]string{
			"autotermination_minutes": "15",
			"cluster_name":            "Shared Autoscaling",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             "100",
		},
		State: map[string]any{
			"autotermination_minutes": 15,
			"cluster_name":            "Shared Autoscaling",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             100,
			"is_pinned":               true,
		},
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id(), "Id should be the same as in reading")
}

func TestResourceClusterUpdate_LibrariesChangeOnTerminatedCluster(t *testing.T) {
	terminated := qa.HTTPFixture{
		Method:   "GET",
		Resource: "/api/2.1/clusters/get?cluster_id=abc",
		Response: compute.ClusterDetails{
			ClusterId:    "abc",
			NumWorkers:   100,
			SparkVersion: "7.1-scala12",
			NodeTypeId:   "i3.xlarge",
			State:        compute.StateTerminated,
			StateMessage: "Terminated for test reasons",
		},
	}
	newLibs := qa.HTTPFixture{
		Method:   "GET",
		Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
		Response: compute.ClusterLibraryStatuses{
			ClusterId: "abc",
			LibraryStatuses: []compute.LibraryFullStatus{
				{
					Library: &compute.Library{
						Jar: "dbfs://foo.jar",
					},
					Status: compute.LibraryInstallStatusInstalled,
				},
				{
					Library: &compute.Library{
						Egg: "dbfs://bar.egg",
					},
					Status: compute.LibraryInstallStatusInstalled,
				},
			},
		},
	}
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			nothingPinned,
			terminated, // 1 of ...
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/edit",
				ExpectedRequest: compute.EditCluster{
					AutoterminationMinutes: 60,
					ClusterId:              "abc",
					NumWorkers:             100,
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: compute.ClusterLibraryStatuses{
					ClusterId: "abc",
					LibraryStatuses: []compute.LibraryFullStatus{
						{
							Library: &compute.Library{
								Egg: "dbfs://bar.egg",
							},
							Status: compute.LibraryInstallStatusInstalled,
						},
						{
							Library: &compute.Library{
								Pypi: &compute.PythonPyPiLibrary{
									Package: "requests",
								},
							},
							Status: compute.LibraryInstallStatusInstalled,
						},
					},
				},
			},
			{ // check to see if cluster is restarting (if so wait)
				Method:   "GET",
				Resource: "/api/2.1/clusters/get?cluster_id=abc",
				Response: compute.ClusterDetails{
					ClusterId:    "abc",
					NumWorkers:   100,
					SparkVersion: "7.1-scala12",
					NodeTypeId:   "i3.xlarge",
					State:        compute.StateTerminated,
				},
			},
			{ // start cluster before libs install
				Method:   "POST",
				Resource: "/api/2.1/clusters/start",
				ExpectedRequest: compute.StartCluster{
					ClusterId: "abc",
				},
			},
			{ // 2 of ...
				Method:   "GET",
				Resource: "/api/2.1/clusters/get?cluster_id=abc",
				Response: compute.ClusterDetails{
					ClusterId:    "abc",
					NumWorkers:   100,
					SparkVersion: "7.1-scala12",
					NodeTypeId:   "i3.xlarge",
					State:        compute.StateRunning,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/libraries/uninstall",
				ExpectedRequest: compute.UninstallLibraries{
					ClusterId: "abc",
					Libraries: []compute.Library{
						{
							Pypi: &compute.PythonPyPiLibrary{
								Package: "requests",
							},
						},
					},
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/libraries/install",
				ExpectedRequest: compute.InstallLibraries{
					ClusterId: "abc",
					Libraries: []compute.Library{
						{
							Jar: "dbfs://foo.jar",
						},
					},
				},
			},
			newLibs,
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/delete",
				ExpectedRequest: compute.DeleteCluster{
					ClusterId: "abc",
				},
			},
			terminated, // 3 of 4
			// read
			terminated, // 4 of 4
			newLibs,
		},
		ID:       "abc",
		Update:   true,
		Resource: ResourceCluster(),
		HCL: `num_workers = 100
		spark_version = "7.1-scala12"
		node_type_id = "i3.xlarge"

		library {
			jar = "dbfs://foo.jar"
		}

		library {
			egg = "dbfs://bar.egg"
		}`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id(), "Id should be the same as in reading")
}

func TestResourceClusterUpdate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/clusters/get?cluster_id=abc",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		ID:       "abc",
		Update:   true,
		Resource: ResourceCluster(),
		State: map[string]any{
			"autotermination_minutes": 15,
			"cluster_name":            "Shared Autoscaling",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             100,
		},
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterUpdate_AutoAz(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			nothingPinned,
			{
				Method:       "GET",
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				ReuseRequest: true,
				Response: compute.ClusterDetails{
					ClusterId:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  compute.StateRunning,
					AwsAttributes: &compute.AwsAttributes{
						Availability:  "SPOT",
						FirstOnDemand: 1,
						ZoneId:        "us-west-2a",
					},
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/start",
				ExpectedRequest: compute.StartCluster{
					ClusterId: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: compute.ClusterLibraryStatuses{
					LibraryStatuses: []compute.LibraryFullStatus{},
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/edit",
				ExpectedRequest: compute.EditCluster{
					AutoterminationMinutes: 15,
					ClusterId:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AwsAttributes: &compute.AwsAttributes{
						Availability:  "SPOT",
						FirstOnDemand: 1,
						ZoneId:        "auto",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: compute.ClusterLibraryStatuses{
					LibraryStatuses: []compute.LibraryFullStatus{},
				},
			},
		},
		ID:       "abc",
		Update:   true,
		Resource: ResourceCluster(),
		HCL: `
		autotermination_minutes = 15
		cluster_name = "Shared Autoscaling"
		spark_version = "7.1-scala12"
		node_type_id = "i3.xlarge"
		num_workers = 100,
		aws_attributes {
			availability            = "SPOT"
			zone_id                 = "auto"
			first_on_demand         = 1
		}
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id(), "Id should be the same as in reading")
}

func TestResourceClusterDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/delete",
				ExpectedRequest: compute.DeleteCluster{
					ClusterId: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/clusters/get?cluster_id=abc",
				Response: compute.ClusterDetails{
					State: compute.StateTerminated,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/permanent-delete",
				ExpectedRequest: compute.PermanentDeleteCluster{
					ClusterId: "abc",
				},
			},
		},
		Resource: ResourceCluster(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/permanent-delete",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceCluster(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterCreate_SingleNode(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			nothingPinned,
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/create",
				ExpectedRequest: compute.CreateCluster{
					NumWorkers:             0,
					ClusterName:            "Single Node Cluster",
					SparkVersion:           "7.3.x-scala12",
					NodeTypeId:             "Standard_F4s",
					AutoterminationMinutes: 120,
					SparkConf: map[string]string{
						"spark.master":                     "local[*]",
						"spark.databricks.cluster.profile": "singleNode",
					},
					CustomTags: map[string]string{
						"ResourceClass": "SingleNode",
					},
					ForceSendFields: []string{"NumWorkers"},
				},
				Response: compute.ClusterDetails{
					ClusterId: "abc",
					State:     compute.StateRunning,
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				Response: compute.ClusterDetails{
					ClusterId:              "abc",
					ClusterName:            "Single Node Cluster",
					SparkVersion:           "7.3.x-scala12",
					NodeTypeId:             "Standard_F4s",
					AutoterminationMinutes: 120,
					State:                  compute.StateRunning,
					SparkConf: map[string]string{
						"spark.master":                     "local[*]",
						"spark.databricks.cluster.profile": "singleNode",
					},
					CustomTags: map[string]string{
						"ResourceClass": "SingleNode",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: compute.ClusterLibraryStatuses{
					LibraryStatuses: []compute.LibraryFullStatus{},
				},
			},
		},
		Create:   true,
		Resource: ResourceCluster(),
		State: map[string]any{
			"autotermination_minutes": 120,
			"cluster_name":            "Single Node Cluster",
			"spark_version":           "7.3.x-scala12",
			"node_type_id":            "Standard_F4s",
			"is_pinned":               false,
			"spark_conf": map[string]any{
				"spark.master":                     "local[*]",
				"spark.databricks.cluster.profile": "singleNode",
			},
			"custom_tags": map[string]any{
				"ResourceClass": "SingleNode",
			},
		},
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, 0, d.Get("num_workers"))
}

func TestResourceClusterCreate_SingleNodeAutoPropertiesCustomizeDiff(t *testing.T) {
	testCases := []struct {
		name          string
		hcl           string
		instanceState map[string]string
		expectedDiff  map[string]*terraform.ResourceAttrDiff
	}{
		{
			"resource with no custom_tags or spark_conf",
			"",
			map[string]string{
				"custom_tags.ResourceClass":                   "SingleNode",
				"spark_conf.spark.master":                     "local[*]",
				"spark_conf.spark.databricks.cluster.profile": "singleNode",
			},
			nil,
		},
		{
			"resource with custom_tags and spark_conf",
			`custom_tags = {
				"ClusterName" = "SingleNode"
			}
			spark_conf = {
				"spark.databricks.delta.preview.enabled" = "true"
			}`,
			map[string]string{
				"custom_tags.ClusterName":                           "SingleNode",
				"custom_tags.ResourceClass":                         "SingleNode",
				"spark_conf.spark.master":                           "local[*]",
				"spark_conf.spark.databricks.cluster.profile":       "singleNode",
				"spark_conf.spark.databricks.delta.preview.enabled": "true",
			},
			nil,
		},
		{
			"resource with custom_tags and spark_conf changes",
			`custom_tags = {
				"ClusterName" = "MonoNode"
			}
			spark_conf = {
				"spark.databricks.delta.preview.enabled" = "false"
			}`,
			map[string]string{
				"custom_tags.ClusterName":                           "SingleNode",
				"custom_tags.ResourceClass":                         "SingleNode",
				"spark_conf.spark.master":                           "local[*]",
				"spark_conf.spark.databricks.cluster.profile":       "singleNode",
				"spark_conf.spark.databricks.delta.preview.enabled": "true",
			},
			map[string]*terraform.ResourceAttrDiff{
				"custom_tags.ClusterName":                           {New: "MonoNode", Old: "SingleNode"},
				"spark_conf.spark.databricks.delta.preview.enabled": {New: "false", Old: "true"},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			expectedDiff := testCase.expectedDiff

			if expectedDiff == nil {
				expectedDiff = make(map[string]*terraform.ResourceAttrDiff)
			}

			expectedDiff["default_tags.%"] = &terraform.ResourceAttrDiff{Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false}
			expectedDiff["driver_instance_pool_id"] = &terraform.ResourceAttrDiff{Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false}
			expectedDiff["driver_node_type_id"] = &terraform.ResourceAttrDiff{Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false}
			expectedDiff["enable_elastic_disk"] = &terraform.ResourceAttrDiff{Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false}
			expectedDiff["enable_local_disk_encryption"] = &terraform.ResourceAttrDiff{Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false}
			expectedDiff["state"] = &terraform.ResourceAttrDiff{Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false}
			expectedDiff["url"] = &terraform.ResourceAttrDiff{Old: "", New: "", NewComputed: true, NewRemoved: false, RequiresNew: false, Sensitive: false}

			instanceState := testCase.instanceState
			instanceState["cluster_id"] = "abc"
			instanceState["autotermination_minutes"] = "60"
			instanceState["cluster_name"] = "Single Node Cluster"
			instanceState["data_security_mode"] = "SINGLE_USER"
			instanceState["spark_version"] = "15.4.x-scala2.12"
			instanceState["single_user_name"] = "testuser"
			instanceState["runtime_engine"] = "STANDARD"
			instanceState["num_workers"] = "0"
			instanceState["node_type_id"] = "Standard_F4s"
			instanceState["kind"] = "CLASSIC_PREVIEW"
			instanceState["is_single_node"] = "true"

			qa.ResourceFixture{
				HCL: `
					spark_version           = "15.4.x-scala2.12"
					runtime_engine          = "STANDARD"
					node_type_id            = "Standard_F4s"
					kind                    = "CLASSIC_PREVIEW"
					cluster_name            = "Single Node Cluster"
					data_security_mode      = "SINGLE_USER"
					autotermination_minutes = 60
					is_single_node          = true
					single_user_name        = "testuser"
				` + testCase.hcl,
				InstanceState: instanceState,
				ExpectedDiff:  expectedDiff,
				Resource:      ResourceCluster(),
			}.ApplyNoError(t)
		})
	}
}

func TestResourceClusterCreate_NegativeNumWorkers(t *testing.T) {
	_, err := qa.ResourceFixture{
		Create:   true,
		Resource: ResourceCluster(),
		State: map[string]any{
			"autotermination_minutes": 120,
			"cluster_name":            "Broken Cluster",
			"spark_version":           "7.3.x-scala12",
			"node_type_id":            "Standard_F4s",
			"num_workers":             -10,
		},
	}.Apply(t)
	assert.Error(t, err)
	require.Equal(t, true, strings.Contains(err.Error(), "expected num_workers to be at least (0)"))
}

func TestResourceClusterCreate_NumWorkersIsZero(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			nothingPinned,
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/create",
				ExpectedRequest: compute.CreateCluster{
					NumWorkers:             0,
					ClusterName:            "Zero workers cluster",
					SparkVersion:           "7.3.x-scala12",
					NodeTypeId:             "Standard_F4s",
					AutoterminationMinutes: 120,
					ForceSendFields:        []string{"NumWorkers"},
				},
				Response: compute.ClusterDetails{
					ClusterId: "abc",
					State:     compute.StateRunning,
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				Response: compute.ClusterDetails{
					ClusterId:              "abc",
					ClusterName:            "Zero workers cluster",
					SparkVersion:           "7.3.x-scala12",
					NodeTypeId:             "Standard_F4s",
					AutoterminationMinutes: 120,
					State:                  compute.StateRunning,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: compute.ClusterLibraryStatuses{
					LibraryStatuses: []compute.LibraryFullStatus{},
				},
			},
		},
		Create:   true,
		Resource: ResourceCluster(),
		State: map[string]any{
			"autotermination_minutes": 120,
			"cluster_name":            "Zero workers cluster",
			"spark_version":           "7.3.x-scala12",
			"node_type_id":            "Standard_F4s",
			"is_pinned":               false,
		},
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, 0, d.Get("num_workers"))
}

func TestModifyClusterRequestAws(t *testing.T) {
	c := compute.CreateCluster{
		InstancePoolId: "a",
		AwsAttributes: &compute.AwsAttributes{
			InstanceProfileArn: "b",
			ZoneId:             "c",
		},
		EnableElasticDisk: true,
		NodeTypeId:        "d",
		DriverNodeTypeId:  "e",
	}
	err := ModifyRequestOnInstancePool(&c)
	assert.NoError(t, err)
	assert.Equal(t, "", c.AwsAttributes.ZoneId)
	assert.Equal(t, "", c.NodeTypeId)
	assert.Equal(t, "", c.DriverNodeTypeId)
	assert.Equal(t, false, c.EnableElasticDisk)
}

func TestModifyClusterRequestAzure(t *testing.T) {
	c := compute.CreateCluster{
		InstancePoolId: "a",
		AzureAttributes: &compute.AzureAttributes{
			FirstOnDemand: 1,
		},
		EnableElasticDisk: true,
		NodeTypeId:        "d",
		DriverNodeTypeId:  "e",
	}
	err := ModifyRequestOnInstancePool(&c)
	assert.NoError(t, err)
	assert.Equal(t, &compute.AzureAttributes{}, c.AzureAttributes)
	assert.Equal(t, "", c.NodeTypeId)
	assert.Equal(t, "", c.DriverNodeTypeId)
	assert.Equal(t, false, c.EnableElasticDisk)
}

func TestModifyClusterRequestGcp(t *testing.T) {
	c := compute.CreateCluster{
		InstancePoolId: "a",
		GcpAttributes: &compute.GcpAttributes{
			UsePreemptibleExecutors: true,
		},
		EnableElasticDisk: true,
		NodeTypeId:        "d",
		DriverNodeTypeId:  "e",
	}
	err := ModifyRequestOnInstancePool(&c)
	assert.NoError(t, err)
	assert.Equal(t, false, c.GcpAttributes.UsePreemptibleExecutors)
	assert.Equal(t, "", c.NodeTypeId)
	assert.Equal(t, "", c.DriverNodeTypeId)
	assert.Equal(t, false, c.EnableElasticDisk)
}

// https://github.com/databricks/terraform-provider-databricks/issues/952
func TestReadOnStoppedClusterWithLibrariesDoesNotFail(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceCluster(),
		Fixtures: []qa.HTTPFixture{
			nothingPinned,
			{
				Method:   "GET",
				Resource: "/api/2.1/clusters/get?cluster_id=foo",
				Response: compute.ClusterDetails{
					State: compute.StateTerminated,
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/libraries/cluster-status?cluster_id=foo",
				Response: compute.ClusterLibraryStatuses{
					ClusterId: "foo",
					LibraryStatuses: []compute.LibraryFullStatus{
						{
							Library: &compute.Library{
								Jar: "foo.bar",
							},
							Status: compute.LibraryInstallStatusPending,
						},
					},
				},
			},
		},
		Read: true,
		ID:   "foo",
	}.ApplyNoError(t)
}

// https://github.com/databricks/terraform-provider-databricks/issues/599
func TestRefreshOnRunningClusterWithFailedLibraryUninstallsIt(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceCluster(),
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/clusters/get?cluster_id=foo",
				Response: compute.ClusterDetails{
					State: compute.StateRunning,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=foo",
				Response: compute.ClusterLibraryStatuses{
					ClusterId: "foo",
					LibraryStatuses: []compute.LibraryFullStatus{
						{
							Library: &compute.Library{
								Jar: "foo.bar",
							},
							Status:   compute.LibraryInstallStatusFailed,
							Messages: []string{"fails for the test"},
						},
						{
							Library: &compute.Library{
								Whl: "bar.whl",
							},
							Status: compute.LibraryInstallStatusInstalled,
						},
					},
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/libraries/uninstall",
				ExpectedRequest: compute.UninstallLibraries{
					ClusterId: "foo",
					Libraries: []compute.Library{
						{
							Jar: "foo.bar",
						},
					},
				},
			},
			nothingPinned,
		},
		Read: true,
		ID:   "foo",
	}.ApplyNoError(t)
}

func TestResourceClusterUpdate_LocalSsdCount(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			nothingPinned,
			{
				Method:       "GET",
				Resource:     "/api/2.1/clusters/get?cluster_id=abc",
				ReuseRequest: true,
				Response: compute.ClusterDetails{
					ClusterId:              "abc",
					NumWorkers:             100,
					ClusterName:            "Non Autoscaling Cluster",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  compute.StateTerminated,
					GcpAttributes: &compute.GcpAttributes{
						LocalSsdCount: 2,
					},
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.1/clusters/edit",
				ExpectedRequest: compute.ClusterDetails{
					AutoterminationMinutes: 15,
					ClusterId:              "abc",
					NumWorkers:             100,
					ClusterName:            "Non Autoscaling Cluster",
					SparkVersion:           "7.1-scala12",
					NodeTypeId:             "i3.xlarge",
					GcpAttributes: &compute.GcpAttributes{
						LocalSsdCount: 0,
					},
				},
			},
		},
		ID:       "abc",
		Update:   true,
		Resource: ResourceCluster(),
		InstanceState: map[string]string{
			"autotermination_minutes": "15",
			"cluster_name":            "Shared Autoscaling",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             "100",
			"gcp_attributes": `"{
				local_ssd_count = 2
			}"`,
		},
		HCL: `
		autotermination_minutes = 15,
		cluster_name =            "Non Autoscaling Cluster"
		spark_version =           "7.1-scala12"
		node_type_id =            "i3.xlarge"
		num_workers =             100
		gcp_attributes = {
			local_ssd_count = 0
		},
		`,
	}.Apply(t)

	assert.NoError(t, err)
}
