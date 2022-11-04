package exporter

import (
	"fmt"
	"regexp"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type regexFix struct {
	Regex       *regexp.Regexp
	Replacement string
}

type instanceApproximation struct {
	// not really interested in other than strings...
	Attributes map[string]any `json:"attributes"`
}

type resourceApproximation struct {
	Type      string                  `json:"type"`
	Name      string                  `json:"name"`
	Provider  string                  `json:"provider"`
	Mode      string                  `json:"mode"`
	Module    string                  `json:"module,omitempty"`
	Instances []instanceApproximation `json:"instances"`
}

type stateApproximation struct {
	Resources []resourceApproximation `json:"resources"`
}

type importable struct {
	// Logical (file) group that resources belong to
	Service string
	// Semantic resource block name
	Name func(d *schema.ResourceData) string
	// Method to perform depth-first search and emit resources
	List func(ic *importContext) error
	// Search resource by non-ID attribute
	Search func(ic *importContext, r *resource) error
	// Emit additional resources after Resource.ReadContext is called
	Import func(ic *importContext, r *resource) error
	// Define logical dependencies between resources
	Depends []reference
	// Custom HCL writer for resource body
	Body func(ic *importContext, body *hclwrite.Body, r *resource) error
	// Function to detect if the given resource should be ignored or not
	Ignore func(ic *importContext, r *resource) bool
	// Function to check if the field in the given resource should be omitted or not
	ShouldOmitField func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool
	// Defines which API version should be used for this specific resource
	ApiVersion common.ApiVersion
}

type reference struct {
	Path     string
	Resource string
	Match    string
	Variable bool
	File     bool
}

type resource struct {
	Resource  string
	ID        string
	Attribute string
	Value     string
	Name      string
	Mode      string
	Data      *schema.ResourceData
}

func (r *resource) MatchPair() (string, string) {
	k := "id"
	v := r.ID
	if r.ID == "" && r.Attribute != "" && r.Value != "" {
		k = r.Attribute
		v = r.Value
	}
	return k, v
}

// String doubles its use as "hashcode"
func (r *resource) String() string {
	n := r.Name
	if n == "" {
		n = "<unknown>"
	}
	k, v := r.MatchPair()
	return fmt.Sprintf("%s[%s] (%s: %s)", r.Resource, n, k, v)
}

func (r *resource) ImportCommand(ic *importContext) string {
	m := ""
	if ic.Module != "" {
		m = ic.Module + "."
	}
	return fmt.Sprintf(`terraform import %s%s.%s "%s"`, m, r.Resource, r.Name, r.ID)
}

type importedResources []*resource

func (a importedResources) Len() int {
	return len(a)
}
func (a importedResources) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a importedResources) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}
