package pipelines

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestResourcePipelineReadPreservesExpectedLastModified(t *testing.T) {
	for _, expected := range []int64{0, 123456} {
		t.Run(fmt.Sprintf("expected_%d", expected), func(t *testing.T) {
			qa.ResourceFixture{
				MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
					w.GetMockPipelinesAPI().EXPECT().Get(mock.Anything, pipelines.GetPipelineRequest{
						PipelineId: "pipeline-1",
					}).Return(&pipelines.GetPipelineResponse{
						PipelineId:   "pipeline-1",
						LastModified: 123457,
						Spec:         &pipelines.PipelineSpec{Name: "test-pipeline"},
					}, nil).Once()
				},
				Resource: ResourcePipeline(),
				Read:     true,
				ID:       "pipeline-1",
				HCL: fmt.Sprintf(`
					name = "test-pipeline"
					expected_last_modified = %d
				`, expected),
			}.ApplyAndExpectData(t, map[string]any{
				"expected_last_modified": int(expected),
				"last_modified":          123457,
			})
		})
	}
}
