package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestViewsData(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTablesAPI().EXPECT().
				ListAll(mock.Anything, catalog.ListTablesRequest{
					CatalogName: "a",
					SchemaName:  "b",
				}).
				Return([]catalog.TableInfo{
					{
						CatalogName: "a",
						SchemaName:  "b",
						Name:        "c",
						FullName:    "a.b.c",
						TableType:   "MANAGED",
					},
					{
						CatalogName: "a",
						SchemaName:  "b",
						Name:        "d",
						FullName:    "a.b.d",
						TableType:   "VIEW",
					},
				}, nil)
		},
		Resource: DataSourceViews(),
		HCL: `
		catalog_name = "a"
		schema_name = "b"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": []string{"a.b.d"},
	})
}

func TestViewsData_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTablesAPI().EXPECT().
				ListAll(mock.Anything, catalog.ListTablesRequest{
					CatalogName: "",
					SchemaName:  "",
				}).
				Return(nil, &apierr.APIError{
					ErrorCode: "BAD_REQUEST",
					Message:   "Bad request: unable to list views",
				})
		},
		Resource:    DataSourceViews(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "Bad request: unable to list views")
}
