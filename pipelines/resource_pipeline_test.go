package pipelines

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

var basicPipelineSpec = pipelineSpec{
	Name:    "test-pipeline",
	Storage: "/test/storage",
	Configuration: map[string]string{
		"key1": "value1",
		"key2": "value2",
	},
	Clusters: []pipelineCluster{
		{
			Label: "default",
			CustomTags: map[string]string{
				"cluster_tag1": "cluster_value1",
			},
		},
	},
	Libraries: []pipelineLibrary{
		{
			Notebook: &notebookLibrary{
				Path: "/Test",
			},
		},
	},
	Filters: &filters{
		Include: []string{"com.databricks.include"},
		Exclude: []string{"com.databricks.exclude"},
	},
}

func TestResourcePipelineCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/pipelines",
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
		cluster {
		  label = "default"
		  custom_tags = {
			"cluster_tag1" = "cluster_value1"
		  }
		}
		library {
		  notebook {
			path = "/Test"
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

func TestResourcePipelineCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/pipelines",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourcePipeline(),
		HCL: `name = "test"
		storage = "/test/storage"
		library {
			notebook {
				path = "/Test"
			}
		}
		filters {
			include = ["a"]
		}
		`,
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourcePipelineCreate_ErrorWhenWaitingFailedCleanup(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/pipelines",
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
					"state": "FAILED",
				},
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.0/pipelines/abcd",
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: common.APIErrorBody{
					ErrorCode: "INTERNAL_ERROR",
					Message:   "Internal error",
				},
				Status: 500,
			},
		},
		Resource: ResourcePipeline(),
		HCL: `name = "test"
		storage = "/test/storage"
		library {
			notebook {
				path = "/Test"
			}
		}
		filters {
			include = ["a"]
		}
		`,
		Create: true,
	}.ExpectError(t, "multiple errors occurred when creating pipeline. "+
		"Error while waiting for creation: \"pipeline abcd has failed\"; "+
		"error while attempting to clean up failed pipeline: \"Internal error\"")
}

func TestResourcePipelineCreate_ErrorWhenWaitingSuccessfulCleanup(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/pipelines",
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
					"state": "FAILED",
				},
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.0/pipelines/abcd",
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: common.APIErrorBody{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   "No such resource",
				},
				Status: 404,
			},
		},
		Resource: ResourcePipeline(),
		HCL: `name = "test"
		storage = "/test/storage"
		library {
			notebook {
				path = "/Test"
			}
		}
		filters {
			include = ["a"]
		}
		`,
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "pipeline abcd has failed")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourcePipelineRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: PipelineInfo{
					PipelineID: "abcd",
					Spec:       &basicPipelineSpec,
				},
			},
		},
		Resource: ResourcePipeline(),
		Read:     true,
		New:      true,
		ID:       "abcd",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abcd", d.Id(), "Id should not be empty")
	assert.Equal(t, "/test/storage", d.Get("storage"))
	assert.Equal(t, "value1", d.Get("configuration.key1"))
	assert.Equal(t, "com.databricks.include", d.Get("filters.0.include.0"))
	assert.Equal(t, false, d.Get("continuous"))
}

func TestResourcePipelineRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourcePipeline(),
		Read:     true,
		Removed:  true,
		ID:       "abcd",
	}.ApplyNoError(t)
}

func TestResourcePipelineRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourcePipeline(),
		Read:     true,
		ID:       "abcd",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abcd", d.Id(), "Id should not be empty for error reads")
}

func TestResourcePipelineUpdate(t *testing.T) {
	state := StateRunning
	spec := pipelineSpec{
		ID:      "abcd",
		Name:    "test",
		Storage: "/test/storage",
		Libraries: []pipelineLibrary{
			{
				Notebook: &notebookLibrary{
					Path: "/Test",
				},
			},
		},
		Filters: &filters{
			Include: []string{"com.databricks.include"},
		},
		Channel: "current",
		Edition: "advanced",
	}
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PUT",
				Resource:        "/api/2.0/pipelines/abcd",
				ExpectedRequest: spec,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: PipelineInfo{
					PipelineID: "abcd",
					Spec:       &spec,
					State:      &state,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: PipelineInfo{
					PipelineID: "abcd",
					Spec:       &spec,
					State:      &state,
				},
			},
		},
		Resource: ResourcePipeline(),
		HCL: `name = "test"
		storage = "/test/storage"
		library {
			notebook {
				path = "/Test"
			}
		}
		filters {
			include = [ "com.databricks.include" ]
		}`,
		InstanceState: map[string]string{
			"name":    "test",
			"storage": "/test/storage",
		},
		Update: true,
		ID:     "abcd",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abcd", d.Id(), "Id should be the same as in reading")
}

func TestResourcePipelineUpdate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for better stub url...
				Method:   "PUT",
				Resource: "/api/2.0/pipelines/abcd",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourcePipeline(),
		HCL: `name = "test"
		storage = "/test/storage"
		library {
			notebook {
				path = "/Test"
			}
		}
		filters {
			include = [ "com.databricks.include" ]
		}`,
		Update: true,
		InstanceState: map[string]string{
			"name":    "test",
			"storage": "/test/storage",
		},
		ID: "abcd",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abcd", d.Id())
}

func TestResourcePipelineUpdate_FailsAfterUpdate(t *testing.T) {
	state := StateFailed
	spec := pipelineSpec{
		ID:      "abcd",
		Name:    "test",
		Storage: "/test/storage",
		Libraries: []pipelineLibrary{
			{
				Notebook: &notebookLibrary{
					Path: "/Test",
				},
			},
		},
		Filters: &filters{
			Include: []string{"com.databricks.include"},
		},
		Channel: "current",
		Edition: "advanced",
	}
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "PUT",
				Resource:        "/api/2.0/pipelines/abcd",
				ExpectedRequest: spec,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: PipelineInfo{
					PipelineID: "abcd",
					Spec:       &spec,
					State:      &state,
				},
			},
		},
		Resource: ResourcePipeline(),
		HCL: `name = "test"
		storage = "/test/storage"
		library {
			notebook {
				path = "/Test"
			}
		}
		filters {
			include = [ "com.databricks.include" ]
		}`,
		Update: true,
		InstanceState: map[string]string{
			"name":    "test",
			"storage": "/test/storage",
		},
		ID: "abcd",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "pipeline abcd has failed")
	assert.Equal(t, "abcd", d.Id(), "Id should be the same as in reading")
}

func TestResourcePipelineDelete(t *testing.T) {
	state := StateRunning
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/pipelines/abcd",
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: PipelineInfo{
					PipelineID: "abcd",
					Spec:       &basicPipelineSpec,
					State:      &state,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: common.APIErrorBody{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   "No such resource",
				},
				Status: 404,
			},
		},
		Resource: ResourcePipeline(),
		Delete:   true,
		ID:       "abcd",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abcd", d.Id())
}

func TestResourcePipelineDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/pipelines/abcd",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 500,
			},
		},
		Resource: ResourcePipeline(),
		Delete:   true,
		ID:       "abcd",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abcd", d.Id())
}
