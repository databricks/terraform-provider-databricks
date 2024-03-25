package catalog

import (
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestRegisteredModelCornerCases(t *testing.T) {
	qa.ResourceCornerCases(t, ResourceRegisteredModel())
}

func TestRegisteredModelCreate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockRegisteredModelsAPI().EXPECT()
			e.Create(mock.Anything, catalog.CreateRegisteredModelRequest{
				Name:        "model",
				CatalogName: "catalog",
				SchemaName:  "schema",
				Comment:     "comment",
			}).Return(&catalog.RegisteredModelInfo{
				Name:        "model",
				Owner:       "owner",
				CatalogName: "catalog",
				SchemaName:  "schema",
				FullName:    "catalog.schema.model",
				Comment:     "comment",
			}, nil)
			e.GetByFullName(mock.Anything, "catalog.schema.model").Return(&catalog.RegisteredModelInfo{
				Name:        "model",
				CatalogName: "catalog",
				Owner:       "owner",
				SchemaName:  "schema",
				FullName:    "catalog.schema.model",
				Comment:     "comment",
			}, nil)
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
			"id":    "catalog.schema.model",
			"owner": "owner",
		},
	)
}

func TestRegisteredModelCreateWithOwner(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockRegisteredModelsAPI().EXPECT()
			e.Create(mock.Anything, catalog.CreateRegisteredModelRequest{
				Name:        "model",
				CatalogName: "catalog",
				SchemaName:  "schema",
				Comment:     "comment",
			}).Return(&catalog.RegisteredModelInfo{
				Name:        "model",
				Owner:       "old_owner",
				CatalogName: "catalog",
				SchemaName:  "schema",
				FullName:    "catalog.schema.model",
				Comment:     "comment",
			}, nil)
			e.Update(mock.Anything, catalog.UpdateRegisteredModelRequest{
				Owner:    "owner",
				FullName: "catalog.schema.model",
				Comment:  "comment",
			}).Return(&catalog.RegisteredModelInfo{
				Name:        "model",
				Owner:       "owner",
				CatalogName: "catalog",
				SchemaName:  "schema",
				FullName:    "catalog.schema.model",
				Comment:     "comment",
			}, nil)
			e.GetByFullName(mock.Anything, "catalog.schema.model").Return(&catalog.RegisteredModelInfo{
				Name:        "model",
				CatalogName: "catalog",
				Owner:       "owner",
				SchemaName:  "schema",
				FullName:    "catalog.schema.model",
				Comment:     "comment",
			}, nil)
		},
		Resource: ResourceRegisteredModel(),
		HCL: `
			name = "model"
			owner = "owner"
			catalog_name = "catalog"
			schema_name = "schema"
			comment = "comment"
			`,
		Create: true,
	}.ApplyAndExpectData(t,
		map[string]any{
			"id":    "catalog.schema.model",
			"owner": "owner",
		},
	)
}

func TestRegisteredModelCreate_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockRegisteredModelsAPI().EXPECT()
			e.Create(mock.Anything, catalog.CreateRegisteredModelRequest{}).Return(nil, errors.New("Internal error happened"))
		},
		Resource: ResourceRegisteredModel(),
		Create:   true,
	}.ExpectError(t, "Internal error happened")
}

func TestRegisteredModelRead(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockRegisteredModelsAPI().EXPECT()
			e.GetByFullName(mock.Anything, "catalog.schema.model").Return(&catalog.RegisteredModelInfo{
				Name:        "model",
				CatalogName: "catalog",
				SchemaName:  "schema",
				FullName:    "catalog.schema.model",
				Comment:     "comment",
			}, nil)
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockRegisteredModelsAPI().EXPECT()
			e.GetByFullName(mock.Anything, "catalog.schema.model").Return(nil, errors.New("Internal error happened"))
		},
		Resource: ResourceRegisteredModel(),
		Read:     true,
		ID:       "catalog.schema.model",
	}.ExpectError(t, "Internal error happened")
}

func TestRegisteredModelUpdate(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockRegisteredModelsAPI().EXPECT()
			e.Update(mock.Anything, catalog.UpdateRegisteredModelRequest{
				FullName: "catalog.schema.model",
				Comment:  "new comment",
			}).Return(&catalog.RegisteredModelInfo{
				Name:        "model",
				CatalogName: "catalog",
				SchemaName:  "schema",
				FullName:    "catalog.schema.model",
				Comment:     "new comment",
			}, nil)
			e.GetByFullName(mock.Anything, "catalog.schema.model").Return(&catalog.RegisteredModelInfo{
				Name:        "model",
				CatalogName: "catalog",
				SchemaName:  "schema",
				FullName:    "catalog.schema.model",
				Comment:     "new comment",
			}, nil)
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

func TestRegisteredModelUpdateOwner(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockRegisteredModelsAPI().EXPECT()
			e.Update(mock.Anything, catalog.UpdateRegisteredModelRequest{
				FullName: "catalog.schema.model",
				Owner:    "new_owner",
			}).Return(&catalog.RegisteredModelInfo{
				Name:        "model",
				CatalogName: "catalog",
				SchemaName:  "schema",
				Owner:       "new_owner",
				FullName:    "catalog.schema.model",
				Comment:     "comment",
			}, nil)
			e.Update(mock.Anything, catalog.UpdateRegisteredModelRequest{
				FullName: "catalog.schema.model",
				Comment:  "new comment",
			}).Return(&catalog.RegisteredModelInfo{
				Name:        "model",
				CatalogName: "catalog",
				SchemaName:  "schema",
				FullName:    "catalog.schema.model",
				Comment:     "new comment",
			}, nil)
			e.GetByFullName(mock.Anything, "catalog.schema.model").Return(&catalog.RegisteredModelInfo{
				Name:        "model",
				CatalogName: "catalog",
				SchemaName:  "schema",
				Owner:       "new_owner",
				FullName:    "catalog.schema.model",
				Comment:     "new comment",
			}, nil)
		},
		Resource: ResourceRegisteredModel(),
		Update:   true,
		ID:       "catalog.schema.model",
		InstanceState: map[string]string{
			"name":         "model",
			"catalog_name": "catalog",
			"owner":        "owner",
			"schema_name":  "schema",
			"comment":      "comment",
		},
		HCL: `
			name = "model"
			catalog_name = "catalog"
			owner = "new_owner"
			schema_name = "schema"
			comment = "new comment"
			`,
	}.ApplyAndExpectData(t, map[string]any{
		"id":           "catalog.schema.model",
		"owner":        "new_owner",
		"comment":      "new comment",
		"name":         "model",
		"catalog_name": "catalog",
		"schema_name":  "schema",
	})
}

func TestRegisteredModelUpdateRollback(t *testing.T) {
	_, err := qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockRegisteredModelsAPI().EXPECT()
			e.Update(mock.Anything, catalog.UpdateRegisteredModelRequest{
				FullName: "catalog.schema.model",
				Owner:    "new_owner",
			}).Return(&catalog.RegisteredModelInfo{
				Name:        "model",
				CatalogName: "catalog",
				SchemaName:  "schema",
				Owner:       "new_owner",
				FullName:    "catalog.schema.model",
				Comment:     "comment",
			}, nil)
			e.Update(mock.Anything, catalog.UpdateRegisteredModelRequest{
				FullName: "catalog.schema.model",
				Comment:  "new comment",
			}).Return(nil, errors.New("Something unexpected"))
			e.Update(mock.Anything, catalog.UpdateRegisteredModelRequest{
				FullName: "catalog.schema.model",
				Owner:    "owner",
			}).Return(&catalog.RegisteredModelInfo{
				Name:        "model",
				CatalogName: "catalog",
				SchemaName:  "schema",
				Owner:       "owner",
				FullName:    "catalog.schema.model",
				Comment:     "comment",
			}, nil)
		},
		Resource: ResourceRegisteredModel(),
		Update:   true,
		ID:       "catalog.schema.model",
		InstanceState: map[string]string{
			"name":         "model",
			"catalog_name": "catalog",
			"owner":        "owner",
			"schema_name":  "schema",
			"comment":      "comment",
		},
		HCL: `
			name = "model"
			catalog_name = "catalog"
			owner = "new_owner"
			schema_name = "schema"
			comment = "new comment"
			`,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Something unexpected")
}

func TestRegisteredModelUpdate_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockRegisteredModelsAPI().EXPECT()
			e.Update(mock.Anything, catalog.UpdateRegisteredModelRequest{
				FullName: "catalog.schema.model",
				Comment:  "new comment",
			}).Return(nil, errors.New("Internal error happened"))
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
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockRegisteredModelsAPI().EXPECT()
			e.DeleteByFullName(mock.Anything, "catalog.schema.model").Return(nil)
		},
		Resource: ResourceRegisteredModel(),
		Delete:   true,
		ID:       "catalog.schema.model",
	}.ApplyNoError(t)
}

func TestRegisteredModelDelete_Error(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			e := w.GetMockRegisteredModelsAPI().EXPECT()
			e.DeleteByFullName(mock.Anything, "catalog.schema.model").Return(errors.New("Internal error happened"))
		},
		Resource: ResourceRegisteredModel(),
		Delete:   true,
		ID:       "catalog.schema.model",
	}.ExpectError(t, "Internal error happened")
}
