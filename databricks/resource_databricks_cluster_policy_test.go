package databricks

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/stretchr/testify/assert"

	"github.com/databrickslabs/databricks-terraform/client/service"
)

func TestAccClusterPolicyResourceFullLifecycle(t *testing.T) {
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// create a resource
				Config: testExternalMetastore(randomName),
				Check: resource.ComposeTestCheckFunc(
					epoch.ResourceCheck("databricks_cluster_policy.external_metastore",
						func(client *service.DatabricksClient, id string) error {
							policy, err := client.ClusterPolicies().Get(id)
							assert.NoError(t, err)
							if policy.Definition == "" {
								return fmt.Errorf("Empty policy definition found")
							}
							return nil
						}),
					resource.TestCheckResourceAttr("databricks_cluster_policy.external_metastore",
						"name", fmt.Sprintf("Terraform policy %s", randomName)),
				),
			},
			{
				// add add the name for it
				Config: testExternalMetastore(randomName + ": UPDATED"),
				Check: resource.TestCheckResourceAttr("databricks_cluster_policy.external_metastore",
					"name", fmt.Sprintf("Terraform policy %s", randomName+": UPDATED")),
			},
			{
				Config:  testExternalMetastore(randomName + ": UPDATED"),
				Destroy: true,
				Check: epoch.ResourceCheck("databricks_cluster_policy.external_metastore",
					func(client *service.DatabricksClient, id string) error {
						resp, err := client.ClusterPolicies().Get(id)
						if err == nil {
							return fmt.Errorf("Resource must have been deleted but: %v", resp)
						}
						return nil
					}),
			},
			{
				// and create it again
				Config: testExternalMetastore(randomName + ": UPDATED"),
				Check: epoch.ResourceCheck("databricks_cluster_policy.external_metastore",
					func(client *service.DatabricksClient, id string) error {
						_, err := client.ClusterPolicies().Get(id)
						if err != nil {
							return err
						}
						return nil
					}),
			},
		},
	})
}

func testExternalMetastore(name string) string {
	return fmt.Sprintf(`
	resource "databricks_cluster_policy" "external_metastore" {
		name = "Terraform policy %s"
		definition = jsonencode({
			"spark_conf.spark.hadoop.javax.jdo.option.ConnectionURL": {
				"type": "fixed",
				"value": "jdbc:sqlserver://<jdbc-url>"
			},
			"spark_conf.spark.hadoop.javax.jdo.option.ConnectionDriverName": {
				"type": "fixed",
				"value": "com.microsoft.sqlserver.jdbc.SQLServerDriver"
			},
			"spark_conf.spark.databricks.delta.preview.enabled": {
				"type": "fixed",
				"value": true
			},
			"spark_conf.spark.hadoop.javax.jdo.option.ConnectionUserName": {
				"type": "fixed",
				"value": "<metastore-user>"
			},
			"spark_conf.spark.hadoop.javax.jdo.option.ConnectionPassword": {
				"type": "fixed",
				"value": "<metastore-password>"
			}
		  })
	}`, name)
}
