package catalog

import (
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestConnectionsCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceExternalLocation())
}

func TestConnectionsCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.1/unity-catalog/connections",
				ExpectedRequest: catalog.CreateConnection{
					Name:           "testConnectionName",
					ConnectionType: catalog.ConnectionType("testConnectionType"),
					Comment:        "This is a test comment.",
					Options: map[string]string{
						"host": "test.com",
					},
					Properties: map[string]string{
						"purpose": "testing",
					},
					Owner: "InitialOwner",
				},
				Response: catalog.ConnectionInfo{
					Name:           "testConnectionName",
					ConnectionType: catalog.ConnectionType("testConnectionType"),
					Comment:        "This is a test comment.",
					FullName:       "testConnectionName",
					Owner:          "InitialOwner",
					Options: map[string]string{
						"host": "test.com",
					},
					Properties: map[string]string{
						"purpose": "testing",
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/connections/testConnectionName?",
				Response: catalog.ConnectionInfo{
					Name:           "testConnectionName",
					ConnectionType: catalog.ConnectionType("testConnectionType"),
					Comment:        "This is a test comment.",
					FullName:       "testConnectionName",
					Owner:          "InitialOwner",
					Options: map[string]string{
						"host": "test.com",
					},
					Properties: map[string]string{
						"purpose": "testing",
					},
				},
			},
		},
		Resource: ResourceConnection(),
		Create:   true,
		HCL: `
		name = "testConnectionName"
		connection_type = "testConnectionType"
		options = {
			host     = "test.com"
		}
		properties = {
			purpose = "testing"
		}		
		comment = "This is a test comment."
		owner = "InitialOwner"
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "testConnectionName", d.Get("name"))
	assert.Equal(t, "testConnectionType", d.Get("connection_type"))
	assert.Equal(t, "This is a test comment.", d.Get("comment"))
	assert.Equal(t, map[string]interface{}{"host": "test.com"}, d.Get("options"))
	assert.Equal(t, map[string]interface{}{"purpose": "testing"}, d.Get("properties"))
}

func TestConnectionsCreate_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.1/unity-catalog/connections",
				ExpectedRequest: catalog.CreateConnection{
					Name:           "testConnectionName",
					ConnectionType: catalog.ConnectionType("testConnectionType"),
					Comment:        "This is a test comment.",
					Options: map[string]string{
						"host": "test.com",
					},
					Owner: "testOwner",
				},
				Response: apierr.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
		},
		Resource: ResourceConnection(),
		Create:   true,
		HCL: `
		name = "testConnectionName"
		owner = "testOwner"
		connection_type = "testConnectionType"
		options = {
			host     = "test.com"
		}		
		comment = "This is a test comment."
		`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected")
}

func TestConnectionsRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/connections/testConnectionName?",
				Response: catalog.ConnectionInfo{
					Name:           "testConnectionName",
					ConnectionType: catalog.ConnectionType("testConnectionType"),
					Comment:        "This is a test comment.",
					FullName:       "testConnectionName",
					Options: map[string]string{
						"host": "test.com",
					},
				},
			},
		},
		Resource: ResourceConnection(),
		Read:     true,
		ID:       "testConnectionName",
		HCL: `
		name = "testConnectionName"
		connection_type = "testConnectionType"
		options = {
			host     = "test.com"
		}		
		comment = "This is a test comment."
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "testConnectionName", d.Get("name"))
	assert.Equal(t, "testConnectionType", d.Get("connection_type"))
	assert.Equal(t, "This is a test comment.", d.Get("comment"))
	assert.Equal(t, map[string]interface{}{"host": "test.com"}, d.Get("options"))
}

func TestResourceConnectionRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/connections/testConnectionName?",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceConnection(),
		Read:     true,
		ID:       "testConnectionName",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "testConnectionName", d.Id(), "Id should not be empty for error reads")
}

func TestConnectionsUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/connections/testConnectionName?",
				Response: catalog.ConnectionInfo{
					Name:           "testConnectionName",
					ConnectionType: catalog.ConnectionType("testConnectionType"),
					Comment:        "testComment",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/connections/testConnectionName",
				ExpectedRequest: catalog.UpdateConnection{
					Name: "testConnectionNameNew",
					Options: map[string]string{
						"host": "test.com",
					},
				},
				Response: catalog.ConnectionInfo{
					Name:           "testConnectionNameNew",
					ConnectionType: catalog.ConnectionType("testConnectionType"),
					Comment:        "testComment",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/connections/testConnectionNameNew?",
				Response: catalog.ConnectionInfo{
					Name:           "testConnectionNameNew",
					ConnectionType: catalog.ConnectionType("testConnectionType"),
					Comment:        "testComment",
				},
			},
		},
		Resource: ResourceConnection(),
		Update:   true,
		ID:       "testConnectionName",
		InstanceState: map[string]string{
			"connection_type": "testConnectionType",
			"comment":         "testComment",
		},
		HCL: `
		name = "testConnectionNameNew"
		connection_type = "testConnectionType"
		comment = "testComment"
		options = {
			host     = "test.com"
		}
		`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "testConnectionNameNew", d.Get("name"))
	assert.Equal(t, "testConnectionType", d.Get("connection_type"))
	assert.Equal(t, "testComment", d.Get("comment"))
}

func TestConnectionUpdate_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/connections/testConnectionName",
				ExpectedRequest: catalog.UpdateConnection{
					Name: "testConnectionNameNew",
					Options: map[string]string{
						"host": "test.com",
					},
				},
				Response: apierr.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
		},
		Resource: ResourceConnection(),
		Update:   true,
		ID:       "testConnectionName",
		InstanceState: map[string]string{
			"connection_type": "testConnectionType",
			"comment":         "testComment",
		},
		HCL: `
		name = "testConnectionNameNew"
		connection_type = "testConnectionType"
		options = {
			host     = "test.com"
		}		
		comment = "testComment"
		`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected")
}

func TestConnectionDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.1/unity-catalog/connections/testConnectionName?",
			},
		},
		Resource: ResourceConnection(),
		Delete:   true,
		ID:       "testConnectionName",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "testConnectionName", d.Id())
}

func TestConnectionDelete_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.1/unity-catalog/connections/testConnectionName?",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_STATE",
					Message:   "Something went wrong",
				},
				Status: 400,
			},
		},
		Resource: ResourceConnection(),
		Delete:   true,
		Removed:  true,
		ID:       "testConnectionName",
	}.ExpectError(t, "Something went wrong")
}
