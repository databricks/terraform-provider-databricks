package mws

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
)

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
						WorkspaceIdsFilter:     []int64{1111111111111111, 222222222222222},
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
						WorkspaceIdsFilter:     []int64{1111111111111111, 222222222222222},
					},
				},
			},
		},
		Resource: ResourceMwsLogDelivery(),
		HCL: `
		account_id = "abc"
		credentials_id = "bcd"
		storage_configuration_id = "def"
		config_name = "Audit logs"
		log_type = "AUDIT_LOGS"
		output_format = "JSON"
		delivery_path_prefix = "/a/b"
		workspace_ids_filter = [1111111111111111, 222222222222222]
		delivery_start_time = "2020-10"`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc|nid", d.Id())
	assert.Equal(t, "nid", d.Get("config_id"))
}

func TestResourceLogDeliveryCreateDisabled(t *testing.T) {
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
						Status:                 "DISABLED",
						WorkspaceIdsFilter:     []int64{1111111111111111, 222222222222222},
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
						Status:                 "DISABLED",
						WorkspaceIdsFilter:     []int64{1111111111111111, 222222222222222},
					},
				},
			},
		},
		Resource: ResourceMwsLogDelivery(),
		HCL: `
		account_id = "abc"
		credentials_id = "bcd"
		storage_configuration_id = "def"
		config_name = "Audit logs"
		log_type = "AUDIT_LOGS"
		output_format = "JSON"
		delivery_path_prefix = "/a/b"
		workspace_ids_filter = [1111111111111111, 222222222222222]
		delivery_start_time = "2020-10"
		status = "DISABLED"`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
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
		Resource: ResourceMwsLogDelivery(),
		HCL: `
		account_id = "abc"
		credentials_id = "bcd"
		storage_configuration_id = "def"
		config_name = "Audit logs"
		log_type = "AUDIT_LOGS"
		output_format = "JSON"
		delivery_path_prefix = "/a/b"
		workspace_ids_filter = [1111111111111111, 222222222222222]
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
						WorkspaceIdsFilter:     []int64{1111111111111111, 222222222222222},
					},
				},
			},
		},
		Resource: ResourceMwsLogDelivery(),
		Read:     true,
		New:      true,
		ID:       "abc|nid",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc|nid", d.Id(), "Id should not be empty")
	assert.Equal(t, "bcd", d.Get("credentials_id"))
	assert.Equal(t, "def", d.Get("storage_configuration_id"))
}

func TestResourceLogDeliveryRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
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
		Resource: ResourceMwsLogDelivery(),
		Read:     true,
		Removed:  true,
		ID:       "abc|nid",
	}.ApplyNoError(t)
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
		Resource: ResourceMwsLogDelivery(),
		Read:     true,
		ID:       "abc|nid",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc|nid", d.Id(), "Id should not be empty for error reads")
}

func TestUpdateLogDelivery(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/accounts/abc/log-delivery/nid",
				ExpectedRequest: map[string]any{
					"status": "ENABLED",
				},
			},
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
					},
				},
			},
		},
		Resource:    ResourceMwsLogDelivery(),
		ID:          "abc|nid",
		Update:      true,
		RequiresNew: true,
		InstanceState: map[string]string{
			"account_id":               "abc",
			"config_id":                "nid",
			"config_name":              "Audit logs",
			"credentials_id":           "bcd",
			"delivery_path_prefix":     "/a/b",
			"delivery_start_time":      "2020-10",
			"id":                       "abc|nid",
			"log_type":                 "AUDIT_LOGS",
			"output_format":            "JSON",
			"status":                   "DISABLED",
			"storage_configuration_id": "def",
		},
		HCL: `
		account_id = "abc"
		credentials_id = "bcd"
		storage_configuration_id = "def"
		config_name = "Audit logs"
		log_type = "AUDIT_LOGS"
		output_format = "JSON"
		delivery_path_prefix = "/a/b"
		delivery_start_time = "2020-10"
		status = "ENABLED"
		`,
	}.ApplyNoError(t)
}

func TestUpdateLogDeliveryError(t *testing.T) {
	_, err := qa.ResourceFixture{
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
		Resource:    ResourceMwsLogDelivery(),
		ID:          "abc|nid",
		Update:      true,
		RequiresNew: true,
		InstanceState: map[string]string{
			"account_id":               "abc",
			"config_id":                "nid",
			"config_name":              "Audit logs",
			"credentials_id":           "bcd",
			"delivery_path_prefix":     "/a/b",
			"delivery_start_time":      "2020-10",
			"id":                       "abc|nid",
			"log_type":                 "AUDIT_LOGS",
			"output_format":            "JSON",
			"status":                   "DISABLED",
			"storage_configuration_id": "def",
		},
		HCL: `
		account_id = "abc"
		credentials_id = "bcd"
		storage_configuration_id = "def"
		config_name = "Audit logs"
		log_type = "AUDIT_LOGS"
		output_format = "JSON"
		delivery_path_prefix = "/a/b"
		delivery_start_time = "2020-10"
		status = "ENABLED"
		`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
}

func TestResourceLogDeliveryDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/accounts/abc/log-delivery/nid",
				ExpectedRequest: map[string]string{
					"status": "DISABLED",
				},
			},
		},
		Resource: ResourceMwsLogDelivery(),
		Delete:   true,
		ID:       "abc|nid",
	}.ApplyNoError(t)
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
		Resource: ResourceMwsLogDelivery(),
		Delete:   true,
		ID:       "abc|nid",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc|nid", d.Id())
}
