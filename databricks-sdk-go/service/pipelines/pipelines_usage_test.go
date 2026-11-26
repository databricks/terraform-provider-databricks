// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package pipelines_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
)

func ExamplePipelinesAPI_Create_pipelines() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebookPath := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	created, err := w.Pipelines.Create(ctx, pipelines.CreatePipeline{
		Continuous: false,
		Name:       fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Libraries: []pipelines.PipelineLibrary{{
			Notebook: &pipelines.NotebookLibrary{
				Path: notebookPath,
			},
		}},
		Clusters: []pipelines.PipelineCluster{{
			InstancePoolId: os.Getenv("TEST_INSTANCE_POOL_ID"),
			Label:          "default",
			NumWorkers:     1,
			CustomTags:     map[string]string{"cluster_type": "default"},
		}},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = w.Pipelines.DeleteByPipelineId(ctx, created.PipelineId)
	if err != nil {
		panic(err)
	}

}

func ExamplePipelinesAPI_Get_pipelines() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebookPath := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	created, err := w.Pipelines.Create(ctx, pipelines.CreatePipeline{
		Continuous: false,
		Name:       fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Libraries: []pipelines.PipelineLibrary{{
			Notebook: &pipelines.NotebookLibrary{
				Path: notebookPath,
			},
		}},
		Clusters: []pipelines.PipelineCluster{{
			InstancePoolId: os.Getenv("TEST_INSTANCE_POOL_ID"),
			Label:          "default",
			NumWorkers:     1,
			CustomTags:     map[string]string{"cluster_type": "default"},
		}},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	byId, err := w.Pipelines.GetByPipelineId(ctx, created.PipelineId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	err = w.Pipelines.DeleteByPipelineId(ctx, created.PipelineId)
	if err != nil {
		panic(err)
	}

}

func ExamplePipelinesAPI_ListPipelineEvents_pipelines() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebookPath := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	created, err := w.Pipelines.Create(ctx, pipelines.CreatePipeline{
		Continuous: false,
		Name:       fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Libraries: []pipelines.PipelineLibrary{{
			Notebook: &pipelines.NotebookLibrary{
				Path: notebookPath,
			},
		}},
		Clusters: []pipelines.PipelineCluster{{
			InstancePoolId: os.Getenv("TEST_INSTANCE_POOL_ID"),
			Label:          "default",
			NumWorkers:     1,
			CustomTags:     map[string]string{"cluster_type": "default"},
		}},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	events, err := w.Pipelines.ListPipelineEventsAll(ctx, pipelines.ListPipelineEventsRequest{
		PipelineId: created.PipelineId,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", events)

	// cleanup

	err = w.Pipelines.DeleteByPipelineId(ctx, created.PipelineId)
	if err != nil {
		panic(err)
	}

}

func ExamplePipelinesAPI_ListPipelines_pipelines() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.Pipelines.ListPipelinesAll(ctx, pipelines.ListPipelinesRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExamplePipelinesAPI_Update_pipelines() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebookPath := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	created, err := w.Pipelines.Create(ctx, pipelines.CreatePipeline{
		Continuous: false,
		Name:       fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Libraries: []pipelines.PipelineLibrary{{
			Notebook: &pipelines.NotebookLibrary{
				Path: notebookPath,
			},
		}},
		Clusters: []pipelines.PipelineCluster{{
			InstancePoolId: os.Getenv("TEST_INSTANCE_POOL_ID"),
			Label:          "default",
			NumWorkers:     1,
			CustomTags:     map[string]string{"cluster_type": "default"},
		}},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	err = w.Pipelines.Update(ctx, pipelines.EditPipeline{
		PipelineId: created.PipelineId,
		Name:       fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Libraries: []pipelines.PipelineLibrary{{
			Notebook: &pipelines.NotebookLibrary{
				Path: notebookPath,
			},
		}},
		Clusters: []pipelines.PipelineCluster{{
			InstancePoolId: os.Getenv("TEST_INSTANCE_POOL_ID"),
			Label:          "default",
			NumWorkers:     1,
			CustomTags:     map[string]string{"cluster_type": "default"},
		}},
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Pipelines.DeleteByPipelineId(ctx, created.PipelineId)
	if err != nil {
		panic(err)
	}

}
