package mws

import (
	"context"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

func TestMwsAccStorageConfigurations(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	acctID := qa.GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	storageAPI := NewStorageConfigurationsAPI(context.Background(), common.CommonEnvironmentClient())
	storageConfigsList, err := storageAPI.List(acctID)
	assert.NoError(t, err, err)
	t.Log(storageConfigsList)

	storageConfig, err := storageAPI.Create(acctID, "sri-mws-terraform-storage-root-bucket", "sri-root-s3-bucket")
	assert.NoError(t, err, err)

	myStorageConfig, err := storageAPI.Read(acctID, storageConfig.StorageConfigurationID)
	assert.NoError(t, err, err)
	t.Log(myStorageConfig.RootBucketInfo.BucketName)

	defer func() {
		err = storageAPI.Delete(acctID, storageConfig.StorageConfigurationID)
		assert.NoError(t, err, err)
	}()
}

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
		Resource: ResourceStorageConfiguration(),
		State: map[string]interface{}{
			"account_id":                 "abc",
			"bucket_name":                "bucket",
			"storage_configuration_name": "Main Storage",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
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
		Resource: ResourceStorageConfiguration(),
		State: map[string]interface{}{
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
		Resource: ResourceStorageConfiguration(),
		Read:     true,
		ID:       "abc/scid",
	}.Apply(t)
	assert.NoError(t, err, err)
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
		Resource: ResourceStorageConfiguration(),
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
		Resource: ResourceStorageConfiguration(),
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
		Resource: ResourceStorageConfiguration(),
		Delete:   true,
		ID:       "abc/scid",
	}.Apply(t)
	assert.NoError(t, err, err)
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
		Resource: ResourceStorageConfiguration(),
		Delete:   true,
		ID:       "abc/scid",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/scid", d.Id())
}
