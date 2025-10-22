package catalog

import (
	"errors"
	"net/http"
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

func TestUpdateMetastore_OnlyOwnerChanges(t *testing.T) {
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

func TestUpdateMetastore_OwnerAndOtherChanges(t *testing.T) {
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
				DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
				DeltaSharingRecipientTokenLifetimeInSeconds: 1004,
				ForceSendFields: []string{"DeltaSharingRecipientTokenLifetimeInSeconds"},
			}).Return(&catalog.MetastoreInfo{
				Owner:             "updatedOwner",
				DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
				DeltaSharingRecipientTokenLifetimeInSeconds: 1004,
			}, nil)
			e.GetById(mock.Anything, "abc").Return(&catalog.MetastoreInfo{
				Name:              "abc",
				Owner:             "updatedOwner",
				DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
				DeltaSharingRecipientTokenLifetimeInSeconds: 1004,
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
		delta_sharing_recipient_token_lifetime_in_seconds = 1004
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
				DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
				DeltaSharingRecipientTokenLifetimeInSeconds: 1004,
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
		delta_sharing_recipient_token_lifetime_in_seconds = 1004
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
				MetastoreInfo: &catalog.CreateAccountsMetastore{
					StorageRoot: "s3://b",
					Name:        "a",
				},
			}).Return(&catalog.AccountsCreateMetastoreResponse{
				MetastoreInfo: &catalog.MetastoreInfo{
					MetastoreId: "abc",
				},
			}, nil)
			e.GetByMetastoreId(mock.Anything, "abc").Return(&catalog.AccountsGetMetastoreResponse{
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
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockAccountMetastoresAPI().EXPECT()
			e.Create(mock.Anything, catalog.AccountsCreateMetastore{
				MetastoreInfo: &catalog.CreateAccountsMetastore{
					StorageRoot: "s3://b",
					Name:        "a",
				},
			}).Return(&catalog.AccountsCreateMetastoreResponse{
				MetastoreInfo: &catalog.MetastoreInfo{
					MetastoreId: "abc",
				},
			}, nil)
			e.Update(mock.Anything, catalog.AccountsUpdateMetastore{
				MetastoreId: "abc",
				MetastoreInfo: &catalog.UpdateAccountsMetastore{
					Owner: "administrators",
				},
			}).Return(&catalog.AccountsUpdateMetastoreResponse{
				MetastoreInfo: &catalog.MetastoreInfo{
					Name:  "a",
					Owner: "administrators",
				},
			}, nil)
			e.GetByMetastoreId(mock.Anything, "abc").Return(&catalog.AccountsGetMetastoreResponse{
				MetastoreInfo: &catalog.MetastoreInfo{
					StorageRoot: "s3://b/abc",
					Name:        "a",
					Owner:       "administrators",
				},
			}, nil)
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
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockAccountMetastoresAPI().EXPECT()
			e.Create(mock.Anything, catalog.AccountsCreateMetastore{
				MetastoreInfo: &catalog.CreateAccountsMetastore{
					StorageRoot: "s3://b",
					Name:        "a",
				},
			}).Return(&catalog.AccountsCreateMetastoreResponse{
				MetastoreInfo: &catalog.MetastoreInfo{
					MetastoreId: "abc",
				},
			}, nil)
			e.Update(mock.Anything, catalog.AccountsUpdateMetastore{
				MetastoreId: "abc",
				MetastoreInfo: &catalog.UpdateAccountsMetastore{
					Owner:                        "administrators",
					DeltaSharingOrganizationName: "acme",
					DeltaSharingScope:            "INTERNAL_AND_EXTERNAL",
					ForceSendFields:              []string{"DeltaSharingRecipientTokenLifetimeInSeconds"},
				},
			}).Return(&catalog.AccountsUpdateMetastoreResponse{
				MetastoreInfo: &catalog.MetastoreInfo{
					Name:                         "a",
					Owner:                        "administrators",
					DeltaSharingOrganizationName: "acme",
					DeltaSharingScope:            "INTERNAL_AND_EXTERNAL",
				},
			}, nil)
			e.GetByMetastoreId(mock.Anything, "abc").Return(&catalog.AccountsGetMetastoreResponse{
				MetastoreInfo: &catalog.MetastoreInfo{
					StorageRoot: "s3://b/abc",
					Name:        "a",
					Owner:       "administrators",
				},
			}, nil)
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
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockAccountMetastoresAPI().EXPECT().Delete(mock.Anything, catalog.DeleteAccountMetastoreRequest{
				MetastoreId: "abc",
			}).Return(&catalog.AccountsDeleteMetastoreResponse{}, nil)
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
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockAccountMetastoresAPI().EXPECT()
			e.GetByMetastoreId(mock.Anything, "abc").Return(&catalog.AccountsGetMetastoreResponse{
				MetastoreInfo: &catalog.MetastoreInfo{
					StorageRoot: "s3://b/abc",
					Name:        "a",
					Owner:       "administrators",
				},
			}, nil)
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
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockAccountMetastoresAPI().EXPECT()
			e.Update(mock.Anything, catalog.AccountsUpdateMetastore{
				MetastoreId: "abc",
				MetastoreInfo: &catalog.UpdateAccountsMetastore{
					Owner: "updatedOwner",
				},
			}).Return(&catalog.AccountsUpdateMetastoreResponse{
				MetastoreInfo: &catalog.MetastoreInfo{
					Name:  "abc",
					Owner: "updatedOwner",
				},
			}, nil)
			e.GetByMetastoreId(mock.Anything, "abc").Return(&catalog.AccountsGetMetastoreResponse{
				MetastoreInfo: &catalog.MetastoreInfo{
					StorageRoot: "s3://b/abc",
					Name:        "a",
					Owner:       "administrators",
				},
			}, nil)
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
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockAccountMetastoresAPI().EXPECT()
			e.Update(mock.Anything, catalog.AccountsUpdateMetastore{
				MetastoreId: "abc",
				MetastoreInfo: &catalog.UpdateAccountsMetastore{
					Owner: "updatedOwner",
				},
			}).Return(&catalog.AccountsUpdateMetastoreResponse{
				MetastoreInfo: &catalog.MetastoreInfo{
					Name:  "abc",
					Owner: "updatedOwner",
				},
			}, nil)
			e.Update(mock.Anything, catalog.AccountsUpdateMetastore{
				MetastoreId: "abc",
				MetastoreInfo: &catalog.UpdateAccountsMetastore{
					DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
					DeltaSharingRecipientTokenLifetimeInSeconds: 1004,
					ForceSendFields: []string{"DeltaSharingRecipientTokenLifetimeInSeconds"},
				},
			}).Return(nil, errors.New("Something unexpected happened"))
			e.Update(mock.Anything, catalog.AccountsUpdateMetastore{
				MetastoreId: "abc",
				MetastoreInfo: &catalog.UpdateAccountsMetastore{
					Owner: "admin",
				},
			}).Return(&catalog.AccountsUpdateMetastoreResponse{
				MetastoreInfo: &catalog.MetastoreInfo{
					Name:  "abc",
					Owner: "admin",
				},
			}, nil)
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
		delta_sharing_recipient_token_lifetime_in_seconds = 1004
		`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected happened")
}

func TestUpdateAccountMetastore_DeltaSharingScopeOnly(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			e := a.GetMockAccountMetastoresAPI().EXPECT()
			e.Update(mock.Anything, catalog.AccountsUpdateMetastore{
				MetastoreId: "abc",
				MetastoreInfo: &catalog.UpdateAccountsMetastore{
					DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
					DeltaSharingRecipientTokenLifetimeInSeconds: 1002,
					ForceSendFields: []string{"DeltaSharingRecipientTokenLifetimeInSeconds"},
				},
			}).Return(&catalog.AccountsUpdateMetastoreResponse{
				MetastoreInfo: &catalog.MetastoreInfo{
					Name:              "abc",
					DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
					DeltaSharingRecipientTokenLifetimeInSeconds: 1002,
				},
			}, nil)
			e.GetByMetastoreId(mock.Anything, "abc").Return(&catalog.AccountsGetMetastoreResponse{
				MetastoreInfo: &catalog.MetastoreInfo{
					StorageRoot:       "s3://b/abc",
					Name:              "abc",
					DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
					DeltaSharingRecipientTokenLifetimeInSeconds: 1002,
				},
			}, nil)

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
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockAccountMetastoresAPI().EXPECT().GetByMetastoreId(mock.Anything, "abc").Return(&catalog.AccountsGetMetastoreResponse{
				MetastoreInfo: &catalog.MetastoreInfo{
					StorageRoot: "s3://b/abc",
					Name:        "a",
					Region:      "us-east1",
				},
			}, nil)
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
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockAccountMetastoresAPI().EXPECT().GetByMetastoreId(mock.Anything, "abc").Return(nil, &apierr.APIError{
				StatusCode: http.StatusNotFound,
			})
		},
		Resource:  ResourceMetastore(),
		AccountID: "100",
		ID:        "abc",
		Read:      true,
	}.ExpectError(t, "resource is not expected to be removed")
}

func TestReadMetastore(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockMetastoresAPI().EXPECT()
			e.GetById(mock.Anything, "abc").Return(&catalog.MetastoreInfo{
				StorageRoot: "s3://b/abc",
				Name:        "a",
			}, nil)
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockMetastoresAPI().EXPECT()
			e.GetById(mock.Anything, "abc").Return(nil, &apierr.APIError{
				StatusCode: http.StatusNotFound,
			})
		},
		Resource: ResourceMetastore(),
		ID:       "abc",
		Read:     true,
		New:      true,
	}.ExpectError(t, "resource is not expected to be removed")
}
