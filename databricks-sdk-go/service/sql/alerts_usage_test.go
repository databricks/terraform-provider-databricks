// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package sql_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/sql"
)

func ExampleAlertsAPI_Create_alerts() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	srcs, err := w.DataSources.List(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", srcs)

	query, err := w.Queries.Create(ctx, sql.CreateQueryRequest{
		Query: &sql.CreateQueryRequestQuery{
			DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			WarehouseId: srcs[0].WarehouseId,
			Description: "test query from Go SDK",
			QueryText:   "SELECT 1",
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", query)

	alert, err := w.Alerts.Create(ctx, sql.CreateAlertRequest{
		Alert: &sql.CreateAlertRequestAlert{
			Condition: &sql.AlertCondition{
				Operand: &sql.AlertConditionOperand{
					Column: &sql.AlertOperandColumn{
						Name: "1",
					},
				},
				Op: sql.AlertOperatorEqual,
				Threshold: &sql.AlertConditionThreshold{
					Value: &sql.AlertOperandValue{
						DoubleValue: 1,
					},
				},
			},
			DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			QueryId:     query.Id,
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", alert)

	// cleanup

	err = w.Queries.DeleteById(ctx, query.Id)
	if err != nil {
		panic(err)
	}
	err = w.Alerts.DeleteById(ctx, alert.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleAlertsAPI_Get_alerts() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	srcs, err := w.DataSources.List(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", srcs)

	query, err := w.Queries.Create(ctx, sql.CreateQueryRequest{
		Query: &sql.CreateQueryRequestQuery{
			DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			WarehouseId: srcs[0].WarehouseId,
			Description: "test query from Go SDK",
			QueryText:   "SELECT 1",
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", query)

	alert, err := w.Alerts.Create(ctx, sql.CreateAlertRequest{
		Alert: &sql.CreateAlertRequestAlert{
			Condition: &sql.AlertCondition{
				Operand: &sql.AlertConditionOperand{
					Column: &sql.AlertOperandColumn{
						Name: "1",
					},
				},
				Op: sql.AlertOperatorEqual,
				Threshold: &sql.AlertConditionThreshold{
					Value: &sql.AlertOperandValue{
						DoubleValue: 1,
					},
				},
			},
			DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			QueryId:     query.Id,
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", alert)

	byId, err := w.Alerts.GetById(ctx, alert.Id)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	err = w.Queries.DeleteById(ctx, query.Id)
	if err != nil {
		panic(err)
	}
	err = w.Alerts.DeleteById(ctx, alert.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleAlertsAPI_List_alerts() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.Alerts.ListAll(ctx, sql.ListAlertsRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleAlertsAPI_Update_alerts() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	srcs, err := w.DataSources.List(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", srcs)

	query, err := w.Queries.Create(ctx, sql.CreateQueryRequest{
		Query: &sql.CreateQueryRequestQuery{
			DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			WarehouseId: srcs[0].WarehouseId,
			Description: "test query from Go SDK",
			QueryText:   "SELECT 1",
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", query)

	alert, err := w.Alerts.Create(ctx, sql.CreateAlertRequest{
		Alert: &sql.CreateAlertRequestAlert{
			Condition: &sql.AlertCondition{
				Operand: &sql.AlertConditionOperand{
					Column: &sql.AlertOperandColumn{
						Name: "1",
					},
				},
				Op: sql.AlertOperatorEqual,
				Threshold: &sql.AlertConditionThreshold{
					Value: &sql.AlertOperandValue{
						DoubleValue: 1,
					},
				},
			},
			DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			QueryId:     query.Id,
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", alert)

	_, err = w.Alerts.Update(ctx, sql.UpdateAlertRequest{
		Id: alert.Id,
		Alert: &sql.UpdateAlertRequestAlert{
			DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		},
		UpdateMask: "display_name",
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Queries.DeleteById(ctx, query.Id)
	if err != nil {
		panic(err)
	}
	err = w.Alerts.DeleteById(ctx, alert.Id)
	if err != nil {
		panic(err)
	}

}
