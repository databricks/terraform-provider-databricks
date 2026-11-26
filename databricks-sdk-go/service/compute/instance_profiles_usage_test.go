// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package compute_test

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/compute"
)

func ExampleInstanceProfilesAPI_Add_awsInstanceProfiles() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	arn := "arn:aws:iam::000000000000:instance-profile/abc"

	err = w.InstanceProfiles.Add(ctx, compute.AddInstanceProfile{
		InstanceProfileArn: arn,
		SkipValidation:     true,
		IamRoleArn:         "arn:aws:iam::000000000000:role/bcd",
	})
	if err != nil {
		panic(err)
	}

}

func ExampleInstanceProfilesAPI_Edit_awsInstanceProfiles() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	arn := "arn:aws:iam::000000000000:instance-profile/abc"

	err = w.InstanceProfiles.Edit(ctx, compute.InstanceProfile{
		InstanceProfileArn: arn,
		IamRoleArn:         "arn:aws:iam::000000000000:role/bcdf",
	})
	if err != nil {
		panic(err)
	}

}

func ExampleInstanceProfilesAPI_List_awsInstanceProfiles() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.InstanceProfiles.ListAll(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}
