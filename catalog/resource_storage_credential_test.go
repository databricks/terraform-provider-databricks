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
				ExpectedRequest: map[string]any{
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
				ExpectedRequest: map[string]any{
					"aws_iam_role": map[string]any{
						"role_arn": "CHANGED",
					},
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

func TestCreateStorageCredentialWithGcpSA(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/storage-credentials",
				ExpectedRequest: StorageCredentialInfo{
					Name: "a",
					GcpSAKey: &GcpServiceAccountKey{
						Email:        "a@example.com",
						PrivateKeyId: "b",
						PrivateKey:   "abcdefg",
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
					GcpSAKey: &GcpServiceAccountKey{
						Email:        "a@example.com",
						PrivateKeyId: "b",
					},
					MetastoreID: "d",
				},
			},
		},
		Resource: ResourceStorageCredential(),
		Create:   true,
		HCL: `
		name = "a"
		gcp_service_account_key {
			email = "a@example.com"
			private_key_id = "b"
			private_key = "abcdefg"
		}
		comment = "c"
		`,
	}.ApplyNoError(t)
}

func TestUpdateAzStorageCredentials(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				ExpectedRequest: map[string]any{
					"azure_service_principal": map[string]any{
						"directory_id":   "CHANGED",
						"application_id": "CHANGED",
						"client_secret":  "CHANGED",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				Response: StorageCredentialInfo{
					Name: "a",
					Azure: &AzureServicePrincipal{
						DirectoryID:   "CHANGED",
						ApplicationID: "CHANGED",
						ClientSecret:  "CHANGED",
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
		azure_service_principal {
			directory_id   = "CHANGED"
			application_id = "CHANGED"
			client_secret  = "CHANGED"
		}
		comment = "c"
		`,
	}.ApplyNoError(t)
}

func TestUpdateAzStorageCredentialMI(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				ExpectedRequest: map[string]any{
					"azure_managed_identity": map[string]any{
						"access_connector_id": "CHANGED",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				Response: StorageCredentialInfo{
					Name: "a",
					AzMI: &AzureManagedIdentity{
						AccessConnectorID: "CHANGED",
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
		azure_managed_identity {
			access_connector_id = "CHANGED"
		}
		comment = "c"
		`,
	}.ApplyNoError(t)
}

func TestUpdateGcpSAStorageCredential(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				ExpectedRequest: map[string]any{
					"gcp_service_account_key": map[string]any{
						"email":          "a@example.com",
						"private_key_id": "b",
						"private_key":    "abcdefg",
					},
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
					GcpSAKey: &GcpServiceAccountKey{
						Email:        "a@example.com",
						PrivateKeyId: "b",
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
		gcp_service_account_key {
			email = "a@example.com"
			private_key_id = "b"
			private_key = "abcdefg"
		}
		comment = "c"
		`,
	}.ApplyNoError(t)
}

func TestUpdateGcpSAStorageCredentialNoChange(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				Response: StorageCredentialInfo{
					Name: "a",
					GcpSAKey: &GcpServiceAccountKey{
						Email:        "a@example.com",
						PrivateKeyId: "b",
					},
					MetastoreID: "d",
				},
			},
		},
		Resource: ResourceStorageCredential(),
		Update:   true,
		ID:       "a",
		InstanceState: map[string]string{
			"name":                                   "a",
			"gcp_service_account_key.email":          "a@example.com",
			"gcp_service_account_key.private_key_id": "b",
			"gcp_service_account_key.private_key":    "",
			"comment":                                "c",
		},
		HCL: `
		name = "a"
		gcp_service_account_key {
			email = "a@example.com"
			private_key_id = "b"
			private_key = "abcdefg"
		}
		comment = "c"
		`,
	}.ApplyNoError(t)
}
