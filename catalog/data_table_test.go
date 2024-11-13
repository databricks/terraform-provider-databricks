package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestTableData(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockTablesAPI().EXPECT()
			e.GetByFullName(mock.Anything, "a.b.c").Return(&catalog.TableInfo{
				FullName:  "a.b.c",
				Name:      "c",
				Owner:     "account users",
				TableType: catalog.TableTypeExternal,
			}, nil)
		},
		Resource: DataSourceTable(),
		HCL: `
		name="a.b.c"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"name":                    "a.b.c",
		"table_info.0.full_name":  "a.b.c",
		"table_info.0.name":       "c",
		"table_info.0.owner":      "account users",
		"table_info.0.table_type": "EXTERNAL",
	})
}

func TestTableData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceTable(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
