package pipelines

import (
	// "context"
	"errors"
	// "fmt"
	"testing"

	// "github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/stretchr/testify/mock"
)

var createRequest = pipelines.CreatePipeline{
	Name:    "test-pipeline",
	Storage: "/test/storage",
	Configuration: map[string]string{
		"key1": "value1",
		"key2": "value2",
	},
	Clusters: []pipelines.PipelineCluster{
		{
			Label: "default",
			CustomTags: map[string]string{
				"cluster_tag1": "cluster_value1",
			},
		},
	},
	Libraries: []pipelines.PipelineLibrary{
		{
			Notebook: &pipelines.NotebookLibrary{
				Path: "/Test",
			},
		},
	},
	Filters: &pipelines.Filters{
		Include: []string{"com.databricks.include"},
		Exclude: []string{"com.databricks.exclude"},
	},
	Deployment: &pipelines.PipelineDeployment{
		Kind:             "BUNDLE",
		MetadataFilePath: "/foo/bar",
	},
	Edition: "ADVANCED",
	Channel: "CURRENT",
}

var updateRequest = pipelines.EditPipeline{
	Id:      "abcd",
	Name:    "test",
	Storage: "/test/storage",
	Libraries: []pipelines.PipelineLibrary{
		{
			Notebook: &pipelines.NotebookLibrary{
				Path: "/Test",
			},
		},
	},
	Filters: &pipelines.Filters{
		Include: []string{"com.databricks.include"},
	},
	Channel: "CURRENT",
	Edition: "ADVANCED",
}

var basicPipelineSpec = pipelines.PipelineSpec{
	Name:    "test-pipeline",
	Storage: "/test/storage",
	Configuration: map[string]string{
		"key1": "value1",
		"key2": "value2",
	},
	Clusters: []pipelines.PipelineCluster{
		{
			Label: "default",
			CustomTags: map[string]string{
				"cluster_tag1": "cluster_value1",
			},
		},
	},
	Libraries: []pipelines.PipelineLibrary{
		{
			Notebook: &pipelines.NotebookLibrary{
				Path: "/Test",
			},
		},
	},
	Filters: &pipelines.Filters{
		Include: []string{"com.databricks.include"},
		Exclude: []string{"com.databricks.exclude"},
	},
	Deployment: &pipelines.PipelineDeployment{
		Kind:             "BUNDLE",
		MetadataFilePath: "/foo/bar",
	},
	Edition: "ADVANCED",
	Channel: "CURRENT",
}

func TestResourcePipelineCreate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockPipelinesAPI().EXPECT()
			e.Create(mock.Anything, createRequest).Return(&pipelines.CreatePipelineResponse{
				PipelineId: "abcd",
			}, nil)
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "abcd",
			}).Return(&pipelines.GetPipelineResponse{
				PipelineId: "abcd",
				Name:       "test-pipeline",
				State:      StateDeploying,
				Spec:       &basicPipelineSpec,
			}, nil).Once()
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "abcd",
			}).Return(&pipelines.GetPipelineResponse{
				PipelineId: "abcd",
				Name:       "test-pipeline",
				State:      StateRunning,
				Spec:       &basicPipelineSpec,
			}, nil).Once()

		},
		Resource: ResourcePipeline(),
		Create:   true,
		HCL: `
			name = "test-pipeline"
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
	}.ApplyAndExpectData(t, map[string]any{
		"id": "abcd",
	})
}

func TestResourcePipelineCreate_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockPipelinesAPI().EXPECT()
			e.Create(mock.Anything, mock.Anything).Return(nil, errors.New("Internal error happened"))
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
	}.ExpectError(t, "Internal error happened")
}

func TestResourcePipelineCreate_ErrorWhenWaitingFailedCleanup(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockPipelinesAPI().EXPECT()
			e.Create(mock.Anything, mock.Anything).Return(&pipelines.CreatePipelineResponse{
				PipelineId: "abcd",
			}, nil)
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "abcd",
			}).Return(&pipelines.GetPipelineResponse{
				PipelineId: "abcd",
				Name:       "test-pipeline",
				State:      StateFailed,
			}, nil).Once()
			e.Delete(mock.Anything, pipelines.DeletePipelineRequest{
				PipelineId: "abcd",
			}).Return(errors.New("Internal error"))
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "abcd",
			}).Return(nil, errors.New("Internal error"))
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
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockPipelinesAPI().EXPECT()
			e.Create(mock.Anything, mock.Anything).Return(&pipelines.CreatePipelineResponse{
				PipelineId: "abcd",
			}, nil)

			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "abcd",
			}).Return(&pipelines.GetPipelineResponse{
				PipelineId: "abcd",
				Name:       "test-pipeline",
				State:      StateFailed,
			}, nil).Once()

			e.Delete(mock.Anything, pipelines.DeletePipelineRequest{
				PipelineId: "abcd",
			}).Return(nil)

			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "abcd",
			}).Return(nil, apierr.ErrNotFound)
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
	}.ExpectError(t, "pipeline abcd has failed")
}

func TestResourcePipelineRead(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockPipelinesAPI().EXPECT()
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "abcd",
			}).Return(&pipelines.GetPipelineResponse{
				PipelineId: "abcd",
				Spec:       &basicPipelineSpec,
			}, nil)
		},
		Resource: ResourcePipeline(),
		Read:     true,
		New:      true,
		ID:       "abcd",
	}.ApplyAndExpectData(t, map[string]any{
		"id":      "abcd",
		"storage": "/test/storage",
		"configuration": map[string]any{
			"key1": "value1",
			"key2": "value2",
		},
		"filters.0.include.0": "com.databricks.include",
		"continuous":          false,
	})
}

func TestResourcePipelineRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockPipelinesAPI().EXPECT()
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "abcd",
			}).Return(nil, apierr.ErrNotFound)
		},
		Resource: ResourcePipeline(),
		Read:     true,
		Removed:  true,
		ID:       "abcd",
	}.ApplyNoError(t)
}

func TestResourcePipelineRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockPipelinesAPI().EXPECT()
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "abcd",
			}).Return(nil, errors.New("Internal error happened"))
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
	spec := pipelines.PipelineSpec{
		Id:      "abcd",
		Name:    "test",
		Storage: "/test/storage",
		Libraries: []pipelines.PipelineLibrary{
			{
				Notebook: &pipelines.NotebookLibrary{
					Path: "/Test",
				},
			},
		},
		Filters: &pipelines.Filters{
			Include: []string{"com.databricks.include"},
		},
		Channel: "CURRENT",
		Edition: "ADVANCED",
	}
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockPipelinesAPI().EXPECT()
			e.Update(mock.Anything, updateRequest).Return(nil)
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "abcd",
			}).Return(&pipelines.GetPipelineResponse{
				PipelineId: "abcd",
				Spec:       &spec,
				State:      state,
			}, nil).Once()
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "abcd",
			}).Return(&pipelines.GetPipelineResponse{
				PipelineId: "abcd",
				Spec:       &spec,
				State:      state,
			}, nil)
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
	}.ApplyAndExpectData(t, map[string]any{
		"id": "abcd",
	})
}

func TestResourcePipelineUpdate_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockPipelinesAPI().EXPECT()
			e.Update(mock.Anything, mock.Anything).Return(errors.New("Internal error happened"))
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
	}.ExpectError(t, "Internal error happened")
}

func TestResourcePipelineUpdate_FailsAfterUpdate(t *testing.T) {
	state := StateFailed
	spec := pipelines.PipelineSpec{
		Id:      "abcd",
		Name:    "test",
		Storage: "/test/storage",
		Libraries: []pipelines.PipelineLibrary{
			{
				Notebook: &pipelines.NotebookLibrary{
					Path: "/Test",
				},
			},
		},
		Filters: &pipelines.Filters{
			Include: []string{"com.databricks.include"},
		},
		Channel: "CURRENT",
		Edition: "ADVANCED",
	}
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockPipelinesAPI().EXPECT()
			e.Update(mock.Anything, updateRequest).Return(nil)
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "abcd",
			}).Return(&pipelines.GetPipelineResponse{
				PipelineId: "abcd",
				Spec:       &spec,
				State:      state,
			}, nil)
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
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockPipelinesAPI().EXPECT()
			e.Delete(mock.Anything, pipelines.DeletePipelineRequest{
				PipelineId: "abcd",
			}).Return(nil)
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "abcd",
			}).Return(&pipelines.GetPipelineResponse{
				PipelineId: "abcd",
				Spec:       &basicPipelineSpec,
				State:      state,
			}, nil).Once()
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "abcd",
			}).Return(nil, apierr.ErrNotFound)
		},
		Resource: ResourcePipeline(),
		Delete:   true,
		ID:       "abcd",
	}.ApplyAndExpectData(t, map[string]any{
		"id": "abcd",
	})
}

func TestResourcePipelineDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockPipelinesAPI().EXPECT()
			e.Delete(mock.Anything, mock.Anything).Return(errors.New("Internal error happened"))
		},
		Resource: ResourcePipeline(),
		Delete:   true,
		ID:       "abcd",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abcd", d.Id())
}

// type pipelineListResponse struct {
// 	Statuses      []pipelines.PipelineStateInfo `json:"statuses"`
// 	NextPageToken string                        `json:"next_page_token"`
// 	PrevPageToken string                        `json:"prev_page_token"`
// }

// func listPipelines(ctx context.Context, client *common.DatabricksClient, maxResults int, filter string) ([]pipelines.PipelineStateInfo, error) {
// 	w, err2 := client.WorkspaceClient()
// 	if err2 != nil {
// 		return nil, err2
// 	}

// 	data, err := w.Pipelines.ListPipelinesAll(ctx, pipelines.ListPipelinesRequest{
// 		MaxResults: maxResults,
// 		Filter:     filter,
// 	})
// 	fmt.Println(data)
// 	return data, err
// }

// // Giving error (don't know how to resolve)
// func TestListPipelines(t *testing.T) {
// 	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
// 		{
// 			Method:   "GET",
// 			Resource: "/api/2.0/pipelines?max_results=1",
// 			Response: pipelineListResponse{
// 				Statuses: []pipelines.PipelineStateInfo{
// 					{
// 						PipelineId:      "123",
// 						Name:            "Pipeline1",
// 						CreatorUserName: "user1",
// 					},
// 				},
// 				NextPageToken: "token1",
// 			},
// 		},
// 		{
// 			Method:   "GET",
// 			Resource: "/api/2.0/pipelines?max_results=1&page_token=token1",
// 			Response: pipelineListResponse{
// 				Statuses: []pipelines.PipelineStateInfo{
// 					{
// 						PipelineId:      "456",
// 						Name:            "Pipeline2",
// 						CreatorUserName: "user2",
// 					},
// 				},
// 				PrevPageToken: "token0",
// 			},
// 		},
// 		// {
// 		// 	Method:   "GET",
// 		// 	Resource: "/api/2.0/pipelines?max_results=2",
// 		// 	Response: pipelineListResponse{
// 		// 		Statuses: []pipelines.PipelineStateInfo{
// 		// 			{
// 		// 				PipelineId:      "123",
// 		// 				Name:            "Pipeline1",
// 		// 				CreatorUserName: "user1",
// 		// 			},
// 		// 			{
// 		// 				PipelineId:      "456",
// 		// 				Name:            "Pipeline2",
// 		// 				CreatorUserName: "user2",
// 		// 			},
// 		// 		},
// 		// 	},
// 		// },
// 	})
// 	defer server.Close()
// 	require.NoError(t, err)

// 	ctx := context.Background()
// 	data, err := listPipelines(ctx, client, 1, "")
// 	require.NoError(t, err)
// 	require.Equal(t, 2, len(data))
// 	require.Equal(t, "Pipeline1", data[0].Name)
// 	require.Equal(t, "456", data[1].PipelineId)
// }

// // Not sure about this test
// func TestListPipelinesWithFilter(t *testing.T) {
// 	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
// 		{
// 			Method:   "GET",
// 			Resource: "/api/2.0/pipelines?filter=name+LIKE+%27Pipeline1%27&max_results=1",
// 			Response: pipelineListResponse{
// 				Statuses: []pipelines.PipelineStateInfo{
// 					{
// 						PipelineId:      "123",
// 						Name:            "Pipeline1",
// 						CreatorUserName: "user1",
// 					},
// 				},
// 			},
// 		},
// 	})
// 	defer server.Close()
// 	require.NoError(t, err)

// 	ctx := context.Background()
// 	data, err := listPipelines(ctx, client, 1, "name LIKE 'Pipeline1'")
// 	// data, err := NewPipelinesAPI(ctx, client).List(1, "name LIKE 'Pipeline1'")
// 	require.NoError(t, err)
// 	require.Equal(t, 1, len(data))
// }

func TestStorageSuppressDiff(t *testing.T) {
	k := "storage"
	generated := "dbfs:/pipelines/c609bbb0-2e42-4bc8-bb4e-a1c26d6e9403"
	require.True(t, suppressStorageDiff(k, generated, "", nil))
	require.False(t, suppressStorageDiff(k, generated, "/tmp/abc", nil))
	require.False(t, suppressStorageDiff(k, "/tmp/abc", "", nil))
}

func TestResourcePipelineCreateServerless(t *testing.T) {
	var serverlessPipelineSpec = pipelines.PipelineSpec{
		Name:    "test-pipeline-serverless",
		Storage: "/test/storage",
		Configuration: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
		Clusters: []pipelines.PipelineCluster{
			{
				Label: "default",
				CustomTags: map[string]string{
					"cluster_tag1": "cluster_value1",
				},
			},
		},
		Libraries: []pipelines.PipelineLibrary{
			{
				Notebook: &pipelines.NotebookLibrary{
					Path: "/TestServerless",
				},
			},
		},
		Filters: &pipelines.Filters{
			Include: []string{"com.databricks.include"},
			Exclude: []string{"com.databricks.exclude"},
		},
		Serverless: true,
	}
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockPipelinesAPI().EXPECT()
			e.Create(mock.Anything, mock.Anything).Return(&pipelines.CreatePipelineResponse{
				PipelineId: "serverless",
			}, nil)
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "serverless",
			}).Return(&pipelines.GetPipelineResponse{
				PipelineId: "serverless",
				Name:       "test-pipeline-serverless",
				State:      StateDeploying,
				Spec:       &serverlessPipelineSpec,
			}, nil).Once()
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "serverless",
			}).Return(&pipelines.GetPipelineResponse{
				PipelineId: "serverless",
				Name:       "test-pipeline-serverless",
				State:      StateRunning,
				Spec:       &serverlessPipelineSpec,
			}, nil).Once()
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "serverless",
			}).Return(&pipelines.GetPipelineResponse{
				PipelineId: "serverless",
				Name:       "test-pipeline-serverless",
				State:      StateRunning,
				Spec:       &serverlessPipelineSpec,
			}, nil)
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
	}.ApplyAndExpectData(t, map[string]any{
		"id": "serverless",
	})
}

func TestZeroWorkers(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockPipelinesAPI().EXPECT()
			e.Create(mock.Anything, pipelines.CreatePipeline{
				Name:    "test-pipeline",
				Channel: "CURRENT",
				Edition: "ADVANCED",
				Clusters: []pipelines.PipelineCluster{
					{
						Label:      "default",
						NumWorkers: 0,
						SparkConf: map[string]string{
							"spark.databricks.cluster.profile": "singleNode",
						},
						ForceSendFields: []string{"NumWorkers"},
					},
				},
			}).Return(&pipelines.CreatePipelineResponse{
				PipelineId: "abcd",
			}, nil)
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "abcd",
			}).Return(&pipelines.GetPipelineResponse{
				PipelineId: "abcd",
				Name:       "test-pipeline",
				State:      StateRunning,
				Spec:       &basicPipelineSpec,
			}, nil).Once()
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "abcd",
			}).Return(&pipelines.GetPipelineResponse{
				PipelineId: "abcd",
				Name:       "test-pipeline",
				State:      StateRunning,
				Spec:       &basicPipelineSpec,
			}, nil).Once()
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
	}.ApplyAndExpectData(t, map[string]any{
		"id": "abcd",
	})
}

func TestAutoscaling(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockPipelinesAPI().EXPECT()
			e.Create(mock.Anything, pipelines.CreatePipeline{
				Name:    "test-pipeline",
				Channel: "CURRENT",
				Edition: "ADVANCED",
				Clusters: []pipelines.PipelineCluster{
					{
						Label: "default",
						Autoscale: &pipelines.PipelineClusterAutoscale{
							MinWorkers: 2,
							MaxWorkers: 10,
						},
					},
				},
			}).Return(&pipelines.CreatePipelineResponse{
				PipelineId: "abcd",
			}, nil)
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "abcd",
			}).Return(&pipelines.GetPipelineResponse{
				PipelineId: "abcd",
				Name:       "test-pipeline",
				State:      StateRunning,
				Spec:       &basicPipelineSpec,
			}, nil).Once()
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "abcd",
			}).Return(&pipelines.GetPipelineResponse{
				PipelineId: "abcd",
				Name:       "test-pipeline",
				State:      StateRunning,
				Spec:       &basicPipelineSpec,
			}, nil).Once()
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
	}.ApplyAndExpectData(t, map[string]any{
		"id": "abcd",
	})
}

func TestDefault(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockPipelinesAPI().EXPECT()
			e.Create(mock.Anything, pipelines.CreatePipeline{
				Name:    "test-pipeline",
				Channel: "CURRENT",
				Edition: "ADVANCED",
				Clusters: []pipelines.PipelineCluster{
					{
						Label: "default",
					},
				},
			}).Return(&pipelines.CreatePipelineResponse{
				PipelineId: "abcd",
			}, nil)
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "abcd",
			}).Return(&pipelines.GetPipelineResponse{
				PipelineId: "abcd",
				Name:       "test-pipeline",
				State:      StateRunning,
				Spec:       &basicPipelineSpec,
			}, nil).Once()
			e.Get(mock.Anything, pipelines.GetPipelineRequest{
				PipelineId: "abcd",
			}).Return(&pipelines.GetPipelineResponse{
				PipelineId: "abcd",
				Name:       "test-pipeline",
				State:      StateRunning,
				Spec:       &basicPipelineSpec,
			}, nil).Once()
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
