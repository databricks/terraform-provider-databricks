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
				},
				Response: catalog.ConnectionInfo{
					Name:           "testConnectionName",
					ConnectionType: catalog.ConnectionType("testConnectionType"),
					Comment:        "This is a test comment.",
					FullName:       "testConnectionName",
					MetastoreId:    "abc",
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
					MetastoreId:    "abc",
					Options: map[string]string{
						"host": "test.com",
					},
					Properties: map[string]string{
						"purpose": "testing",
					},
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/connections/testConnectionName",
				ExpectedRequest: catalog.UpdateConnection{
					Name: "testConnectionName",
					Options: map[string]string{
						"host": "test.com",
					},
					Owner: "InitialOwner",
				},
				Response: catalog.ConnectionInfo{
					Name:           "testConnectionName",
					ConnectionType: catalog.ConnectionType("testConnectionType"),
					Comment:        "This is a test comment.",
					FullName:       "testConnectionName",
					MetastoreId:    "abc",
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
					MetastoreId:    "abc",
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

func TestConnectionsCreate_BuiltinHms(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.1/unity-catalog/connections",
				ExpectedRequest: catalog.CreateConnection{
					Name:           "hms",
					ConnectionType: catalog.ConnectionType("HIVE_METASTORE"),
					Options: map[string]string{
						"builtin": "true",
					},
				},
				Response: catalog.ConnectionInfo{
					Name:           "hms",
					ConnectionType: catalog.ConnectionType("HIVE_METASTORE"),
					FullName:       "hms",
					MetastoreId:    "abc",
					Options: map[string]string{
						"builtin": "true",
						"host":    "test.com",
						"port":    "3306",
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/connections/hms?",
				Response: catalog.ConnectionInfo{
					Name:           "hms",
					ConnectionType: catalog.ConnectionType("HIVE_METASTORE"),
					FullName:       "hms",
					MetastoreId:    "abc",
					Options: map[string]string{
						"builtin": "true",
						"host":    "test.com",
						"port":    "3306",
					},
				},
			},
		},
		Resource: ResourceConnection(),
		Create:   true,
		HCL: `
		name = "hms"
		connection_type = "HIVE_METASTORE"
		options = {
			builtin = "true"
		}
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"name":            "hms",
		"connection_type": "HIVE_METASTORE",
	})
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
				},
				Response: apierr.APIError{
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
					MetastoreId:    "abc",
					Options: map[string]string{
						"host": "test.com",
					},
				},
			},
		},
		Resource: ResourceConnection(),
		Read:     true,
		ID:       "abc|testConnectionName",
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

func TestConnectionRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/connections/testConnectionName?",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceConnection(),
		Read:     true,
		ID:       "abc|testConnectionName",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc|testConnectionName", d.Id(), "Id should not be empty for error reads")
}

func TestConnectionsUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/connections/testConnectionName?",
				Response: catalog.ConnectionInfo{
					Name:           "testConnectionName",
					ConnectionType: catalog.ConnectionType("testConnectionType"),
					MetastoreId:    "abc",
					Comment:        "testComment",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/connections/testConnectionName",
				ExpectedRequest: catalog.UpdateConnection{
					Name: "testConnectionName",
					Options: map[string]string{
						"host": "test.com",
					},
				},
				Response: catalog.ConnectionInfo{
					Name:           "testConnectionName",
					ConnectionType: catalog.ConnectionType("testConnectionType"),
					Comment:        "testComment",
					MetastoreId:    "abc",
					Options: map[string]string{
						"host": "test.com",
					},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/connections/testConnectionName?",
				Response: catalog.ConnectionInfo{
					Name:           "testConnectionName",
					ConnectionType: catalog.ConnectionType("testConnectionType"),
					Comment:        "testComment",
					MetastoreId:    "abc",
					Options: map[string]string{
						"host": "test.com",
					},
				},
			},
		},
		Resource: ResourceConnection(),
		Update:   true,
		ID:       "abc|testConnectionName",
		InstanceState: map[string]string{
			"connection_type": "testConnectionType",
			"comment":         "testComment",
		},
		HCL: `
		name = "testConnectionName"
		connection_type = "testConnectionType"
		comment = "testComment"
		options = {
			host     = "test.com"
		}
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"name":            "testConnectionName",
		"connection_type": "testConnectionType",
		"comment":         "testComment",
	})
}

func TestConnectionsUpdateOwnerAndOtherFields(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/connections/testConnectionName?",
				Response: catalog.ConnectionInfo{
					Name:           "testConnectionName",
					ConnectionType: catalog.ConnectionType("testConnectionType"),
					MetastoreId:    "abc",
					Comment:        "testComment",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/connections/testConnectionName",
				ExpectedRequest: catalog.UpdateConnection{
					Name:  "testConnectionName",
					Owner: "admin",
				},
				Response: catalog.ConnectionInfo{
					Name:           "testConnectionName",
					ConnectionType: catalog.ConnectionType("testConnectionType"),
					Comment:        "testComment",
					MetastoreId:    "abc",
					Options: map[string]string{
						"host": "test.com",
					},
					Owner: "admin",
				},
			},
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/connections/testConnectionName",
				ExpectedRequest: catalog.UpdateConnection{
					Name: "testConnectionName",
					Options: map[string]string{
						"host": "test.com",
					},
				},
				Response: catalog.ConnectionInfo{
					Name:           "testConnectionName",
					ConnectionType: catalog.ConnectionType("testConnectionType"),
					Comment:        "testComment",
					MetastoreId:    "abc",
					Options: map[string]string{
						"host": "test.com",
					},
					Owner: "admin",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/connections/testConnectionName?",
				Response: catalog.ConnectionInfo{
					Name:           "testConnectionName",
					ConnectionType: catalog.ConnectionType("testConnectionType"),
					Comment:        "testComment",
					MetastoreId:    "abc",
					Options: map[string]string{
						"host": "test.com",
					},
					Owner: "admin",
				},
			},
		},
		Resource: ResourceConnection(),
		Update:   true,
		ID:       "abc|testConnectionName",
		InstanceState: map[string]string{
			"connection_type": "testConnectionType",
			"comment":         "testComment",
		},
		HCL: `
		name = "testConnectionName"
		connection_type = "testConnectionType"
		comment = "testComment"
		options = {
			host     = "test.com"
		}
		owner = "admin"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"name":            "testConnectionName",
		"connection_type": "testConnectionType",
		"comment":         "testComment",
		"owner":           "admin",
	})
}

func TestConnectionUpdate_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/connections/testConnectionName",
				ExpectedRequest: catalog.UpdateConnection{
					Name: "testConnectionName",
					Options: map[string]string{
						"host": "test.com",
					},
				},
				Response: apierr.APIError{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
		},
		Resource: ResourceConnection(),
		Update:   true,
		ID:       "abc|testConnectionName",
		InstanceState: map[string]string{
			"connection_type": "testConnectionType",
			"comment":         "testComment",
		},
		HCL: `
		name = "testConnectionName"
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
		ID:       "abc|testConnectionName",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc|testConnectionName", d.Id())
}

func TestConnectionDelete_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.1/unity-catalog/connections/testConnectionName?",
				Response: apierr.APIError{
					ErrorCode: "INVALID_STATE",
					Message:   "Something went wrong",
				},
				Status: 400,
			},
		},
		Resource: ResourceConnection(),
		Delete:   true,
		Removed:  true,
		ID:       "abc|testConnectionName",
	}.ExpectError(t, "Something went wrong")
}
