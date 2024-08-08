package mws

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceStorageConfigurationCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/storage-configurations",
				ExpectedRequest: StorageConfiguration{
					StorageConfigurationName: "Main Storage",
					RootBucketInfo: &RootBucketInfo{
						BucketName: "bucket",
					},
				},
				Response: StorageConfiguration{
					StorageConfigurationID: "scid",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/storage-configurations/scid",
				Response: StorageConfiguration{
					StorageConfigurationID:   "scid",
					StorageConfigurationName: "Main Storage",
					RootBucketInfo: &RootBucketInfo{
						BucketName: "bucket",
					},
				},
			},
		},
		Resource: ResourceMwsStorageConfigurations(),
		State: map[string]any{
			"account_id":                 "abc",
			"bucket_name":                "bucket",
			"storage_configuration_name": "Main Storage",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/scid", d.Id())
}

func TestResourceStorageConfigurationCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/storage-configurations",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceMwsStorageConfigurations(),
		State: map[string]any{
			"account_id":                 "abc",
			"bucket_name":                "bucket",
			"storage_configuration_name": "Main Storage",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceStorageConfigurationRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/storage-configurations/scid",
				Response: StorageConfiguration{
					StorageConfigurationID:   "scid",
					StorageConfigurationName: "Main Storage",
					RootBucketInfo: &RootBucketInfo{
						BucketName: "bucket",
					},
				},
			},
		},
		Resource: ResourceMwsStorageConfigurations(),
		Read:     true,
		ID:       "abc/scid",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/scid", d.Id(), "Id should not be empty")
	assert.Equal(t, "bucket", d.Get("bucket_name"))
	assert.Equal(t, 0, d.Get("creation_time"))
	assert.Equal(t, "scid", d.Get("storage_configuration_id"))
	assert.Equal(t, "Main Storage", d.Get("storage_configuration_name"))
}

func TestResourceStorageConfigurationRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/storage-configurations/scid",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceMwsStorageConfigurations(),
		Read:     true,
		Removed:  true,
		ID:       "abc/scid",
	}.ApplyNoError(t)
}

func TestResourceStorageConfigurationRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/storage-configurations/scid",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceMwsStorageConfigurations(),
		Read:     true,
		ID:       "abc/scid",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/scid", d.Id(), "Id should not be empty for error reads")
}

func TestResourceStorageConfigurationDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for better stub url...
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/abc/storage-configurations/scid",
			},
		},
		Resource: ResourceMwsStorageConfigurations(),
		Delete:   true,
		ID:       "abc/scid",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/scid", d.Id())
}

func TestResourceStorageConfigurationDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/abc/storage-configurations/scid",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceMwsStorageConfigurations(),
		Delete:   true,
		ID:       "abc/scid",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/scid", d.Id())
}
