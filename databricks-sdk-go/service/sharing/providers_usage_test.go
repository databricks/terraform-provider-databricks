// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package sharing_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/sharing"
)

func ExampleProvidersAPI_Create_providers() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	publicShareRecipient := `{
        "shareCredentialsVersion":1,
        "bearerToken":"dapiabcdefghijklmonpqrstuvwxyz",
        "endpoint":"https://sharing.delta.io/delta-sharing/"
    }
`

	created, err := w.Providers.Create(ctx, sharing.CreateProvider{
		Name:                fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		RecipientProfileStr: publicShareRecipient,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = w.Providers.DeleteByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

}

func ExampleProvidersAPI_Get_providers() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	publicShareRecipient := `{
        "shareCredentialsVersion":1,
        "bearerToken":"dapiabcdefghijklmonpqrstuvwxyz",
        "endpoint":"https://sharing.delta.io/delta-sharing/"
    }
`

	created, err := w.Providers.Create(ctx, sharing.CreateProvider{
		Name:                fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		RecipientProfileStr: publicShareRecipient,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.Providers.GetByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Providers.DeleteByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

}

func ExampleProvidersAPI_List_providers() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.Providers.ListAll(ctx, sharing.ListProvidersRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleProvidersAPI_ListShares_providers() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	publicShareRecipient := `{
        "shareCredentialsVersion":1,
        "bearerToken":"dapiabcdefghijklmonpqrstuvwxyz",
        "endpoint":"https://sharing.delta.io/delta-sharing/"
    }
`

	created, err := w.Providers.Create(ctx, sharing.CreateProvider{
		Name:                fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		RecipientProfileStr: publicShareRecipient,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	shares, err := w.Providers.ListSharesAll(ctx, sharing.ListSharesRequest{
		Name: created.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", shares)

	// cleanup

	err = w.Providers.DeleteByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

}

func ExampleProvidersAPI_Update_providers() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	publicShareRecipient := `{
        "shareCredentialsVersion":1,
        "bearerToken":"dapiabcdefghijklmonpqrstuvwxyz",
        "endpoint":"https://sharing.delta.io/delta-sharing/"
    }
`

	created, err := w.Providers.Create(ctx, sharing.CreateProvider{
		Name:                fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		RecipientProfileStr: publicShareRecipient,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.Providers.Update(ctx, sharing.UpdateProvider{
		Name:    created.Name,
		Comment: "Comment for update",
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Providers.DeleteByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

}
