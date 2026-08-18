package exporter

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zclconf/go-cty/cty"
)

func TestMaybeAddQuoteCharacter(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{`plain`, `plain`},
		{`with"quote`, `with\"quote`},
		{`with\backslash`, `with\\backslash`},
		{`inter${polation}`, `inter$${polation}`},
		{`template%{if}`, `template%%{if}`},
		{`a"b\c${d}%{e}`, `a\"b\\c$${d}%%{e}`},
	}
	for _, c := range cases {
		assert.Equal(t, c.want, maybeAddQuoteCharacter(c.in))
	}
}

// TestReferenceFileHclInjection is a regression test for the HCL injection
// sibling of SEC-21613 (H1-3750315): a file name derived from an attacker
// controlled object path flows into a `File` reference and is emitted into a
// generated *.tf file. The value must be escaped so it cannot break out of the
// string literal and inject arbitrary HCL, while `${path.module}` must remain a
// real interpolation.
func TestReferenceFileHclInjection(t *testing.T) {
	ic := &importContext{}
	i := importable{
		Depends: []reference{{Path: "source", File: true}},
	}

	// A benign file name still renders as a path.module-relative reference.
	benign := "dbfs_files/_0cc175b9c0f1b6a831c399e269772661_a"
	benignTokens := ic.reference(i, []string{"source"}, benign, cty.StringVal(benign), &resource{})
	assert.Equal(t, `"${path.module}/dbfs_files/_0cc175b9c0f1b6a831c399e269772661_a"`,
		renderTokens(t, benignTokens))

	// A malicious file name attempting to break out of the string literal and
	// inject a resource block must be neutralized.
	malicious := "dbfs_files/_md5_evil\"}\nresource \"null_resource\" \"x\" {}\n#"
	tokens := ic.reference(i, []string{"source"}, malicious, cty.StringVal(malicious), &resource{})
	rendered := renderTokens(t, tokens)

	// The embedded quote must be escaped, so no bare `"}` closes the string.
	assert.Contains(t, rendered, `\"}`)
	// path.module interpolation is preserved.
	assert.True(t, strings.HasPrefix(rendered, `"${path.module}/`))

	// Most importantly: parsing an attribute built from these tokens must not
	// yield any injected block.
	f := hclwrite.NewEmptyFile()
	f.Body().SetAttributeRaw("source", tokens)
	parsed, diags := hclwrite.ParseConfig(f.Bytes(), "test.tf", hcl.Pos{Line: 1, Column: 1})
	require.False(t, diags.HasErrors(), "generated HCL failed to parse: %s", diags.Error())
	assert.Empty(t, parsed.Body().Blocks(),
		"HCL injection succeeded: an unexpected block was generated from the file name")
	assert.NotNil(t, parsed.Body().GetAttribute("source"))
}

// TestReferenceFileHclInterpolationInjection ensures a file name that tries to
// inject an HCL interpolation (`${...}`) is escaped and rendered literally.
func TestReferenceFileHclInterpolationInjection(t *testing.T) {
	ic := &importContext{}
	i := importable{Depends: []reference{{Path: "source", File: true}}}

	malicious := `dbfs_files/${file("/etc/passwd")}`
	tokens := ic.reference(i, []string{"source"}, malicious, cty.StringVal(malicious), &resource{})
	rendered := renderTokens(t, tokens)

	// The injected interpolation must be escaped to a literal `$${`.
	assert.Contains(t, rendered, `$${file(`)
	assert.NotContains(t, rendered, `/${file(`)
	// The intended path.module interpolation is still present exactly once.
	assert.Equal(t, 1, strings.Count(rendered, "${path.module}"))
}

// TestReferenceVariableRegexpSubstitution verifies that a regexp `Variable`
// reference substitutes only the captured substring with a shared variable and
// registers that variable with the captured value as its default.
func TestReferenceVariableRegexpSubstitution(t *testing.T) {
	ic := &importContext{}
	i := importable{
		Depends: []reference{
			{Path: "parent", MatchType: MatchRegexp, Regexp: regexp.MustCompile(`^accounts/([^/]+)$`),
				Variable: true, VariableName: accountIdVariableName},
		},
	}
	value := "accounts/abc-123"
	tokens := ic.reference(i, []string{"parent"}, value, cty.StringVal(value), &resource{})
	assert.Equal(t, `"accounts/${var.databricks_account_id}"`, renderTokens(t, tokens))
	assert.Contains(t, ic.variables, accountIdVariableName)
	assert.Equal(t, "abc-123", ic.variableDefaults[accountIdVariableName])
}

// TestReferenceVariableRegexpWithSuffix verifies substitution keeps text on both
// sides of the captured group.
func TestReferenceVariableRegexpWithSuffix(t *testing.T) {
	ic := &importContext{}
	i := importable{
		Depends: []reference{
			{Path: "name", MatchType: MatchRegexp, Regexp: regexp.MustCompile(`^accounts/([^/]+)/.*$`),
				Variable: true, VariableName: accountIdVariableName},
		},
	}
	value := "accounts/abc-123/ruleSets/default"
	tokens := ic.reference(i, []string{"name"}, value, cty.StringVal(value), &resource{})
	assert.Equal(t, `"accounts/${var.databricks_account_id}/ruleSets/default"`, renderTokens(t, tokens))
	assert.Equal(t, "abc-123", ic.variableDefaults[accountIdVariableName])
}

// TestReferenceVariableRegexpNoMatchEmitsLiteral verifies that when the regexp
// doesn't match, the literal value is emitted and no variable is registered.
func TestReferenceVariableRegexpNoMatchEmitsLiteral(t *testing.T) {
	ic := &importContext{}
	i := importable{
		Depends: []reference{
			{Path: "parent", MatchType: MatchRegexp, Regexp: regexp.MustCompile(`^accounts/([^/]+)$`),
				Variable: true, VariableName: accountIdVariableName},
		},
	}
	value := "something-else"
	tokens := ic.reference(i, []string{"parent"}, value, cty.StringVal(value), &resource{})
	assert.Equal(t, `"something-else"`, renderTokens(t, tokens))
	assert.NotContains(t, ic.variables, accountIdVariableName)
}

// TestGenerateVariablesWithDefault verifies a registered default is emitted into
// the generated variable block.
func TestGenerateVariablesWithDefault(t *testing.T) {
	tmpDir := t.TempDir()
	ic := &importContext{
		Directory:        tmpDir,
		variables:        map[string]string{accountIdVariableName: accountIdVariableDescription},
		variableDefaults: map[string]string{accountIdVariableName: "abc-123"},
	}
	require.NoError(t, ic.generateVariables())
	content, err := os.ReadFile(tmpDir + "/vars.tf")
	require.NoError(t, err)
	s := string(content)
	assert.Contains(t, s, `variable "databricks_account_id"`)
	assert.Contains(t, s, `default`)
	assert.Contains(t, s, `abc-123`)
}

// TestGenerateFieldWithRegexpVariableReferencePreservesValue is a regression
// test: a regexp `Variable` reference on a field must not cause the field value
// to be overwritten with the sanitized resource name during field extraction
// (as the legacy whole-field variable path does for sensitive fields). Only the
// matched substring should be substituted, keeping the rest of the value intact.
func TestGenerateFieldWithRegexpVariableReferencePreservesValue(t *testing.T) {
	ic := importContextForTest()
	value := "accounts/0d26daa6-5e44-4c97-a497-ef015f91254a/budgetPolicies/0f254aa7-4f27-3b88-a489-1b77a27e7371/ruleSets/default"
	r := &resource{
		Resource: "databricks_access_control_rule_set",
		ID:       value,
		Name:     "rs_test",
		Data: ic.Resources["databricks_access_control_rule_set"].Data(&terraform.InstanceState{
			ID:         value,
			Attributes: map[string]string{"name": value},
		}),
	}
	f := hclwrite.NewEmptyFile()
	block := f.Body().AppendNewBlock("resource", []string{r.Resource, r.Name})
	require.NoError(t, ic.unifiedDataToHcl(ic.Importables[r.Resource], []string{}, r, block.Body()))
	rendered := string(hclwrite.Format(f.Bytes()))
	assert.Contains(t, rendered,
		`name = "accounts/${var.databricks_account_id}/budgetPolicies/0f254aa7-4f27-3b88-a489-1b77a27e7371/ruleSets/default"`)
	assert.Equal(t, "0d26daa6-5e44-4c97-a497-ef015f91254a", ic.variableDefaults[accountIdVariableName])
}

// TestGenerateWholeFieldVariableReference verifies that a whole-field `Variable`
// reference with an explicit VariableName replaces the attribute value with a
// reference to the shared variable (e.g. `account_id = var.databricks_account_id`)
// and registers that variable with the value as its default, without rewriting
// the value to the resource name.
func TestGenerateWholeFieldVariableReference(t *testing.T) {
	ic := importContextForTest()
	acct := "0d26daa6-5e44-4c97-a497-ef015f91254a"
	r := &resource{
		Resource: "databricks_mws_credentials",
		ID:       acct + "/cred-1",
		Name:     "cred_test",
		Data: ic.Resources["databricks_mws_credentials"].Data(&terraform.InstanceState{
			ID: acct + "/cred-1",
			Attributes: map[string]string{
				"account_id":       acct,
				"credentials_name": "my-cred",
				"role_arn":         "arn:aws:iam::123456789012:role/x",
			},
		}),
	}
	f := hclwrite.NewEmptyFile()
	block := f.Body().AppendNewBlock("resource", []string{r.Resource, r.Name})
	require.NoError(t, ic.unifiedDataToHcl(ic.Importables[r.Resource], []string{}, r, block.Body()))
	rendered := string(hclwrite.Format(f.Bytes()))
	assert.Contains(t, rendered, "account_id       = var.databricks_account_id")
	assert.NotContains(t, rendered, acct)
	assert.Equal(t, acct, ic.variableDefaults[accountIdVariableName])
}

// TestReadListFromDataResolvesIntReferences is a regression test: elements of an
// integer list (e.g. workspace IDs) must be resolved to references to other
// resources when a matching Depends exists, not emitted as bare literals.
// Unmatched values still render as plain numbers.
func TestReadListFromDataResolvesIntReferences(t *testing.T) {
	ic := importContextForTest()
	ic.State.Append(resourceApproximation{
		Type: "databricks_mws_workspaces",
		Name: "ws_test",
		Instances: []instanceApproximation{
			{Attributes: map[string]any{"workspace_id": "123"}},
		},
	})
	imp := importable{Depends: []reference{
		{Path: "binding_workspace_ids", Resource: "databricks_mws_workspaces", Match: "workspace_id"},
	}}
	f := hclwrite.NewEmptyFile()
	sch := &schema.Schema{Type: schema.TypeList, Elem: &schema.Schema{Type: schema.TypeInt}}
	require.NoError(t, ic.readListFromData(imp, []string{"binding_workspace_ids"}, &resource{},
		[]any{123, 999}, f.Body(), sch, strconv.Itoa))
	assert.Contains(t, string(hclwrite.Format(f.Bytes())),
		"binding_workspace_ids = [databricks_mws_workspaces.ws_test.workspace_id, 999]")
}

// TestPluginFrameworkIntListResolvesReferences is the Plugin Framework counterpart:
// integer list elements in a PF resource must also resolve to references.
func TestPluginFrameworkIntListResolvesReferences(t *testing.T) {
	ic := importContextForTest()
	ic.State.Append(resourceApproximation{
		Type: "databricks_mws_workspaces",
		Name: "ws_test",
		Instances: []instanceApproximation{
			{Attributes: map[string]any{"workspace_id": "123"}},
		},
	})
	imp := ic.Importables["databricks_budget_policy"]
	sw := &PluginFrameworkSchemaWrapper{schema: ic.PluginFrameworkSchemas["databricks_budget_policy"]}
	fieldSchema := sw.GetField("binding_workspace_ids")
	require.NotNil(t, fieldSchema)
	f := hclwrite.NewEmptyFile()
	require.NoError(t, ic.pluginFrameworkFieldToHcl(imp, []string{}, "binding_workspace_ids", fieldSchema,
		[]interface{}{int64(123), int64(999)}, &resource{}, f.Body(), nil))
	assert.Contains(t, string(hclwrite.Format(f.Bytes())),
		"binding_workspace_ids = [databricks_mws_workspaces.ws_test.workspace_id, 999]")
}

// TestCombineReferencesRuleSetName verifies that when a field matches both the
// ContinueMatch account-id reference and a more-specific resource reference, both
// substitutions are combined into a single value.
func TestCombineReferencesRuleSetName(t *testing.T) {
	ic := importContextForTest()
	acct := "0d26daa6-5e44-4c97-a497-ef015f91254a"
	policy := "0f254aa7-4f27-3b88-a489-1b77a27e7371"
	ic.State.Append(resourceApproximation{
		Type: "databricks_budget_policy",
		Name: "pol_test",
		Instances: []instanceApproximation{
			{Attributes: map[string]any{"policy_id": policy}},
		},
	})
	value := "accounts/" + acct + "/budgetPolicies/" + policy + "/ruleSets/default"
	r := &resource{
		Resource: "databricks_access_control_rule_set", ID: value, Name: "rs_test",
		Data: ic.Resources["databricks_access_control_rule_set"].Data(&terraform.InstanceState{
			ID: value, Attributes: map[string]string{"name": value},
		}),
	}
	f := hclwrite.NewEmptyFile()
	block := f.Body().AppendNewBlock("resource", []string{r.Resource, r.Name})
	require.NoError(t, ic.unifiedDataToHcl(ic.Importables[r.Resource], []string{}, r, block.Body()))
	assert.Contains(t, string(hclwrite.Format(f.Bytes())),
		`name = "accounts/${var.databricks_account_id}/budgetPolicies/${databricks_budget_policy.pol_test.policy_id}/ruleSets/default"`)
}

// TestCombineReferencesChain verifies that a chain of ContinueMatch references
// (d1, d2 with ContinueMatch, terminal d3) combines all their substitutions,
// ordered by position in the value.
func TestCombineReferencesChain(t *testing.T) {
	ic := importContextForTest()
	value := "a/XX/b/YY/c/ZZ"
	imp := importable{Depends: []reference{
		{Path: "f", MatchType: MatchRegexp, Regexp: regexp.MustCompile(`^a/([^/]+)/.*$`), Variable: true, VariableName: "v1", ContinueMatch: true},
		{Path: "f", MatchType: MatchRegexp, Regexp: regexp.MustCompile(`^a/[^/]+/b/([^/]+)/.*$`), Variable: true, VariableName: "v2", ContinueMatch: true},
		{Path: "f", MatchType: MatchRegexp, Regexp: regexp.MustCompile(`^a/[^/]+/b/[^/]+/c/([^/]+)$`), Variable: true, VariableName: "v3"},
	}}
	toks := ic.reference(imp, []string{"f"}, value, cty.StringVal(value), &resource{})
	assert.Equal(t, `"a/${var.v1}/b/${var.v2}/c/${var.v3}"`, renderTokens(t, toks))
}

func renderTokens(t *testing.T, tokens hclwrite.Tokens) string {
	t.Helper()
	f := hclwrite.NewEmptyFile()
	f.Body().SetAttributeRaw("v", tokens)
	src := string(hclwrite.Format(f.Bytes()))
	// Extract the value after `v = `.
	_, rhs, ok := strings.Cut(src, "= ")
	require.True(t, ok, "unexpected rendered attribute: %q", src)
	return strings.TrimSpace(rhs)
}
