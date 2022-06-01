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
					Owner:   "administrators",
					Comment: "c",
				},
				Response: StorageCredentialInfo{
					Name:    "a",
					Owner:   "not_admin",
					Comment: "c",
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
			{
				Method:   "PATCH",
				Resource: "/api/2.0/unity-catalog/storage-credentials/a",
				ExpectedRequest: map[string]interface{}{
					"owner": "administrators",
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
				Resource: "/api/2.0/unity-catalog/storage-credentials/a",
				ExpectedRequest: map[string]interface{}{
					"aws_iam_role": []interface{}{map[string]interface{}{
						"rolearn": "CHANGED",
					}},
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

func TestUpdateStorageCredentialsName(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/unity-catalog/storage-credentials/a",
				ExpectedRequest: map[string]interface{}{
					"name": "b",
					"aws":  "CHANGED",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/storage-credentials/b",
				Response: StorageCredentialInfo{
					Name: "b",
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
		name = "b"
		aws_iam_role {
			role_arn = "CHANGED"
		}
		comment = "c"
		`,
	}.ApplyNoError(t)
}
