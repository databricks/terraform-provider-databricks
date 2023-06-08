package sql

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func commonFixtureWithStatusResponse(response Statement) []qa.HTTPFixture {
	return []qa.HTTPFixture{
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.0/sql/warehouses/abc",
			Response: SQLEndpoint{
				ID:    "abc",
				Name:  "terraform-sql-table",
				State: "RUNNING",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/sql/statements",
			ExpectedRequest: map[string]any{
				"warehouse_id": "abc",
				"statement":    "SELECT 'done'",
			},
			Response: Statement{
				StatementID: "statement1",
				Status: &Status{
					State: "RUNNING",
				},
			},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.0/sql/statements?statement_id=statement1&warehouse_id=",
			Response:     response,
		},
	}
}

func TestStatementAPI_StatementFailed(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, commonFixtureWithStatusResponse(Statement{
		StatementID: "statement2",
		Status: &Status{
			State: "FAILED",
		},
	}))
	defer server.Close()
	require.NoError(t, err)
	ctx := context.Background()
	statements := NewStatementAPI(ctx, client)

	err = statements.Execute("abc", `SELECT 'done'`)
	assert.EqualError(t, err, "Statement cannot finish: FAILED")
}

func TestStatementAPI_Execute(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, commonFixtureWithStatusResponse(Statement{
		StatementID: "statement3",
		Status: &Status{
			State: "SUCCEEDED",
		},
	}))
	defer server.Close()
	require.NoError(t, err)
	ctx := context.Background()
	statements := NewStatementAPI(ctx, client)

	err = statements.Execute("abc", `SELECT 'done'`)
	assert.NoError(t, err)
}
