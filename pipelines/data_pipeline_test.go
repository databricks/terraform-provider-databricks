package pipelines

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

// func TestDataSourcePipeline(t *testing.T) {
// 	qa.ResourceFixture{
// 		Fixtures: []qa.HTTPFixture{
// 			{
// 				Method:   "GET",
// 				Resource: "/api/2.0/pipelines?max_results=100",
// 				//Resource: "/api/2.0/pipelines?filter=name%20LIKE%20%27%25abc%25%27",

// 				Response: map[string]any{
// 					"id":    "123",
// 					"name":  "abc",
// 				},
// 					// {
// 					// 	PipelineID: "456",
// 					// 	Name:       "def",
// 					// },
// 				},
// 			},
// 		},
// 		Resource:    DataSourcePipeline(),
// 		Read:        true,
// 		NonWritable: true,
// 		ID:          "_",
// 	}.ApplyAndExpectData(t, map[string]any{
// 		"ids": map[string]any{
// 			"abc": "123",
// 			//"def": "456",
// 		},
// 	})
// }

func TestDataSourcePipeline_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourcePipeline(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "I'm a teapot")
}
