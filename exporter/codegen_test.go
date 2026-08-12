package exporter

import (
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
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
