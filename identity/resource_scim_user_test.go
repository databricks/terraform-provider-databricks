package identity

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceScimUserCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Resource: ResourceScimUser(),
		Create:   true,
		HCL: `
		default_roles = ["a", "b", "c"]
		display_name = "Tom and Jerry"
		entitlements = ["allow-cluster-create"]
		set_admin = true
		user_name = "foo@bar.com"
		`,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "...", d.Id())
}
func TestResourceScimUserCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for better stub url...
				Method:   "POST",
				Resource: "/api/2.0/...",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceScimUser(),
		Create:   true,
		HCL: `
		default_roles = "..."
		display_name = "..."
		entitlements = "..."
		inherited_roles = "..."
		roles = "..."
		set_admin = "..."
		user_name = "..."
		
		`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}
func TestResourceScimUserRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			// read log output of test util for further stubs...
		},
		Resource: ResourceScimUser(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should not be empty")
	assert.Equal(t, "...default_roles", d.Get("default_roles"))
	assert.Equal(t, "...display_name", d.Get("display_name"))
	assert.Equal(t, "...entitlements", d.Get("entitlements"))
	assert.Equal(t, "...inherited_roles", d.Get("inherited_roles"))
	assert.Equal(t, "...roles", d.Get("roles"))
	assert.Equal(t, "...set_admin", d.Get("set_admin"))
	assert.Equal(t, "...user_name", d.Get("user_name"))

}
func TestResourceScimUserRead_NotFound(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/...",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceScimUser(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}
func TestResourceScimUserRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/...",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceScimUser(),
		Read:     true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id(), "Id should not be empty for error reads")
}
func TestResourceScimUserUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			// request #1 - most likely POST
			// request #2 - same as in TestResourceScimUserRead
		},
		Resource: ResourceScimUser(),
		Update:   true,
		ID:       "abc",
		HCL: `
		default_roles = "..."
		display_name = "..."
		entitlements = "..."
		inherited_roles = "..."
		roles = "..."
		set_admin = "..."
		user_name = "..."
		
		`,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id(), "Id should be the same as in reading")
}
func TestResourceScimUserUpdate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for better stub url...
				Method:   "POST",
				Resource: "/api/2.0/.../edit",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceScimUser(),
		Update:   true,
		ID:       "abc",
		HCL: `
		default_roles = "..."
		display_name = "..."
		entitlements = "..."
		inherited_roles = "..."
		roles = "..."
		set_admin = "..."
		user_name = "..."
		
		`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}
func TestResourceScimUserDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for better stub url...
				Method:   "POST",
				Resource: "/api/2.0/.../delete",
				ExpectedRequest: map[string]string{
					"...id": "abc",
				},
			},
		},
		Resource: ResourceScimUser(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc", d.Id())
}
func TestResourceScimUserDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/.../delete",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceScimUser(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}
