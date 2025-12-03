package exporter

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	sdk_dataquality "github.com/databricks/databricks-sdk-go/service/dataquality"
	"github.com/databricks/databricks-sdk-go/service/qualitymonitorv2"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/stretchr/testify/assert"
)

func TestQualityMonitorV2Export(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		meAdminFixture,
		noCurrentMetastoreAttached,
		emptyDataQualityMonitors,
		{
			Method:   "GET",
			Resource: "/api/2.0/quality-monitors?",
			Response: qualitymonitorv2.ListQualityMonitorResponse{
				QualityMonitors: []qualitymonitorv2.QualityMonitor{
					{
						ObjectType: "schema",
						ObjectId:   "abc-123-def",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/quality-monitors/schema/abc-123-def?",
			Response: qualitymonitorv2.QualityMonitor{
				ObjectType: "schema",
				ObjectId:   "abc-123-def",
				AnomalyDetectionConfig: &qualitymonitorv2.AnomalyDetectionConfig{
					LastRunId:       "run-123",
					LatestRunStatus: "ANOMALY_DETECTION_RUN_STATUS_SUCCESS",
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)

		ic := newImportContext(client)
		ic.noFormat = true
		ic.Directory = tmpDir
		ic.enableListing("dq")
		ic.enableServices("dq")

		err := ic.Run()
		assert.NoError(t, err)

		// Verify that the quality monitor v2 was exported
		content, err := os.ReadFile(tmpDir + "/dq.tf")
		assert.NoError(t, err)
		contentStr := string(content)

		assert.Contains(t, contentStr, `resource "databricks_quality_monitor_v2"`)
		assert.Contains(t, contentStr, `object_type = "schema"`)
		assert.Contains(t, contentStr, `object_id   = "abc-123-def"`)
	})
}

func TestQualityMonitorV2ExportWithMultipleMonitors(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.0/preview/scim/v2/Me",
			Response: scim.User{
				Groups: []scim.ComplexValue{
					{
						Display: "admins",
					},
				},
				UserName: "user@domain.com",
			},
		},
		noCurrentMetastoreAttached,
		emptyDataQualityMonitors,
		{
			Method:   "GET",
			Resource: "/api/2.0/quality-monitors?",
			Response: qualitymonitorv2.ListQualityMonitorResponse{
				QualityMonitors: []qualitymonitorv2.QualityMonitor{
					{
						ObjectType: "schema",
						ObjectId:   "schema-uuid-1",
					},
					{
						ObjectType: "schema",
						ObjectId:   "schema-uuid-2",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/quality-monitors/schema/schema-uuid-1?",
			Response: qualitymonitorv2.QualityMonitor{
				ObjectType: "schema",
				ObjectId:   "schema-uuid-1",
				AnomalyDetectionConfig: &qualitymonitorv2.AnomalyDetectionConfig{
					LastRunId:       "run-1",
					LatestRunStatus: "ANOMALY_DETECTION_RUN_STATUS_SUCCESS",
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/quality-monitors/schema/schema-uuid-2?",
			Response: qualitymonitorv2.QualityMonitor{
				ObjectType: "schema",
				ObjectId:   "schema-uuid-2",
				AnomalyDetectionConfig: &qualitymonitorv2.AnomalyDetectionConfig{
					LastRunId:       "run-2",
					LatestRunStatus: "ANOMALY_DETECTION_RUN_STATUS_FAILED",
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)

		ic := newImportContext(client)
		ic.noFormat = true
		ic.Directory = tmpDir
		ic.enableListing("dq")
		ic.enableServices("dq")

		err := ic.Run()
		assert.NoError(t, err)

		// Verify that both quality monitors v2 were exported
		content, err := os.ReadFile(tmpDir + "/dq.tf")
		assert.NoError(t, err)
		contentStr := string(content)

		assert.Contains(t, contentStr, `object_id   = "schema-uuid-1"`)
		assert.Contains(t, contentStr, `object_id   = "schema-uuid-2"`)
	})
}

func TestDataQualityMonitorsExport(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		meAdminFixture,
		noCurrentMetastoreAttached,
		emptyQualityMonitorsV2,
		{
			Method:   "GET",
			Resource: "/api/data-quality/v1/monitors?",
			Response: sdk_dataquality.ListMonitorResponse{
				Monitors: []sdk_dataquality.Monitor{
					{
						ObjectId:   "123",
						ObjectType: "table",
					},
				},
			},
			ReuseRequest: true,
		},
		{
			Method:   "GET",
			Resource: "/api/data-quality/v1/monitors/table/123?",
			Response: sdk_dataquality.Monitor{
				ObjectId:               "123",
				ObjectType:             "table",
				AnomalyDetectionConfig: &sdk_dataquality.AnomalyDetectionConfig{},
			},
			ReuseRequest: true,
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)

		ic := newImportContext(client)
		ic.noFormat = true
		ic.Directory = tmpDir
		ic.enableListing("dq")
		ic.enableServices("dq")

		err := ic.Run()
		assert.NoError(t, err)

		content, err := os.ReadFile(tmpDir + "/dq.tf")
		assert.NoError(t, err)
		contentStr := string(content)
		assert.True(t, strings.Contains(contentStr, `resource "databricks_data_quality_monitor" "table_123" {
  object_type              = "table"
  object_id                = "123"
  anomaly_detection_config = {}
}`))
	})
}
