package databricks

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func testGetAwsAttributes(attributesMap map[string]string) string {
	var awsAttr bytes.Buffer
	awsAttr.WriteString("aws_attributes {\n")
	for attr, value := range attributesMap {
		awsAttr.WriteString(fmt.Sprintf("%s = \"%s\"\n", attr, value))
	}
	awsAttr.WriteString("}")
	return awsAttr.String()
}

func testGetClusterInstancePoolConfig(instancePoolID string) string {
	if reflect.ValueOf(instancePoolID).IsZero() {
		return ""
	}
	return fmt.Sprintf("instance_pool_id = \"%s\"\n", instancePoolID)
}

func testDefaultZones() string {
	return `data "databricks_zones" "default_zones" {}`
}

func testDefaultAwsInstancePoolResource(awsAttributes, name string) string {
	return fmt.Sprintf(`
resource "databricks_instance_pool" "my_pool" {
  instance_pool_name = "%s"
  min_idle_instances = 0
  max_capacity = 5
  node_type_id = "i3.xlarge"
  %s
  idle_instance_autotermination_minutes = 10
  disk_spec {
    ebs_volume_type = "GENERAL_PURPOSE_SSD"
    disk_size = 80
    disk_count = 1
  }
}
`, name, awsAttributes)
}

func testDefaultClusterResource(instancePool, awsAttributes string) string {
	return fmt.Sprintf(`
	resource "databricks_cluster" "test_cluster" {
		cluster_name = "test-cluster-browser"
		%s
		spark_version = "6.6.x-scala2.11"
		autoscale {
		min_workers = 1
		max_workers = 2
		}
		%s
		autotermination_minutes = 10
		spark_conf = {
		"spark.databricks.cluster.profile" = "serverless"
		"spark.databricks.repl.allowedLanguages" = "sql,python,r"
		"spark.hadoop.fs.s3a.canned.acl" = "BucketOwnerFullControl"
		"spark.hadoop.fs.s3a.acl.default" = "BucketOwnerFullControl"
		}
		enable_elastic_disk = true
		custom_tags = {
		"ResourceClass" = "Serverless"
		}
	}`, instancePool, awsAttributes)
}

func TestAwsAccClusterResource_ValidatePlan(t *testing.T) {
	// TODO: refactor for common instance pool & AZ CLI
	awsAttrNoZoneID := map[string]string{}
	awsAttrInstanceProfile := map[string]string{
		"instance_profile_arn": "my_instance_profile_arn",
	}
	instancePoolLine := testGetClusterInstancePoolConfig("demo_instance_pool_id")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             testDefaultClusterResource(instancePoolLine, testGetAwsAttributes(awsAttrNoZoneID)),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config:             testDefaultClusterResource(instancePoolLine, testGetAwsAttributes(awsAttrInstanceProfile)),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAwsAccClusterResource_CreateClusterViaInstancePool(t *testing.T) {
	awsAttrInstancePool := map[string]string{
		"zone_id":      "${data.databricks_zones.default_zones.default_zone}",
		"availability": "SPOT",
	}
	randomInstancePoolName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	randomStr := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	instanceProfile := fmt.Sprintf("arn:aws:iam::999999999999:instance-profile/%s", randomStr)
	var clusterInfo model.ClusterInfo
	awsAttrCluster := map[string]string{
		"instance_profile_arn": "${databricks_instance_profile.my_instance_profile.id}",
	}
	instancePoolLine := testGetClusterInstancePoolConfig("${databricks_instance_pool.my_pool.id}")
	resourceConfig := testDefaultZones() +
		testAWSDatabricksInstanceProfile(instanceProfile) +
		testDefaultAwsInstancePoolResource(testGetAwsAttributes(awsAttrInstancePool), randomInstancePoolName) +
		testDefaultClusterResource(instancePoolLine, "")

	resourceInstanceProfileConfig := testDefaultZones() +
		testAWSDatabricksInstanceProfile(instanceProfile) +
		testDefaultAwsInstancePoolResource(testGetAwsAttributes(awsAttrInstancePool), randomInstancePoolName) +
		testDefaultClusterResource(instancePoolLine, testGetAwsAttributes(awsAttrCluster))

	resourceEmptyAttrConfig := testDefaultZones() +
		testAWSDatabricksInstanceProfile(instanceProfile) +
		testDefaultAwsInstancePoolResource(testGetAwsAttributes(awsAttrInstancePool), randomInstancePoolName) +
		testDefaultClusterResource(instancePoolLine, "aws_attributes {}")

	resource.Test(t, resource.TestCase{
		IsUnitTest: debugIfCloudEnvSet(),
		Providers:  testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: resourceConfig,
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testClusterExistsAndTerminateForFutureTests("databricks_cluster.test_cluster", &clusterInfo, t),
				),
			},
			{
				Config: resourceInstanceProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testClusterExistsAndTerminateForFutureTests("databricks_cluster.test_cluster", &clusterInfo, t),
				),
			},
			{
				Config: resourceEmptyAttrConfig,
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testClusterExistsAndTerminateForFutureTests("databricks_cluster.test_cluster", &clusterInfo, t),
				),
				ExpectNonEmptyPlan: true,
				Destroy:            true,
			},
			{
				Config: "",
			},
		},
	})
}

func testClusterExistsAndTerminateForFutureTests(n string, cluster *model.ClusterInfo, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := testAccProvider.Meta().(*service.DatabricksClient)
		resp, err := conn.Clusters().Get(rs.Primary.ID)
		if err != nil {
			return err
		}
		return conn.Clusters().Terminate(resp.ClusterID)
	}
}

func TestResourceClusterCreate(t *testing.T) {
	d, err := ResourceFixture{
		Fixtures: []HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/clusters/create",
				ExpectedRequest: model.Cluster{
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeID:             "i3.xlarge",
					AutoterminationMinutes: 15,
				},
				Response: model.ClusterInfo{
					ClusterID: "abc",
					State:     model.ClusterStateRunning,
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=abc",
				Response: model.ClusterInfo{
					ClusterID:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeID:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  model.ClusterStateRunning,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: model.ClusterLibraryStatuses{
					LibraryStatuses: []model.LibraryStatus{},
				},
			},
		},
		Create:   true,
		Resource: resourceCluster(),
		State: map[string]interface{}{
			"autotermination_minutes": 15,
			"cluster_name":            "Shared Autoscaling",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             100,
		},
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterCreate_WithLibraries(t *testing.T) {
	d, err := ResourceFixture{
		Fixtures: []HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/clusters/create",
				ExpectedRequest: model.Cluster{
					NumWorkers:             100,
					SparkVersion:           "7.1-scala12",
					NodeTypeID:             "i3.xlarge",
					AutoterminationMinutes: 60,
				},
				Response: model.ClusterInfo{
					ClusterID: "abc",
					State:     model.ClusterStateRunning,
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=abc",
				Response: model.ClusterInfo{
					ClusterID:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeID:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  model.ClusterStateRunning,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/libraries/install",
				ExpectedRequest: model.ClusterLibraryList{
					ClusterID: "abc",
					Libraries: []model.Library{
						{
							Pypi: &model.PyPi{
								Package: "seaborn==1.2.4",
							},
						},
						{
							Whl: "dbfs://baz.whl",
						},
						{
							Maven: &model.Maven{
								Coordinates: "foo:bar:baz:0.1.0",
								Exclusions:  []string{"org.apache:flink:base"},
								Repo:        "s3://maven-repo-in-s3/release",
							},
						},
						{
							Egg: "dbfs://bar.egg",
						},
						{
							Jar: "dbfs://foo.jar",
						},
						{
							Cran: &model.Cran{
								Package: "rkeops",
								Repo:    "internal",
							},
						},
					},
				},
			},
			{
				Method: "GET",
				// 1 of 3 requests
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: model.ClusterLibraryStatuses{
					LibraryStatuses: []model.LibraryStatus{
						{
							Library: &model.Library{
								Pypi: &model.PyPi{
									Package: "seaborn==1.2.4",
								},
							},
							Status: "PENDING",
						},
						{
							Library: &model.Library{
								Whl: "dbfs://baz.whl",
							},
							Status: "INSTALLED",
						},
					},
				},
			},
			{
				Method: "GET",
				// 2 of 3 requests
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: model.ClusterLibraryStatuses{
					LibraryStatuses: []model.LibraryStatus{
						{
							Library: &model.Library{
								Pypi: &model.PyPi{
									Package: "seaborn==1.2.4",
								},
							},
							Status: "INSTALLED",
						},
						{
							Library: &model.Library{
								Whl: "dbfs://baz.whl",
							},
							Status: "INSTALLED",
						},
					},
				},
			},
			{
				Method: "GET",
				// 3 of 3 requests
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: model.ClusterLibraryStatuses{
					LibraryStatuses: []model.LibraryStatus{
						{
							Library: &model.Library{
								Pypi: &model.PyPi{
									Package: "seaborn==1.2.4",
								},
							},
							Status: "INSTALLED",
						},
						{
							Library: &model.Library{
								Whl: "dbfs://baz.whl",
							},
							Status: "INSTALLED",
						},
					},
				},
			},
		},
		Create:   true,
		Resource: resourceCluster(),
		HCL: `num_workers = 100
		spark_version = "7.1-scala12"
		node_type_id = "i3.xlarge"

		libraries {
			jar = "dbfs://foo.jar"
		}

		libraries {
			egg = "dbfs://bar.egg"
		}

		libraries {
			whl = "dbfs://baz.whl"
		}

		libraries {
			pypi {
				package = "seaborn==1.2.4"
			}
		}

		libraries {
			maven {
				coordinates = "foo:bar:baz:0.1.0"
				repo = "s3://maven-repo-in-s3/release"
				exclusions = [
					"org.apache:flink:base"
				]
			}
		}

		libraries {
			cran {
				package = "rkeops"
				repo = "internal"
			}
		}`,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterCreate_Error(t *testing.T) {
	d, err := ResourceFixture{
		Fixtures: []HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/clusters/create",
				Response: service.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Create:   true,
		Resource: resourceCluster(),
		State: map[string]interface{}{
			"autotermination_minutes": 15,
			"cluster_name":            "Shared Autoscaling",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             100,
		},
	}.Apply(t)
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceClusterRead(t *testing.T) {
	d, err := ResourceFixture{
		Fixtures: []HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=abc",
				Response: model.ClusterInfo{
					ClusterID:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeID:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  model.ClusterStateRunning,
					AutoScale: &model.AutoScale{
						MaxWorkers: 4,
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: model.ClusterLibraryStatuses{
					LibraryStatuses: []model.LibraryStatus{
						{
							Library: &model.Library{
								Pypi: &model.PyPi{
									Package: "requests",
								},
							},
							Status: "INSTALLED",
						},
					},
				},
			},
		},
		Resource: resourceCluster(),
		Read:     true,
		ID:       "abc",
		New:      true,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, 15, d.Get("autotermination_minutes"))
	assert.Equal(t, "Shared Autoscaling", d.Get("cluster_name"))
	assert.Equal(t, "i3.xlarge", d.Get("node_type_id"))
	assert.Equal(t, 4, d.Get("autoscale.0.max_workers"))
	assert.Equal(t, "requests", d.Get("libraries.754562683.pypi.0.package"))
	assert.Equal(t, "RUNNING", d.Get("state"))

	for k, v := range d.State().Attributes {
		fmt.Printf("assert.Equal(t, %#v, d.Get(%#v))\n", v, k)
	}
}

func TestResourceClusterRead_NotFound(t *testing.T) {
	d, err := ResourceFixture{
		Fixtures: []HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=abc",
				Response: service.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: resourceCluster(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceClusterRead_Error(t *testing.T) {
	d, err := ResourceFixture{
		Fixtures: []HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=abc",
				Response: service.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: resourceCluster(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id(), "Id should not be empty for error reads")
}

func TestResourceClusterUpdate(t *testing.T) {
	d, err := ResourceFixture{
		Fixtures: []HTTPFixture{
			{
				Method:       "GET",
				Resource:     "/api/2.0/clusters/get?cluster_id=abc",
				ReuseRequest: true,
				Response: model.ClusterInfo{
					ClusterID:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeID:             "i3.xlarge",
					AutoterminationMinutes: 15,
					State:                  model.ClusterStateRunning,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/clusters/start",
				ExpectedRequest: model.ClusterID{
					ClusterID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: model.ClusterLibraryStatuses{
					LibraryStatuses: []model.LibraryStatus{},
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/clusters/edit",
				ExpectedRequest: model.Cluster{
					AutoterminationMinutes: 15,
					ClusterID:              "abc",
					NumWorkers:             100,
					ClusterName:            "Shared Autoscaling",
					SparkVersion:           "7.1-scala12",
					NodeTypeID:             "i3.xlarge",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: model.ClusterLibraryStatuses{
					LibraryStatuses: []model.LibraryStatus{},
				},
			},
		},
		ID:       "abc",
		Update:   true,
		Resource: resourceCluster(),
		State: map[string]interface{}{
			"autotermination_minutes": 15,
			"cluster_name":            "Shared Autoscaling",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             100,
		},
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should be the same as in reading")
}

func TestResourceClusterUpdate_LibrariesChangeOnTerminatedCluster(t *testing.T) {
	terminated := HTTPFixture{
		Method:   "GET",
		Resource: "/api/2.0/clusters/get?cluster_id=abc",
		Response: model.ClusterInfo{
			ClusterID:    "abc",
			NumWorkers:   100,
			SparkVersion: "7.1-scala12",
			NodeTypeID:   "i3.xlarge",
			State:        model.ClusterStateTerminated,
			StateMessage: "Terminated for test reasons",
		},
	}
	newLibs := HTTPFixture{
		Method:   "GET",
		Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
		Response: model.ClusterLibraryStatuses{
			ClusterID: "abc",
			LibraryStatuses: []model.LibraryStatus{
				{
					Library: &model.Library{
						Jar: "dbfs://foo.jar",
					},
					Status: "INSTALLED",
				},
				{
					Library: &model.Library{
						Egg: "dbfs://bar.egg",
					},
					Status: "INSTALLED",
				},
			},
		},
	}
	d, err := ResourceFixture{
		Fixtures: []HTTPFixture{
			terminated, // 1 of ...
			{
				Method:   "POST",
				Resource: "/api/2.0/clusters/edit",
				ExpectedRequest: model.Cluster{
					AutoterminationMinutes: 60,
					ClusterID:              "abc",
					NumWorkers:             100,
					SparkVersion:           "7.1-scala12",
					NodeTypeID:             "i3.xlarge",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
				Response: model.ClusterLibraryStatuses{
					ClusterID: "abc",
					LibraryStatuses: []model.LibraryStatus{
						{
							Library: &model.Library{
								Egg: "dbfs://bar.egg",
							},
							Status: "INSTALLED",
						},
						{
							Library: &model.Library{
								Pypi: &model.PyPi{
									Package: "requests",
								},
							},
							Status: "INSTALLED",
						},
					},
				},
			},
			{ // start cluster before libs install
				Method:   "POST",
				Resource: "/api/2.0/clusters/start",
				ExpectedRequest: model.ClusterID{
					ClusterID: "abc",
				},
			},
			{ // 2 of ...
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=abc",
				Response: model.ClusterInfo{
					ClusterID:    "abc",
					NumWorkers:   100,
					SparkVersion: "7.1-scala12",
					NodeTypeID:   "i3.xlarge",
					State:        model.ClusterStateRunning,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/libraries/uninstall",
				ExpectedRequest: model.ClusterLibraryList{
					ClusterID: "abc",
					Libraries: []model.Library{
						{
							Pypi: &model.PyPi{
								Package: "requests",
							},
						},
					},
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/libraries/install",
				ExpectedRequest: model.ClusterLibraryList{
					ClusterID: "abc",
					Libraries: []model.Library{
						{
							Jar: "dbfs://foo.jar",
						},
					},
				},
			},
			newLibs,
			{
				Method:   "POST",
				Resource: "/api/2.0/clusters/delete",
				ExpectedRequest: model.ClusterID{
					ClusterID: "abc",
				},
			},
			terminated, // 3 of 4
			// read
			terminated, // 4 of 4
			newLibs,
		},
		ID:       "abc",
		Update:   true,
		Resource: resourceCluster(),
		HCL: `num_workers = 100
		spark_version = "7.1-scala12"
		node_type_id = "i3.xlarge"

		libraries {
			jar = "dbfs://foo.jar"
		}

		libraries {
			egg = "dbfs://bar.egg"
		}`,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should be the same as in reading")
}

func TestResourceClusterUpdate_Error(t *testing.T) {
	d, err := ResourceFixture{
		Fixtures: []HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=abc",
				Response: service.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		ID:       "abc",
		Update:   true,
		Resource: resourceCluster(),
		State: map[string]interface{}{
			"autotermination_minutes": 15,
			"cluster_name":            "Shared Autoscaling",
			"spark_version":           "7.1-scala12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             100,
		},
	}.Apply(t)
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterDelete(t *testing.T) {
	d, err := ResourceFixture{
		Fixtures: []HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/clusters/delete",
				ExpectedRequest: map[string]string{
					"cluster_id": "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=abc",
				Response: model.ClusterInfo{
					State: model.ClusterStateTerminated,
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/clusters/permanent-delete",
				ExpectedRequest: map[string]string{
					"cluster_id": "abc",
				},
			},
		},
		Resource: resourceCluster(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterDelete_Error(t *testing.T) {
	d, err := ResourceFixture{
		Fixtures: []HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/clusters/delete",
				Response: service.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: resourceCluster(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}
