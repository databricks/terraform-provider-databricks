package httplog

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"
)

func TestBodyLoggerCommon(t *testing.T) {
	var body = map[string]any{
		"k1": map[string]any{
			"a": 1,
			"b": 2,
			"c": []string{
				"el1",
				"el2",
			},
		},
		"k2": []string{
			"el1",
			"el2",
			"el3",
			"el4",
		},
		"k3": "this is a very long string",
	}

	buf, err := json.Marshal(body)
	require.NoError(t, err)

	dump := bodyLogger{
		debugTruncateBytes: 4,
		maxBytes:           16,
	}.redactedDump("", buf)

	var out any
	err = json.Unmarshal([]byte(dump), &out)
	require.NoError(t, err)

	// Top level map contains all keys
	keys := maps.Keys(out.(map[string]any))
	assert.Contains(t, keys, "k1")
	assert.Contains(t, keys, "k2")
	assert.Contains(t, keys, "k3")

	// Array under k1 is included in whole
	k1c := out.(map[string]any)["k1"].(map[string]any)["c"].([]any)
	assert.Equal(t, k1c, []any{"el1", "el2"})

	// Array under k2 is not included in whole
	k2 := out.(map[string]any)["k2"].([]any)
	assert.Equal(t, k2, []any{"el1", "... (3 additional elements)"})

	// String under k3 is truncated
	k3 := out.(map[string]any)["k3"].(string)
	assert.Equal(t, k3, "this... (22 more bytes)")
}

func TestBodyLoggerEmpty(t *testing.T) {
	dump := bodyLogger{}.redactedDump("", []byte(""))
	assert.Empty(t, dump)
}

func TestBodyLoggerNotJson(t *testing.T) {
	dump := bodyLogger{debugTruncateBytes: 10}.redactedDump("", []byte("<html>"))
	assert.Equal(t, dump, "<html>")
}

func TestBodyLoggerNonJSONNewline(t *testing.T) {
	dump := bodyLogger{debugTruncateBytes: 100}.redactedDump("< ", []byte(`<html>
	<body>
		Hello world!
	</body>
</html>`))
	assert.Equal(t, dump, `< <html>
< 	<body>
< 		Hello world!
< 	</body>
< </html>`)
}
