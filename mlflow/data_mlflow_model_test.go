package mlflow

import (
	"fmt"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/qa"
	"testing"
)

func TestDataSourceModelSpecificVersion(t *testing.T) {
	modelName := "databricks-model"
	modelVersion := "42"

	var model = ml.ModelVersion{
		CreationTimestamp:    0,
		CurrentStage:         "Staging",
		Description:          "Staging model",
		LastUpdatedTimestamp: 0,
		Name:                 modelName,
		RunId:                "1",
		RunLink:              "",
		Source:               "dbfs:/path/to/model",
		Status:               "READY",
		StatusMessage:        "",
		Tags:                 nil,
		UserId:               "me@databricks.com",
		Version:              modelVersion,
	}

	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/model-versions/get?name=databricks-model&version=42",
				Response: ml.GetModelVersionResponse{ModelVersion: &model},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceModel(),
		ID:          ".",
		HCL: fmt.Sprintf(`
		name = "%s"
		version = "%s"
		`, modelName, modelVersion),
	}.ApplyAndExpectData(t, map[string]interface{}{
		"model_versions": []interface{}{
			map[string]any{
				"creation_timestamp":     0,
				"current_stage":          "Staging",
				"description":            "Staging model",
				"last_updated_timestamp": 0,
				"name":                   modelName,
				"run_id":                 "1",
				"run_link":               "",
				"source":                 "dbfs:/path/to/model",
				"status":                 "READY",
				"status_message":         "",
				"tags":                   []interface{}{},
				"user_id":                "me@databricks.com",
				"version":                modelVersion,
			},
		},
	})
}

func TestDataSourceModelSpecificVersionNotFound(t *testing.T) {
	modelName := "databricks-model-non-existent-version"
	modelVersion := "99"

	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/model-versions/get?name=databricks-model-non-existent-version&version=99",
				Status:   404,
				Response: apierr.APIErrorBody{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   fmt.Sprintf("No model with version %s for model with name: %s.", modelVersion, modelName),
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceModel(),
		ID:          ".",
		HCL: fmt.Sprintf(`
		name = "%s"
		version = "%s"
		`, modelName, modelVersion),
	}.ExpectError(t, fmt.Sprintf("No model with version %s for model with name: %s.", modelVersion, modelName))
}

func TestDataSourceModelLatestVersion(t *testing.T) {
	modelName := "databricks-model-latest-version"

	var models []ml.ModelVersion

	models = append(models, ml.ModelVersion{
		CreationTimestamp:    0,
		CurrentStage:         "Staging",
		Description:          "Staging model",
		LastUpdatedTimestamp: 0,
		Name:                 modelName,
		RunId:                "1",
		RunLink:              "",
		Source:               "dbfs:/path/to/model",
		Status:               "READY",
		StatusMessage:        "",
		Tags:                 nil,
		UserId:               "me@databricks.com",
		Version:              "1",
	})

	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/mlflow/registered-models/get-latest-versions",
				Response: ml.GetLatestVersionsResponse{ModelVersions: models},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceModel(),
		ID:          ".",
		HCL:         fmt.Sprintf(`name = "%s"`, modelName),
	}.ApplyAndExpectData(t, map[string]interface{}{
		"model_versions": []interface{}{
			map[string]any{
				"creation_timestamp":     0,
				"current_stage":          "Staging",
				"description":            "Staging model",
				"last_updated_timestamp": 0,
				"name":                   modelName,
				"run_id":                 "1",
				"run_link":               "",
				"source":                 "dbfs:/path/to/model",
				"status":                 "READY",
				"status_message":         "",
				"tags":                   []interface{}{},
				"user_id":                "me@databricks.com",
				"version":                "1",
			},
		},
	})
}

func TestDataSourceModelByNameNotFound(t *testing.T) {
	modelName := "databricks-model-non-existent"

	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/mlflow/registered-models/get-latest-versions",
				Status:   404,
				Response: apierr.APIErrorBody{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   fmt.Sprintf("RegisteredModel '%s' does not exist. It might have been deleted.", modelName),
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceModel(),
		ID:          ".",
		HCL:         fmt.Sprintf(`name = "%s"`, modelName),
	}.ExpectError(t, fmt.Sprintf("RegisteredModel '%s' does not exist. It might have been deleted.", modelName))
}
