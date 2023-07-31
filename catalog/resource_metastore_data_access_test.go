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
				ExpectedRequest: map[string]any{
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
				ExpectedRequest: map[string]any{
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

func TestCreateDacWithGcpSA(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/metastores/abc/data-access-configurations",
				ExpectedRequest: DataAccessConfiguration{
					Name: "bcd",
					GcpSAKey: &GcpServiceAccountKey{
						Email:        "a@example.com",
						PrivateKeyId: "b",
						PrivateKey:   "abcdefg",
					},
				},
				Response: DataAccessConfiguration{
					ID: "efg",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/metastores/abc",
				ExpectedRequest: map[string]any{
					"default_data_access_config_id": "efg",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/metastores/abc/data-access-configurations/efg",
				Response: DataAccessConfiguration{
					Name: "bcd",
					GcpSAKey: &GcpServiceAccountKey{
						Email:        "a@example.com",
						PrivateKeyId: "b",
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
		gcp_service_account_key {
			email = "a@example.com"
			private_key_id = "b"
			private_key = "abcdefg"
		}
		`,
	}.ApplyNoError(t)
}

func TestCreateDacWithDbGcpSA(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/metastores/abc/data-access-configurations",
				ExpectedRequest: DataAccessConfiguration{
					Name:    "bcd",
					DBGcpSA: &DbGcpServiceAccount{},
				},
				Response: DataAccessConfiguration{
					ID: "efg",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/metastores/abc",
				ExpectedRequest: map[string]any{
					"default_data_access_config_id": "efg",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/metastores/abc/data-access-configurations/efg",
				Response: DataAccessConfiguration{
					Name: "bcd",
					DBGcpSA: &DbGcpServiceAccount{
						Email: "a@example.com",
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
		databricks_gcp_service_account {}
		`,
	}.ApplyNoError(t)
}
