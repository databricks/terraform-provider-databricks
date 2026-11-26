// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package compute_test

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/compute"
)

func ExampleGlobalInitScriptsAPI_Create_globalInitScripts() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.GlobalInitScripts.Create(ctx, compute.GlobalInitScriptCreateRequest{
		Name:     fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Script:   base64.StdEncoding.EncodeToString([]byte(("echo 1"))),
		Enabled:  true,
		Position: 10,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = w.GlobalInitScripts.DeleteByScriptId(ctx, created.ScriptId)
	if err != nil {
		panic(err)
	}

}

func ExampleGlobalInitScriptsAPI_Get_globalInitScripts() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.GlobalInitScripts.Create(ctx, compute.GlobalInitScriptCreateRequest{
		Name:     fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Script:   base64.StdEncoding.EncodeToString([]byte(("echo 1"))),
		Enabled:  true,
		Position: 10,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	byId, err := w.GlobalInitScripts.GetByScriptId(ctx, created.ScriptId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	err = w.GlobalInitScripts.DeleteByScriptId(ctx, created.ScriptId)
	if err != nil {
		panic(err)
	}

}

func ExampleGlobalInitScriptsAPI_List_globalInitScripts() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.GlobalInitScripts.ListAll(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleGlobalInitScriptsAPI_Update_globalInitScripts() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.GlobalInitScripts.Create(ctx, compute.GlobalInitScriptCreateRequest{
		Name:     fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Script:   base64.StdEncoding.EncodeToString([]byte(("echo 1"))),
		Enabled:  true,
		Position: 10,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	err = w.GlobalInitScripts.Update(ctx, compute.GlobalInitScriptUpdateRequest{
		ScriptId: created.ScriptId,
		Name:     fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Script:   base64.StdEncoding.EncodeToString([]byte(("echo 2"))),
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.GlobalInitScripts.DeleteByScriptId(ctx, created.ScriptId)
	if err != nil {
		panic(err)
	}

}
