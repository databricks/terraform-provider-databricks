package mws

import (
	"context"
	"os"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/databrickslabs/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMwsAccVPCEndpointIntegration(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv != "MWS" {
		t.Skip("Cannot run test on non-MWS environment")
	}
	acctID := qa.GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	awsvreID := qa.GetEnvOrSkipTest(t, "TEST_REST_API_VPC_ENDPOINT")
	awsRegion := qa.GetEnvOrSkipTest(t, "TEST_REGION")
	client := common.CommonEnvironmentClient()
	ctx := context.Background()
	vpcEndpointAPI := NewVPCEndpointAPI(ctx, client)
	endpointList, err := vpcEndpointAPI.List(acctID)
	require.NoError(t, err, err)
	t.Logf("VPC Endpoints: %v", endpointList)

	vpcEndpoint := VPCEndpoint{
		AccountID:        acctID,
		VPCEndpointName:  qa.RandomName("tf-"),
		AwsVPCEndpointID: awsvreID,
		Region:           awsRegion,
	}
	err = vpcEndpointAPI.Create(&vpcEndpoint)
	require.NoError(t, err, err)
	defer func() {
		err = vpcEndpointAPI.Delete(acctID, vpcEndpoint.VPCEndpointID)
		assert.NoError(t, err, err)
	}()
	thisEndpoint, err := vpcEndpointAPI.Read(acctID, vpcEndpoint.VPCEndpointID)
	assert.NoError(t, err, err)
	assert.Equal(t, "available", thisEndpoint.State)
}

func TestResourceVPCEndpointCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/vpc-endpoints",
				ExpectedRequest: VPCEndpoint{
					AccountID:        "abc",
					VPCEndpointName:  "ve_name",
					Region:           "ar",
					AwsVPCEndpointID: "ave_id",
				},
				Response: VPCEndpoint{
					VPCEndpointID: "ve_id",
				},
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/accounts/abc/vpc-endpoints/ve_id",
				ReuseRequest: true,
				Response: VPCEndpoint{
					AccountID:        "abc",
					VPCEndpointName:  "ve_name",
					Region:           "ar",
					AwsVPCEndpointID: "ave_id",
					VPCEndpointID:    "ve_id",
					State:            "Available",
				},
			},
		},
		Resource: ResourceVPCEndpoint(),
		HCL: `
		account_id = "abc"
		vpc_endpoint_name = "ve_name"
		region = "ar"
		aws_vpc_endpoint_id = "ave_id"
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/ve_id", d.Id())
}

func TestResourceVPCEndpointCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/vpc-endpoints",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceVPCEndpoint(),
		State: map[string]interface{}{
			"account_id":          "abc",
			"vpc_endpoint_name":   "ve_name",
			"region":              "ar",
			"aws_vpc_endpoint_id": "ave_id",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceVPCEndpointRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/vpc-endpoints/veid",
				Response: VPCEndpoint{
					AccountID:       "veid",
					VPCEndpointName: "ve_name",
					Region:          "ar",
					VPCEndpointID:   "ave_id",
				},
			},
		},
		Resource: ResourceVPCEndpoint(),
		Read:     true,
		New:      true,
		ID:       "abc/veid",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/veid", d.Id(), "Id should not be empty")
	assert.Equal(t, "veid", d.Get("account_id"))
	assert.Equal(t, "ve_name", d.Get("vpc_endpoint_name"))
	assert.Equal(t, "ar", d.Get("region"))
	assert.Equal(t, "ave_id", d.Get("vpc_endpoint_id"))
}

func TestResourceVPCEndpointRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/vpc-endpoints/veid",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceVPCEndpoint(),
		Read:     true,
		Removed:  true,
		ID:       "abc/veid",
	}.ApplyNoError(t)
}

func TestResourceVPCEndpoint_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/vpc-endpoints/veid",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceVPCEndpoint(),
		Read:     true,
		ID:       "abc/veid",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/veid", d.Id(), "Id should not be empty for error reads")
}

func TestResourceVPCEndpointDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/abc/vpc-endpoints/veid",
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/vpc-endpoints/veid",
				Response: VPCEndpoint{
					AccountID:       "abc",
					VPCEndpointName: "ve_name",
					Region:          "ar",
					VPCEndpointID:   "ave_id",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/vpc-endpoints/veid",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Yes, it's not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceVPCEndpoint(),
		Delete:   true,
		ID:       "abc/veid",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/veid", d.Id())
}

func TestResourceVPCEndpointDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/abc/vpc-endpoints/veid",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceVPCEndpoint(),
		Delete:   true,
		ID:       "abc/veid",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/veid", d.Id())
}

func TestResourceVPCEndpointList(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/abc/vpc-endpoints",
			Response: []VPCEndpoint{},
		},
	})
	require.NoError(t, err)
	defer server.Close()

	l, err := NewVPCEndpointAPI(context.Background(), client).List("abc")
	require.NoError(t, err)
	assert.Len(t, l, 0)
}

func TestResourceVPCEndpointCreatePendingAndFails(t *testing.T) {
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/accounts/a/vpc-endpoints",
			ExpectedRequest: VPCEndpoint{
				AwsVPCEndpointID: "a",
				VPCEndpointName:  "a",
				Region:           "a",
				AccountID:        "a",
			},
			Response: VPCEndpoint{
				VPCEndpointID: "b",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/a/vpc-endpoints/b",
			Response: VPCEndpoint{
				State: "PENDING",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/accounts/a/vpc-endpoints/b",
			Response: VPCEndpoint{
				AwsVPCEndpointID: "x",
				State:            "bad thing",
			},
		},
	})
	require.NoError(t, err)
	defer server.Close()

	err = NewVPCEndpointAPI(context.Background(), client).Create(&VPCEndpoint{
		AccountID:        "a",
		AwsVPCEndpointID: "a",
		VPCEndpointName:  "a",
		Region:           "a",
	})
	require.EqualError(t, err, "cannot register x: bad thing")
}
