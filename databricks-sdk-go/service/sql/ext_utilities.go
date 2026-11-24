package sql

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/retries"
)

type statementExecutionAPIUtilities interface {
	ExecuteAndWait(ctx context.Context, request ExecuteStatementRequest) (*StatementResponse, error)
}

// [EXPERIMENTAL] Execute a query and wait for results to be available
func (a *StatementExecutionAPI) ExecuteAndWait(ctx context.Context, request ExecuteStatementRequest) (*StatementResponse, error) {
	immediateResponse, err := a.ExecuteStatement(ctx, request)
	if err != nil {
		return nil, err
	}
	status := immediateResponse.Status
	switch status.State {
	case StatementStateSucceeded:
		return immediateResponse, nil
	case StatementStateFailed, StatementStateCanceled, StatementStateClosed:
		msg := status.State.String()
		if status.Error != nil {
			msg = fmt.Sprintf("%s: %s %s", msg, status.Error.ErrorCode, status.Error.Message)
		}
		return nil, fmt.Errorf("%s", msg)
	default:
		// TODO: parse request.WaitTimeout and use it here
		return retries.Poll[StatementResponse](ctx, 20*time.Minute,
			func() (*StatementResponse, *retries.Err) {
				res, err := a.GetStatementByStatementId(ctx, immediateResponse.StatementId)
				if err != nil {
					return nil, retries.Halt(err)
				}
				status := res.Status
				switch status.State {
				case StatementStateSucceeded:
					return &StatementResponse{
						Manifest:    res.Manifest,
						Result:      res.Result,
						StatementId: res.StatementId,
						Status:      res.Status,
					}, nil
				case StatementStateFailed, StatementStateCanceled, StatementStateClosed:
					msg := status.State.String()
					if status.Error != nil {
						msg = fmt.Sprintf("%s: %s %s", msg, status.Error.ErrorCode, status.Error.Message)
					}
					return nil, retries.Halt(fmt.Errorf("%s", msg))
				default:
					return nil, retries.Continues(status.State.String())
				}
			})
	}
}
