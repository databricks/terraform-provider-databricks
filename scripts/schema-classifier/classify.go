// Rule taxonomy: see scripts/schema-classifier/README.md for the full
// human-readable table of every rule, its verdict, and the rationale.
// classify_test.go pins each rule with a passing test row.

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
)

type Severity int

const (
	NonBreaking Severity = iota
	Breaking
)

func (s Severity) String() string {
	if s == Breaking {
		return "BREAKING"
	}
	return "non-breaking"
}

type Change struct {
	Path     string
	Kind     string
	Severity Severity
	Message  string
}

// Classify walks base and head provider schemas and returns all detected
// changes, ordered by path. Severity == Breaking is what gates the CI check.
func Classify(base, head *ProviderSchemas) []Change {
	var out []Change
	baseProvs := keys(base.ProviderSchemas)
	headProvs := keys(head.ProviderSchemas)
	for _, p := range union(baseProvs, headProvs) {
		b, bok := base.ProviderSchemas[p]
		h, hok := head.ProviderSchemas[p]
		switch {
		case !bok:
			out = append(out, Change{Path: p, Kind: "ProviderAdded", Severity: NonBreaking,
				Message: fmt.Sprintf("Provider %q added", p)})
		case !hok:
			out = append(out, Change{Path: p, Kind: "ProviderRemoved", Severity: Breaking,
				Message: fmt.Sprintf("Provider %q removed", p)})
		default:
			out = append(out, diffProvider(p, b, h)...)
		}
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].Path < out[j].Path })
	return out
}

func diffProvider(prov string, base, head ProviderSchema) []Change {
	var out []Change
	out = append(out, diffSchema(prov+".provider", base.Provider, head.Provider)...)
	out = append(out, diffSchemaMap(prov+".resource", base.ResourceSchemas, head.ResourceSchemas, "Resource")...)
	out = append(out, diffSchemaMap(prov+".data_source", base.DataSourceSchemas, head.DataSourceSchemas, "DataSource")...)
	return out
}

func diffSchemaMap(prefix string, base, head map[string]Schema, kindLabel string) []Change {
	var out []Change
	for _, name := range union(keys(base), keys(head)) {
		path := prefix + "." + name
		b, bok := base[name]
		h, hok := head[name]
		switch {
		case !bok:
			out = append(out, Change{Path: path, Kind: kindLabel + "Added", Severity: NonBreaking,
				Message: fmt.Sprintf("%s %q added", kindLabel, name)})
		case !hok:
			out = append(out, Change{Path: path, Kind: kindLabel + "Removed", Severity: Breaking,
				Message: fmt.Sprintf("%s %q removed", kindLabel, name)})
		default:
			out = append(out, diffSchema(path, b, h)...)
		}
	}
	return out
}

func diffSchema(path string, base, head Schema) []Change {
	return diffBlock(path, base.Block, head.Block)
}

func diffBlock(path string, base, head Block) []Change {
	var out []Change
	out = append(out, diffAttributes(path, base.Attributes, head.Attributes)...)
	out = append(out, diffBlockTypes(path, base.BlockTypes, head.BlockTypes)...)
	return out
}

func diffAttributes(path string, base, head map[string]Attribute) []Change {
	var out []Change
	for _, name := range union(keys(base), keys(head)) {
		p := path + ".attributes." + name
		b, bok := base[name]
		h, hok := head[name]
		switch {
		case !bok:
			out = append(out, classifyAttributeAdded(p, name, h))
		case !hok:
			out = append(out, Change{Path: p, Kind: "AttributeRemoved", Severity: Breaking,
				Message: fmt.Sprintf("Attribute %q removed", name)})
		default:
			out = append(out, classifyAttributeChanged(p, name, b, h)...)
		}
	}
	return out
}

func classifyAttributeAdded(path, name string, a Attribute) Change {
	switch {
	case a.Required:
		return Change{Path: path, Kind: "RequiredAttributeAdded", Severity: Breaking,
			Message: fmt.Sprintf("New required attribute %q added (existing configs become invalid)", name)}
	case a.IsComputedOnly():
		return Change{Path: path, Kind: "ComputedAttributeAdded", Severity: NonBreaking,
			Message: fmt.Sprintf("New computed-only attribute %q added", name)}
	default:
		return Change{Path: path, Kind: "OptionalAttributeAdded", Severity: NonBreaking,
			Message: fmt.Sprintf("New optional attribute %q added", name)}
	}
}

func classifyAttributeChanged(path, name string, base, head Attribute) []Change {
	var out []Change

	if !jsonEqual(base.Type, head.Type) {
		out = append(out, Change{Path: path, Kind: "TypeChanged", Severity: Breaking,
			Message: fmt.Sprintf("Attribute %q type changed (any type delta is breaking)", name)})
	}

	// `nested_type` is a structured object — recurse into its attributes instead
	// of comparing raw JSON. A naive equality check would mis-flag benign deltas
	// (e.g. adding an optional sub-attribute) as TypeChanged.
	out = append(out, diffNestedType(path, name, base.NestedType, head.NestedType)...)

	switch {
	case base.Optional && head.Required:
		out = append(out, Change{Path: path, Kind: "OptionalToRequired", Severity: Breaking,
			Message: fmt.Sprintf("Attribute %q became required (existing configs missing it now error)", name)})
	case base.IsComputedOnly() && head.Required:
		out = append(out, Change{Path: path, Kind: "ComputedOnlyToRequired", Severity: Breaking,
			Message: fmt.Sprintf("Attribute %q was computed-only and is now required (users must now set it; existing configs error)", name)})
	case base.Required && head.Optional:
		out = append(out, Change{Path: path, Kind: "RequiredToOptional", Severity: NonBreaking,
			Message: fmt.Sprintf("Attribute %q became optional (strictly looser)", name)})
	}

	switch {
	case !base.Computed && head.Computed:
		out = append(out, Change{Path: path, Kind: "ComputedAdded", Severity: NonBreaking,
			Message: fmt.Sprintf("Attribute %q became computed", name)})
	case base.Computed && !head.Computed && base.IsSettable() && head.IsSettable():
		out = append(out, Change{Path: path, Kind: "ComputedRemoved", Severity: Breaking,
			Message: fmt.Sprintf("Attribute %q is no longer computed (drift behavior changes)", name)})
	}

	// Settable -> computed-only is effectively a removal of the user-facing attribute.
	if base.IsSettable() && head.IsComputedOnly() {
		out = append(out, Change{Path: path, Kind: "BecameComputedOnly", Severity: Breaking,
			Message: fmt.Sprintf("Attribute %q is no longer user-settable", name)})
	}

	switch {
	case !base.Sensitive && head.Sensitive:
		out = append(out, Change{Path: path, Kind: "SensitiveAdded", Severity: NonBreaking,
			Message: fmt.Sprintf("Attribute %q became sensitive (output now masked)", name)})
	case base.Sensitive && !head.Sensitive:
		out = append(out, Change{Path: path, Kind: "SensitiveRemoved", Severity: Breaking,
			Message: fmt.Sprintf("Attribute %q is no longer sensitive (un-masks previously hidden values)", name)})
	}

	return out
}

// diffNestedType compares two `nested_type` blocks structurally — like diffBlock
// for BlockTypes, but for Plugin Framework's NestedType (no child block_types).
// Recurses into the attributes so that benign sub-attribute changes (e.g.,
// adding an optional sub-attribute) are classified by their own rules instead
// of being lumped under TypeChanged.
func diffNestedType(path, name string, base, head *NestedType) []Change {
	if base == nil && head == nil {
		return nil
	}
	if base == nil || head == nil {
		return []Change{{
			Path: path, Kind: "TypeChanged", Severity: Breaking,
			Message: fmt.Sprintf("Attribute %q nested-type shape changed (nested_type added or removed)", name),
		}}
	}
	var out []Change
	if base.NestingMode != head.NestingMode {
		out = append(out, Change{Path: path, Kind: "NestingModeChanged", Severity: Breaking,
			Message: fmt.Sprintf("Attribute %q nesting_mode changed: %s -> %s", name, base.NestingMode, head.NestingMode)})
	}
	if head.MinItems > base.MinItems {
		out = append(out, Change{Path: path, Kind: "MinItemsIncreased", Severity: Breaking,
			Message: fmt.Sprintf("Attribute %q min_items increased: %d -> %d", name, base.MinItems, head.MinItems)})
	} else if head.MinItems < base.MinItems {
		out = append(out, Change{Path: path, Kind: "MinItemsDecreased", Severity: NonBreaking,
			Message: fmt.Sprintf("Attribute %q min_items decreased: %d -> %d", name, base.MinItems, head.MinItems)})
	}
	baseMax, headMax := base.MaxItems, head.MaxItems
	if baseMax != 0 && (headMax == 0 || headMax > baseMax) {
		out = append(out, Change{Path: path, Kind: "MaxItemsRelaxed", Severity: NonBreaking,
			Message: fmt.Sprintf("Attribute %q max_items relaxed: %d -> %d", name, baseMax, headMax)})
	} else if headMax != 0 && (baseMax == 0 || headMax < baseMax) {
		out = append(out, Change{Path: path, Kind: "MaxItemsDecreased", Severity: Breaking,
			Message: fmt.Sprintf("Attribute %q max_items decreased: %d -> %d", name, baseMax, headMax)})
	}
	// Recurse into the nested attributes. Sub-attribute add/remove/change is
	// classified by the same rules as top-level attributes.
	out = append(out, diffAttributes(path, base.Attributes, head.Attributes)...)
	return out
}

func diffBlockTypes(path string, base, head map[string]BlockType) []Change {
	var out []Change
	for _, name := range union(keys(base), keys(head)) {
		p := path + ".block_types." + name
		b, bok := base[name]
		h, hok := head[name]
		switch {
		case !bok:
			sev := NonBreaking
			msg := fmt.Sprintf("New optional nested block %q added", name)
			kind := "BlockTypeAdded"
			if h.MinItems > 0 {
				sev = Breaking
				msg = fmt.Sprintf("New required nested block %q added (min_items=%d)", name, h.MinItems)
				kind = "RequiredBlockTypeAdded"
			}
			out = append(out, Change{Path: p, Kind: kind, Severity: sev, Message: msg})
		case !hok:
			out = append(out, Change{Path: p, Kind: "BlockTypeRemoved", Severity: Breaking,
				Message: fmt.Sprintf("Nested block %q removed", name)})
		default:
			out = append(out, classifyBlockTypeChanged(p, name, b, h)...)
		}
	}
	return out
}

func classifyBlockTypeChanged(path, name string, base, head BlockType) []Change {
	var out []Change
	if base.NestingMode != head.NestingMode {
		out = append(out, Change{Path: path, Kind: "NestingModeChanged", Severity: Breaking,
			Message: fmt.Sprintf("Nested block %q nesting_mode changed: %s -> %s", name, base.NestingMode, head.NestingMode)})
	}
	if head.MinItems > base.MinItems {
		out = append(out, Change{Path: path, Kind: "MinItemsIncreased", Severity: Breaking,
			Message: fmt.Sprintf("Nested block %q min_items increased: %d -> %d", name, base.MinItems, head.MinItems)})
	} else if head.MinItems < base.MinItems {
		out = append(out, Change{Path: path, Kind: "MinItemsDecreased", Severity: NonBreaking,
			Message: fmt.Sprintf("Nested block %q min_items decreased: %d -> %d", name, base.MinItems, head.MinItems)})
	}
	// max_items==0 means "unbounded" in this schema dump.
	baseMax, headMax := base.MaxItems, head.MaxItems
	if baseMax != 0 && (headMax == 0 || headMax > baseMax) {
		out = append(out, Change{Path: path, Kind: "MaxItemsRelaxed", Severity: NonBreaking,
			Message: fmt.Sprintf("Nested block %q max_items relaxed: %d -> %d", name, baseMax, headMax)})
	} else if headMax != 0 && (baseMax == 0 || headMax < baseMax) {
		out = append(out, Change{Path: path, Kind: "MaxItemsDecreased", Severity: Breaking,
			Message: fmt.Sprintf("Nested block %q max_items decreased: %d -> %d", name, baseMax, headMax)})
	}
	out = append(out, diffBlock(path, base.Block, head.Block)...)
	return out
}

func jsonEqual(a, b json.RawMessage) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	var av, bv any
	if err := json.Unmarshal(a, &av); err != nil {
		return false
	}
	if err := json.Unmarshal(b, &bv); err != nil {
		return false
	}
	return reflect.DeepEqual(av, bv)
}

func keys[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func union(a, b []string) []string {
	seen := make(map[string]struct{}, len(a)+len(b))
	for _, s := range a {
		seen[s] = struct{}{}
	}
	for _, s := range b {
		seen[s] = struct{}{}
	}
	out := make([]string, 0, len(seen))
	for s := range seen {
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}

// HasBreaking returns true when any change in the slice is BREAKING.
func HasBreaking(changes []Change) bool {
	for _, c := range changes {
		if c.Severity == Breaking {
			return true
		}
	}
	return false
}
