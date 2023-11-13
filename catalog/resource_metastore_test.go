package catalog

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestMetastoreCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceMetastore())
}

func TestCreateMetastore(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/metastores",
				ExpectedRequest: catalog.CreateMetastore{
					StorageRoot: "s3://b",
					Name:        "a",
				},
				Response: MetastoreInfo{
					MetastoreID: "abc",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/metastores/abc",
				ExpectedRequest: catalog.UpdateMetastore{
					Name: "a",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/metastores/abc?",
				Response: MetastoreInfo{
					StorageRoot: "s3://b/abc",
					Name:        "a",
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/metastores",
				ExpectedRequest: catalog.CreateMetastore{
					StorageRoot: "s3://b",
					Name:        "a",
				},
				Response: MetastoreInfo{
					MetastoreID: "abc",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/metastores/abc",
				ExpectedRequest: catalog.UpdateMetastore{
					Name:  "a",
					Owner: "administrators",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/metastores/abc?",
				Response: MetastoreInfo{
					StorageRoot: "s3://b/abc",
					Name:        "a",
					Owner:       "administrators",
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/metastores",
				ExpectedRequest: catalog.CreateMetastore{
					StorageRoot: "s3://b",
					Name:        "a",
				},
				Response: MetastoreInfo{
					MetastoreID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/metastores/abc?",
				Response: MetastoreInfo{
					StorageRoot: "s3://b/abc",
					Name:        "a",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/metastores/abc",
				ExpectedRequest: catalog.UpdateMetastore{
					Name:              "a",
					Owner:             "administrators",
					DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
					DeltaSharingRecipientTokenLifetimeInSeconds: 0,
					DeltaSharingOrganizationName:                "acme",
					ForceSendFields:                             []string{"DeltaSharingRecipientTokenLifetimeInSeconds"},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/metastores/abc?",
			},
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/metastores/abc?force=true",
			},
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/metastores/abc?",
				Response: MetastoreInfo{
					StorageRoot: "s3://b/abc",
					Name:        "a",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/metastores/abc",
				ExpectedRequest: catalog.UpdateMetastore{
					Name:              "abc",
					Owner:             "admin",
					DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
					DeltaSharingRecipientTokenLifetimeInSeconds: 1002,
				},
			},
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

func TestUpdateMetastore_DeltaSharingScopeOnly(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/metastores/abc",
				ExpectedRequest: catalog.UpdateMetastore{
					Name:              "abc",
					Owner:             "admin",
					DeltaSharingScope: "INTERNAL_AND_EXTERNAL",
					DeltaSharingRecipientTokenLifetimeInSeconds: 1002,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/metastores/abc?",
				Response: MetastoreInfo{
					StorageRoot: "s3://b/abc",
					Name:        "a",
				},
			},
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
						Name: "a",
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
						Owner:             "admin",
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

func TestUpdateAccountMetastore_DeltaSharingScopeOnly(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PUT",
				Resource: "/api/2.0/accounts/100/metastores/abc",
				ExpectedRequest: catalog.AccountsUpdateMetastore{
					MetastoreInfo: &catalog.UpdateMetastore{
						Name:              "abc",
						Owner:             "admin",
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
