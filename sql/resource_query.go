package sql

import (
	"context"
	"log"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Need a struct for Query because there are aliases we need and it'll be needed in the create method.
type QueryStruct struct {
	sql.Query
	common.Namespace
}

var queryAliasMap = map[string]string{
	"parameters": "parameter",
}

func (QueryStruct) Aliases() map[string]map[string]string {
	return map[string]map[string]string{
		"sql.QueryStruct": queryAliasMap,
	}
}

func (QueryStruct) CustomizeSchema(m *common.CustomizableSchema) *common.CustomizableSchema {
	common.NamespaceCustomizeSchema(m)
	m.SchemaPath("display_name").SetRequired().SetValidateFunc(validation.StringIsNotWhiteSpace)
	m.SchemaPath("query_text").SetRequired()
	m.SchemaPath("warehouse_id").SetRequired().SetValidateFunc(validation.StringIsNotWhiteSpace)
	m.SchemaPath("parent_path").SetCustomSuppressDiff(common.WorkspaceOrEmptyPathPrefixDiffSuppress).SetForceNew()
	m.SchemaPath("owner_user_name").SetSuppressDiff()
	m.SchemaPath("run_as_mode").SetSuppressDiff().SetValidateFunc(validation.StringInSlice([]string{"VIEWER", "OWNER"}, false))
	m.SchemaPath("id").SetReadOnly()
	m.SchemaPath("create_time").SetReadOnly()
	m.SchemaPath("lifecycle_state").SetReadOnly()
	m.SchemaPath("last_modifier_user_name").SetReadOnly()
	m.SchemaPath("update_time").SetReadOnly()

	// customize parameters
	m.SchemaPath("parameter", "name").SetRequired().SetValidateFunc(validation.StringIsNotWhiteSpace)
	m.SchemaPath("parameter", "date_range_value", "precision").SetSuppressDiff()
	m.SchemaPath("parameter", "date_value", "precision").SetSuppressDiff()
	m.SchemaPath("parameter", "query_backed_value", "query_id").SetRequired()
	m.SchemaPath("parameter", "text_value", "value").SetRequired()
	m.SchemaPath("parameter", "numeric_value", "value").SetRequired()
	// TODO: fix setting of AtLeastOneOf
	// valuesAlof := []string{
	// 	"parameter.0.date_range_value",
	// 	"parameter.0.date_value",
	// 	"parameter.0.query_backed_value",
	// 	"parameter.0.text_value",
	// 	"parameter.0.numeric_value",
	// 	"parameter.0.enum_value",
	// }
	// for _, f := range valuesAlof {
	// 	m.SchemaPath("parameter", strings.TrimPrefix(f, "parameter.0.")).SetAtLeastOneOf(valuesAlof)
	// }
	return m
}

type queryCreateStruct struct {
	sql.CreateQueryRequestQuery
}

func (queryCreateStruct) Aliases() map[string]map[string]string {
	return map[string]map[string]string{
		"sql.queryCreateStruct": queryAliasMap,
	}
}

func (queryCreateStruct) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	return s
}

type queryUpdateStruct struct {
	sql.UpdateQueryRequestQuery
}

func (queryUpdateStruct) Aliases() map[string]map[string]string {
	return map[string]map[string]string{
		"sql.queryUpdateStruct": queryAliasMap,
	}
}

func (queryUpdateStruct) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	return s
}

func ResourceQuery() common.Resource {
	s := common.StructToSchema(QueryStruct{}, nil)
	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var q queryCreateStruct
			common.DataToStructPointer(d, s, &q)
			apiQuery, err := w.Queries.Create(ctx, sql.CreateQueryRequest{
				AutoResolveDisplayName: false,
				Query:                  &q.CreateQueryRequestQuery,
				ForceSendFields:        []string{"AutoResolveDisplayName"},
			})
			if err != nil {
				return err
			}
			d.SetId(apiQuery.Id)
			owner := d.Get("owner_user_name").(string)
			if owner != "" {
				_, err = w.Queries.Update(ctx, sql.UpdateQueryRequest{
					Query: &sql.UpdateQueryRequestQuery{
						OwnerUserName: owner,
					},
					Id:         apiQuery.Id,
					UpdateMask: "owner_user_name",
				})
			}
			return err
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			apiQuery, err := w.Queries.GetById(ctx, d.Id())
			if err != nil {
				log.Printf("[WARN] error getting query by ID: %v", err)
				return err
			}
			parentPath := d.Get("parent_path").(string)
			if parentPath != "" && strings.HasPrefix(apiQuery.ParentPath, "/Workspace") && !strings.HasPrefix(parentPath, "/Workspace") {
				apiQuery.ParentPath = strings.TrimPrefix(parentPath, "/Workspace")
			}
			return common.StructToData(QueryStruct{Query: *apiQuery}, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var q queryUpdateStruct
			common.DataToStructPointer(d, s, &q)
			updateMask := "display_name,query_text,warehouse_id,parameters"
			for _, f := range []string{"run_as_mode", "owner_user_name", "description", "tags",
				"apply_auto_limit", "catalog", "schema"} {
				if d.HasChange(f) {
					updateMask += "," + f
				}
			}
			_, err = w.Queries.Update(ctx, sql.UpdateQueryRequest{
				Query:                  &q.UpdateQueryRequestQuery,
				Id:                     d.Id(),
				UpdateMask:             updateMask,
				AutoResolveDisplayName: false,
				ForceSendFields:        []string{"AutoResolveDisplayName"},
			})
			return err
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return w.Queries.DeleteById(ctx, d.Id())
		},
		Schema: s,
	}
}
