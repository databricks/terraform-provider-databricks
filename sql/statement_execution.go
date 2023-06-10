package sql

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Statement is the struct that contains what the 2.0 api returns for the statement api
type Statement struct {
	StatementID string  `json:"statement_id,omitempty"`
	Status      *Status `json:"status,omitempty"`
}

// Status captures the state of a Statement
type Status struct {
	State string `json:"state,omitempty"`
}

// NewStatementAPI creates StatementAPI instance from provider meta
func NewStatementAPI(ctx context.Context, m any) StatementAPI {
	return StatementAPI{
		client:  m.(*common.DatabricksClient),
		context: context.WithValue(ctx, common.Api, common.API_2_0),
	}
}

// StatementAPI exposes the SQL Statement Execution API
type StatementAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Execute runs a SQL statement on a SQL Warehouse/Endpoint and waits for it to finish.
// If the SQL warehouse is not currently running, it will be started.
// Returns the completed statement ID.
func (a StatementAPI) Execute(warehouseID, statementStr string) error {
	// this is the place, where API version propagation through context looks strange
	ctx := context.WithValue(a.context, common.Api, common.API_2_0)
	sqlEndpointsAPI := NewSQLEndpointsAPI(ctx, a.client)
	warehouse, err := sqlEndpointsAPI.Get(warehouseID)
	if err != nil {
		return err
	}

	// start the sql warehouse if it is not running
	if warehouse.State != "RUNNING" {
		err := sqlEndpointsAPI.Start(warehouseID, *schema.DefaultTimeout(30 * time.Minute))
		if err != nil {
			return err
		}
	}
	log.Printf("[INFO] Executing SQL statement on %s:\n%s", warehouseID, statementStr)
	result, err := a.executeStatement(warehouseID, statementStr)
	if err != nil {
		return err
	}

	err = a.waitForStatementFinished(result.StatementID)
	if err != nil {
		return err
	}
	return nil
}

type genericStatementRequest struct {
	Format        string `json:"format,omitempty" url:"format,omitempty"`
	Catalog       string `json:"catalog,omitempty" url:"catalog,omitempty"`
	WaitTimeout   string `json:"wait_timeout,omitempty" url:"wait_timeout,omitempty"`
	Disposition   string `json:"disposition,omitempty" url:"disposition,omitempty"`
	Statement     string `json:"statement,omitempty" url:"statement,omitempty"`
	StatementID   string `json:"statement_id,omitempty" url:"statement_id,omitempty"`
	OnWaitTimeout string `json:"on_wait_timeout,omitempty" url:"on_wait_timeout,omitempty"`
	Schema        string `json:"schema,omitempty" url:"schema,omitempty"`
	ByteLimit     int    `json:"byte_limit,omitempty" url:"byte_limit,omitempty"`
	WarehouseID   string `json:"warehouse_id,omitempty" url:"warehouse_id"`
}

func (a StatementAPI) executeStatement(warehouseID, statementStr string) (*Statement, error) {
	var statement Statement
	err := a.client.Post(a.context, "/sql/statements", genericStatementRequest{
		WarehouseID: warehouseID,
		Statement:   statementStr,
	}, &statement)
	return &statement, err
}

func (a StatementAPI) getStatement(statementID string) (*Statement, error) {
	var statement Statement
	err := a.client.Get(a.context, fmt.Sprintf("/sql/statements/%s", statementID), genericStatementRequest{
		StatementID: statementID,
	}, &statement)
	return &statement, err
}

func (a StatementAPI) waitForStatementFinished(statementID string) error {
	return resource.RetryContext(a.context, 10*time.Minute, func() *resource.RetryError {
		statementInfo, err := a.getStatement(statementID)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		state := statementInfo.Status.State
		switch state {
		case "CANCELED", "FAILED":
			return resource.NonRetryableError(fmt.Errorf("Statement cannot finish: %s", state))
		case "SUCCEEDED", "CLOSED":
			return nil
		}
		log.Printf("[DEBUG] Statement is in %s state", state)
		return resource.RetryableError(fmt.Errorf(state))
	})
}
