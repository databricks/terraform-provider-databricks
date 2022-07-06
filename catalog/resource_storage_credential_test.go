package catalog

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestStorageCredentialsCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceStorageCredential())
}

func TestCreateStorageCredentials(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/storage-credentials",
				ExpectedRequest: StorageCredentialInfo{
					Name: "a",
					Aws: &AwsIamRole{
						RoleARN: "def",
					},
					Comment: "c",
				},
				Response: StorageCredentialInfo{
					Name: "a",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				Response: StorageCredentialInfo{
					Name: "a",
					Aws: &AwsIamRole{
						RoleARN: "def",
					},
					MetastoreID: "d",
				},
			},
		},
		Resource: ResourceStorageCredential(),
		Create:   true,
		HCL: `
		name = "a"
		aws_iam_role {
			role_arn = "def"
		}
		comment = "c"
		`,
	}.ApplyNoError(t)
}

func TestCreateStorageCredentialWithOwner(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/storage-credentials",
				ExpectedRequest: StorageCredentialInfo{
					Name: "a",
					Aws: &AwsIamRole{
						RoleARN: "def",
					},
					Comment: "c",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				ExpectedRequest: map[string]interface{}{
					"owner": "administrators",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				Response: StorageCredentialInfo{
					Name: "a",
					Aws: &AwsIamRole{
						RoleARN: "def",
					},
					MetastoreID: "d",
				},
			},
		},
		Resource: ResourceStorageCredential(),
		Create:   true,
		HCL: `
		name = "a"
		aws_iam_role {
			role_arn = "def"
		}
		comment = "c"
		owner = "administrators"
		`,
	}.ApplyNoError(t)
}

func TestUpdateStorageCredentials(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				ExpectedRequest: map[string]interface{}{
					"aws_iam_role": []interface{}{map[string]interface{}{
						"role_arn": "CHANGED",
					}},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				Response: StorageCredentialInfo{
					Name: "a",
					Aws: &AwsIamRole{
						RoleARN: "CHANGED",
					},
					MetastoreID: "d",
				},
			},
		},
		Resource: ResourceStorageCredential(),
		Update:   true,
		ID:       "a",
		InstanceState: map[string]string{
			"name":    "a",
			"comment": "c",
		},
		HCL: `
		name = "a"
		aws_iam_role {
			role_arn = "CHANGED"
		}
		comment = "c"
		`,
	}.ApplyNoError(t)
}

func TestCreateStorageCredentialWithAzMI(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/storage-credentials",
				ExpectedRequest: StorageCredentialInfo{
					Name: "a",
					AzMI: &AzureManagedIdentity{
						AccessConnectorID: "def",
					},
					Comment: "c",
				},
				Response: StorageCredentialInfo{
					Name: "a",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				Response: StorageCredentialInfo{
					Name: "a",
					AzMI: &AzureManagedIdentity{
						AccessConnectorID: "def",
					},
					MetastoreID: "d",
				},
			},
		},
		Resource: ResourceStorageCredential(),
		Create:   true,
		HCL: `
		name = "a"
		azure_managed_identity {
			access_connector_id = "def"
		}
		comment = "c"
		`,
	}.ApplyNoError(t)
}
