package exporter

import (
	"fmt"
	"log"

	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/common"
	tf_sharing "github.com/databricks/terraform-provider-databricks/sharing"
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
	resourceInfo := ic.Resources["databricks_share"]
	if resourceInfo == nil {
		// Fallback to direct data access if schema is not available
		objectsList := r.Data.Get("object").([]any)
		ic.emitUCGrantsWithOwner("share/"+r.ID, r)
		for _, objRaw := range objectsList {
			obj := objRaw.(map[string]any)
			dataObjectType := obj["data_object_type"].(string)
			name := obj["name"].(string)

			switch dataObjectType {
			case "TABLE":
				ic.Emit(&resource{
					Resource: "databricks_sql_table",
					ID:       name,
				})
			case "VOLUME":
				ic.Emit(&resource{
					Resource: "databricks_volume",
					ID:       name,
				})
			case "MODEL":
				ic.Emit(&resource{
					Resource: "databricks_registered_model",
					ID:       name,
				})
			default:
				log.Printf("[INFO] Object type '%s' (name: '%s') isn't supported in share '%s'",
					dataObjectType, name, r.ID)
			}
		}
		return nil
	}

	var share tf_sharing.ShareInfo
	s := resourceInfo.Schema
	common.DataToStructPointer(r.Data, s, &share)
	// TODO: how to link recipients to share?
	ic.emitUCGrantsWithOwner("share/"+r.ID, r)
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
	return nil
}
