package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/unitycatalog"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestMetastoreAssignmentCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceMetastoreAssignment(),
		qa.CornerCaseID("1000200030004|aaaaaa-bb-cc"),
		qa.CornerCaseSkipCRUD("create"))
}

func TestMetastoreAssignment_Create(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PUT",
				Resource: "/api/2.1/unity-catalog/workspaces/123/metastore",
				ExpectedRequest: map[string]interface{}{
					"default_catalog_name": "hive_metastore",
					"metastore_id":         "a",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/current-metastore-assignment",
				Response: unitycatalog.MetastoreAssignment{
					MetastoreId:        "a",
					WorkspaceId:        "123",
					DefaultCatalogName: "hive_metastore",
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

func TestAccountMetastoreAssignment_Create(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/100/workspaces/123/metastores/a",
				ExpectedRequest: map[string]interface{}{
					"default_catalog_name": "hive_metastore",
					"metastore_id":         "a",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/100/workspaces/123/metastore?",
				Response: unitycatalog.MetastoreAssignment{
					MetastoreId:        "a",
					WorkspaceId:        "123",
					DefaultCatalogName: "hive_metastore",
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
