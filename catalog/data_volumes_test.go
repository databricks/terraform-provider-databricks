package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestDataSourceVolumes(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockVolumesAPI().EXPECT().
				ListAll(mock.Anything, catalog.ListVolumesRequest{
					CatalogName: "a",
					SchemaName:  "b",
				}).
				Return([]catalog.VolumeInfo{
					{
						FullName: "a.b.c",
						Name:     "a",
					},
				}, nil)
		},
		Resource: DataSourceVolumes(),
		HCL: `
		catalog_name = "a"
		schema_name = "b"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": []string{"a.b.c"},
	})
}

func TestDataSourceVolumes_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockVolumesAPI().EXPECT().
				ListAll(mock.Anything, catalog.ListVolumesRequest{
					CatalogName: "",
					SchemaName:  "",
				}).
				Return(nil, &apierr.APIError{
					ErrorCode: "BAD_REQUEST",
					Message:   "Bad request: unable to list volumes",
				})
		},
		Resource:    DataSourceVolumes(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "Bad request: unable to list volumes")
}
