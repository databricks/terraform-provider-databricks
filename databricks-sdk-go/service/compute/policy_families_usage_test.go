// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package compute_test

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/compute"
)

func ExamplePolicyFamiliesAPI_Get_clusterPolicyFamilies() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.PolicyFamilies.ListAll(ctx, compute.ListPolicyFamiliesRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

	firstFamily, err := w.PolicyFamilies.Get(ctx, compute.GetPolicyFamilyRequest{
		PolicyFamilyId: all[0].PolicyFamilyId,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", firstFamily)

}

func ExamplePolicyFamiliesAPI_List_clusterPolicyFamilies() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.PolicyFamilies.ListAll(ctx, compute.ListPolicyFamiliesRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}
