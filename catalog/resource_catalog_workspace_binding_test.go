package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestCatalogWorkspaceBindingsCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceCatalogWorkspaceBinding(),
		qa.CornerCaseID("my_catalog|1234567890101112"),
		qa.CornerCaseSkipCRUD("create"))
}

func TestCatalogWorkspaceBindings_Create(t *testing.T) {
	resource := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/workspace-bindings/catalogs/my_catalog",
				ExpectedRequest: catalog.UpdateWorkspaceBindings{
					Name:             "my_catalog",
					AssignWorkspaces: []int64{1234567890101112},
				},
				Response: catalog.CurrentWorkspaceBindings{
					Workspaces: []int64{1234567890101112},
				},
			}, {
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/workspace-bindings/catalogs/my_catalog?",
				Response: catalog.CurrentWorkspaceBindings{
					Workspaces: []int64{1234567890101112},
				},
			},
		},
		Resource: ResourceCatalogWorkspaceBinding(),
		Create:   true,
		HCL: `
		catalog_name = "my_catalog"
		workspace_id = "1234567890101112"
		`,
	}
	resource.ApplyNoError(t)
}
