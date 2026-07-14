package exporter

import (
	"fmt"
	"log"

	"github.com/databricks/databricks-sdk-go/service/database"
	database_instance_resource "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/database_instance"
)

func listDatabaseInstances(ic *importContext) error {
	instances, err := ic.workspaceClient.Database.ListDatabaseInstancesAll(ic.Context, database.ListDatabaseInstancesRequest{})
	if err != nil {
		return err
	}
	i := 0
	for _, instance := range instances {
		if !ic.MatchesName(instance.Name) {
			log.Printf("[INFO] Skipping database instance %s because it doesn't match %s", instance.Name, ic.match)
			continue
		}
		ic.EmitIfUpdatedAfterMillis(&resource{
			Resource: "databricks_database_instance",
			ID:       instance.Name,
		}, 0, fmt.Sprintf("database instance '%s'", instance.Name))
		i++
	}
	if i > 0 {
		log.Printf("[INFO] Scanned %d Database Instances", i)
	}
	return nil
}

func importDatabaseInstance(ic *importContext, r *resource) error {
	// Copy values from effective_* fields to their input counterparts using converter-based approach
	// This works by:
	// 1. Converting TF state to Go SDK struct
	// 2. Copying effective_* fields to input fields using reflection
	// 3. Converting back to TF state
	// This automatically handles all types (simple and complex) including custom_tags!
	copyEffectiveFieldsToInputFieldsWithConverters[database_instance_resource.DatabaseInstance](
		ic, r, database.DatabaseInstance{})

	// Emit permissions for the database instance
	ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/database-instances/%s", r.ID),
		"database_instance_"+r.Name)
	return nil
}
