package common

import (
	"encoding/json"
	"strings"
)

// includeNode is a node in the tree of include fields.
type includeNode struct {
	// children are the fields that are nested under this node
	children map[string]*includeNode
	// exact is true if the field is an exact match
	exact bool
}

// parseIncludeFields parses the includeFields parameter into a tree of includeNode objects.
// Entries can use dot notation to specify a nested fields.
func parseIncludeFields(includeFields []string) includeNode {
	var root includeNode
	for _, field := range includeFields {
		parts := strings.Split(field, ".")
		current := &root
		for i, part := range parts {
			if current.children == nil {
				current.children = make(map[string]*includeNode)
			}
			if _, ok := current.children[part]; !ok {
				current.children[part] = &includeNode{}
			}
			current = current.children[part]
			if i == len(parts)-1 {
				current.exact = true
			}
		}
	}
	return root
}

// filterMap recursively filters a map[string]any to only include allowed fields.
func (n includeNode) filterMap(m map[string]any) map[string]any {
	result := make(map[string]any)
	for k, v := range n.children {
		// If the field is not in the map, skip.
		val, ok := m[k]
		if !ok {
			continue
		}
		// If the field is an exact match, copy it.
		if v.exact {
			result[k] = val
			continue
		}
		// If the field is a nested map, filter it.
		if submap, ok := val.(map[string]any); ok {
			filtered := v.filterMap(submap)
			if len(filtered) > 0 {
				result[k] = filtered
			} else {
				// If the source was non-nil, preserve as empty map to keep non-nil struct
				result[k] = map[string]any{}
			}
			continue
		}
	}
	return result
}

// CopyViaJSON copies the fields of the source struct to the destination struct.
// The fields to copy are specified by the includeFields parameter.
// Each entry in the includeFields parameter is a JSON path to a field in the source struct.
// It supports nested fields and maps via dot notation.
func CopyViaJSON[T any](src T, includeFields []string) T {
	var dst T
	b1, err := json.Marshal(src)
	if err != nil {
		return dst
	}
	var m map[string]any
	if err := json.Unmarshal(b1, &m); err != nil {
		return dst
	}
	filtered := parseIncludeFields(includeFields).filterMap(m)
	b2, err := json.Marshal(filtered)
	if err != nil {
		return dst
	}
	_ = json.Unmarshal(b2, &dst)
	return dst
}
