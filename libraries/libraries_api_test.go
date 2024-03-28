package libraries

import (
	"testing"
)

func TestNewLibraryFromInstanceState(t *testing.T) {
	tests := []struct {
		want string
		give any
	}{
		{"jar:a", map[string]any{"jar": "a"}},
		{"egg:b", map[string]any{"egg": "b"}},
		{"whl:c", map[string]any{"whl": "c"}},
		{"pypi:d", map[string]any{"pypi": []any{
			map[string]any{"package": "d"},
		}}},
		{"mvn:e", map[string]any{"maven": []any{
			map[string]any{"coordinates": "e"},
		}}},
		{"cran:f", map[string]any{"cran": []any{
			map[string]any{"package": "f"},
		}}},
		{"unknown", map[string]any{"bottle": "g"}},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := NewLibraryFromInstanceState(tt.give); got.String() != tt.want {
				t.Errorf("NewLibraryFromInstanceState() = %v, want %v", got, tt.want)
			}
		})
	}
}
