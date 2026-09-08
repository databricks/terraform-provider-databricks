package providers

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Verifies the muxed provider serves an identical schema from both underlying
// providers, including the new user_agent_extra attribute.
func TestMuxedProviderSchemaIncludesUserAgentExtra(t *testing.T) {
	ctx := context.Background()
	server, err := GetProviderServer(ctx)
	require.NoError(t, err)
	resp, err := server.GetProviderSchema(ctx, &tfprotov6.GetProviderSchemaRequest{})
	require.NoError(t, err)
	for _, d := range resp.Diagnostics {
		assert.NotEqual(t, tfprotov6.DiagnosticSeverityError, d.Severity, d.Summary+": "+d.Detail)
	}
	found := false
	for _, attr := range resp.Provider.Block.Attributes {
		if attr.Name == "user_agent_extra" {
			found = true
		}
	}
	assert.True(t, found, "user_agent_extra should be in the muxed provider schema")
}
