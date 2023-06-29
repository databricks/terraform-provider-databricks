package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestMetastoreData(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/testaccount/metastores/abc?",
				Response: catalog.AccountsMetastoreInfo{
					MetastoreInfo: &catalog.MetastoreInfo{
						Name:                      "a",
						StorageRoot:               "",
						DefaultDataAccessConfigId: "sth",
						Owner:                     "John.Doe@example.com",
						MetastoreId:               "abc",
						Region:                    "",
						Cloud:                     "",
						GlobalMetastoreId:         "",
						CreatedAt:                 0,
						CreatedBy:                 "",
						UpdatedAt:                 0,
						UpdatedBy:                 "",
						DeltaSharingScope:         "",
						DeltaSharingRecipientTokenLifetimeInSeconds: 0,
						DeltaSharingOrganizationName:                "",
					},
				},
			},
		},
		Resource:    DataSourceMetastore(),
		Read:        true,
		NonWritable: true,
		ID:          "abc",
		AccountID:   "testaccount",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
}

func TestMetastoreData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceMetastore(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		AccountID:   "_",
	}.ExpectError(t, "I'm a teapot")
}
