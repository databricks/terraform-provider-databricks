package importer

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type regexFix struct {
	Regex       *regexp.Regexp
	Replacement string
}

type instanceApproximation struct {
	// not really interested in other than strings...
	Attributes map[string]interface{} `json:"attributes"`
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
	Name    func(d *schema.ResourceData) string
	List    func(ic *importContext) error
	Search  func(ic *importContext, r *resource) error
	Import  func(ic *importContext, d *schema.ResourceData) error
	Depends []reference
	Service string
	Body    func(ic *importContext, body *hclwrite.Body, r *resource) error
}

type reference struct {
	Path     string
	Resource string
	Match    string
	Pick     string
}

type resource struct {
	Resource  string
	ID        string
	Attribute string
	Value     string
	Name      string
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