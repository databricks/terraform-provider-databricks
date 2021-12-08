package catalog

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
)

func TestDacCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceDataAccessConfiguration(),
		qa.CornerCaseID("a|b"))
}

func TestCreateDac(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/unity-catalog/metastores/abc/data-access-configurations",
				ExpectedRequest: DataAccessConfiguration{
					Name: "bcd",
					Aws: &AwsIamRole{
						RoleARN: "def",
					},
				},
				Response: DataAccessConfiguration{
					ID: "efg",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.0/unity-catalog/metastores/abc",
				ExpectedRequest: map[string]interface {}{
					"default_data_access_config_id":"efg",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/metastores/abc/data-access-configurations/efg",
				Response: DataAccessConfiguration{
					Name: "bcd",
					Aws: &AwsIamRole{
						RoleARN: "def",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/metastores/abc",
				Response: MetastoreInfo {
					DefaultDacID: "efg",
				},
			},
		},
		Create:   true,
		Resource: ResourceDataAccessConfiguration(),
		HCL: `
		metastore_id = "abc"
		name = "bcd"
		is_default = true
		aws_iam_role {
			role_arn = "def"
		}
		`,
	}.ApplyNoError(t)
}
