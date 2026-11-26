// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package billing_test

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
)

func ExampleLogDeliveryAPI_Create_logDelivery() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	bucket, err := a.Storage.Create(ctx, provisioning.CreateStorageConfigurationRequest{
		StorageConfigurationName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		RootBucketInfo: provisioning.RootBucketInfo{
			BucketName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", bucket)

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

	created, err := a.LogDelivery.Create(ctx, billing.WrappedCreateLogDeliveryConfiguration{
		LogDeliveryConfiguration: billing.CreateLogDeliveryConfigurationParams{
			ConfigName:             fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			CredentialsId:          creds.CredentialsId,
			StorageConfigurationId: bucket.StorageConfigurationId,
			LogType:                billing.LogTypeAuditLogs,
			OutputFormat:           billing.OutputFormatJson,
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	_, err = a.Storage.DeleteByStorageConfigurationId(ctx, bucket.StorageConfigurationId)
	if err != nil {
		panic(err)
	}
	_, err = a.Credentials.DeleteByCredentialsId(ctx, creds.CredentialsId)
	if err != nil {
		panic(err)
	}
	err = a.LogDelivery.PatchStatus(ctx, billing.UpdateLogDeliveryConfigurationStatusRequest{
		LogDeliveryConfigurationId: created.LogDeliveryConfiguration.ConfigId,
		Status:                     billing.LogDeliveryConfigStatusDisabled,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleLogDeliveryAPI_Get_logDelivery() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	bucket, err := a.Storage.Create(ctx, provisioning.CreateStorageConfigurationRequest{
		StorageConfigurationName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		RootBucketInfo: provisioning.RootBucketInfo{
			BucketName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", bucket)

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

	created, err := a.LogDelivery.Create(ctx, billing.WrappedCreateLogDeliveryConfiguration{
		LogDeliveryConfiguration: billing.CreateLogDeliveryConfigurationParams{
			ConfigName:             fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			CredentialsId:          creds.CredentialsId,
			StorageConfigurationId: bucket.StorageConfigurationId,
			LogType:                billing.LogTypeAuditLogs,
			OutputFormat:           billing.OutputFormatJson,
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	byId, err := a.LogDelivery.GetByLogDeliveryConfigurationId(ctx, created.LogDeliveryConfiguration.ConfigId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	_, err = a.Storage.DeleteByStorageConfigurationId(ctx, bucket.StorageConfigurationId)
	if err != nil {
		panic(err)
	}
	_, err = a.Credentials.DeleteByCredentialsId(ctx, creds.CredentialsId)
	if err != nil {
		panic(err)
	}
	err = a.LogDelivery.PatchStatus(ctx, billing.UpdateLogDeliveryConfigurationStatusRequest{
		LogDeliveryConfigurationId: created.LogDeliveryConfiguration.ConfigId,
		Status:                     billing.LogDeliveryConfigStatusDisabled,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleLogDeliveryAPI_List_logDelivery() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	all, err := a.LogDelivery.ListAll(ctx, billing.ListLogDeliveryRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}
