package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestCatalogWorkspaceBindingsCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceCatalogWorkspaceBinding())
}

func TestCatalogWorkspaceBindingsAssign(t *testing.T) {
	resource := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/workspace-bindings/catalogs/my_catalog",
				Response: catalog.CurrentWorkspaceBindings{
					Workspaces: []int64{1234567890101112},
				},
			},
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
			},
		},
		Resource: ResourceCatalogWorkspaceBinding(),
		Create:   true,
		HCL: `
		name = "my_catalog"
		workspace = 1234567890101112
		`,
	}

	// panic: interface conversion: interface {} is nil, not int64
	resource.ApplyNoError(t)
	// resource.ApplyAndExpectData(t, map[string]any{"workspaces": []int64{1234567890}})
}
