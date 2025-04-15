// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package common

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// DiagToString converts a slice of diag.Diagnostics to a string.
func DiagToString(d diag.Diagnostics) string {
	b := strings.Builder{}
	for _, diag := range d {
		b.WriteString(fmt.Sprintf("[%s] %s: %s\n", diag.Severity(), diag.Summary(), diag.Detail()))
	}
	return b.String()
}
