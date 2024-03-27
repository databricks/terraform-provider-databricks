package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestStorageCredentialDataVerify(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockStorageCredentialsAPI().EXPECT()
			e.GetByName(mock.Anything, "abc").Return(
				&catalog.StorageCredentialInfo{
					Name:  "abc",
					Owner: "admin",
					AwsIamRole: &catalog.AwsIamRoleResponse{
						RoleArn: "test",
					},
					AzureManagedIdentity: &catalog.AzureManagedIdentity{
						AccessConnectorId: "test",
					},
					DatabricksGcpServiceAccount: &catalog.DatabricksGcpServiceAccountResponse{
						Email: "test",
					},
				},
				nil)
		},
		Resource:    DataSourceStorageCredential(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL: `
		name = "abc"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"storage_credential_info.0.owner":                                        "admin",
		"storage_credential_info.0.aws_iam_role.0.role_arn":                      "test",
		"storage_credential_info.0.azure_managed_identity.0.access_connector_id": "test",
		"storage_credential_info.0.databricks_gcp_service_account.0.email":       "test",
	})
}

func TestStorageCredentialDataError(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceStorageCredential(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
