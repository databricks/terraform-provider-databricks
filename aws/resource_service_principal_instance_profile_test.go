package aws

import (
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/scim"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceServicePrincipalInstanceProfileCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				ExpectedRequest: scim.PatchRequest(
					"add",
					"roles",
					"arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile"),
				Response: scim.User{
					ID: "abc",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Response: scim.User{
					Schemas:     []scim.URN{scim.ServicePrincipalSchema},
					DisplayName: "ABC SP",
					Roles: []scim.ComplexValue{
						{
							Value: "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
						},
					},
					ID: "abc",
				},
			},
		},
		Resource: ResourceServicePrincipalInstanceProfile(),
		State: map[string]interface{}{
			"application_id":      "abc",
			"instance_profile_id": "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id())
}

func TestResourceServicePrincipalInstanceProfileCreate_Error_BadARN(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceServicePrincipalInstanceProfile(),
		State: map[string]interface{}{
			"application_id":      "abc",
			"instance_profile_id": "fake",
		},
		Create: true,
	}.Apply(t)
	assert.EqualError(t, err, "invalid config supplied. [instance_profile_id] Invalid ARN")
}

func TestResourceServicePrincipalInstanceProfileCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceServicePrincipalInstanceProfile(),
		State: map[string]interface{}{
			"application_id":      "abc",
			"instance_profile_id": "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceServicePrincipalInstanceProfileRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Response: scim.User{
					Schemas:     []scim.URN{scim.ServicePrincipalSchema},
					DisplayName: "ABC SP",
					Roles: []scim.ComplexValue{
						{
							Value: "arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
						},
					},
					ID: "abc",
				},
			},
		},
		Resource: ResourceServicePrincipalInstanceProfile(),
		Read:     true,
		ID:       "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile", d.Id(), "Id should not be empty")
}

func TestResourceServicePrincipalInstanceProfileRead_NoRole(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Response: scim.User{
					Schemas:     []scim.URN{scim.ServicePrincipalSchema},
					DisplayName: "ABC SP",
					ID:          "abc",
				},
			},
		},
		Resource: ResourceServicePrincipalInstanceProfile(),
		Read:     true,
		Removed:  true,
		ID:       "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.ApplyNoError(t)
}

func TestResourceServicePrincipalInstanceProfileRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/ServicePrincipals/abc",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceServicePrincipalInstanceProfile(),
		Read:     true,
		Removed:  true,
		ID:       "abc|arn:aws:iam::999999999999:instance-profile/my-fake-instance-profile",
	}.ApplyNoError(t)
}
