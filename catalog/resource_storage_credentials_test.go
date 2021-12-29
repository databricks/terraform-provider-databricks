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
				ExpectedRequest: StorageCredentialConfig{
					Name: "a",
					Aws: &AwsIamRole{
						RoleARN: "def",
					},
					Comment: "c",
				},
				Response: StorageCredentialConfig{
					Name: "a",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/storage-credentials/a",
				Response: StorageCredentialConfig{
					Name: "a",
					Aws: &AwsIamRole{
						RoleARN: "def",
					},
					MetastoreID: "d",
					Owner:       "e",
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
