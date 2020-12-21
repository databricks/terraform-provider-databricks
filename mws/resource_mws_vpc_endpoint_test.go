package mws

import (
	"context"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"

	"github.com/stretchr/testify/assert"
)

func TestVPCEndpoint(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	acctID := qa.GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	client := common.CommonEnvironmentClient()
	ctx := context.Background()
	vpcEndpointAPI := NewVPCEndpointAPI(ctx, client)
	vpcEndpointList, err := vpcEndpointAPI.List(acctID)
	assert.NoError(t, err, err)
	t.Log(vpcEndpointList)

	vpcEndpoint := VPCEndpoint{
		AccountID:        acctID,
		VPCEndpointName:  qa.RandomName(),
		AWSVPCEndpointID: "",
	}
	err = vpcEndpointAPI.Create(&vpcEndpoint)
	assert.NoError(t, err, err)
	defer func() {
		err = vpcEndpointAPI.Delete(acctID, vpcEndpoint.AWSVPCEndpointID)
		assert.NoError(t, err, err)
	}()

	myVpcEndpoints, err := vpcEndpointAPI.Read(acctID, vpcEndpoint.AWSVPCEndpointID)
	assert.NoError(t, err, err)
	t.Log(myVpcEndpoints)
}
