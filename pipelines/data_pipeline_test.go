package pipelines

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDataSourcePipeline_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourcePipeline(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "I'm a teapot")
}

func TestDataSourcePipelines(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?max_results=100",
				Response: PipelineListResponse{
					Statuses: []PipelineStateInfo{
						{
							PipelineID:      "123",
							Name:            "Pipeline1",
							CreatorUserName: "user1",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?max_results=100&page_token=token1",
				Response: PipelineListResponse{
					Statuses: []PipelineStateInfo{
						{
							PipelineID:      "123",
							Name:            "Pipeline1",
							CreatorUserName: "user1",
						},
					},
					NextPageToken: "token1",
				},
			},
		},
		Resource:    DataSourcePipeline(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": map[string]any{
			"Pipeline1": "123",
		},
	})
}

func TestDataSourcePipelines_Search(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?filter=name+LIKE+%27Pipeline1%27&max_results=100",
				Response: PipelineListResponse{
					Statuses: []PipelineStateInfo{
						{
							PipelineID:      "123",
							Name:            "Pipeline1",
							CreatorUserName: "user1",
						},
					},
				},
			},
		},
		Resource:    DataSourcePipeline(),
		HCL:         `pipeline_name = "Pipeline1"`,
		Read:        true,
		NonWritable: true,
		//Create:      true,
		ID: "_",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": map[string]any{
			"Pipeline1": "123",
		},
	})
}

func TestDataSourcePipelines_SearchError(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?filter=name+LIKE+%27Pipeline2%27&max_results=100",
				Response: PipelineListResponse{},
			},
		},
		Resource:    DataSourcePipeline(),
		HCL:         `pipeline_name = "Pipeline2"`,
		Read:        true,
		NonWritable: true,
		//Create:      true,
		ID: "_",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "there is no pipeline with name LIKE")
}

func TestDataSourcePipelines_NoneFound(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?max_results=100",
				Response: PipelineListResponse{},
			},
		},
		Resource:    DataSourcePipeline(),
		HCL:         `pipeline_name = ""`,
		Read:        true,
		NonWritable: true,
		//Create:      true,
		ID: "_",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "no pipelines found")
}
