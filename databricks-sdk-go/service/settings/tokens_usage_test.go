// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package settings_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/settings"
)

func ExampleTokensAPI_Create_tokens() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	token, err := w.Tokens.Create(ctx, settings.CreateTokenRequest{
		Comment:         fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		LifetimeSeconds: 300,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", token)

	// cleanup

	err = w.Tokens.DeleteByTokenId(ctx, token.TokenInfo.TokenId)
	if err != nil {
		panic(err)
	}

}

func ExampleTokensAPI_GetByComment_tokens() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	token, err := w.Tokens.Create(ctx, settings.CreateTokenRequest{
		Comment:         fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		LifetimeSeconds: 300,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", token)

	byName, err := w.Tokens.GetByComment(ctx, token.TokenInfo.Comment)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byName)

	// cleanup

	err = w.Tokens.DeleteByTokenId(ctx, token.TokenInfo.TokenId)
	if err != nil {
		panic(err)
	}

}

func ExampleTokensAPI_List_tokens() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.Tokens.ListAll(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}
