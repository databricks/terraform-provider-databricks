package databricks

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/stretchr/testify/assert"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
)

func TestAccClusterPolicyResourceFullLifecycle(t *testing.T) {
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resource.Test(t, resource.TestCase{
		IsUnitTest: debugIfCloudEnvSet(),
		Providers:  testAccProviders,
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

func TestResourceClusterPolicyRead(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/policies/clusters/get?policy_id=abc",
			Response: model.ClusterPolicy{
				PolicyID:           "abc",
				Name:               "Dummy",
				Definition:         "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
				CreatedAtTimeStamp: 0,
			},
		},
	}, resourceClusterPolicy, nil, actionWithID("abc", resourceClusterPolicyRead))
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, "Dummy", d.Get("name"))
	assert.Equal(t, "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}", d.Get("definition"))
	assert.Equal(t, "abc", d.Get("policy_id"))
}

func TestResourceClusterPolicyRead_NotFound(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{ // read log output for correct url...
			Method:   "GET",
			Resource: "/api/2.0/policies/clusters/get?policy_id=abc",
			Response: service.APIErrorBody{
				ErrorCode: "NOT_FOUND",
				Message:   "Item not found",
			},
			Status: 404,
		},
	}, resourceClusterPolicy, nil, actionWithID("abc", resourceClusterPolicyRead))
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceClusterPolicyRead_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{ // read log output for correct url...
			Method:   "GET",
			Resource: "/api/2.0/policies/clusters/get?policy_id=abc",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceClusterPolicy, nil, actionWithID("abc", resourceClusterPolicyRead))
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id(), "Id should not be empty for error reads")
}

func TestResourceClusterPolicyCreate(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/policies/clusters/create",
			ExpectedRequest: model.ClusterPolicy{
				Name:               "Dummy",
				Definition:         "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
				CreatedAtTimeStamp: 0,
			},
			Response: model.ClusterPolicy{
				PolicyID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/policies/clusters/get?policy_id=abc",
			Response: model.ClusterPolicy{
				PolicyID:           "abc",
				Name:               "Dummy",
				Definition:         "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
				CreatedAtTimeStamp: 0,
			},
		},
	}, resourceClusterPolicy, map[string]interface{}{
		"definition": `{"spark_conf.foo": {"type": "fixed", "value": "bar"}}`,
		"name":       "Dummy",
	}, resourceClusterPolicyCreate)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterPolicyCreate_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/policies/clusters/create",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceClusterPolicy, map[string]interface{}{
		"definition": `{"spark_conf.foo": {"type": "fixed", "value": "bar"}}`,
		"name":       "Dummy",
	}, resourceClusterPolicyCreate)
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id())
}

func TestResourceClusterPolicyUpdate(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/policies/clusters/edit",
			ExpectedRequest: model.ClusterPolicy{
				PolicyID:           "abc",
				Name:               "Dummy Updated",
				Definition:         "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
				CreatedAtTimeStamp: 0,
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/policies/clusters/get?policy_id=abc",
			Response: model.ClusterPolicy{
				PolicyID:           "abc",
				Name:               "Dummy Updated",
				Definition:         "{\"spark_conf.foo\": {\"type\": \"fixed\", \"value\": \"bar\"}}",
				CreatedAtTimeStamp: 0,
			},
		},
	}, resourceClusterPolicy, map[string]interface{}{
		"definition": `{"spark_conf.foo": {"type": "fixed", "value": "bar"}}`,
		"name":       "Dummy Updated",
	}, actionWithID("abc", resourceClusterPolicyUpdate))
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterPolicyUpdate_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/policies/clusters/edit",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceClusterPolicy, map[string]interface{}{
		"definition": `{"spark_conf.foo": {"type": "fixed", "value": "bar"}}`,
		"name":       "Dummy Updated",
	}, actionWithID("abc", resourceClusterPolicyUpdate))
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterPolicyDelete(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/policies/clusters/delete",
			ExpectedRequest: map[string]string{
				"policy_id": "abc",
			},
		},
	}, resourceClusterPolicy, nil, actionWithID("abc", resourceClusterPolicyDelete))
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestResourceClusterPolicyDelete_Error(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/policies/clusters/delete",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourceClusterPolicy, nil, actionWithID("abc", resourceClusterPolicyDelete))
	assertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}
