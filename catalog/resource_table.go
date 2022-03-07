package catalog

import (
	"context"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type TablesAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewTablesAPI(ctx context.Context, m interface{}) TablesAPI {
	return TablesAPI{m.(*common.DatabricksClient), ctx}
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
	Name                  string            `json:"name" tf:"force_new"`
	CatalogName           string            `json:"catalog_name"`
	SchemaName            string            `json:"schema_name"`
	TableType             string            `json:"table_type"`
	DataSourceFormat      string            `json:"data_source_format"`
	ColumnInfos           []ColumnInfo      `json:"columns" tf:"alias:column"`
	StorageLocation       string            `json:"storage_location,omitempty"`
	StorageCredentialName string            `json:"storage_credential_name,omitempty"`
	ViewDefinition        string            `json:"view_definition,omitempty"`
	Owner                 string            `json:"owner,omitempty" tf:"computed"`
	Comment               string            `json:"comment,omitempty"`
	Properties            map[string]string `json:"properties,omitempty"`
}

func (ti TableInfo) FullName() string {
	return fmt.Sprintf("%s.%s.%s", ti.CatalogName, ti.SchemaName, ti.Name)
}

type Tables struct {
	Tables []TableInfo `json:"tables"`
}

func (a TablesAPI) listTables(catalogName, schemaName string) (tables Tables, err error) {
	err = a.client.Get(a.context, "/unity-catalog/tables/", map[string]string{
		"catalog_name": catalogName,
		"schema_name":  schemaName,
	}, &tables)
	return
}

func (a TablesAPI) createTable(ti *TableInfo) error {
	return a.client.Post(a.context, "/unity-catalog/tables", ti, ti)
}

func (a TablesAPI) getTable(name string) (ti TableInfo, err error) {
	err = a.client.Get(a.context, "/unity-catalog/tables/"+name, nil, &ti)
	return
}

func (a TablesAPI) updateTable(ti TableInfo) error {
	return a.client.Patch(a.context, "/unity-catalog/tables/"+ti.Name, ti)
}

func (a TablesAPI) deleteTable(name string) error {
	return a.client.Delete(a.context, "/unity-catalog/tables/"+name, nil)
}

func ResourceTable() *schema.Resource {
	tableSchema := common.StructToSchema(TableInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return m
		})
	return common.Resource{
		Schema: tableSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ti TableInfo
			common.DataToStructPointer(d, tableSchema, &ti)
			if err := NewTablesAPI(ctx, c).createTable(&ti); err != nil {
				return err
			}
			d.SetId(ti.FullName())
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			ti, err := NewTablesAPI(ctx, c).getTable(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(ti, tableSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ti TableInfo
			common.DataToStructPointer(d, tableSchema, &ti)
			return NewTablesAPI(ctx, c).updateTable(ti)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewTablesAPI(ctx, c).deleteTable(d.Id())
		},
	}.ToResource()
}
