package pipelines

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestDataSourcePipeline_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourcePipelines(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}

func TestDataSourcePipelines(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockPipelinesAPI().EXPECT().
				ListPipelinesAll(mock.Anything, pipelines.ListPipelinesRequest{MaxResults: 100}).
				Return([]pipelines.PipelineStateInfo{
					{
						PipelineId:      "123",
						Name:            "Pipeline1",
						CreatorUserName: "user1",
					},
				}, nil)
		},
		Resource:    DataSourcePipelines(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": []string{
			"123",
		},
	})
}

func TestDataSourcePipelines_Search(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockPipelinesAPI().EXPECT().
				ListPipelinesAll(mock.Anything, pipelines.ListPipelinesRequest{
					Filter:     "name LIKE 'Pipeline1'",
					MaxResults: 100,
				}).
				Return([]pipelines.PipelineStateInfo{
					{
						PipelineId:      "123",
						Name:            "Pipeline1",
						CreatorUserName: "user1",
					},
				}, nil)
		},
		Resource:    DataSourcePipelines(),
		HCL:         `pipeline_name = "Pipeline1"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": []string{
			"123",
		},
	})
}

func TestDataSourcePipelines_SearchError(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockPipelinesAPI().EXPECT().
				ListPipelinesAll(mock.Anything, pipelines.ListPipelinesRequest{
					Filter:     "name LIKE 'Pipeline2'",
					MaxResults: 100,
				}).
				Return([]pipelines.PipelineStateInfo{}, nil)
		},
		Resource:    DataSourcePipelines(),
		HCL:         `pipeline_name = "Pipeline2"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyNoError(t)
}

func TestDataSourcePipelines_NoneFound(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockPipelinesAPI().EXPECT().
				ListPipelinesAll(mock.Anything, pipelines.ListPipelinesRequest{MaxResults: 100}).
				Return([]pipelines.PipelineStateInfo{}, nil)
		},
		Resource:    DataSourcePipelines(),
		HCL:         `pipeline_name = ""`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyNoError(t)
}
