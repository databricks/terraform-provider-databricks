package sql

import (
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

var (
	alertResponse = sql.Alert{
		Id:            "7890",
		QueryId:       "123456",
		DisplayName:   "TF new alert",
		OwnerUserName: "user@domain.com",
		Condition: &sql.AlertCondition{
			Op: "GREATER_THAN",
			Operand: &sql.AlertConditionOperand{
				Column: &sql.AlertOperandColumn{
					Name: "value",
				},
			},
			Threshold: &sql.AlertConditionThreshold{
				Value: &sql.AlertOperandValue{
					DoubleValue: 42,
				},
			},
		},
		ParentPath: "/Workspace/Shared/Alerts",
		NotifyOnOk: true,
	}
	createHcl = `query_id = "123456"
  display_name = "TF new alert"
  parent_path = "/Shared/Alerts"
  owner_user_name = "user@domain.com"
  condition {
    op = "GREATER_THAN"
    operand {
      column {
        name = "value"
      }
    }
    threshold {
      value {
        double_value = 42
      }
    }
  }`
	createAlertRequest = sql.CreateAlertRequest{
		AutoResolveDisplayName: false,
		Alert: &sql.CreateAlertRequestAlert{
			QueryId:     "123456",
			DisplayName: "TF new alert",
			ParentPath:  "/Shared/Alerts",
			NotifyOnOk:  true,
			Condition: &sql.AlertCondition{
				Op: "GREATER_THAN",
				Operand: &sql.AlertConditionOperand{
					Column: &sql.AlertOperandColumn{
						Name: "value",
					},
				},
				Threshold: &sql.AlertConditionThreshold{
					Value: &sql.AlertOperandValue{
						DoubleValue: 42,
					},
				},
			},
		},
		ForceSendFields: []string{"AutoResolveDisplayName"},
	}
)

func TestAlertCreate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockAlertsAPI().EXPECT()
			e.Create(mock.Anything, createAlertRequest).Return(&alertResponse, nil)
			e.Update(mock.Anything, sql.UpdateAlertRequest{
				Id:         "7890",
				UpdateMask: "owner_user_name",
				Alert: &sql.UpdateAlertRequestAlert{
					OwnerUserName: "user@domain.com",
				},
			}).Return(&alertResponse, nil)
			e.GetById(mock.Anything, "7890").Return(&alertResponse, nil)
		},
		Resource: ResourceAlert(),
		Create:   true,
		HCL:      createHcl,
	}.ApplyAndExpectData(t, map[string]any{
		"id":              "7890",
		"query_id":        "123456",
		"display_name":    "TF new alert",
		"owner_user_name": "user@domain.com",
	})
}

func TestAlertCreate_BackendError(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockAlertsAPI().EXPECT()
			e.Create(mock.Anything, createAlertRequest).Return(nil, &apierr.APIError{
				StatusCode: http.StatusBadRequest,
				Message:    "Node named 'TF new alert' already exists",
			})
		},
		Resource: ResourceAlert(),
		Create:   true,
		HCL:      createHcl,
	}.ExpectError(t, "Node named 'TF new alert' already exists")
}

func TestAlertCreate_ErrorMultipleValues(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceAlert(),
		Create:   true,
		HCL: `query_id = "123456"
  display_name = "TF new alert"
  parent_path = "/Shared/Alerts"
  owner_user_name = "user@domain.com"
  condition {
    op = "GREATER_THAN"
    operand {
      column {
        name = "value"
      }
    }
    threshold {
      value {
        double_value = 42
      }
    }
    threshold {
      value {
        bool_value = 42
      }
    }
}`,
	}.ExpectError(t, "invalid config supplied. [condition.#.threshold] Too many list items")
}

func TestAlertRead_Import(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockAlertsAPI().EXPECT().GetById(mock.Anything, "7890").Return(&alertResponse, nil)
		},
		Resource: ResourceAlert(),
		Read:     true,
		ID:       "7890",
		New:      true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":              "7890",
		"query_id":        "123456",
		"display_name":    "TF new alert",
		"owner_user_name": "user@domain.com",
	})
}

func TestAlertRead_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockAlertsAPI().EXPECT().GetById(mock.Anything, "7890").Return(nil, &apierr.APIError{
				StatusCode: http.StatusBadRequest,
				Message:    "bad payload",
			})
		},
		Resource: ResourceAlert(),
		Read:     true,
		ID:       "7890",
		New:      true,
	}.ExpectError(t, "bad payload")
}

func TestAlertDelete(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockAlertsAPI().EXPECT().DeleteById(mock.Anything, "7890").Return(nil)
		},
		Resource: ResourceAlert(),
		Delete:   true,
		ID:       "7890",
		New:      true,
	}.ApplyNoError(t)
}

func TestAlertUpdate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockAlertsAPI().EXPECT()
			e.Update(mock.Anything, sql.UpdateAlertRequest{
				Id:         "7890",
				UpdateMask: "display_name,query_id,seconds_to_retrigger,condition,custom_body,custom_subject,owner_user_name,notify_on_ok",
				Alert: &sql.UpdateAlertRequestAlert{
					QueryId:       "123456",
					DisplayName:   "TF new alert",
					OwnerUserName: "user@domain.com",
					NotifyOnOk:    false,
					Condition: &sql.AlertCondition{
						Op: "GREATER_THAN",
						Operand: &sql.AlertConditionOperand{
							Column: &sql.AlertOperandColumn{
								Name: "value",
							},
						},
						Threshold: &sql.AlertConditionThreshold{
							Value: &sql.AlertOperandValue{
								DoubleValue: 42,
							},
						},
					},
					ForceSendFields: []string{"NotifyOnOk"},
				}}).Return(&alertResponse, nil)
			e.GetById(mock.Anything, "7890").Return(&alertResponse, nil)
		},
		Resource: ResourceAlert(),
		Update:   true,
		ID:       "7890",
		New:      true,
		InstanceState: map[string]string{
			"id":           "7890",
			"query_id":     "123456",
			"notify_on_ok": "true",
		},
		HCL: `query_id = "123456"
  display_name = "TF new alert"
  owner_user_name = "user@domain.com"
  notify_on_ok = false
  condition {
    op = "GREATER_THAN"
    operand {
      column {
        name = "value"
      }
    }
    threshold {
      value {
        double_value = 42
      }
    }
  }`,
	}.ApplyAndExpectData(t, map[string]any{
		"id":              "7890",
		"query_id":        "123456",
		"display_name":    "TF new alert",
		"owner_user_name": "user@domain.com",
	})
}
