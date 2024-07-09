package catalog

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestExternalLocationCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceExternalLocation())
}

func TestCreateExternalLocation(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/external-locations",
				ExpectedRequest: catalog.CreateExternalLocation{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
				},
				Response: catalog.ExternalLocationInfo{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc?",
				Response: catalog.ExternalLocationInfo{
					Owner:       "efg",
					MetastoreId: "fgh",
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Create:   true,
		HCL: `
		name = "abc"
		url = "s3://foo/bar"
		credential_name = "bcd"
		comment = "def"
		`,
	}.ApplyNoError(t)
}

func TestCreateIsolatedExternalLocation(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockExternalLocationsAPI().EXPECT()
			e.Create(mock.Anything, catalog.CreateExternalLocation{
				Name:           "abc",
				Url:            "s3://foo/bar",
				CredentialName: "bcd",
				Comment:        "def",
			}).Return(&catalog.ExternalLocationInfo{
				Name:           "abc",
				Url:            "s3://foo/bar",
				CredentialName: "bcd",
				Comment:        "def",
				MetastoreId:    "e",
				Owner:          "f",
			}, nil)
			e.Update(mock.Anything, catalog.UpdateExternalLocation{
				Name:           "abc",
				Url:            "s3://foo/bar",
				CredentialName: "bcd",
				Comment:        "def",
				IsolationMode:  "ISOLATED",
			}).Return(&catalog.ExternalLocationInfo{
				Name:           "abc",
				Url:            "s3://foo/bar",
				CredentialName: "bcd",
				Comment:        "def",
				IsolationMode:  "ISOLATED",
				MetastoreId:    "e",
				Owner:          "f",
			}, nil)
			w.GetMockMetastoresAPI().EXPECT().Current(mock.Anything).Return(&catalog.MetastoreAssignment{
				MetastoreId: "e",
				WorkspaceId: 123456789101112,
			}, nil)
			w.GetMockWorkspaceBindingsAPI().EXPECT().UpdateBindings(mock.Anything, catalog.UpdateWorkspaceBindingsParameters{
				SecurableName: "abc",
				SecurableType: "external-location",
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
			e.GetByName(mock.Anything, "abc").Return(&catalog.ExternalLocationInfo{
				Name:           "abc",
				Url:            "s3://foo/bar",
				CredentialName: "bcd",
				Comment:        "def",
				IsolationMode:  "ISOLATED",
				MetastoreId:    "e",
				Owner:          "f",
			}, nil)
		},
		Resource: ResourceExternalLocation(),
		Create:   true,
		HCL: `
		name = "abc"
		url = "s3://foo/bar"
		credential_name = "bcd"
		comment = "def"
		isolation_mode = "ISOLATED"
		`,
	}.ApplyNoError(t)
}

func TestCreateExternalLocationWithOwner(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/external-locations",
				ExpectedRequest: catalog.CreateExternalLocation{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
				},
				Response: catalog.ExternalLocationInfo{
					Name:           "abc",
					Url:            "s3://foo/bar",
					Owner:          "x",
					CredentialName: "bcd",
					Comment:        "def",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
					Owner:          "administrators",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc?",
				Response: catalog.ExternalLocationInfo{
					Owner:       "administrators",
					MetastoreId: "fgh",
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Create:   true,
		HCL: `
		name = "abc"
		url = "s3://foo/bar"
		credential_name = "bcd"
		owner = "administrators"
		comment = "def"
		`,
	}.ApplyNoError(t)
}

func TestCreateExternalLocationReadOnly(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/external-locations",
				ExpectedRequest: catalog.CreateExternalLocation{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
					ReadOnly:       true,
				},
				Response: catalog.ExternalLocationInfo{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
					ReadOnly:       true,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc?",
				Response: catalog.ExternalLocationInfo{
					Owner:       "efg",
					MetastoreId: "fgh",
					ReadOnly:    true,
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Create:   true,
		HCL: `
		name = "abc"
		url = "s3://foo/bar"
		credential_name = "bcd"
		comment = "def"
		read_only = true
		`,
	}.ApplyNoError(t)
}

func TestCreateExternalLocationWithAPAndEncryptionDetails(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/external-locations",
				ExpectedRequest: catalog.CreateExternalLocation{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					AccessPoint:    "some_access_point",
					EncryptionDetails: &catalog.EncryptionDetails{
						SseEncryptionDetails: &catalog.SseEncryptionDetails{
							Algorithm:    "AWS_SSE_KMS",
							AwsKmsKeyArn: "some_key_arn",
						},
					},
					Comment: "def",
				},
				Response: catalog.ExternalLocationInfo{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					AccessPoint:    "some_access_point",
					EncryptionDetails: &catalog.EncryptionDetails{
						SseEncryptionDetails: &catalog.SseEncryptionDetails{
							Algorithm:    "AWS_SSE_KMS",
							AwsKmsKeyArn: "some_key_arn",
						},
					},
					Comment: "def",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc?",
				Response: catalog.ExternalLocationInfo{
					Owner:       "efg",
					MetastoreId: "fgh",
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Create:   true,
		HCL: `
		name = "abc"
		url = "s3://foo/bar"
		credential_name = "bcd"
		comment = "def"
		access_point = "some_access_point"
	    encryption_details {
          sse_encryption_details {
            algorithm     = "AWS_SSE_KMS"
            aws_kms_key_arn = "some_key_arn"
		  }
        }
		`,
	}.ApplyNoError(t)
}

func TestUpdateExternalLocation(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc?",
				Response: catalog.ExternalLocationInfo{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Update:   true,
		ID:       "abc",
		InstanceState: map[string]string{
			"name":            "abc",
			"url":             "s3://foo/bar",
			"credential_name": "abc",
			"comment":         "def",
		},
		HCL: `
		name = "abc"
		url = "s3://foo/bar"
		credential_name = "bcd"
		comment = "def"
		`,
	}.ApplyNoError(t)
}

func TestUpdateExternalLocationOnlyOwner(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Owner: "updatedOwner",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Url:            "s3://foo/bar",
					CredentialName: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc?",
				Response: catalog.ExternalLocationInfo{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
					Owner:          "updatedOwner",
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Update:   true,
		ID:       "abc",
		InstanceState: map[string]string{
			"name":            "abc",
			"url":             "s3://foo/bar",
			"credential_name": "abc",
			"comment":         "def",
			"owner":           "administrators",
		},
		HCL: `
		name = "abc"
		url = "s3://foo/bar",
		owner = "updatedOwner"
		credential_name = "abc",
		`,
	}.ApplyNoError(t)
}

func TestUpdateExternalLocationOwnerAndOtherFields(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Owner: "updatedOwner",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Url:            "s3://foo/bar",
					CredentialName: "xyz",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc?",
				Response: catalog.ExternalLocationInfo{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
					Owner:          "updatedOwner",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc?",
				Response: catalog.ExternalLocationInfo{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "xyz",
					Comment:        "def",
					Owner:          "updatedOwner",
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Update:   true,
		ID:       "abc",
		InstanceState: map[string]string{
			"name":            "abc",
			"url":             "s3://foo/bar",
			"credential_name": "abc",
			"comment":         "def",
			"owner":           "administrators",
		},
		HCL: `
		name = "abc"
		url = "s3://foo/bar",
		owner = "updatedOwner"
		credential_name = "xyz",
		`,
	}.ApplyNoError(t)
}

func TestUpdateExternalLocationRollback(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Owner: "updatedOwner",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Url:            "s3://foo/bar",
					CredentialName: "xyz",
				},
				Response: common.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Owner: "administrators",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc?",
				Response: catalog.ExternalLocationInfo{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "abc",
					Comment:        "def",
					Owner:          "administrators",
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Update:   true,
		ID:       "abc",
		InstanceState: map[string]string{
			"name":            "abc",
			"url":             "s3://foo/bar",
			"credential_name": "abc",
			"comment":         "def",
			"owner":           "administrators",
		},
		HCL: `
		name = "abc"
		url = "s3://foo/bar",
		owner = "updatedOwner"
		credential_name = "xyz",
		`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected happened")
}

func TestUpdateExternalLocationRollbackError(t *testing.T) {
	serverErrMessage := "Something unexpected happened"
	rollbackErrMessage := "Internal error happened"
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Owner: "updatedOwner",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Url:            "s3://foo/bar",
					CredentialName: "xyz",
				},
				Response: common.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Owner: "administrators",
				},
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceExternalLocation(),
		Update:   true,
		ID:       "abc",
		InstanceState: map[string]string{
			"name":            "abc",
			"url":             "s3://foo/bar",
			"credential_name": "abc",
			"comment":         "def",
			"owner":           "administrators",
		},
		HCL: `
		name = "abc"
		url = "s3://foo/bar",
		owner = "updatedOwner"
		credential_name = "xyz",
		`,
	}.Apply(t)
	errOccurred := fmt.Sprintf("%s. Owner rollback also failed: %s", serverErrMessage, rollbackErrMessage)
	qa.AssertErrorStartsWith(t, err, errOccurred)
}

func TestUpdateExternalLocationForce(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/external-locations/abc",
				ExpectedRequest: catalog.UpdateExternalLocation{
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
					Force:          true,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/external-locations/abc?",
				Response: catalog.ExternalLocationInfo{
					Name:           "abc",
					Url:            "s3://foo/bar",
					CredentialName: "bcd",
					Comment:        "def",
				},
			},
		},
		Resource: ResourceExternalLocation(),
		Update:   true,
		ID:       "abc",
		InstanceState: map[string]string{
			"name":            "abc",
			"url":             "s3://foo/bar",
			"credential_name": "abc",
			"comment":         "def",
		},
		HCL: `
		name = "abc"
		url = "s3://foo/bar"
		credential_name = "bcd"
		comment = "def"
		force_update = true
		`,
	}.ApplyNoError(t)
}
