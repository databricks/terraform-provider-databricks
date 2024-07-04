package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestCatalogWorkspaceBindingsCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceCatalogWorkspaceBinding(),
		qa.CornerCaseID("1234567890101112|catalog|my_catalog"),
		qa.CornerCaseSkipCRUD("create"))
}

func TestCatalogWorkspaceBindings_Create(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/bindings/catalog/my_catalog",
				ExpectedRequest: catalog.UpdateWorkspaceBindingsParameters{
					Add: []catalog.WorkspaceBinding{
						{
							BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadWrite,
							WorkspaceId: int64(1234567890101112),
						},
					},
					SecurableName: "my_catalog",
					SecurableType: "catalog",
				},
				Response: catalog.WorkspaceBindingsResponse{
					Bindings: []catalog.WorkspaceBinding{
						{
							BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadWrite,
							WorkspaceId: int64(1234567890101112),
						},
					},
				},
			}, {
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/bindings/catalog/my_catalog?",
				Response: catalog.WorkspaceBindingsResponse{
					Bindings: []catalog.WorkspaceBinding{
						{
							BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadWrite,
							WorkspaceId: int64(1234567890101112),
						},
					},
				},
			},
		},
		Resource: ResourceCatalogWorkspaceBinding(),
		Create:   true,
		HCL: `
		catalog_name = "my_catalog"
		workspace_id = "1234567890101112"
		`,
	}.ApplyNoError(t)
}

func TestCatalogWorkspaceBindingsReadOnly_Create(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/bindings/catalog/my_catalog",
				ExpectedRequest: catalog.UpdateWorkspaceBindingsParameters{
					Add: []catalog.WorkspaceBinding{
						{
							BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadOnly,
							WorkspaceId: int64(1234567890101112),
						},
					},
					SecurableName: "my_catalog",
					SecurableType: "catalog",
				},
				Response: catalog.WorkspaceBindingsResponse{
					Bindings: []catalog.WorkspaceBinding{
						{
							BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadOnly,
							WorkspaceId: int64(1234567890101112),
						},
					},
				},
			}, {
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/bindings/catalog/my_catalog?",
				Response: catalog.WorkspaceBindingsResponse{
					Bindings: []catalog.WorkspaceBinding{
						{
							BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadOnly,
							WorkspaceId: int64(1234567890101112),
						},
					},
				},
			},
		},
		Resource: ResourceCatalogWorkspaceBinding(),
		Create:   true,
		HCL: `
		catalog_name = "my_catalog"
		workspace_id = "1234567890101112"
		binding_type = "BINDING_TYPE_READ_ONLY"
		`,
	}.ApplyNoError(t)
}

func TestCatalogWorkspaceBindingsReadImport(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/bindings/catalog/my_catalog?",
				Response: catalog.WorkspaceBindingsResponse{
					Bindings: []catalog.WorkspaceBinding{
						{
							BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadOnly,
							WorkspaceId: int64(1234567890101112),
						},
					},
				},
			},
		},
		Resource: ResourceCatalogWorkspaceBinding(),
		ID:       "1234567890101112|catalog|my_catalog",
		New:      true,
		Read:     true,
	}.ApplyAndExpectData(t, map[string]any{
		"workspace_id":   1234567890101112,
		"securable_type": "catalog",
		"securable_name": "my_catalog",
	})
}

func TestCatalogWorkspaceBindingsReadErrors(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceCatalogWorkspaceBinding(),
		ID:       "1234567890101112|catalog",
		New:      true,
		Read:     true,
	}.ExpectError(t, "incorrect binding id: 1234567890101112|catalog. Correct format: <workspace_id>|<securable_type>|<securable_name>")

	qa.ResourceFixture{
		Resource: ResourceCatalogWorkspaceBinding(),
		ID:       "A234567890101112|catalog|my_catalog",
		New:      true,
		Read:     true,
	}.ExpectError(t, "can't parse workspace_id: strconv.ParseInt: parsing \"A234567890101112\": invalid syntax")
}
