// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package workspace_test

import (
	"context"
	"encoding/base64"
	"fmt"
	"path/filepath"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/workspace"
)

func ExampleWorkspaceAPI_Export_workspaceIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebook := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	exportResponse, err := w.Workspace.Export(ctx, workspace.ExportRequest{
		Format: workspace.ExportFormatSource,
		Path:   notebook,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", exportResponse)

}

func ExampleWorkspaceAPI_GetStatus_genericPermissions() {
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

	obj, err := w.Workspace.GetStatusByPath(ctx, notebookPath)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", obj)

}

func ExampleWorkspaceAPI_GetStatus_workspaceIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebook := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	getStatusResponse, err := w.Workspace.GetStatusByPath(ctx, notebook)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", getStatusResponse)

}

func ExampleWorkspaceAPI_Import_jobsApiFullIntegration() {
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

	err = w.Workspace.Import(ctx, workspace.Import{
		Path:      notebookPath,
		Overwrite: true,
		Format:    workspace.ImportFormatSource,
		Language:  workspace.LanguagePython,
		Content: base64.StdEncoding.EncodeToString([]byte((`import time
time.sleep(10)
dbutils.notebook.exit('hello')
`))),
	})
	if err != nil {
		panic(err)
	}

}

func ExampleWorkspaceAPI_Import_genericPermissions() {
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

	err = w.Workspace.Import(ctx, workspace.Import{
		Path:      notebookPath,
		Overwrite: true,
		Format:    workspace.ImportFormatSource,
		Language:  workspace.LanguagePython,
		Content: base64.StdEncoding.EncodeToString([]byte((`print(1)
`))),
	})
	if err != nil {
		panic(err)
	}

}

func ExampleWorkspaceAPI_Import_pipelines() {
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

	err = w.Workspace.Import(ctx, workspace.Import{
		Content:   base64.StdEncoding.EncodeToString([]byte(("CREATE LIVE TABLE dlt_sample AS SELECT 1"))),
		Format:    workspace.ImportFormatSource,
		Language:  workspace.LanguageSql,
		Overwrite: true,
		Path:      notebookPath,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleWorkspaceAPI_Import_workspaceIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebook := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	err = w.Workspace.Import(ctx, workspace.Import{
		Path:      notebook,
		Format:    workspace.ImportFormatSource,
		Language:  workspace.LanguagePython,
		Content:   base64.StdEncoding.EncodeToString([]byte(("# Databricks notebook source\nprint('hello from job')"))),
		Overwrite: true,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleWorkspaceAPI_List_workspaceIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebook := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	objects, err := w.Workspace.ListAll(ctx, workspace.ListWorkspaceRequest{
		Path: filepath.Dir(notebook),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", objects)

}
