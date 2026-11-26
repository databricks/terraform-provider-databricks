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

func ExampleWorkspacesAPI_Create_workspaces() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	storage, err := a.Storage.Create(ctx, provisioning.CreateStorageConfigurationRequest{
		StorageConfigurationName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		RootBucketInfo: provisioning.RootBucketInfo{
			BucketName: os.Getenv("TEST_ROOT_BUCKET"),
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", storage)

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

	created, err := a.Workspaces.CreateAndWait(ctx, provisioning.CreateWorkspaceRequest{
		WorkspaceName:          fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		AwsRegion:              os.Getenv("AWS_REGION"),
		CredentialsId:          role.CredentialsId,
		StorageConfigurationId: storage.StorageConfigurationId,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	_, err = a.Storage.DeleteByStorageConfigurationId(ctx, storage.StorageConfigurationId)
	if err != nil {
		panic(err)
	}
	_, err = a.Credentials.DeleteByCredentialsId(ctx, role.CredentialsId)
	if err != nil {
		panic(err)
	}
	_, err = a.Workspaces.DeleteByWorkspaceId(ctx, created.WorkspaceId)
	if err != nil {
		panic(err)
	}

}

func ExampleWorkspacesAPI_Get_workspaces() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	storage, err := a.Storage.Create(ctx, provisioning.CreateStorageConfigurationRequest{
		StorageConfigurationName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		RootBucketInfo: provisioning.RootBucketInfo{
			BucketName: os.Getenv("TEST_ROOT_BUCKET"),
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", storage)

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

	created, err := a.Workspaces.CreateAndWait(ctx, provisioning.CreateWorkspaceRequest{
		WorkspaceName:          fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		AwsRegion:              os.Getenv("AWS_REGION"),
		CredentialsId:          role.CredentialsId,
		StorageConfigurationId: storage.StorageConfigurationId,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	byId, err := a.Workspaces.GetByWorkspaceId(ctx, created.WorkspaceId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	_, err = a.Storage.DeleteByStorageConfigurationId(ctx, storage.StorageConfigurationId)
	if err != nil {
		panic(err)
	}
	_, err = a.Credentials.DeleteByCredentialsId(ctx, role.CredentialsId)
	if err != nil {
		panic(err)
	}
	_, err = a.Workspaces.DeleteByWorkspaceId(ctx, created.WorkspaceId)
	if err != nil {
		panic(err)
	}

}

func ExampleWorkspacesAPI_List_workspaces() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	all, err := a.Workspaces.List(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleWorkspacesAPI_Update_workspaces() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	storage, err := a.Storage.Create(ctx, provisioning.CreateStorageConfigurationRequest{
		StorageConfigurationName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		RootBucketInfo: provisioning.RootBucketInfo{
			BucketName: os.Getenv("TEST_ROOT_BUCKET"),
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", storage)

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

	updateRole, err := a.Credentials.Create(ctx, provisioning.CreateCredentialRequest{
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
	logger.Infof(ctx, "found %v", updateRole)

	created, err := a.Workspaces.CreateAndWait(ctx, provisioning.CreateWorkspaceRequest{
		WorkspaceName:          fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		AwsRegion:              os.Getenv("AWS_REGION"),
		CredentialsId:          role.CredentialsId,
		StorageConfigurationId: storage.StorageConfigurationId,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = a.Workspaces.UpdateAndWait(ctx, provisioning.UpdateWorkspaceRequest{
		WorkspaceId: created.WorkspaceId,
		CustomerFacingWorkspace: provisioning.Workspace{
			CredentialsId: updateRole.CredentialsId,
		},
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	_, err = a.Storage.DeleteByStorageConfigurationId(ctx, storage.StorageConfigurationId)
	if err != nil {
		panic(err)
	}
	_, err = a.Credentials.DeleteByCredentialsId(ctx, role.CredentialsId)
	if err != nil {
		panic(err)
	}
	_, err = a.Credentials.DeleteByCredentialsId(ctx, updateRole.CredentialsId)
	if err != nil {
		panic(err)
	}
	_, err = a.Workspaces.DeleteByWorkspaceId(ctx, created.WorkspaceId)
	if err != nil {
		panic(err)
	}

}
