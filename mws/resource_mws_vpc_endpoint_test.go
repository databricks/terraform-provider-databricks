package mws

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
		Resource: ResourceMwsVpcEndpoint(),
		HCL: `
		account_id = "abc"
		vpc_endpoint_name = "ve_name"
		region = "ar"
		aws_vpc_endpoint_id = "ave_id"
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/ve_id", d.Id())
}

func TestResourceVPCEndpointCreate_GCP(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/vpc-endpoints",
				ExpectedRequest: VPCEndpoint{
					AccountID:       "abc",
					VPCEndpointName: "ve_name",
					GcpVpcEndpointInfo: &GcpVpcEndpointInfo{
						ProjectId:       "project_a",
						PscEndpointName: "psc_endpoint_a",
						EndpointRegion:  "region_a",
					},
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
					AccountID:       "abc",
					VPCEndpointName: "ve_name",
					GcpVpcEndpointInfo: &GcpVpcEndpointInfo{
						ProjectId:           "project_a",
						PscEndpointName:     "psc_endpoint_a",
						EndpointRegion:      "region_a",
						PscConnectionId:     "120938102938209",
						ServiceAttachmentId: "service_attachment_a",
					},
				},
			},
		},
		Resource: ResourceMwsVpcEndpoint(),
		HCL: `
		account_id = "abc"
		vpc_endpoint_name = "ve_name"
		gcp_vpc_endpoint_info {
			project_id = "project_a"
			psc_endpoint_name = "psc_endpoint_a"
			endpoint_region = "region_a"
        }
		`,
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceVPCEndpointCreate_ConflictErrors(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Resource: ResourceMwsVpcEndpoint(),
		HCL: `
		account_id = "abc"
		vpc_endpoint_name = "ve_name"
		region = "ar"
		aws_vpc_endpoint_id = "ave_id"
		gcp_vpc_endpoint_info {
			project_id = "project_a"
			psc_endpoint_name = "psc_endpoint_a"
			endpoint_region = "region_a"
        }
		`,
		Create: true,
	}.Apply(t)
	assert.ErrorContains(t, err, "[gcp_vpc_endpoint_info] Conflicting configuration arguments")
	assert.ErrorContains(t, err, "[region] Invalid combination of arguments")
	assert.ErrorContains(t, err, "[aws_vpc_endpoint_id] Invalid combination of arguments")
}

func TestResourceVPCEndpointCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/vpc-endpoints",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceMwsVpcEndpoint(),
		State: map[string]any{
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
		Resource: ResourceMwsVpcEndpoint(),
		Read:     true,
		New:      true,
		ID:       "abc/veid",
	}.Apply(t)
	assert.NoError(t, err)
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
				Response: apierr.APIError{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceMwsVpcEndpoint(),
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
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceMwsVpcEndpoint(),
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
				Response: apierr.APIError{
					ErrorCode: "NOT_FOUND",
					Message:   "Yes, it's not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceMwsVpcEndpoint(),
		Delete:   true,
		ID:       "abc/veid",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/veid", d.Id())
}

func TestResourceVPCEndpointDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/abc/vpc-endpoints/veid",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceMwsVpcEndpoint(),
		Delete:   true,
		ID:       "abc/veid",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc/veid", d.Id())
}

func TestResourceVPCEndpointCreate_NoAccountIDInResource(t *testing.T) {
	qa.ResourceFixture{
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
		Resource: ResourceMwsVpcEndpoint(),
		HCL: `
		vpc_endpoint_name = "ve_name"
		region = "ar"
		aws_vpc_endpoint_id = "ave_id"
		`,
		AccountID: "abc",
		Create:    true,
	}.ApplyAndExpectData(t, map[string]any{
		"account_id": "abc",
		"id":         "abc/ve_id",
	})
}

func TestResourceVPCEndpointCreate_NoAccountID(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceMwsVpcEndpoint(),
		HCL: `
		vpc_endpoint_name = "ve_name"
		region = "ar"
		aws_vpc_endpoint_id = "ave_id"
		`,
		Create: true,
	}.ExpectError(t, "account_id is required in the provider block or in the resource")
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
