package exporter

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
	"sync"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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

func (ra *resourceApproximation) Get(attr string) (any, bool) {
	for _, i := range ra.Instances {
		v, found := i.Attributes[attr]
		if found {
			return v, found
		}
	}
	return nil, false
}

// TODO: think if something like trie may help here...
type resourceApproximationHolder struct {
	mutex      sync.RWMutex
	resources  []*resourceApproximation
	attributes map[string]*resourceApproximation
}

func makeMatchPair(k, v string) string {
	return k + "/" + v
}

func (rah *resourceApproximationHolder) Has(r *resource) bool {
	k, v := r.MatchPair()
	matchPairKey := makeMatchPair(k, v)
	// log.Printf("[DEBUG] resourceApproximationHolder.Has is called for %s/%s", k, v)
	rah.mutex.RLocker().Lock()
	defer rah.mutex.RLocker().Unlock()
	_, exists := rah.attributes[matchPairKey]
	// log.Printf("[DEBUG] resourceApproximationHolder.Has is finished. Resource found? %v", exists)
	return exists
}

func (rah *resourceApproximationHolder) Get(attr, value string) *resourceApproximation {
	matchPairKey := makeMatchPair(attr, value)
	rah.mutex.RLocker().Lock()
	defer rah.mutex.RLocker().Unlock()
	ra := rah.attributes[matchPairKey]
	return ra
}

func (rah *resourceApproximationHolder) Append(ra resourceApproximation) {
	// log.Printf("[DEBUG] resourceApproximationHolder.Append: %v. Instances count=%d", ra, len(ra.Instances))
	rah.mutex.Lock()
	defer rah.mutex.Unlock()
	rah.resources = append(rah.resources, &ra)
	for _, i := range ra.Instances {
		for k, v := range i.Attributes {
			tv, ok := v.(string)
			if ok {
				rah.attributes[makeMatchPair(k, tv)] = &ra
			}
		}
	}
}

type stateApproximation struct {
	rmap map[string]*resourceApproximationHolder
}

func newStateApproximation(suppported_resources []string) *stateApproximation {
	sa := stateApproximation{rmap: map[string]*resourceApproximationHolder{}}
	for _, k := range suppported_resources {
		sa.rmap[k] = &resourceApproximationHolder{attributes: map[string]*resourceApproximation{}}
	}
	return &sa
}

func (s *stateApproximation) Resources(resource_type string) *[]*resourceApproximation {
	rah := s.rmap[resource_type]
	if rah != nil {
		return &rah.resources
	}
	panic(fmt.Sprintf("There is no support for resource type %s", resource_type))
}

func (s *stateApproximation) Has(r *resource) bool {
	rah, exist := s.rmap[r.Resource]
	if !exist {
		panic(fmt.Sprintf("There is no support for resource type %s", r.Resource))
	}
	return rah.Has(r)
}

func (s *stateApproximation) Get(resource_type, attr, value string) *resourceApproximation {
	rah, exist := s.rmap[resource_type]
	if !exist {
		panic(fmt.Sprintf("There is no support for resource type %s", resource_type))
	}
	return rah.Get(attr, value)
}

func (s *stateApproximation) Append(ra resourceApproximation) {
	rah, exist := s.rmap[ra.Type]
	if !exist {
		panic(fmt.Sprintf("There is no support for resource type %s", ra.Type))
	}
	rah.Append(ra)
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
	// Defines if specific service is account level resource
	AccountLevel bool
	// Defines if specific service is workspace level resource
	WorkspaceLevel bool
}

type MatchType string

const (
	// MatchExact is to specify that whole value should match
	MatchExact = "exact"
	// MatchDefault - same meaning as MatchExact
	MatchDefault = ""
	// MatchCaseInsensitive is to specify that whole value should match, but case-insensitive
	MatchCaseInsensitive = "caseinsensitive"
	// MatchPrefix is to specify that prefix of value should match
	MatchPrefix = "prefix"
	// MatchLongestPrefix is to specify that prefix of value should match, and select the longest value from list of candidates
	MatchLongestPrefix = "longestprefix"
	// MatchRegexp is to specify that the group extracted from value should match
	MatchRegexp = "regexp"
)

type valueTransformFunc func(string) string
type isValidAproximationFunc func(ic *importContext, res *resource, sr *resourceApproximation, origPath string) bool

type reference struct {
	// path to a given field, like, `cluster_id`, `access_control.user_name``, ... For references blocks/arrays, the `.N` component isn't required
	Path string
	// resource type: databricks_directory, ...
	Resource string
	// on what field in the target resource to match
	Match string
	// type of match, `prefix` - reference is embedded into string, `` (or `exact`) - full match, ...
	MatchType MatchType
	// true if given reference denote a variable
	Variable bool
	// true if given reference denote a reference to a generated file
	File bool
	// regular expression (if MatchType == "regexp") must define a group that will be used to extract value to match
	Regexp *regexp.Regexp
	// functions to transform match and current search value
	MatchValueTransformFunc  valueTransformFunc
	SearchValueTransformFunc valueTransformFunc
	// function to evaluate fit of the resource approximation found to the resource...
	IsValidApproximation isValidAproximationFunc
	// if we should skip direct lookups (for example, we need it for UC schemas matching)
	SkipDirectLookup bool
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
	// Name of the resource: `databricks_cluster`, `databricks_job`, etc.
	Resource string
	// ID of the resource (could be omitted - then we need to perform search by Attribute/Value)
	ID string
	// Name of the attribute to search by
	Attribute string
	// Value of a given attribute to search by
	Value string
	// Terraform resource name
	Name string
	// If not specified, then we generate a normal resource block, or we can generate a data block if it's set to "data"
	Mode        string
	Incremental bool
	// Actual Terraform data
	Data *schema.ResourceData
	// Arbitrary data to be used by importable
	ExtraData map[string]any
	// References to dependencies - it could be fully resolved resource, with Data, etc., or it could be just resource type + ID
	DependsOn []*resource
}

func (r *resource) AddExtraData(key string, value any) {
	if r.ExtraData == nil {
		r.ExtraData = map[string]any{}
	}
	r.ExtraData[key] = value
}

func (r *resource) AddDependsOn(dep *resource) {
	r.DependsOn = append(r.DependsOn, dep)
}

func (r *resource) GetExtraData(key string) (any, bool) {
	if r.ExtraData == nil {
		return nil, false
	}
	v, ok := r.ExtraData[key]
	return v, ok
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

func (r *resource) ImportResource(ic *importContext) {
	defer ic.waitGroup.Done()
	pr, ok := ic.Resources[r.Resource]
	if !ok {
		log.Printf("[ERROR] %s is not available in provider", r)
		return
	}
	ir, ok := ic.Importables[r.Resource]
	if !ok {
		log.Printf("[ERROR] %s is not available for import", r)
		return
	}
	if ic.HasInState(r, true) {
		log.Printf("[DEBUG] %s already imported", r)
		return
	}

	if r.ID == "" {
		if ir.Search == nil {
			log.Printf("[ERROR] Searching %s is not available", r)
			return
		}
		err := runWithRetries(func() error {
			return ir.Search(ic, r)
		},
			fmt.Sprintf("searching of %v", r))
		if err != nil {
			log.Printf("[ERROR] Error searching %s#%s: %v", r.Resource, r.ID, err)
			return
		}
		if r.ID == "" {
			log.Printf("[WARN] Cannot find %s", r)
			ic.addIgnoredResource(fmt.Sprintf("%s: %s=%s", r.Resource, r.Attribute, r.Value))
			return
		}
	}
	if r.Data == nil {
		// empty data with resource schema
		r.Data = pr.Data(&terraform.InstanceState{
			Attributes: map[string]string{},
			ID:         r.ID,
		})
		r.Data.MarkNewResource()
		resource := strings.ReplaceAll(r.Resource, "databricks_", "")
		ctx := context.WithValue(ic.Context, common.ResourceName, resource)
		apiVersion := ic.Importables[r.Resource].ApiVersion
		if apiVersion != "" {
			ctx = context.WithValue(ctx, common.Api, apiVersion)
		}
		dia := runWithRetries(func() diag.Diagnostics {
			return pr.ReadContext(ctx, r.Data, ic.Client)
		},
			fmt.Sprintf("reading %s#%s", r.Resource, r.ID))
		if dia != nil {
			log.Printf("[ERROR] Error reading %s#%s: %v", r.Resource, r.ID, dia)
			return
		}
		if r.Data.Id() == "" {
			r.Data.SetId(r.ID)
		}
	}
	r.Name = ic.ResourceName(r)
	if ir.Import != nil {
		err := runWithRetries(func() error {
			return ir.Import(ic, r)
		},
			fmt.Sprintf("importing of %s#%s", r.Resource, r.ID))
		if err != nil {
			log.Printf("[ERROR] Failed custom import of %s: %s", r, err)
			return
		}
	}
	ic.Add(r)
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
	c := make(resourcesList, len(a.resources))
	copy(c, a.resources)
	sort.Sort(c)
	return c
}

func (a *importedResources) FindById(resourceType, id string) *resource {
	defer a.mutex.RLocker().Unlock()
	a.mutex.RLocker().Lock()
	for _, r := range a.resources {
		if r.Resource == resourceType && r.ID == id {
			return r
		}
	}

	return nil
}
