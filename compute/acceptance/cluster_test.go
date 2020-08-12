package acceptance

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	. "github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"

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
	return "data \"databricks_zones\" \"default_zones\" {}\n"
}

func testDefaultInstancePoolResource(awsAttributes, name string) string {
	var nodeTypeIdStatement string
	var diskSpecTypeStatement string
	nodeTypeIdStatement = "node_type_id = \"m4.large\""
	diskSpecTypeStatement = "ebs_volume_type = \"GENERAL_PURPOSE_SSD\"\n"
	if strings.ToLower(os.Getenv("CLOUD_ENV")) == "azure" {
		nodeTypeIdStatement = "node_type_id = \"Standard_DS3_v2\""
		diskSpecTypeStatement = "azure_disk_volume_type = \"PREMIUM_LRS\"\n"
	}

	return fmt.Sprintf(`
resource "databricks_instance_pool" "my_pool" {
	instance_pool_name = "%s"
	min_idle_instances = 0
	max_capacity = 5
	%s
	enable_elastic_disk = true
	%s
	idle_instance_autotermination_minutes = 10
	disk_spec {
		%s
		disk_size = 80
		disk_count = 1
	}
}
`, name, nodeTypeIdStatement, awsAttributes, diskSpecTypeStatement)
}

func testDefaultClusterResource(instancePool, awsAttributes string) string {
	return fmt.Sprintf(`
	resource "databricks_cluster" "test_cluster" {
		cluster_name = "test-cluster-instance-pool-test"
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

func TestAwsAccClusterResource_ValidatePlan(t *testing.T) {
	// TODO: refactor for common instance pool & AZ CLI
	awsAttrNoZoneID := map[string]string{}
	awsAttrInstanceProfile := map[string]string{
		"instance_profile_arn": "my_instance_profile_arn",
	}
	instancePoolLine := testGetClusterInstancePoolConfig("demo_instance_pool_id")
	resource.Test(t, resource.TestCase{
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
	var clusterInfo ClusterInfo
	awsAttrCluster := map[string]string{
		"instance_profile_arn": "${databricks_instance_profile.my_instance_profile.id}",
	}
	instancePoolLine := testGetClusterInstancePoolConfig("${databricks_instance_pool.my_pool.id}")
	resourceConfig := testDefaultZones() +
		testAWSDatabricksInstanceProfile(instanceProfile) +
		testDefaultInstancePoolResource(testGetAwsAttributes(awsAttrInstancePool), randomInstancePoolName) +
		testDefaultClusterResource(instancePoolLine, "aws_attributes {}")

	resourceInstanceProfileConfig := testDefaultZones() +
		testAWSDatabricksInstanceProfile(instanceProfile) +
		testDefaultInstancePoolResource(testGetAwsAttributes(awsAttrInstancePool), randomInstancePoolName) +
		testDefaultClusterResource(instancePoolLine, testGetAwsAttributes(awsAttrCluster))

	resourceEmptyAttrConfig := testDefaultZones() +
		testAWSDatabricksInstanceProfile(instanceProfile) +
		testDefaultInstancePoolResource(testGetAwsAttributes(awsAttrInstancePool), randomInstancePoolName) +
		testDefaultClusterResource(instancePoolLine, "aws_attributes {}")

	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: resourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testClusterExistsAndTerminateForFutureTests("databricks_cluster.test_cluster", &clusterInfo, t),
				),
			},
			{
				Config: resourceInstanceProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testClusterExistsAndTerminateForFutureTests("databricks_cluster.test_cluster", &clusterInfo, t),
				),
			},
			{
				Config: resourceEmptyAttrConfig,
				Check: resource.ComposeTestCheckFunc(
					testClusterExistsAndTerminateForFutureTests("databricks_cluster.test_cluster", &clusterInfo, t),
				),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: resourceEmptyAttrConfig,
				Check: resource.ComposeTestCheckFunc(
					testClusterExistsAndTerminateForFutureTests("databricks_cluster.test_cluster", &clusterInfo, t),
				),
			},
		},
	})
}

func TestAzureAccClusterResource_CreateClusterViaInstancePool(t *testing.T) {
	randomInstancePoolName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	var clusterInfo ClusterInfo
	instancePoolLine := testGetClusterInstancePoolConfig("${databricks_instance_pool.my_pool.id}")
	resourceConfig :=
		testDefaultInstancePoolResource("", randomInstancePoolName) +
			testDefaultClusterResource(instancePoolLine, "")

	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: resourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testClusterExistsAndTerminateForFutureTests("databricks_cluster.test_cluster", &clusterInfo, t),
				),
			},
		},
	})
}

func testClusterExistsAndTerminateForFutureTests(n string, cluster *ClusterInfo, t *testing.T) resource.TestCheckFunc {
	return acceptance.ResourceCheck(n, func(client *common.DatabricksClient, id string) error {
		clusters := NewClustersAPI(client)
		_, err := clusters.Get(id)
		if err != nil {
			return err
		}
		return clusters.Terminate(id)
	})
}

func testAWSDatabricksInstanceProfile(instanceProfile string) string {
	return fmt.Sprintf(`
		resource "databricks_instance_profile" "my_instance_profile" {
			instance_profile_arn = "%s"
			skip_validation = true
		}
		`, instanceProfile)
}
