package catalog

import (
	"context"
	"sort"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type recipientsData struct {
	Name   string  `json:"name"`
	Latest bool    `json:"latest,omitempty" tf:"default:false"`
	Tokens []Token `json:"tokens,omitempty" tf:"computed"`
}

func DataSourceRecipientTokens() *schema.Resource {

	return common.DataResource(recipientsData{}, func(ctx context.Context, e interface{}, c *common.DatabricksClient) error {
		data := e.(*recipientsData)
		recipientsAPI := NewRecipientsAPI(ctx, c)
		recipient, err := recipientsAPI.getRecipient(data.Name)
		if err != nil {
			return err
		}
		// sort the tokens descending order
		sort.Slice(recipient.Tokens, func(i, j int) bool {
			return recipient.Tokens[i].CreatedAt > recipient.Tokens[j].CreatedAt
		})
		// get latest token
		if data.Latest && len(recipient.Tokens) > 0 {
			data.Tokens = append(data.Tokens, recipient.Tokens[0])
			// break off
			return nil
		}
		data.Tokens = append(data.Tokens, recipient.Tokens...)
		return nil
	})
}
