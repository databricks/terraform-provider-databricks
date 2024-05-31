package sharing

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestShareData(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/shares/a?include_shared_data=true",
				Response: ShareInfo{
					Name: "a",
					Objects: []SharedDataObject{
						{
							Name:                     "a",
							DataObjectType:           "TABLE",
							Comment:                  "c",
							CDFEnabled:               false,
							StartVersion:             0,
							SharedAs:                 "",
							AddedAt:                  0,
							AddedBy:                  "",
							HistoryDataSharingStatus: "DISABLED",
							Status:                   "ACTIVE",
							Partitions:               []Partition{},
						},
					},
					CreatedBy: "bob",
					CreatedAt: 1921321,
				},
			},
		},
		Resource:    DataSourceShare(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL: `
		name = "a"
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "bob", d.Get("created_by"))
	assert.Equal(t, 1921321, d.Get("created_at"))

	expectedShare := d.Get("object").(*schema.Set).List()[0].(map[string]interface{})

	assert.Equal(t, 0, expectedShare["added_at"])
	assert.Equal(t, "", expectedShare["added_by"])
	assert.Equal(t, "c", expectedShare["comment"])
	assert.Equal(t, "TABLE", expectedShare["data_object_type"])
	assert.Equal(t, "a", expectedShare["name"])
	assert.Equal(t, "", expectedShare["shared_as"])
	assert.Equal(t, 0, expectedShare["start_version"])
	assert.Equal(t, false, expectedShare["cdf_enabled"])
	assert.Equal(t, "ACTIVE", expectedShare["status"])
	assert.Equal(t, "DISABLED", expectedShare["history_data_sharing_status"])
	assert.Nil(t, expectedShare["partitions"])
}

func TestShareData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceShare(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
