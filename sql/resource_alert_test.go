package sql

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestAlertCreate(t *testing.T) {
	alertResponse := sql.Alert{
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
	}
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockAlertsAPI().EXPECT()
			e.Create(mock.Anything, sql.CreateAlertRequest{
				Alert: &sql.CreateAlertRequestAlert{
					QueryId:     "123456",
					DisplayName: "TF new alert",
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
				}}).Return(&alertResponse, nil)
			e.GetById(mock.Anything, "7890").Return(&alertResponse, nil)
		},
		Resource: ResourceAlert(),
		Create:   true,
		HCL: `query_id = "123456"
  display_name = "TF new alert"
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
