package sharing_test

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

const initialShareConfig = `
	resource "databricks_share" "this" {
		name    = "fixture-share"
		owner   = "account users"
		comment = "initial comment"
		object {
			name             = "main.default.first"
			comment          = "first object"
			data_object_type = "TABLE"
		}
		object {
			name             = "main.default.second"
			comment          = "second object"
			data_object_type = "TABLE"
		}
	}
`

const updatedShareConfig = `
	resource "databricks_share" "this" {
		name    = "fixture-share"
		owner   = "account users"
		comment = "updated comment"
		object {
			name             = "main.default.first"
			comment          = "updated first object"
			data_object_type = "TABLE"
		}
		object {
			name             = "main.default.third"
			comment          = "third object"
			data_object_type = "TABLE"
		}
	}
`

var initialShareObjects = []sharing.SharedDataObject{
	{Name: "main.default.first", Comment: "first object", DataObjectType: sharing.SharedDataObjectDataObjectTypeTable},
	{Name: "main.default.second", Comment: "second object", DataObjectType: sharing.SharedDataObjectDataObjectTypeTable},
}

var updatedShareObjects = []sharing.SharedDataObject{
	{Name: "main.default.first", Comment: "updated first object", DataObjectType: sharing.SharedDataObjectDataObjectTypeTable},
	{Name: "main.default.third", Comment: "third object", DataObjectType: sharing.SharedDataObjectDataObjectTypeTable},
}

func shareGetFixtures(count int, share sharing.ShareInfo) []qa.HTTPFixture {
	fixtures := make([]qa.HTTPFixture, count)
	for i := range fixtures {
		fixtures[i] = qa.HTTPFixture{
			Method:          "GET",
			Resource:        "/api/2.1/unity-catalog/shares/fixture-share?include_shared_data=true",
			ExpectedHeaders: map[string]string{"Authorization": "Bearer ..."},
			Response:        share,
		}
	}
	return fixtures
}

func TestShareLifecycleWithHTTPFixtures(t *testing.T) {
	initialShare := sharing.ShareInfo{
		Name:      "fixture-share",
		Owner:     "account users",
		Comment:   "initial comment",
		CreatedAt: 100,
		CreatedBy: "creator@example.com",
		Objects:   initialShareObjects,
	}
	updatedShare := sharing.ShareInfo{
		Name:      "fixture-share",
		Owner:     "account users",
		Comment:   "updated comment",
		CreatedAt: 100,
		CreatedBy: "creator@example.com",
		UpdatedAt: 200,
		UpdatedBy: "updater@example.com",
		Objects:   updatedShareObjects,
	}

	fixtures := []qa.HTTPFixture{
		{
			Method:       "GET",
			Resource:     "/.well-known/databricks-config",
			Response:     map[string]string{"host_type": "WORKSPACE_HOST", "workspace_id": "12345"},
			ReuseRequest: true,
		},
		{
			Method:          "POST",
			Resource:        "/api/2.1/unity-catalog/shares",
			ExpectedHeaders: map[string]string{"Authorization": "Bearer ..."},
			ExpectedRequest: sharing.CreateShare{Name: "fixture-share", Comment: "initial comment"},
			Response:        sharing.ShareInfo{Name: "fixture-share"},
		},
		{
			Method:          "PATCH",
			Resource:        "/api/2.1/unity-catalog/shares/fixture-share",
			ExpectedHeaders: map[string]string{"Authorization": "Bearer ..."},
			ExpectedRequest: sharing.UpdateShare{
				Owner: "account users",
				Updates: []sharing.SharedDataObjectUpdate{
					{Action: sharing.SharedDataObjectUpdateActionAdd, DataObject: &initialShareObjects[0]},
					{Action: sharing.SharedDataObjectUpdateActionAdd, DataObject: &initialShareObjects[1]},
				},
			},
			Response: initialShare,
		},
	}
	fixtures = append(fixtures, shareGetFixtures(5, initialShare)...)
	fixtures = append(fixtures,
		qa.HTTPFixture{
			Method:          "PATCH",
			Resource:        "/api/2.1/unity-catalog/shares/fixture-share",
			ExpectedHeaders: map[string]string{"Authorization": "Bearer ..."},
			ExpectedRequest: sharing.UpdateShare{Owner: "account users"},
			Response:        initialShare,
		},
		qa.HTTPFixture{
			Method:          "PATCH",
			Resource:        "/api/2.1/unity-catalog/shares/fixture-share",
			ExpectedHeaders: map[string]string{"Authorization": "Bearer ..."},
			ExpectedRequest: sharing.UpdateShare{
				Comment: "updated comment",
				Updates: []sharing.SharedDataObjectUpdate{
					{Action: sharing.SharedDataObjectUpdateActionRemove, DataObject: &initialShareObjects[1]},
					{Action: sharing.SharedDataObjectUpdateActionUpdate, DataObject: &updatedShareObjects[0]},
					{Action: sharing.SharedDataObjectUpdateActionAdd, DataObject: &updatedShareObjects[1]},
				},
			},
			Response: updatedShare,
		},
	)
	fixtures = append(fixtures, shareGetFixtures(3, updatedShare)...)
	fixtures = append(fixtures, qa.HTTPFixture{
		Method:          "DELETE",
		Resource:        "/api/2.1/unity-catalog/shares/fixture-share?",
		ExpectedHeaders: map[string]string{"Authorization": "Bearer ..."},
	})

	terraformtest.HTTPFixtures(t, fixtures,
		resource.TestStep{
			Config: initialShareConfig,
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("databricks_share.this", "id", "fixture-share"),
				resource.TestCheckResourceAttr("databricks_share.this", "name", "fixture-share"),
				resource.TestCheckResourceAttr("databricks_share.this", "comment", "initial comment"),
				resource.TestCheckResourceAttr("databricks_share.this", "object.#", "2"),
			),
		},
		resource.TestStep{
			Config: initialShareConfig,
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
			},
		},
		resource.TestStep{
			Config: updatedShareConfig,
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("databricks_share.this", "id", "fixture-share"),
				resource.TestCheckResourceAttr("databricks_share.this", "comment", "updated comment"),
				resource.TestCheckResourceAttr("databricks_share.this", "updated_by", "updater@example.com"),
				resource.TestCheckResourceAttr("databricks_share.this", "object.#", "2"),
				resource.TestCheckResourceAttr("databricks_share.this", "object.0.name", "main.default.first"),
				resource.TestCheckResourceAttr("databricks_share.this", "object.1.name", "main.default.third"),
			),
		},
		resource.TestStep{
			Config: updatedShareConfig,
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
			},
		},
	)
}
