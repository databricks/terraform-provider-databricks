package main

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

// rawType is a tiny helper to build a cty type literal.
func rawType(t string) json.RawMessage { return json.RawMessage(`"` + t + `"`) }

// providerWith wraps a single resource into a full ProviderSchemas for the table-driven tests.
func providerWith(resource string, s Schema) *ProviderSchemas {
	return &ProviderSchemas{
		ProviderSchemas: map[string]ProviderSchema{
			"registry.terraform.io/databricks/databricks": {
				ResourceSchemas: map[string]Schema{resource: s},
			},
		},
	}
}

// schemaWith returns a single-attribute schema for table-driven tests.
func schemaWith(attr string, a Attribute) Schema {
	return Schema{Block: Block{Attributes: map[string]Attribute{attr: a}}}
}

func schemaWithBlock(blockName string, bt BlockType) Schema {
	return Schema{Block: Block{BlockTypes: map[string]BlockType{blockName: bt}}}
}

func TestClassify_RuleTaxonomy(t *testing.T) {
	tests := []struct {
		name        string
		base        *ProviderSchemas
		head        *ProviderSchemas
		wantKind    string // expected change kind to be present (empty if no changes expected)
		wantSev     Severity
		wantNoOther bool // when true, fail if any OTHER kind is detected
	}{
		// Resource-level
		{
			name:     "remove resource is breaking",
			base:     providerWith("databricks_foo", schemaWith("x", Attribute{Type: rawType("string"), Optional: true})),
			head:     providerWith("databricks_bar", schemaWith("x", Attribute{Type: rawType("string"), Optional: true})),
			wantKind: "ResourceRemoved",
			wantSev:  Breaking,
		},
		{
			name: "add resource is non-breaking",
			base: &ProviderSchemas{ProviderSchemas: map[string]ProviderSchema{
				"registry.terraform.io/databricks/databricks": {ResourceSchemas: map[string]Schema{}},
			}},
			head:     providerWith("databricks_foo", schemaWith("x", Attribute{Type: rawType("string"), Optional: true})),
			wantKind: "ResourceAdded",
			wantSev:  NonBreaking,
		},
		// Attribute add/remove
		{
			name:     "remove attribute is breaking",
			base:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Optional: true})),
			head:     providerWith("r", Schema{Block: Block{Attributes: map[string]Attribute{}}}),
			wantKind: "AttributeRemoved",
			wantSev:  Breaking,
		},
		{
			name:     "add required attribute is breaking",
			base:     providerWith("r", Schema{}),
			head:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Required: true})),
			wantKind: "RequiredAttributeAdded",
			wantSev:  Breaking,
		},
		{
			name:     "add optional attribute is non-breaking",
			base:     providerWith("r", Schema{}),
			head:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Optional: true})),
			wantKind: "OptionalAttributeAdded",
			wantSev:  NonBreaking,
		},
		{
			name:     "add computed-only attribute is non-breaking",
			base:     providerWith("r", Schema{}),
			head:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Computed: true})),
			wantKind: "ComputedAttributeAdded",
			wantSev:  NonBreaking,
		},
		// Required <-> Optional
		{
			name:     "optional to required is breaking",
			base:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Optional: true})),
			head:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Required: true})),
			wantKind: "OptionalToRequired",
			wantSev:  Breaking,
		},
		{
			name:     "required to optional is non-breaking",
			base:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Required: true})),
			head:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Optional: true})),
			wantKind: "RequiredToOptional",
			wantSev:  NonBreaking,
		},
		{
			name:     "computed-only to required is breaking",
			base:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Computed: true})),
			head:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Required: true})),
			wantKind: "ComputedOnlyToRequired",
			wantSev:  Breaking,
		},
		{
			name:     "computed-only to required+computed is breaking",
			base:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Computed: true})),
			head:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Required: true, Computed: true})),
			wantKind: "ComputedOnlyToRequired",
			wantSev:  Breaking,
		},
		// Computed transitions
		{
			name:     "adding computed to settable attr is non-breaking",
			base:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Optional: true})),
			head:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Optional: true, Computed: true})),
			wantKind: "ComputedAdded",
			wantSev:  NonBreaking,
		},
		{
			name:     "removing computed from settable attr is breaking",
			base:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Optional: true, Computed: true})),
			head:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Optional: true})),
			wantKind: "ComputedRemoved",
			wantSev:  Breaking,
		},
		{
			name:     "settable becomes computed-only is breaking",
			base:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Optional: true})),
			head:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Computed: true})),
			wantKind: "BecameComputedOnly",
			wantSev:  Breaking,
		},
		// Sensitive
		{
			name:     "adding sensitive is non-breaking",
			base:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Optional: true})),
			head:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Optional: true, Sensitive: true})),
			wantKind: "SensitiveAdded",
			wantSev:  NonBreaking,
		},
		{
			name:     "removing sensitive is breaking",
			base:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Optional: true, Sensitive: true})),
			head:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Optional: true})),
			wantKind: "SensitiveRemoved",
			wantSev:  Breaking,
		},
		// Type changes
		{
			name:     "type change (string to number) is breaking",
			base:     providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Optional: true})),
			head:     providerWith("r", schemaWith("a", Attribute{Type: rawType("number"), Optional: true})),
			wantKind: "TypeChanged",
			wantSev:  Breaking,
		},
		{
			name: "tighten element type (list string to list number) is breaking",
			base: providerWith("r", schemaWith("a", Attribute{
				Type: json.RawMessage(`["list","string"]`), Optional: true,
			})),
			head: providerWith("r", schemaWith("a", Attribute{
				Type: json.RawMessage(`["list","number"]`), Optional: true,
			})),
			wantKind: "TypeChanged",
			wantSev:  Breaking,
		},
		// Nested-type (Plugin Framework structured attribute) transitions
		{
			name: "adding optional sub-attribute to nested type is non-breaking",
			base: providerWith("r", schemaWith("a", Attribute{
				Optional: true,
				NestedType: &NestedType{
					NestingMode: "single",
					Attributes: map[string]Attribute{
						"x": {Type: rawType("string"), Optional: true},
					},
				},
			})),
			head: providerWith("r", schemaWith("a", Attribute{
				Optional: true,
				NestedType: &NestedType{
					NestingMode: "single",
					Attributes: map[string]Attribute{
						"x": {Type: rawType("string"), Optional: true},
						"y": {Type: rawType("string"), Optional: true},
					},
				},
			})),
			wantKind: "OptionalAttributeAdded",
			wantSev:  NonBreaking,
		},
		{
			name: "removing sub-attribute from nested type is breaking",
			base: providerWith("r", schemaWith("a", Attribute{
				Optional: true,
				NestedType: &NestedType{
					NestingMode: "single",
					Attributes: map[string]Attribute{
						"x": {Type: rawType("string"), Optional: true},
						"y": {Type: rawType("string"), Optional: true},
					},
				},
			})),
			head: providerWith("r", schemaWith("a", Attribute{
				Optional: true,
				NestedType: &NestedType{
					NestingMode: "single",
					Attributes: map[string]Attribute{
						"x": {Type: rawType("string"), Optional: true},
					},
				},
			})),
			wantKind: "AttributeRemoved",
			wantSev:  Breaking,
		},
		{
			name: "sub-attribute optional to required in nested type is breaking",
			base: providerWith("r", schemaWith("a", Attribute{
				Optional: true,
				NestedType: &NestedType{
					NestingMode: "single",
					Attributes: map[string]Attribute{
						"x": {Type: rawType("string"), Optional: true},
					},
				},
			})),
			head: providerWith("r", schemaWith("a", Attribute{
				Optional: true,
				NestedType: &NestedType{
					NestingMode: "single",
					Attributes: map[string]Attribute{
						"x": {Type: rawType("string"), Required: true},
					},
				},
			})),
			wantKind: "OptionalToRequired",
			wantSev:  Breaking,
		},
		{
			name: "nested-type nesting_mode change is breaking",
			base: providerWith("r", schemaWith("a", Attribute{
				Optional: true,
				NestedType: &NestedType{
					NestingMode: "single",
					Attributes:  map[string]Attribute{"x": {Type: rawType("string"), Optional: true}},
				},
			})),
			head: providerWith("r", schemaWith("a", Attribute{
				Optional: true,
				NestedType: &NestedType{
					NestingMode: "list",
					Attributes:  map[string]Attribute{"x": {Type: rawType("string"), Optional: true}},
				},
			})),
			wantKind: "NestingModeChanged",
			wantSev:  Breaking,
		},
		{
			name: "identical nested type emits no change",
			base: providerWith("r", schemaWith("a", Attribute{
				Optional: true,
				NestedType: &NestedType{
					NestingMode: "single",
					Attributes:  map[string]Attribute{"x": {Type: rawType("string"), Optional: true}},
				},
			})),
			head: providerWith("r", schemaWith("a", Attribute{
				Optional: true,
				NestedType: &NestedType{
					NestingMode: "single",
					Attributes:  map[string]Attribute{"x": {Type: rawType("string"), Optional: true}},
				},
			})),
			wantKind:    "",
			wantNoOther: true,
		},
		// Description-only change is silent
		{
			name: "description change is non-breaking and emits no change",
			base: providerWith("r", schemaWith("a", Attribute{
				Type: rawType("string"), Optional: true, Description: "old",
			})),
			head: providerWith("r", schemaWith("a", Attribute{
				Type: rawType("string"), Optional: true, Description: "new",
			})),
			wantKind:    "",
			wantNoOther: true,
		},
		// block_types
		{
			name:     "remove nested block is breaking",
			base:     providerWith("r", schemaWithBlock("b", BlockType{NestingMode: "list"})),
			head:     providerWith("r", Schema{}),
			wantKind: "BlockTypeRemoved",
			wantSev:  Breaking,
		},
		{
			name:     "add optional nested block is non-breaking",
			base:     providerWith("r", Schema{}),
			head:     providerWith("r", schemaWithBlock("b", BlockType{NestingMode: "list"})),
			wantKind: "BlockTypeAdded",
			wantSev:  NonBreaking,
		},
		{
			name:     "add required nested block (min_items>0) is breaking",
			base:     providerWith("r", Schema{}),
			head:     providerWith("r", schemaWithBlock("b", BlockType{NestingMode: "list", MinItems: 1})),
			wantKind: "RequiredBlockTypeAdded",
			wantSev:  Breaking,
		},
		{
			name:     "nesting mode change is breaking",
			base:     providerWith("r", schemaWithBlock("b", BlockType{NestingMode: "list"})),
			head:     providerWith("r", schemaWithBlock("b", BlockType{NestingMode: "set"})),
			wantKind: "NestingModeChanged",
			wantSev:  Breaking,
		},
		{
			name:     "min_items increase is breaking",
			base:     providerWith("r", schemaWithBlock("b", BlockType{NestingMode: "list", MinItems: 0})),
			head:     providerWith("r", schemaWithBlock("b", BlockType{NestingMode: "list", MinItems: 1})),
			wantKind: "MinItemsIncreased",
			wantSev:  Breaking,
		},
		{
			name:     "max_items decrease is breaking",
			base:     providerWith("r", schemaWithBlock("b", BlockType{NestingMode: "list", MaxItems: 5})),
			head:     providerWith("r", schemaWithBlock("b", BlockType{NestingMode: "list", MaxItems: 2})),
			wantKind: "MaxItemsDecreased",
			wantSev:  Breaking,
		},
		{
			name:     "max_items relaxed (5 to 10) is non-breaking",
			base:     providerWith("r", schemaWithBlock("b", BlockType{NestingMode: "list", MaxItems: 5})),
			head:     providerWith("r", schemaWithBlock("b", BlockType{NestingMode: "list", MaxItems: 10})),
			wantKind: "MaxItemsRelaxed",
			wantSev:  NonBreaking,
		},
		{
			name:     "max_items removed (capped to unbounded) is non-breaking",
			base:     providerWith("r", schemaWithBlock("b", BlockType{NestingMode: "list", MaxItems: 5})),
			head:     providerWith("r", schemaWithBlock("b", BlockType{NestingMode: "list"})),
			wantKind: "MaxItemsRelaxed",
			wantSev:  NonBreaking,
		},
		// No change
		{
			name:        "identical schemas produce no changes",
			base:        providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Optional: true})),
			head:        providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Optional: true})),
			wantKind:    "",
			wantNoOther: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Classify(tc.base, tc.head)
			if tc.wantKind == "" {
				if tc.wantNoOther && len(got) != 0 {
					t.Fatalf("expected no changes, got: %+v", got)
				}
				return
			}
			var found *Change
			for i := range got {
				if got[i].Kind == tc.wantKind {
					found = &got[i]
					break
				}
			}
			if found == nil {
				t.Fatalf("expected kind %q in changes, got: %+v", tc.wantKind, got)
			}
			if found.Severity != tc.wantSev {
				t.Errorf("kind %q severity = %s, want %s", tc.wantKind, found.Severity, tc.wantSev)
			}
		})
	}
}

func TestHasBreaking(t *testing.T) {
	cases := []struct {
		in   []Change
		want bool
	}{
		{nil, false},
		{[]Change{{Severity: NonBreaking}}, false},
		{[]Change{{Severity: NonBreaking}, {Severity: Breaking}}, true},
	}
	for _, c := range cases {
		if got := HasBreaking(c.in); got != c.want {
			t.Errorf("HasBreaking(%+v) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestWriteMarkdown_EscapesPipes(t *testing.T) {
	var sb strings.Builder
	writeMarkdown(&sb, []Change{{Path: "p", Kind: "K", Severity: Breaking, Message: "a|b"}})
	if !strings.Contains(sb.String(), `a\|b`) {
		t.Fatalf("pipe not escaped: %q", sb.String())
	}
}

func TestRun_ExitsOnBreaking(t *testing.T) {
	dir := t.TempDir()
	base := mustWrite(t, dir, "base.json", providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Optional: true})))
	head := mustWrite(t, dir, "head.json", providerWith("r", Schema{}))

	var out, errOut strings.Builder
	err := run([]string{"--base", base, "--head", head}, &out, &errOut)
	if _, ok := err.(errBreaking); !ok {
		t.Fatalf("expected errBreaking, got %v", err)
	}
	if !strings.Contains(out.String(), "AttributeRemoved") {
		t.Errorf("expected AttributeRemoved in output, got: %s", out.String())
	}
}

func TestRun_AllowBreakingFlag(t *testing.T) {
	dir := t.TempDir()
	base := mustWrite(t, dir, "base.json", providerWith("r", schemaWith("a", Attribute{Type: rawType("string"), Optional: true})))
	head := mustWrite(t, dir, "head.json", providerWith("r", Schema{}))

	var out, errOut strings.Builder
	err := run([]string{"--base", base, "--head", head, "--allow-breaking"}, &out, &errOut)
	if err != nil {
		t.Fatalf("expected nil error with --allow-breaking, got %v", err)
	}
}

func mustWrite(t *testing.T, dir, name string, ps *ProviderSchemas) string {
	t.Helper()
	path := dir + "/" + name
	b, err := json.Marshal(ps)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if err := os.WriteFile(path, b, 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	return path
}
