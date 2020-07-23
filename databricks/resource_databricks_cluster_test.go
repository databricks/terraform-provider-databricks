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

func testGetClusterInstancePoolConfig(instancePoolId string) string {
	if reflect.ValueOf(instancePoolId).IsZero() {
		return ""
	}
	return fmt.Sprintf("instance_pool_id = \"%s\"\n", instancePoolId)
}

func testDefaultZones() string {
	return `
data "databricks_zones" "default_zones" {}
`
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
								  custom_tags = {
									"ResourceClass" = "Serverless"
								  }
								}`, instancePool, awsAttributes)
}

func TestAccAwsClusterResource_ValidatePlan(t *testing.T) {
	awsAttrNoZoneId := map[string]string{}
	awsAttrInstanceProfile := map[string]string{
		"instance_profile_arn": "my_instance_profile_arn",
	}
	instancePoolLine := testGetClusterInstancePoolConfig("demo_instance_pool_id")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             testDefaultClusterResource(instancePoolLine, testGetAwsAttributes(awsAttrNoZoneId)),
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

func TestAccAwsClusterResource_CreateClusterViaInstancePool(t *testing.T) {
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
		Providers: testAccProviders,
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
			},
		},
	})
}

func testDefaultAzureInstancePoolResource(awsAttributes, name string) string {
	return fmt.Sprintf(`
resource "databricks_instance_pool" "my_pool" {
  instance_pool_name = "%s"
  min_idle_instances = 0
  max_capacity = 5
  node_type_id = "Standard_DS3_v2"
  %s
  idle_instance_autotermination_minutes = 10
}
`, name, awsAttributes)
}

func TestAccAzureClusterResource_CreateClusterViaInstancePool(t *testing.T) {
	randomInstancePoolName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	randomStr := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	instanceProfile := fmt.Sprintf("arn:aws:iam::999999999999:instance-profile/%s", randomStr)
	var clusterInfo model.ClusterInfo

	instancePoolLine := testGetClusterInstancePoolConfig("${databricks_instance_pool.my_pool.id}")
	resourceConfig := testAWSDatabricksInstanceProfile(instanceProfile) +
		testDefaultAzureInstancePoolResource("", randomInstancePoolName) +
		testDefaultClusterResource(instancePoolLine, "")

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: resourceConfig,
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testClusterExistsAndTerminateForFutureTests("databricks_cluster.test_cluster", &clusterInfo, t),
				),
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
		conn := testAccProvider.Meta().(*service.DBApiClient)
		resp, err := conn.Clusters().Get(rs.Primary.ID)
		if err != nil {
			return err
		}
		err = conn.Clusters().Terminate(resp.ClusterID)
		if err != nil {
			return err
		}
		err = conn.Clusters().WaitForClusterTerminated(resp.ClusterID, 10, 10)
		if err != nil {
			return err
		}
		return nil
	}
}
