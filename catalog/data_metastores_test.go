package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestMetastoresDataContainsName(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockAccountMetastoresAPI().EXPECT().
				ListAll(mock.Anything).
				Return([]catalog.MetastoreInfo{
					{
						Name:                      "a",
						StorageRoot:               "abc",
						DefaultDataAccessConfigId: "sth",
						Owner:                     "John.Doe@example.com",
						MetastoreId:               "abc",
					},
					{
						Name:                      "b",
						StorageRoot:               "dcw",
						DefaultDataAccessConfigId: "sth",
						Owner:                     "John.Doe@example.com",
						MetastoreId:               "ded",
					},
				}, nil)
		},
		Resource:    DataSourceMetastores(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		AccountID:   "testaccount",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": map[string]interface{}{"a": "abc", "b": "ded"},
	})
}

func TestMetastoresData_Error(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockAccountMetastoresAPI().EXPECT().
				ListAll(mock.Anything).
				Return(nil, &apierr.APIError{
					ErrorCode: "BAD_REQUEST",
					Message:   "Bad request: unable to list metastores",
				})
		},
		Resource:    DataSourceMetastores(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		AccountID:   "_",
	}.ExpectError(t, "Bad request: unable to list metastores")
}
