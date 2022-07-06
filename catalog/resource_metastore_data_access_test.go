package catalog

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDacCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceMetastoreDataAccess(),
		qa.CornerCaseID("a|b"))
}

func TestCreateDac(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/metastores/abc/data-access-configurations",
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
				Resource: "/api/2.1/unity-catalog/metastores/abc",
				ExpectedRequest: map[string]interface{}{
					"default_data_access_config_id": "efg",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/metastores/abc/data-access-configurations/efg",
				Response: DataAccessConfiguration{
					Name: "bcd",
					Aws: &AwsIamRole{
						RoleARN: "def",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/metastores/abc",
				Response: MetastoreInfo{
					DefaultDacID: "efg",
				},
			},
		},
		Create:   true,
		Resource: ResourceMetastoreDataAccess(),
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

func TestCreateDacWithAzMI(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/metastores/abc/data-access-configurations",
				ExpectedRequest: DataAccessConfiguration{
					Name: "bcd",
					AzMI: &AzureManagedIdentity{
						AccessConnectorID: "def",
					},
				},
				Response: DataAccessConfiguration{
					ID: "efg",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/metastores/abc",
				ExpectedRequest: map[string]interface{}{
					"default_data_access_config_id": "efg",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/metastores/abc/data-access-configurations/efg",
				Response: DataAccessConfiguration{
					Name: "bcd",
					AzMI: &AzureManagedIdentity{
						AccessConnectorID: "def",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/metastores/abc",
				Response: MetastoreInfo{
					DefaultDacID: "efg",
				},
			},
		},
		Create:   true,
		Resource: ResourceMetastoreDataAccess(),
		HCL: `
		metastore_id = "abc"
		name = "bcd"
		is_default = true
		azure_managed_identity {
			access_connector_id = "def"
		}
		`,
	}.ApplyNoError(t)
}
