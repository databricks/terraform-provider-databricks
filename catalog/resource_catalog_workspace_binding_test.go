package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/catalog/bindings"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestCatalogWorkspaceBindingsCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceCatalogWorkspaceBinding(),
		qa.CornerCaseID("1234567890101112|catalog|my_catalog"),
		qa.CornerCaseSkipCRUD("create"))
}

func TestCatalogWorkspaceBindings_Create(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockWorkspaceBindingsAPI().EXPECT()
			e.UpdateBindings(mock.Anything, catalog.UpdateWorkspaceBindingsParameters{
				Add: []catalog.WorkspaceBinding{
					{
						BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadWrite,
						WorkspaceId: int64(1234567890101112),
					},
				},
				SecurableName: "my_catalog",
				SecurableType: string(bindings.BindingsSecurableTypeCatalog),
			}).Return(&catalog.UpdateWorkspaceBindingsResponse{
				Bindings: []catalog.WorkspaceBinding{
					{
						BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadWrite,
						WorkspaceId: int64(1234567890101112),
					},
				},
			}, nil)
			e.GetBindingsBySecurableTypeAndSecurableName(mock.Anything, string(bindings.BindingsSecurableTypeCatalog), "my_catalog").
				Return(&catalog.GetWorkspaceBindingsResponse{
					Bindings: []catalog.WorkspaceBinding{
						{
							BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadWrite,
							WorkspaceId: int64(1234567890101112),
						},
					},
				}, nil)
		},
		Resource: ResourceCatalogWorkspaceBinding(),
		Create:   true,
		HCL: `
		catalog_name = "my_catalog"
		workspace_id = "1234567890101112"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"workspace_id":   1234567890101112,
		"securable_type": "catalog",
		"securable_name": "my_catalog",
		"binding_type":   "BINDING_TYPE_READ_WRITE",
	})
}

func TestCatalogWorkspaceBindingsReadOnly_Create(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockWorkspaceBindingsAPI().EXPECT()
			e.UpdateBindings(mock.Anything, catalog.UpdateWorkspaceBindingsParameters{
				Add: []catalog.WorkspaceBinding{
					{
						BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadOnly,
						WorkspaceId: int64(1234567890101112),
					},
				},
				SecurableName: "my_catalog",
				SecurableType: string(bindings.BindingsSecurableTypeCatalog),
			}).Return(&catalog.UpdateWorkspaceBindingsResponse{
				Bindings: []catalog.WorkspaceBinding{
					{
						BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadOnly,
						WorkspaceId: int64(1234567890101112),
					},
				},
			}, nil)
			e.GetBindingsBySecurableTypeAndSecurableName(mock.Anything, "catalog", "my_catalog").
				Return(&catalog.GetWorkspaceBindingsResponse{
					Bindings: []catalog.WorkspaceBinding{
						{
							BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadOnly,
							WorkspaceId: int64(1234567890101112),
						},
					},
				}, nil)
		},
		Resource: ResourceCatalogWorkspaceBinding(),
		Create:   true,
		HCL: `
		catalog_name = "my_catalog"
		workspace_id = "1234567890101112"
		binding_type = "BINDING_TYPE_READ_ONLY"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"workspace_id":   1234567890101112,
		"securable_type": "catalog",
		"securable_name": "my_catalog",
		"binding_type":   "BINDING_TYPE_READ_ONLY",
	})
}

func TestCatalogWorkspaceBindingsReadImport(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockWorkspaceBindingsAPI().EXPECT().
				GetBindingsBySecurableTypeAndSecurableName(mock.Anything, string(bindings.BindingsSecurableTypeCatalog), "my_catalog").
				Return(&catalog.GetWorkspaceBindingsResponse{
					Bindings: []catalog.WorkspaceBinding{
						{
							BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadOnly,
							WorkspaceId: int64(1234567890101112),
						},
					},
				}, nil)
		},
		Resource: ResourceCatalogWorkspaceBinding(),
		ID:       "1234567890101112|catalog|my_catalog",
		New:      true,
		Read:     true,
	}.ApplyAndExpectData(t, map[string]any{
		"workspace_id":   1234567890101112,
		"securable_type": "catalog",
		"securable_name": "my_catalog",
		"binding_type":   "BINDING_TYPE_READ_ONLY",
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
