package exporter

import (
	"fmt"
	"log"

	"github.com/databricks/databricks-sdk-go/service/sharing"
	pluginfw_sharing "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/sharing"
)

func listUcShares(ic *importContext) error {
	it := ic.workspaceClient.Shares.List(ic.Context, sharing.ListSharesRequest{})
	for it.HasNext(ic.Context) {
		share, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		ic.EmitIfUpdatedAfterMillisAndNameMatches(&resource{
			Resource: "databricks_share",
			ID:       share.Name,
		}, share.Name, share.UpdatedAt, fmt.Sprintf("share '%s'", share.Name))
	}
	return nil
}

func importUcShare(ic *importContext, r *resource) error {
	log.Printf("[DEBUG] Importing share: %s", r.ID)
	// Convert Plugin Framework state to Go SDK struct
	var share sharing.ShareInfo
	if err := convertPluginFrameworkToGoSdk(ic, r.DataWrapper,
		pluginfw_sharing.ShareInfoExtended{}, &share); err != nil {
		return err
	}

	// Emit UC grants with owner
	ic.emitUCGrantsWithOwner("share/"+r.ID, r)

	// Emit RFA access request destinations if configured
	ic.emitRfaAccessRequestDestinations("SHARE", r.ID)

	// Emit dependencies for each object in the share
	for _, obj := range share.Objects {
		switch obj.DataObjectType {
		case "TABLE":
			ic.Emit(&resource{
				Resource: "databricks_sql_table",
				ID:       obj.Name,
			})
		case "VOLUME":
			ic.Emit(&resource{
				Resource: "databricks_volume",
				ID:       obj.Name,
			})
		case "MODEL":
			ic.Emit(&resource{
				Resource: "databricks_registered_model",
				ID:       obj.Name,
			})
		case "SCHEMA":
			ic.Emit(&resource{
				Resource: "databricks_schema",
				ID:       obj.Name,
			})
		default:
			log.Printf("[INFO] Object type '%s' (name: '%s') isn't supported in share '%s'",
				obj.DataObjectType, obj.Name, r.ID)
		}
	}
	return nil
}

func listUcRecipients(ic *importContext) error {
	it := ic.workspaceClient.Recipients.List(ic.Context, sharing.ListRecipientsRequest{})
	for it.HasNext(ic.Context) {
		rec, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		ic.EmitIfUpdatedAfterMillisAndNameMatches(&resource{
			Resource: "databricks_recipient",
			ID:       rec.Name,
		}, rec.Name, rec.UpdatedAt, fmt.Sprintf("recipient '%s'", rec.Name))
	}
	return nil
}

func importUcRecipient(ic *importContext, r *resource) error {
	owner := r.Data.Get("owner").(string)
	if owner != "" {
		emitUserSpOrGroup(ic, owner)
	}

	// Emit RFA access request destinations if configured
	ic.emitRfaAccessRequestDestinations("RECIPIENT", r.ID)

	return nil
}

func listUcProviders(ic *importContext) error {
	it := ic.workspaceClient.Providers.List(ic.Context, sharing.ListProvidersRequest{})
	for it.HasNext(ic.Context) {
		provider, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		ic.EmitIfUpdatedAfterMillisAndNameMatches(&resource{
			Resource: "databricks_provider",
			ID:       provider.Name,
		}, provider.Name, provider.UpdatedAt, fmt.Sprintf("provider '%s'", provider.Name))
	}
	return nil
}

func importUcProvider(ic *importContext, r *resource) error {
	// Emit RFA access request destinations if configured
	ic.emitRfaAccessRequestDestinations("PROVIDER", r.ID)
	return nil
}
