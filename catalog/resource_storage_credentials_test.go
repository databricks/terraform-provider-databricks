package catalog

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
)

func TestStorageCredentialsCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceStorageCredential())
}

func TestCreateStorageCredentials(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/unity-catalog/storage-credentials",
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
				Resource: "/api/2.0/unity-catalog/storage-credentials/a",
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

func TestUpdateStorageCredentials(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/unity-catalog/storage-credentials/a",
				ExpectedRequest: StorageCredentialInfo{
					Name: "a",
					Aws: &AwsIamRole{
						RoleARN: "CHANGED",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/storage-credentials/a",
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
