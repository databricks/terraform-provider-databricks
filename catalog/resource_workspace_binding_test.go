package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/catalog/bindings"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestWorkspaceBindingsCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceWorkspaceBinding(),
		qa.CornerCaseID("1234567890101112|catalog|my_catalog"),
		qa.CornerCaseSkipCRUD("create"))
}

func TestWorkspaceBindings_Create(t *testing.T) {
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
				Response: catalog.UpdateWorkspaceBindingsResponse{
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
				Response: catalog.GetWorkspaceBindingsResponse{
					Bindings: []catalog.WorkspaceBinding{
						{
							BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadWrite,
							WorkspaceId: int64(1234567890101112),
						},
					},
				},
			},
		},
		Resource: ResourceWorkspaceBinding(),
		Create:   true,
		HCL: `
		catalog_name = "my_catalog"
		workspace_id = "1234567890101112"
		`,
	}.ApplyNoError(t)
}

func TestWorkspaceBindingsReadOnly_Create(t *testing.T) {
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
				Response: catalog.UpdateWorkspaceBindingsResponse{
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
				Response: catalog.GetWorkspaceBindingsResponse{
					Bindings: []catalog.WorkspaceBinding{
						{
							BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadOnly,
							WorkspaceId: int64(1234567890101112),
						},
					},
				},
			},
		},
		Resource: ResourceWorkspaceBinding(),
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
				Response: catalog.UpdateWorkspaceBindingsResponse{
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
				Response: catalog.GetWorkspaceBindingsResponse{
					Bindings: []catalog.WorkspaceBinding{
						{
							BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadOnly,
							WorkspaceId: int64(1234567890101112),
						},
					},
				},
			},
		},
		Resource: ResourceWorkspaceBinding(),
		Create:   true,
		HCL: `
		securable_name = "my_catalog"
		securable_type = "catalog"
		workspace_id   = "1234567890101112"
		binding_type   = "BINDING_TYPE_READ_ONLY"
		`,
	}.ApplyNoError(t)
}

func TestSecurableWorkspaceBindings_CreateExtLocation(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockWorkspaceBindingsAPI().EXPECT()
			e.UpdateBindings(mock.Anything, catalog.UpdateWorkspaceBindingsParameters{
				Add: []catalog.WorkspaceBinding{{
					BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadWrite,
					WorkspaceId: int64(1234567890101112),
				},
				},
				SecurableName: "external_location",
				SecurableType: string(bindings.BindingsSecurableTypeExternalLocation),
			}).Return(&catalog.UpdateWorkspaceBindingsResponse{
				Bindings: []catalog.WorkspaceBinding{
					{
						BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadWrite,
						WorkspaceId: int64(1234567890101112),
					},
				},
			}, nil)
			e.GetBindingsBySecurableTypeAndSecurableName(mock.Anything, string(bindings.BindingsSecurableTypeExternalLocation), "external_location").Return(&catalog.GetWorkspaceBindingsResponse{
				Bindings: []catalog.WorkspaceBinding{
					{
						WorkspaceId: int64(1234567890101112),
					},
				},
			}, nil)
		},
		Resource: ResourceWorkspaceBinding(),
		Create:   true,
		HCL: `
		securable_name = "external_location"
		securable_type = "external_location"
		workspace_id   = "1234567890101112"
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
				Response: catalog.UpdateWorkspaceBindingsResponse{
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
				Response: catalog.GetWorkspaceBindingsResponse{
					Bindings: []catalog.WorkspaceBinding{
						{
							BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadOnly,
							WorkspaceId: int64(1234567890101112),
						},
					},
				},
			},
		},
		Resource: ResourceWorkspaceBinding(),
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

func TestWorkspaceBindingsRead_OldStyleSecurableType(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/bindings/external_location/my_catalog?",
				Response: catalog.GetWorkspaceBindingsResponse{
					Bindings: []catalog.WorkspaceBinding{
						{
							BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadOnly,
							WorkspaceId: int64(1234567890101112),
						},
					},
				},
			},
		},
		Resource: ResourceWorkspaceBinding(),
		ID:       "1234567890101112|external-location|my_catalog",
		Read:     true,
	}.ApplyAndExpectData(t, map[string]any{
		"workspace_id":   1234567890101112,
		"securable_type": string(bindings.BindingsSecurableTypeExternalLocation),
		"securable_name": "my_catalog",
	})
}

func TestWorkspaceBindingsReadImport(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/bindings/catalog/my_catalog?",
				Response: catalog.GetWorkspaceBindingsResponse{
					Bindings: []catalog.WorkspaceBinding{
						{
							BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadOnly,
							WorkspaceId: int64(1234567890101112),
						},
					},
				},
			},
		},
		Resource: ResourceWorkspaceBinding(),
		ID:       "1234567890101112|catalog|my_catalog",
		New:      true,
		Read:     true,
	}.ApplyAndExpectData(t, map[string]any{
		"workspace_id":   1234567890101112,
		"securable_type": "catalog",
		"securable_name": "my_catalog",
	})
}

func TestWorkspaceBindingsReadErrors(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceWorkspaceBinding(),
		ID:       "1234567890101112|catalog",
		New:      true,
		Read:     true,
	}.ExpectError(t, "incorrect binding id: 1234567890101112|catalog. Correct format: <workspace_id>|<securable_type>|<securable_name>")

	qa.ResourceFixture{
		Resource: ResourceWorkspaceBinding(),
		ID:       "A234567890101112|catalog|my_catalog",
		New:      true,
		Read:     true,
	}.ExpectError(t, "can't parse workspace_id: strconv.ParseInt: parsing \"A234567890101112\": invalid syntax")
}
