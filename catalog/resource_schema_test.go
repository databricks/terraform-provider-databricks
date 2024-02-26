package catalog

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestSchemaCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceSchema())
}

func TestCreateSchema(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/schemas",
				ExpectedRequest: catalog.CreateSchema{
					Name:        "a",
					CatalogName: "b",
					Comment:     "c",
				},
				Response: catalog.SchemaInfo{
					FullName: "b.a",
					Comment:  "c",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/schemas/b.a?",
				Response: catalog.SchemaInfo{
					MetastoreId: "d",
					Comment:     "c",
					Owner:       "e",
				},
			},
		},
		Resource: ResourceSchema(),
		Create:   true,
		HCL: `
		name = "a"
		catalog_name = "b"
		comment = "c"
		`,
	}.ApplyNoError(t)
}

func TestCreateSchemaWithOwner(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.1/unity-catalog/schemas",
				ExpectedRequest: catalog.CreateSchema{
					Name:        "a",
					CatalogName: "b",
					Comment:     "c",
				},
				Response: catalog.SchemaInfo{
					FullName: "b.a",
					Comment:  "c",
					Owner:    "testers",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/schemas/b.a",
				ExpectedRequest: catalog.UpdateSchema{
					Owner:   "administrators",
					Comment: "c",
				},
				Response: catalog.SchemaInfo{
					FullName: "b.a",
					Comment:  "c",
					Owner:    "administrators",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/schemas/b.a?",
				Response: catalog.SchemaInfo{
					MetastoreId: "d",
					Comment:     "c",
					Owner:       "administrators",
				},
			},
		},
		Resource: ResourceSchema(),
		Create:   true,
		HCL: `
		name = "a"
		catalog_name = "b"
		comment = "c"
		owner = "administrators"
		`,
	}.ApplyNoError(t)
}

func TestUpdateSchema(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/current-metastore-assignment",
				Response: catalog.MetastoreAssignment{
					MetastoreId: "d",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/schemas/b.a",
				ExpectedRequest: catalog.UpdateSchema{
					Owner: "administrators",
				},
				Response: catalog.SchemaInfo{
					FullName: "b.a",
					Comment:  "c",
					Owner:    "administrators",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/schemas/b.a",
				ExpectedRequest: catalog.UpdateSchema{
					Comment: "c",
				},
				Response: catalog.SchemaInfo{
					FullName: "b.a",
					Comment:  "c",
					Owner:    "administrators",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/schemas/b.a?",
				Response: catalog.SchemaInfo{
					Name:        "a",
					CatalogName: "b",
					MetastoreId: "d",
					Comment:     "c",
					Owner:       "administrators",
				},
			},
		},
		Resource: ResourceSchema(),
		Update:   true,
		ID:       "b.a",
		InstanceState: map[string]string{
			"metastore_id": "d",
			"name":         "a",
			"catalog_name": "b",
			"comment":      "c",
		},
		HCL: `
		name = "a"
		catalog_name = "b"
		comment = "c"
		owner = "administrators"
		`,
	}.ApplyNoError(t)
}

func TestUpdateSchemaOwnerWithOtherFields(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/schemas/b.a",
				ExpectedRequest: catalog.UpdateSchema{
					Owner: "administrators",
				},
				Response: catalog.SchemaInfo{
					FullName: "b.a",
					Comment:  "c",
					Owner:    "administrators",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/schemas/b.a",
				ExpectedRequest: catalog.UpdateSchema{
					Comment: "d",
				},
				Response: catalog.SchemaInfo{
					FullName: "b.a",
					Comment:  "d",
					Owner:    "administrators",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/schemas/b.a?",
				Response: catalog.SchemaInfo{
					Name:        "a",
					CatalogName: "b",
					MetastoreId: "d",
					Comment:     "d",
					Owner:       "administrators",
				},
			},
		},
		Resource: ResourceSchema(),
		Update:   true,
		ID:       "b.a",
		InstanceState: map[string]string{
			"name":         "a",
			"catalog_name": "b",
			"comment":      "c",
			"owner":        "testOwner",
		},
		HCL: `
		name = "a"
		catalog_name = "b"
		comment = "d"
		owner = "administrators"
		`,
	}.ApplyNoError(t)
}

func TestUpdateSchemaRollback(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/schemas/b.a",
				ExpectedRequest: catalog.UpdateSchema{
					Owner: "administrators",
				},
				Response: catalog.SchemaInfo{
					FullName: "b.a",
					Comment:  "c",
					Owner:    "administrators",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/schemas/b.a",
				ExpectedRequest: catalog.UpdateSchema{
					Comment: "d",
				},
				Response: apierr.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   "Something unexpected happened",
				},
				Status: 500,
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/schemas/b.a",
				ExpectedRequest: catalog.UpdateSchema{
					Owner: "testOwner",
				},
				Response: catalog.SchemaInfo{
					FullName: "b.a",
					Comment:  "c",
					Owner:    "testOwner",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/schemas/b.a?",
				Response: catalog.SchemaInfo{
					Name:        "a",
					CatalogName: "b",
					MetastoreId: "d",
					Comment:     "d",
					Owner:       "testOwner",
				},
			},
		},
		Resource: ResourceSchema(),
		Update:   true,
		ID:       "b.a",
		InstanceState: map[string]string{
			"name":         "a",
			"catalog_name": "b",
			"comment":      "c",
			"owner":        "testOwner",
		},
		HCL: `
		name = "a"
		catalog_name = "b"
		comment = "d"
		owner = "administrators"
		`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected happened")
}

func TestUpdateSchemaRollback_Error(t *testing.T) {
	serverErrMessage := "Something unexpected happened"
	rollbackErrMessage := "Internal error happened"
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/schemas/b.a",
				ExpectedRequest: catalog.UpdateSchema{
					Owner: "administrators",
				},
				Response: catalog.SchemaInfo{
					FullName: "b.a",
					Comment:  "c",
					Owner:    "administrators",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/schemas/b.a",
				ExpectedRequest: catalog.UpdateSchema{
					Comment: "d",
				},
				Response: apierr.APIErrorBody{
					ErrorCode: "SERVER_ERROR",
					Message:   serverErrMessage,
				},
				Status: 500,
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/schemas/b.a",
				ExpectedRequest: catalog.UpdateSchema{
					Owner: "testOwner",
				},
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   rollbackErrMessage,
				},
				Status: 400,
			},
		},
		Resource: ResourceSchema(),
		Update:   true,
		ID:       "b.a",
		InstanceState: map[string]string{
			"name":         "a",
			"catalog_name": "b",
			"comment":      "c",
			"owner":        "testOwner",
		},
		HCL: `
		name = "a"
		catalog_name = "b"
		comment = "d"
		owner = "administrators"
		`,
	}.Apply(t)
	errOccurred := fmt.Sprintf("%s. Owner rollback also failed: %s", serverErrMessage, rollbackErrMessage)
	qa.AssertErrorStartsWith(t, err, errOccurred)
}

func TestUpdateSchemaForceNew(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/schemas/b.a",
				ExpectedRequest: catalog.UpdateSchema{
					Owner: "administrators",
				},
				Response: catalog.SchemaInfo{
					FullName: "b.a",
					Comment:  "c",
					Owner:    "administrators",
				},
			},
			{
				Method:   "PATCH",
				Resource: "/api/2.1/unity-catalog/schemas/b.a",
				ExpectedRequest: catalog.UpdateSchema{
					Comment: "c",
				},
				Response: catalog.SchemaInfo{
					FullName: "b.a",
					Comment:  "c",
					Owner:    "administrators",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/schemas/b.a?",
				Response: catalog.SchemaInfo{
					Name:        "a",
					MetastoreId: "d",
					Comment:     "c",
					Owner:       "administrators",
				},
			},
		},
		RequiresNew: true,
		Resource:    ResourceSchema(),
		Update:      true,
		ID:          "b.a",
		InstanceState: map[string]string{
			"metastore_id": "d",
			"name":         "a",
			"catalog_name": "b",
			"comment":      "c",
		},
		HCL: `
		name = "a"
		catalog_name = "x"
		comment = "c"
		owner = "administrators"
		`,
	}.ApplyNoError(t)
}

func TestDeleteSchema(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/schemas/b.a?",
			},
		},
		Resource: ResourceSchema(),
		Delete:   true,
		ID:       "b.a",
		HCL: `
		name = "a"
		catalog_name = "b"
		comment = "c"
		owner = "administrators"
		`,
	}.ApplyNoError(t)
}

func TestForceDeleteSchema(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/tables?catalog_name=b&schema_name=a",
				Response: catalog.ListTablesResponse{
					Tables: []catalog.TableInfo{
						{
							CatalogName: "b",
							SchemaName:  "a",
							Name:        "c",
							FullName:    "b.a.c",
							TableType:   "MANAGED",
						},
						{
							CatalogName: "b",
							SchemaName:  "a",
							Name:        "d",
							FullName:    "b.a.d",
							TableType:   "VIEW",
						},
					},
				},
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/tables/b.a.c?",
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/tables/b.a.d?",
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/volumes?catalog_name=b&schema_name=a",
				Response: catalog.ListVolumesResponseContent{
					Volumes: []catalog.VolumeInfo{
						{
							CatalogName: "b",
							SchemaName:  "a",
							Name:        "c",
							FullName:    "b.a.c",
							VolumeType:  catalog.VolumeTypeManaged,
						},
						{
							CatalogName: "b",
							SchemaName:  "a",
							Name:        "d",
							FullName:    "b.a.d",
							VolumeType:  catalog.VolumeTypeExternal,
						},
					},
				},
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/volumes/b.a.c?",
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/volumes/b.a.d?",
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/functions?catalog_name=b&schema_name=a",
				Response: catalog.ListFunctionsResponse{
					Functions: []catalog.FunctionInfo{
						{
							CatalogName: "b",
							SchemaName:  "a",
							Name:        "c",
							FullName:    "b.a.c",
						},
					},
				},
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/functions/b.a.c?",
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/unity-catalog/models?catalog_name=b&schema_name=a",
				Response: catalog.ListRegisteredModelsResponse{
					RegisteredModels: []catalog.RegisteredModelInfo{
						{
							CatalogName: "b",
							SchemaName:  "a",
							Name:        "c",
							FullName:    "b.a.c",
						},
					},
				},
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/models/b.a.c?",
			},
			{
				Method:   "DELETE",
				Resource: "/api/2.1/unity-catalog/schemas/b.a?",
			},
		},
		Resource: ResourceSchema(),
		Delete:   true,
		ID:       "b.a",
		HCL: `
		name = "a"
		catalog_name = "b"
		comment = "c"
		owner = "administrators"
		force_destroy = true
		`,
	}.ApplyNoError(t)
}
