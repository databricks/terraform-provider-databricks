// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package qualitymonitorv2

import (
	"context"
)

// Deprecated: Please use the Data Quality Monitoring API instead (REST:
// /api/data-quality/v1/monitors). Manage data quality of UC objects (currently
// support `schema`).
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type QualityMonitorV2Service interface {

	// Deprecated: Use Data Quality Monitoring API instead
	// (/api/data-quality/v1/monitors). Create a quality monitor on UC object.
	CreateQualityMonitor(ctx context.Context, request CreateQualityMonitorRequest) (*QualityMonitor, error)

	// Deprecated: Use Data Quality Monitoring API instead
	// (/api/data-quality/v1/monitors). Delete a quality monitor on UC object.
	DeleteQualityMonitor(ctx context.Context, request DeleteQualityMonitorRequest) error

	// Deprecated: Use Data Quality Monitoring API instead
	// (/api/data-quality/v1/monitors). Read a quality monitor on UC object.
	GetQualityMonitor(ctx context.Context, request GetQualityMonitorRequest) (*QualityMonitor, error)

	// Deprecated: Use Data Quality Monitoring API instead
	// (/api/data-quality/v1/monitors). (Unimplemented) List quality monitors.
	ListQualityMonitor(ctx context.Context, request ListQualityMonitorRequest) (*ListQualityMonitorResponse, error)

	// Deprecated: Use Data Quality Monitoring API instead
	// (/api/data-quality/v1/monitors). (Unimplemented) Update a quality monitor
	// on UC object.
	UpdateQualityMonitor(ctx context.Context, request UpdateQualityMonitorRequest) (*QualityMonitor, error)
}
