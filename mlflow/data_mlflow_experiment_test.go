package mlflow

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestDataSourceExperimentById(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			api := w.GetMockExperimentsAPI()
			api.EXPECT().GetExperiment(mock.Anything, ml.GetExperimentRequest{
				ExperimentId: "1234567890",
			}).Return(&ml.GetExperimentResponse{
				Experiment: &ml.Experiment{
					Name:             "/Users/databricks/my-experiment",
					ExperimentId:     "1234567890",
					ArtifactLocation: "dbfs:/databricks/mlflow-tracking/1234567890",
					LifecycleStage:   "active",
				},
			}, nil)
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			api := w.GetMockExperimentsAPI()
			api.EXPECT().GetExperiment(mock.Anything, ml.GetExperimentRequest{
				ExperimentId: "0987654321",
			}).Return(nil, &apierr.APIError{
				ErrorCode:  "RESOURCE_DOES_NOT_EXIST",
				Message:    "Node ID 0987654321 does not exist.",
				StatusCode: 404,
			})
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			api := w.GetMockExperimentsAPI()
			api.EXPECT().GetByName(mock.Anything, ml.GetByNameRequest{
				ExperimentName: experimentName,
			}).Return(&ml.GetExperimentByNameResponse{
				Experiment: &ml.Experiment{
					Name:             experimentName,
					ExperimentId:     "1234567890",
					ArtifactLocation: "dbfs:/databricks/mlflow-tracking/1234567890",
					LifecycleStage:   "active",
				},
			}, nil)
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			api := w.GetMockExperimentsAPI()
			api.EXPECT().GetByName(mock.Anything, ml.GetByNameRequest{
				ExperimentName: experimentName,
			}).Return(nil, &apierr.APIError{
				ErrorCode:  "RESOURCE_DOES_NOT_EXIST",
				Message:    "Node /Users/databricks/non-existent-experiment does not exist.",
				StatusCode: 404,
			})
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceExperiment(),
		ID:          ".",
		HCL:         fmt.Sprintf(`name = "%s"`, experimentName),
	}.ExpectError(t, "Node /Users/databricks/non-existent-experiment does not exist.")
}

func TestDataSourceExperimentByShortName(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			api := w.GetMockExperimentsAPI()
			api.EXPECT().SearchExperimentsAll(mock.Anything, ml.SearchExperiments{
				Filter: "name LIKE '%/my-experiment'",
			}).Return([]ml.Experiment{
				{
					Name:         "/Users/databricks/another-my-experiment",
					ExperimentId: "1111111111",
				},
				{
					Name:             "/Users/databricks/my-experiment",
					ExperimentId:     "1234567890",
					ArtifactLocation: "dbfs:/databricks/mlflow-tracking/1234567890",
					LifecycleStage:   "active",
				},
			}, nil)
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceExperiment(),
		ID:          ".",
		HCL:         `name = "my-experiment"`,
	}.ApplyAndExpectData(t, map[string]interface{}{
		"id":                "1234567890",
		"artifact_location": "dbfs:/databricks/mlflow-tracking/1234567890",
		"experiment_id":     "1234567890",
		"lifecycle_stage":   "active",
		"name":              "/Users/databricks/my-experiment",
	})
}

func TestDataSourceExperimentByRelativePath(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			api := w.GetMockExperimentsAPI()
			api.EXPECT().SearchExperimentsAll(mock.Anything, ml.SearchExperiments{
				Filter: "name LIKE '%/team/my-experiment'",
			}).Return([]ml.Experiment{
				{
					Name:         "/Shared/team/my-experiment",
					ExperimentId: "1234567890",
				},
			}, nil)
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceExperiment(),
		ID:          ".",
		HCL:         `name = "team/my-experiment"`,
	}.ApplyAndExpectData(t, map[string]interface{}{
		"id":   "1234567890",
		"name": "/Shared/team/my-experiment",
	})
}

func TestDataSourceExperimentByShortNameEscapesFilter(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			api := w.GetMockExperimentsAPI()
			api.EXPECT().SearchExperimentsAll(mock.Anything, ml.SearchExperiments{
				Filter: "name LIKE '%/joe_s _100__ experiment'",
			}).Return([]ml.Experiment{
				{
					Name:         "/Users/databricks/joe-s [100]] experiment",
					ExperimentId: "1111111111",
				},
				{
					Name:         `/Users/databricks/joe's "100%" experiment`,
					ExperimentId: "1234567890",
				},
			}, nil)
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceExperiment(),
		ID:          ".",
		HCL:         `name = "joe's \"100%\" experiment"`,
	}.ApplyAndExpectData(t, map[string]interface{}{
		"id":   "1234567890",
		"name": `/Users/databricks/joe's "100%" experiment`,
	})
}

func TestDataSourceExperimentByShortNameNotFound(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			api := w.GetMockExperimentsAPI()
			api.EXPECT().SearchExperimentsAll(mock.Anything, ml.SearchExperiments{
				Filter: "name LIKE '%/my_experiment'",
			}).Return([]ml.Experiment{
				{
					Name:         "/Users/databricks/my-experiment",
					ExperimentId: "1234567890",
				},
			}, nil)
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceExperiment(),
		ID:          ".",
		HCL:         `name = "my_experiment"`,
	}.ExpectError(t, "experiment with name 'my_experiment' does not exist")
}

func TestDataSourceExperimentByShortNameAmbiguous(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			api := w.GetMockExperimentsAPI()
			api.EXPECT().SearchExperimentsAll(mock.Anything, ml.SearchExperiments{
				Filter: "name LIKE '%/my-experiment'",
			}).Return([]ml.Experiment{
				{
					Name:         "/Users/second@example.com/my-experiment",
					ExperimentId: "2222222222",
				},
				{
					Name:         "/Users/first@example.com/my-experiment",
					ExperimentId: "1111111111",
				},
			}, nil)
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceExperiment(),
		ID:          ".",
		HCL:         `name = "my-experiment"`,
	}.ExpectError(t, "there are 2 experiments with name 'my-experiment': "+
		"/Users/first@example.com/my-experiment, /Users/second@example.com/my-experiment. "+
		"Use the full path of the experiment instead")
}

func TestDataSourceExperimentByShortNameSearchError(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			api := w.GetMockExperimentsAPI()
			api.EXPECT().SearchExperimentsAll(mock.Anything, ml.SearchExperiments{
				Filter: "name LIKE '%/my-experiment'",
			}).Return(nil, &apierr.APIError{
				ErrorCode:  "PERMISSION_DENIED",
				Message:    "Permission denied",
				StatusCode: 403,
			})
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceExperiment(),
		ID:          ".",
		HCL:         `name = "my-experiment"`,
	}.ExpectError(t, "Permission denied")
}
