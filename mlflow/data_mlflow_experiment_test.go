package mlflow

import (
	"fmt"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"net/url"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDataSourceExperiment(t *testing.T) {
	experimentName := "/Users/databricks/my-experiment"

	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/mlflow/experiments/get-by-name?experiment_name=%s", url.QueryEscape(experimentName)),
				Response: ml.GetExperimentByNameResponse{
					Experiment: &ml.Experiment{
						Name:             experimentName,
						ExperimentId:     "1234567890",
						ArtifactLocation: "dbfs:/databricks/mlflow-tracking/1234567890",
						LifecycleStage:   "active",
					},
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceExperiment(),
		ID:          ".",
		HCL:         fmt.Sprintf(`name = "%s"`, experimentName),
	}.ApplyAndExpectData(t, map[string]any{
		"artifact_location": "dbfs:/databricks/mlflow-tracking/1234567890",
		"id":                "1234567890",
		"lifecycle_stage":   "active",
		"name":              experimentName,
	})
}

func TestDataSourceExperimentNotFound(t *testing.T) {
	experimentName := "/Users/databricks/non-existent-experiment"

	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/mlflow/experiments/get-by-name?experiment_name=%s", url.QueryEscape(experimentName)),
				Status:   404,
				Response: apierr.APIErrorBody{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   "Node /Users/databricks/non-existent-experiment does not exist.",
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceExperiment(),
		ID:          ".",
		HCL:         fmt.Sprintf(`name = "%s"`, experimentName),
	}.ExpectError(t, "Node /Users/databricks/non-existent-experiment does not exist.")
}

func TestDataSourceExperimentInvalidPath(t *testing.T) {
	experimentName := "invalid_path"

	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/mlflow/experiments/get-by-name?experiment_name=%s", url.QueryEscape(experimentName)),
				Status:   404,
				Response: apierr.APIErrorBody{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   "Got an invalid experiment name 'invalid_path'. An experiment name must be an absolute path within the Databricks workspace, e.g. '/Users/<some-username>/my-experiment'.",
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceExperiment(),
		ID:          ".",
		HCL:         fmt.Sprintf(`name = "%s"`, experimentName),
	}.ExpectError(t, "Got an invalid experiment name 'invalid_path'. An experiment name must be an absolute path within the Databricks workspace, e.g. '/Users/<some-username>/my-experiment'.")
}
