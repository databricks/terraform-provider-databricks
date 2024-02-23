package exporter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourceApproaximationGet(t *testing.T) {
	_, found := (&resourceApproximation{}).Get("test")
	assert.False(t, found)

	v, found := (&resourceApproximation{
		Instances: []instanceApproximation{
			{Attributes: map[string]any{"test": "42"}},
		},
	}).Get("test")
	require.True(t, found)
	assert.Equal(t, "42", v.(string))
}
