package databricks

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
)

func TestAccClusterPolicyResourceFullLifecycle(t *testing.T) {
	var policy model.ClusterPolicy
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// create a resource
				Config: testExternalMetastore(randomName),
				Check: resource.ComposeTestCheckFunc(
					testAccIDCallback(t, "databricks_cluster_policy.external_metastore",
						func(client *service.DBApiClient, id string) error {
							resp, err := client.ClusterPolicies().Get(id)
							if err != nil {
								return err
							}
							policy = *resp
							return nil
						}),
					func(s *terraform.State) error {
						if policy.Definition == "" {
							return fmt.Errorf("Empty policy definition found")
						}
						return nil
					},
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
				Check: testAccIDCallback(t, "databricks_cluster_policy.external_metastore",
					func(client *service.DBApiClient, id string) error {
						resp, err := client.ClusterPolicies().Get(id)
						if err == nil {
							return fmt.Errorf("Resource must have been deleted but: %v", resp)
						}
						return nil
					}),
			},
		},
	})
}

func testAccIDCallback(t *testing.T, name string, cb func(client *service.DBApiClient, id string) error) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}
		client := testAccProvider.Meta().(*service.DBApiClient)
		err := cb(client, rs.Primary.ID)
		return err
	}
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
