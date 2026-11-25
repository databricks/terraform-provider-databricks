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

func ExampleIpAccessListsAPI_Create_ipAccessLists() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.IpAccessLists.Create(ctx, settings.CreateIpAccessList{
		Label:       fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		IpAddresses: []string{"1.0.0.0/16"},
		ListType:    settings.ListTypeBlock,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = w.IpAccessLists.DeleteByIpAccessListId(ctx, created.IpAccessList.ListId)
	if err != nil {
		panic(err)
	}

}

func ExampleIpAccessListsAPI_Get_ipAccessLists() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.IpAccessLists.Create(ctx, settings.CreateIpAccessList{
		Label:       fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		IpAddresses: []string{"1.0.0.0/16"},
		ListType:    settings.ListTypeBlock,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	byId, err := w.IpAccessLists.GetByIpAccessListId(ctx, created.IpAccessList.ListId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	err = w.IpAccessLists.DeleteByIpAccessListId(ctx, created.IpAccessList.ListId)
	if err != nil {
		panic(err)
	}

}

func ExampleIpAccessListsAPI_List_ipAccessLists() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.IpAccessLists.ListAll(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleIpAccessListsAPI_Replace_ipAccessLists() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.IpAccessLists.Create(ctx, settings.CreateIpAccessList{
		Label:       fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		IpAddresses: []string{"1.0.0.0/16"},
		ListType:    settings.ListTypeBlock,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	err = w.IpAccessLists.Replace(ctx, settings.ReplaceIpAccessList{
		IpAccessListId: created.IpAccessList.ListId,
		Label:          fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		IpAddresses:    []string{"1.0.0.0/24"},
		ListType:       settings.ListTypeBlock,
		Enabled:        false,
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.IpAccessLists.DeleteByIpAccessListId(ctx, created.IpAccessList.ListId)
	if err != nil {
		panic(err)
	}

}
