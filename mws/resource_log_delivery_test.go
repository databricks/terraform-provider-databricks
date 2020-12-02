package mws

import (
	"context"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMwsAccLogDelivery(t *testing.T) {
	acctID := qa.GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	roleARN := qa.GetEnvOrSkipTest(t, "TEST_LOGDELIVERY_ARN")
	bucket := qa.GetEnvOrSkipTest(t, "TEST_LOGDELIVERY_BUCKET")
	client := common.CommonEnvironmentClient()
	randomName := qa.RandomName("tf-logdelivery-")

	ctx := context.Background()
	logDeliveryAPI := NewLogDeliveryAPI(ctx, client)
	credentialsAPI := NewCredentialsAPI(ctx, client)
	storageConfigurationsAPI := NewStorageConfigurationsAPI(ctx, client)

	creds, err := credentialsAPI.Create(acctID, randomName, roleARN)
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, credentialsAPI.Delete(acctID, creds.CredentialsID))
	}()

	storageConfig, err := storageConfigurationsAPI.Create(acctID, randomName, bucket)
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, storageConfigurationsAPI.Delete(acctID, storageConfig.StorageConfigurationID))
	}()

	configID, err := logDeliveryAPI.Create(LogDeliveryConfiguration{
		AccountID:              acctID,
		CredentialsID:          creds.CredentialsID,
		StorageConfigurationID: storageConfig.StorageConfigurationID,
		ConfigName:             randomName,
		DeliveryPathPrefix:     randomName,
		LogType:                "AUDIT_LOGS",
		OutputFormat:           "JSON",
	})
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, logDeliveryAPI.Disable(acctID, configID))
	}()

	ldc, err := logDeliveryAPI.Read(acctID, configID)
	require.NoError(t, err)
	assert.Equal(t, "ENABLED", ldc.Status)
}

func TestResourceLogDeliveryCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/log-delivery",
				ExpectedRequest: LogDelivery{
					LogDeliveryConfiguration: LogDeliveryConfiguration{
						AccountID:              "abc",
						ConfigName:             "Audit logs",
						CredentialsID:          "bcd",
						DeliveryPathPrefix:     "/a/b",
						LogType:                "AUDIT_LOGS",
						OutputFormat:           "JSON",
						StorageConfigurationID: "def",
						DeliveryStartTime:      "2020-10",
						WorkspaceIdsFilter:     []string{"e", "f"},
					},
				},
				Response: LogDelivery{
					LogDeliveryConfiguration: LogDeliveryConfiguration{
						ConfigID: "nid",
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/log-delivery/nid",
				Response: LogDelivery{
					LogDeliveryConfiguration: LogDeliveryConfiguration{
						ConfigID:               "nid",
						AccountID:              "abc",
						ConfigName:             "Audit logs",
						CredentialsID:          "bcd",
						DeliveryPathPrefix:     "/a/b",
						LogType:                "AUDIT_LOGS",
						OutputFormat:           "JSON",
						StorageConfigurationID: "def",
						DeliveryStartTime:      "2020-10",
						WorkspaceIdsFilter:     []string{"e", "f"},
					},
				},
			},
		},
		Resource: ResourceLogDelivery(),
		HCL: `
		account_id = "abc"
		credentials_id = "bcd"
		storage_configuration_id = "def"
		config_name = "Audit logs"
		log_type = "AUDIT_LOGS"
		output_format = "JSON"
		delivery_path_prefix = "/a/b"
		workspace_ids_filter = ["e", "f"]
		delivery_start_time = "2020-10"`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc|nid", d.Id())
	assert.Equal(t, "nid", d.Get("config_id"))
}

func TestResourceLogDeliveryCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/log-delivery",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceLogDelivery(),
		HCL: `
		account_id = "abc"
		credentials_id = "bcd"
		storage_configuration_id = "def"
		config_name = "Audit logs"
		log_type = "AUDIT_LOGS"
		output_format = "JSON"
		delivery_path_prefix = "/a/b"
		workspace_ids_filter = ["e", "f"]
		delivery_start_time = "2020-10"`,
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceLogDeliveryRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/log-delivery/nid",
				Response: LogDelivery{
					LogDeliveryConfiguration: LogDeliveryConfiguration{
						ConfigID:               "nid",
						Status:                 "ENABLED",
						AccountID:              "abc",
						ConfigName:             "Audit logs",
						CredentialsID:          "bcd",
						DeliveryPathPrefix:     "/a/b",
						LogType:                "AUDIT_LOGS",
						OutputFormat:           "JSON",
						StorageConfigurationID: "def",
						DeliveryStartTime:      "2020-10",
						WorkspaceIdsFilter:     []string{"e", "f"},
					},
				},
			},
		},
		Resource: ResourceLogDelivery(),
		Read:     true,
		New:      true,
		ID:       "abc|nid",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc|nid", d.Id(), "Id should not be empty")
	assert.Equal(t, "bcd", d.Get("credentials_id"))
	assert.Equal(t, "def", d.Get("storage_configuration_id"))
}

func TestResourceLogDeliveryRead_NotFound(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/log-delivery/nid",
				Response: LogDelivery{
					LogDeliveryConfiguration: LogDeliveryConfiguration{
						Status: "DISABLED",
					},
				},
			},
		},
		Resource: ResourceLogDelivery(),
		Read:     true,
		ID:       "abc|nid",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceLogDeliveryRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/log-delivery/nid",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceLogDelivery(),
		Read:     true,
		ID:       "abc|nid",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc|nid", d.Id(), "Id should not be empty for error reads")
}

func TestResourceLogDeliveryDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/accounts/abc/log-delivery/nid",
				ExpectedRequest: map[string]string{
					"status": "DISABLED",
				},
			},
		},
		Resource: ResourceLogDelivery(),
		Delete:   true,
		ID:       "abc|nid",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc|nid", d.Id())
}

func TestResourceLogDeliveryDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/accounts/abc/log-delivery/nid",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceLogDelivery(),
		Delete:   true,
		ID:       "abc|nid",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc|nid", d.Id())
}
