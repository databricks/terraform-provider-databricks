package compute

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

var basicPipelineSpec = map[string]interface{}{
	"name":    "test-pipeline",
	"storage": "/test/storage",
	"configuration": map[string]string{
		"key1": "value1",
		"key2": "value2",
	},
	"clusters": []map[string]interface{}{
		{
			"label":       "default",
			"num_workers": 2,
			"custom_tags": map[string]string{
				"cluster_tag1": "cluster_value1",
			},
		},
	},
	"libraries": []map[string]interface{}{
		{
			"jar": "dbfs:/pipelines/code/abcde.jar",
		},
		{
			"maven": map[string]string{
				"coordinates": "com.microsoft.azure:azure-eventhubs-spark_2.12:2.3.18",
			},
		},
	},
	"filters": map[string][]string{
		"include": {"com.databricks.include"},
		"exclude": {"com.databricks.exclude"},
	},
}

func TestResourcePipelineCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "POST",
				Resource:        "/api/2.0/pipelines",
				ExpectedRequest: basicPipelineSpec,
				Response: createPipelineResponse{
					PipelineID: "abcd",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: map[string]interface{}{
					"id":    "abcd",
					"name":  "test-pipeline",
					"state": "DEPLOYING",
					"spec":  basicPipelineSpec,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: map[string]interface{}{
					"id":    "abcd",
					"name":  "test-pipeline",
					"state": "RUNNING",
					"spec":  basicPipelineSpec,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: map[string]interface{}{
					"id":    "abcd",
					"name":  "test-pipeline",
					"state": "RUNNING",
					"spec":  basicPipelineSpec,
				},
			},
		},
		Create:   true,
		Resource: ResourcePipeline(),
		HCL: `name = "test-pipeline"
		storage = "/test/storage"
		configuration = {
		  key1 = "value1"
		  key2 = "value2"
		}
		clusters {
		  label = "default"
		  num_workers = 2
		  custom_tags = {
			"cluster_tag1" = "cluster_value1"
		  }
		}
		libraries {
		  jar = "dbfs:/pipelines/code/abcde.jar"
		}
		libraries {
		  maven {
			coordinates = "com.microsoft.azure:azure-eventhubs-spark_2.12:2.3.18"
		  }
		}
		filters {
		  include = ["com.databricks.include"]
		  exclude = ["com.databricks.exclude"]
		}
		continuous = false
		`,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abcd", d.Id())
}
