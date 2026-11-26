// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package workspace_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/workspace"
)

func ExampleSecretsAPI_CreateScope_secrets() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	keyName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	scopeName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	err = w.Secrets.CreateScope(ctx, workspace.CreateScope{
		Scope: scopeName,
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Secrets.DeleteSecret(ctx, workspace.DeleteSecret{
		Scope: scopeName,
		Key:   keyName,
	})
	if err != nil {
		panic(err)
	}
	err = w.Secrets.DeleteScopeByScope(ctx, scopeName)
	if err != nil {
		panic(err)
	}

}

func ExampleSecretsAPI_ListAcls_secrets() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	keyName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	scopeName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	err = w.Secrets.CreateScope(ctx, workspace.CreateScope{
		Scope: scopeName,
	})
	if err != nil {
		panic(err)
	}

	acls, err := w.Secrets.ListAclsByScope(ctx, scopeName)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", acls)

	// cleanup

	err = w.Secrets.DeleteSecret(ctx, workspace.DeleteSecret{
		Scope: scopeName,
		Key:   keyName,
	})
	if err != nil {
		panic(err)
	}
	err = w.Secrets.DeleteScopeByScope(ctx, scopeName)
	if err != nil {
		panic(err)
	}

}

func ExampleSecretsAPI_ListScopes_secrets() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	scopes, err := w.Secrets.ListScopesAll(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", scopes)

}

func ExampleSecretsAPI_ListSecrets_secrets() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	keyName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	scopeName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	err = w.Secrets.CreateScope(ctx, workspace.CreateScope{
		Scope: scopeName,
	})
	if err != nil {
		panic(err)
	}

	scrts, err := w.Secrets.ListSecretsByScope(ctx, scopeName)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", scrts)

	// cleanup

	err = w.Secrets.DeleteSecret(ctx, workspace.DeleteSecret{
		Scope: scopeName,
		Key:   keyName,
	})
	if err != nil {
		panic(err)
	}
	err = w.Secrets.DeleteScopeByScope(ctx, scopeName)
	if err != nil {
		panic(err)
	}

}

func ExampleSecretsAPI_PutAcl_secrets() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	keyName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	group, err := w.Groups.Create(ctx, iam.Group{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", group)

	scopeName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	err = w.Secrets.CreateScope(ctx, workspace.CreateScope{
		Scope: scopeName,
	})
	if err != nil {
		panic(err)
	}

	err = w.Secrets.PutAcl(ctx, workspace.PutAcl{
		Scope:      scopeName,
		Permission: workspace.AclPermissionManage,
		Principal:  group.DisplayName,
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Groups.DeleteById(ctx, group.Id)
	if err != nil {
		panic(err)
	}
	err = w.Secrets.DeleteSecret(ctx, workspace.DeleteSecret{
		Scope: scopeName,
		Key:   keyName,
	})
	if err != nil {
		panic(err)
	}
	err = w.Secrets.DeleteScopeByScope(ctx, scopeName)
	if err != nil {
		panic(err)
	}

}

func ExampleSecretsAPI_PutSecret_secrets() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	keyName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	scopeName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	err = w.Secrets.CreateScope(ctx, workspace.CreateScope{
		Scope: scopeName,
	})
	if err != nil {
		panic(err)
	}

	err = w.Secrets.PutSecret(ctx, workspace.PutSecret{
		Scope:       scopeName,
		Key:         keyName,
		StringValue: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Secrets.DeleteSecret(ctx, workspace.DeleteSecret{
		Scope: scopeName,
		Key:   keyName,
	})
	if err != nil {
		panic(err)
	}
	err = w.Secrets.DeleteScopeByScope(ctx, scopeName)
	if err != nil {
		panic(err)
	}

}
