// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package provisioning_test

import (
	"context"
	"os"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/provisioning"
)

func ExampleEncryptionKeysAPI_Create_encryptionKeys() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	created, err := a.EncryptionKeys.Create(ctx, provisioning.CreateCustomerManagedKeyRequest{
		AwsKeyInfo: &provisioning.CreateAwsKeyInfo{
			KeyArn:   os.Getenv("TEST_MANAGED_KMS_KEY_ARN"),
			KeyAlias: os.Getenv("TEST_STORAGE_KMS_KEY_ALIAS"),
		},
		UseCases: []provisioning.KeyUseCase{provisioning.KeyUseCaseManagedServices},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	_, err = a.EncryptionKeys.DeleteByCustomerManagedKeyId(ctx, created.CustomerManagedKeyId)
	if err != nil {
		panic(err)
	}

}

func ExampleEncryptionKeysAPI_Get_encryptionKeys() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	created, err := a.EncryptionKeys.Create(ctx, provisioning.CreateCustomerManagedKeyRequest{
		AwsKeyInfo: &provisioning.CreateAwsKeyInfo{
			KeyArn:   os.Getenv("TEST_MANAGED_KMS_KEY_ARN"),
			KeyAlias: os.Getenv("TEST_STORAGE_KMS_KEY_ALIAS"),
		},
		UseCases: []provisioning.KeyUseCase{provisioning.KeyUseCaseManagedServices},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	byId, err := a.EncryptionKeys.GetByCustomerManagedKeyId(ctx, created.CustomerManagedKeyId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	_, err = a.EncryptionKeys.DeleteByCustomerManagedKeyId(ctx, created.CustomerManagedKeyId)
	if err != nil {
		panic(err)
	}

}

func ExampleEncryptionKeysAPI_List_encryptionKeys() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	all, err := a.EncryptionKeys.List(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}
