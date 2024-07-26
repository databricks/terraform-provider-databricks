package pipelines

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/terraform-provider-databricks/qa"
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?max_results=100",
				Response: pipelines.ListPipelinesResponse{
					Statuses: []pipelines.PipelineStateInfo{
						{
							PipelineId:      "123",
							Name:            "Pipeline1",
							CreatorUserName: "user1",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?max_results=100&page_token=token1",
				Response: pipelines.ListPipelinesResponse{
					Statuses: []pipelines.PipelineStateInfo{
						{
							PipelineId:      "123",
							Name:            "Pipeline1",
							CreatorUserName: "user1",
						},
					},
					NextPageToken: "token1",
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?filter=name+LIKE+%27Pipeline1%27&max_results=100",
				Response: pipelines.ListPipelinesResponse{
					Statuses: []pipelines.PipelineStateInfo{
						{
							PipelineId:      "123",
							Name:            "Pipeline1",
							CreatorUserName: "user1",
						},
					},
				},
			},
		},
		Resource:    DataSourcePipelines(),
		HCL:         `pipeline_name = "Pipeline1"`,
		Read:        true,
		NonWritable: true,
		//Create:      true,
		ID: "_",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": []string{
			"123",
		},
	})
}

func TestDataSourcePipelines_SearchError(t *testing.T) {
	//_, err := qa.ResourceFixture{
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?filter=name+LIKE+%27Pipeline2%27&max_results=100",
				Response: pipelines.ListPipelineEventsResponse{},
			},
		},
		Resource:    DataSourcePipelines(),
		HCL:         `pipeline_name = "Pipeline2"`,
		Read:        true,
		NonWritable: true,
		//Create:      true,
		ID: "_",
	}.ApplyNoError(t)
}

func TestDataSourcePipelines_NoneFound(t *testing.T) {
	//_, err := qa.ResourceFixture{
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?max_results=100",
				Response: pipelines.ListPipelineEventsResponse{},
			},
		},
		Resource:    DataSourcePipelines(),
		HCL:         `pipeline_name = ""`,
		Read:        true,
		NonWritable: true,
		//Create:      true,
		ID: "_",
	}.ApplyNoError(t)
}
