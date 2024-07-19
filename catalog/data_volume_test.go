package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestDataSourceVolume(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockVolumesAPI().EXPECT()
			e.ReadByName(mock.Anything, "a.b.c").Return(&catalog.VolumeInfo{
				FullName:    "a.b.c",
				CatalogName: "a",
				SchemaName:  "b",
				Name:        "c",
				Owner:       "account users",
				VolumeType:  catalog.VolumeTypeManaged,
			}, nil)
		},
		Resource: DataSourceVolume(),
		HCL: `
		name="a.b.c"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"name":                       "a.b.c",
		"volume_info.0.full_name":    "a.b.c",
		"volume_info.0.catalog_name": "a",
		"volume_info.0.schema_name":  "b",
		"volume_info.0.name":         "c",
		"volume_info.0.owner":        "account users",
		"volume_info.0.volume_type":  "MANAGED",
	})
}

func TestDataSourceVolume_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceVolume(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
