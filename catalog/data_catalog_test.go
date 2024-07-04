package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestCatalogData(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockCatalogsAPI().EXPECT()
			e.GetByName(mock.Anything, "a").Return(&catalog.CatalogInfo{
				Name:        "a",
				Owner:       "account users",
				CatalogType: "MANAGED_CATALOG",
			}, nil)
		},
		Resource: DataSourceCatalog(),
		HCL: `
		name="a"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"name":                        "a",
		"id":                          "a",
		"catalog_info.0.name":         "a",
		"catalog_info.0.owner":        "account users",
		"catalog_info.0.catalog_type": "MANAGED_CATALOG",
	})
}

func TestCatalogData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceCatalog(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
