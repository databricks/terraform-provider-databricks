package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestTablesData(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTablesAPI().EXPECT().
				ListAll(mock.Anything, catalog.ListTablesRequest{
					CatalogName: "a",
					SchemaName:  "b",
				}).
				Return([]catalog.TableInfo{
					{
						FullName:  "a.b.c",
						Name:      "c",
						TableType: "TABLE",
					},
					{
						FullName:  "a.b.d",
						Name:      "d",
						TableType: "TABLE",
					},
				}, nil)
		},
		Resource: DataSourceTables(),
		HCL: `
		catalog_name = "a"
		schema_name = "b"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": []string{"a.b.c", "a.b.d"},
	})
}

// https://github.com/databricks/terraform-provider-databricks/issues/1264
func TestTablesDataIssue1264(t *testing.T) {
	r := DataSourceTables()
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTablesAPI().EXPECT().
				ListAll(mock.Anything, catalog.ListTablesRequest{
					CatalogName: "a",
					SchemaName:  "b",
				}).
				Return([]catalog.TableInfo{
					{
						Name:      "a",
						FullName:  "a.b.a",
						TableType: "TABLE",
					},
					{
						Name:      "b",
						FullName:  "a.b.b",
						TableType: "TABLE",
					},
				}, nil)
		},
		Resource: r,
		HCL: `
		catalog_name = "a"
		schema_name = "b"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	require.NoError(t, err)
	s := d.Get("ids").(*schema.Set)
	assert.Equal(t, 2, s.Len())
	assert.True(t, s.Contains("a.b.a"))

	d, err = qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTablesAPI().EXPECT().
				ListAll(mock.Anything, catalog.ListTablesRequest{
					CatalogName: "a",
					SchemaName:  "b",
				}).
				Return([]catalog.TableInfo{
					{
						Name:      "c",
						FullName:  "a.b.c",
						TableType: "TABLE",
					},
					{
						Name:      "d",
						FullName:  "a.b.d",
						TableType: "TABLE",
					},
				}, nil)
		},
		Resource: r,
		HCL: `
		catalog_name = "a"
		schema_name = "b"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	require.NoError(t, err)
	s = d.Get("ids").(*schema.Set)
	assert.Equal(t, 2, s.Len())
	assert.True(t, s.Contains("a.b.c"))
}

func TestTablesData_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockTablesAPI().EXPECT().
				ListAll(mock.Anything, catalog.ListTablesRequest{
					CatalogName: "",
					SchemaName:  "",
				}).
				Return(nil, &apierr.APIError{
					ErrorCode: "BAD_REQUEST",
					Message:   "Bad request: unable to list tables",
				})
		},
		Resource:    DataSourceTables(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "Bad request: unable to list tables")
}
