package httplog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type bodyLogger struct {
	debugTruncateBytes int
	maxBytes           int
}

var redactKeys = map[string]bool{
	"string_value":  true,
	"token_value":   true,
	"content":       true,
	"access_token":  true,
	"refresh_token": true,
	"id_token":      true,
	"token":         true,
	"password":      true,
}

func (b bodyLogger) mask(m map[string]any) {
	for k := range m {
		if _, ok := redactKeys[k]; ok {
			m[k] = "**REDACTED**"
		}
	}
}

// mapKeys returns sorted slice of keys in the map.
// Used for deterministic iteration over maps.
func mapKeys(m map[string]any) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (b bodyLogger) recursiveMarshalMap(m map[string]any, budget int) (json.RawMessage, error) {
	var out = make(map[string]any)

	// Redact sensitive values.
	b.mask(m)

	// Each key in the map appears in the output, regardless of character budget.
	for _, k := range mapKeys(m) {
		raw, err := b.recursiveMarshal(m[k], budget)
		if err != nil {
			return nil, err
		}
		out[k] = raw
		budget -= len(raw)
	}

	return json.Marshal(out)
}

func (b bodyLogger) recursiveMarshalSlice(s []any, budget int) (json.RawMessage, error) {
	var out []any

	// The first element of a slice appears in the output, regardless of character budget.
	// Subsequent elements are included if the budget allows.
	for i := range s {
		// If we're out of character budget, include trailer.
		if i > 0 && budget <= 0 {
			out = append(out, fmt.Sprintf("... (%d additional elements)", len(s)-len(out)))
			break
		}

		raw, err := b.recursiveMarshal(s[i], budget)
		if err != nil {
			return nil, err
		}
		out = append(out, raw)
		budget -= len(raw)
	}

	return json.Marshal(out)
}

func (b bodyLogger) recursiveMarshal(v any, budget int) (json.RawMessage, error) {
	if m, ok := v.(map[string]any); ok {
		return b.recursiveMarshalMap(m, budget)
	}

	if s, ok := v.([]any); ok {
		return b.recursiveMarshalSlice(s, budget)
	}

	// Truncate strings to max length.
	if s, ok := v.(string); ok {
		return json.Marshal(onlyNBytes(s, b.debugTruncateBytes))
	}

	// Other types can be marshalled as is.
	return json.Marshal(v)
}

func (b bodyLogger) redactedDump(prefix string, body []byte) string {
	if len(body) == 0 {
		return ""
	}

	// Unmarshal body into primitive types.
	var tmp any
	err := json.Unmarshal(body, &tmp)
	if err != nil {
		// Unable to unmarshal means the body isn't JSON. Split on newlines and
		// truncate to maxBytes.
		truncated := onlyNBytes(string(body), b.debugTruncateBytes)
		splitByNewlines := strings.Split(truncated, "\n")
		sb := strings.Builder{}
		for i, line := range splitByNewlines {
			if i != 0 {
				sb.WriteString("\n")
			}
			line = strings.Trim(line, "\r")
			sb.WriteString(fmt.Sprintf("%s%s", prefix, line))
		}
		return sb.String()
	}

	maxBytes := b.maxBytes
	if maxBytes == 0 {
		maxBytes = 1024
	}
	if b.debugTruncateBytes > maxBytes {
		maxBytes = b.debugTruncateBytes
	}

	// Re-marshal body taking redaction and character limit into account.
	raw, err := b.recursiveMarshal(tmp, maxBytes)
	if err != nil {
		return fmt.Sprintf("[unable to marshal: %s]", err.Error())
	}

	var buf bytes.Buffer
	buf.WriteString(prefix)
	err = json.Indent(&buf, raw, prefix, "  ")
	if err != nil {
		return fmt.Sprintf("[unable to indent: %s]", err.Error())
	}

	return buf.String()
}

func onlyNBytes(j string, numBytes int) string {
	diff := len([]byte(j)) - numBytes
	if diff > 0 {
		return fmt.Sprintf("%s... (%d more bytes)", j[:numBytes], diff)
	}
	return j
}
