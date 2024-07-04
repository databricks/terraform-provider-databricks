package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
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
					AwsIamRole: &catalog.AwsIamRoleRequest{
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
					AwsIamRole: &catalog.AwsIamRoleResponse{
						RoleArn:    "def",
						ExternalId: "123",
					},
					MetastoreId: "d",
					Id:          "1234-5678",
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
		"storage_credential_id":      "1234-5678",
	})
}

func TestCreateIsolatedStorageCredential(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockStorageCredentialsAPI().EXPECT()
			e.Create(mock.Anything, catalog.CreateStorageCredential{
				Name: "a",
				AwsIamRole: &catalog.AwsIamRoleRequest{
					RoleArn: "def",
				},
				Comment: "c",
			}).Return(&catalog.StorageCredentialInfo{
				Name: "a",
				AwsIamRole: &catalog.AwsIamRoleResponse{
					RoleArn:    "def",
					ExternalId: "123",
				},
				MetastoreId: "d",
				Id:          "1234-5678",
				Owner:       "f",
			}, nil)
			e.Update(mock.Anything, catalog.UpdateStorageCredential{
				Name: "a",
				AwsIamRole: &catalog.AwsIamRoleRequest{
					RoleArn: "def",
				},
				Comment:       "c",
				IsolationMode: "ISOLATION_MODE_ISOLATED",
			}).Return(&catalog.StorageCredentialInfo{
				Name: "a",
				AwsIamRole: &catalog.AwsIamRoleResponse{
					RoleArn:    "def",
					ExternalId: "123",
				},
				MetastoreId:   "d",
				Id:            "1234-5678",
				Owner:         "f",
				IsolationMode: "ISOLATION_MODE_ISOLATED",
			}, nil)
			w.GetMockMetastoresAPI().EXPECT().Current(mock.Anything).Return(&catalog.MetastoreAssignment{
				MetastoreId: "e",
				WorkspaceId: 123456789101112,
			}, nil)
			w.GetMockWorkspaceBindingsAPI().EXPECT().UpdateBindings(mock.Anything, catalog.UpdateWorkspaceBindingsParameters{
				SecurableName: "a",
				SecurableType: "storage-credential",
				Add: []catalog.WorkspaceBinding{
					{
						WorkspaceId: int64(123456789101112),
						BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadWrite,
					},
				},
			}).Return(&catalog.WorkspaceBindingsResponse{
				Bindings: []catalog.WorkspaceBinding{
					{
						WorkspaceId: int64(123456789101112),
						BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadWrite,
					},
				},
			}, nil)
			e.GetByName(mock.Anything, "a").Return(&catalog.StorageCredentialInfo{
				Name: "a",
				AwsIamRole: &catalog.AwsIamRoleResponse{
					RoleArn:    "def",
					ExternalId: "123",
				},
				MetastoreId:   "d",
				Id:            "1234-5678",
				Owner:         "f",
				IsolationMode: "ISOLATION_MODE_ISOLATED",
			}, nil)
		},
		Resource: ResourceStorageCredential(),
		Create:   true,
		HCL: `
		name = "a"
		aws_iam_role {
			role_arn = "def"
		}
		comment = "c"
		isolation_mode = "ISOLATION_MODE_ISOLATED"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"aws_iam_role.0.external_id": "123",
		"aws_iam_role.0.role_arn":    "def",
		"name":                       "a",
		"storage_credential_id":      "1234-5678",
		"isolation_mode":             "ISOLATION_MODE_ISOLATED",
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
					AwsIamRole: &catalog.AwsIamRoleRequest{
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
					AwsIamRole: &catalog.AwsIamRoleRequest{
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
					AwsIamRole: &catalog.AwsIamRoleResponse{
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

func TestCreateAccountStorageCredentialWithOwner(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/account_id/metastores/metastore_id/storage-credentials",
				ExpectedRequest: &catalog.AccountsCreateStorageCredential{
					MetastoreId: "metastore_id",
					CredentialInfo: &catalog.CreateStorageCredential{
						Name: "storage_credential_name",
						AwsIamRole: &catalog.AwsIamRoleRequest{
							RoleArn: "arn:aws:iam::1234567890:role/MyRole-AJJHDSKSDF",
						},
					},
				},
				Response: catalog.AccountsStorageCredentialInfo{
					CredentialInfo: &catalog.StorageCredentialInfo{
						Name: "storage_credential_name",
					},
				},
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/accounts/account_id/metastores/metastore_id/storage-credentials/storage_credential_name",
				ExpectedRequest: &catalog.AccountsUpdateStorageCredential{
					CredentialInfo: &catalog.UpdateStorageCredential{
						Name:  "storage_credential_name",
						Owner: "administrators",
						AwsIamRole: &catalog.AwsIamRoleRequest{
							RoleArn: "arn:aws:iam::1234567890:role/MyRole-AJJHDSKSDF",
						},
					},
				},
				Response: &catalog.AccountsStorageCredentialInfo{
					CredentialInfo: &catalog.StorageCredentialInfo{
						Name: "storage_credential_name",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/account_id/metastores/metastore_id/storage-credentials/storage_credential_name?",
				Response: &catalog.AccountsStorageCredentialInfo{
					CredentialInfo: &catalog.StorageCredentialInfo{
						Name: "storage_credential_name",
						AwsIamRole: &catalog.AwsIamRoleResponse{
							RoleArn: "arn:aws:iam::1234567890:role/MyRole-AJJHDSKSDF",
						},
						Id: "1234-5678",
					},
				},
			},
		},
		Resource:  ResourceStorageCredential(),
		AccountID: "account_id",
		Create:    true,
		HCL: `
		name = "storage_credential_name"
		metastore_id = "metastore_id"
		aws_iam_role {
			role_arn = "arn:aws:iam::1234567890:role/MyRole-AJJHDSKSDF"
		}
		owner = "administrators"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"storage_credential_id": "1234-5678",
	})
}

func TestCreateStorageCredentialsReadOnly(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/storage-credentials",
				ExpectedRequest: catalog.CreateStorageCredential{
					Name: "a",
					AwsIamRole: &catalog.AwsIamRoleRequest{
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
					AwsIamRole: &catalog.AwsIamRoleRequest{
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
					AwsIamRole: &catalog.AwsIamRoleResponse{
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
					AwsIamRole: &catalog.AwsIamRoleRequest{
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
					AwsIamRole: &catalog.AwsIamRoleResponse{
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
					AwsIamRole: &catalog.AwsIamRoleRequest{
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
					AwsIamRole: &catalog.AwsIamRoleResponse{
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
					AwsIamRole: &catalog.AwsIamRoleRequest{
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
					AwsIamRole: &catalog.AwsIamRoleResponse{
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
					AwsIamRole: &catalog.AwsIamRoleRequest{
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
					AwsIamRole: &catalog.AwsIamRoleResponse{
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
					AzureManagedIdentity: &catalog.AzureManagedIdentityRequest{
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
					AzureManagedIdentity: &catalog.AzureManagedIdentityResponse{
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
					AzureManagedIdentity: &catalog.AzureManagedIdentityResponse{
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
					DatabricksGcpServiceAccount: &catalog.DatabricksGcpServiceAccountRequest{},
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
					AzureManagedIdentity: &catalog.AzureManagedIdentityResponse{
						AccessConnectorId: "CHANGED",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a?",
				Response: catalog.StorageCredentialInfo{
					Name: "a",
					AzureManagedIdentity: &catalog.AzureManagedIdentityResponse{
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

func TestUpdateAzStorageCredentialSpn(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/storage-credentials/a",
				ExpectedRequest: catalog.UpdateStorageCredential{
					Comment: "c",
					AzureServicePrincipal: &catalog.AzureServicePrincipal{
						ApplicationId: "SAME",
						DirectoryId:   "SAME",
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
						ApplicationId: "SAME",
						DirectoryId:   "SAME",
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
			application_id = "SAME"
			directory_id = "SAME"
			client_secret = "CHANGED"
		}
		comment = "c"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"azure_service_principal.0.application_id": "SAME",
		"azure_service_principal.0.directory_id":   "SAME",
		"azure_service_principal.0.client_secret":  "CHANGED",
	})
}
