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

func TestSecurableWorkspaceBindings_Create(t *testing.T) {
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
		securable_name = "my_catalog"
		securable_type = "catalog"
		workspace_id   = "1234567890101112"
		binding_type   = "BINDING_TYPE_READ_ONLY"
		`,
	}.ApplyNoError(t)
}

func TestSecurableWorkspaceBindings_Delete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/bindings/catalog/my_catalog",
				ExpectedRequest: catalog.UpdateWorkspaceBindingsParameters{
					Remove: []catalog.WorkspaceBinding{
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
		Delete:   true,
		ID:       "1234567890101112|catalog|my_catalog",
		HCL: `
		securable_name = "my_catalog"
		securable_type = "catalog"
		workspace_id   = "1234567890101112"
		binding_type   = "BINDING_TYPE_READ_ONLY"
		`,
	}.ApplyNoError(t)
}
