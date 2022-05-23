package catalog

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceRecipients() *schema.Resource {
	type recipientsData struct {
		Recipients []string `json:"recipients,omitempty" tf:"computed,slice_set"`
	}
	return common.DataResource(recipientsData{}, func(ctx context.Context, e interface{}, c *common.DatabricksClient) error {
		data := e.(*recipientsData)
		recipientsAPI := NewRecipientsAPI(ctx, c)
		recipients, err := recipientsAPI.list()
		if err != nil {
			return err
		}
		for _, recipient := range recipients.Recipients {
			data.Recipients = append(data.Recipients, recipient.Name)
		}
		return nil
	})
}
