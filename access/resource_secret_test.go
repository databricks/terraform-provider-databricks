package access

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceSecretRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/list?scope=foo",
				Response: SecretsList{
					Secrets: []SecretMetadata{
						{
							Key:                  "bar",
							LastUpdatedTimestamp: 12345678,
						},
					},
				},
			},
		},
		Resource: ResourceSecret(),
		Read:     true,
		ID:       "foo|||bar",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "foo|||bar", d.Id())
	assert.Equal(t, "bar", d.Get("key"))
	assert.Equal(t, 12345678, d.Get("last_updated_timestamp"))
	assert.Equal(t, "foo", d.Get("scope"))
	assert.Equal(t, "", d.Get("string_value"))
}

func TestResourceSecretRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/list?scope=foo",
				Response: SecretsList{
					Secrets: []SecretMetadata{
						{
							Key:                  "bar",
							LastUpdatedTimestamp: 12345678,
						},
					},
				},
			},
		},
		Resource: ResourceSecret(),
		Read:     true,
		Removed:  true,
		ID:       "foo|||missing",
	}.ApplyNoError(t)
}

func TestResourceSecretRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/list?scope=foo",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceSecret(),
		Read:     true,
		ID:       "foo|||bar",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "foo|||bar", d.Id(), "Id should not be empty for error reads")
}

func TestResourceSecretCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/put",
				ExpectedRequest: SecretsRequest{
					StringValue: "SparkIsTh3Be$t",
					Scope:       "foo",
					Key:         "bar",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/list?scope=foo",
				Response: SecretsList{
					Secrets: []SecretMetadata{
						{
							Key:                  "bar",
							LastUpdatedTimestamp: 12345678,
						},
					},
				},
			},
		},
		Resource: ResourceSecret(),
		State: map[string]interface{}{
			"scope":        "foo",
			"key":          "bar",
			"string_value": "SparkIsTh3Be$t",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "foo|||bar", d.Id())
}

func TestResourceSecretCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/put",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceSecret(),
		State: map[string]interface{}{
			"key":          "...",
			"scope":        "...",
			"string_value": "...",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceSecretDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for better stub url...
				Method:   "POST",
				Resource: "/api/2.0/secrets/delete",
				ExpectedRequest: map[string]string{
					"scope": "foo",
					"key":   "bar",
				},
			},
		},
		Resource: ResourceSecret(),
		Delete:   true,
		ID:       "foo|||bar",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "foo|||bar", d.Id())
}

func TestResourceSecretDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/secrets/delete",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceSecret(),
		Delete:   true,
		ID:       "foo|||bar",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "foo|||bar", d.Id())
}
