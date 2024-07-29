package pipelines

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var basicPipelineSpec = PipelineSpec{
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
	Libraries: []PipelineLibrary{
		{
			Notebook: &NotebookLibrary{
				Path: "/Test",
			},
		},
	},
	Filters: &filters{
		Include: []string{"com.databricks.include"},
		Exclude: []string{"com.databricks.exclude"},
	},
	Deployment: &PipelineDeployment{
		Kind:             "BUNDLE",
		MetadataFilePath: "/foo/bar",
	},
	Edition: "ADVANCED",
	Channel: "CURRENT",
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
				Response: map[string]any{
					"id":    "abcd",
					"name":  "test-pipeline",
					"state": "DEPLOYING",
					"spec":  basicPipelineSpec,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: map[string]any{
					"id":    "abcd",
					"name":  "test-pipeline",
					"state": "RUNNING",
					"spec":  basicPipelineSpec,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: map[string]any{
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
		deployment {
			kind = "BUNDLE"
			metadata_file_path = "/foo/bar"
		}
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abcd", d.Id())
}

func TestResourcePipelineCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/pipelines",
				Response: apierr.APIError{
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
				Response: map[string]any{
					"id":    "abcd",
					"name":  "test-pipeline",
					"state": "FAILED",
				},
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.0/pipelines/abcd?",
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: apierr.APIError{
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
				Response: map[string]any{
					"id":    "abcd",
					"name":  "test-pipeline",
					"state": "FAILED",
				},
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.0/pipelines/abcd?",
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: apierr.APIError{
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
	assert.NoError(t, err)
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
				Response: apierr.APIError{
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
				Response: apierr.APIError{
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
	spec := PipelineSpec{
		ID:      "abcd",
		Name:    "test",
		Storage: "/test/storage",
		Libraries: []PipelineLibrary{
			{
				Notebook: &NotebookLibrary{
					Path: "/Test",
				},
			},
		},
		Filters: &filters{
			Include: []string{"com.databricks.include"},
		},
		Channel: "CURRENT",
		Edition: "ADVANCED",
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
	assert.NoError(t, err)
	assert.Equal(t, "abcd", d.Id(), "Id should be the same as in reading")
}

func TestResourcePipelineUpdate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for better stub url...
				Method:   "PUT",
				Resource: "/api/2.0/pipelines/abcd",
				Response: apierr.APIError{
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
	spec := PipelineSpec{
		ID:      "abcd",
		Name:    "test",
		Storage: "/test/storage",
		Libraries: []PipelineLibrary{
			{
				Notebook: &NotebookLibrary{
					Path: "/Test",
				},
			},
		},
		Filters: &filters{
			Include: []string{"com.databricks.include"},
		},
		Channel: "CURRENT",
		Edition: "ADVANCED",
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
				Resource: "/api/2.0/pipelines/abcd?",
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
				Response: apierr.APIError{
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
	assert.NoError(t, err)
	assert.Equal(t, "abcd", d.Id())
}

func TestResourcePipelineDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/pipelines/abcd?",
				Response: apierr.APIError{
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

func TestListPipelines(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/pipelines?max_results=1",
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
		{
			Method:   "GET",
			Resource: "/api/2.0/pipelines?max_results=1&page_token=token1",
			Response: PipelineListResponse{
				Statuses: []PipelineStateInfo{
					{
						PipelineID:      "456",
						Name:            "Pipeline2",
						CreatorUserName: "user2",
					},
				},
				PrevPageToken: "token0",
			},
		},
	})
	defer server.Close()
	require.NoError(t, err)

	ctx := context.Background()
	data, err := NewPipelinesAPI(ctx, client).List(1, "")
	require.NoError(t, err)
	require.Equal(t, 2, len(data))
	require.Equal(t, "Pipeline1", data[0].Name)
	require.Equal(t, "456", data[1].PipelineID)
}

func TestListPipelinesWithFilter(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/pipelines?filter=name%20LIKE%20%27Pipeline1%27&max_results=1",
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
	})
	defer server.Close()
	require.NoError(t, err)

	ctx := context.Background()
	data, err := NewPipelinesAPI(ctx, client).List(1, "name LIKE 'Pipeline1'")
	require.NoError(t, err)
	require.Equal(t, 1, len(data))
}

func TestStorageSuppressDiff(t *testing.T) {
	k := "storage"
	generated := "dbfs:/pipelines/c609bbb0-2e42-4bc8-bb4e-a1c26d6e9403"
	require.True(t, suppressStorageDiff(k, generated, "", nil))
	require.False(t, suppressStorageDiff(k, generated, "/tmp/abc", nil))
	require.False(t, suppressStorageDiff(k, "/tmp/abc", "", nil))
}

func TestResourcePipelineCreateServerless(t *testing.T) {
	var serverlessPipelineSpec = PipelineSpec{
		Name:    "test-pipeline-serverless",
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
		Libraries: []PipelineLibrary{
			{
				Notebook: &NotebookLibrary{
					Path: "/TestServerless",
				},
			},
		},
		Filters: &filters{
			Include: []string{"com.databricks.include"},
			Exclude: []string{"com.databricks.exclude"},
		},
		Serverless: true,
	}
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/pipelines",
				Response: createPipelineResponse{
					PipelineID: "serverless",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/serverless",
				Response: map[string]any{
					"id":    "serverless",
					"name":  "test-pipeline-serverless",
					"state": "DEPLOYING",
					"spec":  serverlessPipelineSpec,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/serverless",
				Response: map[string]any{
					"id":    "serverless",
					"name":  "test-pipeline-serverless",
					"state": "RUNNING",
					"spec":  serverlessPipelineSpec,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/serverless",
				Response: map[string]any{
					"id":    "serverless",
					"name":  "test-pipeline-serverless",
					"state": "RUNNING",
					"spec":  serverlessPipelineSpec,
				},
			},
		},
		Create:   true,
		Resource: ResourcePipeline(),
		HCL: `name = "test-pipeline-serverless"
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
			path = "/TestServerless"
		  }
		}
		filters {
		  include = ["com.databricks.include"]
		  exclude = ["com.databricks.exclude"]
		}
		continuous = false
		serverless = true
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "serverless", d.Id())
}

func TestZeroWorkers(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/pipelines",
				ExpectedRequest: PipelineSpec{
					Name:    "test-pipeline",
					Channel: "CURRENT",
					Edition: "ADVANCED",
					Clusters: []pipelineCluster{
						{
							Label:      "default",
							NumWorkers: 0,
							SparkConf: map[string]string{
								"spark.databricks.cluster.profile": "singleNode",
							},
							ForceSendFields: []string{"NumWorkers"},
						},
					},
				},
				Response: createPipelineResponse{
					PipelineID: "abcd",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: map[string]any{
					"id":    "abcd",
					"name":  "test-pipeline",
					"state": "RUNNING",
					"spec":  basicPipelineSpec,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: map[string]any{
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
		cluster {
		  label = "default"
		  num_workers = 0
		  spark_conf = {
			spark.databricks.cluster.profile = "singleNode"	
		  }
		}
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abcd", d.Id())
}

func TestAutoscaling(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/pipelines",
				ExpectedRequest: PipelineSpec{
					Name:    "test-pipeline",
					Channel: "CURRENT",
					Edition: "ADVANCED",
					Clusters: []pipelineCluster{
						{
							Label: "default",

							Autoscale: &dltAutoScale{
								MinWorkers: 2,
								MaxWorkers: 10,
							},
						},
					},
				},
				Response: createPipelineResponse{
					PipelineID: "abcd",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: map[string]any{
					"id":    "abcd",
					"name":  "test-pipeline",
					"state": "RUNNING",
					"spec":  basicPipelineSpec,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: map[string]any{
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
		cluster {
		  label = "default"
		  autoscale {
			min_workers = 2
			max_workers = 10
		  }
		}
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abcd", d.Id())
}

func TestDefault(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/pipelines",
				ExpectedRequest: PipelineSpec{
					Name:    "test-pipeline",
					Channel: "CURRENT",
					Edition: "ADVANCED",
					Clusters: []pipelineCluster{
						{
							Label: "default",
						},
					},
				},
				Response: createPipelineResponse{
					PipelineID: "abcd",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: map[string]any{
					"id":    "abcd",
					"name":  "test-pipeline",
					"state": "RUNNING",
					"spec":  basicPipelineSpec,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/abcd",
				Response: map[string]any{
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
		cluster {
		  label = "default"
		}
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abcd", d.Id())
}
