package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestExternalLocationsData(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockExternalLocationsAPI().EXPECT()
			e.ListAll(mock.Anything, catalog.ListExternalLocationsRequest{}).Return(
				[]catalog.ExternalLocationInfo{
					{Name: "a"}, {Name: "b"},
				},
				nil)
		},
		Resource:    DataSourceExternalLocations(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"names": []interface{}{"a", "b"},
	})
}

func TestExternalLocationsData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceExternalLocations(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
