package databricks

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stretchr/testify/assert"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
)

var (
	TestingUser      = "ben"
	TestingAdminUser = "admin"
)

func TestPermissionsRead(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/preview/permissions/clusters/abc?",
			Response: model.ObjectACL{
				ObjectID:   "/clusters/abc",
				ObjectType: "clusters",
				AccessControlList: []*model.AccessControl{
					{
						UserName: &TestingUser,
						AllPermissions: []*model.Permission{
							{
								PermissionLevel: "CAN_READ",
								Inherited:       false,
							},
						},
					},
					{
						UserName: &TestingAdminUser,
						AllPermissions: []*model.Permission{
							{
								PermissionLevel: "CAN_MANAGE",
								Inherited:       false,
							},
						},
					},
				},
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/preview/scim/v2/Me?",
			Response: model.User{
				UserName: TestingAdminUser,
			},
		},
	}, resourcePermissions, nil, func(d *schema.ResourceData, c interface{}) error {
		d.SetId("/clusters/abc")
		return resourcePermissionsRead(d, c)
	})
	assert.NoError(t, err, err)
	assert.Equal(t, "/clusters/abc", d.Id())
	assert.Equal(t, TestingUser, d.Get("access_control.0.user_name"))
	assert.Equal(t, "CAN_READ", d.Get("access_control.0.permission_level"))
	assert.Equal(t, 1, d.Get("access_control.#"))
}

func TestPermissionsRead_some_error(t *testing.T) {
	_, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/preview/permissions/clusters/abc?",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourcePermissions, map[string]interface{}{},
		func(d *schema.ResourceData, c interface{}) error {
			d.SetId("/clusters/abc")
			return resourcePermissionsRead(d, c)
		})
	assert.Error(t, err)
}

func TestPermissionsRead_ErrorOnScimMe(t *testing.T) {
	_, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/preview/permissions/clusters/abc?",
			Response: model.ObjectACL{
				ObjectID:   "/clusters/abc",
				ObjectType: "clusters",
				AccessControlList: []*model.AccessControl{
					{
						UserName: &TestingUser,
						AllPermissions: []*model.Permission{
							{
								PermissionLevel: "CAN_READ",
								Inherited:       false,
							},
						},
					},
					{
						UserName: &TestingAdminUser,
						AllPermissions: []*model.Permission{
							{
								PermissionLevel: "CAN_MANAGE",
								Inherited:       false,
							},
						},
					},
				},
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/preview/scim/v2/Me?",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourcePermissions, map[string]interface{}{},
		func(d *schema.ResourceData, c interface{}) error {
			d.SetId("/clusters/abc")
			return resourcePermissionsRead(d, c)
		})
	assert.Error(t, err)
}

func TestPermissionsDelete(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:          http.MethodPut,
			Resource:        "/api/2.0/preview/permissions/clusters/abc",
			ExpectedRequest: model.ObjectACL{},
		},
	}, resourcePermissions, nil, func(d *schema.ResourceData, c interface{}) error {
		d.SetId("/clusters/abc")
		return resourcePermissionsDelete(d, c)
	})
	assert.NoError(t, err, err)
	assert.Equal(t, "/clusters/abc", d.Id())
}

func TestPermissionsDelete_error(t *testing.T) {
	_, err := ResourceTester(t, []HTTPFixture{
		{
			Method:          http.MethodPut,
			Resource:        "/api/2.0/preview/permissions/clusters/abc",
			ExpectedRequest: model.ObjectACL{},
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourcePermissions, nil, func(d *schema.ResourceData, c interface{}) error {
		d.SetId("/clusters/abc")
		return resourcePermissionsDelete(d, c)
	})
	assert.Error(t, err)
}

func TestPermissionsCreate_invalid(t *testing.T) {
	_, err := ResourceTester(t, []HTTPFixture{}, resourcePermissions,
		nil, resourcePermissionsCreate)
	assert.EqualError(t, err, "At least one type of resource identifiers must be set")
}

func TestPermissionsCreate_no_access_control(t *testing.T) {
	_, err := ResourceTester(t, []HTTPFixture{}, resourcePermissions,
		map[string]interface{}{
			"cluster_id": "abc",
		}, resourcePermissionsCreate)
	assert.EqualError(t, err, "Invalid config supplied. access_control: required field is not set")
}

func TestPermissionsCreate_conflicting_fields(t *testing.T) {
	_, err := ResourceTester(t, []HTTPFixture{}, resourcePermissions,
		map[string]interface{}{
			"cluster_id":    "abc",
			"notebook_path": "/Init",
			"access_control": []interface{}{
				map[string]interface{}{
					"user_name":        TestingUser,
					"permission_level": "CAN_READ",
				},
			},
		}, resourcePermissionsCreate)
	assert.EqualError(t, err, "Invalid config supplied. cluster_id: conflicts with notebook_path. notebook_path: conflicts with cluster_id")
}

func TestPermissionsCreate(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodPatch,
			Resource: "/api/2.0/preview/permissions/clusters/abc",
			ExpectedRequest: model.AccessControlChangeList{
				AccessControlList: []*model.AccessControlChange{
					{
						UserName:        &TestingUser,
						PermissionLevel: "CAN_READ",
					},
				},
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/preview/permissions/clusters/abc?",
			Response: model.ObjectACL{
				ObjectID:   "/clusters/abc",
				ObjectType: "clusters",
				AccessControlList: []*model.AccessControl{
					{
						UserName: &TestingUser,
						AllPermissions: []*model.Permission{
							{
								PermissionLevel: "CAN_READ",
								Inherited:       false,
							},
						},
					},
					{
						UserName: &TestingAdminUser,
						AllPermissions: []*model.Permission{
							{
								PermissionLevel: "CAN_MANAGE",
								Inherited:       false,
							},
						},
					},
				},
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/preview/scim/v2/Me?",
			Response: model.User{
				UserName: TestingAdminUser,
			},
		},
	}, resourcePermissions, map[string]interface{}{
		"cluster_id": "abc",
		"access_control": []interface{}{
			map[string]interface{}{
				"user_name":        TestingUser,
				"permission_level": "CAN_READ",
			},
		},
	}, resourcePermissionsCreate)
	assert.NoError(t, err, err)
	assert.Equal(t, TestingUser, d.Get("access_control.0.user_name"))
	assert.Equal(t, "CAN_READ", d.Get("access_control.0.permission_level"))
	assert.Equal(t, 1, d.Get("access_control.#"))
}

func TestPermissionsCreate_NotebookPath_NotExists(t *testing.T) {
	_, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/workspace/get-status?path=%2FDevelopment%2FInit",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourcePermissions, map[string]interface{}{
		"notebook_path": "/Development/Init",
		"access_control": []interface{}{
			map[string]interface{}{
				"user_name":        TestingUser,
				"permission_level": "CAN_USE",
			},
		},
	}, resourcePermissionsCreate)

	assert.Error(t, err)
}

func TestPermissionsCreate_NotebookPath(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/workspace/get-status?path=%2FDevelopment%2FInit",
			Response: model.WorkspaceObjectStatus{
				ObjectID:   988765,
				ObjectType: "NOTEBOOK",
			},
		},
		{
			Method:   http.MethodPatch,
			Resource: "/api/2.0/preview/permissions/notebooks/988765",
			ExpectedRequest: model.AccessControlChangeList{
				AccessControlList: []*model.AccessControlChange{
					{
						UserName:        &TestingUser,
						PermissionLevel: "CAN_USE",
					},
				},
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/preview/permissions/notebooks/988765?",
			Response: model.ObjectACL{
				ObjectID:   "/notebooks/988765",
				ObjectType: "notebooks",
				AccessControlList: []*model.AccessControl{
					{
						UserName: &TestingUser,
						AllPermissions: []*model.Permission{
							{
								PermissionLevel: "CAN_USE",
								Inherited:       false,
							},
						},
					},
					{
						UserName: &TestingAdminUser,
						AllPermissions: []*model.Permission{
							{
								PermissionLevel: "CAN_MANAGE",
								Inherited:       false,
							},
						},
					},
				},
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/preview/scim/v2/Me?",
			Response: model.User{
				UserName: TestingAdminUser,
			},
		},
	}, resourcePermissions, map[string]interface{}{
		"notebook_path": "/Development/Init",
		"access_control": []interface{}{
			map[string]interface{}{
				"user_name":        TestingUser,
				"permission_level": "CAN_USE",
			},
		},
	}, resourcePermissionsCreate)

	assert.NoError(t, err, err)
	assert.Equal(t, TestingUser, d.Get("access_control.0.user_name"))
	assert.Equal(t, "CAN_USE", d.Get("access_control.0.permission_level"))
	assert.Equal(t, 1, d.Get("access_control.#"))
}

func TestPermissionsCreate_error(t *testing.T) {
	_, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodPatch,
			Resource: "/api/2.0/preview/permissions/clusters/abc",
			Response: service.APIErrorBody{
				ErrorCode: "INVALID_REQUEST",
				Message:   "Internal error happened",
			},
			Status: 400,
		},
	}, resourcePermissions, map[string]interface{}{
		"cluster_id": "abc",
		"access_control": []interface{}{
			map[string]interface{}{
				"user_name":        TestingUser,
				"permission_level": "CAN_USE",
			},
		},
	}, resourcePermissionsCreate)
	if assert.Error(t, err) {
		if e, ok := err.(service.APIError); ok {
			assert.Equal(t, "INVALID_REQUEST", e.ErrorCode)
		}
	}
}

func TestAccDatabricksPermissionsResourceFullLifecycle(t *testing.T) {
	var permissions model.ObjectACL
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// create a resource
				Config: testClusterPolicyPermissions(randomName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("databricks_permissions.dummy_can_use",
						"object_type", "cluster-policy"),
					testAccIDCallback(t, "databricks_permissions.dummy_can_use",
						func(client *service.DBApiClient, id string) error {
							resp, err := client.Permissions().Read(id)
							if err != nil {
								return err
							}
							permissions = *resp
							assert.Len(t, permissions.AccessControlList, 3)
							return nil
						}),
				),
			},
			{
				Config: testClusterPolicyPermissionsSecondGroupAdded(randomName),
				Check: testAccIDCallback(t, "databricks_permissions.dummy_can_use",
					func(client *service.DBApiClient, id string) error {
						resp, err := client.Permissions().Read(id)
						if err != nil {
							return err
						}
						permissions = *resp
						assert.Len(t, permissions.AccessControlList, 3)
						return nil
					}),
			},
		},
	})
}

func testClusterPolicyPermissions(name string) string {
	return fmt.Sprintf(`
	resource "databricks_cluster_policy" "something_simple" {
		name = "Terraform Policy %[1]s"
		definition = jsonencode({
			"spark_conf.spark.hadoop.javax.jdo.option.ConnectionURL": {
				"type": "forbidden"
			}
		})
	}
	resource "databricks_scim_group" "dummy_group" {
		display_name = "Terraform Group %[1]s"
	}
	resource "databricks_permissions" "dummy_can_use" {
		cluster_policy_id = databricks_cluster_policy.something_simple.id
		access_control {
			group_name = databricks_scim_group.dummy_group.display_name
			permission_level = "CAN_USE"
		}
	}
	`, name)
}

func testClusterPolicyPermissionsSecondGroupAdded(name string) string {
	return fmt.Sprintf(`
	resource "databricks_cluster_policy" "something_simple" {
		name = "Terraform Policy %[1]s"
		definition = jsonencode({
			"spark_conf.spark.hadoop.javax.jdo.option.ConnectionURL": {
				"type": "forbidden"
			},
			"spark_conf.spark.secondkey": {
				"type": "forbidden"
			}
		})
	}
	resource "databricks_scim_group" "dummy_group" {
		display_name = "Terraform Group %[1]s"
	}
	resource "databricks_scim_group" "second_group" {
		display_name = "Terraform Second Group %[1]s"
	}
	resource "databricks_permissions" "dummy_can_use" {
		cluster_policy_id = databricks_cluster_policy.something_simple.id
		access_control {
			group_name = databricks_scim_group.dummy_group.display_name
			permission_level = "CAN_USE"
		}
		access_control {
			group_name = databricks_scim_group.second_group.display_name
			permission_level = "CAN_USE"
		}
	}
	`, name)
}

func TestAccNotebookPermissions(t *testing.T) {
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// create a resource
				Config: fmt.Sprintf(`
				resource "databricks_notebook" "dummy" {
					content = base64encode("# Databricks notebook source\nprint(1)")
					path = "/Beginning/Init"
					overwrite = true
					mkdirs = true
					language = "PYTHON"
					format = "SOURCE"
				}
				resource "databricks_scim_group" "dummy_group" {
					display_name = "Terraform Group %[1]s"
				}
				resource "databricks_permissions" "dummy_can_use" {
					directory_path = "/Beginning"
					access_control {
						group_name = databricks_scim_group.dummy_group.display_name
						permission_level = "CAN_MANAGE"
					}
				}
				`, randomName),
			},
		},
	})
}
