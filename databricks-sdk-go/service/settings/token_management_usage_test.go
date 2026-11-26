// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package settings_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/settings"
)

func ExampleTokenManagementAPI_CreateOboToken_createOboTokenOnAws() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	groups, err := w.Groups.GroupDisplayNameToIdMap(ctx, iam.ListGroupsRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", groups)

	spn, err := w.ServicePrincipals.Create(ctx, iam.ServicePrincipal{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Groups: []iam.ComplexValue{{
			Value: groups["admins"],
		}},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", spn)

	obo, err := w.TokenManagement.CreateOboToken(ctx, settings.CreateOboTokenRequest{
		ApplicationId:   spn.ApplicationId,
		LifetimeSeconds: 60,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", obo)

	// cleanup

	err = w.ServicePrincipals.DeleteById(ctx, spn.Id)
	if err != nil {
		panic(err)
	}
	err = w.TokenManagement.DeleteByTokenId(ctx, obo.TokenInfo.TokenId)
	if err != nil {
		panic(err)
	}

}

func ExampleTokenManagementAPI_Get_createOboTokenOnAws() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	groups, err := w.Groups.GroupDisplayNameToIdMap(ctx, iam.ListGroupsRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", groups)

	spn, err := w.ServicePrincipals.Create(ctx, iam.ServicePrincipal{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Groups: []iam.ComplexValue{{
			Value: groups["admins"],
		}},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", spn)

	obo, err := w.TokenManagement.CreateOboToken(ctx, settings.CreateOboTokenRequest{
		ApplicationId:   spn.ApplicationId,
		LifetimeSeconds: 60,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", obo)

	byId, err := w.TokenManagement.GetByTokenId(ctx, obo.TokenInfo.TokenId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	err = w.ServicePrincipals.DeleteById(ctx, spn.Id)
	if err != nil {
		panic(err)
	}
	err = w.TokenManagement.DeleteByTokenId(ctx, obo.TokenInfo.TokenId)
	if err != nil {
		panic(err)
	}

}

func ExampleTokenManagementAPI_List_createOboTokenOnAws() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.TokenManagement.ListAll(ctx, settings.ListTokenManagementRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}
