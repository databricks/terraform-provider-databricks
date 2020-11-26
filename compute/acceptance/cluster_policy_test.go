package acceptance

import (
	"context"
	"fmt"
	"testing"

	. "github.com/databrickslabs/databricks-terraform/compute"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

func TestAccClusterPolicyResourceFullLifecycle(t *testing.T) {
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				// create a resource
				Config: testExternalMetastore(randomName),
				Check: resource.ComposeTestCheckFunc(
					acceptance.ResourceCheck("databricks_cluster_policy.external_metastore",
						func(client *common.DatabricksClient, id string) error {
							policy, err := NewClusterPoliciesAPI(context.Background(), client).Get(id)
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
