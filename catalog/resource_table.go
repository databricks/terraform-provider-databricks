package catalog

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type TablesAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewTablesAPI(ctx context.Context, m any) TablesAPI {
	return TablesAPI{m.(*common.DatabricksClient), context.WithValue(ctx, common.Api, common.API_2_1)}
}

type ColumnInfo struct {
	Name             string `json:"name"`
	TypeText         string `json:"type_text"`
	TypeJson         string `json:"type_json,omitempty"`
	TypeName         string `json:"type_name"`
	TypePrecision    int32  `json:"type_precision,omitempty"`
	TypeScale        int32  `json:"type_scale,omitempty"`
	TypeIntervalType string `json:"type_interval_type,omitempty"`
	Position         int32  `json:"position"`
	Comment          string `json:"comment,omitempty"`
	Nullable         bool   `json:"nullable,omitempty" tf:"default:true"`
	PartitionIndex   int32  `json:"partition_index,omitempty"`
}

type TableInfo struct {
	Name                  string            `json:"name"`
	CatalogName           string            `json:"catalog_name" tf:"force_new"`
	SchemaName            string            `json:"schema_name" tf:"force_new"`
	TableType             string            `json:"table_type" tf:"force_new"`
	DataSourceFormat      string            `json:"data_source_format"`
	ColumnInfos           []ColumnInfo      `json:"columns" tf:"alias:column"`
	StorageLocation       string            `json:"storage_location,omitempty" tf:"suppress_diff"`
	StorageCredentialName string            `json:"storage_credential_name,omitempty" tf:"force_new"`
	ViewDefinition        string            `json:"view_definition,omitempty"`
	Owner                 string            `json:"owner,omitempty" tf:"computed"`
	Comment               string            `json:"comment,omitempty"`
	Properties            map[string]string `json:"properties,omitempty"`
}

func (ti TableInfo) FullName() string {
	return fmt.Sprintf("%s.%s.%s", ti.CatalogName, ti.SchemaName, ti.Name)
}

func (a TablesAPI) createTable(ti *TableInfo) error {
	return a.client.Post(a.context, "/unity-catalog/tables", ti, ti)
}

func (a TablesAPI) getTable(name string) (ti TableInfo, err error) {
	err = a.client.Get(a.context, "/unity-catalog/tables/"+name, nil, &ti)
	return
}

func (a TablesAPI) deleteTable(name string) error {
	return a.client.Delete(a.context, "/unity-catalog/tables/"+name, nil)
}

func ResourceTable() common.Resource {
	tableSchema := common.StructToSchema(TableInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			caseInsensitiveFields := []string{"name", "catalog_name", "schema_name"}
			for _, field := range caseInsensitiveFields {
				m[field].DiffSuppressFunc = common.EqualFoldDiffSuppress
			}
			return m
		})
	update := updateFunctionFactory("/unity-catalog/tables", []string{
		"owner", "name", "data_source_format", "columns", "storage_location",
		"view_definition", "comment", "properties"})
	return common.Resource{
		Schema: tableSchema,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			if d.Get("table_type") != "EXTERNAL" {
				return nil
			}
			return nil
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ti TableInfo
			common.DataToStructPointer(d, tableSchema, &ti)
			if err := NewTablesAPI(ctx, c).createTable(&ti); err != nil {
				return err
			}
			d.SetId(ti.FullName())
			return update(ctx, d, c)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			ti, err := NewTablesAPI(ctx, c).getTable(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(ti, tableSchema, d)
		},
		Update: update,
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewTablesAPI(ctx, c).deleteTable(d.Id())
		},
	}
}
