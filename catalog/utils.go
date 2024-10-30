package catalog

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// UC catalog resources accept an optional metastore_id parameter. This is required for account-level operations, but it is not used
// for workspace-level. However, to avoid confusion for customers when they specify an id at workspace-level, we validate that the
// id matches the metastore assigned to the workspace.
// This must be done for all operations that modify the resource (create, update, delete) but not when reading the resource to avoid
// breaking the terraform resource if the current compute value is not correct (due to this validation not exiting previously).
func validateMetastoreId(ctx context.Context, w *databricks.WorkspaceClient, metastoreId string) error {
	if metastoreId == "" {
		return nil
	}
	cat, err := w.Metastores.Current(ctx)
	if err != nil {
		return err
	}
	if cat.MetastoreId != metastoreId {
		return fmt.Errorf("metastore_id must be empty or equal to the metastore id assigned to the workspace: %s. "+
			"If the metastore assigned to the workspace has changed, the new metastore id must be explicitly set", cat.MetastoreId)
	}
	return nil
}

// Given a slice of SqlColumnInfo, construct the corresponding HCL string.
func GetSqlColumnInfoHCL(columnInfos []SqlColumnInfo) string {
	columnsTemplate := ""

	for _, ci := range columnInfos {
		ciTemplate := fmt.Sprintf(
			`
			column {
				name      = "%s"
				type      = "%s"
				nullable  = %t
				comment   = "%s"
			}
			`, ci.Name, ci.Type, ci.Nullable, ci.Comment,
		)
		columnsTemplate += ciTemplate
	}
	return columnsTemplate
}

// Given a slice of SqlKeyConstraintInfo, construct the corresponding HCL string.
func GetSqlKeyConstraintInfoHCL(keyConstraintInfos []SqlKeyConstraintInfo) string {
	keyConstraintsTemplate := ""

	for _, kci := range keyConstraintInfos {
		switch kci.SqlKeyConstraint.(type) {
		case SqlPrimaryKeyConstraint:
			pkci := kci.SqlKeyConstraint.(SqlPrimaryKeyConstraint)
			pkciTemplate := fmt.Sprintf(
				`
			{
				key_constraint {
					name      = "%s"
					primary_key      = "%s"
					rely  = %t
				}
			}
			`, pkci.Name, pkci.PrimaryKey, pkci.Rely,
			)
			keyConstraintsTemplate += pkciTemplate
		case SqlForeignKeyConstraint:
			fkci := kci.SqlKeyConstraint.(SqlForeignKeyConstraint)
			fkciTemplate := fmt.Sprintf(
				`
			{
				key_constraint {
					name      = "%s"
					referenced_key      = "%s"
					referenced_catalog  = "%s"
					referenced_schema  = "%s"
					referenced_table  = "%s"
					referenced_foreign_key  = "%s"
				}
			}
			`, fkci.Name, fkci.ReferencedKey, fkci.ReferencedCatalog, fkci.ReferencedSchema, fkci.ReferencedTable, fkci.ReferencedForeignKey,
			)
			keyConstraintsTemplate += fkciTemplate
		}
	}
	return keyConstraintsTemplate
}

// check if a UC resource needs the additional update call during create operation
func updateRequired(d *schema.ResourceData, updateOnly []string) bool {
	for _, key := range updateOnly {
		if d.Get(key) != "" {
			return true
		}
	}
	return false
}
