package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestExternalLocationDataVerify(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockExternalLocationsAPI().EXPECT()
			e.GetByName(mock.Anything, "abc").Return(
				&catalog.ExternalLocationInfo{
					Name:           "abc",
					Owner:          "admin",
					CredentialName: "test",
					Url:            "s3://test",
					ReadOnly:       true,
				},
				nil)
		},
		Resource:    DataSourceExternalLocation(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL: `
		name = "abc"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"external_location_info.0.owner":           "admin",
		"external_location_info.0.url":             "s3://test",
		"external_location_info.0.credential_name": "test",
		"external_location_info.0.read_only":       true,
	})
}

func TestExternalLocationDataError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceExternalLocation(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
