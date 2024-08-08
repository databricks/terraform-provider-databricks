package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestDataSourceSchema(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockSchemasAPI().EXPECT()
			e.GetByFullName(mock.Anything, "a.b").Return(&catalog.SchemaInfo{
				FullName:    "a.b",
				CatalogName: "a",
				Name:        "b",
				Owner:       "account users",
			}, nil)
		},
		Resource: DataSourceSchema(),
		HCL: `
		name="a.b"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"name":                       "a.b",
		"id":                         "a.b",
		"schema_info.0.full_name":    "a.b",
		"schema_info.0.catalog_name": "a",
		"schema_info.0.name":         "b",
		"schema_info.0.owner":        "account users",
	})
}

func TestDataSourceSchema_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceSchema(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
