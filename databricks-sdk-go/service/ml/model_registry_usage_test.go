// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package ml_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/ml"
)

func ExampleModelRegistryAPI_CreateComment_modelVersionComments() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	model, err := w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", model)

	mv, err := w.ModelRegistry.CreateModelVersion(ctx, ml.CreateModelVersionRequest{
		Name:   model.RegisteredModel.Name,
		Source: "dbfs:/tmp",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", mv)

	created, err := w.ModelRegistry.CreateComment(ctx, ml.CreateComment{
		Comment: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Name:    mv.ModelVersion.Name,
		Version: mv.ModelVersion.Version,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = w.ModelRegistry.DeleteComment(ctx, ml.DeleteCommentRequest{
		Id: created.Comment.Id,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleModelRegistryAPI_CreateModel_models() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

}

func ExampleModelRegistryAPI_CreateModel_modelVersions() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	model, err := w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", model)

}

func ExampleModelRegistryAPI_CreateModel_modelVersionComments() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	model, err := w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", model)

}

func ExampleModelRegistryAPI_CreateModelVersion_modelVersions() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	model, err := w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", model)

	created, err := w.ModelRegistry.CreateModelVersion(ctx, ml.CreateModelVersionRequest{
		Name:   model.RegisteredModel.Name,
		Source: "dbfs:/tmp",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

}

func ExampleModelRegistryAPI_CreateModelVersion_modelVersionComments() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	model, err := w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", model)

	mv, err := w.ModelRegistry.CreateModelVersion(ctx, ml.CreateModelVersionRequest{
		Name:   model.RegisteredModel.Name,
		Source: "dbfs:/tmp",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", mv)

}

func ExampleModelRegistryAPI_CreateWebhook_registryWebhooks() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.ModelRegistry.CreateWebhook(ctx, ml.CreateRegistryWebhook{
		Description: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Events:      []ml.RegistryWebhookEvent{ml.RegistryWebhookEventModelVersionCreated},
		HttpUrlSpec: &ml.HttpUrlSpec{
			Url: w.Config.CanonicalHostName(),
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = w.ModelRegistry.DeleteWebhook(ctx, ml.DeleteWebhookRequest{
		Id: created.Webhook.Id,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleModelRegistryAPI_GetModel_models() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	model, err := w.ModelRegistry.GetModel(ctx, ml.GetModelRequest{
		Name: created.RegisteredModel.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", model)

}

func ExampleModelRegistryAPI_ListModels_models() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.ModelRegistry.ListModelsAll(ctx, ml.ListModelsRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleModelRegistryAPI_ListWebhooks_registryWebhooks() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.ModelRegistry.ListWebhooksAll(ctx, ml.ListWebhooksRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleModelRegistryAPI_UpdateComment_modelVersionComments() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	model, err := w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", model)

	mv, err := w.ModelRegistry.CreateModelVersion(ctx, ml.CreateModelVersionRequest{
		Name:   model.RegisteredModel.Name,
		Source: "dbfs:/tmp",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", mv)

	created, err := w.ModelRegistry.CreateComment(ctx, ml.CreateComment{
		Comment: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Name:    mv.ModelVersion.Name,
		Version: mv.ModelVersion.Version,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.ModelRegistry.UpdateComment(ctx, ml.UpdateComment{
		Comment: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Id:      created.Comment.Id,
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.ModelRegistry.DeleteComment(ctx, ml.DeleteCommentRequest{
		Id: created.Comment.Id,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleModelRegistryAPI_UpdateModel_models() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	model, err := w.ModelRegistry.GetModel(ctx, ml.GetModelRequest{
		Name: created.RegisteredModel.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", model)

	_, err = w.ModelRegistry.UpdateModel(ctx, ml.UpdateModelRequest{
		Name:        model.RegisteredModelDatabricks.Name,
		Description: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}

}

func ExampleModelRegistryAPI_UpdateModelVersion_modelVersions() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	model, err := w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", model)

	created, err := w.ModelRegistry.CreateModelVersion(ctx, ml.CreateModelVersionRequest{
		Name:   model.RegisteredModel.Name,
		Source: "dbfs:/tmp",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.ModelRegistry.UpdateModelVersion(ctx, ml.UpdateModelVersionRequest{
		Description: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Name:        created.ModelVersion.Name,
		Version:     created.ModelVersion.Version,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleModelRegistryAPI_UpdateWebhook_registryWebhooks() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.ModelRegistry.CreateWebhook(ctx, ml.CreateRegistryWebhook{
		Description: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Events:      []ml.RegistryWebhookEvent{ml.RegistryWebhookEventModelVersionCreated},
		HttpUrlSpec: &ml.HttpUrlSpec{
			Url: w.Config.CanonicalHostName(),
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.ModelRegistry.UpdateWebhook(ctx, ml.UpdateRegistryWebhook{
		Id:          created.Webhook.Id,
		Description: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.ModelRegistry.DeleteWebhook(ctx, ml.DeleteWebhookRequest{
		Id: created.Webhook.Id,
	})
	if err != nil {
		panic(err)
	}

}
