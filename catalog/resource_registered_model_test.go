package catalog

import (
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestRegisteredModelCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceRegisteredModel())
}

func TestRegisteredModelCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.1/unity-catalog/models",
				ExpectedRequest: catalog.CreateRegisteredModelRequest{
					Name:        "model",
					CatalogName: "catalog",
					SchemaName:  "schema",
					Comment:     "comment",
				},
				Response: catalog.RegisteredModelInfo{
					Name:        "model",
					CatalogName: "catalog",
					SchemaName:  "schema",
					FullName:    "catalog.schema.model",
					Comment:     "comment",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/models/catalog.schema.model?",
				Response: catalog.RegisteredModelInfo{
					Name:        "model",
					CatalogName: "catalog",
					SchemaName:  "schema",
					FullName:    "catalog.schema.model",
					Comment:     "comment",
				},
			},
		},
		Resource: ResourceRegisteredModel(),
		HCL: `
			name = "model"
			catalog_name = "catalog"
			schema_name = "schema"
			comment = "comment"
			`,
		Create: true,
	}.ApplyAndExpectData(t,
		map[string]any{
			"id": "catalog.schema.model",
		},
	)
}

func TestRegisteredModelCreate_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.1/unity-catalog/models",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceRegisteredModel(),
		Create:   true,
	}.ExpectError(t, "Internal error happened")
}

func TestRegisteredModelRead(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/models/catalog.schema.model?",
				Response: catalog.RegisteredModelInfo{
					Name:        "model",
					CatalogName: "catalog",
					SchemaName:  "schema",
					FullName:    "catalog.schema.model",
					Comment:     "comment",
				},
			},
		},
		Resource: ResourceRegisteredModel(),
		Read:     true,
		ID:       "catalog.schema.model",
	}.ApplyAndExpectData(t,
		map[string]any{
			"id": "catalog.schema.model",
		},
	)
}

func TestRegisteredModelRead_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/models/catalog.schema.model?",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceRegisteredModel(),
		Read:     true,
		ID:       "catalog.schema.model",
	}.ExpectError(t, "Internal error happened")
}

func TestRegisteredModelUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/models/catalog.schema.model",
				ExpectedRequest: catalog.UpdateRegisteredModelRequest{
					FullName: "catalog.schema.model",
					Comment:  "new comment",
				},
				Response: catalog.RegisteredModelInfo{
					Name:        "model",
					CatalogName: "catalog",
					SchemaName:  "schema",
					FullName:    "catalog.schema.model",
					Comment:     "new comment",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.1/unity-catalog/models/catalog.schema.model?",
				Response: catalog.RegisteredModelInfo{
					Name:        "model",
					CatalogName: "catalog",
					SchemaName:  "schema",
					FullName:    "catalog.schema.model",
					Comment:     "new comment",
				},
			},
		},
		Resource: ResourceRegisteredModel(),
		Update:   true,
		ID:       "catalog.schema.model",
		InstanceState: map[string]string{
			"name":         "model",
			"catalog_name": "catalog",
			"schema_name":  "schema",
			"comment":      "comment",
		},
		HCL: `
			name = "model"
			catalog_name = "catalog"
			schema_name = "schema"
			comment = "new comment"
			`,
	}.ApplyNoError(t)
}

func TestRegisteredModelUpdate_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPatch,
				Resource: "/api/2.1/unity-catalog/models/catalog.schema.model",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceRegisteredModel(),
		Update:   true,
		ID:       "catalog.schema.model",
		InstanceState: map[string]string{
			"name":         "model",
			"catalog_name": "catalog",
			"schema_name":  "schema",
			"comment":      "comment",
		},
		HCL: `
			name = "model"
			catalog_name = "catalog"
			schema_name = "schema"
			comment = "new comment"
			`,
	}.ExpectError(t, "Internal error happened")
}

func TestRegisteredModelDelete(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.1/unity-catalog/models/catalog.schema.model?",
				Response: "",
			},
		},
		Resource: ResourceRegisteredModel(),
		Delete:   true,
		ID:       "catalog.schema.model",
	}.ApplyNoError(t)
}

func TestRegisteredModelDelete_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodDelete,
				Resource: "/api/2.1/unity-catalog/models/catalog.schema.model?",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceRegisteredModel(),
		Delete:   true,
		ID:       "catalog.schema.model",
	}.ExpectError(t, "Internal error happened")
}
