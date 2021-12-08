package catalog

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
)

func TestMetastoreAssignmentCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceMetastoreAssignment(),
		qa.CornerCaseID("1000200030004|aaaaaa-bb-cc"),
		qa.CornerCaseSkipCRUD("read"))
}

func TestMetastoreAssignment_Create(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PUT",
				Resource: "/api/2.0/unity-catalog/workspaces/123/metastore",
				ExpectedRequest: MetastoreAssignment{
					WorkspaceID:        123,
					MetastoreID:        "a",
					DefaultCatalogName: "main",
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
