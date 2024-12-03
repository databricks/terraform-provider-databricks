package validators

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// Identical to objectvalidators.ExactlyOneOf() except that it doesn't count the current attribute.
// Great for migrating ExactlyOneOf and defining it on the containing schema.
type ExactlyOneAttributeValidator struct {
	paths path.Expressions
}

func ExactlyOneOf(expressions ...path.Expression) validator.Object {
	return ExactlyOneAttributeValidator{
		paths: expressions,
	}
}

// Description implements validator.Object.
func (e ExactlyOneAttributeValidator) Description(ctx context.Context) string {
	return e.MarkdownDescription(ctx)
}

// MarkdownDescription implements validator.Object.
func (e ExactlyOneAttributeValidator) MarkdownDescription(context.Context) string {
	return fmt.Sprintf("Requires exactly one of the following attributes on the specified object to be set: " + e.paths.String())
}

// ValidateObject implements validator.Object.
func (e ExactlyOneAttributeValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, res *validator.ObjectResponse) {
	count := 0
	expressions := req.PathExpression.MergeExpressions(e.paths...)

	// If current attribute is unknown, delay validation
	if req.ConfigValue.IsUnknown() {
		return
	}

	for _, expression := range expressions {
		matchedPaths, diags := req.Config.PathMatches(ctx, expression)

		res.Diagnostics.Append(diags...)

		// Collect all errors
		if diags.HasError() {
			continue
		}

		for _, mp := range matchedPaths {
			// If the user specifies the same attribute this validator is applied to,
			// also as part of the input, skip it
			if mp.Equal(req.Path) {
				continue
			}

			var mpVal attr.Value
			diags := req.Config.GetAttribute(ctx, mp, &mpVal)
			res.Diagnostics.Append(diags...)

			// Collect all errors
			if diags.HasError() {
				continue
			}

			// Delay validation until all involved attribute have a known value
			if mpVal.IsUnknown() {
				return
			}

			if !mpVal.IsNull() {
				count++
			}
		}
	}

	if count == 0 {
		res.Diagnostics.Append(validatordiag.InvalidAttributeCombinationDiagnostic(
			req.Path,
			fmt.Sprintf("No attribute specified when one (and only one) of %s is required", expressions),
		))
	}

	if count > 1 {
		res.Diagnostics.Append(validatordiag.InvalidAttributeCombinationDiagnostic(
			req.Path,
			fmt.Sprintf("%d attributes specified when one (and only one) of %s is required", count, expressions),
		))
	}
}

var _ validator.Object = ExactlyOneAttributeValidator{}
