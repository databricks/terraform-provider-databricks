package acceptance

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	. "github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

type cloudSpecificHCLStatements struct {
	nodeTypeId   string
	diskSpecType string
}

var (
	azureClusterDefaults = cloudSpecificHCLStatements{
		nodeTypeId:   "node_type_id = \"Standard_DS3_v2\"",
		diskSpecType: "azure_disk_volume_type = \"PREMIUM_LRS\"\n",
	}
	awsClusterDefaults = cloudSpecificHCLStatements{
		nodeTypeId:   "node_type_id = \"i3.xlarge\"",
		diskSpecType: "ebs_volume_type = \"GENERAL_PURPOSE_SSD\"\n",
	}
)

func getCloudSpecificHCLStatements() cloudSpecificHCLStatements {
	switch strings.ToLower(os.Getenv("CLOUD_ENV")) {
	case "azure":
		return azureClusterDefaults
	default:
		return awsClusterDefaults
	}
}

func testDefaultZones() string {
	return "data \"databricks_zones\" \"default_zones\" {}\n"
}

type instancePoolHCLBuilder struct {
	Name          string
	identifier    string
	nodeTypeId    string
	diskSpecType  string
	awsAttributes string
}

func newInstancePoolHCLBuilder(name string) *instancePoolHCLBuilder {
	return &instancePoolHCLBuilder{
		Name:       name,
		identifier: name,
	}
}

func (i *instancePoolHCLBuilder) withCloudEnv() *instancePoolHCLBuilder {
	cloudHCLStatements := getCloudSpecificHCLStatements()
	i.nodeTypeId = cloudHCLStatements.nodeTypeId
	i.diskSpecType = cloudHCLStatements.diskSpecType
	return i
}

func (i *instancePoolHCLBuilder) build() string {
	return fmt.Sprintf(`
resource "databricks_instance_pool" "%[1]s" {
	instance_pool_name = "tf-instance-pool-%[1]s"
	min_idle_instances = 0
	max_capacity = 5
	%[2]s
	enable_elastic_disk = true
	%[3]s
	idle_instance_autotermination_minutes = 10
	disk_spec {
		%[4]s
		disk_size = 80
		disk_count = 1
	}
}
`, i.Name, i.nodeTypeId, i.awsAttributes, i.diskSpecType)
}

func getAwsAttributes(attributesMap map[string]string) string {
	if os.Getenv("CLOUD_ENV") != "AWS" {
		return ""
	}
	var awsAttr bytes.Buffer
	awsAttr.WriteString("aws_attributes {\n")
	for attr, value := range attributesMap {
		awsAttr.WriteString(fmt.Sprintf("%s = \"%s\"\n", attr, value))
	}
	awsAttr.WriteString("\t}")
	return awsAttr.String()
}

func (i *instancePoolHCLBuilder) withAwsAttributes(attributesMap map[string]string) *instancePoolHCLBuilder {
	i.awsAttributes = getAwsAttributes(attributesMap)
	return i
}

func getCommonLibraries() string {
	return `
	library_maven {
		coordinates = "org.jsoup:jsoup:1.7.2"
		repo = "https://mavencentral.org"
		exclusions = ["slf4j:slf4j"]
	}
	library_pypi {
		package = "faker"
		repo = "https://pypi.org"
	}
	library {
		pypi {
			package = "networkx"
		}
	}
	library {
		maven {
			coordinates = "com.microsoft.azure:azure-eventhubs-spark_2.11:2.3.7"
		}
	}
`
}

type clusterHCLBuilder struct {
	Name          string
	awsAttributes string
	instancePool  string
	libraries     string
	nodeTypeId    string
	diskSpec      string
}

func newClusterHCLBuilder(name string) *clusterHCLBuilder {
	return &clusterHCLBuilder{Name: name}
}

func (c *clusterHCLBuilder) withAwsAttributes(attributesMap map[string]string) *clusterHCLBuilder {
	c.awsAttributes = getAwsAttributes(attributesMap)
	return c
}

func (c *clusterHCLBuilder) withInstancePool(instancePoolID string) *clusterHCLBuilder {
	if instancePoolID == "" {
		return c
	}
	c.instancePool = fmt.Sprintf("instance_pool_id = %s\n", instancePoolID)
	return c
}

func (c *clusterHCLBuilder) withDefaultLibraries() *clusterHCLBuilder {
	c.libraries = getCommonLibraries()
	return c
}

func (c *clusterHCLBuilder) withCloudDiskSpec() *clusterHCLBuilder {
	cloudHCLStatements := getCloudSpecificHCLStatements()
	c.diskSpec = fmt.Sprintf(`
	disk_spec {
		%s
		disk_size = 80
		disk_count = 1
	}
`, cloudHCLStatements.diskSpecType)
	return c
}

func (c *clusterHCLBuilder) build() string {
	return fmt.Sprintf(`
resource "databricks_cluster" "%[1]s" {
	cluster_name = "%[1]s"
	%[2]s
	spark_version = "6.6.x-scala2.11"
	autoscale {
		min_workers = 1
		max_workers = 2
	}
	%[3]s
	%[4]s
	%[5]s
	autotermination_minutes = 10
	spark_conf = {
		"spark.databricks.cluster.profile" = "serverless"
		"spark.databricks.repl.allowedLanguages" = "sql,python,r"
		"spark.databricks.delta.preview.enabled" = "true"
		"spark.hadoop.fs.s3a.canned.acl" = "BucketOwnerFullControl"
		"spark.hadoop.fs.s3a.acl.default" = "BucketOwnerFullControl"
	}
	custom_tags = {
		"ResourceClass" = "Serverless"
	}
}`, c.Name, c.instancePool, c.awsAttributes, c.libraries, c.nodeTypeId, c.diskSpec)
}

func TestAwsAccClusterResource_ValidatePlan(t *testing.T) {
	// TODO: refactor for common instance pool & AZ CLI
	planName := "my-plan-cluster"
	awsAttrNoZoneID := map[string]string{}
	awsAttrInstanceProfile := map[string]string{
		"instance_profile_arn": "my_instance_profile_arn",
	}
	awsConfigWithNoZoneId := newClusterHCLBuilder(planName).
		withAwsAttributes(awsAttrNoZoneID).
		withCloudDiskSpec().
		build()
	awsConfigWithInstanceProfile := newClusterHCLBuilder(planName).
		withAwsAttributes(awsAttrInstanceProfile).
		withCloudDiskSpec().
		build()
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config:             awsConfigWithNoZoneId,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config:             awsConfigWithInstanceProfile,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAwsAccClusterResource_CreateClusterViaInstancePool(t *testing.T) {
	randomInstancePoolSuffix := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	randomInstancePoolName := fmt.Sprintf("pool-%s", randomInstancePoolSuffix)
	randomInstancePoolInterpolation := fmt.Sprintf("databricks_instance_pool.%s.id", randomInstancePoolName)
	randomClusterSuffix := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	randomClusterName := fmt.Sprintf("cluster-%s", randomClusterSuffix)
	randomClusterId := fmt.Sprintf("databricks_cluster.%s", randomClusterName)
	awsAttrInstancePool := map[string]string{
		"zone_id":      "${data.databricks_zones.default_zones.default_zone}",
		"availability": "SPOT",
	}
	randomStr := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	instanceProfileRName := "my-tf-test-instance-profile"
	instanceProfile := fmt.Sprintf("arn:aws:iam::999999999999:instance-profile/tf-test-%s", randomStr)
	awsAttrCluster := map[string]string{
		"instance_profile_arn": fmt.Sprintf("${databricks_instance_profile.%s.id}", instanceProfileRName),
	}

	clusterNoInstanceProfileConfig := testDefaultZones() +
		testAWSDatabricksInstanceProfile(instanceProfile, instanceProfileRName) +
		newInstancePoolHCLBuilder(randomInstancePoolName).
			withAwsAttributes(awsAttrInstancePool).withCloudEnv().
			build() +
		newClusterHCLBuilder(randomClusterName).
			withAwsAttributes(nil).
			withInstancePool(randomInstancePoolInterpolation).
			build()

	clusterWithInstanceProfileConfig := testDefaultZones() +
		testAWSDatabricksInstanceProfile(instanceProfile, instanceProfileRName) +
		newInstancePoolHCLBuilder(randomInstancePoolName).
			withAwsAttributes(awsAttrInstancePool).withCloudEnv().
			build() +
		newClusterHCLBuilder(randomClusterName).
			withAwsAttributes(awsAttrCluster).
			withInstancePool(randomInstancePoolInterpolation).
			build()

	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: clusterNoInstanceProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testClusterCheckAndTerminateForFutureTests(randomClusterId, t),
				),
			},
			{
				Config: clusterWithInstanceProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testClusterCheckAndTerminateForFutureTests(randomClusterId, t),
				),
			},
			{
				Config: clusterNoInstanceProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testClusterCheckAndTerminateForFutureTests(randomClusterId, t),
				),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: clusterNoInstanceProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testClusterCheckAndTerminateForFutureTests(randomClusterId, t),
				),
			},
		},
	})
}

func TestAzureAccClusterResource_CreateClusterViaInstancePool(t *testing.T) {
	randomInstancePoolName := fmt.Sprintf("pool_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	randomClusterName := fmt.Sprintf("cluster_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	defaultAzureInstancePoolClusterTest := newInstancePoolHCLBuilder(randomInstancePoolName).withCloudEnv().build() +
		newClusterHCLBuilder(randomClusterName).withInstancePool(
			fmt.Sprintf("databricks_instance_pool.%s.id", randomInstancePoolName)).build()
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: defaultAzureInstancePoolClusterTest,
				Check: resource.ComposeTestCheckFunc(
					testClusterCheckAndTerminateForFutureTests(
						fmt.Sprintf("databricks_cluster.%s", randomClusterName), t),
				),
			},
		},
	})
}

func TestAccClusterResource_CreateClusterWithLibraries(t *testing.T) {
	if os.Getenv("CLOUD_ENV") == "" {
		return
	}
	randomName := fmt.Sprintf("cluster-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	randomClusterID := fmt.Sprintf("databricks_cluster.%s", randomName)
	var clusterInfo ClusterInfo

	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: newClusterHCLBuilder(randomName).
					withInstancePool(fmt.Sprintf("%#v", CommonInstancePoolID())).
					withAwsAttributes(nil).
					build(),
				Check: resource.ComposeTestCheckFunc(
					testClusterCheckExists(randomClusterID, &clusterInfo, t),
				),
			},
			{
				PreConfig: func() {
					client := common.CommonEnvironmentClient()
					ctx := context.Background()
					err := NewClustersAPI(ctx, client).Terminate(clusterInfo.ClusterID)
					assert.NoError(t, err)
				},
				Config: newClusterHCLBuilder(randomName).
					withInstancePool(fmt.Sprintf("%#v", CommonInstancePoolID())).
					withAwsAttributes(nil).
					withDefaultLibraries().
					build(),
			},
		},
	})
}

func testClusterCheckExists(n string, cluster *ClusterInfo, t *testing.T) resource.TestCheckFunc {
	ctx := context.Background()
	return acceptance.ResourceCheck(n, func(client *common.DatabricksClient, id string) error {
		clusters := NewClustersAPI(ctx, client)
		c, err := clusters.Get(id)
		*cluster = c
		return err
	})
}

func testClusterCheckAndTerminateForFutureTests(n string, t *testing.T) resource.TestCheckFunc {
	ctx := context.Background()
	return acceptance.ResourceCheck(n, func(client *common.DatabricksClient, id string) error {
		return NewClustersAPI(ctx, client).Terminate(id)
	})
}

func testAWSDatabricksInstanceProfile(instanceProfile string, name string) string {
	return fmt.Sprintf(`
		resource "databricks_instance_profile" "%s" {
			instance_profile_arn = "%s"
			skip_validation = true
		}
		`, name, instanceProfile)
}
