package permissions

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/stretchr/testify/assert"
)

// fakeAttributeGetter is a minimal attributeGetter backed by a map; absent keys
// return "" so KeyValue skips them.
type fakeAttributeGetter map[string]any

func (f fakeAttributeGetter) Get(key string) any {
	if v, ok := f[key]; ok {
		return v
	}
	return ""
}

func TestNormalizeSecurableNameStripsAiGatewayPrefix(t *testing.T) {
	for field, prefix := range securableResourceNamePrefixes {
		// The databricks_ai_gateway_*.name attribute carries the prefix.
		assert.Equal(t, "main.default.x", normalizeSecurableName(field, prefix+"main.default.x"))
		// A bare full name is left unchanged.
		assert.Equal(t, "main.default.x", normalizeSecurableName(field, "main.default.x"))
	}
	// Non-AI-Gateway securables are never touched, even if the value happens to
	// look prefixed.
	assert.Equal(t, "model-provider-services/x",
		normalizeSecurableName("catalog", "model-provider-services/x"))
	assert.Equal(t, "main.default.tbl", normalizeSecurableName("table", "main.default.tbl"))
}

func TestSuppressResourceNamePrefixDiff(t *testing.T) {
	f := SuppressResourceNamePrefixDiff("model_provider_service")
	assert.NotNil(t, f)
	// config (prefixed, from .name) vs state (bare, from read) => suppressed.
	assert.True(t, f("model_provider_service", "main.default.x",
		"model-provider-services/main.default.x", nil))
	// genuinely different securables are not suppressed.
	assert.False(t, f("model_provider_service", "main.default.x",
		"model-provider-services/main.default.y", nil))
	// non-AI-Gateway securables get no suppressor.
	assert.Nil(t, SuppressResourceNamePrefixDiff("catalog"))
}

func TestSecurableMappingKeyValueStripsPrefix(t *testing.T) {
	sm := SecurableMapping{"model_provider_service": catalog.SecurableType("model_provider_service")}
	d := fakeAttributeGetter{
		"model_provider_service": "model-provider-services/main.default.openai_prod",
	}
	field, name := sm.KeyValue(d)
	assert.Equal(t, "model_provider_service", field)
	// The prefix is stripped so the permissions API receives the bare full name,
	// and the resource ID does not gain an extra "/" segment.
	assert.Equal(t, "main.default.openai_prod", name)
}
