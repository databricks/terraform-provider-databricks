package mws

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func testLogDeliveryCreateRequest(status billing.LogDeliveryConfigStatus) billing.WrappedCreateLogDeliveryConfiguration {
	return billing.WrappedCreateLogDeliveryConfiguration{
		LogDeliveryConfiguration: billing.CreateLogDeliveryConfigurationParams{
			ConfigName:             "Audit logs",
			CredentialsId:          "bcd",
			DeliveryPathPrefix:     "/a/b",
			LogType:                billing.LogTypeAuditLogs,
			OutputFormat:           billing.OutputFormatJson,
			StorageConfigurationId: "def",
			DeliveryStartTime:      "2020-10",
			Status:                 status,
			WorkspaceIdsFilter:     []int64{1111111111111111, 222222222222222},
		},
	}
}

func testLogDeliveryResponse(status billing.LogDeliveryConfigStatus) *billing.GetLogDeliveryConfigurationResponse {
	return &billing.GetLogDeliveryConfigurationResponse{
		LogDeliveryConfiguration: &billing.LogDeliveryConfiguration{
			ConfigId:               "nid",
			AccountId:              "abc",
			ConfigName:             "Audit logs",
			CredentialsId:          "bcd",
			DeliveryPathPrefix:     "/a/b",
			LogType:                billing.LogTypeAuditLogs,
			OutputFormat:           billing.OutputFormatJson,
			StorageConfigurationId: "def",
			DeliveryStartTime:      "2020-10",
			Status:                 status,
			WorkspaceIdsFilter:     []int64{1111111111111111, 222222222222222},
		},
	}
}

func testLogDeliveryHCL(status string) string {
	hcl := `
		account_id = "abc"
		credentials_id = "bcd"
		storage_configuration_id = "def"
		config_name = "Audit logs"
		log_type = "AUDIT_LOGS"
		output_format = "JSON"
		delivery_path_prefix = "/a/b"
		workspace_ids_filter = [1111111111111111, 222222222222222]
		delivery_start_time = "2020-10"`
	if status != "" {
		hcl += `
		status = "` + status + `"`
	}
	return hcl
}

func testLogDeliveryInstanceState(status string) map[string]string {
	return map[string]string{
		"account_id":               "abc",
		"config_id":                "nid",
		"config_name":              "Audit logs",
		"credentials_id":           "bcd",
		"delivery_path_prefix":     "/a/b",
		"delivery_start_time":      "2020-10",
		"id":                       "abc|nid",
		"log_type":                 "AUDIT_LOGS",
		"output_format":            "JSON",
		"status":                   status,
		"storage_configuration_id": "def",
	}
}

func TestResourceLogDeliveryCreate(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			api := a.GetMockLogDeliveryAPI().EXPECT()
			api.Create(mock.Anything, testLogDeliveryCreateRequest("")).
				Return(&billing.WrappedLogDeliveryConfiguration{
					LogDeliveryConfiguration: &billing.LogDeliveryConfiguration{ConfigId: "nid"},
				}, nil)
			api.GetByLogDeliveryConfigurationId(mock.Anything, "nid").Return(testLogDeliveryResponse(""), nil)
		},
		Resource: ResourceMwsLogDelivery(),
		HCL:      testLogDeliveryHCL(""),
		Create:   true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":        "abc|nid",
		"config_id": "nid",
	})
}

func TestResourceLogDeliveryCreateDisabled(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			api := a.GetMockLogDeliveryAPI().EXPECT()
			api.Create(mock.Anything, testLogDeliveryCreateRequest(billing.LogDeliveryConfigStatusDisabled)).
				Return(&billing.WrappedLogDeliveryConfiguration{
					LogDeliveryConfiguration: &billing.LogDeliveryConfiguration{ConfigId: "nid"},
				}, nil)
			api.GetByLogDeliveryConfigurationId(mock.Anything, "nid").
				Return(testLogDeliveryResponse(billing.LogDeliveryConfigStatusDisabled), nil)
		},
		Resource: ResourceMwsLogDelivery(),
		HCL:      testLogDeliveryHCL("DISABLED"),
		Create:   true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":        "abc|nid",
		"config_id": "nid",
		"status":    "DISABLED",
	})
}

func TestResourceLogDeliveryCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockLogDeliveryAPI().EXPECT().Create(mock.Anything, mock.Anything).
				Return(nil, &apierr.APIError{
					ErrorCode:  "INVALID_REQUEST",
					Message:    "Internal error happened",
					StatusCode: 400,
				})
		},
		Resource: ResourceMwsLogDelivery(),
		HCL:      testLogDeliveryHCL(""),
		Create:   true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceLogDeliveryRead(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockLogDeliveryAPI().EXPECT().GetByLogDeliveryConfigurationId(mock.Anything, "nid").
				Return(testLogDeliveryResponse(billing.LogDeliveryConfigStatusEnabled), nil)
		},
		Resource: ResourceMwsLogDelivery(),
		Read:     true,
		New:      true,
		ID:       "abc|nid",
	}.ApplyAndExpectData(t, map[string]any{
		"id":                       "abc|nid",
		"credentials_id":           "bcd",
		"storage_configuration_id": "def",
	})
}

func TestResourceLogDeliveryRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockLogDeliveryAPI().EXPECT().GetByLogDeliveryConfigurationId(mock.Anything, "nid").
				Return(&billing.GetLogDeliveryConfigurationResponse{
					LogDeliveryConfiguration: &billing.LogDeliveryConfiguration{
						Status: billing.LogDeliveryConfigStatusDisabled,
					},
				}, nil)
		},
		Resource: ResourceMwsLogDelivery(),
		Read:     true,
		Removed:  true,
		ID:       "abc|nid",
	}.ApplyNoError(t)
}

func TestResourceLogDeliveryRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockLogDeliveryAPI().EXPECT().GetByLogDeliveryConfigurationId(mock.Anything, "nid").
				Return(nil, &apierr.APIError{
					ErrorCode:  "INVALID_REQUEST",
					Message:    "Internal error happened",
					StatusCode: 400,
				})
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
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			api := a.GetMockLogDeliveryAPI().EXPECT()
			api.PatchStatus(mock.Anything, billing.UpdateLogDeliveryConfigurationStatusRequest{
				LogDeliveryConfigurationId: "nid",
				Status:                     billing.LogDeliveryConfigStatusEnabled,
			}).Return(nil)
			api.GetByLogDeliveryConfigurationId(mock.Anything, "nid").
				Return(testLogDeliveryResponse(billing.LogDeliveryConfigStatusEnabled), nil)
		},
		Resource:      ResourceMwsLogDelivery(),
		ID:            "abc|nid",
		Update:        true,
		RequiresNew:   true,
		InstanceState: testLogDeliveryInstanceState("DISABLED"),
		HCL:           testLogDeliveryHCL("ENABLED"),
	}.ApplyAndExpectData(t, map[string]any{
		"id":     "abc|nid",
		"status": "ENABLED",
	})
}

func TestUpdateLogDeliveryError(t *testing.T) {
	_, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockLogDeliveryAPI().EXPECT().PatchStatus(mock.Anything, billing.UpdateLogDeliveryConfigurationStatusRequest{
				LogDeliveryConfigurationId: "nid",
				Status:                     billing.LogDeliveryConfigStatusEnabled,
			}).Return(&apierr.APIError{
				ErrorCode:  "INVALID_REQUEST",
				Message:    "Internal error happened",
				StatusCode: 400,
			})
		},
		Resource:      ResourceMwsLogDelivery(),
		ID:            "abc|nid",
		Update:        true,
		RequiresNew:   true,
		InstanceState: testLogDeliveryInstanceState("DISABLED"),
		HCL:           testLogDeliveryHCL("ENABLED"),
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
}

func TestResourceLogDeliveryDelete(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockLogDeliveryAPI().EXPECT().PatchStatus(mock.Anything, billing.UpdateLogDeliveryConfigurationStatusRequest{
				LogDeliveryConfigurationId: "nid",
				Status:                     billing.LogDeliveryConfigStatusDisabled,
			}).Return(nil)
		},
		Resource: ResourceMwsLogDelivery(),
		Delete:   true,
		ID:       "abc|nid",
	}.ApplyAndExpectData(t, nil)
}

func TestResourceLogDeliveryDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockLogDeliveryAPI().EXPECT().PatchStatus(mock.Anything, billing.UpdateLogDeliveryConfigurationStatusRequest{
				LogDeliveryConfigurationId: "nid",
				Status:                     billing.LogDeliveryConfigStatusDisabled,
			}).Return(&apierr.APIError{
				ErrorCode:  "INVALID_REQUEST",
				Message:    "Internal error happened",
				StatusCode: 400,
			})
		},
		Resource: ResourceMwsLogDelivery(),
		Delete:   true,
		ID:       "abc|nid",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc|nid", d.Id())
}
