package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
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
				ExpectedRequest: catalog.CreateStorageCredential{
					Name: "a",
					AwsIamRole: &catalog.AwsIamRole{
						RoleArn: "def",
					},
					Comment: "c",
				},
				Response: catalog.StorageCredentialInfo{
					Name: "a",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a?",
				Response: catalog.StorageCredentialInfo{
					Name: "a",
					AwsIamRole: &catalog.AwsIamRole{
						RoleArn:    "def",
						ExternalId: "123",
					},
					MetastoreId: "d",
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
	}.ApplyAndExpectData(t, map[string]any{
		"aws_iam_role.0.external_id": "123",
		"aws_iam_role.0.role_arn":    "def",
		"name":                       "a",
	})
}

func TestCreateStorageCredentialWithOwner(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/storage-credentials",
				ExpectedRequest: catalog.CreateStorageCredential{
					Name: "a",
					AwsIamRole: &catalog.AwsIamRole{
						RoleArn: "def",
					},
					Comment: "c",
				},
				Response: catalog.StorageCredentialInfo{
					Name: "a",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				ExpectedRequest: catalog.UpdateStorageCredential{
					AwsIamRole: &catalog.AwsIamRole{
						RoleArn: "def",
					},
					Comment: "c",
					Owner:   "administrators",
				},
				Response: catalog.StorageCredentialInfo{
					Name: "a",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a?",
				Response: catalog.StorageCredentialInfo{
					Name: "a",
					AwsIamRole: &catalog.AwsIamRole{
						RoleArn: "def",
					},
					MetastoreId: "d",
					Owner:       "administrators",
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

func TestCreateStorageCredentialsReadOnly(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/storage-credentials",
				ExpectedRequest: catalog.CreateStorageCredential{
					Name: "a",
					AwsIamRole: &catalog.AwsIamRole{
						RoleArn: "def",
					},
					Comment:  "c",
					ReadOnly: true,
				},
				Response: catalog.StorageCredentialInfo{
					Name: "a",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				ExpectedRequest: catalog.UpdateStorageCredential{
					AwsIamRole: &catalog.AwsIamRole{
						RoleArn: "def",
					},
					Comment:  "c",
					ReadOnly: true,
				},
				Response: catalog.StorageCredentialInfo{
					Name: "a",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a?",
				Response: catalog.StorageCredentialInfo{
					Name: "a",
					AwsIamRole: &catalog.AwsIamRole{
						RoleArn: "def",
					},
					MetastoreId: "d",
					ReadOnly:    true,
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
		read_only = true
		`,
	}.ApplyNoError(t)
}

func TestUpdateStorageCredentials(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				ExpectedRequest: catalog.UpdateStorageCredential{
					AwsIamRole: &catalog.AwsIamRole{
						RoleArn: "CHANGED",
					},
					Comment: "c",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a?",
				Response: catalog.StorageCredentialInfo{
					Name: "a",
					AwsIamRole: &catalog.AwsIamRole{
						RoleArn: "CHANGED",
					},
					MetastoreId: "d",
					Comment:     "c",
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

func TestUpdateStorageCredentialsWithOwnerOnly(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				ExpectedRequest: catalog.UpdateStorageCredential{
					Owner: "updatedOwner",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				ExpectedRequest: catalog.UpdateStorageCredential{
					Comment: "c",
					AwsIamRole: &catalog.AwsIamRole{
						RoleArn: "INITIAL",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a?",
				Response: catalog.StorageCredentialInfo{
					Name:        "a",
					MetastoreId: "d",
					Comment:     "c",
					Owner:       "updatedOwner",
					AwsIamRole: &catalog.AwsIamRole{
						RoleArn: "INITIAL",
					},
				},
			},
		},
		Resource: ResourceStorageCredential(),
		Update:   true,
		ID:       "a",
		InstanceState: map[string]string{
			"name":                    "a",
			"comment":                 "c",
			"aws_iam_role.#":          "1",
			"aws_iam_role.0.role_arn": "INITIAL",
		},
		HCL: `
		name = "a"
		comment = "c"
		aws_iam_role {
			role_arn = "INITIAL"
		}
		owner = "updatedOwner"
		`,
	}.ApplyNoError(t)
}

func TestUpdateStorageCredentialsWithOwnerAndOtherFields(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				ExpectedRequest: catalog.UpdateStorageCredential{
					Owner: "updatedOwner",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				ExpectedRequest: catalog.UpdateStorageCredential{
					Comment: "e",
					AwsIamRole: &catalog.AwsIamRole{
						RoleArn: "CHANGED",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a?",
				Response: catalog.StorageCredentialInfo{
					Name:        "a",
					MetastoreId: "d",
					Comment:     "e",
					Owner:       "updatedOwner",
					AwsIamRole: &catalog.AwsIamRole{
						RoleArn: "CHANGED",
					},
				},
			},
		},
		Resource: ResourceStorageCredential(),
		Update:   true,
		ID:       "a",
		InstanceState: map[string]string{
			"name":                    "a",
			"comment":                 "c",
			"aws_iam_role.#":          "1",
			"aws_iam_role.0.role_arn": "INITIAL",
		},
		HCL: `
		name = "a"
		comment = "e"
		aws_iam_role {
			role_arn = "CHANGED"
		}
		owner = "updatedOwner"
		`,
	}.ApplyNoError(t)
}

func TestUpdateStorageCredentialsRollback(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				ExpectedRequest: catalog.UpdateStorageCredential{
					Owner: "updatedOwner",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				ExpectedRequest: catalog.UpdateStorageCredential{
					Comment: "d",
					AwsIamRole: &catalog.AwsIamRole{
						RoleArn: "CHANGED",
					},
				},
				Response: apierr.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				ExpectedRequest: catalog.UpdateStorageCredential{
					Owner: "admin",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a?",
				Response: catalog.StorageCredentialInfo{
					Name:        "a",
					MetastoreId: "d",
					Comment:     "c",
					Owner:       "admin",
					AwsIamRole: &catalog.AwsIamRole{
						RoleArn: "INITIAL",
					},
				},
			},
		},
		Resource: ResourceStorageCredential(),
		Update:   true,
		ID:       "a",
		InstanceState: map[string]string{
			"name":                    "a",
			"comment":                 "c",
			"owner":                   "admin",
			"aws_iam_role.#":          "1",
			"aws_iam_role.0.role_arn": "INITIAL",
		},
		HCL: `
		name = "a"
		comment = "d"
		aws_iam_role {
			role_arn = "CHANGED"
		}
		owner = "updatedOwner"
		`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected happened")
}

func TestCreateStorageCredentialWithAzMI(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/storage-credentials",
				ExpectedRequest: catalog.CreateStorageCredential{
					Name: "a",
					AzureManagedIdentity: &catalog.AzureManagedIdentity{
						AccessConnectorId: "def",
					},
					Comment: "c",
				},
				Response: catalog.StorageCredentialInfo{
					Name: "a",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				ExpectedRequest: catalog.UpdateStorageCredential{
					AzureManagedIdentity: &catalog.AzureManagedIdentity{
						AccessConnectorId: "def",
					},
					Comment: "c",
				},
				Response: catalog.StorageCredentialInfo{
					Name: "a",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a?",
				Response: catalog.StorageCredentialInfo{
					Name: "a",
					AzureManagedIdentity: &catalog.AzureManagedIdentity{
						AccessConnectorId: "def",
					},
					MetastoreId: "d",
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

func TestUpdateAzStorageCredentials(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				ExpectedRequest: catalog.UpdateStorageCredential{
					Comment: "c",
					AzureServicePrincipal: &catalog.AzureServicePrincipal{
						DirectoryId:   "CHANGED",
						ApplicationId: "CHANGED",
						ClientSecret:  "CHANGED",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a?",
				Response: catalog.StorageCredentialInfo{
					Name: "a",
					AzureServicePrincipal: &catalog.AzureServicePrincipal{
						DirectoryId:   "CHANGED",
						ApplicationId: "CHANGED",
						ClientSecret:  "CHANGED",
					},
					MetastoreId: "d",
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

func TestCreateStorageCredentialWithDbGcpSA(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/storage-credentials",
				ExpectedRequest: catalog.CreateStorageCredential{
					Name:                        "a",
					Comment:                     "c",
					DatabricksGcpServiceAccount: struct{}{},
				},
				Response: catalog.StorageCredentialInfo{
					Name: "a",
					DatabricksGcpServiceAccount: &catalog.DatabricksGcpServiceAccountResponse{
						Email: "a@example.com",
					},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				ExpectedRequest: catalog.UpdateStorageCredential{
					Comment: "c",
				},
				Response: catalog.StorageCredentialInfo{
					Name: "a",
					DatabricksGcpServiceAccount: &catalog.DatabricksGcpServiceAccountResponse{
						Email: "a@example.com",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a?",
				Response: catalog.StorageCredentialInfo{
					Name: "a",
					DatabricksGcpServiceAccount: &catalog.DatabricksGcpServiceAccountResponse{
						Email: "a@example.com",
					},
					MetastoreId: "d",
				},
			},
		},
		Resource: ResourceStorageCredential(),
		Create:   true,
		HCL: `
		name = "a"
		databricks_gcp_service_account {}
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
				ExpectedRequest: catalog.UpdateStorageCredential{
					Comment: "c",
					AzureManagedIdentity: &catalog.AzureManagedIdentity{
						AccessConnectorId: "CHANGED",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a?",
				Response: catalog.StorageCredentialInfo{
					Name: "a",
					AzureManagedIdentity: &catalog.AzureManagedIdentity{
						AccessConnectorId: "CHANGED",
					},
					MetastoreId: "d",
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
