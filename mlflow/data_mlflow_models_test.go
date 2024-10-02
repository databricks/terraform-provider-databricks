package mlflow

import (
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDataSourceModels(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			api := w.GetMockModelRegistryAPI()
			api.EXPECT().ListModelsAll(mock.Anything, ml.ListModelsRequest{}).Return([]ml.Model{
				{
					Name: "model-01",
				},
				{
					Name: "model-02",
				},
			}, nil)
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceModels(),
		ID:          ".",
	}.ApplyAndExpectData(t, map[string]interface{}{
		"names": []interface{}{"model-01", "model-02"},
	})
}
