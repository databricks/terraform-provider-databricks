package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestMetastoreAssignmentCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceMetastoreAssignment(),
		qa.CornerCaseID("1000200030004|aaaaaa-bb-cc"),
	)
}

func TestMetastoreAssignment_Create(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PUT",
				Resource: "/api/2.1/unity-catalog/workspaces/123/metastore",
				ExpectedRequest: catalog.CreateMetastoreAssignment{
					MetastoreId: "a",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/current-metastore-assignment",
				Response: catalog.MetastoreAssignment{
					MetastoreId: "a",
					WorkspaceId: 123,
				},
			},
		},
		Resource: ResourceMetastoreAssignment(),
		Create:   true,
		HCL: `
		workspace_id = 123
		metastore_id = "a"
		`,
	}.ApplyNoError(t)
}

func TestMetastoreAssignment_Import(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/current-metastore-assignment",
				Response: catalog.MetastoreAssignment{
					MetastoreId:        "a",
					WorkspaceId:        123,
					DefaultCatalogName: "test_metastore",
				},
			},
		},
		Resource: ResourceMetastoreAssignment(),
		Read:     true,
		ID:       "123|a",
	}.ApplyAndExpectData(t, map[string]any{
		"workspace_id": 123,
		"metastore_id": "a",
	})
}

func TestMetastoreAssignmentAccount_Create(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/100/workspaces/123/metastores/a",
				ExpectedRequest: catalog.AccountsCreateMetastoreAssignment{
					MetastoreAssignment: &catalog.CreateMetastoreAssignment{
						MetastoreId: "a",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/100/workspaces/123/metastore?",
				Response: catalog.AccountsMetastoreAssignment{
					MetastoreAssignment: &catalog.MetastoreAssignment{
						MetastoreId: "a",
						WorkspaceId: 123,
					},
				},
			},
		},
		Resource:  ResourceMetastoreAssignment(),
		AccountID: "100",
		Create:    true,
		HCL: `
		workspace_id = 123
		metastore_id = "a"
		`,
	}.ApplyNoError(t)
}

func TestMetastoreAssignmentAccount_Update(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PUT",
				Resource: "/api/2.0/accounts/100/workspaces/123/metastores/b",
				//CreateMetastoreAssignment needs to have default_catalog_name marked as omitempty
				ExpectedRequest: map[string]any{
					"metastore_assignment": map[string]any{
						"metastore_id": "b",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/100/workspaces/123/metastore?",
				Response: catalog.AccountsMetastoreAssignment{
					MetastoreAssignment: &catalog.MetastoreAssignment{
						MetastoreId: "b",
						WorkspaceId: 123,
					},
				},
			},
		},
		Resource:    ResourceMetastoreAssignment(),
		AccountID:   "100",
		ID:          "123|a",
		Update:      true,
		RequiresNew: true,
		InstanceState: map[string]string{
			"workspace_id": "123",
			"metastore_id": "a",
		},
		HCL: `
		workspace_id = 123
		metastore_id = "b"
		`,
	}.ApplyNoError(t)
}

func TestMetastoreAssignmentWorskpace_Update(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PUT",
				Resource: "/api/2.0/accounts/100/workspaces/124/metastores/a",
				//CreateMetastoreAssignment needs to have default_catalog_name marked as omitempty
				ExpectedRequest: map[string]any{
					"metastore_assignment": map[string]any{
						"metastore_id": "a",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/100/workspaces/123/metastore?",
				Response: catalog.AccountsMetastoreAssignment{
					MetastoreAssignment: &catalog.MetastoreAssignment{
						MetastoreId: "a",
						WorkspaceId: 123,
					},
				},
			},
		},
		Resource:    ResourceMetastoreAssignment(),
		AccountID:   "100",
		ID:          "123|a",
		Update:      true,
		RequiresNew: true,
		InstanceState: map[string]string{
			"workspace_id": "123",
			"metastore_id": "a",
		},
		HCL: `
		workspace_id = 124
		metastore_id = "a"
		`,
	}.ApplyNoError(t)
}
