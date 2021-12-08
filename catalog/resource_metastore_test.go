package catalog

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
)

func TestMetastoreCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceMetastore())
}

func TestCreateMetastore(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/unity-catalog/metastores",
				ExpectedRequest: MetastoreInfo{
					StorageRoot: "s3://b",
					Name:        "a",
				},
				Response: MetastoreInfo{
					MetastoreID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/metastores/abc",
				Response: MetastoreInfo{
					StorageRoot: "s3://b/abc",
					Name:        "a",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.0/unity-catalog/metastores/abc",
				ExpectedRequest: map[string]interface{}{
					"owner": "administrators",
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

func TestDeleteMetastore(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/unity-catalog/metastores/abc",
				ExpectedRequest: map[string]bool{
					"force": false,
				},
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
				Resource: "/api/2.0/unity-catalog/metastores/abc",
				ExpectedRequest: map[string]bool{
					"force": true,
				},
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
