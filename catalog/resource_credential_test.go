package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestCredentialsCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceCredential())
}

func TestCreateCredential(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockCredentialsAPI().EXPECT()
			e.CreateCredential(mock.Anything, catalog.CreateCredentialRequest{
				Name: "a",
				AwsIamRole: &catalog.AwsIamRole{
					RoleArn: "def",
				},
				Comment: "c",
				Purpose: "SERVICE",
			}).Return(&catalog.CredentialInfo{
				Name: "a",
				AwsIamRole: &catalog.AwsIamRole{
					RoleArn: "def",
				},
				Purpose: "SERVICE",
				Comment: "c",
			}, nil)
			e.GetCredentialByNameArg(mock.Anything, "a").Return(&catalog.CredentialInfo{
				Name: "a",
				AwsIamRole: &catalog.AwsIamRole{
					RoleArn:    "def",
					ExternalId: "123",
				},
				Purpose:       "SERVICE",
				MetastoreId:   "d",
				Id:            "1234-5678",
				Owner:         "f",
				IsolationMode: "ISOLATION_MODE_ISOLATED",
			}, nil)
		},
		Resource: ResourceCredential(),
		Create:   true,
		HCL: `
		name = "a"
		aws_iam_role {
			role_arn = "def"
		}
		purpose = "SERVICE"
		comment = "c"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"aws_iam_role.0.external_id": "123",
		"aws_iam_role.0.role_arn":    "def",
		"name":                       "a",
		"purpose":                    "SERVICE",
	})
}

func TestCreateIsolatedCredential(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockCredentialsAPI().EXPECT()
			e.CreateCredential(mock.Anything, catalog.CreateCredentialRequest{
				Name: "a",
				AwsIamRole: &catalog.AwsIamRole{
					RoleArn: "def",
				},
				Comment: "c",
				Purpose: "SERVICE",
			}).Return(&catalog.CredentialInfo{
				Name: "a",
				AwsIamRole: &catalog.AwsIamRole{
					RoleArn:    "def",
					ExternalId: "123",
				},
				Purpose:     "SERVICE",
				MetastoreId: "d",
				Id:          "1234-5678",
				Owner:       "f",
			}, nil)
			e.UpdateCredential(mock.Anything, catalog.UpdateCredentialRequest{
				NameArg: "a",
				AwsIamRole: &catalog.AwsIamRole{
					RoleArn: "def",
				},
				Comment:       "c",
				IsolationMode: "ISOLATION_MODE_ISOLATED",
			}).Return(&catalog.CredentialInfo{
				Name: "a",
				AwsIamRole: &catalog.AwsIamRole{
					RoleArn:    "def",
					ExternalId: "123",
				},
				Purpose:       "SERVICE",
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
				SecurableType: catalog.UpdateBindingsSecurableTypeCredential,
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
			e.GetCredentialByNameArg(mock.Anything, "a").Return(&catalog.CredentialInfo{
				Name: "a",
				AwsIamRole: &catalog.AwsIamRole{
					RoleArn:    "def",
					ExternalId: "123",
				},
				Purpose:       "SERVICE",
				MetastoreId:   "d",
				Id:            "1234-5678",
				Owner:         "f",
				IsolationMode: "ISOLATION_MODE_ISOLATED",
			}, nil)
		},
		Resource: ResourceCredential(),
		Create:   true,
		HCL: `
		name = "a"
		aws_iam_role {
			role_arn = "def"
		}
		comment = "c"
		purpose = "SERVICE"
		isolation_mode = "ISOLATION_MODE_ISOLATED"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"aws_iam_role.0.external_id": "123",
		"aws_iam_role.0.role_arn":    "def",
		"name":                       "a",
		"isolation_mode":             "ISOLATION_MODE_ISOLATED",
	})
}
