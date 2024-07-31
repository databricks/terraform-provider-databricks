package mlflow

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/ml"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDataSourceExperimentById(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/experiments/get?experiment_id=1234567890",
				Response: ml.GetExperimentResponse{
					Experiment: &ml.Experiment{
						Name:             "/Users/databricks/my-experiment",
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
		ID:          "_",
		HCL:         fmt.Sprintf(`experiment_id = "%s"`, "1234567890"),
	}.ApplyAndExpectData(t, map[string]interface{}{
		"artifact_location": "dbfs:/databricks/mlflow-tracking/1234567890",
		"creation_time":     0,
		"experiment_id":     "1234567890",
		"last_update_time":  0,
		"lifecycle_stage":   "active",
		"name":              "/Users/databricks/my-experiment",
		"tags":              []interface{}{},
	})
}

func TestDataSourceExperimentByIdNotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/experiments/get?experiment_id=0987654321",
				Status:   404,
				Response: apierr.APIError{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   "Node ID 0987654321 does not exist.",
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceExperiment(),
		ID:          ".",
		HCL:         fmt.Sprintf(`experiment_id = "%s"`, "0987654321"),
	}.ExpectError(t, "Node ID 0987654321 does not exist.")
}

func TestDataSourceExperimentByName(t *testing.T) {
	experimentName := "/Users/databricks/my-experiment"

	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/mlflow/experiments/get-by-name?experiment_name=%s", url.QueryEscape(experimentName)),
				Response: ml.GetExperimentResponse{
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
	}.ApplyAndExpectData(t, map[string]interface{}{
		"artifact_location": "dbfs:/databricks/mlflow-tracking/1234567890",
		"creation_time":     0,
		"experiment_id":     "1234567890",
		"last_update_time":  0,
		"lifecycle_stage":   "active",
		"name":              experimentName,
		"tags":              []interface{}{},
	})
}

func TestDataSourceExperimentByNameNotFound(t *testing.T) {
	experimentName := "/Users/databricks/non-existent-experiment"

	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/mlflow/experiments/get-by-name?experiment_name=%s", url.QueryEscape(experimentName)),
				Status:   404,
				Response: apierr.APIError{
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

func TestDataSourceExperimentByNameInvalidPath(t *testing.T) {
	experimentName := "invalid_path"

	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/mlflow/experiments/get-by-name?experiment_name=%s", url.QueryEscape(experimentName)),
				Status:   404,
				Response: apierr.APIError{
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
