package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestCreateMetastore_ApiFieldAccount(t *testing.T) {
	// Even without AccountID set on provider, api = "account" routes to account client
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
					DeltaSharingRecipientTokenLifetimeInSeconds: maxDeltaSharingRecipientTokenLifetimeInSeconds,
				},
			}).Return(&catalog.AccountsUpdateMetastoreResponse{
				MetastoreInfo: &catalog.MetastoreInfo{
					Name: "a",
				},
			}, nil)
			e.GetByMetastoreId(mock.Anything, "abc").Return(&catalog.AccountsGetMetastoreResponse{
				MetastoreInfo: &catalog.MetastoreInfo{
					StorageRoot: "s3://b/abc",
					Name:        "a",
				},
			}, nil)
		},
		Resource: ResourceMetastore(),
		Create:   true,
		HCL: `
		name = "a"
		storage_root = "s3://b"
		api = "account"
		`,
	}.ApplyNoError(t)
}

func TestCreateMetastore_ApiFieldWorkspace(t *testing.T) {
	// api = "workspace" routes to workspace client even when AccountID is set
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
				Id: "abc",
				DeltaSharingRecipientTokenLifetimeInSeconds: maxDeltaSharingRecipientTokenLifetimeInSeconds,
			}).Return(&catalog.MetastoreInfo{
				Name: "a",
			}, nil)
			e.GetById(mock.Anything, "abc").Return(&catalog.MetastoreInfo{
				StorageRoot: "s3://b/abc",
				Name:        "a",
			}, nil)
		},
		Resource:            ResourceMetastore(),
		AccountID:           "100",
		ProviderWorkspaceID: "12345",
		Host:                "https://accounts.cloud.databricks.com",
		Create:              true,
		HCL: `
		name = "a"
		storage_root = "s3://b"
		api = "workspace"
		`,
	}.ApplyNoError(t)
}

func TestReadMetastore_ApiFieldAccount(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockAccountMetastoresAPI().EXPECT().
				GetByMetastoreId(mock.Anything, "abc").
				Return(&catalog.AccountsGetMetastoreResponse{
					MetastoreInfo: &catalog.MetastoreInfo{
						StorageRoot: "s3://b/abc",
						Name:        "a",
					},
				}, nil)
		},
		Resource: ResourceMetastore(),
		Read:     true,
		New:      true,
		ID:       "abc",
		HCL: `
		name = "a"
		storage_root = "s3://b"
		api = "account"
		`,
	}.ApplyNoError(t)
}

func TestDeleteMetastore_ApiFieldAccount(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockAccountMetastoresAPI().EXPECT().
				Delete(mock.Anything, catalog.DeleteAccountMetastoreRequest{
					MetastoreId: "abc",
				}).Return(&catalog.AccountsDeleteMetastoreResponse{}, nil)
		},
		Resource: ResourceMetastore(),
		Delete:   true,
		ID:       "abc",
		HCL: `
		name = "a"
		storage_root = "s3://b"
		api = "account"
		`,
	}.ApplyNoError(t)
}
