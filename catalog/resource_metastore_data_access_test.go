package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
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
				Resource: "/api/2.1/unity-catalog/storage-credentials",
				ExpectedRequest: catalog.CreateStorageCredential{
					Name: "bcd",
					AwsIamRole: &catalog.AwsIamRoleRequest{
						RoleArn: "def",
					},
				},
				Response: catalog.StorageCredentialInfo{
					Id: "bcd",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/metastores/abc",
				ExpectedRequest: catalog.UpdateMetastore{
					StorageRootCredentialId: "bcd",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/bcd?",
				Response: catalog.StorageCredentialInfo{
					Name: "bcd",
					AwsIamRole: &catalog.AwsIamRoleResponse{
						RoleArn: "def",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/metastore_summary",
				Response: catalog.MetastoreInfo{
					StorageRootCredentialId: "bcd",
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
				Resource: "/api/2.1/unity-catalog/storage-credentials",
				ExpectedRequest: catalog.CreateStorageCredential{
					Name: "bcd",
					AzureManagedIdentity: &catalog.AzureManagedIdentityRequest{
						AccessConnectorId: "def",
						ManagedIdentityId: "/..../subscription",
					},
				},
				Response: catalog.StorageCredentialInfo{
					Id: "bcd",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/metastores/abc",
				ExpectedRequest: catalog.UpdateMetastore{
					StorageRootCredentialId: "bcd",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/bcd?",
				Response: catalog.StorageCredentialInfo{
					Name: "bcd",
					AzureManagedIdentity: &catalog.AzureManagedIdentityResponse{
						AccessConnectorId: "def",
						ManagedIdentityId: "/..../subscription",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/metastore_summary",
				Response: catalog.MetastoreInfo{
					StorageRootCredentialId: "bcd",
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
			managed_identity_id = "/..../subscription"
		}
		`,
	}.ApplyNoError(t)
}

func TestCreateDacWithDbGcpSA(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/storage-credentials",
				ExpectedRequest: catalog.CreateStorageCredential{
					Name:                        "bcd",
					DatabricksGcpServiceAccount: &catalog.DatabricksGcpServiceAccountRequest{},
				},
				Response: catalog.StorageCredentialInfo{
					Id: "bcd",
					DatabricksGcpServiceAccount: &catalog.DatabricksGcpServiceAccountResponse{
						Email: "a@example.com",
					},
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/metastores/abc",
				ExpectedRequest: catalog.UpdateMetastore{
					StorageRootCredentialId: "bcd",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/storage-credentials/bcd?",
				Response: catalog.StorageCredentialInfo{
					Name: "bcd",
					DatabricksGcpServiceAccount: &catalog.DatabricksGcpServiceAccountResponse{
						Email: "a@example.com",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/metastore_summary",
				Response: catalog.MetastoreInfo{
					StorageRootCredentialId: "bcd",
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

func TestCreateAccountDacWithAws(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/100/metastores/abc/storage-credentials",
				ExpectedRequest: catalog.AccountsCreateStorageCredential{
					MetastoreId: "abc",
					CredentialInfo: &catalog.CreateStorageCredential{
						Name: "bcd",
						AwsIamRole: &catalog.AwsIamRoleRequest{
							RoleArn: "def",
						},
					},
				},
				Response: catalog.AccountsStorageCredentialInfo{
					CredentialInfo: &catalog.StorageCredentialInfo{
						Id: "bcd",
					},
				},
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/accounts/100/metastores/abc",
				ExpectedRequest: catalog.AccountsUpdateMetastore{
					MetastoreId: "abc",
					MetastoreInfo: &catalog.UpdateMetastore{
						StorageRootCredentialId: "bcd",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/100/metastores/abc/storage-credentials/bcd?",
				Response: catalog.AccountsStorageCredentialInfo{
					CredentialInfo: &catalog.StorageCredentialInfo{
						Name: "bcd",
						AwsIamRole: &catalog.AwsIamRoleResponse{
							RoleArn: "def",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/100/metastores/abc?",
				Response: catalog.AccountsMetastoreInfo{
					MetastoreInfo: &catalog.MetastoreInfo{
						StorageRootCredentialId: "bcd",
					},
				},
			},
		},
		Create:    true,
		AccountID: "100",
		Resource:  ResourceMetastoreDataAccess(),
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

func TestCreateAccountDacWithAzMI(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/100/metastores/abc/storage-credentials",
				ExpectedRequest: catalog.AccountsCreateStorageCredential{
					MetastoreId: "abc",
					CredentialInfo: &catalog.CreateStorageCredential{
						Name: "bcd",
						AzureManagedIdentity: &catalog.AzureManagedIdentityRequest{
							AccessConnectorId: "def",
						},
					},
				},
				Response: catalog.AccountsStorageCredentialInfo{
					CredentialInfo: &catalog.StorageCredentialInfo{
						Id: "bcd",
					},
				},
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/accounts/100/metastores/abc",
				ExpectedRequest: catalog.AccountsUpdateMetastore{
					MetastoreId: "abc",
					MetastoreInfo: &catalog.UpdateMetastore{
						StorageRootCredentialId: "bcd",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/100/metastores/abc/storage-credentials/bcd?",
				Response: catalog.AccountsStorageCredentialInfo{
					CredentialInfo: &catalog.StorageCredentialInfo{
						Name: "bcd",
						AzureManagedIdentity: &catalog.AzureManagedIdentityResponse{
							AccessConnectorId: "def",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/100/metastores/abc?",
				Response: catalog.AccountsMetastoreInfo{
					MetastoreInfo: &catalog.MetastoreInfo{
						StorageRootCredentialId: "bcd",
					},
				},
			},
		},
		Create:    true,
		AccountID: "100",
		Resource:  ResourceMetastoreDataAccess(),
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

func TestCreateAccountDacWithDbGcpSA(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/100/metastores/abc/storage-credentials",
				ExpectedRequest: catalog.AccountsCreateStorageCredential{
					MetastoreId: "abc",
					CredentialInfo: &catalog.CreateStorageCredential{
						Name:                        "bcd",
						DatabricksGcpServiceAccount: &catalog.DatabricksGcpServiceAccountRequest{},
					},
				},
				Response: catalog.AccountsStorageCredentialInfo{
					CredentialInfo: &catalog.StorageCredentialInfo{
						Id: "bcd",
						DatabricksGcpServiceAccount: &catalog.DatabricksGcpServiceAccountResponse{
							Email: "a@example.com",
						},
					},
				},
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/accounts/100/metastores/abc",
				ExpectedRequest: catalog.AccountsUpdateMetastore{
					MetastoreId: "abc",
					MetastoreInfo: &catalog.UpdateMetastore{
						StorageRootCredentialId: "bcd",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/100/metastores/abc/storage-credentials/bcd?",
				Response: catalog.AccountsStorageCredentialInfo{
					CredentialInfo: &catalog.StorageCredentialInfo{
						Name: "bcd",
						DatabricksGcpServiceAccount: &catalog.DatabricksGcpServiceAccountResponse{
							Email: "a@example.com",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/100/metastores/abc?",
				Response: catalog.AccountsMetastoreInfo{
					MetastoreInfo: &catalog.MetastoreInfo{
						StorageRootCredentialId: "bcd",
					},
				},
			},
		},
		Create:    true,
		AccountID: "100",
		Resource:  ResourceMetastoreDataAccess(),
		HCL: `
		metastore_id = "abc"
		name = "bcd"
		is_default = true
		databricks_gcp_service_account {}
		`,
	}.ApplyAndExpectData(t,
		map[string]any{
			"databricks_gcp_service_account.#":       1,
			"databricks_gcp_service_account.0.email": "a@example.com",
		})
}

func TestUpdateDacWithOwner(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockAccountStorageCredentialsAPI().EXPECT()
			e.Update(mock.Anything, catalog.AccountsUpdateStorageCredential{
				MetastoreId:           "metastore_id",
				StorageCredentialName: "a",
				CredentialInfo: &catalog.UpdateStorageCredential{
					Owner: "updatedOwner",
					Name:  "a",
				},
			}).Return(&catalog.AccountsStorageCredentialInfo{
				CredentialInfo: &catalog.StorageCredentialInfo{
					Name:        "a",
					MetastoreId: "d",
					Comment:     "c",
					Owner:       "updatedOwner",
					AwsIamRole: &catalog.AwsIamRoleResponse{
						RoleArn: "INITIAL",
					},
				},
			}, nil)
			e.Update(mock.Anything, catalog.AccountsUpdateStorageCredential{
				MetastoreId:           "metastore_id",
				StorageCredentialName: "a",
				CredentialInfo: &catalog.UpdateStorageCredential{
					Name:    "a",
					Comment: "c",
					AwsIamRole: &catalog.AwsIamRoleRequest{
						RoleArn: "INITIAL",
					},
					Force: true,
				},
			}).Return(&catalog.AccountsStorageCredentialInfo{
				CredentialInfo: &catalog.StorageCredentialInfo{
					Name:        "a",
					MetastoreId: "d",
					Comment:     "c",
					Owner:       "updatedOwner",
					AwsIamRole: &catalog.AwsIamRoleResponse{
						RoleArn: "INITIAL",
					},
				},
			}, nil)
			e.GetByMetastoreIdAndStorageCredentialName(mock.Anything, "metastore_id", "a").Return(&catalog.AccountsStorageCredentialInfo{
				CredentialInfo: &catalog.StorageCredentialInfo{
					Name:        "a",
					MetastoreId: "d",
					Comment:     "c",
					Owner:       "updatedOwner",
					AwsIamRole: &catalog.AwsIamRoleResponse{
						RoleArn: "INITIAL",
					},
				},
			}, nil)
			a.GetMockAccountMetastoresAPI().EXPECT().GetByMetastoreId(mock.Anything, "metastore_id").Return(&catalog.AccountsMetastoreInfo{
				MetastoreInfo: &catalog.MetastoreInfo{
					StorageRootCredentialName: "a",
				},
			}, nil)
		},
		Resource:  ResourceMetastoreDataAccess(),
		Update:    true,
		ID:        "metastore_id|a",
		AccountID: "account_id",
		InstanceState: map[string]string{
			"name":                    "a",
			"comment":                 "c",
			"aws_iam_role.#":          "1",
			"aws_iam_role.0.role_arn": "INITIAL",
			"is_default":              "true",
		},
		HCL: `
		name = "a"
		comment = "c"
		aws_iam_role {
			role_arn = "INITIAL"
		}
		owner = "updatedOwner"
		force_update = true
		is_default = true
		`,
	}.ApplyNoError(t)
}

func TestDeleteDac(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockAccountStorageCredentialsAPI().EXPECT().Delete(mock.Anything, catalog.DeleteAccountStorageCredentialRequest{
				MetastoreId:           "metastore_id",
				StorageCredentialName: "a",
				Force:                 true,
			}).Return(nil)
		},
		Resource:  ResourceMetastoreDataAccess(),
		AccountID: "account_id",
		Delete:    true,
		ID:        "metastore_id|a",
		HCL: `
		name = "a"
		metastore_id = "metastore_id"
		aws_iam_role {
			role_arn = "def"
		}
		comment = "c"
		force_destroy = true
		`,
	}.ApplyNoError(t)
}
