// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package provisioning_test

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/provisioning"
)

func ExampleCredentialsAPI_Create_logDelivery() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	creds, err := a.Credentials.Create(ctx, provisioning.CreateCredentialRequest{
		CredentialsName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		AwsCredentials: provisioning.CreateCredentialAwsCredentials{
			StsRole: &provisioning.CreateCredentialStsRole{
				RoleArn: os.Getenv("TEST_LOGDELIVERY_ARN"),
			},
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", creds)

	// cleanup

	_, err = a.Credentials.DeleteByCredentialsId(ctx, creds.CredentialsId)
	if err != nil {
		panic(err)
	}

}

func ExampleCredentialsAPI_Create_credentials() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	role, err := a.Credentials.Create(ctx, provisioning.CreateCredentialRequest{
		CredentialsName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		AwsCredentials: provisioning.CreateCredentialAwsCredentials{
			StsRole: &provisioning.CreateCredentialStsRole{
				RoleArn: os.Getenv("TEST_CROSSACCOUNT_ARN"),
			},
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", role)

	// cleanup

	_, err = a.Credentials.DeleteByCredentialsId(ctx, role.CredentialsId)
	if err != nil {
		panic(err)
	}

}

func ExampleCredentialsAPI_Create_workspaces() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	role, err := a.Credentials.Create(ctx, provisioning.CreateCredentialRequest{
		CredentialsName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		AwsCredentials: provisioning.CreateCredentialAwsCredentials{
			StsRole: &provisioning.CreateCredentialStsRole{
				RoleArn: os.Getenv("TEST_CROSSACCOUNT_ARN"),
			},
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", role)

	// cleanup

	_, err = a.Credentials.DeleteByCredentialsId(ctx, role.CredentialsId)
	if err != nil {
		panic(err)
	}

}

func ExampleCredentialsAPI_Get_credentials() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	role, err := a.Credentials.Create(ctx, provisioning.CreateCredentialRequest{
		CredentialsName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		AwsCredentials: provisioning.CreateCredentialAwsCredentials{
			StsRole: &provisioning.CreateCredentialStsRole{
				RoleArn: os.Getenv("TEST_CROSSACCOUNT_ARN"),
			},
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", role)

	byId, err := a.Credentials.GetByCredentialsId(ctx, role.CredentialsId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	_, err = a.Credentials.DeleteByCredentialsId(ctx, role.CredentialsId)
	if err != nil {
		panic(err)
	}

}

func ExampleCredentialsAPI_List_credentials() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	configs, err := a.Credentials.List(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", configs)

}
