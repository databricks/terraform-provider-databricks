package mlflow

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDataSourceModel(t *testing.T) {
	modelName := "databricks-model"

	var model = ml.ModelDatabricks{
		CreationTimestamp:    0,
		Description:          "Staging model",
		Id:                   "1",
		LastUpdatedTimestamp: 0,
		LatestVersions:       nil,
		Name:                 modelName,
		PermissionLevel:      "CAN_MANAGE",
		Tags:                 nil,
		UserId:               "me@databricks.com",
	}

	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/mlflow/databricks/registered-models/get?name=%s", modelName),
				Response: ml.GetModelResponse{RegisteredModelDatabricks: &model},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceModel(),
		ID:          ".",
		HCL: fmt.Sprintf(`
		name = "%s"
		`, modelName),
	}.ApplyAndExpectData(t, map[string]interface{}{
		"name":             modelName,
		"user_id":          "me@databricks.com",
		"description":      "Staging model",
		"latest_versions":  []interface{}{},
		"id":               "1",
		"permission_level": "CAN_MANAGE",
		"tags":             []interface{}{},
	})
}

func TestDataSourceModelNotFound(t *testing.T) {
	modelName := "databricks-model-non-existent"

	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/mlflow/databricks/registered-models/get?name=%s", modelName),
				Status:   404,
				Response: apierr.APIError{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   fmt.Sprintf("RegisteredModel '%s' does not exist. It might have been deleted.", modelName),
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceModel(),
		ID:          ".",
		HCL: fmt.Sprintf(`
		name = "%s"
		`, modelName),
	}.ExpectError(t, fmt.Sprintf("RegisteredModel '%s' does not exist. It might have been deleted.", modelName))
}
