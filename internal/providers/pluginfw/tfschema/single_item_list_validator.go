package tfschema

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// MaxItemsValidator is a custom validator that enforces the maximum number of items in a list.
func MaxItemsValidator(maxItems int) validator.List {
	return maxItemsValidator{
		maxItems: maxItems,
	}
}

type maxItemsValidator struct {
	maxItems int
}

// Description provides a description of the validator for Terraform documentation.
func (v maxItemsValidator) Description(ctx context.Context) string {
	return fmt.Sprintf("Validates that the list contains at most %d items.", v.maxItems)
}

// MarkdownDescription provides a markdown description for the Terraform documentation.
func (v maxItemsValidator) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("Validates that the list contains at most %d items.", v.maxItems)
}

// Validate ensures that the attribute plan contains at most the maximum allowed items.
func (v maxItemsValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	// Check if the value is a List type
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	listVal := req.ConfigValue

	// Check if the list exceeds the maximum allowed items
	if len(listVal.Elements()) > v.maxItems {
		resp.Diagnostics.AddError(
			"Too many items in the list",
			fmt.Sprintf("This list can contain at most %d items. Found %d items.", v.maxItems, len(listVal.Elements())),
		)
	}
}
