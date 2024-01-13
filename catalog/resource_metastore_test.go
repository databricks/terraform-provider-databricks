package catalog

import (
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestMetastoreCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceMetastore())
}

func TestCreateMetastore(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockMetastoresAPI().EXPECT()
			e.Create(mock.Anything, catalog.CreateMetastore{
				StorageRoot: "s3://b",
				Name:        "a",
			}).Return(&catalog.MetastoreInfo{
				MetastoreId: "abc",
			}, nil)
			e.Update(mock.Anything, catalog.UpdateMetastore{
				Id:   "abc",
				Name: "a",
			}).Return(&catalog.MetastoreInfo{
				Name: "a",
			}, nil)
			e.GetById(mock.Anything, "abc").Return(&catalog.MetastoreInfo{
				StorageRoot: "s3://b/abc",
				Name:        "a",
			}, nil)
		},
		Resource: ResourceMetastore(),
		Create:   true,
		HCL: `
		name = "a"
		storage_root = "s3://b"
		`,
	}.ApplyNoError(t)
}

func TestCreateMetastoreWithOwner(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockMetastoresAPI().EXPECT()
			e.Create(mock.Anything, catalog.CreateMetastore{
				StorageRoot: "s3://b",
				Name:        "a",
			}).Return(&catalog.MetastoreInfo{
				MetastoreId: "abc",
			}, nil)
			e.Update(mock.Anything, catalog.UpdateMetastore{
				Id:    "abc",
				Name:  "a",
				Owner: "administrators",
			}).Return(&catalog.MetastoreInfo{
				Name:  "a",
				Owner: "administrators",
			}, nil)
			e.GetById(mock.Anything, "abc").Return(&catalog.MetastoreInfo{
				StorageRoot: "s3://b/abc",
				Name:        "a",
				Owner:       "administrators",
			}, nil)
		},
		Resource: ResourceMetastore(),
		Create:   true,
		HCL: `
		name = "a"
		storage_root = "s3://b"
		owner = "administrators"
		`,
	}.ApplyNoError(t)
}

func TestCreateMetastore_DeltaSharing(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockMetastoresAPI().EXPECT()
			e.Create(mock.Anything, catalog.CreateMetastore{
				StorageRoot: "s3://b",
				Name:        "a",
			}).Return(&catalog.MetastoreInfo{
				MetastoreId: "abc",
			}, nil)
			e.Update(mock.Anything, catalog.UpdateMetastore{
				Id:                "abc",
				Name:              "a",
				Owner:             "administrators",
				DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
				DeltaSharingRecipientTokenLifetimeInSeconds: 0,
				DeltaSharingOrganizationName:                "acme",
				ForceSendFields:                             []string{"DeltaSharingRecipientTokenLifetimeInSeconds"},
			}).Return(&catalog.MetastoreInfo{
				Name:              "a",
				Owner:             "administrators",
				DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
				DeltaSharingRecipientTokenLifetimeInSeconds: 0,
				DeltaSharingOrganizationName:                "acme",
				ForceSendFields:                             []string{"DeltaSharingRecipientTokenLifetimeInSeconds"},
			}, nil)
			e.GetById(mock.Anything, "abc").Return(&catalog.MetastoreInfo{
				StorageRoot:       "s3://b/abc",
				Name:              "a",
				Owner:             "administrators",
				DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
				DeltaSharingRecipientTokenLifetimeInSeconds: 0,
				DeltaSharingOrganizationName:                "acme",
				ForceSendFields:                             []string{"DeltaSharingRecipientTokenLifetimeInSeconds"},
			}, nil)
		},
		Resource: ResourceMetastore(),
		Create:   true,
		HCL: `
		name = "a"
		storage_root = "s3://b"
		owner = "administrators"
		delta_sharing_scope = "INTERNAL_AND_EXTERNAL"
		delta_sharing_recipient_token_lifetime_in_seconds = 0
		delta_sharing_organization_name = "acme"
		`,
	}.ApplyNoError(t)
}

func TestDeleteMetastore(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockMetastoresAPI().EXPECT().Delete(mock.Anything, catalog.DeleteMetastoreRequest{Id: "abc"}).Return(nil)
		},
		Resource: ResourceMetastore(),
		Delete:   true,
		ID:       "abc",
		HCL: `
		name = "a"
		storage_root = "s3://b"
		`,
	}.ApplyNoError(t)
}

func TestForceDeleteMetastore(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(mwc *mocks.MockWorkspaceClient) {
			mwc.GetMockMetastoresAPI().EXPECT().Delete(mock.Anything, catalog.DeleteMetastoreRequest{
				Id:    "abc",
				Force: true,
			}).Return(nil)
		},
		Resource: ResourceMetastore(),
		Delete:   true,
		ID:       "abc",
		HCL: `
		name = "a"
		storage_root = "s3://b"

		force_destroy = true
		`,
	}.ApplyNoError(t)
}

func TestUpdateMetastore_NoChanges(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockMetastoresAPI().EXPECT()
			e.GetById(mock.Anything, "abc").Return(&catalog.MetastoreInfo{
				StorageRoot: "s3://b/abc",
				Name:        "abc",
			}, nil)
			e.Update(mock.Anything, catalog.UpdateMetastore{
				Id:                "abc",
				Name:              "abc",
				DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
				DeltaSharingRecipientTokenLifetimeInSeconds: 1002,
				ForceSendFields: []string{"DeltaSharingRecipientTokenLifetimeInSeconds"},
			}).Return(&catalog.MetastoreInfo{
				Name:              "a",
				DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
				DeltaSharingRecipientTokenLifetimeInSeconds: 1002,
			}, nil)
		},
		Resource:    ResourceMetastore(),
		ID:          "abc",
		Update:      true,
		RequiresNew: true,
		InstanceState: map[string]string{
			"name":                "abc",
			"storage_root":        "s3:/a",
			"owner":               "admin",
			"delta_sharing_scope": "INTERNAL_AND_EXTERNAL",
			"delta_sharing_recipient_token_lifetime_in_seconds": "1002",
		},
		HCL: `
		name = "abc"
		storage_root = "s3:/a"
		owner = "admin"
		delta_sharing_scope = "INTERNAL_AND_EXTERNAL"
		delta_sharing_recipient_token_lifetime_in_seconds = 1002
		`,
	}.ApplyNoError(t)
}

func TestUpdateMetastore_OwnerChanges(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockMetastoresAPI().EXPECT()
			e.Update(mock.Anything, catalog.UpdateMetastore{
				Id:    "abc",
				Owner: "updatedOwner",
			}).Return(&catalog.MetastoreInfo{
				Name:  "abc",
				Owner: "updatedOwner",
			}, nil)
			e.Update(mock.Anything, catalog.UpdateMetastore{
				Id:                "abc",
				Name:              "abc",
				DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
				DeltaSharingRecipientTokenLifetimeInSeconds: 1002,
				ForceSendFields: []string{"DeltaSharingRecipientTokenLifetimeInSeconds"},
			}).Return(&catalog.MetastoreInfo{
				Name:              "a",
				Owner:             "updatedOwner",
				DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
				DeltaSharingRecipientTokenLifetimeInSeconds: 1002,
			}, nil)
			e.GetById(mock.Anything, "abc").Return(&catalog.MetastoreInfo{
				Name:              "abc",
				Owner:             "updatedOwner",
				DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
				DeltaSharingRecipientTokenLifetimeInSeconds: 1002,
			}, nil)
		},
		Resource:    ResourceMetastore(),
		ID:          "abc",
		Update:      true,
		RequiresNew: true,
		InstanceState: map[string]string{
			"name":                "abc",
			"storage_root":        "s3:/a",
			"owner":               "admin",
			"delta_sharing_scope": "INTERNAL_AND_EXTERNAL",
			"delta_sharing_recipient_token_lifetime_in_seconds": "1002",
		},
		HCL: `
		name = "abc"
		storage_root = "s3:/a"
		owner = "updatedOwner"
		delta_sharing_scope = "INTERNAL_AND_EXTERNAL"
		delta_sharing_recipient_token_lifetime_in_seconds = 1002
		`,
	}.ApplyNoError(t)
}

func TestUpdateMetastore_Rollback(t *testing.T) {
	_, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockMetastoresAPI().EXPECT()
			e.Update(mock.Anything, catalog.UpdateMetastore{
				Id:    "abc",
				Owner: "updatedOwner",
			}).Return(&catalog.MetastoreInfo{
				Name:  "abc",
				Owner: "updatedOwner",
			}, nil)
			e.Update(mock.Anything, catalog.UpdateMetastore{
				Id:                "abc",
				Name:              "abc",
				DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
				DeltaSharingRecipientTokenLifetimeInSeconds: 1002,
				ForceSendFields: []string{"DeltaSharingRecipientTokenLifetimeInSeconds"},
			}).Return(nil, errors.New("Something unexpected happened"))
			e.Update(mock.Anything, catalog.UpdateMetastore{
				Id:    "abc",
				Owner: "admin",
			}).Return(&catalog.MetastoreInfo{
				Name:  "abc",
				Owner: "admin",
			}, nil)
		},
		Resource:    ResourceMetastore(),
		ID:          "abc",
		Update:      true,
		RequiresNew: true,
		InstanceState: map[string]string{
			"name":                "abc",
			"storage_root":        "s3:/a",
			"owner":               "admin",
			"delta_sharing_scope": "INTERNAL_AND_EXTERNAL",
			"delta_sharing_recipient_token_lifetime_in_seconds": "1002",
		},
		HCL: `
		name = "abc"
		storage_root = "s3:/a"
		owner = "updatedOwner"
		delta_sharing_scope = "INTERNAL_AND_EXTERNAL"
		delta_sharing_recipient_token_lifetime_in_seconds = 1002
		`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected happened")
}

func TestUpdateMetastore_DeltaSharingScopeOnly(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockMetastoresAPI().EXPECT()
			e.Update(mock.Anything, catalog.UpdateMetastore{
				Id:                "abc",
				Name:              "abc",
				DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
				DeltaSharingRecipientTokenLifetimeInSeconds: 1002,
				ForceSendFields: []string{"DeltaSharingRecipientTokenLifetimeInSeconds"},
			}).Return(&catalog.MetastoreInfo{
				Name:              "abc",
				DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
				DeltaSharingRecipientTokenLifetimeInSeconds: 1002,
			}, nil)
			e.GetById(mock.Anything, "abc").Return(&catalog.MetastoreInfo{
				Name:              "abc",
				DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
				DeltaSharingRecipientTokenLifetimeInSeconds: 1002,
			}, nil)
		},
		Resource:    ResourceMetastore(),
		ID:          "abc",
		Update:      true,
		RequiresNew: true,
		InstanceState: map[string]string{
			"name":                "abc",
			"storage_root":        "s3:/a",
			"owner":               "admin",
			"delta_sharing_scope": "INTERNAL",
			"delta_sharing_recipient_token_lifetime_in_seconds": "1002",
		},
		HCL: `
		name = "abc"
		storage_root = "s3:/a"
		owner = "admin"
		delta_sharing_scope = "INTERNAL_AND_EXTERNAL"
		delta_sharing_recipient_token_lifetime_in_seconds = 1002
		`,
	}.ApplyNoError(t)
}

func TestCreateAccountMetastore(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockAccountMetastoresAPI().EXPECT()
			e.Create(mock.Anything, catalog.AccountsCreateMetastore{
				MetastoreInfo: &catalog.CreateMetastore{
					StorageRoot: "s3://b",
					Name:        "a",
				},
			}).Return(&catalog.AccountsMetastoreInfo{
				MetastoreInfo: &catalog.MetastoreInfo{
					MetastoreId: "abc",
				},
			}, nil)
			e.Update(mock.Anything, catalog.AccountsUpdateMetastore{
				MetastoreId: "abc",
				MetastoreInfo: &catalog.UpdateMetastore{
					Name: "a",
				},
			}).Return(&catalog.AccountsMetastoreInfo{
				MetastoreInfo: &catalog.MetastoreInfo{
					Name: "a",
				},
			}, nil)
			e.GetByMetastoreId(mock.Anything, "abc").Return(&catalog.AccountsMetastoreInfo{
				MetastoreInfo: &catalog.MetastoreInfo{
					StorageRoot: "s3://b/abc",
					Name:        "a",
				},
			}, nil)
		},
		Resource:  ResourceMetastore(),
		AccountID: "100",
		Create:    true,
		HCL: `
		name = "a"
		storage_root = "s3://b"
		`,
	}.ApplyNoError(t)
}

func TestCreateAccountMetastoreWithOwner(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/100/metastores",
				ExpectedRequest: catalog.AccountsCreateMetastore{
					MetastoreInfo: &catalog.CreateMetastore{
						StorageRoot: "s3://b",
						Name:        "a",
					},
				},
				Response: catalog.AccountsMetastoreInfo{
					MetastoreInfo: &catalog.MetastoreInfo{
						MetastoreId: "abc",
					},
				},
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/accounts/100/metastores/abc",
				ExpectedRequest: catalog.AccountsUpdateMetastore{
					MetastoreInfo: &catalog.UpdateMetastore{
						Name:  "a",
						Owner: "administrators",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/100/metastores/abc?",
				Response: catalog.AccountsMetastoreInfo{
					MetastoreInfo: &catalog.MetastoreInfo{
						StorageRoot: "s3://b/abc",
						Name:        "a",
						Owner:       "administrators",
					},
				},
			},
		},
		Resource:  ResourceMetastore(),
		AccountID: "100",
		Create:    true,
		HCL: `
		name = "a"
		storage_root = "s3://b"
		owner = "administrators"
		`,
	}.ApplyNoError(t)
}

func TestCreateAccountMetastore_DeltaSharing(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/100/metastores",
				ExpectedRequest: catalog.AccountsCreateMetastore{
					MetastoreInfo: &catalog.CreateMetastore{
						StorageRoot: "s3://b",
						Name:        "a",
					},
				},
				Response: catalog.AccountsMetastoreInfo{
					MetastoreInfo: &catalog.MetastoreInfo{
						MetastoreId: "abc",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/100/metastores/abc?",
				Response: catalog.AccountsMetastoreInfo{
					MetastoreInfo: &catalog.MetastoreInfo{
						StorageRoot: "s3://b/abc",
						Name:        "a",
					},
				},
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/accounts/100/metastores/abc",
				ExpectedRequest: catalog.AccountsUpdateMetastore{
					MetastoreInfo: &catalog.UpdateMetastore{
						Name:              "a",
						Owner:             "administrators",
						DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
						DeltaSharingRecipientTokenLifetimeInSeconds: 0,
						DeltaSharingOrganizationName:                "acme",
						ForceSendFields:                             []string{"DeltaSharingRecipientTokenLifetimeInSeconds"},
					},
				},
			},
		},
		Resource:  ResourceMetastore(),
		AccountID: "100",
		Create:    true,
		HCL: `
		name = "a"
		storage_root = "s3://b"
		owner = "administrators"
		delta_sharing_scope = "INTERNAL_AND_EXTERNAL"
		delta_sharing_recipient_token_lifetime_in_seconds = 0
		delta_sharing_organization_name = "acme"
		`,
	}.ApplyNoError(t)
}

func TestDeleteAccountMetastore(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/100/metastores/abc?",
			},
		},
		Resource:  ResourceMetastore(),
		AccountID: "100",
		Delete:    true,
		ID:        "abc",
		HCL: `
		name = "a"
		storage_root = "s3://b"
		`,
	}.ApplyNoError(t)
}

func TestUpdateAccountMetastore_NoChanges(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/100/metastores/abc?",
				Response: catalog.AccountsMetastoreInfo{
					MetastoreInfo: &catalog.MetastoreInfo{
						StorageRoot: "s3://b/abc",
						Name:        "a",
					},
				},
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/accounts/100/metastores/abc",
				ExpectedRequest: catalog.AccountsUpdateMetastore{
					MetastoreInfo: &catalog.UpdateMetastore{
						Name:              "abc",
						DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
						DeltaSharingRecipientTokenLifetimeInSeconds: 1002,
					},
				},
			},
		},
		Resource:    ResourceMetastore(),
		AccountID:   "100",
		ID:          "abc",
		Update:      true,
		RequiresNew: true,
		InstanceState: map[string]string{
			"name":                "abc",
			"storage_root":        "s3:/a",
			"owner":               "admin",
			"delta_sharing_scope": "INTERNAL_AND_EXTERNAL",
			"delta_sharing_recipient_token_lifetime_in_seconds": "1002",
		},
		HCL: `
		name = "abc"
		storage_root = "s3:/a"
		owner = "admin"
		delta_sharing_scope = "INTERNAL_AND_EXTERNAL"
		delta_sharing_recipient_token_lifetime_in_seconds = 1002
		`,
	}.ApplyNoError(t)
}

func TestUpdateAccountMetastore_OwnerChanges(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/100/metastores/abc?",
				Response: catalog.AccountsMetastoreInfo{
					MetastoreInfo: &catalog.MetastoreInfo{
						StorageRoot: "s3://b/abc",
						Name:        "a",
					},
				},
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/accounts/100/metastores/abc",
				ExpectedRequest: catalog.AccountsUpdateMetastore{
					MetastoreInfo: &catalog.UpdateMetastore{
						Owner: "updatedOwner",
					},
				},
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/accounts/100/metastores/abc",
				ExpectedRequest: catalog.AccountsUpdateMetastore{
					MetastoreInfo: &catalog.UpdateMetastore{
						Name:              "abc",
						DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
						DeltaSharingRecipientTokenLifetimeInSeconds: 1002,
					},
				},
			},
		},
		Resource:    ResourceMetastore(),
		AccountID:   "100",
		ID:          "abc",
		Update:      true,
		RequiresNew: true,
		InstanceState: map[string]string{
			"name":                "abc",
			"storage_root":        "s3:/a",
			"owner":               "admin",
			"delta_sharing_scope": "INTERNAL_AND_EXTERNAL",
			"delta_sharing_recipient_token_lifetime_in_seconds": "1002",
		},
		HCL: `
		name = "abc"
		storage_root = "s3:/a"
		owner = "updatedOwner"
		delta_sharing_scope = "INTERNAL_AND_EXTERNAL"
		delta_sharing_recipient_token_lifetime_in_seconds = 1002
		`,
	}.ApplyNoError(t)
}

func TestUpdateAccountMetastore_Rollback(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/100/metastores/abc?",
				Response: catalog.AccountsMetastoreInfo{
					MetastoreInfo: &catalog.MetastoreInfo{
						StorageRoot: "s3://b/abc",
						Name:        "a",
					},
				},
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/accounts/100/metastores/abc",
				ExpectedRequest: catalog.AccountsUpdateMetastore{
					MetastoreInfo: &catalog.UpdateMetastore{
						Owner: "updatedOwner",
					},
				},
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/accounts/100/metastores/abc",
				ExpectedRequest: catalog.AccountsUpdateMetastore{
					MetastoreInfo: &catalog.UpdateMetastore{
						Name:              "abc",
						DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
						DeltaSharingRecipientTokenLifetimeInSeconds: 1002,
					},
				},
				Response: apierr.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
			{
				Method:   "PUT",
				Resource: "/api/2.0/accounts/100/metastores/abc",
				ExpectedRequest: catalog.AccountsUpdateMetastore{
					MetastoreInfo: &catalog.UpdateMetastore{
						Owner: "admin",
					},
				},
			},
		},
		Resource:    ResourceMetastore(),
		AccountID:   "100",
		ID:          "abc",
		Update:      true,
		RequiresNew: true,
		InstanceState: map[string]string{
			"name":                "abc",
			"storage_root":        "s3:/a",
			"owner":               "admin",
			"delta_sharing_scope": "INTERNAL_AND_EXTERNAL",
			"delta_sharing_recipient_token_lifetime_in_seconds": "1002",
		},
		HCL: `
		name = "abc"
		storage_root = "s3:/a"
		owner = "updatedOwner"
		delta_sharing_scope = "INTERNAL_AND_EXTERNAL"
		delta_sharing_recipient_token_lifetime_in_seconds = 1002
		`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected happened")
}

func TestUpdateAccountMetastore_DeltaSharingScopeOnly(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PUT",
				Resource: "/api/2.0/accounts/100/metastores/abc",
				ExpectedRequest: catalog.AccountsUpdateMetastore{
					MetastoreInfo: &catalog.UpdateMetastore{
						Name:              "abc",
						DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
						DeltaSharingRecipientTokenLifetimeInSeconds: 1002,
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/100/metastores/abc?",
				Response: catalog.AccountsMetastoreInfo{
					MetastoreInfo: &catalog.MetastoreInfo{
						StorageRoot: "s3://b/abc",
						Name:        "a",
					},
				},
			},
		},
		Resource:    ResourceMetastore(),
		AccountID:   "100",
		ID:          "abc",
		Update:      true,
		RequiresNew: true,
		InstanceState: map[string]string{
			"name":                "abc",
			"storage_root":        "s3:/a",
			"owner":               "admin",
			"delta_sharing_scope": "INTERNAL",
			"delta_sharing_recipient_token_lifetime_in_seconds": "1002",
		},
		HCL: `
		name = "abc"
		storage_root = "s3:/a"
		owner = "admin"
		delta_sharing_scope = "INTERNAL_AND_EXTERNAL"
		delta_sharing_recipient_token_lifetime_in_seconds = 1002
		`,
	}.ApplyNoError(t)
}

func TestReadAccountMetastore(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/100/metastores/abc?",
				Response: catalog.AccountsMetastoreInfo{
					MetastoreInfo: &catalog.MetastoreInfo{
						StorageRoot: "s3://b/abc",
						Name:        "a",
						Region:      "us-east1",
					},
				},
			},
		},
		Resource:  ResourceMetastore(),
		AccountID: "100",
		ID:        "abc",
		Read:      true,
		New:       true,
	}.ApplyAndExpectData(t,
		map[string]any{
			"id":           "abc",
			"storage_root": "s3://b/abc",
			"name":         "a",
			"region":       "us-east1",
		})
}

func TestReadAccountMetastore_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/100/metastores/abc?",
				Response: apierr.APIErrorBody{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   "Metastore with the given ID could not be found.",
				},
				Status: 404,
			},
		},
		Resource:  ResourceMetastore(),
		AccountID: "100",
		ID:        "abc",
		Read:      true,
	}.ExpectError(t, "resource is not expected to be removed")
}

func TestReadMetastore(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/metastores/abc?",
				Response: catalog.MetastoreInfo{
					StorageRoot: "s3://b/abc",
					Name:        "a",
				},
			},
		},
		Resource: ResourceMetastore(),
		ID:       "abc",
		Read:     true,
		New:      true,
	}.ApplyAndExpectData(t,
		map[string]any{
			"id":           "abc",
			"storage_root": "s3://b/abc",
			"name":         "a",
		})
}

func TestReadMetastore_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/metastores/abc?",
				Response: apierr.APIErrorBody{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   "Metastore with the given ID could not be found.",
				},
				Status: 404,
			},
		},
		Resource: ResourceMetastore(),
		ID:       "abc",
		Read:     true,
		New:      true,
	}.ExpectError(t, "resource is not expected to be removed")
}
