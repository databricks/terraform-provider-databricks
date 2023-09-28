package exporter

import (
	"fmt"
	"regexp"
	"sort"
	"sync"

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
	mutex sync.RWMutex
	// TODO: use map by type -> should speedup Has function?
	resources []resourceApproximation
}

// TODO: check if it's used directly by multiple threads?
func (s *stateApproximation) Resources() []resourceApproximation {
	s.mutex.RLocker().Lock()
	defer s.mutex.RLocker().Unlock()
	c := make([]resourceApproximation, len(s.resources))
	copy(c, s.resources)
	return c
}

func (s *stateApproximation) Has(r *resource) bool {
	s.mutex.RLocker().Lock()
	defer s.mutex.RLocker().Unlock()
	k, v := r.MatchPair()
	for _, sr := range s.resources {
		if sr.Type != r.Resource {
			continue
		}
		for _, i := range sr.Instances {
			tv, ok := i.Attributes[k].(string)
			if ok && tv == v {
				return true
			}
		}
	}
	return false
}

func (s *stateApproximation) Append(ra resourceApproximation) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.resources = append(s.resources, ra)
}

type importable struct {
	// Logical (file) group that resources belong to
	Service string
	// Semantic resource block name
	Name func(ic *importContext, d *schema.ResourceData) string
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
	// Defines if specific service is account level
	AccountLevel bool
}

type MatchType string

const (
	// MatchExact is to specify that whole value should match
	MatchExact = "exact"
	// MatchDefault - same meaning as MatchExact
	MatchDefault = ""
	// MatchPrefix is to specify that prefix of value should match
	MatchPrefix = "prefix"
	// MatchRegexp is to specify that the group extracted from value should match
	MatchRegexp = "regexp"
)

type reference struct {
	Path      string
	Resource  string
	Match     string
	MatchType MatchType // type of match, `prefix` - reference is embedded into string, `` (or `exact`) - full match, ...
	Variable  bool
	File      bool
	Regexp    *regexp.Regexp // regular expression must define a group that will be used to extract value to match
}

func (r reference) MatchAttribute() string {
	if r.Match != "" {
		return r.Match
	}
	return "id"
}

func (r reference) MatchTypeValue() MatchType {
	if r.MatchType == "" {
		return MatchExact
	}
	return r.MatchType
}

type resource struct {
	Resource    string
	ID          string
	Attribute   string
	Value       string
	Name        string
	Mode        string
	Incremental bool
	Data        *schema.ResourceData
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

// TODO: split resources into a map of resource type -> list of resources (guarded by RW locks)
type resourcesList []*resource

type importedResources struct {
	resources resourcesList
	mutex     sync.RWMutex
}

func (a *importedResources) Append(r *resource) {
	defer a.mutex.Unlock()
	a.mutex.Lock()
	a.resources = append(a.resources, r)
}

func (a *importedResources) Len() int {
	defer a.mutex.RLocker().Unlock()
	a.mutex.RLocker().Lock()
	return len(a.resources)
}

func (a resourcesList) Len() int {
	return len(a)
}

func (a resourcesList) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a resourcesList) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func (a *importedResources) Sorted() []*resource {
	defer a.mutex.Unlock()
	a.mutex.Lock()
	// TODO: make a copy...
	sort.Sort(a.resources)
	return a.resources
}
