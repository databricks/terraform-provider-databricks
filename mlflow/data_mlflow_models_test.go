package mlflow

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDataSourceModels(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/mlflow/registered-models/list?",
				Response: ml.ListModelsResponse{
					NextPageToken: "",
					RegisteredModels: []ml.Model{
						{
							Name: "model-01",
						},
						{
							Name: "model-02",
						},
					},
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceModels(),
		ID:          ".",
	}.ApplyAndExpectData(t, map[string]interface{}{
		"names": []interface{}{"model-01", "model-02"},
	})
}
