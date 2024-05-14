package catalog

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestMetastoreDataById(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(m *mocks.MockAccountClient) {
			e := m.GetMockAccountMetastoresAPI().EXPECT()
			e.GetByMetastoreId(mock.Anything, "abc").Return(&catalog.AccountsMetastoreInfo{
				MetastoreInfo: &catalog.MetastoreInfo{
					Name:        "xyz",
					MetastoreId: "abc",
					Owner:       "pqr",
					Region:      "unknown",
				},
			}, nil)
		},
		Resource:    DataSourceMetastore(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		AccountID:   "testaccount",
		HCL: `
		metastore_id = "abc"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"name":                          "xyz",
		"region":                        "unknown",
		"metastore_info.0.name":         "xyz",
		"metastore_info.0.owner":        "pqr",
		"metastore_info.0.metastore_id": "abc",
	})
}

func TestMetastoreDataErrorNoParams(t *testing.T) {
	qa.ResourceFixture{
		Resource:    DataSourceMetastore(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		AccountID:   "_",
	}.ExpectError(t, "one of metastore_id, name or region must be provided")
}

func TestMetastoreDataErrorMultipleParams(t *testing.T) {
	qa.ResourceFixture{
		Resource:    DataSourceMetastore(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		AccountID:   "_",
		HCL: `
		metastore_id = "abc"
		name         = "abc"
		`,
	}.ExpectError(t, "only one of metastore_id, name or region must be provided")
}

func TestMetastoreDataError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceMetastore(),
		Read:        true,
		NonWritable: true,
		ID:          "id",
		AccountID:   "_",
		HCL: `
		metastore_id = "abc"
		`,
	}.ExpectError(t, "i'm a teapot")
}

func TestMetastoreByName(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(m *mocks.MockAccountClient) {
			e := m.GetMockAccountMetastoresAPI().EXPECT()
			e.ListAll(mock.Anything).Return([]catalog.MetastoreInfo{
				{
					Name:        "xyz",
					MetastoreId: "abc",
					Owner:       "pqr",
					Region:      "unknown",
				},
			}, nil)
		},
		Resource:    DataSourceMetastore(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		AccountID:   "testaccount",
		HCL: `
		name = "xyz"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"name":                          "xyz",
		"region":                        "unknown",
		"metastore_info.0.name":         "xyz",
		"metastore_info.0.owner":        "pqr",
		"metastore_info.0.metastore_id": "abc",
	})
}

func TestMetastoreByRegion(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(m *mocks.MockAccountClient) {
			e := m.GetMockAccountMetastoresAPI().EXPECT()
			e.ListAll(mock.Anything).Return([]catalog.MetastoreInfo{
				{
					Name:        "xyz",
					MetastoreId: "abc",
					Owner:       "pqr",
					Region:      "westus",
				},
			}, nil)
		},
		Resource:    DataSourceMetastore(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		AccountID:   "testaccount",
		HCL: `
		region = "westus"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"id":                            "abc",
		"name":                          "xyz",
		"region":                        "westus",
		"metastore_info.0.name":         "xyz",
		"metastore_info.0.owner":        "pqr",
		"metastore_info.0.metastore_id": "abc",
	})
}

func TestMetastoreByNameNoData(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(m *mocks.MockAccountClient) {
			e := m.GetMockAccountMetastoresAPI().EXPECT()
			e.ListAll(mock.Anything).Return([]catalog.MetastoreInfo{}, nil)
		},
		Resource:    DataSourceMetastore(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		AccountID:   "testaccount",
		HCL: `
		name = "test"
		`,
	}.ExpectError(t, "a metastore with name 'test' or in region '' is not found")
}

func TestMetastoreByNameListError(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(m *mocks.MockAccountClient) {
			e := m.GetMockAccountMetastoresAPI().EXPECT()
			e.ListAll(mock.Anything).Return([]catalog.MetastoreInfo{}, fmt.Errorf("i'm a teapot"))
		},
		Resource:    DataSourceMetastore(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		AccountID:   "testaccount",
		HCL: `
		name = "test"
		`,
	}.ExpectError(t, "i'm a teapot")
}

func TestMetastoreByRegionMultipleEntries(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(m *mocks.MockAccountClient) {
			e := m.GetMockAccountMetastoresAPI().EXPECT()
			e.ListAll(mock.Anything).Return([]catalog.MetastoreInfo{
				{
					Name:        "xyz",
					MetastoreId: "abc",
					Owner:       "pqr",
					Region:      "westus",
				},
				{
					Name:        "xyz2",
					MetastoreId: "abc2",
					Owner:       "pqr",
					Region:      "westus",
				},
			}, nil)
		},
		Resource:    DataSourceMetastore(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		AccountID:   "testaccount",
		HCL: `
		region = "westus"
		`,
	}.ExpectError(t, "there are 2 metastores with name '' in region 'westus'")
}
