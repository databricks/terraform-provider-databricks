package catalog

import (
	"errors"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestSchemaCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceSchema())
}

func TestCreateSchema(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockSchemasAPI().EXPECT()
			e.Create(mock.Anything, catalog.CreateSchema{
				Name:        "a",
				CatalogName: "b",
				Comment:     "c",
			}).Return(&catalog.SchemaInfo{
				FullName: "b.a",
				Comment:  "c",
				Owner:    "testers",
			}, nil)
			e.GetByFullName(mock.Anything, "b.a").Return(&catalog.SchemaInfo{
				MetastoreId: "d",
				Comment:     "c",
				Owner:       "testers",
			}, nil)
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockSchemasAPI().EXPECT()
			e.Create(mock.Anything, catalog.CreateSchema{
				Name:        "a",
				CatalogName: "b",
				Comment:     "c",
			}).Return(&catalog.SchemaInfo{
				FullName: "b.a",
				Comment:  "c",
				Owner:    "testers",
			}, nil)
			e.Update(mock.Anything, catalog.UpdateSchema{
				FullName:                     "b.a",
				Comment:                      "c",
				Owner:                        "administrators",
				EnablePredictiveOptimization: "ENABLE",
			}).Return(&catalog.SchemaInfo{
				FullName: "b.a",
				Comment:  "c",
				Owner:    "administrators",
			}, nil)
			e.GetByFullName(mock.Anything, "b.a").Return(&catalog.SchemaInfo{
				MetastoreId:                  "d",
				Comment:                      "c",
				Owner:                        "administrators",
				EnablePredictiveOptimization: "ENABLE",
			}, nil)
		},
		Resource: ResourceSchema(),
		Create:   true,
		HCL: `
		name = "a"
		catalog_name = "b"
		comment = "c"
		owner = "administrators"
		enable_predictive_optimization = "ENABLE"
		`,
	}.ApplyNoError(t)
}

func TestUpdateSchema(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockMetastoresAPI().EXPECT().Current(mock.Anything).Return(&catalog.MetastoreAssignment{
				MetastoreId: "d",
			}, nil)
			e := w.GetMockSchemasAPI().EXPECT()
			e.Update(mock.Anything, catalog.UpdateSchema{
				FullName: "b.a",
				Owner:    "administrators",
			}).Return(&catalog.SchemaInfo{
				FullName: "b.a",
				Comment:  "c",
				Owner:    "administrators",
			}, nil)
			e.Update(mock.Anything, catalog.UpdateSchema{
				FullName: "b.a",
				Comment:  "d",
			}).Return(&catalog.SchemaInfo{
				FullName: "b.a",
				Comment:  "d",
				Owner:    "administrators",
			}, nil)
			e.GetByFullName(mock.Anything, "b.a").Return(&catalog.SchemaInfo{
				Name:        "a",
				CatalogName: "b",
				MetastoreId: "d",
				Comment:     "d",
				Owner:       "administrators",
			}, nil)
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
		comment = "d"
		owner = "administrators"
		`,
	}.ApplyNoError(t)
}

func TestUpdateSchemaSetEmptyComment(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockMetastoresAPI().EXPECT().Current(mock.Anything).Return(&catalog.MetastoreAssignment{
				MetastoreId: "d",
			}, nil)
			e := w.GetMockSchemasAPI().EXPECT()
			e.Update(mock.Anything, catalog.UpdateSchema{
				FullName:        "b.a",
				Comment:         "",
				ForceSendFields: []string{"Comment"},
			}).Return(&catalog.SchemaInfo{
				FullName: "b.a",
				Owner:    "administrators",
			}, nil)
			e.GetByFullName(mock.Anything, "b.a").Return(&catalog.SchemaInfo{
				Name:        "a",
				CatalogName: "b",
				MetastoreId: "d",
				Owner:       "administrators",
			}, nil)
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
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"comment": "",
	})
}

func TestUpdateSchemaChangeForceDestroy(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockMetastoresAPI().EXPECT().Current(mock.Anything).Return(&catalog.MetastoreAssignment{
				MetastoreId: "d",
			}, nil)
			e := w.GetMockSchemasAPI().EXPECT()
			e.GetByFullName(mock.Anything, "b.a").Return(&catalog.SchemaInfo{
				Name:        "a",
				CatalogName: "b",
				MetastoreId: "d",
				Owner:       "administrators",
			}, nil)
		},
		Resource: ResourceSchema(),
		Update:   true,
		ID:       "b.a",
		InstanceState: map[string]string{
			"metastore_id":  "d",
			"name":          "a",
			"catalog_name":  "b",
			"force_destroy": "true",
		},
		HCL: `
		name = "a"
		catalog_name = "b"
		force_destroy = false
		`,
	}.ApplyNoError(t)
}

func TestUpdateSchemaOwnerWithOtherFields(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockSchemasAPI().EXPECT()
			e.Update(mock.Anything, catalog.UpdateSchema{
				FullName: "b.a",
				Owner:    "administrators",
			}).Return(&catalog.SchemaInfo{
				FullName: "b.a",
				Comment:  "c",
				Owner:    "administrators",
			}, nil)
			e.Update(mock.Anything, catalog.UpdateSchema{
				FullName: "b.a",
				Comment:  "d",
			}).Return(&catalog.SchemaInfo{
				FullName: "b.a",
				Comment:  "d",
				Owner:    "administrators",
			}, nil)
			e.GetByFullName(mock.Anything, "b.a").Return(&catalog.SchemaInfo{
				Name:        "a",
				CatalogName: "b",
				MetastoreId: "d",
				Comment:     "d",
				Owner:       "administrators",
			}, nil)
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockSchemasAPI().EXPECT()
			e.Update(mock.Anything, catalog.UpdateSchema{
				FullName: "b.a",
				Owner:    "administrators",
			}).Return(&catalog.SchemaInfo{
				FullName: "b.a",
				Comment:  "c",
				Owner:    "administrators",
			}, nil)
			e.Update(mock.Anything, catalog.UpdateSchema{
				FullName: "b.a",
				Comment:  "d",
			}).Return(nil, errors.New("Something unexpected happened"))
			e.Update(mock.Anything, catalog.UpdateSchema{
				FullName: "b.a",
				Owner:    "testOwner",
			}).Return(&catalog.SchemaInfo{
				FullName: "b.a",
				Comment:  "c",
				Owner:    "testOwner",
			}, nil)
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockSchemasAPI().EXPECT()
			e.Update(mock.Anything, catalog.UpdateSchema{
				FullName: "b.a",
				Owner:    "administrators",
			}).Return(&catalog.SchemaInfo{
				FullName: "b.a",
				Comment:  "c",
				Owner:    "administrators",
			}, nil)
			e.Update(mock.Anything, catalog.UpdateSchema{
				FullName: "b.a",
				Comment:  "d",
			}).Return(nil, errors.New(serverErrMessage))
			e.Update(mock.Anything, catalog.UpdateSchema{
				FullName: "b.a",
				Owner:    "testOwner",
			}).Return(nil, errors.New(rollbackErrMessage))
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockSchemasAPI().EXPECT()
			e.Update(mock.Anything, catalog.UpdateSchema{
				FullName: "b.a",
				Owner:    "administrators",
			}).Return(&catalog.SchemaInfo{
				FullName: "b.a",
				Comment:  "c",
				Owner:    "administrators",
			}, nil)
			e.Update(mock.Anything, catalog.UpdateSchema{
				FullName: "b.a",
				Comment:  "c",
			}).Return(&catalog.SchemaInfo{
				FullName: "b.a",
				Comment:  "c",
				Owner:    "administrators",
			}, nil)
			e.GetByFullName(mock.Anything, "b.a").Return(&catalog.SchemaInfo{
				Name:        "a",
				MetastoreId: "d",
				Comment:     "c",
				Owner:       "administrators",
			}, nil)
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockSchemasAPI().EXPECT().DeleteByFullName(mock.Anything, "b.a").Return(nil)
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			t := w.GetMockTablesAPI().EXPECT()
			t.ListAll(mock.Anything, catalog.ListTablesRequest{
				CatalogName: "b",
				SchemaName:  "a",
			}).Return([]catalog.TableInfo{
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
			}, nil)
			t.DeleteByFullName(mock.Anything, "b.a.c").Return(nil)
			t.DeleteByFullName(mock.Anything, "b.a.d").Return(nil)
			v := w.GetMockVolumesAPI().EXPECT()
			v.ListAll(mock.Anything, catalog.ListVolumesRequest{
				CatalogName: "b",
				SchemaName:  "a",
			}).Return([]catalog.VolumeInfo{
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
			}, nil)
			v.DeleteByName(mock.Anything, "b.a.c").Return(nil)
			v.DeleteByName(mock.Anything, "b.a.d").Return(nil)
			f := w.GetMockFunctionsAPI().EXPECT()
			f.ListAll(mock.Anything, catalog.ListFunctionsRequest{
				CatalogName: "b",
				SchemaName:  "a",
			}).Return([]catalog.FunctionInfo{
				{
					CatalogName: "b",
					SchemaName:  "a",
					Name:        "c",
					FullName:    "b.a.c",
				},
			}, nil)
			f.DeleteByName(mock.Anything, "b.a.c").Return(nil)
			m := w.GetMockRegisteredModelsAPI().EXPECT()
			m.ListAll(mock.Anything, catalog.ListRegisteredModelsRequest{
				CatalogName: "b",
				SchemaName:  "a",
			}).Return([]catalog.RegisteredModelInfo{
				{
					CatalogName: "b",
					SchemaName:  "a",
					Name:        "c",
					FullName:    "b.a.c",
				},
			}, nil)
			m.DeleteByFullName(mock.Anything, "b.a.c").Return(nil)
			w.GetMockSchemasAPI().EXPECT().DeleteByFullName(mock.Anything, "b.a").Return(nil)
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
